package main

type Deployment struct {
	ApiVersion        string            `yaml:"apiVersion"`
	Kind              string            `yaml:"kind"`
	Metadata          Metadata          `yaml:"metadata"`
	Spec              SpecDeployment    `yaml:"spec"`
	Image             string            `yaml:"image"`
	Repo              string            `yaml:"-"`
	Version           string            `yaml:"-"`
	Port              int               `yaml:"port"`
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
	Strategy             map[string]string            `yaml:"strategy"`
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
	Probes       []Probe                      `yaml:"-"`
	Env          []Env                        `yaml:"env,omitempty"`
}
type Probe struct {
	HttpGet             HttpGet `yaml:"httpGet"`
	FailureThreshold    int     `yaml:"failureThreshold,omitempty"`
	PeriodSeconds       int     `yaml:"periodSeconds,omitempty"`
	InitialDelaySeconds int     `yaml:"initialDelaySeconds,omitempty"`
}
type HttpGet struct {
	Path string
	Port string
}
type EnvPlain struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}
type EnvKeyRef struct {
	Name      string    `yaml:"name"`
	ValueFrom ValueFrom `yaml:"valueFrom"`
}
type ValueFrom struct {
	SecretKeyRef    KeyRef `yaml:"secretKeyRef,omitempty"`
	ConfigMapKeyRef KeyRef `yaml:"configMapKeyRef,omitempty"`
}
type KeyRef struct {
	Name string `yaml:"name"`
	Key  string `yaml:"key"`
}
type Env interface{}

func NewDeployment(kind string, metadata Metadata) *Deployment {
	return &Deployment{
		ApiVersion: "apps/v1",
		Kind:       kind,
		Metadata:   metadata,
	}
}
