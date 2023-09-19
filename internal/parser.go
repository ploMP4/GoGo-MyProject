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
	args     []string // Arguments passed
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
func (p *Parser) parseGadget(filename string) (Gadget, error) {
	var gadget Gadget

	yamlFile, err := os.Open(
		fmt.Sprintf("%s/gadgets/%s.yaml", PROJECT_ROOT_DIR_NAME, filename),
	)
	if err != nil {
		yamlFile, err = os.Open(fmt.Sprintf("%s/%s.yaml", p.settings.GadgetPath, filename))
		if err != nil {
			return gadget, err
		}
	}
	defer yamlFile.Close()

	yamlData, err := io.ReadAll(yamlFile)
	if err != nil {
		return gadget, err
	}

	if err = yaml.Unmarshal(yamlData, &gadget); err != nil {
		return gadget, err
	}

	return gadget, nil
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
		gadget, _ := p.parseGadget(filename)
		helpCommands = append(helpCommands, fmt.Sprintf("\n%30s   - %s", filename, gadget.Help))
	}

	return helpCommands
}

// Parse and return help for the subcommands of a gadget
func (p Parser) getSubHelp(filename string) ([]string, error) {
	helpCommands := []string{}

	gadget, err := p.parseGadget(filename)
	if err != nil {
		return nil, err
	}

	for name, command := range gadget.SubCommands {
		helpCommands = append(helpCommands, fmt.Sprintf("\n%31s   - %s", name, command.Help))
	}

	return helpCommands, nil
}

// Use the parsed gadget and the args to construct
// the dirs, main and sub commands and return them
func (p *Parser) parseArgs(gadget Gadget) (Gadget, []string, bool, string) {
	commands := gadget.Commands
	subCommands := make(SubCommands)

	all, verbose, appname := p.parseFlagsAndPlaceholders(&gadget)

	if all {
		p.parseAll(&gadget, &commands, subCommands)
	} else {
		p.parseCmd(&gadget, &commands, subCommands)
	}

	gadget.Commands = commands
	gadget.SubCommands = subCommands
	dirs := gadget.Dirs

	return gadget, dirs, verbose, appname
}

func (p *Parser) parseFlagsAndPlaceholders(gadget *Gadget) (all, verbose bool, appname string) {
	all = false
	verbose = false
	appname = ""

	for idx, arg := range p.args {
		switch arg {
		case SHORT_ALL_FLAG, ALL_FLAG:
			all = true
			if idx < len(p.args) {
				p.args = append(p.args[:idx], p.args[idx:]...)
			} else {
				p.args = p.args[:idx]
			}

		case SHORT_EXCLUDE_FLAG, EXLCUDE_FLAG:
			if subcommand, exists := gadget.SubCommands[p.args[idx+1]]; exists {
				subcommand.Exclude = true
				gadget.SubCommands[p.args[idx+1]] = subcommand
				if idx < len(p.args) {
					p.args = append(p.args[:idx+1], p.args[idx+1:]...)
				} else {
					p.args = p.args[:idx+1]
				}
			}

		case SHORT_VERBOSE_FLAG, VERBOSE_FLAG:
			verbose = true
			if idx < len(p.args) {
				p.args = append(p.args[:idx], p.args[idx:]...)
			} else {
				p.args = p.args[:idx]
			}

		case PLACEHOLDER_APPNAME:
			appname = p.args[idx+1]
			if idx < len(p.args) {
				p.args = append(p.args[:idx+1], p.args[idx+1:]...)
			} else {
				p.args = p.args[:idx+1]
			}
		}
	}

	return all, verbose, appname
}

func (p *Parser) parseAll(gadget *Gadget, commands *Commands, subCommands SubCommands) {
	for key, value := range gadget.SubCommands {
		if value.Exclude {
			continue
		} else if value.Override {
			*commands = value.Commands
			showMessage("Using", value.Name)
		} else {
			subCommands[key] = value
		}
	}
}

func (p *Parser) parseCmd(gadget *Gadget, commands *Commands, subCommands SubCommands) {
	for _, arg := range p.args {
		if value, exists := gadget.SubCommands[arg]; exists {
			if value.Override {
				*commands = value.Commands
				showMessage("Using", value.Name)
			} else {
				subCommands[arg] = value
			}
		}
	}
}
