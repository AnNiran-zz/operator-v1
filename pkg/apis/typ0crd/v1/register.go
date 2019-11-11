package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// GroupName is the group name used in this package
	GroupName string = "crd.devcluster.com"

	// Kind is the Typ0crd kind
	Kind string = "Typ0crd"

	// GroupVersion is the version
	GroupVersion = "v1"

	// Plural is the plural of Typ0crd
	Plural = "typ0crds"

	// Singular is the singular of Typ0crd
	Singular = "typ0crd"

	// CRDName is the CRD name of Type0crd
	CRDName string = Plural + "." + GroupName
)

var (
	// SchemeBuilder ...
	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	localSchemeBuilder = &SchemeBuilder

	// SchemeGroupVersion - the "group" and the "version" that uniquely identitifes the API
	SchemeGroupVersion = schema.GroupVersion{
		Group:   GroupName,
		Version: GroupVersion,
	}
	// AddToScheme represents the applied functions to the scheme
	AddToScheme = SchemeBuilder.AddToScheme
)

func init() {
	localSchemeBuilder.Register(addKnownTypes)
}

// Resource takes an unqualified resource and returns a Group qualified resource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// addKnownTypes adds the set of types defined in this package to the supplied scheme
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&Typ0crd{},
		&Typ0crdList{},
		&metav1.Status{},
	)

	metav1.AddToGroupVersion(
		scheme,
		SchemeGroupVersion,
	)

	return nil
}
