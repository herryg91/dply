package k8s_repository

import (
	"fmt"
	"strconv"

	"github.com/herryg91/dply/dply-server/entity"
	v1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	apiv1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func NewDeploymentParam(param entity.Deployment) *v1.Deployment {
	envars := []apiv1.EnvVar{}
	for k, v := range param.Envar.Variables {
		envars = append(envars, apiv1.EnvVar{
			Name:  k,
			Value: fmt.Sprintf("%v", v),
		})
	}
	ports := []apiv1.ContainerPort{}
	for _, p := range param.Port.Ports {
		ports = append(ports, apiv1.ContainerPort{
			Name:          p.Name,
			ContainerPort: int32(p.Port),
			Protocol:      apiv1.Protocol(p.Protocol),
		})
	}

	var liveness_probe, readiness_probe, startup_probe *corev1.Probe = nil, nil, nil
	if param.DeploymentConfig.LivenessProbe != nil && param.DeploymentConfig.LivenessProbe.Path != "" && param.DeploymentConfig.LivenessProbe.Port > 0 {
		liveness_probe = &corev1.Probe{
			Handler: corev1.Handler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: param.DeploymentConfig.LivenessProbe.Path,
					Port: intstr.FromInt(param.DeploymentConfig.LivenessProbe.Port),
				},
			},
			FailureThreshold:    int32(param.DeploymentConfig.LivenessProbe.FailureThreshold),
			InitialDelaySeconds: int32(param.DeploymentConfig.LivenessProbe.InitialDelaySeconds),
			PeriodSeconds:       int32(param.DeploymentConfig.LivenessProbe.PeriodSeconds),
		}
	}
	if param.DeploymentConfig.ReadinessProbe != nil && param.DeploymentConfig.ReadinessProbe.Path != "" && param.DeploymentConfig.ReadinessProbe.Port > 0 {
		readiness_probe = &corev1.Probe{
			Handler: corev1.Handler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: param.DeploymentConfig.ReadinessProbe.Path,
					Port: intstr.FromInt(param.DeploymentConfig.ReadinessProbe.Port),
				},
			},
			FailureThreshold:    int32(param.DeploymentConfig.ReadinessProbe.FailureThreshold),
			InitialDelaySeconds: int32(param.DeploymentConfig.ReadinessProbe.InitialDelaySeconds),
			PeriodSeconds:       int32(param.DeploymentConfig.ReadinessProbe.PeriodSeconds),
		}
	}
	if param.DeploymentConfig.StartupProbe != nil && param.DeploymentConfig.StartupProbe.Path != "" && param.DeploymentConfig.StartupProbe.Port > 0 {
		startup_probe = &corev1.Probe{
			Handler: corev1.Handler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: param.DeploymentConfig.StartupProbe.Path,
					Port: intstr.FromInt(param.DeploymentConfig.StartupProbe.Port),
				},
			},
			FailureThreshold:    int32(param.DeploymentConfig.StartupProbe.FailureThreshold),
			InitialDelaySeconds: int32(param.DeploymentConfig.StartupProbe.InitialDelaySeconds),
			PeriodSeconds:       int32(param.DeploymentConfig.StartupProbe.PeriodSeconds),
		}
	}

	deployment := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      param.Name,
			Namespace: param.Env,
			Labels:    param.GetK8sLabel(),
		},
		Spec: v1.DeploymentSpec{
			Replicas: toInt32Pointer(param.Scale.MinReplica),
			Selector: &metav1.LabelSelector{MatchLabels: map[string]string{
				"app": param.Name,
				"env": param.Env,
			}},
			Strategy: v1.DeploymentStrategy{
				Type: v1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &v1.RollingUpdateDeployment{
					MaxSurge:       &intstr.IntOrString{Type: intstr.Int, IntVal: 25},
					MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 0},
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": param.Name,
						"env": param.Env,
					},
				},
				Spec: apiv1.PodSpec{
					Affinity: &corev1.Affinity{
						NodeAffinity:    &corev1.NodeAffinity{},
						PodAffinity:     &corev1.PodAffinity{},
						PodAntiAffinity: &corev1.PodAntiAffinity{},
					},
					Containers: []apiv1.Container{
						{
							Name:            param.Name,
							Image:           param.DeploymentImage.Image,
							ImagePullPolicy: apiv1.PullIfNotPresent,
							Env:             envars,
							Resources: apiv1.ResourceRequirements{
								Limits: apiv1.ResourceList{
									apiv1.ResourceCPU:    resource.MustParse(strconv.Itoa(param.Scale.MaxCpu) + "m"),
									apiv1.ResourceMemory: resource.MustParse(strconv.Itoa(param.Scale.MaxMemory) + "M"),
								},
								Requests: apiv1.ResourceList{
									apiv1.ResourceCPU:    resource.MustParse(strconv.Itoa(param.Scale.MinCpu) + "m"),
									apiv1.ResourceMemory: resource.MustParse(strconv.Itoa(param.Scale.MinMemory) + "M"),
								},
							},
							Ports:          ports,
							LivenessProbe:  liveness_probe,
							ReadinessProbe: readiness_probe,
							StartupProbe:   startup_probe,
						},
					},
				},
			},
		},
	}

	if len(param.Affinity.NodeAffinity) > 0 {
		requiredNodeAffinity := []corev1.NodeSelectorRequirement{}
		preferredNodeAffinity := []corev1.PreferredSchedulingTerm{}
		for _, nodeAffinity := range param.Affinity.NodeAffinity {
			if nodeAffinity.Mode == entity.AffinityMode_Required {
				requiredNodeAffinity = append(requiredNodeAffinity, corev1.NodeSelectorRequirement{
					Key:      nodeAffinity.Key,
					Operator: corev1.NodeSelectorOperator(nodeAffinity.Operator),
					Values:   nodeAffinity.Values,
				})
			} else {
				preferredNodeAffinity = append(preferredNodeAffinity, corev1.PreferredSchedulingTerm{
					Weight: int32(nodeAffinity.Weight),
					Preference: corev1.NodeSelectorTerm{
						MatchExpressions: []apiv1.NodeSelectorRequirement{
							apiv1.NodeSelectorRequirement{
								Key:      nodeAffinity.Key,
								Operator: corev1.NodeSelectorOperator(nodeAffinity.Operator),
								Values:   nodeAffinity.Values,
							},
						},
					},
				})
			}
		}

		if len(requiredNodeAffinity) > 0 {
			deployment.Spec.Template.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = &corev1.NodeSelector{
				NodeSelectorTerms: []corev1.NodeSelectorTerm{
					corev1.NodeSelectorTerm{
						MatchExpressions: requiredNodeAffinity,
					},
				},
			}
		}
		if len(preferredNodeAffinity) > 0 {
			deployment.Spec.Template.Spec.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution = preferredNodeAffinity
		}
	} else {
		deployment.Spec.Template.Spec.Affinity.NodeAffinity = nil
	}

	if len(param.Affinity.PodAffinity) > 0 {
		requiredPodAffinity := []corev1.PodAffinityTerm{}
		preferredPodAffinity := []corev1.WeightedPodAffinityTerm{}
		for _, podAffinity := range param.Affinity.PodAffinity {
			if podAffinity.Mode == entity.AffinityMode_Required {
				requiredPodAffinity = append(requiredPodAffinity, corev1.PodAffinityTerm{
					TopologyKey: podAffinity.TopologyKey,
					LabelSelector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							metav1.LabelSelectorRequirement{
								Key:      podAffinity.Key,
								Operator: metav1.LabelSelectorOperator(podAffinity.Operator),
								Values:   podAffinity.Values,
							},
						},
					},
				})
			} else {
				preferredPodAffinity = append(preferredPodAffinity, corev1.WeightedPodAffinityTerm{
					Weight: int32(podAffinity.Weight),
					PodAffinityTerm: corev1.PodAffinityTerm{
						TopologyKey: podAffinity.TopologyKey,
						LabelSelector: &metav1.LabelSelector{
							MatchExpressions: []metav1.LabelSelectorRequirement{
								metav1.LabelSelectorRequirement{
									Key:      podAffinity.Key,
									Operator: metav1.LabelSelectorOperator(podAffinity.Operator),
									Values:   podAffinity.Values,
								},
							},
						},
					},
				})
			}
		}

		if len(requiredPodAffinity) > 0 {
			deployment.Spec.Template.Spec.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution = requiredPodAffinity
		}
		if len(preferredPodAffinity) > 0 {
			deployment.Spec.Template.Spec.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution = preferredPodAffinity
		}
	} else {
		deployment.Spec.Template.Spec.Affinity.PodAffinity = nil
	}

	if len(param.Affinity.PodAntiAffinity) > 0 {
		requiredPodAntiAffinity := []corev1.PodAffinityTerm{}
		preferredPodAntiAffinity := []corev1.WeightedPodAffinityTerm{}
		for _, podAntiAffinity := range param.Affinity.PodAntiAffinity {
			if podAntiAffinity.Mode == entity.AffinityMode_Required {
				requiredPodAntiAffinity = append(requiredPodAntiAffinity, corev1.PodAffinityTerm{
					TopologyKey: podAntiAffinity.TopologyKey,
					LabelSelector: &metav1.LabelSelector{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							metav1.LabelSelectorRequirement{
								Key:      podAntiAffinity.Key,
								Operator: metav1.LabelSelectorOperator(podAntiAffinity.Operator),
								Values:   podAntiAffinity.Values,
							},
						},
					},
				})
			} else {
				preferredPodAntiAffinity = append(preferredPodAntiAffinity, corev1.WeightedPodAffinityTerm{
					Weight: int32(podAntiAffinity.Weight),
					PodAffinityTerm: corev1.PodAffinityTerm{
						TopologyKey: podAntiAffinity.TopologyKey,
						LabelSelector: &metav1.LabelSelector{
							MatchExpressions: []metav1.LabelSelectorRequirement{
								metav1.LabelSelectorRequirement{
									Key:      podAntiAffinity.Key,
									Operator: metav1.LabelSelectorOperator(podAntiAffinity.Operator),
									Values:   podAntiAffinity.Values,
								},
							},
						},
					},
				})
			}
		}

		if len(requiredPodAntiAffinity) > 0 {
			deployment.Spec.Template.Spec.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution = requiredPodAntiAffinity
		}
		if len(preferredPodAntiAffinity) > 0 {
			deployment.Spec.Template.Spec.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution = preferredPodAntiAffinity
		}
	} else {
		deployment.Spec.Template.Spec.Affinity.PodAntiAffinity = nil
	}

	return deployment
}

