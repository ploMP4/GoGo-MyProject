package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type SubCommand struct {
	Name     string   `yaml:"name"`     // Name that will be displayed in the Installing status message e.x Installing: React
	Commands Commands `yaml:"commands"` // The commands that will be executed.
	Override bool     `yaml:"override"` // Overrides the last command in the main commands array and runs this instead
	Parallel bool     `yaml:"parallel"` // Sets if the command will be run concurrently with others or not
	Exclude  bool     `yaml:"exclude"`  // If true this command will be ignored when the (a, all) flag is ran
	Files    Files    `yaml:"files"`    // Specify files that you want to change
	Help     string   `yaml:"help"`     // Help text for the command
}

func (s SubCommand) execute() error {
	if s.Commands != nil {
		err := s.runCommands()
		if err != nil {
			return err
		}
	}

	s.handleFiles()

	return nil
}

func (s SubCommand) runCommands() error {
	for _, command := range s.Commands {
		cmd := strings.Fields(command)
		c := exec.Command(cmd[0], cmd[1:]...)

		if app.verbose {
			app.spinner.Stop()
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
		}

		err := c.Run()
		if err != nil {
			return err
		}
	}

	return nil

}

func (s SubCommand) handleFiles() {
	if s.Files != nil {
		for name, file := range s.Files {
			if file.Template {
				templatePath, err := s.getTemplatePath(file, s.Name)
				if err != nil {
					showMessage("Warning", err.Error())
					showMessage("Skiping...", name)
					return
				}

				file.handleTemplate(templatePath, app.parser.args)
			} else {
				file.handleEdit(name, app.appname)
			}

			s.updateFileName(file)
		}
	}
}

func (s SubCommand) updateFileName(file File) {
	for idx, arg := range app.parser.args {
		if arg == PLACEHOLDER_FILENAME && len(s.Files) <= 1 {
			filepathSlice := strings.Split(file.Filepath, "/")
			filepathSlice[len(filepathSlice)-1] = app.parser.args[idx+1]
			filepathNew := strings.Join(filepathSlice, "/")
			if err := os.Rename(file.Filepath, filepathNew); err != nil {
				showMessage("Warning", err.Error())
			}
		}
	}
}

func (s SubCommand) getTemplatePath(file File, commandName string) (string, error) {
	templatePath := fmt.Sprintf(
		"../%s/templates/%s/%s/%s",
		PROJECT_ROOT_DIR_NAME,
		app.gadgetName,
		commandName,
		file.Filepath,
	)

	if !fileExists(templatePath) {
		templatePath = fmt.Sprintf(
			"./%s/templates/%s/%s/%s",
			PROJECT_ROOT_DIR_NAME,
			app.gadgetName,
			commandName,
			file.Filepath,
		)
	}

	if !fileExists(templatePath) {
		templatePath = fmt.Sprintf(
			"%s/%s/%s/%s",
			app.parser.settings.TemplatePath,
			app.gadgetName,
			commandName,
			file.Filepath,
		)
	}

	if !fileExists(templatePath) {
		return "", errors.New("template not found")
	}

	return templatePath, nil
}
