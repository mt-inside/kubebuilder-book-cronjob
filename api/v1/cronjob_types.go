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

package v1

import (
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CronJobSpec defines the desired state of CronJob
type CronJobSpec struct {
	Schedule string `json:"schedule"`

	// +optional
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty"`

	// strings are non-nullable, so for omission detection in all cases we'd need *string, but "" is hardly ever a sensible user-supplied value
	// +optional
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`

	// +optional
	// +kubebuilder:default:=false
	Suspend *bool `json:"suspend,omitempty"`

	JobTemplate batchv1beta1.JobTemplateSpec `json:"jobTemplate"`

	// +optional
	// +kubebuilder:default:=3
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`
	// +optional
	// +kubebuilder:default:=1
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`
}

// +kubebuilder:validation:Enum=Allow;Forbid;Replace
type ConcurrencyPolicy string

const (
	AllowConcurrent   ConcurrencyPolicy = "Allow"
	ForbidConcurrent  ConcurrencyPolicy = "Forbid"
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

// CronJobStatus defines the observed state of CronJob
// GOTCHA! Defining this and putting it in the root struct does NOT give you a Status, you need the +kubebuilder:subresource:status directive too.
//   without it, you get confusing run-time errors from Status().Update about not being able to find the CronJob object itself
type CronJobStatus struct {
	// +optional
	Active []corev1.ObjectReference `json:"active,omitempty"`
	// +optional
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`            // makes it a runtime.Object ?
	metav1.ObjectMeta `json:"metadata,omitempty"` // makes it a metav1.Object ?

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

func init() {
	// there's an apparent race here with main::init(), but
	// recall that init()s are run in the order of which modules import which others
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}
