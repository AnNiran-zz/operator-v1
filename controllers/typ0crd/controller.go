package typ0crd

import (
	"errors"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/klog"

	log "github.com/Sirupsen/logrus"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"

	"operator-v1/pingdom"
	v1 "operator-v1/pkg/apis/typ0crd/v1"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// SuccessSynced is used as part of the Event 'reason' when a Resource is synced
	SuccessSynced = "Synced"
	// ErrResourceExists is used as part of the Event 'reason' when a Resource fails
	// to sync due to an already existing Pingdom Resource
	ErrResourceExists = "ErrResourceExists"

	// MessageResourceExists is the message used for Events when a resource
	// fails to sync due to a Deployment already existing
	MessageResourceExists = "Resource %q already exists and is not managed by Typ0crd"
	// MessageResourceSynced is the message used for an Event fired when a Typ0crd
	// is synced successfully
	MessageResourceSynced = "Resource synced successfully"
)

// Run is the main path of execution for the controller loop
// Sets up the event handlers and types we are interested in
// syncs informer caches and starts workers
// it will block until stopCh is closed - after that it will shut down the workqueue
// and wait for workers to finish processing their current work items
func (c *Controller) Run(thread int, stopCh <-chan struct{}, delch <-chan string) error {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	c.logger.Info("Controller.Run: initiating")

	log.Info("Waiting for informer caches to sync")
	// do the initial synchronization (one time) to populate resources
	if ok := cache.WaitForCacheSync(stopCh, c.typ0crdSynced); !ok {
		log.Error("Error synching cache")
		return errors.New("Waiting for informer caches to sync")
	}

	c.logger.Info("Controller.Run: cache sync complete")
	c.logger.Info("Starting workers")

	for i := 0; i < thread; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	log.Info("Started workers")
	<-stopCh
	log.Info("Shutting down workers")

	return nil
}

// runWorker executes the loop to process new items added to the queue
func (c *Controller) runWorker() {
	log.Info("Controller.runWorker: starting")

	// invoke processNextItem to fetch and consume the next change
	// to a watched or listed resource
	for c.processNextItem() {
		log.Info("Controller.runWorker: processing next item")
	}

	log.Info("Controller.runWorker: completed")
}

// processNextItem retrieves each queued item and takes the
// necessary handler action based off of if the item was
// created or deleted
func (c *Controller) processNextItem() bool {
	log.Info("Controller.processNextItem: start")

	// fetch the next item (blocking) from the queue to process or
	// if a shutdown is requested then return out of this to stop
	// processing
	obj, quit := c.queue.Get()
	// stop the worker loop from running as this indicates we
	// have sent a shutdown message that the queue has indicated
	// from the Get method
	if quit {
		return false
	}

	err := func(obj interface{}) error {
		// c.queue.Forget must be called if we want the item to be re-queued
		// but not in case of transient error
		defer c.queue.Done(obj)

		var key string
		var ok bool
		// we expect strings of the form namespace/name to come out from the workqueue
		// the delayed nature of the workqueue means the items in the informer cache may be
		// more up to date that when the item was initially put onto the workqueue
		if key, ok = obj.(string); !ok {
			// if the item in the queue is invalid in order to stop re-queueing it
			c.queue.Forget(obj)
			log.Error(fmt.Errorf("Expected string in workqueue but got %#v", obj))
			return nil
		}

		// run the synchandler, passing it the namespace/name string of the resource
		// to be synced
		if err := c.syncHandler(key); err != nil {
			log.Error(fmt.Sprintf("error syncing '%s': %s", key, err.Error()))
			return err
		}

		// if no errors occured we forget the item and it will be queued again
		// until another change happens
		c.queue.Forget(obj)
		log.Info(fmt.Sprintf("Successfully synced %s", key))
		return nil
	}(obj)

	if err != nil {
		log.Error(err.Error())
		return true
	}

	// keep the worker loop running by returning true
	return true
}

// syncHandler compares the actual state with the desired state, and attempts to converge the two
// updates the Status Block of the resource with the current status
func (c *Controller) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	obj, err := c.typ0crdLister.Typ0crds(namespace).Get(name)
	if err != nil {
		// resource may no longer exist and we stop processing.
		if k8serrors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("Typ0crd '%s' in work queue no longer exists", key))
			return nil
		}
		return err
	}

	var checkid int
	if obj.Status.PingdomCheckID != 0 {
		log.Info("Pingdom check already created")

		// proceed for update
		log.Info("resource version is:")
		log.Info(obj.ResourceVersion)

	} else {
		checkid, err = pingdom.CreateCheck(obj)
		if err != nil {
			log.Info("Error creating object: ", err.Error())
		}
		log.Info("Obtained checkid is: ", checkid)

		err = c.updateCRDStatus(obj, checkid)
		if err != nil {
			log.Info(fmt.Sprintf("An error occured while updating status: %s", err.Error()))
			return err
		}
	}

	c.recorder.Event(obj, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}

func (c *Controller) updateCRDStatus(typ0crd *v1.Typ0crd, checkid int) error {
	// NEVER modify objects from the store. It's a read-only, local cache.
	// You can use DeepCopy() to make a deep copy of original object and modify this copy
	// Or create a copy manually for better performance

	// add new data to object
	typ0crd.Spec.PingdomCheckID = checkid
	typ0crd.Status.PingdomCheckID = checkid
	typ0crd.Status.Status = SuccessSynced

	res, err := c.typ0crdClientset.CrdV1().Typ0crds(typ0crd.Namespace).Update(typ0crd)
	if err != nil {
		log.Error(err.Error())
	}

	log.Info(res)
	log.Info("Updated object: ", typ0crd.Name, typ0crd.Status.PingdomCheckID)
	return nil
}

// enqueueCRD takes a typ0crd resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than Typ0crd
func (c *Controller) enqueueCRD(obj interface{}) {
	var key string
	var err error
	// convert the resource object into a key (in this case
	// we are just doing it in the format of 'namespace/name')
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	log.Infof("Add typ0crd: %s", key)
	c.queue.Add(key)
}

// handleObject will take any resource implementing metav1.Object and attempt
// to find the Typ0crd resource that 'owns' it. It does this by looking at the
// objects metadata.ownerReferences field for an appropriate OwnerReference.
// It then enqueues that Typ0crd resource to be processed. If the object does not
// have an appropriate OwnerReference, it will simply be skipped
func (c *Controller) handleObject(obj interface{}) {
	var object meta_v1.Object
	var ok bool
	if object, ok = obj.(meta_v1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(meta_v1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		log.Info(fmt.Sprintf("Recovered deleted object '%s' from tombstone", object.GetName()))
	}
	log.Info("Processing object: ", object.GetName())
	if ownerRef := meta_v1.GetControllerOf(object); ownerRef != nil {
		// If this object is not owned by a Typ0crd, we should not do anything more
		// with it.
		if ownerRef.Kind != "Typ0crd" {
			return
		}

		_, err := c.deplymentsLister.Deployments(object.GetNamespace()).Get(ownerRef.Name)
		if err != nil {
			klog.V(4).Infof("ignoring orphaned object '%s' of foo '%s'", object.GetSelfLink(), ownerRef.Name)
			return
		}

		c.enqueueCRD(obj)
		return
	}
}
