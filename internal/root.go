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

type App struct {
	filename string           // Name of the gadget we are executing
	appName  string           // The name that the main app folder will have
	parser   Parser           // Parser
	spinner  *spinner.Spinner // Load Spinner
	verbose  bool             // Verbose output flag
}

func Execute() {
	var message string
	filename, appName, args, err := validateInput()
	s := loadSpinner()

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
func validateInput() (string, string, []string, error) {
	var command, appName string
	var args []string

	if len(os.Args) > 1 {
		command = os.Args[1]

		if len(os.Args) >= 3 {
			appName = os.Args[2]
		}

		if len(os.Args) >= 4 {
			args = os.Args[3:]
		}
	} else {
		showHelp()
		return "", "", nil, errors.New("no command provided")
	}

	return command, appName, args, nil
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

	if app.appName == "" {
		return "", errors.New("appname was not provided")
	}

	mainCommands, otherCommands, dirs, verbose := app.parser.parseArgs()
	mainCommands[len(mainCommands)-1] = fmt.Sprintf(
		"%s %s",
		mainCommands[len(mainCommands)-1],
		app.appName,
	)
	app.verbose = verbose

	app.spinner.Start()

	msg, err := app.runMainCommands(mainCommands)
	if err != nil {
		return msg, err
	}

	app.spinner.Restart()
	showMessage("Created Project", app.appName)

	os.Chdir(app.appName)

	app.createDirs(dirs)
	app.runSubCommands(otherCommands)

	app.spinner.Stop()
	return "\nApp Created Successfully: " + app.appName, nil
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
		if strings.Contains(file.Filepath, "<APPNAME>") {
			path := strings.Split(file.Filepath, "<APPNAME>")
			file.Filepath = app.appName + path[1]
		}

		showMessage("Adding", name, "in", Green(file.Filepath))
		editFile(file.Filepath, file.Change.SplitOn, file.Change.Append)
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
