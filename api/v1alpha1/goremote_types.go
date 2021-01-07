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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GoRemoteSpec defines the desired state of GoRemote
type GoRemoteSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// GoRemoteImage is the base image to run the environment
	GoRemoteImage string `json:"goRemoteImage,omitempty"`
	// GitRepo is the project under development URL in github, gitlab or any git based server
	GitRepo string `json:"gitRepo,omitempty"`
	// ContainerPorts are the ports that should be exposed by the go-remote container
	// for the servive being developed
	ContainerPorts []corev1.ContainerPort `json:"containerPorts,omitempty"`

	// Service ports
	ServicePorts []corev1.ServicePort

	// Extra volumes that should be mounted for the container
	Volumes      []corev1.Volume      `json:"Volumes,omitempty"`
	VolumeMounts []corev1.VolumeMount `json:"VolumeMounts,omitempty"`
}

// GoRemoteStatus defines the observed state of GoRemote
type GoRemoteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GoRemote is the Schema for the goremotes API
type GoRemote struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GoRemoteSpec   `json:"spec,omitempty"`
	Status GoRemoteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GoRemoteList contains a list of GoRemote
type GoRemoteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GoRemote `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GoRemote{}, &GoRemoteList{})
}
