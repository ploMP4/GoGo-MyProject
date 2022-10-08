package pkg

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
const version = "2.0.0"

type App struct {
	filename string           // Name of the config file we are executing
	appName  string           // The name that the main app folder will have
	parser   Parser           // Parser
	spinner  *spinner.Spinner // Load Spinner
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
	}

	if err != nil {
		exitGracefully(err)
	}

	switch filename {
	case "h", "help":
		if appName == "" {
			showHelp()
		} else {
			showSubHelp(appName)
		}

	case "v", "version":
		color.Green("Application version: " + version)

	case "C", "set-config-path":
		app.parser.parseSettings()
		err = app.parser.settings.setConfigPath(appName)
		if err != nil {
			message = fmt.Sprint("Config path set to: " + appName)
		}

	case "T", "set-template-path":
		app.parser.parseSettings()
		err = app.parser.settings.setTemplatePath(appName)
		if err != nil {
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

	err = app.parser.parseConfig(app.filename)
	if err != nil {
		return "", err
	}

	if app.appName == "" {
		return "", errors.New("appname was not provided")
	}

	mainCommands, otherCommands, dirs := app.parser.parseArgs()
	mainCommands[len(mainCommands)-1] = append(mainCommands[len(mainCommands)-1], app.appName)

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
	for _, cmd := range mainCommands {
		app.spinner.Restart()
		showMessage("Running", cmd...)

		c := exec.Command(cmd[0], cmd[1:]...)
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
			go func(wg *sync.WaitGroup, command SubCommand) {
				defer wg.Done()

				err := app.executeSubCommand(command)
				if err != nil {
					color.Yellow("Failed to execute %s\n", command)
					color.Red("Error: %v\n", err)
				}
			}(&wg, command)
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
	if command.Command != nil {
		c := exec.Command(command.Command[0], command.Command[1:]...)
		err := c.Run()
		if err != nil {
			return err
		}
	}

	if command.Files != nil {
		for name, file := range command.Files {
			if file.Template {
				app.spinner.Restart()
				showMessage("Copying", file.Filepath)

				templatePath := fmt.Sprintf("%s/%s/%s/%s", app.parser.settings.TemplatePath, app.filename, command.Name, file.Filepath)
				copyFileFromTemplate(templatePath, file.Filepath)
			} else {
				app.spinner.Restart()

				if strings.Contains(file.Filepath, "<APPNAME>") {
					path := strings.Split(file.Filepath, "<APPNAME>")
					file.Filepath = app.appName + path[1]
				}

				showMessage("Adding", name, "in", Green(file.Filepath))
				editFile(file.Filepath, file.Change.SplitOn, file.Change.Append)
			}
		}
	}

	return nil
}

// Adds a string in the specified file either at the end of the file
// if the splitOn argument is an empty string or by splitting the file by the string
// specified in splitOn and appending it there
func editFile(filename, splitOn, toAppend string) {
	var settings string

	content, err := os.ReadFile(filename)
	if err != nil {
		exitGracefully(err)
	}

	if splitOn != "" { // Append after certain string in the file
		s := strings.Split(string(content), splitOn)
		s[0] += splitOn + toAppend

		settings = strings.Join(s, " ")
	} else { // Append at the end of file
		settings = string(content) + toAppend
	}

	err = os.WriteFile(filename, []byte(settings), 0644)
	if err != nil {
		exitGracefully(err)
	}
}
