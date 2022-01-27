package main

import (
	"errors"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		color.Red("Screen clearing is unsupported on your platform")
	}
}

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
	
	help                                                                                   -show the help menu
	version                                                                                -print application version
	cpp <appname> || c++ <appname>                                                         -create c++ application 
	react <appname> [ts | typescript] [redux] [mui | material-ui | materialUI] [bootstrap] -create react app and optionally add typescript, redux, material-ui, bootstrap
	next <appname> [ts | typescript] [redux]                                               -crete nextJS app and optionally add typescript, redux  

	`)
}
