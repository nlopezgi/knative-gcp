/*
Copyright 2020 Google LLC

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

	pubsubv1alpha1 "github.com/google/knative-gcp/pkg/apis/pubsub/v1alpha1"
	versioned "github.com/google/knative-gcp/pkg/client/clientset/versioned"
	internalinterfaces "github.com/google/knative-gcp/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/google/knative-gcp/pkg/client/listers/pubsub/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TriggerInformer provides access to a shared informer and lister for
// Triggers.
type TriggerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TriggerLister
}

type triggerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTriggerInformer constructs a new informer for Trigger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTriggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTriggerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTriggerInformer constructs a new informer for Trigger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTriggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PubsubV1alpha1().Triggers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PubsubV1alpha1().Triggers(namespace).Watch(options)
			},
		},
		&pubsubv1alpha1.Trigger{},
		resyncPeriod,
		indexers,
	)
}

func (f *triggerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTriggerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *triggerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&pubsubv1alpha1.Trigger{}, f.defaultInformer)
}

func (f *triggerInformer) Lister() v1alpha1.TriggerLister {
	return v1alpha1.NewTriggerLister(f.Informer().GetIndexer())
}
