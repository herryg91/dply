package entity

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type User struct {
	Name     string `json:"name"`
	Usertype string `json:"usertype"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (User) FromFile() *User {
	homeDir, _ := os.UserHomeDir()
	dplyDir := homeDir + "/.dply"
	userFileLoc := dplyDir + "/user.json"
	if _, err := os.Stat(userFileLoc); os.IsNotExist(err) {
		return nil
	}

	resp := &User{}
	userFileData, _ := ioutil.ReadFile(userFileLoc)
	err := json.Unmarshal(userFileData, &resp)
	if err != nil {
		return nil
	}

	return resp
}
