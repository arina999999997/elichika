# Elichika with Kubernetes
The Helm chart offers deployment of Elichika across a Kubernetes cluster utilizing the Elichika docker image.

A Kubernetes cluster is required, along with all prerequisites from Longhorn. More information can be found [here](https://longhorn.io/docs/1.9.0/deploy/install/).

## How to deploy
A public helm chart exists for Elichika, and can be found [here]().

Optionally, the Kubernetes configuration files can be applied inside the `kube/Elichika` directory. Please note that the following files must be applied in a specific order:
```
StorageClasses
PersistentVolumeClaims
Deployments
```
Pods with ReplicaSets should be deployed for services required to run Elichika completely contained within the Kubernetes cluster.

## Updating resources
The helm chart can easily be upgraded with `helm upgrade`; however, a manual deployment of Kubernetes files will require a `rollout restart` for updating the image across the cluster's pods.

## Test environment
Vagrant is utilized with the `libvirt` provider to provision a Kubernetes cluster locally. Additionally, the GitHub workflow produces a sample cluster through Vagrant to ensure resources successfully deploy and have accessible endpoints.

To get started with testing, ensure that all requirements are fulfilled in the [`kube/kubernetes-vagrant/README.md`](../kube/kubernetes-vagrant/README.md).

Then, apply the following commands:
```
vagrant up
```

To cleanup the cluster:
```
vagrant destroy --force
```
A `.kubeconfig` file is accessible at the root of the repository, which should allow the localhost computer to access the cluster normally.

## GitHub Workflow
Upon commits to `main`, a GitHub Workflow is generated to deploy a Kubernetes cluster through the provided `kubernetes-vagrant` submodule.

New images are tested by successfully deploying the container and accessing the WebUI endpoint. The tests are accomplished by adding in Ansible playbooks into the [kube/test](../kube/test) which are used by Vagrant.
