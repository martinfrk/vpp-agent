// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package syncbase contains reusable structures in multiple datasync transports.
// See following reusable structures:
// - Watcher that maintains the registrations/subscriptions.
// - Registry of latest revisions of values per each key synchronized by datasync.
// - Default implementation of Events & Iterators interfaces defined in data_api.go.
//   Events & Iterators in this package are reused (but not in all datasync transports).
package syncbase
