/*
Copyright © 2022 Kostas Artopoulos

*/
package cmd

import (
	"os"
	"os/exec"
	"sync"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type ReactCmd struct {
	typescript bool
	redux      bool
	materialUI bool
	bootstrap  bool
	cmd        *cobra.Command
}

var react ReactCmd

func init() {
	react.cmd = &cobra.Command{
		Use:   "react [appname]",
		Short: Yellow("Create react app and optionally add typescript, redux, material-ui, bootstrap"),
		Args:  NameExists,
		Run:   react.run,
	}

	rootCmd.AddCommand(react.cmd)

	react.cmd.Flags().BoolP("help", "h", false, Yellow("help for react"))
	react.cmd.Flags().BoolVarP(&react.typescript, "typescript", "t", false, Yellow("Uses typescript template"))
	react.cmd.Flags().BoolVarP(&react.redux, "redux", "r", false, Yellow("Install redux and creates boilerplate"))
	react.cmd.Flags().BoolVarP(&react.materialUI, "mui", "m", false, Yellow("Install  Material-UI to use as a UI library"))
	react.cmd.Flags().BoolVarP(&react.bootstrap, "bootstrap", "b", false, Yellow("Install Bootsrap to use as a UI library"))
}

func (r *ReactCmd) run(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup

	appName := args[0]
	color.Green("Creating React app: " + appName)

	// Construct create-react-app command
	command := []string{"npx", "create-react-app", appName}
	if r.typescript {
		command = append(command, "--template")
		command = append(command, "typescript")
		color.Blue("Template: Typescript")
	}

	// Execute create-react-app with provided arguments
	c := exec.Command(command[0], command[1:]...)

	s := LoadSpinner()
	s.Start()
	defer s.Stop()

	err := c.Run()
	if err != nil {
		ExitGracefully(err)
	}

	s.Restart()

	// TODO: create redux boilerplate from template
	if r.redux {
		wg.Add(1)
		color.Blue("State management: Redux")
		go r.installRedux(&wg, appName)
	}

	if r.materialUI {
		wg.Add(1)
		color.Blue("UI Library: Material-UI")
		go r.installMUI(&wg, appName)
	}

	wg.Wait()
	ExitGracefully(nil, Green("\nReact app created successfully under name: "+appName))
}

func (r *ReactCmd) installRedux(wg *sync.WaitGroup, appName string) {
	var cmd *exec.Cmd
	defer wg.Done()

	color.Green("Installing Redux...")

	os.Chdir(appName)

	if react.typescript {
		cmd = exec.Command("npm", "install", "@reduxjs/toolkit", "@types/react-redux", "react-redux")
	} else {
		cmd = exec.Command("npm", "install", "react-redux")
	}

	err := cmd.Run()
	if err != nil {
		ExitGracefully(err, "Unable to install redux")
	}
}

func (r *ReactCmd) installMUI(wg *sync.WaitGroup, appName string) {
	defer wg.Done()
	color.Green("Installing Material UI...")

	os.Chdir(appName)
	cmd := exec.Command("npm", "install", "@mui/material", "@emotion/react", "@emotion/styled")

	err := cmd.Run()
	if err != nil {
		ExitGracefully(err, "Unable to install material-ui")
	}
}
