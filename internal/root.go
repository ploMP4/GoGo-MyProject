package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// Application Version
const APPLICATION_VERSION = "4.4.0"

const (
	SHORT_ALL_FLAG = "a"
	ALL_FLAG       = "all"

	SHORT_EXCLUDE_FLAG = "e"
	EXLCUDE_FLAG       = "exclude"

	SHORT_VERBOSE_FLAG = "vv"
	VERBOSE_FLAG       = "verbose"

	SHORT_VERSION_FLAG = "v"
	VERSION_FLAG       = "version"

	SHORT_HELP_FLAG = "h"
	HELP_FLAG       = "help"

	SHORT_SET_GADGET_PATH_FLAG = "G"
	SET_GADGET_PATH_FLAG       = "gadgetdir"

	SHORT_SET_TEMPATE_PATH_FLAG = "T"
	SET_TEMPATE_PATH_FLAG       = "templatedir"
)

const (
	PLACEHOLDER_FILENAME = "_FILENAME"
	PLACEHOLDER_APPNAME  = "_APPNAME"
)

type App struct {
	filename string           // Name of the gadget we are executing
	appName  string           // The name that the main app folder will have
	parser   Parser           // Parser
	spinner  *spinner.Spinner // Load Spinner
	verbose  bool             // Verbose output flag
}

func Execute() {
	var message string
	var appName string
	filename, args, err := validateInput()
	s := loadSpinner()
	if len(args) >= 1 {
		appName = args[0]
	}

	app := &App{
		filename: filename,
		appName:  appName,
		spinner:  s,
		parser: Parser{
			args: args,
		},
		verbose: false,
	}

	if err != nil {
		exitGracefully(err)
	}

	switch filename {
	case SHORT_HELP_FLAG, HELP_FLAG:
		if appName == "" {
			showHelp()
		} else {
			showSubHelp(appName)
		}

	case SHORT_VERSION_FLAG, VERSION_FLAG:
		color.Green("Application version: " + APPLICATION_VERSION)

	case SHORT_SET_GADGET_PATH_FLAG, SET_GADGET_PATH_FLAG:
		app.parser.parseSettings()
		err = app.parser.settings.setGadgetPath(appName)
		if err == nil {
			message = fmt.Sprint("Config path set to: " + appName)
		}

	case SHORT_SET_TEMPATE_PATH_FLAG, SET_TEMPATE_PATH_FLAG:
		app.parser.parseSettings()
		err = app.parser.settings.setTemplatePath(appName)
		if err == nil {
			message = fmt.Sprint("Template path set to: " + appName)
		}

	default:
		message, err = app.run()
	}

	s.Stop()
	exitGracefully(err, message)
}

// Validate that the user passed what command he wants
// to execute and also return with it the appname and
// the rest of the args for later use
func validateInput() (string, []string, error) {
	var command string
	var args []string

	if len(os.Args) > 1 {
		command = os.Args[1]

		if len(os.Args) >= 3 {
			args = os.Args[2:]
		}
	} else {
		showHelp()
		return "", nil, errors.New("no gadget provided")
	}

	return command, args, nil
}

func (app *App) run() (string, error) {
	err := app.parser.parseSettings()
	if err != nil {
		return "", err
	}

	err = app.parser.parseGadget(app.filename)
	if err != nil {
		return "", err
	}

	if app.appName == "" && !app.parser.gadget.Template {
		return "", errors.New("appname was not provided")
	}

	mainCommands, subCommands, dirs, verbose := app.parser.parseArgs()
	app.verbose = verbose

	app.spinner.Start()

	if !app.parser.gadget.Template {
		mainCommands[len(mainCommands)-1] = fmt.Sprintf(
			"%s %s",
			mainCommands[len(mainCommands)-1],
			app.appName,
		)

		msg, err := app.runMainCommands(mainCommands)
		if err != nil {
			return msg, err
		}

		os.Chdir(app.appName)
		showMessage("Created Project", app.appName)
	}

	app.spinner.Restart()

	if app.parser.gadget.Files != nil {
		for name, file := range app.parser.gadget.Files {
			app.handeMainCommandFiles(name, file)
		}
	}
	app.createDirs(dirs)
	app.runSubCommands(subCommands)

	app.spinner.Stop()
	return "\nGadget Executed Successfully: " + app.filename, nil
}

// Used to run all the main commands and throw an error if
// something goes wrong
func (app *App) runMainCommands(mainCommands MainCommmands) (string, error) {
	for _, command := range mainCommands {
		app.spinner.Restart()
		showMessage("Running", command)

		cmd := strings.Fields(command)
		c := exec.Command(cmd[0], cmd[1:]...)

		if app.verbose {
			app.spinner.Stop()
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
		}

		err := c.Run()
		if err != nil {
			return "Unable to execute command: " + cmd[0], err
		}
	}

	return "", nil
}

