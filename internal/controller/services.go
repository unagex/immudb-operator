package controller

import (
	"context"
	"fmt"

	immudbiov1 "github.com/MathieuCesbron/immudb-operator/api/v1"
	"github.com/MathieuCesbron/immudb-operator/internal/controller/common"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (r *ImmudbReconciler) ManageServices(ctx context.Context, immudb *immudbiov1.Immudb) error {
	err := r.ManageServiceHTTP(ctx, immudb)
	if err != nil {
		return err
	}

	err = r.ManageServiceGRPC(ctx, immudb)
	return err
}

func (r *ImmudbReconciler) ManageServiceHTTP(ctx context.Context, immudb *immudbiov1.Immudb) error {
	svc := &corev1.Service{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: immudb.Namespace,
		Name:      immudb.Name + "-http",
	}, svc)

	// create if service does not exist
	if k8serrors.IsNotFound(err) {
		svc = r.GetServiceHTTP(immudb)
		r.Log.Info("creating service")
		err = r.Create(ctx, svc)
		if err != nil && !k8serrors.IsAlreadyExists(err) {
			return fmt.Errorf("error creating service: %w", err)
		}
		return nil
	}

	// requeue if we cannot Get the resource
	if err != nil {
		return fmt.Errorf("error getting service: %w", err)
	}

	// TODO: check if service config is good and patch if that's the case.

	return nil
}

func (r *ImmudbReconciler) ManageServiceGRPC(ctx context.Context, immudb *immudbiov1.Immudb) error {
	svc := &corev1.Service{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: immudb.Namespace,
		Name:      immudb.Name + "-grpc",
	}, svc)

	// create if service does not exist
	if k8serrors.IsNotFound(err) {
		svc = r.GetServiceGRPC(immudb)
		r.Log.Info("creating service grpc")
		err = r.Create(ctx, svc)
		if err != nil && !k8serrors.IsAlreadyExists(err) {
			return fmt.Errorf("error creating service grpc: %w", err)
		}
		return nil
	}

	// requeue if we cannot Get the resource
	if err != nil {
		return fmt.Errorf("error getting service grpc: %w", err)
	}

	// TODO: check if service config is good and patch if that's the case.

	return nil
}

func (r *ImmudbReconciler) GetServiceHTTP(immudb *immudbiov1.Immudb) *corev1.Service {
	ls := common.GetLabels(immudb.Name)
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:            immudb.Name + "-http",
			Namespace:       immudb.Namespace,
			OwnerReferences: common.GetOwnerReferences(immudb),
			Labels:          ls,
		},
		Spec: corev1.ServiceSpec{
			Selector: ls,
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					TargetPort: intstr.FromString("http"),
					Port:       8080,
				},
				{
					Name:       "metrics",
					TargetPort: intstr.FromString("metrics"),
					Port:       9497,
				}},
		},
	}
}

func (r *ImmudbReconciler) GetServiceGRPC(immudb *immudbiov1.Immudb) *corev1.Service {
	ls := common.GetLabels(immudb.Name)
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:            immudb.Name + "-grpc",
			Namespace:       immudb.Namespace,
			OwnerReferences: common.GetOwnerReferences(immudb),
			Labels:          ls,
		},
		Spec: corev1.ServiceSpec{
			Selector: ls,
			Ports: []corev1.ServicePort{
				{
					Name:       "grpc",
					TargetPort: intstr.FromString("grpc"),
					Port:       3322,
				},
			},
		},
	}
}
