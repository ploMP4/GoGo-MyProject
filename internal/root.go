package internal

import (
	"errors"
	"fmt"
	"os"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// Application Version
const APPLICATION_VERSION = "1.0.0"

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

type Commands []string
type SubCommands map[string]SubCommand
type Files map[string]File
type Placeholders map[string]string

type App struct {
	gadgetName string           // Name of the gadget we are executing
	parser     Parser           // Parser
	spinner    *spinner.Spinner // Load Spinner
	appname    string
	verbose    bool
}

var app *App

func Execute() {
	var message string

	gadgetName, args, err := validateInput()
	s := loadSpinner()

	app = &App{
		gadgetName: gadgetName,
		spinner:    s,
		parser: Parser{
			args: args,
		},
		appname: "",
		verbose: false,
	}

	if err != nil {
		exitGracefully(err)
	}

	switch gadgetName {
	case SHORT_HELP_FLAG, HELP_FLAG:
		if len(args) == 0 {
			showHelp()
		} else {
			showSubHelp(args[0])
		}

	case SHORT_VERSION_FLAG, VERSION_FLAG:
		color.Green("Application version: " + APPLICATION_VERSION)

	case SHORT_SET_GADGET_PATH_FLAG, SET_GADGET_PATH_FLAG:
		if len(args) == 0 {
			showHelp()
			break
		}

		if err = app.parser.parseSettings(); err != nil {
			exitGracefully(err)
		}
		if err = app.parser.settings.setGadgetPath(args[0]); err == nil {
			message = fmt.Sprint("Config path set to: " + args[0])
		}

	case SHORT_SET_TEMPATE_PATH_FLAG, SET_TEMPATE_PATH_FLAG:
		if len(args) == 0 {
			showHelp()
			break
		}

		if err = app.parser.parseSettings(); err != nil {
			exitGracefully(err)
		}
		if err = app.parser.settings.setTemplatePath(args[0]); err == nil {
			message = fmt.Sprint("Template path set to: " + args[0])
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
	var gadget string
	var args []string

	if len(os.Args) > 1 {
		gadget = os.Args[1]

		if len(os.Args) >= 2 {
			args = os.Args[2:]
		}
	} else {
		showHelp()
		return "", nil, errors.New("no gadget provided")
	}

	return gadget, args, nil
}

func (app *App) run() (string, error) {
	err := app.parser.parseSettings()
	if err != nil {
		return "", err
	}

	gadget, err := app.parser.parseGadget(app.gadgetName)
	if err != nil {
		return "", err
	}

	gadget, dirs, verbose, appname := app.parser.parseArgs(gadget)
	app.appname = appname
	app.verbose = verbose

	app.spinner.Start()

	msg, err := gadget.runCommands()
	if err != nil {
		return msg, err
	}

	app.spinner.Restart()

	if gadget.Chdir && app.appname != "" {
		err := os.Chdir(app.appname)
		if err != nil {
			showMessage("Warning", err.Error())
		}
	}

	app.createDirs(dirs)

	app.spinner.Restart()
	gadget.handleFiles()

	gadget.runSubCommands()

	app.spinner.Stop()
	return "\nGadget Executed Successfully: " + app.gadgetName, nil
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
