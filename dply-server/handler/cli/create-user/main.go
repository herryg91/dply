package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	user_usecase "github.com/herryg91/dply/dply-server/app/usecase/user"
	"github.com/herryg91/dply/dply-server/config"
	"github.com/herryg91/dply/dply-server/entity"
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
	var password string
	var name string
	var usertype string = "user"

	pflag.StringVarP(&email, "email", "e", "", "enter email")
	pflag.StringVarP(&password, "password", "p", "", "enter password")
	pflag.StringVarP(&name, "name", "n", "", "enter name")
	// pflag.StringVarP(&usertype, "type", "t", "user", "choice: admin|user")
	pflag.Parse()
	usertype = strings.ToLower(usertype)

	if email == "" {
		logrus.Errorln("`--email` or `-e` is required")
		return
	} else if err := checkmail.ValidateFormat(email); err != nil {
		logrus.Errorln("`--email` or `-e` is not an email format, got: " + email)
		return
	} else if password == "" {
		logrus.Errorln("`--password` or `-p` is required")
		return
	} else if name == "" {
		logrus.Errorln("`--name` or `-n` is required")
		return
	}

	// Action
	err = user_uc.Register(email, password, entity.UserType(usertype), name)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	fmt.Println("User is succesfully created")
}
