# Java/Go 代码生成器
基于[truffle](https://github.com/trufflesuite/ganache-cli)工具，自动读取其生成的json文件，
利用[abigen](https://github.com/ethereum/go-ethereum)和[web3j](https://github.com/web3j/web3j/releases)
生成对应的go和java代码。

```bash
// 生成json文件
truffle compile
```

# 使用方法
在main文件中设置好相关配置：
```go
var Config = GenConfig{
    // truffle项目位置
	TruffleProjectPath: "",
	// abigen工具位置
	ABIGenPath:         "/home/ubuntu/gopath/src/github.com/ethereum/go-ethereum/build/bin/abigen",
	// web3j工具位置
	Web3jPath:          "/home/ubuntu/web3j-3.6.0/bin/web3j",
	// java输出目录
	JavaOutput:         "./",
	// go输出目录
	GoOutPut:           "./",
	// java包名称
	JavaPackage:        "o",
	// go包名称（go文件名称会转为小写）
	GoPackage:          "o",
}

// 文件名字未填写在这的代码不会生成
var FilesFilter = map[string]bool{
	"AccessControl": true,
}
```

# 生成文件

生成可执行文件
```go
go build
```


直接运行
```go
go run *.go
```
