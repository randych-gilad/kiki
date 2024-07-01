package main

type Deployment struct {
	ApiVersion           string                       `yaml:"apiVersion"`
	Kind                 string                       `yaml:"kind"`
	Metadata             Metadata                     `yaml:"metadata"`
	Replicas             int                          `yaml:"replicas"`
	RevisionHistoryLimit int                          `yaml:"revisionHistoryLimit,omitempty"`
	Strategy             map[string]string            `yaml:"strategy"`
	ImagePullSecret      string                       `yaml:"imagePullSecret,omitempty"`
	NodeSelector         string                       `yaml:"nodeSelector,omitempty"`
	VolumeMounts         map[string]string            `yaml:"volumeMounts,omitempty"`
	Image                string                       `yaml:"image"`
	Repo                 string                       `yaml:"-"`
	Version              string                       `yaml:"-"`
	Port                 string                       `yaml:"port"`
	Resources            map[string]map[string]string `yaml:"resources,omitempty"`
	Probe                map[string]string            `yaml:"-"`
	EnvSecretLocal       []string                     `yaml:"-"`
	EnvSecretGlobal      map[string]string            `yaml:"-"`
	EnvSecretSpecific    []string                     `yaml:"-"`
	EnvConfigmap         []string                     `yaml:"-"`
	EnvPlain             map[string]string            `yaml:"-"`
}
type Metadata struct {
	Name        string            `yaml:"name"`
	Namespace   string            `yaml:"namespace"`
	Labels      map[string]string `yaml:"labels"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
}

func NewDeployment(kind string) *Deployment {
	return &Deployment{
		ApiVersion: "apps/v1",
		Kind:       kind,
		Metadata:   Metadata{Name: "kekus-test", Labels: map[string]string{"app": "kekus-test"}, Namespace: "kekus"},
	}
}
