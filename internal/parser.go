package internal

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Used to parse the arguments passed in and
// the yaml file specified
type Parser struct {
	settings Settings // settings.yaml file
	gadget   Gadget   // The gadget file
	args     []string // Arguments passed
}

type Commands []string

type SubCommands map[string]SubCommand

type Files map[string]File

type Placeholders map[string]string

type Gadget struct {
	Commands    Commands    `yaml:"commands"` // Array with the commands that will be executed.
	Chdir       bool        `yaml:"chdir"`
	Dirs        []string    `yaml:"dirs"` // Array with names of directories that will be created
	Files       Files       `yaml:"files"`
	SubCommands SubCommands `yaml:"subCommands"` // Commands that can be passed after the initial command for optional features e.x. ts for typescript in a react command
	Help        string      `yaml:"help"`        // Help text for the command
}

type SubCommand struct {
	Name     string   `yaml:"name"`     // Name that will be displayed in the Installing status message e.x Installing: React
	Commands Commands `yaml:"commands"` // The commands that will be executed.
	Override bool     `yaml:"override"` // Overrides the last command in the main commands array and runs this instead
	Parallel bool     `yaml:"parallel"` // Sets if the command will be run concurrently with others or not
	Exclude  bool     `yaml:"exclude"`  // If true this command will be ignored when the (a, all) flag is ran
	Files    Files    `yaml:"files"`    // Specify files that you want to change
	Help     string   `yaml:"help"`     // Help text for the command
}

type File struct {
	Filepath string     `yaml:"filepath"` // Path where the file we want to edit is located. Path starts from the root file of our project
	Template bool       `yaml:"template"` // Specify if the file will be updated from an existing template
	Change   FileChange `yaml:"change"`   // Properties about changing the file
}

type FileChange struct {
	SplitOn     string       `yaml:"split-on"` // Specify string to split the file on
	Append      string       `yaml:"append"`   // Content that will be appended after the split on
	Placeholder Placeholders `yaml:"placeholder"`
}

// Parse the settings.yaml file that exists in
// the root of the application into the Parser.settings
func (p *Parser) parseSettings() error {
	e, err := os.Executable()
	if err != nil {
		return err
	}

	e_path, err := filepath.EvalSymlinks(e)
	if err != nil {
		return err
	}

	yamlFile, err := os.Open(filepath.Dir(e_path) + "/settings.yaml")
	if err != nil {
		return err
	}
	defer yamlFile.Close()

	yamlData, err := io.ReadAll(yamlFile)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(yamlData, &p.settings); err != nil {
		return err
	}

	return nil
}

// Check if a file with the name passed in by the user exists
// and parse its contents into the Parser.gadget
func (p *Parser) parseGadget(filename string) error {
	yamlFile, err := os.Open(
		fmt.Sprintf("%s/gadgets/%s.yaml", PROJECT_ROOT_DIR_NAME, filename),
	)
	if err != nil {
		yamlFile, err = os.Open(fmt.Sprintf("%s/%s.yaml", p.settings.GadgetPath, filename))
		if err != nil {
			return err
		}
	}
	defer yamlFile.Close()

	yamlData, err := io.ReadAll(yamlFile)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(yamlData, &p.gadget); err != nil {
		return err
	}

	return nil
}

// Parse and return the help commands for all the gadgets
func (p Parser) getHelp() []string {
	helpCommands := []string{}

	files, err := os.ReadDir(p.settings.GadgetPath)
	if err != nil {
		return nil
	}

	localFiles, err := os.ReadDir(PROJECT_ROOT_DIR_NAME + "/gadgets")
	if err == nil {
		files = append(files, localFiles...)
		files = compactFilesSlice(files)
	}

	for _, file := range files {
		filename := strings.Split(file.Name(), ".")[0]
		_ = p.parseGadget(filename)
		helpCommands = append(helpCommands, fmt.Sprintf("\n%30s   - %s", filename, p.gadget.Help))
	}

	return helpCommands
}

// Parse and return help for the subcommands of a gadget
func (p Parser) getSubHelp(filename string) ([]string, error) {
	helpCommands := []string{}

	err := p.parseGadget(filename)
	if err != nil {
		return nil, err
	}

	for name, command := range p.gadget.SubCommands {
		helpCommands = append(helpCommands, fmt.Sprintf("\n%31s   - %s", name, command.Help))
	}

	return helpCommands, nil
}

// Use the parsed gadget and the args to construct
// the dirs, main and sub commands and return them
func (p *Parser) parseArgs() (Commands, []SubCommand, []string, bool, string) {
	commands := p.gadget.Commands
	var subCommands []SubCommand

	all, verbose, appname := p.parseFlagsAndPlaceholders()

	if all {
		p.parseAll(&commands, &subCommands)
	} else {
		p.parseCmd(&commands, &subCommands)
	}

	dirs := p.gadget.Dirs

	return commands, subCommands, dirs, verbose, appname
}

func (p *Parser) parseFlagsAndPlaceholders() (all, verbose bool, appname string) {
	all = false
	verbose = false
	appname = ""

	for idx, arg := range p.args {
		switch arg {
		case SHORT_ALL_FLAG, ALL_FLAG:
			all = true
			if idx < len(p.args) {
				p.args = append(p.args[:idx], p.args[idx+1:]...)
			} else {
				p.args = p.args[:idx]
			}

		case SHORT_EXCLUDE_FLAG, EXLCUDE_FLAG:
			if subcommand, exists := p.gadget.SubCommands[p.args[idx+1]]; exists {
				subcommand.Exclude = true
				p.gadget.SubCommands[p.args[idx+1]] = subcommand
				if idx < len(p.args) {
					p.args = append(p.args[:idx+1], p.args[idx+2:]...)
				} else {
					p.args = p.args[:idx+1]
				}
			}

		case SHORT_VERBOSE_FLAG, VERBOSE_FLAG:
			verbose = true
			if idx < len(p.args) {
				p.args = append(p.args[:idx], p.args[idx+1:]...)
			} else {
				p.args = p.args[:idx]
			}

		case PLACEHOLDER_APPNAME:
			appname = p.args[idx+1]
			if idx < len(p.args) {
				p.args = append(p.args[:idx+1], p.args[idx+2:]...)
			} else {
				p.args = p.args[:idx+1]
			}
		}
	}

	return all, verbose, appname
}

func (p *Parser) parseAll(commands *Commands, subCommands *[]SubCommand) {
	for _, value := range p.gadget.SubCommands {
		if value.Exclude {
			continue
		} else if value.Override {
			*commands = value.Commands
			showMessage("Using", value.Name)
		} else {
			*subCommands = append(*subCommands, value)
		}
	}
}

func (p *Parser) parseCmd(commands *Commands, subCommands *[]SubCommand) {
	for _, arg := range p.args {
		if value, exists := p.gadget.SubCommands[arg]; exists {
			if value.Override {
				*commands = value.Commands
				showMessage("Using", value.Name)
			} else {
				*subCommands = append(*subCommands, value)
			}
		}
	}
}
