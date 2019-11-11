package typ0crd

import (
	"operator-v1/config"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"

	typ0crdClientset "operator-v1/pkg/client/typ0crd/clientset/versioned"
	typ0crdInformers "operator-v1/pkg/client/typ0crd/informers/externalversions"
	typ0crdInformer_v1 "operator-v1/pkg/client/typ0crd/informers/externalversions/typ0crd/v1"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// DelCh receives names of objects to be deleted
var DelCh = make(chan string, 1)

// retrieve the Kubernetes cluster client from outside of the cluster
func getKubernetesClient() (kubernetes.Interface, typ0crdClientset.Interface) {
	// create the config from the path
	config, err := clientcmd.BuildConfigFromFlags("", config.KubeConfigPath)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}

	// generate the client based off of the config
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}

	typ0crdClient, err := typ0crdClientset.NewForConfig(config)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}

	//restCl, _ := rest.RESTClientFor(config)
	log.Info("Successfully constructed k8s client")
	return k8sClient, typ0crdClient
}

// main position
func main() {
	// use a channel to synchronize the finalization for a graceful shutdown
	stopCh := make(chan struct{})
	defer close(stopCh)
	defer close(DelCh)

	// get the Kubernetes client for connectivity
	k8sClient, typ0crdClient := getKubernetesClient()
	typ0crdInformerFactory := typ0crdInformers.NewSharedInformerFactory(typ0crdClient, time.Second*10)

	controller := NewController(k8sClient, typ0crdClient, typ0crdInformerFactory)

	inf := typ0crdInformer_v1.NewTyp0crdInformer(
		typ0crdClient,
		meta_v1.NamespaceAll,
		0,
		cache.Indexers{},
	)
	controller.informer = inf

	go typ0crdInformerFactory.Start(stopCh)

	// run the controller loop to process items
	if err := controller.Run(2, stopCh, DelCh); err != nil {
		log.Error("Error running controller: ", err.Error())
		os.Exit(1)
	}
}
