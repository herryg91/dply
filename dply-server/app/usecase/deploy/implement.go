package deploy_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply-server/app/repository"
	affinity_usecase "github.com/herryg91/dply/dply-server/app/usecase/affinity"
	deployment_config_usecase "github.com/herryg91/dply/dply-server/app/usecase/deployment_config"
	envar_usecase "github.com/herryg91/dply/dply-server/app/usecase/envar"
	image_usecase "github.com/herryg91/dply/dply-server/app/usecase/image"
	port_usecase "github.com/herryg91/dply/dply-server/app/usecase/port"
	scale_usecase "github.com/herryg91/dply/dply-server/app/usecase/scale"
	"github.com/herryg91/dply/dply-server/entity"
)

type usecase struct {
	deploy_repo          repository.DeploymentRepository
	k8s_repo             repository.K8sRepository
	image_uc             image_usecase.UseCase
	envar_uc             envar_usecase.UseCase
	scale_uc             scale_usecase.UseCase
	port_uc              port_usecase.UseCase
	affinity_uc          affinity_usecase.UseCase
	deployment_config_uc deployment_config_usecase.UseCase
}

func New(
	deploy_repo repository.DeploymentRepository,
	k8s_repo repository.K8sRepository,
	image_uc image_usecase.UseCase,
	envar_uc envar_usecase.UseCase,
	scale_uc scale_usecase.UseCase,
	port_uc port_usecase.UseCase,
	affinity_uc affinity_usecase.UseCase,
	deployment_config_uc deployment_config_usecase.UseCase,
) UseCase {
	return &usecase{
		deploy_repo:          deploy_repo,
		k8s_repo:             k8s_repo,
		image_uc:             image_uc,
		envar_uc:             envar_uc,
		scale_uc:             scale_uc,
		port_uc:              port_uc,
		affinity_uc:          affinity_uc,
		deployment_config_uc: deployment_config_uc,
	}
}

func (uc *usecase) DeployImage(project string, env string, name string, digest string, createdBy int) error {
	currentImage, err := uc.image_uc.GetByDigest(digest)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	currentEnvar, _ := uc.envar_uc.Get(project, env, name)
	if currentEnvar == nil {
		currentEnvar = entity.Envar{}.DefaultEnvar(env, name)
	}

	currentScale, _ := uc.scale_uc.Get(project, env, name)
	if currentScale == nil {
		currentScale = entity.Scale{}.DefaultScale(project, env, name)
	}

	currentPort, _ := uc.port_uc.Get(project, env, name)
	if currentPort == nil {
		currentPort = entity.Port{}.DefaultPort(env, name)
	}

	currentAffinity, _ := uc.affinity_uc.Get(project, env, name)
	if currentAffinity == nil {
		currentAffinity = entity.Affinity{}.DefaultAffinity(env, name)
	}

	currentDeploymentConfig, _ := uc.deployment_config_uc.Get(project, env, name)
	if currentDeploymentConfig == nil {
		currentDeploymentConfig = &entity.DeploymentConfig{
			Project:        project,
			Env:            env,
			Name:           name,
			LivenessProbe:  nil,
			ReadinessProbe: nil,
			StartupProbe:   nil,
		}
	}

	newDeployment := entity.Deployment{
		Project:          project,
		Env:              env,
		Name:             name,
		DeploymentImage:  *currentImage,
		Envar:            *currentEnvar,
		Port:             *currentPort,
		Scale:            *currentScale,
		Affinity:         *currentAffinity,
		DeploymentConfig: *currentDeploymentConfig,
		CreatedBy:        createdBy,
	}

	// Check if redeploy or new deployment
	deployAction := "new" // choice: new|redeploy
	currentDeployment, err := uc.deploy_repo.Get(project, env, name)
	if err != nil {
		if !errors.Is(err, repository.ErrDeploymentNotFound) {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	}

	if currentDeployment != nil {
		if newDeployment.IsDifferentDeploymentConfig(*currentDeployment) {
			deployAction = "new"
		} else {
			deployAction = "redeploy"
		}
	} else {
		deployAction = "new"
	}

	// Do Deploy Action + K8s action
	if deployAction == "new" {
		err := uc.deploy_repo.Create(newDeployment)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}

		err = uc.deployK8s(newDeployment)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	} else {
		fmt.Println("Nothing change in deployment, just redeploy")
		err = uc.deployK8s(newDeployment)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	}

	return nil
}

func (uc *usecase) Redeploy(project string, env string, serviceName string, createdBy int) error {
	currentDeploy, err := uc.deploy_repo.Get(project, env, serviceName)
	if err != nil {
		if errors.Is(err, repository.ErrDeploymentNotFound) {
			return fmt.Errorf("%w: %v", ErrUnexpected, "No deployment, cannot redeploy")
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return uc.DeployImage(project, env, serviceName, currentDeploy.DeploymentImage.Digest, createdBy)
}
