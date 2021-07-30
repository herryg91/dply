package main

import (
	"fmt"
	"time"

	image_usecase "github.com/herryg91/dply/dply-server/app/usecase/image"
	"github.com/herryg91/dply/dply-server/config"
	"github.com/herryg91/dply/dply-server/pkg/db/mysql"
	"github.com/herryg91/dply/dply-server/repository/image_repository"
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
	image_repo := image_repository.New(db)
	image_uc := image_usecase.New(image_repo)

	// Init Flag
	var digest string

	pflag.StringVarP(&digest, "digest", "d", "", "image digest")
	pflag.Parse()

	if digest == "" {
		logrus.Errorln("`--digest` or `-d` is required")
		return
	}

	// Action
	err = image_uc.Remove(digest)
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}

	fmt.Println("Image successfully deleted")
}
