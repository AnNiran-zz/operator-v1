package pingdom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	v1 "operator-v1/pkg/apis/typ0crd/v1"

	log "github.com/Sirupsen/logrus"
)

// CreateCheck obtains configuration data from handled object
func CreateCheck(object *v1.Typ0crd) (int, error) {

	objectName := object.GetName()
	objectSelfLink := object.GetSelfLink()

	cmd := exec.Command(os.Getenv("GOPATH")+"/src/operator-v1/scripts/pingdom/create_check.sh", Username, Password, AppKey, objectName, host, objectSelfLink)
	cmd.Stdin = strings.NewReader("")

	var out bytes.Buffer
	cmd.Stdout = &out

	var err bytes.Buffer
	cmd.Stderr = &err

	res := cmd.Run()
	if res != nil {
		log.Info(err.String())
		log.Info(err.String())
		// maybe create object again?
	}

	log.Info(out.String())

	result := make(map[string]map[string]interface{})
	errUn := json.Unmarshal(out.Bytes(), &result)
	if errUn != nil {
		log.Info(errUn.Error())
	}

	var checkid int
	if checkidUint, ok := result["check"]["id"].(float64); ok {
		checkid = int(checkidUint)
	}
	fmt.Println(checkid)
	return checkid, nil
}

// DeleteCheck deletes a pingdom check
func DeleteCheck(checkid string) error {

	cmd := exec.Command(os.Getenv("GOPATH")+"/src/operator-v1/scripts/pingdom/delete_check.sh", Username, Password, AppKey, string(checkid))
	cmd.Stdin = strings.NewReader("")

	var out bytes.Buffer
	cmd.Stdout = &out

	var err bytes.Buffer
	cmd.Stderr = &err

	res := cmd.Run()
	if res != nil {
		fmt.Println(err.String())
		return res
	}

	fmt.Println(out.String())
	return nil
}

// UpdateCheck ...
func UpdateCheck(obj v1.Typ0crd) {

}
