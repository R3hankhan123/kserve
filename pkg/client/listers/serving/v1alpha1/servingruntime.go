/*
Copyright 2023 The KServe Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kserve/kserve/pkg/apis/serving/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ServingRuntimeLister helps list ServingRuntimes.
// All objects returned here must be treated as read-only.
type ServingRuntimeLister interface {
	// List lists all ServingRuntimes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServingRuntime, err error)
	// ServingRuntimes returns an object that can list and get ServingRuntimes.
	ServingRuntimes(namespace string) ServingRuntimeNamespaceLister
	ServingRuntimeListerExpansion
}

// servingRuntimeLister implements the ServingRuntimeLister interface.
type servingRuntimeLister struct {
	listers.ResourceIndexer[*v1alpha1.ServingRuntime]
}

// NewServingRuntimeLister returns a new ServingRuntimeLister.
func NewServingRuntimeLister(indexer cache.Indexer) ServingRuntimeLister {
	return &servingRuntimeLister{listers.New[*v1alpha1.ServingRuntime](indexer, v1alpha1.Resource("servingruntime"))}
}

// ServingRuntimes returns an object that can list and get ServingRuntimes.
func (s *servingRuntimeLister) ServingRuntimes(namespace string) ServingRuntimeNamespaceLister {
	return servingRuntimeNamespaceLister{listers.NewNamespaced[*v1alpha1.ServingRuntime](s.ResourceIndexer, namespace)}
}

// ServingRuntimeNamespaceLister helps list and get ServingRuntimes.
// All objects returned here must be treated as read-only.
type ServingRuntimeNamespaceLister interface {
	// List lists all ServingRuntimes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServingRuntime, err error)
	// Get retrieves the ServingRuntime from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServingRuntime, error)
	ServingRuntimeNamespaceListerExpansion
}

// servingRuntimeNamespaceLister implements the ServingRuntimeNamespaceLister
// interface.
type servingRuntimeNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.ServingRuntime]
}
