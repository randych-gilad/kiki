package main

type Values struct {
	Deployment struct {
		Version              string            `yaml:"version"`
		Kind                 string            `yaml:"kind"`
		Annotations          map[string]string `yaml:"annotations"`
		Replicas             int               `yaml:"replicas"`
		RevisionHistoryLimit int               `yaml:"revisionHistoryLimit"`
		Strategy             map[string]string `yaml:"strategy"`
		VolumeMounts         map[string]string `yaml:"volumeMounts"`
		Image                string            `yaml:"image"`
		Labels               map[string]string `yaml:"labels"`
		Port                 int               `yaml:"port"`
		Limits               struct {
			CPU    string `yaml:"cpu"`
			Memory string `yaml:"memory"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			CPU    string `yaml:"cpu"`
			Memory string `yaml:"memory"`
		} `yaml:"requests,omitempty"`
		NodeSelector string `yaml:"nodeSelector"`
		Probe        struct {
			Readiness struct {
				Path             string `yaml:"path"`
				FailureThreshold int    `yaml:"failureThreshold"`
				PeriodSeconds    int    `yaml:"periodSeconds"`
			} `yaml:"readiness,omitempty"`
			Liveness struct {
				Path             string `yaml:"path"`
				FailureThreshold int    `yaml:"failureThreshold"`
				PeriodSeconds    int    `yaml:"periodSeconds"`
			} `yaml:"liveness,omitempty"`
			Startup struct {
				Path                string `yaml:"path"`
				FailureThreshold    int    `yaml:"failureThreshold"`
				PeriodSeconds       int    `yaml:"periodSeconds"`
				InitialDelaySeconds int    `yaml:"initialDelaySeconds"`
			} `yaml:"startup,omitempty"`
		} `yaml:"probe"`
		Env struct {
			Secret struct {
				Local    []string          `yaml:"local"`
				Global   map[string]string `yaml:"global"`
				Specific []string          `yaml:"specific"`
			} `yaml:"secret"`
			Configmap []string          `yaml:"configmap"`
			Plain     map[string]string `yaml:"plain"`
		} `yaml:"env"`
	} `yaml:"Deployment"`
}
