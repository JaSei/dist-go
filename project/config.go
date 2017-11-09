package project

import (
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

type Config struct {
	Author  string
	License string
}

func (project project) SaveConfig() error {
	configPath := project.distGoTomlPath()
	config := Config{Author: project.author, License: project.license}

	tomlbytes, err := toml.Marshal(config)
	if err != nil {
		return errors.Wrap(err, "SaveConfig")
	}

	return errors.Wrap(configPath.SpewBytes(tomlbytes), "SaveConfig")
}

func (project *project) LoadConfig() error {
	configBytes, err := project.distGoTomlPath().SlurpBytes()
	if err != nil {
		return errors.Wrap(err, "LoadConfig")
	}

	config := Config{}
	if err = toml.Unmarshal(configBytes, &config); err != nil {
		return errors.Wrap(err, "LoadConfig")
	}

	project.author = config.Author
	project.license = config.License

	return nil
}
