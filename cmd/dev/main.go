package main

import (
	"fmt"
	"io/ioutil"

	"github.com/owainlewis/frequency/pkg/parser"
)

func main() {

	content, err := ioutil.ReadFile("./examples/hello/config.yml")

	if err != nil {
		return
	}

	p := parser.NewParser()
	manifest, err := p.ParseManifest(content)
	if err != nil {
		fmt.Printf("Error parsing %s", err)
		return
	}

	fmt.Printf("Manifest is %v", manifest)
}
