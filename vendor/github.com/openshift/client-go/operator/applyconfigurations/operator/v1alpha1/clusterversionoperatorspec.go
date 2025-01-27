// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/openshift/api/operator/v1"
)

// ClusterVersionOperatorSpecApplyConfiguration represents a declarative configuration of the ClusterVersionOperatorSpec type for use
// with apply.
type ClusterVersionOperatorSpecApplyConfiguration struct {
	OperatorLogLevel *v1.LogLevel `json:"operatorLogLevel,omitempty"`
}

// ClusterVersionOperatorSpecApplyConfiguration constructs a declarative configuration of the ClusterVersionOperatorSpec type for use with
// apply.
func ClusterVersionOperatorSpec() *ClusterVersionOperatorSpecApplyConfiguration {
	return &ClusterVersionOperatorSpecApplyConfiguration{}
}

// WithOperatorLogLevel sets the OperatorLogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OperatorLogLevel field is set to the value of the last call.
func (b *ClusterVersionOperatorSpecApplyConfiguration) WithOperatorLogLevel(value v1.LogLevel) *ClusterVersionOperatorSpecApplyConfiguration {
	b.OperatorLogLevel = &value
	return b
}