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

func (r *repository) CheckHPAExist(env, name string) error {
	hpaClient := r.k8scli.AutoscalingV1().HorizontalPodAutoscalers(env)
	_, err := hpaClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return repository_intf.ErrK8sHPANotFound
		}
		return err
	}

	return nil
}

func (r *repository) CreateHPA(param entity.Deployment) error {
	hpaClient := r.k8scli.AutoscalingV1().HorizontalPodAutoscalers(param.Env)
	log.Println("Creating HPA...")
	result, err := hpaClient.Create(context.TODO(), NewHPAParam(param), metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Created HPA %q.\n", result.GetObjectMeta().GetName())
	return nil
}

func (r *repository) UpdateHPA(param entity.Deployment) error {
	hpaClient := r.k8scli.AutoscalingV1().HorizontalPodAutoscalers(param.Env)
	currentHpa, err := hpaClient.Get(context.TODO(), param.Name, metav1.GetOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return repository_intf.ErrK8sHPANotFound
		}
		return err
	}
	isUpdate := false
	if *currentHpa.Spec.MinReplicas != int32(param.Scale.MinReplica) {
		isUpdate = true
		currentHpa.Spec.MinReplicas = toInt32Pointer(param.Scale.MinReplica)
	}
	if currentHpa.Spec.MaxReplicas != int32(param.Scale.MaxReplica) {
		isUpdate = true
		currentHpa.Spec.MaxReplicas = int32(param.Scale.MaxReplica)
	}
	if currentHpa.Spec.TargetCPUUtilizationPercentage != nil && *currentHpa.Spec.TargetCPUUtilizationPercentage != int32(param.Scale.TargetCPUUtilization) {
		isUpdate = true
		currentHpa.Spec.TargetCPUUtilizationPercentage = toInt32Pointer(param.Scale.TargetCPUUtilization)
	}

	if !isUpdate {
		fmt.Println("HPA: Nothing to change...")
		return nil
	}

	result, err := hpaClient.Update(context.TODO(), currentHpa, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Updated HPA %q.\n", result.GetObjectMeta().GetName())
	return nil
}
