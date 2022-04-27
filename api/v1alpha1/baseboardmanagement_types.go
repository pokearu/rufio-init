/*
Copyright 2022 Tinkerbell.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PowerState represents power state the BaseboardManagement.
type PowerState string

// BootDevice represents boot device of the BaseboardManagement.
type BootDevice string

const (
	On  PowerState = "on"
	Off PowerState = "off"
)

const (
	Pxe   BootDevice = "pxe"
	Disk  BootDevice = "disk"
	Bios  BootDevice = "bios"
	Cdrom BootDevice = "cdrom"
	Safe  BootDevice = "safe"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BaseboardManagementSpec defines the desired state of BaseboardManagement
type BaseboardManagementSpec struct {

	// Connection represents the BaseboardManagement connectivity information.
	Connection Connection `json:"connection"`

	// Power is the desired power state of the BaseboardManagement.
	// +kubebuilder:validation:Enum=On;Off
	Power PowerState `json:"power"`

	// Vendor is the vendor name of the BaseboardManagement.
	// +kubebuilder:validation:MinLength=1
	Vendor string `json:"vendor"`
}

type Connection struct {
	// Host is the host IP address or hostname of the BaseboardManagement.
	// +kubebuilder:validation:MinLength=1
	Host string `json:"host"`

	// AuthSecretRef is the SecretReference that contains authentication information of the BaseboardManagement.
	// The Secret must contain username and password keys.
	AuthSecretRef corev1.SecretReference `json:"authSecretRef"`

	// InsecureTLS specifies trusted TLS connections.
	InsecureTLS bool `json:"insecureTLS"`
}

// BaseboardManagementStatus defines the observed state of BaseboardManagement
type BaseboardManagementStatus struct {
	// Power is the current power state of the BaseboardManagement.
	// +kubebuilder:validation:Enum=On;Off
	// +optional
	Power PowerState `json:"powerState,omitempty"`

	// BootDevice is the current first BootDevice of the BaseboardManagement.
	// +optional
	// +kubebuilder:validation:Enum=Pxe;Disk;Bios;Cdrom;Safe
	BootDevice BootDevice `json:"bootState,omitempty"`

	// ErrorMessage represents cause of failure to set desired BaseboardManagement state.
	// +optional
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Version is the current BaseboardManagement version.
	// +optional
	Version string `json:"version,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=baseboardmanagements,scope=Namespaced,categories=tinkerbell,singular=baseboardmanagement,shortName=bm

// BaseboardManagement is the Schema for the baseboardmanagements API
type BaseboardManagement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BaseboardManagementSpec   `json:"spec,omitempty"`
	Status BaseboardManagementStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BaseboardManagementList contains a list of BaseboardManagement
type BaseboardManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BaseboardManagement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BaseboardManagement{}, &BaseboardManagementList{})
}
