package client

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *client) CreateRole(ctx context.Context, name, namespace string) (*rbacv1.Role, error) {
	return c.clientset.RbacV1().Roles(namespace).Create(ctx, &rbacv1.Role{}, metav1.CreateOptions{})
}

func (c *client) CreateClusterRole(ctx context.Context, name string) (*rbacv1.ClusterRole, error) {
	return c.clientset.RbacV1().ClusterRoles().Create(ctx, &rbacv1.ClusterRole{}, metav1.CreateOptions{})
}
