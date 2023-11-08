package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func GetService() []*corev1.Service {

	service1 := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-hw-go",
				"chart":    "hw-go-1.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-hw-go-http",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "http",
					Port:     8080,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "http",
						Type:   intstr.Type(1),
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-hw-go",
				"release": "release-name",
			},
		},
	}

	service2 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"heritage": "Helm",
				"release":  "release-name",
				"app":      "ricplt-hw-go",
				"chart":    "hw-go-1.0.0",
			},
			Name:      "service-ricplt-hw-go-rmr",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Port:     4560,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "rmrdata",
						Type:   intstr.Type(1),
					},
					Name: "rmrdata",
				},
				corev1.ServicePort{
					Name:     "rmrroute",
					Port:     4561,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Type(1),
						StrVal: "rmrroute",
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-hw-go",
				"release": "release-name",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	return []*corev1.Service{service1, service2}
}
