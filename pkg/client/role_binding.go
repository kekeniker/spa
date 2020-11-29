package client

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *client) CreateRoleBinding(ctx context.Context, saName, rName, namespace string) (*rbacv1.RoleBinding, error) {
	return c.clientset.RbacV1().RoleBindings(namespace).Create(ctx, &rbacv1.RoleBinding{}, metav1.CreateOptions{})
}

func (c *client) CreateClusterRoleBinding(ctx context.Context, saName, rName string) (*rbacv1.ClusterRoleBinding, error) {
	return c.clientset.RbacV1().ClusterRoleBindings().Create(ctx, &rbacv1.ClusterRoleBinding{}, metav1.CreateOptions{})
}
