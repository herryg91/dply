package entity

type DeploymentConfig struct {
	Project        string        `json:"-"`
	Env            string        `json:"-"`
	Name           string        `json:"-"`
	LivenessProbe  *HttpGetProbe `json:"liveness_probe"`
	ReadinessProbe *HttpGetProbe `json:"readiness_probe"`
	StartupProbe   *HttpGetProbe `json:"startup_probe"`
}

type HttpGetProbe struct {
	Path                string `json:"path"`
	Port                int    `json:"port"`
	FailureThreshold    int    `json:"failure_threshold"`
	PeriodSeconds       int    `json:"period_seconds"`
	InitialDelaySeconds int    `json:"initial_delay_seconds"`
}
