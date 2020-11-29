package client

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *client) CreateServiceAccount(ctx context.Context, name, namespace string) (*v1.ServiceAccount, *v1.Secret, error) {
	_, err := c.clientset.CoreV1().ServiceAccounts(namespace).Create(ctx, &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{Name: name},
	}, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, nil, err
	}

	watcher, err := c.clientset.CoreV1().ServiceAccounts(namespace).Watch(ctx, metav1.ListOptions{})

	select {
	case <-watcher.ResultChan():
	}

	sa, err := c.clientset.CoreV1().ServiceAccounts(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	secret, err := c.clientset.CoreV1().Secrets(sa.Namespace).Get(ctx, sa.Secrets[0].Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	return sa, secret, nil
}
