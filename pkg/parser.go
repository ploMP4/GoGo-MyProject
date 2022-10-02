package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

// Used to parse the arguments passed in and
// the toml file specified
type Parser struct {
	settings Settings // settings.toml file
	config   Config   // The config file
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

// Describe the main toml config file
type Config struct {
	Commands    MainCommmands `toml:"commands"`    // Array with the commands that will be executed. Note: commands should be passed as an array instead of using spaces e.x ["npx", "create-react-app"]
	Dirs        []string      `toml:"dirs"`        // Array with names of directories that will be created
	SubCommands SubCommands   `toml:"subCommands"` // Commands that can be passed after the initial command for optional features e.x. ts for typescript in a react command
	Help        string        `toml:"help"`        // Help text for the command
}

// Describe a subcommand
type SubCommand struct {
	Name     string    `toml:"name"`     // Name that will be displayed in the Installing status message e.x Installing: React
	Command  []string  `toml:"command"`  // The command that will be executed.  Note: commands should be passed as an array instead of using spaces e.x ["npx", "create-react-app"]
	Override bool      `toml:"override"` // Overrides the last command in the main commands array and runs this instead
	Parallel bool      `toml:"parallel"` // Sets if the command will be run concurrently with others or not
	Exclude  bool      `toml:"exclude"`  // If true this command will be ignored when the (a, all) flag is ran
	Files    FilesType `toml:"files"`    // Specify files that you want to change
	Help     string    `toml:"help"`     // Help text for the command
}

// Describe a file object
type File struct {
	Filepath string     `toml:"filepath"` // Path where the file we want to edit is located. Path starts from the root file of our project
	Template bool       `toml:"template"` // Specify if the file will be updated from an existing template
	Change   FileChange `toml:"change"`   // Properties about changing the file
}

// Describe file change properties object
type FileChange struct {
	SplitOn string `toml:"split-on"` // Specify string to split the file on
	Append  string `toml:"append"`   // Content that will be appended after the split on
}

// Parse the settings.toml file that exists in
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

	tomlFile, err := os.Open(filepath.Dir(e_path) + "/settings.toml")
	if err != nil {
		return err
	}
	defer tomlFile.Close()

	tomlData, err := ioutil.ReadAll(tomlFile)
	if err != nil {
		return err
	}

	if err = toml.Unmarshal(tomlData, &p.settings); err != nil {
		return err
	}

	return nil
}

// Check if a file with the name passed in by the user exists
// and parse its contents into the Parser.config
func (p *Parser) parseConfig(filename string) error {
	tomlFile, err := os.Open(fmt.Sprintf("%s/%s.toml", p.settings.ConfigPath, filename))
	if err != nil {
		return err
	}
	defer tomlFile.Close()

	tomlData, err := ioutil.ReadAll(tomlFile)
	if err != nil {
		return err
	}

	if err = toml.Unmarshal(tomlData, &p.config); err != nil {
		return err
	}

	return nil
}

// Parse and return the help commands for all the config files
func (p Parser) getHelp() []string {
	helpCommands := []string{}

	files, err := ioutil.ReadDir(p.settings.ConfigPath)
	if err != nil {
		return nil
	}

	for _, file := range files {
		filename := strings.Split(file.Name(), ".")[0]
		_ = p.parseConfig(filename)
		helpCommands = append(helpCommands, fmt.Sprintf("\n%30s   - %s", filename, p.config.Help))
	}

	return helpCommands
}

// Parse and return help for the subcommands of a config file
func (p Parser) getSubHelp(filename string) ([]string, error) {
	helpCommands := []string{}

	err := p.parseConfig(filename)
	if err != nil {
		return nil, err
	}

	for name, command := range p.config.SubCommands {
		helpCommands = append(helpCommands, fmt.Sprintf("\n%31s   - %s", name, command.Help))
	}

	return helpCommands, nil
}

// Use the parsed config file and the args to construct
// the dirs, main and sub commands and return them
func (p *Parser) parseArgs() (MainCommmands, []SubCommand, []string) {
	finalCommand := p.config.Commands[len(p.config.Commands)-1]
	var otherCommands []SubCommand

	all := false
	for idx, arg := range p.args {
		if arg == "all" || arg == "a" {
			all = true
		}

		if arg == "exclude" || arg == "e" {
			if subcommand, ok := p.config.SubCommands[p.args[idx+1]]; ok {
				subcommand.Exclude = true

				p.config.SubCommands[p.args[idx+1]] = subcommand
			}
		}
	}

	if all {
		for _, value := range p.config.SubCommands {
			if value.Exclude {
				continue
			} else if value.Override {
				finalCommand = value.Command
				showMessage("Using", value.Name)
			} else {
				otherCommands = append(otherCommands, value)
			}
		}
	} else {
		for _, arg := range p.args {
			if value, isMapContainsKey := p.config.SubCommands[arg]; isMapContainsKey {
				if value.Override {
					finalCommand = value.Command
					showMessage("Using", value.Name)
				} else {
					otherCommands = append(otherCommands, value)
				}
			}
		}
	}

	p.config.Commands[len(p.config.Commands)-1] = finalCommand
	mainCommands := p.config.Commands
	dirs := p.config.Dirs

	return mainCommands, otherCommands, dirs
}
