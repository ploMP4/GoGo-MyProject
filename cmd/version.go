/*
Copyright © 2022 Kostas Artopoulos

*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var version = "1.0.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the application version",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("GoGo Version: " + version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
