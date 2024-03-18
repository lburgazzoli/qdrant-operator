/*
Copyright 2023.

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
	"context"
	time "time"

	qdrantv1alpha1 "github.com/megacamelus/qdrant-operator/api/qdrant/v1alpha1"
	versioned "github.com/megacamelus/qdrant-operator/pkg/client/qdrant/clientset/versioned"
	internalinterfaces "github.com/megacamelus/qdrant-operator/pkg/client/qdrant/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/megacamelus/qdrant-operator/pkg/client/qdrant/listers/qdrant/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CollectionInformer provides access to a shared informer and lister for
// Collections.
type CollectionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CollectionLister
}

type collectionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCollectionInformer constructs a new informer for Collection type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCollectionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCollectionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCollectionInformer constructs a new informer for Collection type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCollectionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QdrantV1alpha1().Collections(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QdrantV1alpha1().Collections(namespace).Watch(context.TODO(), options)
			},
		},
		&qdrantv1alpha1.Collection{},
		resyncPeriod,
		indexers,
	)
}

func (f *collectionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCollectionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *collectionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&qdrantv1alpha1.Collection{}, f.defaultInformer)
}

func (f *collectionInformer) Lister() v1alpha1.CollectionLister {
	return v1alpha1.NewCollectionLister(f.Informer().GetIndexer())
}
