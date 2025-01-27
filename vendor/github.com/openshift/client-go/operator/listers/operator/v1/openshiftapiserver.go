// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// OpenShiftAPIServerLister helps list OpenShiftAPIServers.
// All objects returned here must be treated as read-only.
type OpenShiftAPIServerLister interface {
	// List lists all OpenShiftAPIServers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*operatorv1.OpenShiftAPIServer, err error)
	// Get retrieves the OpenShiftAPIServer from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*operatorv1.OpenShiftAPIServer, error)
	OpenShiftAPIServerListerExpansion
}

// openShiftAPIServerLister implements the OpenShiftAPIServerLister interface.
type openShiftAPIServerLister struct {
	listers.ResourceIndexer[*operatorv1.OpenShiftAPIServer]
}

// NewOpenShiftAPIServerLister returns a new OpenShiftAPIServerLister.
func NewOpenShiftAPIServerLister(indexer cache.Indexer) OpenShiftAPIServerLister {
	return &openShiftAPIServerLister{listers.New[*operatorv1.OpenShiftAPIServer](indexer, operatorv1.Resource("openshiftapiserver"))}
}
