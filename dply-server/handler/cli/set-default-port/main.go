package main

import (
	"encoding/json"
	"fmt"
	"time"

	port_usecase "github.com/herryg91/dply/dply-server/app/usecase/port"
	"github.com/herryg91/dply/dply-server/config"
	"github.com/herryg91/dply/dply-server/entity"
	"github.com/herryg91/dply/dply-server/pkg/db/mysql"
	"github.com/herryg91/dply/dply-server/pkg/editor"
	"github.com/herryg91/dply/dply-server/repository/port_repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"gorm.io/gorm/logger"
)

func main() {
	// Init DB + Usecase
	cfg := config.New()
	db, err := mysql.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName,
		mysql.SetPrintLog(cfg.DBLogEnable, logger.LogLevel(1), time.Duration(1000)*time.Millisecond))
	if err != nil {
		logrus.Panicln("Failed to Initialized mysql DB:", err)
	}

	// Init Flag
	var editorApp string

	pflag.StringVarP(&editorApp, "editor", "e", "vi", "pick your editor (default:vi)")
	pflag.Parse()

	if editorApp == "" {
		logrus.Errorln("`--editor` or `-e` is required")
		return
	}

	port_repo := port_repository.New(db)
	port_uc := port_usecase.New(port_repo)

	currentTemplate, err := port_uc.GetTemplate("default")
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}
	currentTemplateIndented, _ := json.MarshalIndent(currentTemplate.Ports, "", "    ")

	// Get Updated Data via Editor
	updatedData, err := editor.Open(editorApp, "tmp_port_edit", currentTemplateIndented)
	if err != nil {
		logrus.Errorln(fmt.Errorf("Unexpected error: %v", "Error on editor: "+err.Error()))
		return
	}

	if string(currentTemplateIndented) == string(updatedData) {
		fmt.Println("Nothing to changed")
		return
	}

	data := entity.PortTemplate{
		TemplateName: "default",
		Ports:        []entity.PortSpec{},
	}
	err = json.Unmarshal(updatedData, &data.Ports)
	if err != nil {
		logrus.Errorln(fmt.Errorf("Unexpected error: %v", "Error unmarshal: "+string(updatedData)))
		return
	}

	err = port_uc.UpsertTemplate(data)
	if err != nil {
		logrus.Errorln(err.Error())
	}

	fmt.Println("Default port succesfully updated")
}
