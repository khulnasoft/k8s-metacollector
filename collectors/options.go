// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Khulnasoft Authors
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

package collectors

import (
	"github.com/khulnasoft/k8s-metacollector/pkg/subscriber"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type collectorOptions struct {
	externalSource    source.Source
	subscriberChan    subscriber.SubsChan
	podMatchingFields func(metadata *metav1.ObjectMeta) client.ListOption
	ownerSources      map[string]chan<- event.GenericEvent
}

// CollectorOption function used to set options when creating a new meta collector.
type CollectorOption func(opt *collectorOptions)

// WithExternalSource configure external sources that could trigger the reconcile loop of the collector.
func WithExternalSource(src source.Source) CollectorOption {
	return func(opt *collectorOptions) {
		opt.externalSource = src
	}
}

// WithSubscribersChan configures the subscriber channel.
func WithSubscribersChan(sChan subscriber.SubsChan) CollectorOption {
	return func(opt *collectorOptions) {
		opt.subscriberChan = sChan
	}
}

// WithPodMatchingFields configures the field selector used in the list operations.
func WithPodMatchingFields(podMatchingFields func(metadata *metav1.ObjectMeta) client.ListOption) CollectorOption {
	return func(opt *collectorOptions) {
		opt.podMatchingFields = podMatchingFields
	}
}

// WithOwnerSources a map holding channels used to trigger owner's reconciles.
func WithOwnerSources(sources map[string]chan<- event.GenericEvent) CollectorOption {
	return func(opt *collectorOptions) {
		opt.ownerSources = sources
	}
}
