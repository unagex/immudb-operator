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

You should have an EKS cluster already running. See the [official documentation](https://docs.aws.amazon.com/eks/latest/userguide/create-cluster.html) if that's not the case.
1. Install the Amazon EBS CSI driver add-on on your EKS cluster. See the [official documentation](https://docs.aws.amazon.com/eks/latest/userguide/managing-ebs-csi.html) to add it.
2. Grant permissions for your EKS cluster to interact with Amazon EBS volumes, you need to update the IAM roles associated with your EKS nodes. Here is the necessary policy to attach to your cluster role:
```json
{
 "Version": "2012-10-17",
 "Statement": [{
  "Sid": "VisualEditor0",
  "Effect": "Allow",
  "Action": [
   "ec2:CreateVolume",
   "ec2:DeleteVolume",
   "ec2:AttachVolume",
   "ec2:DetachVolume",
   "ec2:DescribeVolumes",
   "ec2:CreateTags",
   "ec2:DeleteTags",
   "ec2:DescribeTags"
  ],
  "Resource": "*"
 }]
}
```
3. Add the helm repo.
```bash
helm repo add immudb-operator-charts https://unagex.github.io/immudb-operator
helm repo update
```
4. Install the operator in the namespace `immudb-operator`. See [operator configuration](#operator-configuration) for more customization.
```bash
helm install immudb-operator immudb-operator-charts/immudb-operator -n immudb-operator --create-namespace
```
Congratulations ! The operator is now installed. To test it, you can deploy a basic immudb (optional):

5. Deploy an immudb database in the namespace `default`. See [immudb configuration](./configuration) for more customization.
```bash
kubectl apply -f https://raw.githubusercontent.com/unagex/immudb-operator/main/config/samples/v1_immudb.yaml
```
6. Access immudb web console on port 8080.
```bash
kubectl port-forward services/immudb-sample-http 8080:8080
```

## Install on Google Kubernetes Engine (GKE)

You should have an GKE cluster already running. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/creating-a-zonal-cluster) if that's not the case.


## Install on Azure Kubernetes Service (AKS)

## Generic installation with kubectl only

# Operator configuration