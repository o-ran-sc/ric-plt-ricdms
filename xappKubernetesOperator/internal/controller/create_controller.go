package controller

import (
	"context"
	"fmt"
)

func (r *XappDepReconciler) CreateAll() {
	var err error
	namespaceProvided := "ricxapp"

	for _, resource := range GetConfigMap() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetConfigMap()| Error --> |", err)
		}
	}

	for _, resource := range GetDeployment() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetDeployment()| Error --> |", err)
		}
	}

	for _, resource := range GetService() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetService()| Error --> |", err)
		}
	}

}
