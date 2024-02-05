# Unagex Kubernetes Operator for immudb

![Go version](https://img.shields.io/github/go-mod/go-version/unagex/immudb-operator)
![Kubernetes Version](https://img.shields.io/badge/Kubernetes-1.18%2B-green.svg)

## Features

- Create immudb clusters defined as custom resources
- Customize storage provisioning for AWS, GCP, AZURE, ...
- Upgrade immudb version and config
- Export metrics to Prometheus

## Documentation

TODO: create documentation

```
helm repo add immudb-operator-charts https://unagex.github.io/immudb-operator
helm repo update
helm install immudb-operator immudb-operator-charts/immudb-operator
```