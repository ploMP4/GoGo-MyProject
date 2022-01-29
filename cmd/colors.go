package cmd

import "github.com/fatih/color"

func Yellow(str string) string {
	return color.YellowString(str)
}

func Red(str string) string {
	return color.RedString(str)
}

func Green(str string) string {
	return color.GreenString(str)
}
