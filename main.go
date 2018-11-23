package main

import (
	"path/filepath"
	"strings"

	"github.com/golang/glog"
)

var Config = GenConfig{
	TruffleProjectPath: "",
	ABIGenPath:         "/home/ubuntu/gopath/src/github.com/ethereum/go-ethereum/build/bin/abigen",
	Web3jPath:          "/home/ubuntu/web3j-3.6.0/bin/web3j",
	JavaOutput:         "./",
	GoOutPut:           "./",
	JavaPackage:        "o",
	GoPackage:          "o",
}

var FilesFilter = map[string]bool{
	"AccessControl": true,
}

func main() {
	// Get contract path
	contractsPath := filepath.Join(Config.TruffleProjectPath, truffleBuild)
	files, err := GetDirFiles(contractsPath)
	if err != nil {
		glog.Error(err)
	}

	// Generate code
	for _, f := range files {
		name := strings.Trim(f, ".json")

		// Filter files.
		if generate, ok := FilesFilter[name]; !ok || !generate {
			continue
		}

		if err := GenerateCode(filepath.Join(contractsPath, f), name, true, true, true); err != nil {
			glog.Error(err)
		}
	}
}
