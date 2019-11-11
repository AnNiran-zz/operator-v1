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

// Typ1crd is the CRD
type Typ1crd struct {
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
	Status Typ1crdStatus `json:"status,omitempty"`

	// specify custom spec here
	Spec Typ1crdSpec `json:"spec,omitempty"`
}

// Typ1crdSpec is a desired state description of Typ1crd
type Typ1crdSpec struct {
	Message        string `json:"message,omitempty"`
	PingdomCheckID int    `json:"pingdomcheckid,omitempty"`
}

// Typ1crdStatus describes lifecycle status of Typ1crd
type Typ1crdStatus struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Typ1crdList is the list of Typ1crds
type Typ1crdList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Typ1crd `json:"items"`
}
