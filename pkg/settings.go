package pkg

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Settings struct {
	ConfigPath   string `toml:"config-path"`   // Path of folder containing toml files
	TemplatePath string `toml:"template-path"` // Path of folder containing templates
}

// Change the config-path value in settings.toml
func (s *Settings) setConfigPath(path string) error {
	buf := new(bytes.Buffer)
	s.ConfigPath = path

	err := toml.NewEncoder(buf).Encode(map[string]string{
		"config-path":   s.ConfigPath,
		"template-path": s.TemplatePath,
	})
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

	err = ioutil.WriteFile(filepath.Dir(e_path)+"/settings.toml", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

// Change the template-path value in settings.toml
func (s *Settings) setTemplatePath(path string) error {
	buf := new(bytes.Buffer)
	s.TemplatePath = path

	err := toml.NewEncoder(buf).Encode(map[string]string{
		"config-path":   s.ConfigPath,
		"template-path": s.TemplatePath,
	})
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

	err = ioutil.WriteFile(filepath.Dir(e_path)+"/settings.toml", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
