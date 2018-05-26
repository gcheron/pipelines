// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubeflow/pipelines/pkg/apis/scheduledworkflow/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScheduledWorkflowLister helps list ScheduledWorkflows.
type ScheduledWorkflowLister interface {
	// List lists all ScheduledWorkflows in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ScheduledWorkflow, err error)
	// ScheduledWorkflows returns an object that can list and get ScheduledWorkflows.
	ScheduledWorkflows(namespace string) ScheduledWorkflowNamespaceLister
	ScheduledWorkflowListerExpansion
}

// scheduledWorkflowLister implements the ScheduledWorkflowLister interface.
type scheduledWorkflowLister struct {
	indexer cache.Indexer
}

// NewScheduledWorkflowLister returns a new ScheduledWorkflowLister.
func NewScheduledWorkflowLister(indexer cache.Indexer) ScheduledWorkflowLister {
	return &scheduledWorkflowLister{indexer: indexer}
}

// List lists all ScheduledWorkflows in the indexer.
func (s *scheduledWorkflowLister) List(selector labels.Selector) (ret []*v1alpha1.ScheduledWorkflow, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScheduledWorkflow))
	})
	return ret, err
}

// ScheduledWorkflows returns an object that can list and get ScheduledWorkflows.
func (s *scheduledWorkflowLister) ScheduledWorkflows(namespace string) ScheduledWorkflowNamespaceLister {
	return scheduledWorkflowNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScheduledWorkflowNamespaceLister helps list and get ScheduledWorkflows.
type ScheduledWorkflowNamespaceLister interface {
	// List lists all ScheduledWorkflows in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ScheduledWorkflow, err error)
	// Get retrieves the ScheduledWorkflow from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ScheduledWorkflow, error)
	ScheduledWorkflowNamespaceListerExpansion
}

// scheduledWorkflowNamespaceLister implements the ScheduledWorkflowNamespaceLister
// interface.
type scheduledWorkflowNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScheduledWorkflows in the indexer for a given namespace.
func (s scheduledWorkflowNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ScheduledWorkflow, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScheduledWorkflow))
	})
	return ret, err
}

// Get retrieves the ScheduledWorkflow from the indexer for a given namespace and name.
func (s scheduledWorkflowNamespaceLister) Get(name string) (*v1alpha1.ScheduledWorkflow, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("scheduledworkflow"), name)
	}
	return obj.(*v1alpha1.ScheduledWorkflow), nil
}
