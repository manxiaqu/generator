package main

import (
	"encoding/json"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
)

const (
	ABI = "abi"
	BIN = "bytecode"
)

// Generate code
func GenerateCode(path, name string, javaFlag, goFlag, delete bool) error {
	raw, err := Read(path)
	if err != nil {
		glog.Error("read file error:", err)
		return err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		glog.Error("unmarshal data error:", err)
		return err
	}

	abiRaw, err := json.Marshal(data[ABI])
	binRaw, err := json.Marshal(data[BIN])
	if err != nil {
		panic(err)
	}

	// Trim ""
	trimBin := []byte(strings.Trim(string(binRaw), "\""))

	// Write abi and bin file
	abiName := name + ".abi"
	binName := name + ".bin"
	// Ignore error.
	Write(abiName, abiRaw)
	Write(binName, trimBin)

	// Generate go code.
	if goFlag {
		// Golang always lower.
		outName := filepath.Join(Config.GoOutPut, strings.ToLower(name)+".go")
		command := exec.Command(Config.ABIGenPath, "--bin", binName, "--abi", abiName, "--pkg", Config.GoPackage, "--out", outName)
		if err = command.Run(); err != nil {
			return err
		}
	}

	// Generate java code.
	if javaFlag {
		command := exec.Command(Config.Web3jPath, "solidity", "generate", binName, abiName, "-p", Config.JavaPackage, "-o", Config.JavaOutput)
		if err = command.Run(); err != nil {
			return err
		}
	}

	// Delete tmp abi/bin files.
	if delete {
		// Ignore error.
		Delete(abiName)
		Delete(binName)
	}

	return nil
}
