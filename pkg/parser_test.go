package pkg

import (
	"testing"
)

var p = Parser{args: []string{"v"}}

func TestParser_parseSettings(t *testing.T) {
	err := p.parseSettings()
	if err == nil {
		t.Error("no error thrown with not existant settings.json file")
	}

	// err = p.parseSettings()
	// if err != nil {
	// 	t.Error(err)
	// }
}

func TestParser_parseConfig(t *testing.T) {
	p.settings.ConfigPath = "./testdata/config"

	err := p.parseConfig("cpp")
	if err != nil {
		t.Error(err)
	}
}

func TestParser_getHelp(t *testing.T) {
	helpCommands := p.getHelp()
	if helpCommands == nil {
		t.Error("help menu wasn't returned with existant commands")
	}

	p2 := p
	p2.settings.ConfigPath = "non_existant"

	helpCommands = p2.getHelp()
	if helpCommands != nil {
		t.Error("")
	}
}

func TestParser_getSubHelp(t *testing.T) {
	_, err := p.getSubHelp("cpp")
	if err != nil {
		t.Error(err)
	}

	_, err = p.getSubHelp("non_existant")
	if err == nil {
		t.Error("non existant config file name didn't throw an error")
	}
}

func TestParser_parseArgs(t *testing.T) {
	mainCommands, otherCommands, dirs := p.parseArgs()
	if mainCommands == nil {
		t.Error("no main commands where returned")
	}

	if otherCommands == nil {
		t.Error("no subcommands where returned")
	}

	if dirs == nil {
		t.Error("no dirs where returned")
	}
}
