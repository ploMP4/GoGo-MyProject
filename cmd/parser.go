package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Parser struct {
	json Json
	args []string
}

type MainCommmands [][]string
type SubCommands map[string]SubCommand

type Json struct {
	Commands    MainCommmands `json:"commands"`
	Dirs        []string      `json:"dirs"`
	SubCommands SubCommands   `json:"subCommands"`
	Help        string        `json:"help"`
}

type SubCommand struct {
	Name     string   `json:"name"`
	Command  []string `json:"command"`
	Override bool     `json:"override"`
	Parallel bool     `json:"parallel"`
	Help     string   `json:"help"`
}

func (p *Parser) parseJson(filename string) error {
	jsonFile, err := os.Open("./config/" + filename + ".json")
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
