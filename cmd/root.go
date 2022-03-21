package cmd

import (
	"errors"
	"os"
	"os/exec"
	"sync"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// Application Version
var version = "0.2.0"

func Execute() {
	var message string
	command, appName, args, err := validateInput()
	s := loadSpinner()

	if err != nil {
		exitGracefully(err)
	}

	switch command {
	case "help":
		showHelp()

	case "version":
		color.Green("Application version: " + version)

	default:
		message, err = run(command, appName, args, s)
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

func run(filename, appName string, args []string, s *spinner.Spinner) (string, error) {
	parser := Parser{args: args}

	err := parser.parseJson(filename)
	if err != nil {
		return "", err
	}

	if appName == "" {
		return "", errors.New("appname was not provided")
	}

	mainCommands, otherCommands := parser.parseArgs()
	mainCommands[len(mainCommands)-1] = append(mainCommands[len(mainCommands)-1], appName)

	s.Start()

	msg, err := runMainCommands(mainCommands, s)
	if err != nil {
		return msg, err
	}

	s.Restart()
	showMessage("Created Project", appName)

	os.Chdir(appName)
	runSubCommands(otherCommands, s)

	s.Stop()
	return "\nApp Created Successfully: " + appName, nil
}

// Used to run all the main commands and throw an error if
// something goes wrong
func runMainCommands(mainCommands MainCommmands, s *spinner.Spinner) (string, error) {
	for _, cmd := range mainCommands {
		s.Restart()
		showMessage("Running", cmd...)

		c := exec.Command(cmd[0], cmd[1:]...)
		err := c.Run()
		if err != nil {
			return "Unable to execute command: " + cmd[0], err
		}
	}

	return "", nil
}

// Used to run all the subcommands either concurrently or by themselves
// based on the value of SubCommand.parallel. Displays a message if
// there is an error
func runSubCommands(subcommands []SubCommand, s *spinner.Spinner) {
	var wg sync.WaitGroup

	for _, command := range subcommands {
		s.Restart()
		showMessage("Installing", command.Name)

		if command.Parallel {
			wg.Add(1)
			go func(wg *sync.WaitGroup, command []string) {
				defer wg.Done()

				c := exec.Command(command[0], command[1:]...)
				err := c.Run()
				if err != nil {
					color.Yellow("Failed to execute %s\n", command)
					color.Red("Error: %v\n", err)
				}
			}(&wg, command.Command)
		} else {
			c := exec.Command(command.Command[0], command.Command[1:]...)
			err := c.Run()
			if err != nil {
				color.Yellow("Failed to execute %s\n", command)
				color.Red("Error: %v\n", err)
			}
		}
	}

	wg.Wait()
}
