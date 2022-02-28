/*
Copyright © 2022 Kostas Artopoulos

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gogo",
	Short: "A CLI tool to create starter boilerplate for you",
	Long: `GoGo is a CLI tool that creates the starter boilerplate 
for your projects and it's really helpfull for people
who use many different programming languages and frameworks.`,
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "help for "+rootCmd.Name())

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		ExitGracefully(err)
	}
}
