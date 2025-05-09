/*
Copyright 2021 The Argoproj Authors.

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
	eventsv1alpha1 "github.com/argoproj/argo-events/pkg/apis/events/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// SensorLister helps list Sensors.
// All objects returned here must be treated as read-only.
type SensorLister interface {
	// List lists all Sensors in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*eventsv1alpha1.Sensor, err error)
	// Sensors returns an object that can list and get Sensors.
	Sensors(namespace string) SensorNamespaceLister
	SensorListerExpansion
}

// sensorLister implements the SensorLister interface.
type sensorLister struct {
	listers.ResourceIndexer[*eventsv1alpha1.Sensor]
}

// NewSensorLister returns a new SensorLister.
func NewSensorLister(indexer cache.Indexer) SensorLister {
	return &sensorLister{listers.New[*eventsv1alpha1.Sensor](indexer, eventsv1alpha1.Resource("sensor"))}
}

// Sensors returns an object that can list and get Sensors.
func (s *sensorLister) Sensors(namespace string) SensorNamespaceLister {
	return sensorNamespaceLister{listers.NewNamespaced[*eventsv1alpha1.Sensor](s.ResourceIndexer, namespace)}
}

// SensorNamespaceLister helps list and get Sensors.
// All objects returned here must be treated as read-only.
type SensorNamespaceLister interface {
	// List lists all Sensors in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*eventsv1alpha1.Sensor, err error)
	// Get retrieves the Sensor from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*eventsv1alpha1.Sensor, error)
	SensorNamespaceListerExpansion
}

// sensorNamespaceLister implements the SensorNamespaceLister
// interface.
type sensorNamespaceLister struct {
	listers.ResourceIndexer[*eventsv1alpha1.Sensor]
}
