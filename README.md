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
var Config = Config{
        // truffle项目位置
	TruffleProject: "",
	// 需要生成的合约名称
	Contracts []string
	// 对应的语言
	DstLang []{
		// 语言名称
		Name
		// 工具路径
		Tool
		// 输出文件夹
		Output
		// 输出文件所在package
		Package
	},
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
