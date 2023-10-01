package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/fatih/color"
)

type Gadget struct {
	Commands    Commands    `yaml:"commands"` // Array with the commands that will be executed.
	Chdir       bool        `yaml:"chdir"`
	Dirs        []string    `yaml:"dirs"` // Array with names of directories that will be created
	Files       Files       `yaml:"files"`
	SubCommands SubCommands `yaml:"subCommands"` // Commands that can be passed after the initial command for optional features e.x. ts for typescript in a react command
	Help        string      `yaml:"help"`        // Help text for the command
}

// Used to run all the main commands and throw an error if
// something goes wrong
func (g Gadget) runCommands() (string, error) {
	if g.Commands == nil {
		return "", nil
	}

	for _, command := range g.Commands {
		if strings.Contains(command, PLACEHOLDER_APPNAME) && app.appname != "" {
			command = strings.ReplaceAll(command, PLACEHOLDER_APPNAME, app.appname)
		}
		app.spinner.Restart()
		showMessage("Running", command)

		cmd := strings.Fields(command)
		c := exec.Command(cmd[0], cmd[1:]...)

		if app.verbose {
			app.spinner.Stop()
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
		}

		err := c.Run()
		if err != nil {
			return "Unable to execute command: " + cmd[0], err
		}
	}

	return "", nil
}

// Used to run all the subcommands either concurrently or by themselves
// based on the value of SubCommand.parallel. Displays a message if
// there is an error
func (g Gadget) runSubCommands() {
	var wg sync.WaitGroup

	for _, subcmd := range g.SubCommands {
		app.spinner.Restart()
		showMessage("Running", subcmd.Name)

		if subcmd.Parallel {
			wg.Add(1)
			go func(subcmd SubCommand) {
				defer wg.Done()

				err := subcmd.execute()
				if err != nil {
					color.Yellow("Failed to execute %s\n", subcmd)
					color.Red("Error: %v\n", err)
				}
			}(subcmd)
		} else {
			err := subcmd.execute()
			if err != nil {
				color.Yellow("Failed to execute %s\n", subcmd)
				color.Red("Error: %v\n", err)
			}
		}
	}

	wg.Wait()
}

func (g Gadget) handleFiles() {
	if g.Files != nil {
		for name, file := range g.Files {
			if file.Template {
				templatePath, err := g.getTemplatePath(file)
				if err != nil {
					showMessage("Warning", err.Error())
					showMessage("Skiping...", name)
					return
				}

				file.handleTemplate(templatePath, app.parser.args)
			} else {
				file.handleEdit(name, app.appname)
			}

			g.updateFileName(file)
		}
	}
}

func (g Gadget) updateFileName(file File) {
	for idx, arg := range app.parser.args {
		if arg == PLACEHOLDER_FILENAME && len(g.Files) <= 1 {
			filepathSlice := strings.Split(file.Filepath, "/")
			filepathSlice[len(filepathSlice)-1] = app.parser.args[idx+1]
			filepathNew := strings.Join(filepathSlice, "/")
			showMessage("Renaming", fmt.Sprintf("%s -> %s", file.Filepath, filepathNew))
			if err := os.Rename(file.Filepath, filepathNew); err != nil {
				showMessage("Warning", err.Error())
			}
		}
	}
}

func (g Gadget) getTemplatePath(file File) (string, error) {
	var templatePath string

	templatePath = fmt.Sprintf(
		"./%s/templates/%s/%s",
		PROJECT_ROOT_DIR_NAME,
		app.gadgetName,
		file.Filepath,
	)

	if !fileExists(templatePath) {
		templatePath = fmt.Sprintf(
			"../%s/templates/%s/%s",
			PROJECT_ROOT_DIR_NAME,
			app.gadgetName,
			file.Filepath,
		)
	}

	if !fileExists(templatePath) {
		templatePath = fmt.Sprintf(
			"%s/%s/%s",
			app.parser.settings.TemplatePath,
			app.gadgetName,
			file.Filepath,
		)
	}

	if !fileExists(templatePath) {
		return "", errors.New("template not found")
	}

	return templatePath, nil
}
