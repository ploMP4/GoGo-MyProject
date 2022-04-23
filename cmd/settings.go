package cmd

import (
	"encoding/json"
	"io/ioutil"
)

type Settings struct {
	ConfigPath string `json:"config-path"` // Path of folder containing json files
}

// Change the config-path value in settings.json
func (s *Settings) setConfigPath(path string) error {
	s.ConfigPath = path

	file, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./settings.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}
