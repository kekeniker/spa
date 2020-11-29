# SPinnaker Admin(SPA) tools

> spa is an swiss knife tool for Spinnaker

## Supported Features

* Create Kubernetes Service Accounts with (Cluster)Role and (Cluster)Binding (and Kubeconfigs)

### Features in detail

Here, I will explain what the commands actually do.

#### `spa service-account get`: Get Service account

It will do the following things:

1. Get service account (and retrieve token)
2. Create dedicated Kubernetes config

#### `spa service-account create`: Create Service Account

It will do the following things:

1. Create a service account with specified name in specified namespace.
2. If not exist, create a (Cluster)Binding that Spinnaker requires.
3. Bind the service account and the cluster binding.
4. (Optional) Create a dedicated Kubernetes config. This can be used to configure Spinnaker Clouddriver. 

## Author

* [KeisukeYamashita](https://github.com/KeisukeYamashita)

## License

Copyright 2020 KeisukeYamashita. marco is released under the Apache License 2.0.
