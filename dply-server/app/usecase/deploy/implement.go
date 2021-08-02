package deploy_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply-server/app/repository"
	affinity_usecase "github.com/herryg91/dply/dply-server/app/usecase/affinity"
	envar_usecase "github.com/herryg91/dply/dply-server/app/usecase/envar"
	image_usecase "github.com/herryg91/dply/dply-server/app/usecase/image"
	port_usecase "github.com/herryg91/dply/dply-server/app/usecase/port"
	scale_usecase "github.com/herryg91/dply/dply-server/app/usecase/scale"
	"github.com/herryg91/dply/dply-server/entity"
)

type usecase struct {
	deploy_repo repository.DeploymentRepository
	k8s_repo    repository.K8sRepository
	image_uc    image_usecase.UseCase
	envar_uc    envar_usecase.UseCase
	scale_uc    scale_usecase.UseCase
	port_uc     port_usecase.UseCase
	affinity_uc affinity_usecase.UseCase
}

func New(
	deploy_repo repository.DeploymentRepository,
	k8s_repo repository.K8sRepository,
	image_uc image_usecase.UseCase,
	envar_uc envar_usecase.UseCase,
	scale_uc scale_usecase.UseCase,
	port_uc port_usecase.UseCase,
	affinity_uc affinity_usecase.UseCase,
) UseCase {
	return &usecase{
		deploy_repo: deploy_repo,
		k8s_repo:    k8s_repo,
		image_uc:    image_uc,
		envar_uc:    envar_uc,
		scale_uc:    scale_uc,
		port_uc:     port_uc,
		affinity_uc: affinity_uc,
	}
}

func (uc *usecase) DeployImage(env string, name string, digest string, createdBy int) error {
	currentImage, err := uc.image_uc.GetByDigest(digest)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	currentEnvar, _ := uc.envar_uc.Get(env, name)
	if currentEnvar == nil {
		currentEnvar = entity.Envar{}.DefaultEnvar(env, name)
	}

	currentScale, _ := uc.scale_uc.Get(env, name)
	if currentScale == nil {
		currentScale = entity.Scale{}.DefaultScale(env, name)
	}

	currentPort, _ := uc.port_uc.Get(env, name)
	if currentPort == nil {
		currentPort = entity.Port{}.DefaultPort(env, name)
	}

	currentAffinity, _ := uc.affinity_uc.Get(env, name)
	if currentAffinity == nil {
		currentAffinity = entity.Affinity{}.DefaultAffinity(env, name)
	}

	newDeployment := entity.Deployment{
		Env:             env,
		Name:            name,
		DeploymentImage: *currentImage,
		Envar:           *currentEnvar,
		Port:            *currentPort,
		Scale:           *currentScale,
		Affinity:        *currentAffinity,
		CreatedBy:       createdBy,
	}

	// Check if redeploy or new deployment
	deployAction := "new" // choice: new|redeploy
	currentDeployment, err := uc.deploy_repo.Get(env, name)
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

func (uc *usecase) Redeploy(env string, serviceName string, createdBy int) error {
	currentDeploy, err := uc.deploy_repo.Get(env, serviceName)
	if err != nil {
		if errors.Is(err, repository.ErrDeploymentNotFound) {
			return fmt.Errorf("%w: %v", ErrUnexpected, "No deployment, cannot redeploy")
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return uc.DeployImage(env, serviceName, currentDeploy.DeploymentImage.Digest, createdBy)
}
