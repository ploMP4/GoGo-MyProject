package internal

import (
	"fmt"
	"os"
	"strings"
)

func copyFileFromTemplate(templatePath, targetFile string) error {
	data, err := os.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile(data, targetFile)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

func copyDataToFile(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func fileExists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}

	return true
}

// Adds a string in the specified file either at the end of the file
// if the splitOn argument is an empty string or by splitting the file by the string
// specified in splitOn and appending it there
func editFile(filename, splitOn, toAppend string) {
	var settings string

	content, err := os.ReadFile(filename)
	if err != nil {
		exitGracefully(err)
	}

	if splitOn != "" { // Append after certain string in the file
		s := strings.Split(string(content), splitOn)
		s[0] += splitOn + toAppend

		settings = strings.Join(s, " ")
	} else { // Append at the end of file
		settings = string(content) + toAppend
	}

	err = os.WriteFile(filename, []byte(settings), 0644)
	if err != nil {
		exitGracefully(err)
	}
}

func replacePlaceholder(filepath, placeholder, value string) {
	showMessage(
		"Replacing",
		fmt.Sprintf("%s -> %s", placeholder, value),
	)

	content, err := os.ReadFile(filepath)
	if err != nil {
		exitGracefully(err)
	}

	replaced := strings.ReplaceAll(string(content), placeholder, value)

	err = os.WriteFile(filepath, []byte(replaced), 0644)
	if err != nil {
		exitGracefully(err)
	}
}
