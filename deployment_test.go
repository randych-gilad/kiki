package main

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestNewDeployment(t *testing.T) {
	tests := []struct {
		kind string
		want *Deployment
	}{
		{
			kind: "Deployment",
			want: &Deployment{
				ApiVersion: "apps/v1",
				Kind:       "Deployment",
			},
		},
		{
			kind: "Service",
			want: &Deployment{
				ApiVersion: "apps/v1",
				Kind:       "Service",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			got := NewDeployment(tt.kind)
			if got.ApiVersion != tt.want.ApiVersion || got.Kind != tt.want.Kind {
				t.Errorf("DeploymentNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshalYAML(t *testing.T) {
	deployment := NewDeployment("Deployment")

	yamlData, err := yaml.Marshal(&deployment)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	expectedYAML := "apiVersion: apps/v1\nkind: Deployment\n"
	if string(yamlData) != expectedYAML {
		t.Errorf("yaml.Marshal() = %s, want %s", yamlData, expectedYAML)
	}
}
