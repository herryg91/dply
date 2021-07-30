package handler

import (
	affinity_usecase "github.com/herryg91/dply/dply-server/app/usecase/affinity"
	envar_usecase "github.com/herryg91/dply/dply-server/app/usecase/envar"
	port_usecase "github.com/herryg91/dply/dply-server/app/usecase/port"
	scale_usecase "github.com/herryg91/dply/dply-server/app/usecase/scale"
	pbSpec "github.com/herryg91/dply/dply-server/handler/grst/spec"
)

type specConfig struct {
	envar_uc    envar_usecase.UseCase
	scale_uc    scale_usecase.UseCase
	port_uc     port_usecase.UseCase
	affinity_uc affinity_usecase.UseCase
	pbSpec.UnimplementedSpecApiServer
}

func NewSpecHandler(envar_uc envar_usecase.UseCase, scale_uc scale_usecase.UseCase, port_uc port_usecase.UseCase, affinity_uc affinity_usecase.UseCase) pbSpec.SpecApiServer {
	return &specConfig{
		envar_uc:    envar_uc,
		scale_uc:    scale_uc,
		port_uc:     port_uc,
		affinity_uc: affinity_uc,
	}
}
