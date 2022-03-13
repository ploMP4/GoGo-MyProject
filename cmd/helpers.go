package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func exitGracefully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}

	if len(message) > 0 {
		color.Yellow(message)
	}

	if err != nil {
		color.Red("Error: %v\n", err)
		os.Exit(1)
	}

	color.Green("Finished")
	os.Exit(0)
}

func showHelp() {
	fmt.Printf(`%s

	gogo <COMMAND> <APPNAME> [args]

%s

	help     -show the help menu
	version  -print application version

`, Yellow("Usage:"), Yellow("Available commands:"))
}

func showMessage(prefix string, message ...string) {
	fmt.Printf("%s: %s\n", Yellow(prefix), message)
}

func loadSpinner() *spinner.Spinner {
	return spinner.New(spinner.CharSets[43], 100*time.Millisecond, spinner.WithHiddenCursor(true)) // Build our new spinner
}
