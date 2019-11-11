package typ0crd

import (
	"errors"
	"fmt"

	"operator-v1/pingdom"
	v1 "operator-v1/pkg/apis/typ0crd/v1"

	log "github.com/Sirupsen/logrus"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectDeleted(obj interface{}) bool
	ObjectUpdated(objOld, objNew interface{})
}

// Typ0Handler ...
type Typ0Handler struct{}

// Init handles any handler initialization
func (t *Typ0Handler) Init() error {
	log.Info("Handler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *Typ0Handler) ObjectCreated(obj interface{}) {}

// ObjectDeleted is called when an object is deleted
func (t *Typ0Handler) ObjectDeleted(obj interface{}) error {
	log.Info("Handler.ObjectDeleted")

	var object *v1.Typ0crd
	var ok bool

	switch obj.(type) {
	case *v1.Typ0crd:
		log.Info("Received type is *v1.Typ0crd")
	case meta_v1.Object:
		log.Info("Received type is meta_v1.Object")
	}

	if object, ok = obj.(*v1.Typ0crd); ok {
		// object is *v1.Typ0crd
		if object.Spec.PingdomCheckID == 0 {
			log.Info("Pingdom check for object missing")
			// why would it be missing
		} else {
			log.Info("Ready to do the delete check now")
			log.Info(object.Spec.PingdomCheckID)
			if err := pingdom.DeleteCheck(string(object.Spec.PingdomCheckID)); err != nil {
				log.Info("An error occurred while deleting from api: %v", err)
				return err
			}
		}

		log.Info(fmt.Sprintf("Uptime check delete: %#v", obj))
		return nil
	}

	return errors.New("Object is not of type typ0crd")
}

// ObjectUpdated is called when an object is updated
func (t *Typ0Handler) ObjectUpdated(objOld, objNew interface{}) error {
	log.Info("Handler.ObjectUpdated")
	return nil
}
