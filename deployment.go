package main

type Deployment struct {
	ApiVersion           string            `yaml:"apiVersion"`
	Kind                 string            `yaml:"kind"`
	Metadata             Metadata          `yaml:"metadata"`
	Replicas             int               `yaml:"replicas"`
	RevisionHistoryLimit int               `yaml:"revisionHistoryLimit"`
	Strategy             map[string]string `yaml:"strategy"`
	ImagePullSecret      string            `yaml:"imagePullSecret"`
	NodeSelector         string            `yaml:"nodeSelector"`
	VolumeMounts         map[string]string `yaml:"volumeMounts"`
	Image                string            `yaml:"image"`
	Repo                 string
	Version              string
	Port                 string                        `yaml:"port"`
	Resources            *map[string]map[string]string `yaml:"resources"`
	Probe                map[string]string
	EnvSecretLocal       []string
	EnvSecretGlobal      map[string]string
	EnvSecretSpecific    []string
	EnvConfigmap         []string
	EnvPlain             map[string]string
}
type Metadata struct {
	Name        string            `yaml:"name"`
	Namespace   string            `yaml:"namespace"`
	Labels      map[string]string `yaml:"labels"`
	Annotations map[string]string `yaml:"annotations"`
}

func NewDeployment(kind string) *Deployment {
	return &Deployment{
		ApiVersion: "apps/v1",
		Kind:       kind,
	}
}
