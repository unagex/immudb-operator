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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ImmudbSpec defines the desired state of Immudb
type ImmudbSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Number of desired immudb pods
	// +kubebuilder:validation:Minimum=1
	Replicas int32 `json:"replicas"`
}

// ImmudbStatus defines the observed state of Immudb
type ImmudbStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
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