func (app *App) handeMainCommandFiles(name string, file File) {
	app.spinner.Restart()

	if file.Template {
		showMessage("Copying", file.Filepath)

		templatePath := fmt.Sprintf(
			"./%s/templates/%s/%s",
			PROJECT_ROOT_DIR_NAME,
			app.filename,
			file.Filepath,
		)

		if !fileExists(templatePath) {
			templatePath = fmt.Sprintf(
				"%s/%s/%s",
				app.parser.settings.TemplatePath,
				app.filename,
				file.Filepath,
			)
		}

		copyFileFromTemplate(templatePath, file.Filepath)

		if file.Change.Placeholder != nil {
			handlePlaceholders(file.Filepath, file.Change.Placeholder, app.parser.args)
		}
	} else {
		if strings.Contains(file.Filepath, PLACEHOLDER_APPNAME) {
			path := strings.Split(file.Filepath, PLACEHOLDER_APPNAME)
			file.Filepath = app.appName + path[1]
		}

		showMessage("Adding", name, "in", Green(file.Filepath))
		editFile(file.Filepath, file.Change.SplitOn, file.Change.Append)
	}

	for idx, arg := range app.parser.args {
		if arg == PLACEHOLDER_FILENAME {
			filepathSlice := strings.Split(file.Filepath, "/")
			filepathSlice[len(filepathSlice)-1] = app.parser.args[idx+1]
			filepathNew := strings.Join(filepathSlice, "/")
			os.Rename(file.Filepath, filepathNew)
		}
	}
}

func (app *App) createDirs(dirs []string) {
	for _, dir := range dirs {
		app.spinner.Restart()
		showMessage("Creating", dir)

		err := os.Mkdir(dir, 0755)
		if err != nil {
			color.Red("Error: %v\n", err)
		}
	}
}

// Used to run all the subcommands either concurrently or by themselves
// based on the value of SubCommand.parallel. Displays a message if
// there is an error
func (app *App) runSubCommands(subcommands []SubCommand) {
	var wg sync.WaitGroup

	for _, command := range subcommands {
		app.spinner.Restart()
		showMessage("Running", command.Name)

		if command.Parallel {
			wg.Add(1)
			go func(command SubCommand) {
				defer wg.Done()

				err := app.executeSubCommand(command)
				if err != nil {
					color.Yellow("Failed to execute %s\n", command)
					color.Red("Error: %v\n", err)
				}
			}(command)
		} else {
			err := app.executeSubCommand(command)
			if err != nil {
				color.Yellow("Failed to execute %s\n", command)
				color.Red("Error: %v\n", err)
			}
		}
	}

	wg.Wait()
}

// Executes a single subcommand
func (app *App) executeSubCommand(command SubCommand) error {
	if command.Command != "" {
		err := app.handleSubCommandCommand(command.Command)

		if err != nil {
			return err
		}
	}

	if command.Files != nil {
		for name, file := range command.Files {
			app.handleSubCommandFiles(command.Name, name, file)
		}
	}

	return nil
}

func (app *App) handleSubCommandCommand(command string) error {
	cmd := strings.Fields(command)
	c := exec.Command(cmd[0], cmd[1:]...)

	if app.verbose {
		app.spinner.Stop()
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}

func (app *App) handleSubCommandFiles(commandName, name string, file File) {
	app.spinner.Restart()

	if file.Template {
		showMessage("Copying", file.Filepath)

		templatePath := fmt.Sprintf(
			"../%s/templates/%s/%s/%s",
			PROJECT_ROOT_DIR_NAME,
			app.filename,
			commandName,
			file.Filepath,
		)

		if !fileExists(templatePath) {
			templatePath = fmt.Sprintf(
				"%s/%s/%s/%s",
				app.parser.settings.TemplatePath,
				app.filename,
				commandName,
				file.Filepath,
			)
		}

		copyFileFromTemplate(templatePath, file.Filepath)

		if file.Change.Placeholder != nil {
			handlePlaceholders(file.Filepath, file.Change.Placeholder, app.parser.args)
		}
	} else {
		if strings.Contains(file.Filepath, PLACEHOLDER_APPNAME) {
			path := strings.Split(file.Filepath, PLACEHOLDER_APPNAME)
			file.Filepath = app.appName + path[1]
		}

		showMessage("Adding", name, "in", Green(file.Filepath))
		editFile(file.Filepath, file.Change.SplitOn, file.Change.Append)
	}

	for idx, arg := range app.parser.args {
		if arg == PLACEHOLDER_FILENAME {
			filepathSlice := strings.Split(file.Filepath, "/")
			filepathSlice[len(filepathSlice)-1] = app.parser.args[idx+1]
			filepathNew := strings.Join(filepathSlice, "/")
			os.Rename(file.Filepath, filepathNew)
		}
	}
}

func handlePlaceholders(filepath string, placeholders Placeholders, args []string) {
	for placeholder, defaultValue := range placeholders {
		found := findAndReplacePlaceholder(filepath, placeholder, args)
		if !found {
			replacePlaceholder(filepath, placeholder, defaultValue)
		}
	}
}

func findAndReplacePlaceholder(filepath, placeholder string, args []string) bool {
	for idx, arg := range args {
		if arg == placeholder {
			replacePlaceholder(filepath, placeholder, args[idx+1])

			return true
		}
	}

	return false
}
