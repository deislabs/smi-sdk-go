/*
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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha4

import (
	"context"
	"time"

	v1alpha4 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha4"
	scheme "github.com/servicemeshinterface/smi-sdk-go/pkg/gen/client/specs/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HTTPRouteGroupsGetter has a method to return a HTTPRouteGroupInterface.
// A group's client should implement this interface.
type HTTPRouteGroupsGetter interface {
	HTTPRouteGroups(namespace string) HTTPRouteGroupInterface
}

// HTTPRouteGroupInterface has methods to work with HTTPRouteGroup resources.
type HTTPRouteGroupInterface interface {
	Create(ctx context.Context, hTTPRouteGroup *v1alpha4.HTTPRouteGroup, opts v1.CreateOptions) (*v1alpha4.HTTPRouteGroup, error)
	Update(ctx context.Context, hTTPRouteGroup *v1alpha4.HTTPRouteGroup, opts v1.UpdateOptions) (*v1alpha4.HTTPRouteGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha4.HTTPRouteGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha4.HTTPRouteGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha4.HTTPRouteGroup, err error)
	HTTPRouteGroupExpansion
}

// hTTPRouteGroups implements HTTPRouteGroupInterface
type hTTPRouteGroups struct {
	client rest.Interface
	ns     string
}

// newHTTPRouteGroups returns a HTTPRouteGroups
func newHTTPRouteGroups(c *SpecsV1alpha4Client, namespace string) *hTTPRouteGroups {
	return &hTTPRouteGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hTTPRouteGroup, and returns the corresponding hTTPRouteGroup object, and an error if there is any.
func (c *hTTPRouteGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha4.HTTPRouteGroup, err error) {
	result = &v1alpha4.HTTPRouteGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httproutegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HTTPRouteGroups that match those selectors.
func (c *hTTPRouteGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha4.HTTPRouteGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha4.HTTPRouteGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httproutegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hTTPRouteGroups.
func (c *hTTPRouteGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("httproutegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hTTPRouteGroup and creates it.  Returns the server's representation of the hTTPRouteGroup, and an error, if there is any.
func (c *hTTPRouteGroups) Create(ctx context.Context, hTTPRouteGroup *v1alpha4.HTTPRouteGroup, opts v1.CreateOptions) (result *v1alpha4.HTTPRouteGroup, err error) {
	result = &v1alpha4.HTTPRouteGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("httproutegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hTTPRouteGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hTTPRouteGroup and updates it. Returns the server's representation of the hTTPRouteGroup, and an error, if there is any.
func (c *hTTPRouteGroups) Update(ctx context.Context, hTTPRouteGroup *v1alpha4.HTTPRouteGroup, opts v1.UpdateOptions) (result *v1alpha4.HTTPRouteGroup, err error) {
	result = &v1alpha4.HTTPRouteGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("httproutegroups").
		Name(hTTPRouteGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hTTPRouteGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hTTPRouteGroup and deletes it. Returns an error if one occurs.
func (c *hTTPRouteGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httproutegroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hTTPRouteGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httproutegroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hTTPRouteGroup.
func (c *hTTPRouteGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha4.HTTPRouteGroup, err error) {
	result = &v1alpha4.HTTPRouteGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("httproutegroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
