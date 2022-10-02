package pkg

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Settings struct {
	ConfigPath   string `toml:"config-path"`   // Path of folder containing toml files
	TemplatePath string `toml:"template-path"` // Path of folder containing templates
}

// Change the config-path value in settings.toml
func (s *Settings) setConfigPath(path string) error {
	s.ConfigPath = path

	file, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	e, err := os.Executable()
	if err != nil {
		return err
	}

	e_path, err := filepath.EvalSymlinks(e)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Dir(e_path)+"/settings.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Change the template-path value in settings.json
func (s *Settings) setTemplatePath(path string) error {
	s.TemplatePath = path

	file, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	e, err := os.Executable()
	if err != nil {
		return err
	}

	e_path, err := filepath.EvalSymlinks(e)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Dir(e_path)+"/settings.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}
