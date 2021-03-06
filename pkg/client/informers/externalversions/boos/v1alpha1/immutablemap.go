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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	boosv1alpha1 "github.com/mattmoor/boo-maps/pkg/apis/boos/v1alpha1"
	versioned "github.com/mattmoor/boo-maps/pkg/client/clientset/versioned"
	internalinterfaces "github.com/mattmoor/boo-maps/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/mattmoor/boo-maps/pkg/client/listers/boos/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ImmutableMapInformer provides access to a shared informer and lister for
// ImmutableMaps.
type ImmutableMapInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ImmutableMapLister
}

type immutableMapInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewImmutableMapInformer constructs a new informer for ImmutableMap type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewImmutableMapInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredImmutableMapInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredImmutableMapInformer constructs a new informer for ImmutableMap type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredImmutableMapInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BoosV1alpha1().ImmutableMaps(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BoosV1alpha1().ImmutableMaps(namespace).Watch(options)
			},
		},
		&boosv1alpha1.ImmutableMap{},
		resyncPeriod,
		indexers,
	)
}

func (f *immutableMapInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredImmutableMapInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *immutableMapInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&boosv1alpha1.ImmutableMap{}, f.defaultInformer)
}

func (f *immutableMapInformer) Lister() v1alpha1.ImmutableMapLister {
	return v1alpha1.NewImmutableMapLister(f.Informer().GetIndexer())
}
