# Go Remote!

A complete containerized golang development environment to work remotely from within an OpenShift worker node.


# Get Started

  1 - Deploying the requirements

    oc apply -f manifests/role_binding.yaml
    oc apply -f manifests/role_scc.yaml
    oc apply -f manifests/role.yaml
    oc apply -f manifests/service_account.yaml
    oc apply -f manifests/service.yaml

  2 - Running the operator

    run vscode debugger

  
