/*
Copyright The Kubernetes Authors.

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

package fake

import (
	typ0crdv1 "operator-v1/pkg/apis/typ0crd/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTyp0crds implements Typ0crdInterface
type FakeTyp0crds struct {
	Fake *FakeCrdV1
	ns   string
}

var typ0crdsResource = schema.GroupVersionResource{Group: "crd.devcluster.com", Version: "v1", Resource: "typ0crds"}

var typ0crdsKind = schema.GroupVersionKind{Group: "crd.devcluster.com", Version: "v1", Kind: "Typ0crd"}

// Get takes name of the typ0crd, and returns the corresponding typ0crd object, and an error if there is any.
func (c *FakeTyp0crds) Get(name string, options v1.GetOptions) (result *typ0crdv1.Typ0crd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(typ0crdsResource, c.ns, name), &typ0crdv1.Typ0crd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*typ0crdv1.Typ0crd), err
}

// List takes label and field selectors, and returns the list of Typ0crds that match those selectors.
func (c *FakeTyp0crds) List(opts v1.ListOptions) (result *typ0crdv1.Typ0crdList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(typ0crdsResource, typ0crdsKind, c.ns, opts), &typ0crdv1.Typ0crdList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &typ0crdv1.Typ0crdList{ListMeta: obj.(*typ0crdv1.Typ0crdList).ListMeta}
	for _, item := range obj.(*typ0crdv1.Typ0crdList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested typ0crds.
func (c *FakeTyp0crds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(typ0crdsResource, c.ns, opts))

}

// Create takes the representation of a typ0crd and creates it.  Returns the server's representation of the typ0crd, and an error, if there is any.
func (c *FakeTyp0crds) Create(typ0crd *typ0crdv1.Typ0crd) (result *typ0crdv1.Typ0crd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(typ0crdsResource, c.ns, typ0crd), &typ0crdv1.Typ0crd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*typ0crdv1.Typ0crd), err
}

// Update takes the representation of a typ0crd and updates it. Returns the server's representation of the typ0crd, and an error, if there is any.
func (c *FakeTyp0crds) Update(typ0crd *typ0crdv1.Typ0crd) (result *typ0crdv1.Typ0crd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(typ0crdsResource, c.ns, typ0crd), &typ0crdv1.Typ0crd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*typ0crdv1.Typ0crd), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTyp0crds) UpdateStatus(typ0crd *typ0crdv1.Typ0crd) (*typ0crdv1.Typ0crd, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(typ0crdsResource, "status", c.ns, typ0crd), &typ0crdv1.Typ0crd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*typ0crdv1.Typ0crd), err
}

// Delete takes name of the typ0crd and deletes it. Returns an error if one occurs.
func (c *FakeTyp0crds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(typ0crdsResource, c.ns, name), &typ0crdv1.Typ0crd{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTyp0crds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(typ0crdsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &typ0crdv1.Typ0crdList{})
	return err
}

// Patch applies the patch and returns the patched typ0crd.
func (c *FakeTyp0crds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *typ0crdv1.Typ0crd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(typ0crdsResource, c.ns, name, pt, data, subresources...), &typ0crdv1.Typ0crd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*typ0crdv1.Typ0crd), err
}
