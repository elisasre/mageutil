package docs

import (
	_ "embed"
	"encoding/json"
)

//go:embed swagger.json
var specBytes []byte

func LoadSpec() (*json.RawMessage, error) {
	spec := &json.RawMessage{}
	err := json.Unmarshal(specBytes, spec)
	if err != nil {
		return nil, err
	}
	return spec, nil
}
