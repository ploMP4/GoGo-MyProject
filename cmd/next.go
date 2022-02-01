package cmd

import (
	"os/exec"
	"sync"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type NextCmd struct {
	typescript bool
	redux      bool
	materialUI bool
	bootstrap  bool
	cmd        *cobra.Command

	react ReactCmd
}

var next NextCmd

func init() {
	next.cmd = &cobra.Command{
		Use:   "next [appname]",
		Short: "Create next app and optionally add typescript, redux, material-ui, bootstrap",
		Args:  NameExists,
		Run:   next.run,
	}

	rootCmd.AddCommand(next.cmd)

	next.cmd.Flags().BoolP("help", "h", false, "help for react")
	next.cmd.Flags().BoolVarP(&next.typescript, "typescript", "t", false, "Uses typescript template")
	next.cmd.Flags().BoolVarP(&next.redux, "redux", "r", false, "Install redux and creates boilerplate")
	next.cmd.Flags().BoolVarP(&next.materialUI, "mui", "m", false, "Install  Material-UI to use as a UI library")
	next.cmd.Flags().BoolVarP(&next.bootstrap, "bootstrap", "b", false, "Install Bootsrap to use as a UI library")
}

func (n *NextCmd) run(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup

	appName := args[0]
	color.Green("Creating Next app: " + appName)

	// Construct create-next-app command
	command := []string{"npx", "create-next-app", appName}

	switch {
	case n.typescript && n.redux:
		command = append(command, "--example")
		command = append(command, "with-redux")
		color.Blue("Template: Typescript")
		color.Blue("State management: Redux")

	case n.redux:
		command = append(command, "--example")
		command = append(command, "with-redux-thunk")
		color.Blue("Template: Javascript")
		color.Blue("State management: Redux")

	case n.typescript:
		command = append(command, "--typescript")
		color.Blue("Template: Typescript")
	}

	c := exec.Command(command[0], command[1:]...)

	s := LoadSpinner()
	s.Start()

	err := c.Run()
	if err != nil {
		ExitGracefully(err, "Unable to create next app")
	}

	if n.materialUI {
		wg.Add(1)
		s.Restart()
		color.Blue("UI Library: Material-UI")
		go n.react.installMUI(&wg, appName)
	}

	if n.bootstrap {
		wg.Add(1)
		s.Restart()
		color.Blue("UI Library: Bootstrap")
		go n.react.installBootstrap(&wg, appName)
	}

	wg.Wait()
	s.Stop()
	ExitGracefully(nil, "\nNext app created successfully under name: "+appName)
}
