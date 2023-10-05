package internal

import (
	"fmt"
	"io/fs"
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
		app.spinner.Stop()
		color.Red("Error: %v\n", err)
		os.Exit(1)
	}

	app.spinner.Stop()
	color.Green("Finished")
	os.Exit(0)
}

// Print the general help menu
func showHelp() {
	p := Parser{}
	if err := p.parseSettings(); err != nil {
		exitGracefully(err)
	}
	helpCommands := p.getHelp()

	fmt.Fprintf(
		color.Output,
		`A CLI tool that allows you to easily blueprint boilerplate and repetitive commands. 

GoGo is a CLI tool that allows you to easily blueprint boilerplate 
and repetitive commands. From basic files to scaffolding 
the entire project structure.

%s

	gogo <COMMAND> <APPNAME> [args]

%s

	     h, help [command]   - show the help menu
	            v, version   - print application version
           G, gadgetdir <path>   - set the gadgets folder path containing your yaml files.
         T, templatedir <path>   - set the template folder path.
	%v

`,
		Yellow("USAGE:"),
		Yellow("AVAILABLE COMMANDS:"),
		strings.Trim(fmt.Sprint(helpCommands), "[]"),
	)
}

// Show the help menu the subcommands of a gadget file
func showSubHelp(filename string) {
	p := Parser{}
	if err := p.parseSettings(); err != nil {
		exitGracefully(err)
	}
	helpCommands, err := p.getSubHelp(filename)
	if err != nil {
		exitGracefully(fmt.Errorf("command %s not found", filename))
	}

	fmt.Fprintf(color.Output, `%s
	%23s   - %s
	%s   - %s
	%23s   - %s
	%v
	
	`, Yellow("AVAILABLE COMMANDS FOR: "+filename),
		"a, all", "Run all subcommands",
		"e, exclude [subcommand]", "Don't run specified subcommand",
		"vv, verbose", "Give verbose output of the commands running",
		strings.Trim(fmt.Sprint(helpCommands), "[]"))
}

// Used to display status messages
// e.x. Running: npm install
func showMessage(prefix string, message ...string) {
	fmt.Fprintf(color.Output, "%s: %s\n", Yellow(prefix), message)
}

func loadSpinner() *spinner.Spinner {
	return spinner.New(
		spinner.CharSets[43],
		100*time.Millisecond,
		spinner.WithHiddenCursor(true),
	) // Build our new spinner
}

// Remove duplicate file entries from a file slice,
// returns a new slice with the entries removed
func compactFilesSlice(files []fs.DirEntry) []fs.DirEntry {
	seen := make(map[string]bool)
	compactFiles := []fs.DirEntry{}

	for _, file := range files {
		if _, ok := seen[file.Name()]; !ok {
			seen[file.Name()] = true
			compactFiles = append(compactFiles, file)
		}
	}

	return compactFiles
}
