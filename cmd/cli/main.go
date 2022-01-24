package main

import (
	"errors"
	"os"

	"github.com/fatih/color"
)

const version = "1.0.0"

type GGP struct {
	path     string
	rootPath string
	appname  string
	cpp      CppProjectType
	django   DjangoProjectType
}

var (
	ggp GGP
)

func main() {
	var message string
	command, appName, args, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	setup(command, appName)

	switch command {
	case "help":
		showHelp()

	case "version":
		color.Yellow("Application version: " + version)

	case "java":
		appNameExists(appName)

	case "cpp", "c++":
		appNameExists(appName)
		ggp.cpp.create(args)
		message = "C++ project created successfully under the name: " + ggp.appname

	case "django":
		appNameExists(appName)
		ggp.django.create()

	default:
		showHelp()
	}

	exitGracefully(nil, message)
}

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

func exitGracefully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}

	if err != nil {
		color.Red("Error: %v\n", err)
	}

	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished")
	}

	os.Exit(0)
}
