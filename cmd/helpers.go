package cmd

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NameExists(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New(Red("no app name provided"))
	}

	return nil
}

func ExitGracefully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}

	if err != nil {
		color.Red("Error: %v\n", err)
		os.Exit(1)
	}

	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished")
	}

	os.Exit(0)
}

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

func LoadSpinner() *spinner.Spinner {
	return spinner.New(spinner.CharSets[43], 100*time.Millisecond) // Build our new spinner
}
