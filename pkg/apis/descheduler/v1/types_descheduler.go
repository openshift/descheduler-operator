package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeDescheduler is the Schema for the deschedulers API
// +k8s:openapi-gen=true
// +genclient
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
type KubeDescheduler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user settable values for configuration
	// +required
	Spec KubeDeschedulerSpec `json:"spec"`
	// status holds observed values from the cluster. They may not be overridden.
	// +optional
	Status KubeDeschedulerStatus `json:"status"`
}

// KubeDeschedulerSpec defines the desired state of KubeDescheduler
type KubeDeschedulerSpec struct {
	operatorv1.OperatorSpec `json:",inline"`

	// Profiles sets which descheduler strategy profiles are enabled
	Profiles []DeschedulerProfile `json:"profiles"`

	// DeschedulingIntervalSeconds is the number of seconds between descheduler runs
	// +optional
	DeschedulingIntervalSeconds *int32 `json:"deschedulingIntervalSeconds,omitempty"`

	// PodLifetime is the length of time after which pods should be evicted
	// This field should be used with profiles that enable the PodLifetime strategy, such as LifecycleAndUtilization
	PodLifetime *time.Duration `json:"podLifetime,omitempty"`
}

// DeschedulerProfile allows configuring the enabled strategy profiles for the descheduler
// it allows multiple profiles to be enabled at once, which will have cumulative effects on the cluster.
// +kubebuilder:validation:Enum=AffinityAndTaints;TopologyAndDuplicates;LifecycleAndUtilization;DevPreviewLongLifecycle
type DeschedulerProfile string

var (
	// AffinityAndTaints enables descheduling strategies that balance pods based on affinity and
	// node taint violations.
	AffinityAndTaints DeschedulerProfile = "AffinityAndTaints"

	// TopologyAndDuplicates attempts to spread pods evenly among nodes based on topology spread
	// constraints and duplicate replicas on the same node.
	TopologyAndDuplicates DeschedulerProfile = "TopologyAndDuplicates"

	// LifecycleAndUtilization attempts to balance pods based on node resource usage, pod age, and pod restarts
	LifecycleAndUtilization DeschedulerProfile = "LifecycleAndUtilization"

	// DevPreviewLongLifecycle handles cluster lifecycle over a long term
	DevPreviewLongLifecycle DeschedulerProfile = "DevPreviewLongLifecycle"
)

// KubeDeschedulerStatus defines the observed state of KubeDescheduler
type KubeDeschedulerStatus struct {
	operatorv1.OperatorStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeDeschedulerList contains a list of KubeDescheduler
type KubeDeschedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeDescheduler `json:"items"`
}
