package client

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *client) CreateServiceAccount(ctx context.Context, name, namespace string) (*v1.ServiceAccount, *v1.Secret, error) {
	sa, err := c.clientset.CoreV1().ServiceAccounts(namespace).Create(ctx, &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{Name: name},
	}, metav1.CreateOptions{})
	if err != nil {
		return nil, nil, err
	}

	secret, err := c.clientset.CoreV1().Secrets(sa.Namespace).Get(ctx, sa.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	return sa, secret, nil
}
