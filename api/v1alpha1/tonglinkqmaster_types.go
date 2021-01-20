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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TongLinkQMasterSpec defines the desired state of TongLinkQMaster
type TongLinkQMasterSpec struct {
	// master image
	Image string `json:"image,omitempty"`
	// master image pull policy
	ImagePullPolicy v1.PullPolicy    `json:"imagePullPolicy,omitempty"`
	Envs            []v1.EnvVar      `json:"env,omitempty"`
	Volumes         []v1.Volume      `json:"volumes,omitempty"`
	VolumeMounts    []v1.VolumeMount `json:"volumeMounts,omitempty"`
	Port            int32            `json:"port,omitempty"`
}

type MasterStatus string

var (
	Healthy   MasterStatus = "Healthy"
	UnHealthy MasterStatus = "UnHealthy"
	Pending   MasterStatus = "Pending"
	Fail      MasterStatus = "Fail"
)

// TongLinkQClusterStatus defines the observed state of TongLinkQCluster
type TongLinkQMasterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Parse MasterStatus `json:"parse,omitempty"`
	// master pod ip
	MasterIp string `json:"masterIp,omitempty"`
	// master pod node host ip
	NodeIp string `json:"nodeIp,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:printcolumn:name="Name",type="string",JSONPath=".metadata.name"
//+kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.parse"
//+kubebuilder:printcolumn:name="MasterIp",type="string",JSONPath=".status.masterIp",priority=10
//+kubebuilder:printcolumn:name="NodeIp",type="string",JSONPath=".status.nodeIp",priority=10

// TongLinkQCluster is the Schema for the tonglinkqclusters API
type TongLinkQMaster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TongLinkQMasterSpec   `json:"spec,omitempty"`
	Status TongLinkQMasterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TongLinkQClusterList contains a list of TongLinkQCluster
type TongLinkQMasterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TongLinkQMaster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TongLinkQMaster{}, &TongLinkQMasterList{})
}
