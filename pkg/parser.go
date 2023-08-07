package pkg

import (
	"fmt"
	"io/ioutil"
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

// We can have many main commands and commands
// come in the form of a string array so they can
// be passed in exec.Command()
type MainCommmands [][]string

// We can have many subcommands and they come in the form
// of an javascript object having what the argument name is going
// to be as the key and another object with the rest of the settings
// as the value so we translate it to a map[string]SubCommand
type SubCommands map[string]SubCommand

// We can change many files with different ways each.
// Files is an object with the key being a small
// description to print out to the user and the properties for the values
type FilesType map[string]File

// Describe the main yaml gadget
type Gadget struct {
	Commands    MainCommmands `yaml:"commands"`    // Array with the commands that will be executed. Note: commands should be passed as an array instead of using spaces e.x ["npx", "create-react-app"]
	Dirs        []string      `yaml:"dirs"`        // Array with names of directories that will be created
	SubCommands SubCommands   `yaml:"subCommands"` // Commands that can be passed after the initial command for optional features e.x. ts for typescript in a react command
	Help        string        `yaml:"help"`        // Help text for the command
}

// Describe a subcommand
type SubCommand struct {
	Name     string    `yaml:"name"`     // Name that will be displayed in the Installing status message e.x Installing: React
	Command  []string  `yaml:"command"`  // The command that will be executed.  Note: commands should be passed as an array instead of using spaces e.x ["npx", "create-react-app"]
	Override bool      `yaml:"override"` // Overrides the last command in the main commands array and runs this instead
	Parallel bool      `yaml:"parallel"` // Sets if the command will be run concurrently with others or not
	Exclude  bool      `yaml:"exclude"`  // If true this command will be ignored when the (a, all) flag is ran
	Files    FilesType `yaml:"files"`    // Specify files that you want to change
	Help     string    `yaml:"help"`     // Help text for the command
}

// Describe a file object
type File struct {
	Filepath string     `yaml:"filepath"` // Path where the file we want to edit is located. Path starts from the root file of our project
	Template bool       `yaml:"template"` // Specify if the file will be updated from an existing template
	Change   FileChange `yaml:"change"`   // Properties about changing the file
}

// Describe file change properties object
type FileChange struct {
	SplitOn string `yaml:"split-on"` // Specify string to split the file on
	Append  string `yaml:"append"`   // Content that will be appended after the split on
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

	yamlData, err := ioutil.ReadAll(yamlFile)
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
	yamlFile, err := os.Open(fmt.Sprintf("%s/%s.yaml", p.settings.GadgetPath, filename))
	if err != nil {
		return err
	}
	defer yamlFile.Close()

	yamlData, err := ioutil.ReadAll(yamlFile)
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

	files, err := ioutil.ReadDir(p.settings.GadgetPath)
	if err != nil {
		return nil
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
func (p *Parser) parseArgs() (MainCommmands, []SubCommand, []string, bool) {
	finalCommand := p.gadget.Commands[len(p.gadget.Commands)-1]
	var otherCommands []SubCommand

	all, verbose := p.parseFlags()
	if all {
		p.parseAll(&finalCommand, &otherCommands)
	} else {
		p.parseCmd(&finalCommand, &otherCommands)
	}

	p.gadget.Commands[len(p.gadget.Commands)-1] = finalCommand
	mainCommands := p.gadget.Commands
	dirs := p.gadget.Dirs

	return mainCommands, otherCommands, dirs, verbose
}

func (p *Parser) parseFlags() (all, verbose bool) {
	all = false
	verbose = false

	for idx, arg := range p.args {
		switch arg {
		case SHORT_ALL_FLAG, ALL_FLAG:
			all = true

		case SHORT_EXCLUDE_FLAG, EXLCUDE_FLAG:
			if subcommand, exists := p.gadget.SubCommands[p.args[idx+1]]; exists {
				subcommand.Exclude = true
				p.gadget.SubCommands[p.args[idx+1]] = subcommand
			}

		case SHORT_VERBOSE_FLAG, VERBOSE_FLAG:
			verbose = true
		}
	}

	return all, verbose
}

func (p *Parser) parseAll(finalCommand *[]string, otherCommands *[]SubCommand) {
	for _, value := range p.gadget.SubCommands {
		if value.Exclude {
			continue
		} else if value.Override {
			*finalCommand = value.Command
			showMessage("Using", value.Name)
		} else {
			*otherCommands = append(*otherCommands, value)
		}
	}
}

func (p *Parser) parseCmd(finalCommand *[]string, otherCommands *[]SubCommand) {
	for _, arg := range p.args {
		if value, exists := p.gadget.SubCommands[arg]; exists {
			if value.Override {
				*finalCommand = value.Command
				showMessage("Using", value.Name)
			} else {
				*otherCommands = append(*otherCommands, value)
			}
		}
	}
}
