package parser

import (
	"github.com/ghodss/yaml"
	types "github.com/owainlewis/frequency/pkg/types"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseManifest(data []byte) (*types.Manifest, error) {
	var manifest *types.Manifest
	err := yaml.Unmarshal(data, &manifest)
	if err != nil {
		return nil, err
	}

	return manifest, nil
}
