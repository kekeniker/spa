package client

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd/api"
)

func (c *client) CreateKubeConfig(secret *v1.Secret) (*api.Config, error) {
	newCfg := api.NewConfig()
	oldCfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	_ = newCfg
	_ = oldCfg

	return newCfg, nil
}

func loadConfig() (*api.Config, error) {
	b, err := ioutil.ReadFile("~/.kube/config")
	if err != nil {
		return nil, err
	}

	cfg := api.NewConfig()
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
