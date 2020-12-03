package client

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func (c *client) CreateKubeConfig(secret *v1.Secret) (*api.Config, error) {
	newCfg := api.NewConfig()
	oldCfg, err := c.loadConfig()
	if err != nil {
		return nil, err
	}

	fmt.Println(oldCfg.APIVersion)

	newCfg.APIVersion = "v1"
	newCfg.Kind = "Config"
	newCfg.Preferences = *api.NewPreferences()

	currentContextName := oldCfg.CurrentContext
	for name, context := range oldCfg.Contexts {
		if name == currentContextName {
			context.AuthInfo = "username"
			newCfg.CurrentContext = name
			newCfg.Contexts[name] = context
			break
		}
	}

	for name, cluster := range oldCfg.Clusters {
		if name == currentContextName {
			newCfg.Clusters[name] = cluster
		}
	}

	authInfo := &api.AuthInfo{
		Token: string(secret.Data["token"]),
	}

	newCfg.AuthInfos[newCfg.CurrentContext] = authInfo
	return newCfg, nil
}

func (c *client) loadConfig() (*api.Config, error) {
	path, err := findKubeConfig()
	if err != nil {
		return nil, err
	}

	return clientcmd.LoadFromFile(path)
}

func findKubeConfig() (string, error) {
	env := os.Getenv("KUBECONFIG")
	if env != "" {
		return env, nil
	}
	path, err := homedir.Expand("~/.kube/config")
	if err != nil {
		return "", err
	}
	return path, nil
}
