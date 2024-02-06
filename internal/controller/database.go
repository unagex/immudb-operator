package controller

import (
	"context"
	"fmt"

	unagexcomv1 "github.com/unagex/immudb-operator/api/v1"
	"github.com/unagex/immudb-operator/internal/controller/common"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/apimachinery/pkg/types"
)

func (r *ImmudbReconciler) ManageDatabase(ctx context.Context, immudb *unagexcomv1.Immudb) error {
	sts := &appsv1.StatefulSet{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: immudb.Namespace,
		Name:      immudb.Name,
	}, sts)

	// create if statefulset does not exist
	if k8serrors.IsNotFound(err) {
		sts := r.GetStatefulset(immudb)
		r.Log.Info("creating statefulset")
		err := r.Create(ctx, sts)
		if err != nil && !k8serrors.IsAlreadyExists(err) {
			return fmt.Errorf("error creating statefulset: %w", err)
		}
		return nil
	}

	if err != nil {
		return fmt.Errorf("error getting statefulset: %w", err)
	}

	// update if statefulset config is wrong
	if *sts.Spec.Replicas != *immudb.Spec.Replicas {
		sts = r.GetStatefulset(immudb)
		r.Log.Info(fmt.Sprintf("updating statefulset field spec.replicas to %d", *immudb.Spec.Replicas))
		err = r.Update(ctx, sts)
		if err != nil {
			return fmt.Errorf("error updating statefulset field spec.replicas: %w", err)
		}
		return nil
	}

	// update status if not sync with statefulset
	diff := immudb.Status.ReadyReplicas != sts.Status.ReadyReplicas
	if diff {
		immudb.Status.ReadyReplicas = sts.Status.ReadyReplicas
		immudb.Status.Ready = false
		// TODO: instance is ready if quorum and not if different.
		if immudb.Status.ReadyReplicas == *immudb.Spec.Replicas {
			immudb.Status.Ready = true
		}
		r.Log.Info(fmt.Sprintf("update immudb field status.readyReplicas to %d/%d",
			immudb.Status.ReadyReplicas, *immudb.Spec.Replicas))
		err = r.Status().Update(ctx, immudb)
		if err != nil {
			return fmt.Errorf("error updating immudb field status.readyReplicas: %w", err)
		}
	}

	return nil
}

func (r *ImmudbReconciler) GetStatefulset(immudb *unagexcomv1.Immudb) *appsv1.StatefulSet {
	ls := common.GetLabels(immudb.Name)
	return &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:            immudb.Name,
			Namespace:       immudb.Namespace,
			OwnerReferences: common.GetOwnerReferences(immudb),
			Labels:          ls,
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: immudb.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			ServiceName: immudb.Name + "-http",
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					// TODO: Add support for AWS, GCP and all
					Volumes: []corev1.Volume{
						{
							Name: immudb.Name + "-storage",
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: immudb.Name,
								},
							},
						},
					},
					Containers: []corev1.Container{
						{
							Image:           immudb.Spec.Image,
							ImagePullPolicy: immudb.Spec.ImagePullPolicy,
							Name:            "immudb",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: 8080,
								},
								{
									Name:          "grpc",
									ContainerPort: 3322,
								},
								{
									Name:          "metrics",
									ContainerPort: 9497,
								},
							},
							ReadinessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path:   "/readyz",
										Port:   intstr.FromString("metrics"),
										Scheme: corev1.URISchemeHTTP,
									},
								},
							},
							LivenessProbe: &corev1.Probe{
								FailureThreshold: 9,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path:   "/readyz",
										Port:   intstr.FromString("metrics"),
										Scheme: corev1.URISchemeHTTP,
									},
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									MountPath: "/var/lib/immudb",
									Name:      immudb.Name + "-storage",
									// TODO: Add a variable to disable SubPath if we want. Check the helm chart for more informations.
									SubPath: "immudb",
								},
							},
						},
					},
				},
			},
			VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: immudb.Name + "-storage",
					},
					Spec: corev1.PersistentVolumeClaimSpec{
						// TODO: Add storageClassName option
						AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								// TODO: Resize would be possible depending on the cloud provider.
								corev1.ResourceStorage: resource.MustParse("1Gi"),
							},
						},
					},
				},
			},
		},
	}
}
