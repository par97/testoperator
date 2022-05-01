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

// TestkindSpec defines the desired state of Testkind
type TestkindSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// type is used to explain how to parse Params map below
	// +kubebuilder:validation:Enum=China;US
	Country string `json:"country"`

	Population int32 `json:"population,omitempty"`

	Cities []string `json:"cities,omitempty"`

	// Params contain provider-specific parameters
	Params map[string]string `json:"params,omitempty"`
}

// TestkindStatus defines the observed state of Testkind
type TestkindStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Phase string `json:"phase,omitempty"`

	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Testkind is the Schema for the testkinds API
type Testkind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestkindSpec   `json:"spec,omitempty"`
	Status TestkindStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TestkindList contains a list of Testkind
type TestkindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Testkind `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Testkind{}, &TestkindList{})
}
