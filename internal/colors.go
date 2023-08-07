package internal

import "github.com/fatih/color"

// Wrapper function for color.YellowString
// Returns a yellow string
func Yellow(str string) string {
	return color.YellowString(str)
}

// Wrapper function for color.RedString
// Returns a red string
func Red(str string) string {
	return color.RedString(str)
}

// Wrapper function for color.GreenString
// Returns a green string
func Green(str string) string {
	return color.GreenString(str)
}
