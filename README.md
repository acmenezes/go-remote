# Go Remote!

A complete containerized Golang development environment to work remotely from within a Kubernetes or OpenShift worker node.

---
## What is it?

It seems quite easy to develop Kubernetes applications from our laptops when we're just making API calls somewhere on the internet or between our own containers. But what if our application works with special hardware? What if our application needs to retrieve some information that is sitting on the low level layers on the worker node? What if I need to install new drivers on the node? Or tune the kernel some way?

Those tasks are normally performed by Kubernetes operators that have special access permissions on the worker node to make configurations required by also special applications.

But even if we don't need to configure the worker node we may want, for example, to debug a micro service running on the Pod network to troubleshoot specific internal connection related problems or be able to stop the code on certain lines for inspection when its running inside the cluster. It could be testing a new webhook for a specific admission controller. Of course we can expose those ports on our laptops and let that connection happen on the inverse direction from the kubeapi to the webhook we're developing. But wouldn't be cool if I can run the code from the node without installing anything to it?

That's what the go-remote project does. It's basically an operator that allow users to define their development environment as a custom resource and then deploy and serve this dev environment as a VS Code remote server allowing us, developers, to debug, set break points, log points, inspect variables, run our code without compiling it and use the Kubernetes or OpenShift worker nodes as compute resource for that. With that we can mimic behaviors that would only work on the node and also talk to the worker node system directly if we have the proper credentials for that.

---
## Who is it for?

- Cluster operator developers in need to deal with the worker node directly are the first use case for that.
<br>

- Developers creating webhooks for admission controllers in need to test them inside the cluster.
<br>

- Developers of micro-services that are separated from their counter part in the application by network policies and are not exposed externally.
<br>
 
- Companies that want to host their development environments as a service on Kubernetes and OpenShift clusters with proper security and backup policies. (Future Features)
<br>

---
## How does it work?

Here is a high level overview of the go remote development environment and how it works with VS Code:

<img src='docs/img/go-remote-arch.png'></img>

