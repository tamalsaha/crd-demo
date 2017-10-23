package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/appscode/go/log"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/tamalsaha/go-oneliners"
)

func main() {
	masterURL := ""
	kubeconfigPath := "/home/tamal/.kube/config"

	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}

	crdClient := apiextensionsclient.NewForConfigOrDie(config)
	crds, err := crdClient.ApiextensionsV1beta1().CustomResourceDefinitions().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, crd := range crds.Items {
		oneliners.FILE(crd.Name)
	}
	prom, err := crdClient.ApiextensionsV1beta1().CustomResourceDefinitions().Get("prometheuses.monitoring.coreos.com", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	oneliners.FILE(prom)
}
