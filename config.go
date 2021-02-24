package generator

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

const truffleBuild = "build/contracts"

// Config contains all necessary info to generate go code using truffle compiled json files.
type Config struct {
	// path of abigen tool.
	AbigenPath string `yaml:"abigenPath"`

	// Path of truffle project.
	TruffleProjectPath string `yaml:"truffleProjectPath"`

	// Project name, also used as pkg name.
	Name string `yaml:"name"`

	OutDir string `yaml:"outDir"`

	// contracts.
	Contracts []string `yaml:"contracts"`
}

// MustLoadConfig loads config and panic if err.
func MustLoadConfig(filepath string) *Config {
	c, err := LoadConfig(filepath)
	if err != nil {
		panic(err)
	}

	return c
}

// LoadConfig loads config.
func LoadConfig(filepath string) (*Config, error) {
	yamlConfig, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := Config{}
	if err = yaml.Unmarshal(yamlConfig, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
