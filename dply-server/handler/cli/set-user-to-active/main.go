package main

import (
	"fmt"
	"time"

	"github.com/badoux/checkmail"
	user_usecase "github.com/herryg91/dply/dply-server/app/usecase/user"
	"github.com/herryg91/dply/dply-server/config"
	"github.com/herryg91/dply/dply-server/pkg/db/mysql"
	password_svc "github.com/herryg91/dply/dply-server/pkg/password"
	"github.com/herryg91/dply/dply-server/repository/user_repository"
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
	user_repo := user_repository.New(db)
	user_uc := user_usecase.New(user_repo, password_svc.NewBcryptPassword(cfg.PasswordSalt))

	// Init Flag
	var email string

	pflag.StringVarP(&email, "email", "e", "", "enter email")
	pflag.Parse()

	if email == "" {
		logrus.Errorln("`--email` or `-e` is required")
		return
	} else if err := checkmail.ValidateFormat(email); err != nil {
		logrus.Errorln("`--email` or `-e` is not an email format, got: " + email)
		return
	}

	// Action
	err = user_uc.EditStatusToActive(email)
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}

	fmt.Println("User has been set to active")
}
