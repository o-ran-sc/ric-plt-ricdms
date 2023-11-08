package controller

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func int32Ptr(val int) *int32 {
	var a int32
	a = int32(val)
	return &a
}

func GetDeployment() []*appsv1.Deployment {

	deployment1 := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-hw-go",
				"chart":    "hw-go-1.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name: "ricplt-hw-go",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "ricplt-hw-go",
					"release": "release-name",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":             "ricplt-hw-go",
						"kubernetes_name": "ricplt_hw-go",
						"release":         "release-name",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{

						corev1.Container{
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 8080,
									Name:          "http",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									Name:          "rmrroute",
									Protocol:      corev1.Protocol("TCP"),
									ContainerPort: 4561,
								},
								corev1.ContainerPort{
									ContainerPort: 4560,
									Name:          "rmrdata",
									Protocol:      corev1.Protocol("TCP"),
								},
							},
							Stdin: false,
							TTY:   false,
							ReadinessProbe: &corev1.Probe{
								PeriodSeconds: 15,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "ric/v1/health/ready",
										Port: intstr.IntOrString{
											IntVal: 8080,
										},
									},
								},
								InitialDelaySeconds: 5,
							},
							StdinOnce: false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/opt/ric/config",
									Name:      "config-volume",
									ReadOnly:  false,
								},
							},
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-hw-go-appenv",
										},
									},
								},
								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "dbaas-appconfig",
										},
									},
								},
							},
							Image:           "nexus3.o-ran-sc.org:10004/o-ran-sc/ric-app-hw-go:1.1.1",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							LivenessProbe: &corev1.Probe{
								InitialDelaySeconds: 5,
								PeriodSeconds:       15,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "ric/v1/health/alive",
										Port: intstr.IntOrString{
											IntVal: 8080,
										},
									},
								},
							},
							Name: "hw-go",
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					Hostname:    "hw-go",
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "nexus3-o-ran-sc-org-10004",
						},
					},
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "config-volume",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-hw-go-appconfig",
									},
								},
							},
						},
					},
				},
			},
			Paused:   false,
			Replicas: int32Ptr(1),
		},
	}

	return []*appsv1.Deployment{deployment1}
}
