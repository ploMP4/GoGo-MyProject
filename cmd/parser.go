package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Used to parse the arguments passed in and
// the json file specified
type Parser struct {
	configPath string   // Path of folder containing json files
	json       Json     // The json file name
	args       []string // Arguments passed
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
// Files is an object with the filename we want to change
// as the key and the properties for the value
type FilesType map[string]File

// Describe the main json config file
type Json struct {
	Commands    MainCommmands `json:"commands"`    // Array with the commands that will be executed. Note: commands should be passed as an array instead of using spaces e.x ["npx", "create-react-app"]
	Dirs        []string      `json:"dirs"`        // Array with names of directories that will be created
	SubCommands SubCommands   `json:"subCommands"` // Commands that can be passed after the initial command for optional features e.x. ts for typescript in a react command
	Help        string        `json:"help"`        // Help text for the command
}

// Describe a subcommand
type SubCommand struct {
	Name     string    `json:"name"`     // Name that will be displayed in the Installing status message e.x Installing: React
	Command  []string  `json:"command"`  // The command that will be executed.  Note: commands should be passed as an array instead of using spaces e.x ["npx", "create-react-app"]
	Override bool      `json:"override"` // Overrides the last command in the main commands array and runs this instead
	Parallel bool      `json:"parallel"` // Sets if the command will be run concurrently with others or not
	Files    FilesType `json:"files"`    // Specify files that you want to change
	Help     string    `json:"help"`     // Help text for the command
}

// Describe a file object
type File struct {
	Template bool       `json:"template"` // Specify if the file will be updated from an existing template
	Change   FileChange `json:"change"`   // Properties about changing the file
}

// Describe file change properties object
type FileChange struct {
	SplitOn string `json:"splitOn"` // Specify string to split the file on
	Append  string `json:"append"`  // Content that will be appended after the split on
}

// Check if a file with the name passed in by the user exists
// and parse its contents into the Parser.json
func (p *Parser) parseJson(filename string) error {
	jsonFile, err := os.Open(fmt.Sprintf("%s/%s.json", p.configPath, filename))
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(jsonData, &p.json); err != nil {
		return err
	}

	return nil
}

// Parse and return the help commands for the json files
func (p Parser) getHelp() []string {
	helpCommands := []string{}

	files, err := ioutil.ReadDir(p.configPath)
	if err != nil {
		return nil
	}

	for _, file := range files {
		filename := strings.Split(file.Name(), ".")[0]
		p.parseJson(filename)
		helpCommands = append(helpCommands, fmt.Sprintf("\n%15s   - %s", filename, p.json.Help))
	}

	return helpCommands
}

// Use the parsed json file and the args to construct
// the main and sub commands and return them
func (p *Parser) parseArgs() (MainCommmands, []SubCommand) {
	finalCommand := p.json.Commands[len(p.json.Commands)-1]
	var otherCommands []SubCommand

	for _, arg := range p.args {
		if value, isMapContainsKey := p.json.SubCommands[arg]; isMapContainsKey {
			if value.Override {
				finalCommand = value.Command
				showMessage("Using", value.Name)
			} else {
				otherCommands = append(otherCommands, value)
			}
		}
	}

	p.json.Commands[len(p.json.Commands)-1] = finalCommand
	mainCommands := p.json.Commands

	return mainCommands, otherCommands
}
