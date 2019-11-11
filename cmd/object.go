package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"operator-v1/config"
	"os/exec"
	"strings"
)

// ./main object add [crd-name] [object-name]
// ./main object remove [crd-name] [object-name]
// ./main object update [crd-name] [object-name]

// createObject creates an object of crd in the cluster
func createObject(args []string) error {
	object := fmt.Sprintf("%s/%s/objs/%s-%s.yaml", config.CRDPath, args[0], args[0], args[1])

	cmd := exec.Command(fmt.Sprintf("%s/crd.sh", config.ScriptsPath), "create", object)
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

// deleteObject deletes an object of crd from the cluster
func deleteObject(args []string) error {
	object := fmt.Sprintf("%s/%s/objs/%s-%s.yaml", config.CRDPath, args[0], args[0], args[1])

	cmd := exec.Command(fmt.Sprintf("%s/crd.sh", config.ScriptsPath), "delete", object)
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

// updateObject updates an object of crd
func updateObject(args []string) error {
	// to be implemented
	return nil
}

// printHelpObject prints out possible command options for an object
func printHelpObject() {
	fmt.Println(`
	Manage objects:
	object [action] [crd-name] [object-name]

	Example:
	Add object to cluster:
	object add typ0crd obj9

	Remove object from cluster:
	object remove typ2crd obj3

	Update object inside the cluster:
	object update typ1crd obj4
	`)

}
