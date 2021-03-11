## Configuring the Dev Environment

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
<br><br>

- gitRepo:
<br>
    Your project certainly stored in some git based repository. Here is where you inform the operator about it. From this URL the operator will `git clone` your entire project on the path `go/src/github.com/project/`.
<br><br>
- containerPorts:
<br>
    If the application being developed will be listening to specific ports here is where you declare it. Of course, port 2222 must be there because it's how VS Code connects to your Pod/Service.
<br><br>

-   VolumeMounts and Volumes
<br>
    If your application needs any kind of volume to be mounted here is where you can put the the mount points and configurations. Be aware that the gitrepo is the one that controls where your project will be copied to and available afterwards. The other examples with hostpath on proc and crio socket come from the pod-network-operator. That specific operator needs those for its applications.
<br><br>

- serviceAccount:
<br>
    Whatever service account is being used on your application it needs to be declared here. That service account will be the same that received the permissions in the first place using the RBAC manifests.
<br><br>

- goRemoteNamespace:
<br>
    This is your dev namespace. It's where the dev environment will be deployed.