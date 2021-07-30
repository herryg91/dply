package main

import (
	"encoding/json"
	"fmt"
	"time"

	affinity_usecase "github.com/herryg91/dply/dply-server/app/usecase/affinity"
	"github.com/herryg91/dply/dply-server/config"
	"github.com/herryg91/dply/dply-server/entity"
	"github.com/herryg91/dply/dply-server/pkg/db/mysql"
	"github.com/herryg91/dply/dply-server/pkg/editor"
	"github.com/herryg91/dply/dply-server/repository/affinity_repository"
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

	affinity_repo := affinity_repository.New(db)
	affinity_uc := affinity_usecase.New(affinity_repo)

	currentTemplate, err := affinity_uc.GetTemplate("default")
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}
	currentTemplateMap := map[string]interface{}{
		"node_affinity":     currentTemplate.NodeAffinity,
		"pod_affinity":      currentTemplate.PodAffinity,
		"pod_anti_affinity": currentTemplate.PodAntiAffinity,
	}

	currentTemplateIndented, _ := json.MarshalIndent(currentTemplateMap, "", "    ")

	// Get Updated Data via Editor
	updatedData, err := editor.Open(editorApp, "tmp_affinity_edit", currentTemplateIndented)
	if err != nil {
		logrus.Errorln(fmt.Errorf("Unexpected error: %v", "Error on editor: "+err.Error()))
		return
	}

	if string(currentTemplateIndented) == string(updatedData) {
		fmt.Println("Nothing to changed")
		return
	}

	data := entity.AffinityTemplate{
		TemplateName:    "default",
		NodeAffinity:    []entity.AffinityTerm{},
		PodAffinity:     []entity.AffinityTerm{},
		PodAntiAffinity: []entity.AffinityTerm{},
	}
	err = json.Unmarshal(updatedData, &data)
	if err != nil {
		logrus.Errorln(fmt.Errorf("Unexpected error: %v", "Error unmarshal: "+string(updatedData)))
		return
	}
	data.TemplateName = "default"

	err = affinity_uc.UpsertTemplate(data)
	if err != nil {
		logrus.Errorln(err.Error())
	}

	fmt.Println("Default affinity succesfully updated")
}
