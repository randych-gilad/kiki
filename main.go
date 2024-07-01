package main

import (
	"fmt"
	"log/slog"

	"gopkg.in/yaml.v2"
)

func main() {
	deployment := NewDeployment(
		"Deployment")

	yamlData, err := yaml.Marshal(&deployment)
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Println(string(yamlData))
}
