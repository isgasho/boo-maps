/*
Copyright 2019 Matt Moore

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
	v1alpha1 "github.com/mattmoor/boo-maps/pkg/apis/boos/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ImmutableMapLister helps list ImmutableMaps.
type ImmutableMapLister interface {
	// List lists all ImmutableMaps in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ImmutableMap, err error)
	// ImmutableMaps returns an object that can list and get ImmutableMaps.
	ImmutableMaps(namespace string) ImmutableMapNamespaceLister
	ImmutableMapListerExpansion
}

// immutableMapLister implements the ImmutableMapLister interface.
type immutableMapLister struct {
	indexer cache.Indexer
}

// NewImmutableMapLister returns a new ImmutableMapLister.
func NewImmutableMapLister(indexer cache.Indexer) ImmutableMapLister {
	return &immutableMapLister{indexer: indexer}
}

// List lists all ImmutableMaps in the indexer.
func (s *immutableMapLister) List(selector labels.Selector) (ret []*v1alpha1.ImmutableMap, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ImmutableMap))
	})
	return ret, err
}

// ImmutableMaps returns an object that can list and get ImmutableMaps.
func (s *immutableMapLister) ImmutableMaps(namespace string) ImmutableMapNamespaceLister {
	return immutableMapNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ImmutableMapNamespaceLister helps list and get ImmutableMaps.
type ImmutableMapNamespaceLister interface {
	// List lists all ImmutableMaps in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ImmutableMap, err error)
	// Get retrieves the ImmutableMap from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ImmutableMap, error)
	ImmutableMapNamespaceListerExpansion
}

// immutableMapNamespaceLister implements the ImmutableMapNamespaceLister
// interface.
type immutableMapNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ImmutableMaps in the indexer for a given namespace.
func (s immutableMapNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ImmutableMap, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ImmutableMap))
	})
	return ret, err
}

// Get retrieves the ImmutableMap from the indexer for a given namespace and name.
func (s immutableMapNamespaceLister) Get(name string) (*v1alpha1.ImmutableMap, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("immutablemap"), name)
	}
	return obj.(*v1alpha1.ImmutableMap), nil
}