/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	immudbiov1 "github.com/MathieuCesbron/immudb-operator/api/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/go-logr/logr"
)

// ImmudbReconciler reconciles a Immudb object
type ImmudbReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

// +kubebuilder:rbac:groups=immudb.io,resources=immudbs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=immudb.io,resources=immudbs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=immudb.io,resources=immudbs/finalizers,verbs=update
func (r *ImmudbReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log = log.FromContext(ctx).WithName("Reconciler")

	immudb := &immudbiov1.Immudb{}
	err := r.Get(ctx, req.NamespacedName, immudb)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			r.Log.Info("immudb cr has been deleted")
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, fmt.Errorf("error getting immudb cr: %w", err)
	}

	err = r.CreateDatabase(ctx, immudb)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ImmudbReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&immudbiov1.Immudb{}).
		Complete(r)
}
