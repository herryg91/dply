package entity

import (
	"encoding/json"
	"strconv"
)

type Deployment struct {
	Id              int      `json:"id"`
	Project         string   `json:"project"`
	Env             string   `json:"env"`
	Name            string   `json:"name"`
	DeploymentImage Image    `json:"image_detail"`
	Envar           Envar    `json:"variables"`
	Port            Port     `json:"ports"`
	Scale           Scale    `json:"scale"`
	Affinity        Affinity `json:"affinity"`
	CreatedBy       int      `json_name:"created_by"`
}

func (d1 *Deployment) IsDifferentDeploymentConfig(d2 Deployment) bool {
	// check image
	if d1.DeploymentImage.Digest != d2.DeploymentImage.Digest {
		return true
	}
	// Check Config
	var1, _ := json.Marshal(&d1.Envar.Variables)
	var2, _ := json.Marshal(&d2.Envar.Variables)
	if string(var1) != string(var2) {
		return true
	}

	port1, _ := json.Marshal(&d1.Port.Ports)
	port2, _ := json.Marshal(&d2.Port.Ports)
	if string(port1) != string(port2) {
		return true
	}

	return false
}

type DeploymentLabel struct {
	Key   string
	Value string
}

func (d *Deployment) GetK8sLabel(additionalLabel ...DeploymentLabel) map[string]string {
	out := map[string]string{
		"app":                          d.Name,
		"env":                          d.Env,
		"repository":                   d.DeploymentImage.Repository,
		"app.kubernetes.io/name":       d.Name,
		"app.kubernetes.io/managed-by": "dply",
		"app.kubernetes.io/created-by": strconv.Itoa(d.CreatedBy),
	}
	for _, addLabel := range additionalLabel {
		out[addLabel.Key] = addLabel.Value
	}
	return out
}
