package main

import (
	"fmt"
	"log/slog"

	"gopkg.in/yaml.v2"
)

type Deployment struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

func ConstructDeployment(kind string) *Deployment {
	return &Deployment{
		ApiVersion: "apps/v1",
		Kind:       kind,
	}
}

func main() {
	deployment := ConstructDeployment("Deployment")

	yamlData, err := yaml.Marshal(&deployment)
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Println(string(yamlData))
}
