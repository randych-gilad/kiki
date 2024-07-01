package main

type Deployment struct {
	ApiVersion        string            `yaml:"apiVersion"`
	Kind              string            `yaml:"kind"`
	Metadata          Metadata          `yaml:"metadata"`
	Strategy          map[string]string `yaml:"strategy"`
	Image             string            `yaml:"image"`
	Repo              string            `yaml:"-"`
	Version           string            `yaml:"-"`
	Port              string            `yaml:"port"`
	EnvSecretLocal    []string          `yaml:"-"`
	EnvSecretGlobal   map[string]string `yaml:"-"`
	EnvSecretSpecific []string          `yaml:"-"`
	EnvConfigmap      []string          `yaml:"-"`
	EnvPlain          map[string]string `yaml:"-"`
}

type Metadata struct {
	Name        string            `yaml:"name"`
	Namespace   string            `yaml:"namespace"`
	Labels      map[string]string `yaml:"labels"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
}
type SpecDeployment struct {
	Replicas             int                          `yaml:"replicas"`
	RevisionHistoryLimit int                          `yaml:"revisionHistoryLimit,omitempty"`
	Selector             map[string]map[string]string `yaml:"selector"`
	Template             Template                     `yaml:"template"`
}
type Template struct {
	Metadata map[string]map[string]string `yaml:"metadata"`
	Spec     SpecTemplate                 `yaml:"spec"`
}
type SpecTemplate struct {
	ImagePullSecrets string              `yaml:"imagePullSecrets,omitempty"`
	NodeSelector     string              `yaml:"nodeSelector,omitempty"`
	Volumes          []map[string]string `yaml:"volumes,omitempty"`
	Containers       []Container         `yaml:"containers"`
}
type Container struct {
	VolumeMounts map[string]string            `yaml:"volumeMounts,omitempty"`
	Resources    map[string]map[string]string `yaml:"resources,omitempty"`
	Probe        map[string]string            `yaml:"-"`
}

func NewDeployment(kind string, metadata Metadata) *Deployment {
	return &Deployment{
		ApiVersion: "apps/v1",
		Kind:       kind,
		Metadata:   metadata,
	}
}
