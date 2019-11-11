package config

import "os"

const (
	// CRDsPath holds path to crds folder
	CRDsPath = "operator-v1/k8s/crd"

	// NamespacesPath holds path to namespaces path
	NamespacesPath = "operator-v1/k8s/namespaces"

	// ClusterHost holds cluster address
	ClusterHost = "" // implement
)

var (
	// KubeConfigPath contains Kube home config path
	KubeConfigPath = os.Getenv("HOME") + "/.kube/config"

	// CRDPath holds path to crd
	CRDPath = os.Getenv("GOPATH") + "/src/operator-v1/k8s/crd"

	// ScriptsPath holds path to bash scripts used across the code
	ScriptsPath = os.Getenv("GOPATH") + "/src/operator-v1/scripts"

	// APISPath holds path to pkg/apis directory tree
	APISPath = os.Getenv("GOPATH") + "/src/operator-v1/pkg/apis"
)
