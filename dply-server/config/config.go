package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME" default:"dply-server"`
	Environment string `envconfig:"ENVIRONMENT" default:"dev"`
	Maintenance bool   `envconfig:"MAINTENANCE" default:"false"`
	RestPort    int    `envconfig:"REST_PORT" default:"50080" required:"true"`
	GrpcPort    int    `envconfig:"GRPC_PORT" default:"50090" required:"true"`

	DBHost         string `envconfig:"DB_HOST" default:"localhost"`
	DBPort         int    `envconfig:"DB_PORT" default:"3306"`
	DBUserName     string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:"password"`
	DBDatabaseName string `envconfig:"DB_DBNAME" default:"dply"`
	DBLogEnable    bool   `envconfig:"DB_LOG_ENABLE" default:"true"`
	DBLogLevel     int    `envconfig:"DB_LOG_LEVEL" default:"3"`
	DBLogThreshold int    `envconfig:"DB_LOG_THRESHOLD" default:"100"`

	PasswordSalt string `envconfig:"PASSWORD_SALT" default:"bMWLKgSqIUhJVdnE"`

	K8SInCluster  bool   `envconfig:"K8S_IN_CLUSTER" default:"false"`
	K8SKubeConfig string `envconfig:"K8S_KUBE_CONFIG" default:"~/.kube/config"`
}

func New() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	// cfg.K8SKubeConfig = "/Users/hg/.kube/config"
	// cfg.DBHost = "dply-mysql.dply.svc.cluster.local"
	// cfg.DBPassword = "mysql@dply"
	return cfg
}
