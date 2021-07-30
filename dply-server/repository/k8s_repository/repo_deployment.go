package k8s_repository

import (
	"context"
	"fmt"
	"strings"

	repository_intf "github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

type repository struct {
	k8scli *kubernetes.Clientset
}

func New(k8scli *kubernetes.Clientset) repository_intf.K8sRepository {
	return &repository{k8scli}
}

func (r *repository) ApplyService() error {
	return nil
}

func (r *repository) CheckDeploymentExist(env, name string) error {
	deploymentsClient := r.k8scli.AppsV1().Deployments(env)
	_, err := deploymentsClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return repository_intf.ErrK8sDeploymentNotFound
		}
		return err
	}
	return err
}

func (r *repository) CreateDeployment(param entity.Deployment) error {
	deploymentsClient := r.k8scli.AppsV1().Deployments(param.Env)
	fmt.Println("Creating deployment...")
	newDeploymentParam := NewDeploymentParam(param)
	result, err := deploymentsClient.Create(context.TODO(), newDeploymentParam, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

	return nil
}

func (r *repository) UpdateDeployment(param entity.Deployment) error {
	deploymentsClient := r.k8scli.AppsV1().Deployments(param.Env)
	fmt.Println("Update deployment...")
	newDeploymentParam := NewDeploymentParam(param)
	result, err := deploymentsClient.Update(context.TODO(), newDeploymentParam, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Updated deployment %q.\n", result.GetObjectMeta().GetName())
	return nil
}
