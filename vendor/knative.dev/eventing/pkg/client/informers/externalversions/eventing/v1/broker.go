/*
Copyright 2020 The Knative Authors

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

package v1

import (
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	eventingv1 "knative.dev/eventing/pkg/apis/eventing/v1"
	versioned "knative.dev/eventing/pkg/client/clientset/versioned"
	internalinterfaces "knative.dev/eventing/pkg/client/informers/externalversions/internalinterfaces"
	v1 "knative.dev/eventing/pkg/client/listers/eventing/v1"
)

// BrokerInformer provides access to a shared informer and lister for
// Brokers.
type BrokerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.BrokerLister
}

type brokerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBrokerInformer constructs a new informer for Broker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBrokerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBrokerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBrokerInformer constructs a new informer for Broker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBrokerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventingV1().Brokers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventingV1().Brokers(namespace).Watch(context.TODO(), options)
			},
		},
		&eventingv1.Broker{},
		resyncPeriod,
		indexers,
	)
}

func (f *brokerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBrokerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *brokerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eventingv1.Broker{}, f.defaultInformer)
}

func (f *brokerInformer) Lister() v1.BrokerLister {
	return v1.NewBrokerLister(f.Informer().GetIndexer())
}
