package client

import (
	"context"
	"encoding/base64"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceAccount represents spin-admin tool service account
type ServiceAccount struct {
}

func (c *client) CreateServiceAccount(ctx context.Context, in *ServiceAccount) (string, error) {
	sa, err := c.clientset.CoreV1().ServiceAccounts("default").Create(ctx, &v1.ServiceAccount{}, metav1.CreateOptions{})
	if err != nil {
		return "", nil
	}

	secret, err := c.clientset.CoreV1().Secrets(sa.Namespace).Get(ctx, sa.Name, metav1.GetOptions{})
	if err != nil {
		return "", nil
	}

	b, err := base64.StdEncoding.DecodeString(string(secret.Data["token"]))
	if err != nil {
		return "", nil
	}

	return string(b), nil
}
