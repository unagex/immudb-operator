package controller

import (
	"context"
	"fmt"

	unagexcomv1 "github.com/unagex/immudb-operator/api/v1"
	"github.com/unagex/immudb-operator/internal/controller/common"
	knetworkingv1 "k8s.io/api/networking/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
)

func (r *ImmudbReconciler) ManageIngress(ctx context.Context, immudb *unagexcomv1.Immudb) error {
	if !immudb.Spec.Ingress.Enabled {
		return nil
	}

	sts := &knetworkingv1.Ingress{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: immudb.Namespace,
		Name:      immudb.Name,
	}, sts)

	// create if ingress does not exist
	if k8serrors.IsNotFound(err) {
		sts := r.GetIngress(immudb)
		err := r.Create(ctx, sts)
		if err != nil && !k8serrors.IsAlreadyExists(err) {
			return fmt.Errorf("error creating ingress: %w", err)
		}
		if err == nil {
			r.Log.Info("ingress created")
		}
		return nil
	}

	if err != nil {
		return fmt.Errorf("error getting ingress: %w", err)
	}

	return nil
}

func (r *ImmudbReconciler) GetIngress(immudb *unagexcomv1.Immudb) *knetworkingv1.Ingress {
	ls := common.GetLabels(immudb.Name)
	return &knetworkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:            immudb.Name,
			Namespace:       immudb.Namespace,
			OwnerReferences: common.GetOwnerReferences(immudb),
			Labels:          ls,
		},
		Spec: knetworkingv1.IngressSpec{
			IngressClassName: immudb.Spec.Ingress.IngressClassName,
			TLS:              immudb.Spec.Ingress.TLS,
			Rules: []knetworkingv1.IngressRule{
				{
					Host: immudb.Spec.Ingress.Host,
					IngressRuleValue: knetworkingv1.IngressRuleValue{
						HTTP: &knetworkingv1.HTTPIngressRuleValue{
							Paths: []knetworkingv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: ptr.To(knetworkingv1.PathTypePrefix),
									Backend: knetworkingv1.IngressBackend{
										Service: &knetworkingv1.IngressServiceBackend{
											Name: immudb.Name + "-http",
											Port: knetworkingv1.ServiceBackendPort{
												Name: "http",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
