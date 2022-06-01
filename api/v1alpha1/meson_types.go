/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MesonSpec defines the desired state of Meson
type MesonSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Docker image name.
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=meson/meson
	Image string `json:"appId"`

	// Replicas indicate the replicas to mantain
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=1
	Replicas int32 `json:"replicas"`
}

// MesonStatus defines the observed state of Meson
type MesonStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Pods are the name of the Pods hosting the App
	Pods []string `json:"pods"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Meson is the Schema for the mesons API
type Meson struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MesonSpec   `json:"spec,omitempty"`
	Status MesonStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MesonList contains a list of Meson
type MesonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Meson `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Meson{}, &MesonList{})
}
