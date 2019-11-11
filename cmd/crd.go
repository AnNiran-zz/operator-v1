package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"operator-v1/config"
	"os"
	"os/exec"
	"strings"
)

// ./main crd add [crd-name]
// ./main crd remove [crd-name]
// ./main crd gen-client [crd-name]
// ./main crd del-client [crd-name]

// addCRD adds a CRD to the cluster using kubectl
func addCRD(crd string) error {
	// kubectl create -f /hdd/program/go/src/operator-v1/k8s/crd
	crdLoc := fmt.Sprintf("%s/%s/%s.yaml", config.CRDPath, crd, crd)

	cmd := exec.Command(fmt.Sprintf("%s/crd.sh", config.ScriptsPath), "create", crdLoc)
	cmd.Stdin = strings.NewReader("")

	var out bytes.Buffer
	cmd.Stdout = &out

	var err bytes.Buffer
	cmd.Stderr = &err

	res := cmd.Run()
	if res != nil {
		fmt.Println(err.String())
		return errors.New(err.String())
	}

	fmt.Println(out.String())
	return nil
}

// removeCRD removes a CRD from the cluster using kubectl
func removeCRD(crd string) error {
	// kubectl delete -f /hdd/program/go/src/operator-v1/k8s/crd
	crdLoc := fmt.Sprintf("%s/%s/%s.yaml", config.CRDPath, crd, crd)

	cmd := exec.Command(fmt.Sprintf("%s/crd.sh", config.ScriptsPath), "delete", crdLoc)
	cmd.Stdin = strings.NewReader("")

	var out bytes.Buffer
	cmd.Stdout = &out

	var err bytes.Buffer
	cmd.Stderr = &err

	res := cmd.Run()
	if res != nil {
		fmt.Println(err.String())
		return errors.New(err.String())
	}

	fmt.Println(out.String())
	return nil
}

// generateCRDClientData generates CRD client data files
func generateCRDClientData(crd string) error {
	// check if crd folder exists in /apis
	apisfolder := fmt.Sprintf("%s/%s", config.APISPath, crd)

	if _, err := os.Stat(apisfolder); os.IsNotExist(err) {
		return errors.New("CRD path inside pkg/apis does not exist")
	}

	cmd := exec.Command(fmt.Sprintf("%s/generate.sh", config.ScriptsPath), crd)
	cmd.Stdin = strings.NewReader("")

	var out bytes.Buffer
	cmd.Stdout = &out

	var err bytes.Buffer
	cmd.Stderr = &err

	res := cmd.Run()
	if res != nil {
		fmt.Println(err.String())
		return errors.New(err.String())
	}

	fmt.Println(out.String())
	return nil
}

// deleteCRDClientData deletes CRD client data files
func deleteCRDClientData(crd string) error {
	cmd := exec.Command(fmt.Sprintf("%s/delete.sh", config.ScriptsPath), crd)
	cmd.Stdin = strings.NewReader("")

	var out bytes.Buffer
	cmd.Stdout = &out

	var err bytes.Buffer
	cmd.Stderr = &err

	res := cmd.Run()
	if res != nil {
		fmt.Println(err.String())
		return errors.New(err.String())
	}

	fmt.Println(out.String())
	return nil
}

// printHelpCRD prints out possible command options for a CRD
func printHelpCRD() {
	fmt.Println(`
	Manage Custom Resource Definitions:
	crd [action] [crd-name]

	Example:
	Add CRD to cluster:
	crd add typ0crd

	Remove CRD from cluster:
	crd remove typ0crd

	Generate CRD pkg/client files:
	crd generate-client typ3crd

	Delete CRD pkg/client files:
	crd delete-client typ2crd
	`)
}
