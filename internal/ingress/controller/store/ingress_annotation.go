/*
Copyright 2015 The Kubernetes Authors.

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

package store

import (
	"github.com/kubernetes-sigs/aws-alb-ingress-controller/internal/ingress/annotations"
	"k8s.io/client-go/tools/cache"
)

// IngressAnnotationsLister makes a Store that lists annotations in Ingress rules.
type IngressAnnotationsLister struct {
	cache.Store
}

// ByKey returns the Ingress annotations matching key in the local Ingress annotations Store.
func (il IngressAnnotationsLister) ByKey(key string) (*annotations.Ingress, error) {
	i, exists, err := il.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, NotExistsError(key)
	}
	if i.(*annotations.Ingress).Error != nil {
		return nil, i.(*annotations.Ingress).Error
	}
	return i.(*annotations.Ingress), nil
}
