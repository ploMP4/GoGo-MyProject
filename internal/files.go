package internal

import (
	"fmt"
	"os"
	"strings"
)

type File struct {
	Filepath string     `yaml:"filepath"` // Path where the file we want to edit is located. Path starts from the root file of our project
	Template bool       `yaml:"template"` // Specify if the file will be updated from an existing template
	Change   FileChange `yaml:"change"`   // Properties about changing the file
}

type FileChange struct {
	SplitOn      string       `yaml:"split-on"` // Specify string to split the file on
	Append       string       `yaml:"append"`   // Content that will be appended after the split on
	Placeholders Placeholders `yaml:"placeholder"`
}

func (f *File) handleTemplate(templatePath string, args []string) {
	showMessage("Copying", f.Filepath)

	app.spinner.Restart()
	if err := f.copyFromTemplate(templatePath); err != nil {
		exitGracefully(err)
	}

	if f.Change.Placeholders != nil {
		for placeholder, defaultValue := range f.Change.Placeholders {
			app.spinner.Restart()
			found := f.findAndReplacePlaceholder(placeholder, args)
			if !found {
				if err := f.replacePlaceholder(placeholder, defaultValue); err != nil {
					exitGracefully(err)
				}
			}
		}
	}
}

func (f *File) handleEdit(name, appname string) {
	if strings.Contains(f.Filepath, PLACEHOLDER_APPNAME) {
		f.Filepath = strings.ReplaceAll(f.Filepath, PLACEHOLDER_APPNAME, appname)
	}

	showMessage("Adding", name, "in", Green(f.Filepath))
	if err := f.edit(); err != nil {
		exitGracefully(err)
	}

}

func (f *File) copyFromTemplate(templatePath string) error {
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	err = f.copyDataToFile(data)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) copyDataToFile(data []byte) error {
	err := os.WriteFile(f.Filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Adds a string in the specified file either at the end of the file
// if the splitOn argument is an empty string or by splitting the file by the string
// specified in splitOn and appending it there
func (f *File) edit() error {
	var settings string

	content, err := os.ReadFile(f.Filepath)
	if err != nil {
		return err
	}

	if f.Change.SplitOn != "" { // Append after certain string in the file
		s := strings.Split(string(content), f.Change.SplitOn)
		s[0] += f.Change.SplitOn + f.Change.Append

		settings = strings.Join(s, " ")
	} else { // Append at the end of file
		settings = string(content) + f.Change.Append
	}

	err = os.WriteFile(f.Filepath, []byte(settings), 0644)
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

func (f *File) findAndReplacePlaceholder(placeholder string, args []string) bool {
	for idx, arg := range args {
		if arg == placeholder {
			if err := f.replacePlaceholder(placeholder, args[idx+1]); err != nil {
				exitGracefully(err)
			}

			return true
		}
	}

	return false
}

func (f *File) replacePlaceholder(placeholder, value string) error {
	showMessage(
		"Replacing",
		fmt.Sprintf("%s -> %s", placeholder, value),
	)

	content, err := os.ReadFile(f.Filepath)
	if err != nil {
		return err
	}

	replaced := strings.ReplaceAll(string(content), placeholder, value)

	err = os.WriteFile(f.Filepath, []byte(replaced), 0644)
	if err != nil {
		return err
	}

	return nil
}
