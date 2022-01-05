package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/herryg91/dply/dply/pkg/editor"
)

type Config struct {
	DplyServerHost         string `json:"server_host"` // dply-server.dply.svc.cluster.local:9090
	Editor                 string `json:"editor"`      // vi | nano
	Project                string `json:"project"`     // default: default
	DockerHost             string `json:"docker_host"` // tcp://dind.dply.svc.cluster.local:2376
	DockerVersion          string `json:"docker_version "`
	DockerCertificatesPath string `json:"docker_certs_path"` // ca.pem, cert.pem, key.pem files
	RegistryHost           string `json:"registry_host"`     // https://gcr.io
	RegistryUsername       string `json:"registry_username"`
	RegistryPassword       string `json:"registry_password"`
	RegistryTagPrefix      string `json:"registry_tag_prefix"` // gcr.io/<folder1>/<folder2>
}

func (Config) FromDefaultConfig() *Config {
	homeDir, _ := os.UserHomeDir()
	dplyDir := homeDir + "/.dply"
	return &Config{
		DplyServerHost:         "dply-server.dply.svc.cluster.local:9090",
		Editor:                 "vi",
		Project:                "default",
		DockerHost:             "tcp://dind.dply.svc.cluster.local:2376",
		DockerVersion:          "1.41",
		DockerCertificatesPath: dplyDir + "/docker-certs",
		RegistryHost:           "https://docker.io",
		RegistryUsername:       "username",
		RegistryPassword:       "password",
		RegistryTagPrefix:      "docker.io/dply/dply-image",
	}
}

func (Config) FromFile() *Config {
	homeDir, _ := os.UserHomeDir()
	dplyDir := homeDir + "/.dply"
	configFileLoc := dplyDir + "/config.json"

	// if ~/.dply not exist, then mkdir ~/.dply
	if _, err := os.Stat(dplyDir); os.IsNotExist(err) {
		os.Mkdir(dplyDir, 0755)
	}

	// if ~/.dply/config.json not exist, then create default config
	if _, err := os.Stat(configFileLoc); os.IsNotExist(err) {
		Config{}.FromDefaultConfig().SaveConfig()
		fmt.Println("config.json file isn't exist, generate default config file: `" + configFileLoc + "`")
	}

	if _, err := os.Stat(configFileLoc); os.IsNotExist(err) {
		return nil
	}

	configFromFile, _ := ioutil.ReadFile(configFileLoc)
	s := Config{}
	_ = json.Unmarshal(configFromFile, &s)

	needToRewrite := false
	if s.DplyServerHost == "" {
		needToRewrite = true
		s.DplyServerHost = "dply-server.dply.svc.cluster.local:9090"
	}
	if s.Editor == "" {
		needToRewrite = true
		s.Editor = "vi"
	}
	if needToRewrite {
		s.SaveConfig()
		fmt.Println("config.json file is broken, repairing to default config file: `" + configFileLoc + "`")
	}

	return &s
}

func (s *Config) SaveConfig() error {
	homeDir, _ := os.UserHomeDir()
	dplyDir := homeDir + "/.dply"
	configFileLoc := dplyDir + "/config.json"

	if s.Project == "" {
		fmt.Println("project is required, changed into default")
		s.Project = "default"
	}

	configJsonMarshalled, _ := json.Marshal(&s)

	err := ioutil.WriteFile(configFileLoc, configJsonMarshalled, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

var ErrConfigNothingChange = fmt.Errorf("Nothing to changes")

func (s *Config) UpdateByEditor() error {
	current_data, _ := json.MarshalIndent(s, "", "    ")

	editor_app := editor.EditorApp(s.Editor)
	updated_data, err := editor.Open(editor_app, "tmp_config_edit", current_data)
	if err != nil {
		return fmt.Errorf("Error editor: " + err.Error())
	}

	// if nothing to change
	if string(current_data) == string(updated_data) {
		return ErrConfigNothingChange
	}

	err = json.Unmarshal(updated_data, &s)
	if err != nil {
		return fmt.Errorf("Error unmarshal: " + string(updated_data))
	}

	err = s.SaveConfig()
	if err != nil {
		return fmt.Errorf("Failed to save config: " + err.Error())
	}

	return nil
}

func (s *Config) UpdateProject(project string) error {
	s.Project = project

	err := s.SaveConfig()
	if err != nil {
		return fmt.Errorf("Failed to save config: " + err.Error())
	}

	return nil
}
