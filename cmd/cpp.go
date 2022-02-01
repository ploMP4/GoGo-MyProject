/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type CppCmd struct {
	cmd *cobra.Command
}

var cpp CppCmd

func init() {
	cpp.cmd = &cobra.Command{
		Use:     "cpp [appname]",
		Aliases: []string{"c++"},
		Short:   "Create c++ application",
		Args:    NameExists,
		Run:     cpp.run,
	}

	rootCmd.AddCommand(cpp.cmd)
}

func (c *CppCmd) run(cmd *cobra.Command, args []string) {
	appName := args[0]

	color.Green("Creating C++ application: " + appName)
	dirs := []string{"bin", "includes", "src"}

	err := os.Mkdir(appName, 0755)
	if err != nil {
		ExitGracefully(err)
	}

	err = os.Chdir(appName)
	if err != nil {
		ExitGracefully(err)
	}

	rootPath, err := os.Getwd()
	if err != nil {
		ExitGracefully(err)
	}

	color.Green("Creating Project Structure...")
	for _, dir := range dirs {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			ExitGracefully(err)
		}
	}

	main := rootPath + "/src/main.cpp"
	err = copyFileFromTemplate("templates/cpp/main.cpp.txt", main)
	if err != nil {
		ExitGracefully(err)
	}

	color.Green("Creating Makefile...")
	makefile := rootPath + "/Makefile"
	err = copyFileFromTemplate("templates/cpp/Makefile", makefile)
	if err != nil {
		ExitGracefully(err)
	}

	ExitGracefully(nil, "C++ project created successfully under the name: "+appName)
}
