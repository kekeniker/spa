package client

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *client) CreateRoleBinding(ctx context.Context, saName, rName, rbName, namespace string) (*rbacv1.RoleBinding, error) {
	b := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      rbName,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     rName,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Namespace: namespace,
				Name:      saName,
			},
		},
	}

	rb, err := c.clientset.RbacV1().RoleBindings(namespace).Create(ctx, b, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	watcher, err := c.clientset.RbacV1().RoleBindings(namespace).Watch(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	select {
	case <-watcher.ResultChan():
	}

	return rb, nil
}

func (c *client) CreateClusterRoleBinding(ctx context.Context, saName, rName, rbName, namespace string) (*rbacv1.ClusterRoleBinding, error) {
	b := &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: rbName,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     rName,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Namespace: namespace,
				Name:      saName,
			},
		},
	}

	rb, err := c.clientset.RbacV1().ClusterRoleBindings().Create(ctx, b, metav1.CreateOptions{})
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			return nil, err
		}
	}

	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	watcher, err := c.clientset.RbacV1().ClusterRoleBindings().Watch(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	select {
	case <-watcher.ResultChan():
	}

	return rb, nil
}
