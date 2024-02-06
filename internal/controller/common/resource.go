package common

import (
	unagexcomv1 "github.com/unagex/immudb-operator/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetLabels(name string) map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       "immudb",
		"app.kubernetes.io/instance":   name,
		"app.kubernetes.io/managed-by": "immudb-operator",
	}
}

func GetOwnerReferences(immudb *unagexcomv1.Immudb) []metav1.OwnerReference {
	return []metav1.OwnerReference{
		{
			APIVersion: immudb.APIVersion,
			Kind:       immudb.Kind,
			Name:       immudb.Name,
			UID:        immudb.UID,
		},
	}
}
