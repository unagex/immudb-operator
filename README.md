# Unagex Kubernetes Operator for immudb

![Go version](https://img.shields.io/github/go-mod/go-version/unagex/immudb-operator)
![Kubernetes Version](https://img.shields.io/badge/Kubernetes-1.18%2B-green.svg)
![Release](https://img.shields.io/github/v/release/unagex/immudb-operator)
[![Go Report Card](https://goreportcard.com/badge/github.com/unagex/immudb-operator)](https://goreportcard.com/report/github.com/unagex/immudb-operator)

## Features

- Create immudb clusters defined as custom resources
- Customize storage provisioning for AWS, GCP, AZURE, ...
- Update immudb version and config (soon)
- Export metrics to Prometheus

## Quickstart
1. Deploy the operator with helm
```
helm repo add immudb-operator-charts https://unagex.github.io/immudb-operator
helm repo update
helm install immudb-operator immudb-operator-charts/immudb-operator
```
2. Deploy a basic immudb database
```
kubectl apply -f https://raw.githubusercontent.com/unagex/immudb-operator/main/config/samples/v1_immudb.yaml
```
⬇ See documentation below for more ⬇

## Documentation

* [Operator Installation](./docs/installation.md)
* [Immudb Configuration](./docs/configuration.md)
* [Operator Overview](./docs/overview.md)
<br></br>
* [Contribution](./docs/contribution.md)
* [Contact](mailto:mathieu.cesbron@protonmail.com)