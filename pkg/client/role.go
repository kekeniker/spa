package client

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	// See https://blog.spinnaker.io/spinnaker-kubernetes-rbac-c40f1f73c172
	defaultPolicyRules = []rbacv1.PolicyRule{
		{
			APIGroups: []string{""},
			Resources: []string{"namespace", "configMaps", "events", "replicationcontrollers", "serviceaccounts", "pods/logs"},
			Verbs:     []string{"get", "list"},
		},
		{
			APIGroups: []string{""},
			Resources: []string{"pods", "services", "secrets"},
			Verbs:     []string{"*"},
		},
		{
			APIGroups: []string{"autoscaling"},
			Resources: []string{"horizontalpodautoscalers"},
			Verbs:     []string{"list", "get"},
		},
		{
			APIGroups: []string{"apps"},
			Resources: []string{"controllerrevisions", "statefulsets"},
			Verbs:     []string{"list"},
		},
		{
			APIGroups: []string{"extensions"},
			Resources: []string{"deployments", "replicasets", "ingresses"},
			Verbs:     []string{"*"},
		},
	}
)

func (c *client) CreateRole(ctx context.Context, name, namespace string) (*rbacv1.Role, error) {
	role := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Rules: defaultPolicyRules,
	}
	role, err := c.clientset.RbacV1().Roles(namespace).Create(ctx, role, metav1.CreateOptions{})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			return role, nil
		}
		return nil, err
	}

	watcher, err := c.clientset.RbacV1().Roles(namespace).Watch(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	select {
	case <-watcher.ResultChan():
	}

	return c.clientset.RbacV1().Roles(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (c *client) CreateClusterRole(ctx context.Context, name string) (*rbacv1.ClusterRole, error) {
	role := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Rules: defaultPolicyRules,
	}
	role, err := c.clientset.RbacV1().ClusterRoles().Create(ctx, role, metav1.CreateOptions{})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			return role, nil
		}
		return nil, err
	}

	watcher, err := c.clientset.RbacV1().ClusterRoles().Watch(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	select {
	case <-watcher.ResultChan():
	}

	return c.clientset.RbacV1().ClusterRoles().Get(ctx, name, metav1.GetOptions{})
}
