package main

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	v, err := ValuesParse("EXAMPLE.yaml")
	v.Assert()
	if err != nil {
		slog.Error(err.Error())
		return
	}
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

func ValuesParse(filepath string) (*Values, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		slog.Error("error reading file: %v", err)
		return nil, err
	}
	values := &Values{}
	err = yaml.Unmarshal(data, &values)
	if err != nil {
		slog.Error("error unmarshaling YAML: %v", err)
		return nil, err
	}
	return values, nil
}

func (v *Values) Assert() error {
	return nil
}

func (d *Deployment) Default(v *Values) error {
	slog.Info("Validate and default values against defined rules...")

	if v.Deployment.Kind == "" {
		slog.Warn("Deployment.Kind not specified, defaulting to Deployment")
		d.Kind = "Deployment"
	} else if v.Deployment.Kind != "Deployment" && v.Deployment.Kind != "DaemonSet" {
		return fmt.Errorf("invalid Deployment.Kind, got: %w, want: Deployment or DaemonSet", v.Deployment.Kind)
	}
	if v.Deployment.Annotations == nil {
		slog.Warn("Deployment.Annotations not specified")
	} else {
		d.Metadata.Annotations = v.Deployment.Annotations
	}
	if v.Deployment.Replicas == 0 {
		slog.Warn("Deployment.Replicas not specified, defaulting to 1")
		d.Spec.Replicas = 1
	} else {
		d.Spec.Replicas = v.Deployment.Replicas
	}
	if v.Deployment.Kind == "DaemonSet" && v.Deployment.Strategy == nil {
		return fmt.Errorf("invalid Deployment.Strategy, cannot assign default for DaemonSet")
	} else {
		d.Spec.Strategy = map[string]string{"type": "RollingUpdate"}
	}

	return nil
}
