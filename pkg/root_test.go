package pkg

import (
	"os"
	"syscall"
	"testing"
)

var app App

func TestRoot_validateInput(t *testing.T) {
	defer func(stdout *os.File) {
		os.Stdout = stdout
	}(os.Stdout)
	os.Stdout = os.NewFile(uintptr(syscall.Stdin), os.DevNull)

	os.Args = []string{}
	_, _, _, err := validateInput()
	if err == nil {
		t.Error("no error returned when nothing was provided")
	}

	os.Args = []string{"", "cpp", "my_app"}
	command, appName, _, err := validateInput()
	if err != nil {
		t.Error(err)
	}

	if command != "cpp" {
		t.Error("function return wrong command name: ", command)
	}

	if appName != "my_app" {
		t.Error("function returned wrong app name: ", appName)
	}

	os.Args = append(os.Args, "v")
	_, _, args, err := validateInput()
	if err != nil {
		t.Error(err)
	}
	if args == nil {
		t.Error("function returned no amount of arguments when it should have returned 1")
	}

	app.filename = command
	app.appName = appName
	app.parser.args = args
	app.spinner = loadSpinner()
}

func TestRoot_runMainCommands(t *testing.T) {
	defer func(stdout *os.File) {
		os.Stdout = stdout
	}(os.Stdout)
	os.Stdout = os.NewFile(uintptr(syscall.Stdin), os.DevNull)

	message, err := app.runMainCommands([][]string{{"echo", app.appName}})
	if err != nil {
		t.Error(message, err)
	}
	app.spinner.Stop()

	_, err = app.runMainCommands([][]string{{"some_command"}})
	if err == nil {
		t.Error("no error returned when false commands where passed")
	}
	app.spinner.Stop()
}

func TestRoot_executeSubCommand(t *testing.T) {
	sub_command := SubCommand{
		Name:    "test",
		Command: []string{"some_command"},
	}
	err := app.executeSubCommand(sub_command)
	if err == nil {
		t.Error("no error returned when false commands where passed")
	}

	sub_command.Command = []string{"echo", "test"}
	err = app.executeSubCommand(sub_command)
	if err != nil {
		t.Error(err)
	}
}
