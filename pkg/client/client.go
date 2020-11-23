package client

import (
	"context"
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Client interface represents the API for the Kubernetes API cluster supported for spin-admin
type Client interface {
	CreateServiceAccount(context.Context, *ServiceAccount) (string, error)
}

type client struct {
	clientset *kubernetes.Clientset
}

// NewClient returns a Kubernetes API client that can be used outside the cluster
func NewClient(ctx context.Context) (Client, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &client{
		clientset: clientset,
	}, nil
}
