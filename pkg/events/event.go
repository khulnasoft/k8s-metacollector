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

package events

import (
	"fmt"

	"github.com/khulnasoft/k8s-metacollector/metadata"
	"github.com/khulnasoft/k8s-metacollector/pkg/fields"
)

const (
	// Create possible type for an event.
	Create = "Create"
	// Update possible type for an event.
	Update = "Update"
	// Delete possible type for an event.
	Delete = "Delete"
)

var _ Interface = &Event{}

// Event generated for watched kubernetes resources.
type Event struct {
	*metadata.Event
	Subs fields.Subscribers
}

// Subscribers returns the destination nodes.
func (ge *Event) Subscribers() fields.Subscribers {
	return ge.Subs
}

// String returns the event in string format.
func (ge *Event) String() string {
	return fmt.Sprintf("Resource Kind %q, event type %q, resource name %q, subscribers %q",
		ge.Kind, ge.Reason, ge.GetMeta(), ge.Subscribers())
}

// Type returns the event type.
func (ge *Event) Type() string {
	return ge.Reason
}

// ResourceKind returns the Kind of the resource for which the event has been crafted.
func (ge *Event) ResourceKind() string {
	return ge.Kind
}

// GRPCMessage returns the grpc message ready to be sent over the grpc connection.
func (ge *Event) GRPCMessage() *metadata.Event {
	return ge.Event
}
