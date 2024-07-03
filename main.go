package main

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	v := ValuesParse("EXAMPLE.yaml")
	if err := v.Default(); err != nil {
		slog.Error(err.Error())
		return
	}
	// fmt.Printf("%#v\n", v)
	deployment := NewDeployment(
		"Deployment",
		Metadata{})
	yamlData, err := yaml.Marshal(&deployment)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	fmt.Println(string(yamlData))
}

func ValuesParse(filepath string) Values {
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

func (v *Values) Default() error {
	slog.Info("Validate and default values against defined rules...")

	if v.Deployment.Kind == "" {
		slog.Warn("Deployment.Kind not specified, defaulting to Deployment")
		v.Deployment.Kind = "Deployment"
	} else if v.Deployment.Kind != "Deployment" && v.Deployment.Kind != "DaemonSet" {
		return fmt.Errorf("invalid Deployment.Kind, got: %v, want: Deployment or DaemonSet", v.Deployment.Kind)
	}

	if v.Deployment.Annotations == nil {
		slog.Warn("Deployment.Annotations not specified")
	}

	if v.Deployment.Replicas == 0 {
		v.Deployment.Replicas = 1
		slog.Warn("Deployment.Replicas not specified, defaulting to 1")
	}
	if v.Deployment.Kind == "DaemonSet" && v.Deployment.Strategy == nil {
		return fmt.Errorf("")
	}
	if v.Deployment.Strategy == nil {
		v.Deployment.Strategy = map[string]string{"type": "RollingUpdate"}
	}

	return nil
}
