package main

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	deployment := NewDeployment(
		"Deployment",
		Metadata{})

	yamlData, err := yaml.Marshal(&deployment)
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Println(string(yamlData))
}
func ParseValues(filepath string) Values {
	data, err := os.ReadFile(filepath)
	if err != nil {
		slog.Error("error reading file: %v", err)
	}
	values := Values{}
	err = yaml.Unmarshal(data, &values)
	if err != nil {
		slog.Error("error unmarshaling YAML: %v", err)
	}
	return values
}
