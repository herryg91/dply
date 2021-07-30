package deploy_usecase

import (
	"errors"
	"fmt"

	"github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
)

func (uc *usecase) deployK8s(deployment entity.Deployment) error {
	// Create / Update Deployment
	deploymentAction := "update"
	err := uc.k8s_repo.CheckDeploymentExist(deployment.Env, deployment.Name)
	if err != nil {
		if errors.Is(err, repository.ErrK8sDeploymentNotFound) {
			deploymentAction = "create"
		} else {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	}

	if deploymentAction == "create" {
		err = uc.k8s_repo.CreateDeployment(deployment)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	} else {
		err = uc.k8s_repo.UpdateDeployment(deployment)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	}

	// Create / Update Service
	serviceAction := "update"
	err = uc.k8s_repo.CheckServiceExist(deployment.Env, deployment.Name)
	if err != nil {
		if errors.Is(err, repository.ErrK8sServiceNotFound) {
			serviceAction = "create"
		} else {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	}
	if serviceAction == "create" {
		err = uc.k8s_repo.CreateService(deployment)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	} else {
		err = uc.k8s_repo.UpdateService(deployment)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	}
	// Create / Update HPA
	hpaAction := "update"
	err = uc.k8s_repo.CheckHPAExist(deployment.Env, deployment.Name)
	if err != nil {
		if errors.Is(err, repository.ErrK8sHPANotFound) {
			hpaAction = "create"
		} else {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	}
	if hpaAction == "create" {
		err = uc.k8s_repo.CreateHPA(deployment)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	} else {
		err = uc.k8s_repo.UpdateHPA(deployment)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
		}
	}

	return nil
}
