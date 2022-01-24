package main

import (
	"errors"
	"os"

	"github.com/fatih/color"
)

func setup(command, appName string) {
	if command != "version" && command != "help" {
		path, err := os.Getwd()
		if err != nil {
			exitGracefully(err)
		}

		ggp.path = path
		ggp.appname = appName
		ggp.rootPath = path + "/" + appName
	}
}

func appNameExists(appName string) {
	if appName == "" {
		exitGracefully(errors.New("no application name provided"))
	}
}

func showHelp() {
	color.Yellow(`Available commands:
	
	help                           -show the help menu
	version                        -print application version
	cpp <appname> || c++ <appname> -create c++ application 

	`)
}
