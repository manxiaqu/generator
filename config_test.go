package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPath = "./test/config.yaml"

func TestLoadConfig(t *testing.T) {
	config := MustLoadConfig(testPath)

	assert.Equal(t, "truffle", config.TruffleProject, "truffle path incorrect")
	assert.Equal(t, 2, len(config.DstLang), "length of dst incorrect")

	godst := config.DstLang[0]
	assert.Equal(t, "go", godst.Name, "go name incorrect")
	assert.Equal(t, "abigen", godst.Tool, "go lang tool incorrect")
	assert.Equal(t, "go", godst.Output, "go output incorrect")
	assert.Equal(t, "gopackage", godst.Package, "go package incorrect")

	javadst := config.DstLang[1]
	assert.Equal(t, "java", javadst.Name, "java name incorrect")
	assert.Equal(t, "web3j", javadst.Tool, "java tool incorrect")
	assert.Equal(t, "java", javadst.Output, "java output incorrect")
	assert.Equal(t, "javapackage", javadst.Package, "java package incorrect")

	assert.Equal(t, 2, len(config.Contracts), "contracts len incorrect")
	assert.Equal(t, "contract1", config.Contracts[0], "contract 1 incorrect")
	assert.Equal(t, "contract2", config.Contracts[1], "contract 2 incorrect")
}
