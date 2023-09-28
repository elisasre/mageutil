package docs

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSpecFail(t *testing.T) {
	specBytes = []byte("this is not json")
	data, err := LoadSpec()
	assert.Nil(t, data)
	assert.Error(t, err)
}
