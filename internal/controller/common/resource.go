package common

import (
	immudbiov1 "github.com/MathieuCesbron/immudb-operator/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetLabels(name string) map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       "immudb",
		"app.kubernetes.io/instance":   name,
		"app.kubernetes.io/managed-by": "immudb-operator",
	}
}

func GetOwnerReferences(immudb *immudbiov1.Immudb) []metav1.OwnerReference {
	return []metav1.OwnerReference{
		{
			APIVersion: immudb.APIVersion,
			Kind:       immudb.Kind,
			Name:       immudb.Name,
			UID:        immudb.UID,
		},
	}
}
