package k8s_repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	repository_intf "github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *repository) CheckServiceExist(env, name string) error {
	svcClient := r.k8scli.CoreV1().Services(env)
	_, err := svcClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return repository_intf.ErrK8sServiceNotFound
		}
		return err
	}

	return nil
}

func (r *repository) CreateService(param entity.Deployment) error {
	svcClient := r.k8scli.CoreV1().Services(param.Env)
	log.Println("Creating Service...")
	result, err := svcClient.Create(context.TODO(), NewServiceParam(param), metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Created Service %q.\n", result.GetObjectMeta().GetName())

	return nil
}

func (r *repository) UpdateService(param entity.Deployment) error {
	svcClient := r.k8scli.CoreV1().Services(param.Env)
	currentSvc, err := svcClient.Get(context.TODO(), param.Name, metav1.GetOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return repository_intf.ErrK8sServiceNotFound
		}
		return err
	}

	fmt.Println("Update Service...")
	result, err := svcClient.Update(context.TODO(), UpdateServiceParam(*currentSvc, param), metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Updated Service %q.\n", result.GetObjectMeta().GetName())
	return nil
}