func NewServiceParam(in entity.Deployment) *corev1.Service {
	ports := []corev1.ServicePort{}
	for _, p := range in.Port.Ports {
		ports = append(ports, corev1.ServicePort{
			Name:       p.Name,
			Port:       int32(p.RemotePort),
			TargetPort: intstr.FromInt(p.Port),
			Protocol:   apiv1.Protocol(p.Protocol),
		})
	}
	resp := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      in.Name,
			Namespace: in.Env,
			Labels:    in.GetK8sLabel(),
		},
		Spec: corev1.ServiceSpec{
			Ports: ports,
			Selector: map[string]string{
				"app": in.Name,
				"env": in.Env,
			},
		},
	}

	if in.Port.AccessType == entity.Access_Type_ClusterIP {
		resp.Spec.Type = apiv1.ServiceTypeClusterIP
		if in.Port.ExternalIP != "" {
			resp.Spec.ExternalIPs = []string{in.Port.ExternalIP}
		}
	} else if in.Port.AccessType == entity.Access_Type_LoadBalancer {
		resp.Spec.Type = apiv1.ServiceTypeLoadBalancer
	} else {
		resp.Spec.Type = apiv1.ServiceTypeClusterIP
	}

	return resp
}

func UpdateServiceParam(old corev1.Service, in entity.Deployment) *corev1.Service {
	ports := []corev1.ServicePort{}
	for _, p := range in.Port.Ports {
		ports = append(ports, corev1.ServicePort{
			Name:       p.Name,
			Port:       int32(p.RemotePort),
			TargetPort: intstr.FromInt(p.Port),
			Protocol:   apiv1.Protocol(p.Protocol),
		})
	}
	old.Spec.Ports = ports
	if in.Port.AccessType == entity.Access_Type_ClusterIP {
		old.Spec.Type = apiv1.ServiceTypeClusterIP
	} else if in.Port.AccessType == entity.Access_Type_LoadBalancer {
		old.Spec.Type = apiv1.ServiceTypeLoadBalancer
		if in.Port.ExternalIP != "" {
			old.Spec.LoadBalancerIP = in.Port.ExternalIP
		}
	} else {
		old.Spec.Type = apiv1.ServiceTypeClusterIP
	}

	old.Spec.Selector = map[string]string{
		"app": in.Name,
		"env": in.Env,
	}
	old.ObjectMeta.Labels = in.GetK8sLabel()
	return &old
}

func NewHPAParam(param entity.Deployment) *autoscalingv1.HorizontalPodAutoscaler {
	return &autoscalingv1.HorizontalPodAutoscaler{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "autoscaling/v1",
			Kind:       "HorizontalPodAutoscaler",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      param.Name,
			Namespace: param.Env,
			Labels:    param.GetK8sLabel(),
		},
		Spec: autoscalingv1.HorizontalPodAutoscalerSpec{
			MinReplicas: toInt32Pointer(param.Scale.MinReplica),
			MaxReplicas: int32(param.Scale.MaxReplica),
			ScaleTargetRef: autoscalingv1.CrossVersionObjectReference{
				Kind:       "Deployment",
				APIVersion: "apps/v1",
				Name:       param.Name,
			},
			TargetCPUUtilizationPercentage: toInt32Pointer(param.Scale.TargetCPUUtilization),
		},
	}
}
