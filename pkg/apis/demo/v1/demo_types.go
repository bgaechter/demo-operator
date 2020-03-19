package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DemoSpec defines the desired state of Demo
type DemoSpec struct {
	Message string `json:"message"`
	Count   int32  `json:"count"`
}

// DemoStatus defines the observed state of Demo
type DemoStatus struct {
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Demo is the Schema for the demos API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=demos,scope=Namespaced
type Demo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoSpec   `json:"spec,omitempty"`
	Status DemoStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DemoList contains a list of Demo
type DemoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Demo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Demo{}, &DemoList{})
}
