module github.com/herryg91/dply/dply-server

go 1.16

require (
	github.com/badoux/checkmail v1.2.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v1.8.5
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.3.0
	github.com/herryg91/cdd/grst v0.0.0-20210412145958-a1a833e40c22
	github.com/herryg91/cdd/protoc-gen-cdd v0.0.0-20210412145958-a1a833e40c22
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/mcuadros/go-defaults v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
	google.golang.org/genproto v0.0.0-20210224155714-063164c882e6
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.25.1-0.20201208041424-160c7477e0e8
	gopkg.in/validator.v2 v2.0.0-20210331031555-b37d688a7fb0
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.9
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
)
