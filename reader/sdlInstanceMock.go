//
// Copyright 2019 AT&T Intellectual Property
// Copyright 2019 Nokia
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

package reader

import "github.com/stretchr/testify/mock"

//Deprecated: Will be removed in a future release, please use MockSdlSyncStorage instead.
type MockSdlInstance struct {
	mock.Mock
}

//Deprecated: Will be removed in a future release, please use instead the SubscribeChannel
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) SubscribeChannel(cb func(string, ...string), channels ...string) error {
	a := m.Called(cb, channels)
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the UnsubscribeChannel
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) UnsubscribeChannel(channels ...string) error {
	a := m.Called(channels)
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the SetAndPublish
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) SetAndPublish(channelsAndEvents []string, pairs ...interface{}) error {
	a := m.Called(channelsAndEvents, pairs)
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the SetIfAndPublish
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) SetIfAndPublish(channelsAndEvents []string, key string, oldData, newData interface{}) (bool, error) {
	a := m.Called(channelsAndEvents, key, oldData, newData)
	return a.Bool(0), a.Error(1)
}

//Deprecated: Will be removed in a future release, please use instead the SetIfNotExistsAndPublish
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) SetIfNotExistsAndPublish(channelsAndEvents []string, key string, data interface{}) (bool, error) {
	a := m.Called(channelsAndEvents, key, data)
	return a.Bool(0), a.Error(1)
}

//Deprecated: Will be removed in a future release, please use instead the RemoveAndPublish
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) RemoveAndPublish(channelsAndEvents []string, keys []string) error {
	a := m.Called(channelsAndEvents, keys)
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the RemoveIfAndPublish
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) RemoveIfAndPublish(channelsAndEvents []string, key string, data interface{}) (bool, error) {
	a := m.Called(channelsAndEvents, key, data)
	return a.Bool(0), a.Error(1)
}

//Deprecated: Will be removed in a future release, please use instead the RemoveAllAndPublish
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) RemoveAllAndPublish(channelsAndEvents []string) error {
	a := m.Called(channelsAndEvents)
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the Set
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) Set(pairs ...interface{}) error {
	a := m.Called(pairs)
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the Get
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) Get(keys []string) (map[string]interface{}, error) {
	a := m.Called(keys)
	return a.Get(0).(map[string]interface{}), a.Error(1)
}

//Deprecated: Will be removed in a future release, please use instead the GetAll
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) GetAll() ([]string, error) {
	a := m.Called()
	return a.Get(0).([]string), a.Error(1)
}

//Deprecated: Will be removed in a future release, please use instead the Close
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) Close() error {
	a := m.Called()
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the Remove
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) Remove(keys []string) error {
	a := m.Called(keys)
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the RemoveAll
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) RemoveAll() error {
	a := m.Called()
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the SetIf
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) SetIf(key string, oldData, newData interface{}) (bool, error) {
	a := m.Called(key, oldData, newData)
	return a.Bool(0), a.Error(1)
}

//Deprecated: Will be removed in a future release, please use instead the SetIfNotExists
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) SetIfNotExists(key string, data interface{}) (bool, error) {
	a := m.Called(key, data)
	return a.Bool(0), a.Error(1)
}

//Deprecated: Will be removed in a future release, please use instead the RemoveIf
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) RemoveIf(key string, data interface{}) (bool, error) {
	a := m.Called(key, data)
	return a.Bool(0), a.Error(1)
}

//Deprecated: Will be removed in a future release, please use instead the AddMember
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) AddMember(group string, member ...interface{}) error {
	a := m.Called(group, member)
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the RemoveMember
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) RemoveMember(group string, member ...interface{}) error {
	a := m.Called(group, member)
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the RemoveGroup
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) RemoveGroup(group string) error {
	a := m.Called(group)
	return a.Error(0)
}

//Deprecated: Will be removed in a future release, please use instead the GetMembers
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) GetMembers(group string) ([]string, error) {
	a := m.Called(group)
	return a.Get(0).([]string), a.Error(1)
}

//Deprecated: Will be removed in a future release, please use instead the IsMember
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) IsMember(group string, member interface{}) (bool, error) {
	a := m.Called(group, member)
	return a.Bool(0), a.Error(1)
}

//Deprecated: Will be removed in a future release, please use instead the GroupSize
//receiver function of the MockSdlSyncStorage type.
func (m *MockSdlInstance) GroupSize(group string) (int64, error) {
	a := m.Called(group, )
	return int64(a.Int(0)), a.Error(1)
}
