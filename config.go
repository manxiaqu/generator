package generator

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

const truffleBuild = "build/contracts"

// Config contains all necessary info to generate go/java code by truffle compiled json files.
type Config struct {
	// Path of truffle project.
	TruffleProject string

	// Destination language.
	DstLang []Lang

	// contracts.
	Contracts []string
}

// Lang is the language of output dst from .sol.
type Lang struct {
	// Name of lang.
	Name string

	// path of tool to generate language code.
	Tool string

	// path of output.
	Output string

	// package name.
	Package string
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
