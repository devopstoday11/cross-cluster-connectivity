// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/vmware-tanzu/cross-cluster-connectivity/apis/connectivity/v1alpha1"
	scheme "github.com/vmware-tanzu/cross-cluster-connectivity/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceRecordsGetter has a method to return a ServiceRecordInterface.
// A group's client should implement this interface.
type ServiceRecordsGetter interface {
	ServiceRecords(namespace string) ServiceRecordInterface
}

// ServiceRecordInterface has methods to work with ServiceRecord resources.
type ServiceRecordInterface interface {
	Create(*v1alpha1.ServiceRecord) (*v1alpha1.ServiceRecord, error)
	Update(*v1alpha1.ServiceRecord) (*v1alpha1.ServiceRecord, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServiceRecord, error)
	List(opts v1.ListOptions) (*v1alpha1.ServiceRecordList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceRecord, err error)
	ServiceRecordExpansion
}

// serviceRecords implements ServiceRecordInterface
type serviceRecords struct {
	client rest.Interface
	ns     string
}

// newServiceRecords returns a ServiceRecords
func newServiceRecords(c *ConnectivityV1alpha1Client, namespace string) *serviceRecords {
	return &serviceRecords{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceRecord, and returns the corresponding serviceRecord object, and an error if there is any.
func (c *serviceRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceRecord, err error) {
	result = &v1alpha1.ServiceRecord{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicerecords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceRecords that match those selectors.
func (c *serviceRecords) List(opts v1.ListOptions) (result *v1alpha1.ServiceRecordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceRecordList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicerecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceRecords.
func (c *serviceRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicerecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a serviceRecord and creates it.  Returns the server's representation of the serviceRecord, and an error, if there is any.
func (c *serviceRecords) Create(serviceRecord *v1alpha1.ServiceRecord) (result *v1alpha1.ServiceRecord, err error) {
	result = &v1alpha1.ServiceRecord{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicerecords").
		Body(serviceRecord).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceRecord and updates it. Returns the server's representation of the serviceRecord, and an error, if there is any.
func (c *serviceRecords) Update(serviceRecord *v1alpha1.ServiceRecord) (result *v1alpha1.ServiceRecord, err error) {
	result = &v1alpha1.ServiceRecord{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicerecords").
		Name(serviceRecord.Name).
		Body(serviceRecord).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceRecord and deletes it. Returns an error if one occurs.
func (c *serviceRecords) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicerecords").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicerecords").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceRecord.
func (c *serviceRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceRecord, err error) {
	result = &v1alpha1.ServiceRecord{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicerecords").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
