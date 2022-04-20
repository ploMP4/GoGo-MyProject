package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// Exits the program gracefully and
// displays a message and an error message
// if there are passed any.
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

// Print the help menu
func showHelp() {
	p := Parser{}
	p.parseSettings()
	helpCommands := p.getHelp()

	fmt.Printf(`A CLI tool to create starter boilerplate for you

GoGo is a CLI tool that creates the starter boilerplate 
for your projects and it's really helpfull for people
who use many different programming languages and frameworks.

%s

	gogo <COMMAND> <APPNAME> [args]

%s

	        help [command]   - show the help menu
	               version   - print application version
	set-config-path <path>   - set the config folder path containing your json files.
	%v

`, Yellow("USAGE:"), Yellow("AVAILABLE COMMANDS:"), strings.Trim(fmt.Sprint(helpCommands), "[]"))
}

// Used to display status messages
// e.x. Running: npm install
func showMessage(prefix string, message ...string) {
	fmt.Printf("%s: %s\n", Yellow(prefix), message)
}

func loadSpinner() *spinner.Spinner {
	return spinner.New(spinner.CharSets[43], 100*time.Millisecond, spinner.WithHiddenCursor(true)) // Build our new spinner
}
