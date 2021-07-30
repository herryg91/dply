package editor

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type EditorApp string

const (
	Editor_Vi   EditorApp = "vi"
	Editor_Nano EditorApp = "nano"
)

func Open(editorApp EditorApp, tmpFileName string, initData []byte) ([]byte, error) {
	if tmpFileName == "" {
		tmpFileName = "dply_tmp_" + strconv.FormatInt(time.Now().Unix(), 10)
	}
	tmpFile, err := ioutil.TempFile("", "scaling_edit")
	if err != nil {
		return initData, err
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Write(initData)
	tmpFile.Close()

	// open editor via terminal cmd
	termCmd := exec.Command(string(editorApp), tmpFile.Name())
	termCmd.Stdin = os.Stdin
	termCmd.Stdout = os.Stdout
	termCmd.Stderr = os.Stderr
	if err := termCmd.Run(); err != nil {
		return initData, err
	}

	return ioutil.ReadFile(tmpFile.Name())
}
