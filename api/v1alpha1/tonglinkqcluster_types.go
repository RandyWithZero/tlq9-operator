/*
Copyright 2021.

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

// TongLinkQClusterSpec defines the desired state of TongLinkQCluster
type TongLinkQClusterSpec struct {
	// worker size
	WorkerSize uint8 `json:"workerSize,omitempty"`
	//enable worker backup
	EnableWorkerBackup bool `json:"enableWorkerBackup,omitempty"`
}

// TongLinkQClusterStatus defines the observed state of TongLinkQCluster
type TongLinkQClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:printcolumn:name="Name",type="string",JSONPath=".metadata.name"

// TongLinkQCluster is the Schema for the tonglinkqclusters API
type TongLinkQCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TongLinkQClusterSpec   `json:"spec,omitempty"`
	Status TongLinkQClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TongLinkQClusterList contains a list of TongLinkQCluster
type TongLinkQClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TongLinkQCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TongLinkQCluster{}, &TongLinkQClusterList{})
}
