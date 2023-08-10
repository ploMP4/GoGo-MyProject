package internal

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const PROJECT_ROOT_DIR_NAME = ".gogo"

type Settings struct {
	GadgetPath   string `yaml:"gadget-path"`   // Path of folder containing toml files
	TemplatePath string `yaml:"template-path"` // Path of folder containing templates
}

// Change the gadget-path value in settings.yaml
func (s *Settings) setGadgetPath(path string) error {
	buf := new(bytes.Buffer)
	s.GadgetPath = path

	err := yaml.NewEncoder(buf).Encode(map[string]string{
		"gadget-path":   s.GadgetPath,
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

	err = ioutil.WriteFile(filepath.Dir(e_path)+"/settings.yaml", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

// Change the template-path value in settings.yaml
func (s *Settings) setTemplatePath(path string) error {
	buf := new(bytes.Buffer)
	s.TemplatePath = path

	err := yaml.NewEncoder(buf).Encode(map[string]string{
		"gadget-path":   s.GadgetPath,
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

	err = ioutil.WriteFile(filepath.Dir(e_path)+"/settings.yaml", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
