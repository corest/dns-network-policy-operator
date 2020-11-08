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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DNSNetworkPolicySpec defines the desired state of DNSNetworkPolicy
type DNSNetworkPolicySpec struct {
	// Domains is the list of domain names, which should be resolved
	// before updating reference network policy.
	// e.g. ["kubernetes.local", "http://google.com", "https://twitter.com"]
	Domains []string `json:"domains,omitempty"`
	// targetNetworkPolicy is an existing network policy in the object namespace,
	// which is updated with resolved domains IP addresses.
	// e.g. memcached-network-policy
	TargetNetworkPolicy string `json:"targetNetworkPolicy,omitempty"`
}

// DNSNetworkPolicyStatus defines the observed state of DNSNetworkPolicy
type DNSNetworkPolicyStatus struct {
	// LastTargetNetworkPolicyUpdateTime is the last time targetNetworkPolicy was updated.
	LastTargetNetworkPolicyUpdateTime *metav1.Time `json:"lastTargetNetworkPolicyUpdateTime"`
}

// +kubebuilder:object:root=true

// DNSNetworkPolicy is the Schema for the dnsnetworkpolicies API
type DNSNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSNetworkPolicySpec   `json:"spec,omitempty"`
	Status DNSNetworkPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSNetworkPolicyList contains a list of DNSNetworkPolicy
type DNSNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSNetworkPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DNSNetworkPolicy{}, &DNSNetworkPolicyList{})
}
