package main

import (
	"fmt"
	"io/ioutil"

	"bytes"
	"net/http"

	"github.com/ghodss/yaml"
)

func newTaskFromFile(filepath string) ([]byte, error) {
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return yaml.YAMLToJSON(contents)
}

func main() {
	jsonValue, err := newTaskFromFile("examples/yaml/task.yaml")
	if err != nil {
		fmt.Println("Failed to load task")
	}

	resp, err := http.Post("http://localhost:9000/api/v1/tasks", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("Err %s", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("%d: %s", resp.StatusCode, body)
}
