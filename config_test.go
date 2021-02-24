package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPath = "./test/config.yaml"

func TestLoadConfig(t *testing.T) {
	config := MustLoadConfig(testPath)

	assert.Equal(t, config.TruffleProjectPath, "truffle")
	assert.Equal(t, config.AbigenPath, "abigenPath")
	assert.Equal(t, config.Name, "name")
	assert.Equal(t, config.OutDir, "outDir")

	assert.Equal(t, 2, len(config.Contracts))
	assert.Equal(t, config.Contracts[0], "contract1")
	assert.Equal(t, config.Contracts[1], "contract2")
}
