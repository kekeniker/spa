# SPinnaker Admin(SPA) tools

> spa is an swiss knife tool for Spinnaker

## Install

Install by this command.

```console
$ brew install kekeniker/tap/spa
```

## Supported Features

* Create Kubernetes Service Accounts with (Cluster)Role and (Cluster)Binding (and Kubeconfigs)

### Features in detail

Here, I will explain what the commands actually do.

#### `spa service-account get`: Get Service account

It will do the following things:

1. Get service account (and retrieve token)
2. (Optional)Create dedicated Kubernetes config
3. If no `output` option is specified, it will only print the access token

#### `spa service-account create`: Create Service Account

It will do the following things:

1. Create a service account with specified name in specified namespace.
2. If not exist, create a (Cluster)Binding that Spinnaker requires.
3. Bind the service account and the cluster binding.
4. (Optional) Create a dedicated Kubernetes config. This can be used to configure Spinnaker Clouddriver. 

## Author

* [KeisukeYamashita](https://github.com/KeisukeYamashita)

## License

Copyright 2020 KeisukeYamashita.  
SPA is released under the Apache License 2.0.
