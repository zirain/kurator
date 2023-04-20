/*
Copyright Kurator Authors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kurator.dev/kurator/pkg/apis/telemetry/v1alpha1"
)

// TelemetryLister helps list Telemetries.
// All objects returned here must be treated as read-only.
type TelemetryLister interface {
	// List lists all Telemetries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Telemetry, err error)
	// Telemetries returns an object that can list and get Telemetries.
	Telemetries(namespace string) TelemetryNamespaceLister
	TelemetryListerExpansion
}

// telemetryLister implements the TelemetryLister interface.
type telemetryLister struct {
	indexer cache.Indexer
}

// NewTelemetryLister returns a new TelemetryLister.
func NewTelemetryLister(indexer cache.Indexer) TelemetryLister {
	return &telemetryLister{indexer: indexer}
}

// List lists all Telemetries in the indexer.
func (s *telemetryLister) List(selector labels.Selector) (ret []*v1alpha1.Telemetry, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Telemetry))
	})
	return ret, err
}

// Telemetries returns an object that can list and get Telemetries.
func (s *telemetryLister) Telemetries(namespace string) TelemetryNamespaceLister {
	return telemetryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TelemetryNamespaceLister helps list and get Telemetries.
// All objects returned here must be treated as read-only.
type TelemetryNamespaceLister interface {
	// List lists all Telemetries in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Telemetry, err error)
	// Get retrieves the Telemetry from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Telemetry, error)
	TelemetryNamespaceListerExpansion
}

// telemetryNamespaceLister implements the TelemetryNamespaceLister
// interface.
type telemetryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Telemetries in the indexer for a given namespace.
func (s telemetryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Telemetry, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Telemetry))
	})
	return ret, err
}

// Get retrieves the Telemetry from the indexer for a given namespace and name.
func (s telemetryNamespaceLister) Get(name string) (*v1alpha1.Telemetry, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("telemetry"), name)
	}
	return obj.(*v1alpha1.Telemetry), nil
}
