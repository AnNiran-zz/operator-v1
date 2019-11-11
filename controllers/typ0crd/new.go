package typ0crd

import (
	corev1 "k8s.io/api/core/v1"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	appslisters "k8s.io/client-go/listers/apps/v1"

	typ0crdClientset "operator-v1/pkg/client/typ0crd/clientset/versioned"
	typ0crdScheme "operator-v1/pkg/client/typ0crd/clientset/versioned/scheme"
	typ0crdInformers "operator-v1/pkg/client/typ0crd/informers/externalversions"
	typ0crdListers "operator-v1/pkg/client/typ0crd/listers/typ0crd/v1"

	log "github.com/Sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

// Controller struct defines how a controller should encapsulate
// logging, client connectivity, informing (list and watching)
// queueing, and handling of resource changes
type Controller struct {
	logger           *log.Entry
	typ0crdClientset typ0crdClientset.Interface
	k8sClientset     kubernetes.Interface

	queue            workqueue.RateLimitingInterface
	informer         cache.SharedIndexInformer
	typ0crdLister    typ0crdListers.Typ0crdLister
	typ0crdSynced    cache.InformerSynced
	deplymentsLister appslisters.DeploymentLister
	handler          *Typ0Handler
	recorder         record.EventRecorder
	restClient       *rest.RESTClient
}

const controllerAgentName = "typ0crd-controller"

// NewController construct the Controller object which has all of the necessary components to
// handle logging, connections, informing (listing and watching), the queue,
// and the handler
func NewController(
	k8sClient kubernetes.Interface,
	typ0crdClient typ0crdClientset.Interface,
	typ0crdInformerFactory typ0crdInformers.SharedInformerFactory,
) *Controller {

	typ0crdInformer := typ0crdInformerFactory.Crd().V1().Typ0crds()

	typ0crdScheme.AddToScheme(scheme.Scheme)
	log.Info("Creating event broadcaster")

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(log.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: k8sClient.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	// create a new queue so that when the informer gets a resource that is either
	// a result of listing or watching, we can add an idenfitying key to the queue
	// so that it can be handled in the handler
	queue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Pingdom Resources")

	controller := &Controller{
		logger:           log.NewEntry(log.New()),
		k8sClientset:     k8sClient,
		typ0crdClientset: typ0crdClient,
		typ0crdLister:    typ0crdInformer.Lister(),
		typ0crdSynced:    typ0crdInformer.Informer().HasSynced,
		queue:            queue,
		recorder:         recorder,
		handler:          &Typ0Handler{},
	}

	typ0crdInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueCRD,
		UpdateFunc: func(old, new interface{}) {
			err := controller.handler.ObjectUpdated(old, new)
			if err != nil {
				log.Info(err.Error())
			}
		},
		DeleteFunc: func(obj interface{}) {
			err := controller.handler.ObjectDeleted(obj)
			if err != nil {
				log.Info(err.Error())
			}

		},
	})

	return controller
}
