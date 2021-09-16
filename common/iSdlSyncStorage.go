//
// Copyright 2021 AT&T Intellectual Property
// Copyright 2021 Nokia
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

//  This source code is part of the near-RT RIC (RAN Intelligent Controller)
//  platform project (RICP).

package common

// ISdlSyncStorage integrates (wraps) the functionality that sdlgo SyncStorage provides
type ISdlSyncStorage interface {
	SubscribeChannel(ns string, cb func(string, ...string), channels ...string) error
	UnsubscribeChannel(ns string, channels ...string) error
	Close() error
	SetAndPublish(ns string, channelsAndEvents []string, pairs ...interface{}) error
	Set(ns string, pairs ...interface{}) error
	Get(ns string, keys []string) (map[string]interface{}, error)
	SetIfAndPublish(ns string, channelsAndEvents []string, key string, oldData, newData interface{}) (bool, error)
	SetIf(ns string, key string, oldData, newData interface{}) (bool, error)
	SetIfNotExistsAndPublish(ns string, channelsAndEvents []string, key string, data interface{}) (bool, error)
	SetIfNotExists(ns string, key string, data interface{}) (bool, error)
	RemoveAndPublish(ns string, channelsAndEvents []string, keys []string) error
	Remove(ns string, keys []string) error
	RemoveIfAndPublish(ns string, channelsAndEvents []string, key string, data interface{}) (bool, error)
	RemoveIf(ns string, key string, data interface{}) (bool, error)
	GetAll(ns string) ([]string, error)
	RemoveAll(ns string) error
	RemoveAllAndPublish(ns string, channelsAndEvents []string) error
	AddMember(ns string, group string, member ...interface{}) error
	RemoveMember(ns string, group string, member ...interface{}) error
	RemoveGroup(ns string, group string) error
	GetMembers(ns string, group string) ([]string, error)
	IsMember(ns string, group string, member interface{}) (bool, error)
	GroupSize(ns string, group string) (int64, error)
}
