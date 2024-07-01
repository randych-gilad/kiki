package main

type Deployment struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

func NewDeployment(kind string) *Deployment {
	return &Deployment{
		ApiVersion: "apps/v1",
		Kind:       kind,
	}
}
