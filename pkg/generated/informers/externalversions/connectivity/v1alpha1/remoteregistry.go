// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	connectivityv1alpha1 "github.com/vmware-tanzu/cross-cluster-connectivity/apis/connectivity/v1alpha1"
	versioned "github.com/vmware-tanzu/cross-cluster-connectivity/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/vmware-tanzu/cross-cluster-connectivity/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/vmware-tanzu/cross-cluster-connectivity/pkg/generated/listers/connectivity/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RemoteRegistryInformer provides access to a shared informer and lister for
// RemoteRegistries.
type RemoteRegistryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RemoteRegistryLister
}

type remoteRegistryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRemoteRegistryInformer constructs a new informer for RemoteRegistry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRemoteRegistryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRemoteRegistryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRemoteRegistryInformer constructs a new informer for RemoteRegistry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRemoteRegistryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConnectivityV1alpha1().RemoteRegistries(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConnectivityV1alpha1().RemoteRegistries(namespace).Watch(options)
			},
		},
		&connectivityv1alpha1.RemoteRegistry{},
		resyncPeriod,
		indexers,
	)
}

func (f *remoteRegistryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRemoteRegistryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *remoteRegistryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&connectivityv1alpha1.RemoteRegistry{}, f.defaultInformer)
}

func (f *remoteRegistryInformer) Lister() v1alpha1.RemoteRegistryLister {
	return v1alpha1.NewRemoteRegistryLister(f.Informer().GetIndexer())
}
