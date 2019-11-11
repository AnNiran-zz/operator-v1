package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	// StatusCreated defines a successfully created check
	StatusCreated = "Created"

	// StatusError defines a error was obtained when creating the check
	StatusError = "Error"

	// StatusProcessed defines a processed check
	StatusProcessed = "Processed"
)

// +genclient
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Typ0crd is the CRD
type Typ0crd struct {
	// TypeMeta is the metadata for the resouce
	// - kind
	// - apiversion
	metav1.TypeMeta `json:",inline"`

	// ObjectMeta contains the metadata fro the particular object
	// - name
	// - namespace
	// - self link
	// - labels

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status Typ0crdStatus `json:"status,omitempty"`

	// specify custom spec here
	Spec Typ0crdSpec `json:"spec,omitempty"`
}

// Typ0crdStatus describes lifecycle status of Typ0crd
type Typ0crdStatus struct {
	Message        string `json:"message"`
	Status         string `json:"status"`
	PingdomCheckID int    `json:"pingdomcheckid"`
	Version        int    `json:"version"`
}

// Typ0crdSpec is a desired state description of Typ0crd
type Typ0crdSpec struct {
	Name           string `json:"name"`
	Replicas       int    `json:"replicas,omitempty"`
	Message        string `json:"message,omitempty"`
	PingdomCheckID int    `json:"pingdomcheckid,omitempty"`
}

// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Typ0crdList is the list of Typ0crds
type Typ0crdList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Typ0crd `json:"items"`
}