<br><br>
1. VS Code requires an SSH connection to the remote development environment. Through SSH it can run the VS Code Server and coordinate, through extensions, the IDE actions or tasks remotely. The SSH port is exposed externally via a Kubernetes service. If you are running Kubernetes or an OpenShift cluster in a Cloud Service Provider it will spin up a service with type LoadBalancer and expose SSH on port 2222. On how to use it instructions you will see how to get the LB url in order to connect your local VS Code to the remote environment.
<br>
2. The container itself can expose multiple different ports. It will be exposing SSH through port 2222 in order to avoid requesting privileges to bind low number ports. This will be the backend port for the VS Code / Go Remote service. Those ports are exposed through Supervisord that runs inside the Pod. It may be configured to expose multiple different ports for multiple processes/containers in the same Pod. The goal is to be able to mimic any application behavior on the development environment and still have extra ports like the 2222 for remote debugging.
<br>
3. The VS Code server is the one that does the hard work. It will coordinate the debugging actions on the remote code tree that will be accessed from the Go Remote Pod. If you want to understand better how VS Code remote works check out [here](https://code.visualstudio.com/docs/remote/remote-overview). Please remark that remote container development on VS Code documentation at this point in time is not remote Kubernetes development. It is for running VS Code from a local container running on the developers machine. Not on a cluster worker node.
<br>
4. The source code, the github project we are working on, is pulled  into the Go Remote container on `go/src/github.com/project/`. And that's the path we need to open remotely in order to code, debug and test. It's important to note that both the code on VS Code IDE and the terminals opened from VS Code will be running from the Dev Pod's image and from within the worker node. Any command like `kubectl` or `oc` will reach the kubeapi from within the cluster. Therefore the kubeconfig file also needs to mounted on the Pod. Check the how to configure a dev environment with go-remote operator.
<br>
5. At this point there is a set of specific packages that are being used for the development container image. But any development image may be built and set on the go-remote operator CR. So if you have a different set of requirements or a ready to go image you only need to change the go-remote image field on the CR pointing it to your registry. One cool future feature is adding the build capability to the go-remote operator. That would allow for updating the packages and requirements automatically. Another nice element to note is that it's not restricted to Golang. It can be any language supported by VS Code.
<br>
6. Finally Supervisord is the one that allows for multiple processes running inside this Pod and exposing ports as needed giving flexibility to model the service we want the way we want it. "...it is meant to be used to control processes related to a project or a customer, and is meant to start like any other program at boot time." Check http://supervisord.org/ So let's say you want to serve something on a specific port you can add that together with the 2222 that already serves the VS Code server.

---
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
#### Configuring the Dev Environment

Now we can deploy the operand. A.K.A the dev environment we're going to use. For that we need to configure it first.

Below we can find the development environment for pod-network-operator example:

```
apiVersion: go-remote.opdev.io/v1alpha1
kind: GoRemote
metadata:
  name: goremote-sample
spec:
  # Add fields here
  goRemoteImage: quay.io/acmenezes/go-remote:latest
  gitRepo: https://github.com/opdev/pod-network-operator.git
  containerPorts:
    - containerPort: 2222

  VolumeMounts:
    - mountPath: /tmp/proc
      name: proc
    - mountPath: /var/run/crio/crio.sock
      name: crio-sock
    - name: gitrepo
      mountPath: /root/go/src/github.com/project/pod-network-operator/

  Volumes:
    - name: proc
      hostPath:
        # Mounting the proc file system to get process namespaces
        path: /proc
        type: Directory
    - name: crio-sock
      hostPath:
        # Mounting the crio socket to inspect containers
        path: /var/run/crio/crio.sock
        type: Socket
    - name: gitrepo 
      emptyDir: {}
  
  serviceAccount: "pod-network-operator-sa"
  operatorNamespace: "pod-network-operator"
```

- goRemoteImage:
<br>
    This field is for the operand, meaning the development container image. So let's say you have in mind your perfect development environment with selected libraries, packages and tools and you want that container image you built to be the one deployed here. Here is the field where you tell the operator which image to run for your development environment.
<br>

- gitRepo:
<br>
    Your project certainly stored in some git based repository. Here is where you inform the operator about it. From this URL the operator will `git clone` your entire project on the path `go/src/github.com/project/`.
<br>
- containerPorts:
<br>
    If the application being developed will be listening to specific ports here is where you declare it. Of course, port 2222 must be there because it's how VS Code connects to your Pod/Service.
<br>

-   VolumeMounts and Volumes
<br>
    If your application needs any kind of volume to be mounted here is where you can put the the mount points and configurations. Be aware that the gitrepo is the one that controls where your project will be copied to and available afterwards. The other examples with hostpath on proc and crio socket come from the pod-network-operator. That specific operator needs those for its applications.
<br>

- serviceAccount:
<br>
    Whatever service account is being used on your application it needs to be declared here. That service account will be the same that received the permissions in the first place using the RBAC manifests.
<br>

- goRemoteNamespace:
<br>
    This is your dev namespace. It's where the dev environment will be deployed.


#### Deploying Go Remote, your Dev Environment:

Once we have that CR ready we can deploy it like below:

`oc apply -f config/samples/go-remote_v1alpha1_goremote.yaml`

```
goremote.go-remote.opdev.io/goremote-sample created
```

Now we can go to the namespace that we chose for our application and check for that new environment:

`oc get pods -n pod-network-operator`

```
NAME                         READY   STATUS    RESTARTS   AGE
go-remote-6fcbc758fd-mt9tx   1/1     Running   0          71s
```

This first iteration is deploying in OCP on AWS and therefore we need to grab the LB url to ssh into via VS Code into the remote environment. Here it goes:

`oc get svc`

```
NAME            TYPE           CLUSTER-IP     EXTERNAL-IP                                                                 PORT(S)                         AGE
go-remote-svc   LoadBalancer   172.30.17.68   a1590a18e20754dcfb0c671beb412b09-458278264.ca-central-1.elb.amazonaws.com   2222:31175/TCP,2345:31261/TCP   3m16s
```

Finally we may register that URL on our ssh configuration for VS Code to use it.





Add that new host, the ELB, at the end of the file like below:

```
Host  a1590a18e20754dcfb0c671beb412b09-458278264.ca-central-1.elb.amazonaws.com  
  HostName  a1590a18e20754dcfb0c671beb412b09-458278264.ca-central-1.elb.amazonaws.com   
  Port 2222
  User root
```