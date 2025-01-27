// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	configv1alpha1 "github.com/openshift/api/config/v1alpha1"
	applyconfigurationsconfigv1alpha1 "github.com/openshift/client-go/config/applyconfigurations/config/v1alpha1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ClusterMonitoringsGetter has a method to return a ClusterMonitoringInterface.
// A group's client should implement this interface.
type ClusterMonitoringsGetter interface {
	ClusterMonitorings() ClusterMonitoringInterface
}

// ClusterMonitoringInterface has methods to work with ClusterMonitoring resources.
type ClusterMonitoringInterface interface {
	Create(ctx context.Context, clusterMonitoring *configv1alpha1.ClusterMonitoring, opts v1.CreateOptions) (*configv1alpha1.ClusterMonitoring, error)
	Update(ctx context.Context, clusterMonitoring *configv1alpha1.ClusterMonitoring, opts v1.UpdateOptions) (*configv1alpha1.ClusterMonitoring, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, clusterMonitoring *configv1alpha1.ClusterMonitoring, opts v1.UpdateOptions) (*configv1alpha1.ClusterMonitoring, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*configv1alpha1.ClusterMonitoring, error)
	List(ctx context.Context, opts v1.ListOptions) (*configv1alpha1.ClusterMonitoringList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *configv1alpha1.ClusterMonitoring, err error)
	Apply(ctx context.Context, clusterMonitoring *applyconfigurationsconfigv1alpha1.ClusterMonitoringApplyConfiguration, opts v1.ApplyOptions) (result *configv1alpha1.ClusterMonitoring, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, clusterMonitoring *applyconfigurationsconfigv1alpha1.ClusterMonitoringApplyConfiguration, opts v1.ApplyOptions) (result *configv1alpha1.ClusterMonitoring, err error)
	ClusterMonitoringExpansion
}

// clusterMonitorings implements ClusterMonitoringInterface
type clusterMonitorings struct {
	*gentype.ClientWithListAndApply[*configv1alpha1.ClusterMonitoring, *configv1alpha1.ClusterMonitoringList, *applyconfigurationsconfigv1alpha1.ClusterMonitoringApplyConfiguration]
}

// newClusterMonitorings returns a ClusterMonitorings
func newClusterMonitorings(c *ConfigV1alpha1Client) *clusterMonitorings {
	return &clusterMonitorings{
		gentype.NewClientWithListAndApply[*configv1alpha1.ClusterMonitoring, *configv1alpha1.ClusterMonitoringList, *applyconfigurationsconfigv1alpha1.ClusterMonitoringApplyConfiguration](
			"clustermonitorings",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *configv1alpha1.ClusterMonitoring { return &configv1alpha1.ClusterMonitoring{} },
			func() *configv1alpha1.ClusterMonitoringList { return &configv1alpha1.ClusterMonitoringList{} },
		),
	}
}