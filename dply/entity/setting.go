package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Setting struct {
	ServerHostGrpc string `json:"server_host_grpc"`
	Editor         string `json:"editor"`
}

func (Setting) FromDefaultSetting() *Setting {
	return &Setting{
		ServerHostGrpc: "dply-server.dply.svc.cluster.local:9090",
		Editor:         "vi",
	}
}

func (Setting) FromFile() *Setting {
	homeDir, _ := os.UserHomeDir()
	dplyDir := homeDir + "/.dply"
	settingFileLoc := dplyDir + "/setting.json"

	// if ~/.dply not exist, then mkdir ~/.dply
	if _, err := os.Stat(dplyDir); os.IsNotExist(err) {
		os.Mkdir(dplyDir, 0755)
	}

	// if ~/.dply/setting.json not exist, then create default config
	if _, err := os.Stat(settingFileLoc); os.IsNotExist(err) {
		Setting{}.FromDefaultSetting().SaveSetting()
		fmt.Println("setting.json file isn't exist, generate default config file: `" + settingFileLoc + "`")
	}

	if _, err := os.Stat(settingFileLoc); os.IsNotExist(err) {
		return nil
	}

	settingFromFile, _ := ioutil.ReadFile(settingFileLoc)
	s := Setting{}
	_ = json.Unmarshal(settingFromFile, &s)

	needToRewrite := false
	if s.ServerHostGrpc == "" {
		needToRewrite = true
		s.ServerHostGrpc = "dply-server.dply.svc.cluster.local:9090"
	}
	if s.Editor == "" {
		needToRewrite = true
		s.Editor = "vi"
	}
	if needToRewrite {
		s.SaveSetting()
		fmt.Println("setting.json file is broken, repairing to default setting file: `" + settingFileLoc + "`")
	}

	return &s
}

func (s *Setting) SaveSetting() error {
	homeDir, _ := os.UserHomeDir()
	dplyDir := homeDir + "/.dply"
	settingFileLoc := dplyDir + "/setting.json"

	settingJsonMarshalled, _ := json.Marshal(&s)

	err := ioutil.WriteFile(settingFileLoc, settingJsonMarshalled, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
