package generator

import (
	"encoding/json"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/inconshreveable/log15"
)

// GenerateCodeByConfigPath automatically generates go files for contract defined in truffle contracts.
func GenerateCodeByConfigPath(configPath string) {
	GenerateCodeByConfig(MustLoadConfig(configPath))
}

// GenerateCodeByConfig automatically generates go files for contract defined in truffle contracts.
func GenerateCodeByConfig(config *Config) {

	// generate bin and abi files.
	for _, contract := range config.Contracts {
		path := filepath.Join(config.TruffleProject, "build", "contracts", contract+".json")

		generateABIAndBIN(path, contract)
		defer Delete(getABI(contract))
		defer Delete(getBIN(contract))

		for _, lang := range config.DstLang {
			generateCode(contract, lang)
		}
	}
}

func generateABIAndBIN(path, name string) error {
	raw, err := Read(path)
	if err != nil {
		log.Error("read json file failed", "err", err, "file", path)
		return err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		log.Error("json unmarshal failed", "err", err)
		return err
	}

	abiRaw, err := json.Marshal(data["abi"])
	if err != nil {
		return err
	}

	binRaw, err := json.Marshal(data["bytecode"])
	if err != nil {
		return err
	}

	// Trim ""
	trimBin := []byte(strings.Trim(string(binRaw), "\""))

	// Write abi and bin file

	abiName := getABI(name)
	binName := getBIN(name)

	if err := Write(abiName, abiRaw); err != nil {
		return err
	}
	return Write(binName, trimBin)
}

func generateCode(contract string, lang Lang) {
	var commandString []string
	switch lang.Name {
	case "java":
		commandString = getJavaCommand(getBIN(contract), getABI(contract), lang.Package, lang.Output)
	case "go":
		commandString = getGoCommand(getBIN(contract), getABI(contract), lang.Package, filepath.Join(lang.Output, getGoName(contract)), contract)
	default:
		panic("not support")
	}

	command := exec.Command(lang.Tool, commandString...)
	if err := command.Run(); err != nil {
		log.Error("generate code failed", "lang", lang.Name, "err", err)
	}
}

func getBIN(contract string) string {
	return strings.ToLower(contract) + ".bin"
}

func getABI(contract string) string {
	return strings.ToLower(contract) + ".abi"
}

func getGoName(contract string) string {
	return strings.ToLower(contract) + ".go"
}

func getGoCommand(binName, abiName, packageName, dst, contract string) []string {
	return []string{
		"--bin",
		binName,
		"--abi",
		abiName,
		"--pkg",
		packageName,
		"--out",
		dst,
		"--type",
		contract,
	}
}

func getJavaCommand(binName, abiName, packageName, dst string) []string {
	return []string{
		"solidity",
		"generate",
		binName,
		abiName,
		"-p",
		packageName,
		"-o",
		dst,
	}
}
