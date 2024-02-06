# Installation

- [Install on Minikube](#install-on-minikube)
- [Install on Amazon Elastic Kubernetes Service (EKS)](#install-on-amazon-elastic-kubernetes-service-eks)
- [Install on Google Kubernetes Engine (GKE)](#install-on-google-kubernetes-engine-gke)
- [Install on Azure Kubernetes Service (AKS)](#install-on-azure-kubernetes-service-aks)
- [Generic installation with kubectl only](#generic-installation-with-kubectl-only)

## Install on minikube
1. Create minikube cluster.
```bash
minikube start
```
2. Add the helm repo.
```bash
helm repo add immudb-operator-charts https://unagex.github.io/immudb-operator
helm repo update
```
3. Install the operator in the namespace `immudb-operator`. See [operator configuration](#operator-configuration) for more customization.
```bash
helm install immudb-operator immudb-operator-charts/immudb-operator -n immudb-operator --create-namespace
```
Congratulations ! The operator is now installed. To test it, you can deploy a basic immudb (optional):

4. Deploy an immudb database in the namespace `default`. See [immudb configuration](./configuration) for more customization.
```bash
kubectl apply -f https://raw.githubusercontent.com/unagex/immudb-operator/main/config/samples/v1_immudb.yaml
```
5. Access immudb web console.
```bash
 minikube service immudb-sample-http --url
```
Click on the first URL returned to access the immudb web console.


## Install on Amazon Elastic Kubernetes Service (EKS)

## Install on Google Kubernetes Engine (GKE)

## Install on Azure Kubernetes Service (AKS)

## Generic installation with kubectl only

# Operator configuration