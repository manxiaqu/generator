package main

const truffleBuild = "build/contracts"

type GenConfig struct {
	// Tool path
	TruffleProjectPath string
	ABIGenPath         string
	Web3jPath          string

	// Output path
	JavaOutput, GoOutPut   string
	JavaPackage, GoPackage string
}
