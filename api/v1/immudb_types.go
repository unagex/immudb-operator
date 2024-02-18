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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	knetworkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ImmudbSpec defines the desired state of Immudb
type ImmudbSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// The image name to use for PostgreSQL containers.
	// +kubebuilder:default="codenotary/immudb:latest"
	// +kubebuilder:validation:Optional
	Image string `json:"image,omitempty"`

	// ImagePullPolicy is used to determine when Kubernetes will attempt to
	// pull (download) container images.
	// +kubebuilder:validation:Enum={Always,Never,IfNotPresent}
	// +kubebuilder:default="IfNotPresent"
	// +kubebuilder:validation:Optional
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`

	// Number of desired immudb pods. At the moment, you can just have 1 replica of immudb. We are working to raise that limit.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:default=1
	// +kubebuilder:validation:Optional
	Replicas *int32 `json:"replicas"`

	// +kubebuilder:validation:Required
	Volume ImmudbVolumeSpec `json:"volume"`

	// +kubebuilder:validation:Required
	Ingress ImmudbIngressSpec `json:"ingress"`

	// +kubebuilder:validation:Required
	ServiceMonitor ImmudbServiceMonitorSpec `json:"serviceMonitor"`
}

type ImmudbVolumeSpec struct {
	// StorageClassName defined for the volume.
	// +kubebuilder:validation:Optional
	StorageClassName *string `json:"storageClassName,omitempty"`

	// Size of the volume.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^\d+(Gi|Gb|Ki|)$`
	// +kubebuilder:validation:Pattern=`^\d+(Ki|Mi|Gi|Ti|Pi|Ei|m|k|M|G|T|P|E)$`
	Size string `json:"size"`
}

type ImmudbIngressSpec struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=nginx
	IngressClassName *string `json:"ingressClassName"`

	// +kubebuilder:validation:Optional
	TLS []knetworkingv1.IngressTLS `json:"tls"`

	// +kubebuilder:validation:Optional
	Host string `json:"host"`
}

type ImmudbServiceMonitorSpec struct {
	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled"`

	// Labels Prometheus should be configured to watch.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels"`
}

// ImmudbStatus defines the observed state of Immudb
type ImmudbStatus struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Number of ready replicas.
	ReadyReplicas int32 `json:"readyReplicas"`

	// Instance ready to accept connections.
	Ready bool `json:"ready"`

	// Hosts to connect to the database.
	// +kubebuilder:validation:Optional
	Hosts *HostsStatus `json:"hosts"`
}

type HostsStatus struct {
	// +kubebuilder:validation:Required
	HTTP string `json:"HTTP"`

	// +kubebuilder:validation:Required
	Metrics string `json:"Metrics"`

	// +kubebuilder:validation:Required
	GRPC string `json:"GRPC"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Immudb is the Schema for the immudbs API
type Immudb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImmudbSpec   `json:"spec"`
	Status ImmudbStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ImmudbList contains a list of Immudb
type ImmudbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Immudb `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Immudb{}, &ImmudbList{})
}
