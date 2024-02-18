package controller

import (
	"context"
	"fmt"

	unagexcomv1 "github.com/unagex/immudb-operator/api/v1"
	"github.com/unagex/immudb-operator/internal/controller/common"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	promv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

func (r *ImmudbReconciler) ManageServiceMonitor(ctx context.Context, immudb *unagexcomv1.Immudb) error {
	if !immudb.Spec.ServiceMonitor.Enabled {
		return nil
	}

	sts := &promv1.ServiceMonitor{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: immudb.Namespace,
		Name:      immudb.Name,
	}, sts)

	// create if service monitor does not exist
	if k8serrors.IsNotFound(err) {
		sts := r.GetServiceMonitor(immudb)
		err := r.Create(ctx, sts)
		if err != nil && !k8serrors.IsAlreadyExists(err) {
			return fmt.Errorf("error creating service monitor: %w", err)
		}
		if err == nil {
			r.Log.Info("service monitor created")
		}
		return nil
	}

	if err != nil {
		return fmt.Errorf("error getting service monitor: %w", err)
	}

	return nil
}

func (r *ImmudbReconciler) GetServiceMonitor(immudb *unagexcomv1.Immudb) *promv1.ServiceMonitor {
	ls := common.GetLabels(immudb.Name)
	return &promv1.ServiceMonitor{
		ObjectMeta: metav1.ObjectMeta{
			Name:            immudb.Name,
			Namespace:       immudb.Namespace,
			OwnerReferences: common.GetOwnerReferences(immudb),
			Labels:          immudb.Spec.ServiceMonitor.Labels,
		},
		Spec: promv1.ServiceMonitorSpec{
			Selector: metav1.LabelSelector{
				MatchLabels: ls,
			},
			Endpoints: []promv1.Endpoint{
				{Port: "metrics", Path: "/metrics"},
			},
			NamespaceSelector: promv1.NamespaceSelector{
				MatchNames: []string{immudb.Namespace},
			},
		},
	}
}
