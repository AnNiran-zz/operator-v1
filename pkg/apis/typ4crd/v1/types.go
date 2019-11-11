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

// Typ4crd is the CRD
type Typ4crd struct {
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
	Status Typ4crdStatus `json:"status,omitempty"`

	// specify custom spec here
	Spec Typ4crdSpec `json:"spec,omitempty"`
}

// Typ4crdSpec is a desired state description of Typ4crd
type Typ4crdSpec struct {
	Message        string `json:"message,omitempty"`
	PingdomCheckID int    `json:"pingdomcheckid,omitempty"`
}

// Typ4crdStatus describes lifecycle status of Typ4crd
type Typ4crdStatus struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Typ4crdList is the list of Typ4crds
type Typ4crdList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Typ4crd `json:"items"`
}
