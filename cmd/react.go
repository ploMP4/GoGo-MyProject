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
		Short: "Create react app and optionally add typescript, redux, material-ui, bootstrap",
		Args:  NameExists,
		Run:   react.run,
	}

	rootCmd.AddCommand(react.cmd)

	react.cmd.Flags().BoolP("help", "h", false, "help for react")
	react.cmd.Flags().BoolVarP(&react.typescript, "typescript", "t", false, "Uses typescript template")
	react.cmd.Flags().BoolVarP(&react.redux, "redux", "r", false, "Install redux and creates boilerplate")
	react.cmd.Flags().BoolVarP(&react.materialUI, "mui", "m", false, "Install  Material-UI to use as a UI library")
	react.cmd.Flags().BoolVarP(&react.bootstrap, "bootstrap", "b", false, "Install Bootsrap to use as a UI library")
}

func (r *ReactCmd) run(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup

	appName := args[0]
	color.Green("Creating React app: " + appName)

	// Construct create-react-app command
	command := []string{"npx", "create-react-app", appName}

	switch {
	case r.typescript && r.redux:
		command = append(command, "--template")
		command = append(command, "redux-typescript")
		color.Blue("Template: Typescript")
		color.Blue("State management: Redux")

	case r.typescript:
		command = append(command, "--template")
		command = append(command, "typescript")
		color.Blue("Template: Typescript")

	case r.redux:
		command = append(command, "--template")
		command = append(command, "redux")
		color.Blue("Template: Javascript")
		color.Blue("State management: Redux")
	}

	// Execute create-react-app with provided arguments
	c := exec.Command(command[0], command[1:]...)

	s := LoadSpinner()
	s.Start()

	err := c.Run()
	if err != nil {
		ExitGracefully(err, "Unable to create react app")
	}

	if r.materialUI {
		wg.Add(1)
		s.Restart()
		color.Blue("UI Library: Material-UI")
		go r.installMUI(&wg, appName)
	}

	if r.bootstrap {
		wg.Add(1)
		s.Restart()
		color.Blue("UI Library: Bootstrap")
		go r.installBootstrap(&wg, appName)
	}

	wg.Wait()
	s.Stop()
	ExitGracefully(nil, "\nReact app created successfully under name: "+appName)
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

func (r *ReactCmd) installBootstrap(wg *sync.WaitGroup, appName string) {
	defer wg.Done()
	color.Green("Installing Bootstrap...")

	os.Chdir(appName)
	cmd := exec.Command("npm", "install", "react-bootstrap")

	err := cmd.Run()
	if err != nil {
		ExitGracefully(err, "Unable to install bootstrap")
	}
}
