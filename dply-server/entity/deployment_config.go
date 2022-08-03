package entity

type HttpGetProbe struct {
	Path                string `json:"path"`
	Port                int    `json:"port"`
	FailureThreshold    int    `json:"failure_threshold"`
	PeriodSeconds       int    `json:"period_seconds"`
	InitialDelaySeconds int    `json:"initial_delay_seconds"`
}

type DeploymentConfig struct {
	Id             int           `json:"id"`
	Project        string        `json:"project"`
	Env            string        `json:"env"`
	Name           string        `json:"name"`
	LivenessProbe  *HttpGetProbe `json:"liveness_probe"`
	ReadinessProbe *HttpGetProbe `json:"readiness_probe"`
	StartupProbe   *HttpGetProbe `json:"startup_probe"`
	CreatedBy      int           `json:"created_by"`
}
