package deploy_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
)

type usecase struct {
	deploy_repo   repository.DeploymentRepository
	k8s_repo      repository.K8sRepository
	image_repo    repository.ImageRepository
	envar_repo    repository.EnvarRepository
	scale_repo    repository.ScaleRepository
	port_repo     repository.PortRepository
	affinity_repo repository.AffinityRepository
}

func New(
	deploy_repo repository.DeploymentRepository,
	k8s_repo repository.K8sRepository,
	image_repo repository.ImageRepository,
	envar_repo repository.EnvarRepository,
	scale_repo repository.ScaleRepository,
	port_repo repository.PortRepository,
	affinity_repo repository.AffinityRepository,
) UseCase {
	return &usecase{
		deploy_repo:   deploy_repo,
		k8s_repo:      k8s_repo,
		image_repo:    image_repo,
		envar_repo:    envar_repo,
		scale_repo:    scale_repo,
		port_repo:     port_repo,
		affinity_repo: affinity_repo,
	}
}

func (uc *usecase) DeployImage(env string, name string, digest string, createdBy int) error {
	currentImage, err := uc.image_repo.GetByDigest(digest)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	currentEnvar, _ := uc.envar_repo.Get(env, name)
	if currentEnvar == nil {
		currentEnvar = entity.Envar{}.DefaultEnvar(env, name)
	}

	currentScale, _ := uc.scale_repo.Get(env, name)
	if currentScale == nil {
		currentScale = entity.Scale{}.DefaultScale(env, name)
	}

	currentPort, _ := uc.port_repo.Get(env, name)
	if currentPort == nil {
		currentPort = entity.Port{}.DefaultPort(env, name)
	}

	currentAffinity, _ := uc.affinity_repo.Get(env, name)
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
