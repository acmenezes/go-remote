## How to Install the Operator

#### Permissions

When writing an application we need to define service accounts, roles and role bindings for the Pods we are deploying with our code. Therefore, in order to have a similar environment to the actual application we need to have at hand all those manifests. If any SCC must be applied, in the OpenShift case for example, new roles must be built in order to apply those to the dev Pod.

Under the config/rbac_application folder we can find a full example that applies to an operator that requires privileged access to the node. All we need is to replace those with the new application's RBAC files and when running `make deploy` later they should be deployed to the appropriate namespace to receive that application's workloads.

#### Custom Resource Definition

The CRD or Custom Resource Definition for our remote development environment is called goRemote. That will allow us to run `oc get goremote` and list which environments are running on our system and their configuration.

This CRD is installed with the `make deploy` command as well.

#### Finally Deploy the Operator

`git clone https://github.com/opdev/go-remote.git`

`cd go-remote`

`make deploy`

You should see something like below coming up on the screen. Make sure no error messages appear:

```
/usr/local/bin/controller-gen "crd:trivialVersions=true" rbac:roleName=go-remote-manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
cd config/manager && /usr/local/bin/kustomize edit set image controller=quay.io/opdev/go-remote:v0.0.1
/usr/local/bin/kustomize build config/default | kubectl apply -f -
namespace/go-remote-operator created
namespace/pod-network-operator created
customresourcedefinition.apiextensions.k8s.io/goremotes.go-remote.opdev.io created
serviceaccount/go-remote-operator-sa created
serviceaccount/pod-network-operator-sa created
role.rbac.authorization.k8s.io/leader-election-role created
role.rbac.authorization.k8s.io/pod-network-operator-manager-role created
role.rbac.authorization.k8s.io/role-scc-privileged created
clusterrole.rbac.authorization.k8s.io/go-remote-manager-role created
clusterrole.rbac.authorization.k8s.io/pod-network-operator-manager-role created
rolebinding.rbac.authorization.k8s.io/leader-election-rolebinding created
rolebinding.rbac.authorization.k8s.io/manager-rolebinding created
rolebinding.rbac.authorization.k8s.io/rolebinding-priv-scc-pod-network-operator created
clusterrolebinding.rbac.authorization.k8s.io/go-remote-manager-rolebinding created
clusterrolebinding.rbac.authorization.k8s.io/manager-rolebinding created
deployment.apps/go-remote-operator created
```

Finally verify that you have the operator running on the go-remote-operator namespace:

```
oc get pods -n go-remote-operator
NAME                                  READY   STATUS    RESTARTS   AGE
go-remote-operator-856cbdb584-h6bhx   1/1     Running   0          4m1s
```