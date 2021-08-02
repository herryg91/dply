package main

import (
	"fmt"
	"log"
	"time"

	"github.com/herryg91/cdd/grst"
	"github.com/herryg91/cdd/grst/builtin/validationrule"
	loggerInterceptor "github.com/herryg91/cdd/grst/interceptor/logger"
	recoveryInterceptor "github.com/herryg91/cdd/grst/interceptor/recovery"
	sessionInterceptor "github.com/herryg91/cdd/grst/interceptor/session"
	"github.com/herryg91/dply/dply-server/app/repository"
	affinity_usecase "github.com/herryg91/dply/dply-server/app/usecase/affinity"
	deploy_usecase "github.com/herryg91/dply/dply-server/app/usecase/deploy"
	envar_usecase "github.com/herryg91/dply/dply-server/app/usecase/envar"
	image_usecase "github.com/herryg91/dply/dply-server/app/usecase/image"
	port_usecase "github.com/herryg91/dply/dply-server/app/usecase/port"
	scale_usecase "github.com/herryg91/dply/dply-server/app/usecase/scale"
	user_usecase "github.com/herryg91/dply/dply-server/app/usecase/user"
	"github.com/herryg91/dply/dply-server/config"
	"github.com/herryg91/dply/dply-server/entity"
	"github.com/herryg91/dply/dply-server/handler"
	pbDeploy "github.com/herryg91/dply/dply-server/handler/grst/deploy"
	pbImage "github.com/herryg91/dply/dply-server/handler/grst/image"
	pbServer "github.com/herryg91/dply/dply-server/handler/grst/server"
	pbSpec "github.com/herryg91/dply/dply-server/handler/grst/spec"
	pbUser "github.com/herryg91/dply/dply-server/handler/grst/user"
	migration_file "github.com/herryg91/dply/dply-server/migrations"
	"github.com/herryg91/dply/dply-server/pkg/db/mysql"
	"github.com/herryg91/dply/dply-server/pkg/interceptor"
	"github.com/herryg91/dply/dply-server/pkg/password"
	internalvalidationrule "github.com/herryg91/dply/dply-server/pkg/validationrule"
	"github.com/herryg91/dply/dply-server/repository/affinity_repository"
	"github.com/herryg91/dply/dply-server/repository/deployment_repository"
	"github.com/herryg91/dply/dply-server/repository/envar_repository"
	"github.com/herryg91/dply/dply-server/repository/image_repository"
	"github.com/herryg91/dply/dply-server/repository/k8s_repository"
	"github.com/herryg91/dply/dply-server/repository/migration_repository"
	"github.com/herryg91/dply/dply-server/repository/port_repository"
	"github.com/herryg91/dply/dply-server/repository/scale_repository"
	"github.com/herryg91/dply/dply-server/repository/user_repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func initK8sClient(cfg config.Config) (*kubernetes.Clientset, error) {
	var k8sConfig *rest.Config = nil
	if cfg.K8SInCluster {
		var err error
		k8sConfig, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		kubeconfig := cfg.K8SKubeConfig
		k8sConfig, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func main() {
	cfg := config.New()

	db, err := mysql.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBPassword, cfg.DBDatabaseName,
		mysql.SetPrintLog(cfg.DBLogEnable, logger.LogLevel(cfg.DBLogLevel), time.Duration(cfg.DBLogThreshold)*time.Millisecond))
	if err != nil {
		logrus.Panicln("Failed to Initialized mysql DB:", err)
	}

	validationrule.Initialize()
	internalvalidationrule.RegisterJsonRule()

	k8sClient, err := initK8sClient(cfg)
	if err != nil {
		log.Panicln(err)
	}

	scale_repo := scale_repository.New(db)
	scale_uc := scale_usecase.New(scale_repo)

	envar_repo := envar_repository.New(db)
	envar_uc := envar_usecase.New(envar_repo)

	port_repo := port_repository.New(db)
	port_uc := port_usecase.New(port_repo)

	affinity_repo := affinity_repository.New(db)
	affinity_uc := affinity_usecase.New(affinity_repo)

	spec_hndl := handler.NewSpecHandler(envar_uc, scale_uc, port_uc, affinity_uc)

	user_repo := user_repository.New(db)
	user_uc := user_usecase.New(user_repo, password.NewBcryptPassword(cfg.PasswordSalt))
	user_hndl := handler.NewUserHandler(user_uc)

	image_repo := image_repository.New(db)
	image_uc := image_usecase.New(image_repo)
	image_hndl := handler.NewImageHandler(image_uc)

	deploy_repo := deployment_repository.New(db)
	k8s_repo := k8s_repository.New(k8sClient)
	deploy_uc := deploy_usecase.New(deploy_repo, k8s_repo, image_uc, envar_uc, scale_uc, port_uc, affinity_uc)
	deploy_hndl := handler.NewDeployHandler(deploy_uc)

	migration_repo := migration_repository.New(db)
	setupServer(db, migration_repo)

	grstServer, err := grst.NewServer(cfg.GrpcPort, cfg.RestPort, true,
		grst.RegisterGRPCUnaryInterceptor("session", sessionInterceptor.UnaryServerInterceptor()),
		grst.RegisterGRPCUnaryInterceptor("recovery", recoveryInterceptor.UnaryServerInterceptor()),
		grst.RegisterGRPCUnaryInterceptor("log", loggerInterceptor.UnaryServerInterceptor()),
		// grst.RegisterGRPCUnaryInterceptor("login_as_admin", interceptor.LoginAsAdminInterceptor(user_uc, []string{
		// "/image.ImageApi/Remove",
		// "/spec.SpecApi/GetPortTemplate",
		// "/spec.SpecApi/UpdatePortTemplate",
		// "/spec.SpecApi/GetAffinityTemplate",
		// "/spec.SpecApi/UpdateAffinityTemplate",
		// })),
		grst.RegisterGRPCUnaryInterceptor("must_login", interceptor.MustLoginInterceptor(user_uc, []string{
			"/user.UserApi/CheckLogin",
			"/user.UserApi/GetCurrentLogin",
			"/user.UserApi/UpdatePassword",
			"/image.ImageApi/Add",
			"/image.ImageApi/Get",

			"/spec.SpecApi/GetEnvar",
			"/spec.SpecApi/UpsertEnvar",
			"/spec.SpecApi/GetScale",
			"/spec.SpecApi/UpsertScale",
			"/spec.SpecApi/GetPort",
			"/spec.SpecApi/UpsertPort",
			"/spec.SpecApi/GetAffinity",
			"/spec.SpecApi/UpsertAffinity",
			"/deploy.DeployApi/DeployImage",
			"/deploy.DeployApi/Redeploy",
		})),
	)

	if err != nil {
		logrus.Panicln("Failed to Initialize GRPC-REST Server:", err)
	}
	reflection.Register(grstServer.GetGrpcServer())

	pbUser.RegisterUserApiGrstServer(grstServer, user_hndl)
	pbImage.RegisterImageApiGrstServer(grstServer, image_hndl)
	pbDeploy.RegisterDeployApiGrstServer(grstServer, deploy_hndl)
	pbSpec.RegisterSpecApiGrstServer(grstServer, spec_hndl)
	pbServer.RegisterServerApiGrstServer(grstServer, handler.NewServerHandler())
	if err := <-grstServer.ListenAndServeGrst(); err != nil {
		logrus.Panicln("Failed to Run Grpcrest Server:", err)
	}
}

func setupServer(db *gorm.DB, migration_repo repository.MigrationRepository) {
	fmt.Println("1. Check table `migrations`")
	err := setupMigrationTable(migration_repo)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("2. Update migrations")
	err = setupUpdateMigrations(db, migration_repo)
	if err != nil {
		log.Panicln(err)
	}

	// fmt.Println("3. Checking admin account")
	// err = setupAdminAccount(db)
	// if err != nil {
	// 	log.Panicln(err)
	// }
}

func setupMigrationTable(migration_repo repository.MigrationRepository) error {
	exist, err := migration_repo.IsTableExist()
	if err != nil {
		return err
	}

	if exist {
		fmt.Println("OK")
		return nil
	}

	fmt.Println("table `migrations` not found... creating the table...")
	err = migration_repo.CreateTable()
	if err != nil {
		return err
	}

	fmt.Println("table `migrations` succesfully created")
	return nil
}

func setupUpdateMigrations(db *gorm.DB, migration_repo repository.MigrationRepository) error {
	lastMigration, err := migration_repo.GetLast()
	if err != nil {
		return err
	}

	lastMigrationName := ""
	if lastMigration != nil {
		lastMigrationName = lastMigration.Name
	}

	listToMigrate := []migration_file.MigrationFile{
		migration_file.NewMigration0001(db),
	}
	isUpdate := false
	for _, m := range listToMigrate {
		if m.GetName() > lastMigrationName {
			fmt.Println("Migrating `" + m.GetName() + "`...")
			err = m.Up()
			if err != nil {
				return err
			}
			migration_repo.Create(entity.Migration{Name: m.GetName()})
			lastMigrationName = m.GetName()
			isUpdate = true
		}
	}

	if !isUpdate {
		fmt.Println("OK. Already the latest version")
	}

	return nil
}

// func setupAdminAccount(db *gorm.DB) error {
// 	cfg := config.New()
// 	user_repo := user_repository.New(db)
// 	user_uc := user_usecase.New(user_repo, password.NewBcryptPassword(cfg.PasswordSalt))

// 	adminEmail := "admin@dply.com"
// 	adminPassword := helpers.RandomString(16)

// 	alreadyExist := true
// 	_, err := user_repo.GetByEmail(adminEmail)
// 	if err != nil {
// 		if errors.Is(err, repository.ErrUserNotFound) {
// 			alreadyExist = false
// 		} else {
// 			return err
// 		}
// 	}

// 	if alreadyExist {
// 		fmt.Println("OK. No need to generate admin account")
// 		return nil
// 	}

// 	errRegister := user_uc.Register(adminEmail, adminPassword, "admin", "Admin")
// 	if errRegister != nil {
// 		return err
// 	}
// 	fmt.Println(fmt.Sprintf("Admin account created. username: %s | password: %s", adminEmail, adminPassword))

// 	return nil
// }
