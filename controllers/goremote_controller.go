/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	goremotev1alpha1 "github.com/fennec-project/go-remote/api/v1alpha1"
)

// GoRemoteReconciler reconciles a GoRemote object
type GoRemoteReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	GoRemote *goremotev1alpha1.GoRemote
}

// +kubebuilder:rbac:groups=go-remote.fennecproject.io,resources=goremotes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=go-remote.fennecproject.io,resources=goremotes/status,verbs=get;update;patch

// +kubebuilder:rbac:groups="*",resources="*",verbs="*"

// Reconcile runs the logic controlling GoRemote instances
func (r *GoRemoteReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("goremote", req.NamespacedName)

	// get the list of goRemote CRs loop through them
	GoRemote := &goremotev1alpha1.GoRemote{}
	err := r.Get(context.TODO(), req.NamespacedName, GoRemote)
	if err != nil {
		return reconcile.Result{}, err
	}
	r.GoRemote = GoRemote

	// create one deployment per CR

	deploy := r.newDeploymentForGoRemote(GoRemote)
	err = r.Client.Create(context.TODO(), deploy)
	if err != nil {

		return reconcile.Result{}, err
	}

	service := r.newServiceForGoRemote(GoRemote)
	err = r.Client.Create(context.TODO(), service)
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

// SetupWithManager register the controller with K8S manager instance
func (r *GoRemoteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&goremotev1alpha1.GoRemote{}).
		Complete(r)
}

func (r *GoRemoteReconciler) newServiceForGoRemote(goRemote *goremotev1alpha1.GoRemote) runtime.Object {

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "go-remote-svc",
			Namespace: goRemote.Spec.GoRemoteNamespace,
			Labels:    map[string]string{"app": "go-remote"},
		},

		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{Name: "ssh", Port: 2222, TargetPort: intstr.FromInt(2222), Protocol: corev1.ProtocolTCP},
				{Name: "delve", Port: 2345, TargetPort: intstr.FromInt(2345), Protocol: corev1.ProtocolTCP},
			},
			Type:     corev1.ServiceTypeLoadBalancer,
			Selector: map[string]string{"app": "go-remote"},
		},
	}
	// Set GoRemote instance as the owner and controller
	controllerutil.SetControllerReference(goRemote, service, r.Scheme)
	return service
}

func (r *GoRemoteReconciler) newDeploymentForGoRemote(goRemote *goremotev1alpha1.GoRemote) runtime.Object {

	var replicas int32 = 1
	var privileged bool = true
	// var hostPathTypeDir = corev1.HostPathDirectory
	// var hostPathTypeSock = corev1.HostPathSocket
	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "go-remote",
			Namespace: goRemote.Spec.GoRemoteNamespace,
			Labels:    map[string]string{"app": "go-remote"},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "go-remote"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": "go-remote"},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: goRemote.Spec.ServiceAccount,
					NodeSelector:       goRemote.Spec.NodeSelector,
					InitContainers: []corev1.Container{
						{
							Name:    "gitclone",
							Image:   "alpine:3.7",
							Command: []string{"/bin/sh", "-c"},
							Args:    []string{"apk add --no-cache git && git config --global http.sslVerify 'false' && git clone " + goRemote.Spec.GitRepo + " /tmp"},
							VolumeMounts: []corev1.VolumeMount{
								{Name: "gitrepo",
									MountPath: "/tmp",
								},
							},
						},
					},
					Containers: []corev1.Container{
						{
							Name:            "go-remote",
							Image:           goRemote.Spec.GoRemoteImage,
							ImagePullPolicy: corev1.PullAlways,
							SecurityContext: &corev1.SecurityContext{
								Privileged: &privileged,
							},
							Ports: goRemote.Spec.ContainerPorts,

							VolumeMounts: goRemote.Spec.VolumeMounts,
						},
					},
					Volumes: goRemote.Spec.Volumes,
				},
			},
		},
	}
	// Set GoRemote instance as the owner and controller
	controllerutil.SetControllerReference(goRemote, deploy, r.Scheme)
	return deploy
}
