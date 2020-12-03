package client

import (
	"context"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
)

var _ Client = (*client)(nil)

// Client interface represents the API for the Kubernetes API cluster supported for spin-admin
type Client interface {
	CreateServiceAccount(ctx context.Context, name, namespace string) (*corev1.ServiceAccount, *corev1.Secret, error)
	CreateRole(ctx context.Context, roleName, namespace string) (*rbacv1.Role, error)
	CreateRoleBinding(ctx context.Context, serviceAccountName, roleName, roleBindingName, namespace string) (*rbacv1.RoleBinding, error)
	CreateClusterRole(ctx context.Context, roleName string) (*rbacv1.ClusterRole, error)
	CreateClusterRoleBinding(ctx context.Context, serviceAccountName, roleName, rbName, namespace string) (*rbacv1.ClusterRoleBinding, error)

	CreateKubeConfig(secret *corev1.Secret, username string) (*api.Config, error)
}

// ClientOption is for additional client configurations.
type ClientOption func(*client)

// WithDryRun specifies the dry run operation.
func WithDryRun() ClientOption {
	return func(c *client) {
		c.dryRun = true
	}
}

type client struct {
	clientset *kubernetes.Clientset
	dryRun    bool
}

// NewClient returns a Kubernetes API client that can be used outside the cluster
func NewClient(ctx context.Context, opts ...ClientOption) (Client, error) {
	var kubeconfig string
	c := &client{}
	for _, opt := range opts {
		opt(c)
	}

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeconfig = "h"
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	c.clientset = clientset
	return c, nil
}
