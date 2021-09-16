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

package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func initSdlSyncStorageMockTest() (sdlSyncStorageMockTest *MockSdlSyncStorage) {
	sdlSyncStorageMockTest = new(MockSdlSyncStorage)
	return
}

func TestRemoveAllMock(t *testing.T) {
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("RemoveAll", ns).Return(nil)
	err := sdlSyncStorageMockTest.RemoveAll(ns)
	assert.Nil(t, err)
}

func TestRemoveMock(t *testing.T) {
	var data []string
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("Remove", ns, []string(data)).Return(nil)
	err := sdlSyncStorageMockTest.Remove(ns, data)
	assert.Nil(t, err)

}

func TestRemoveIfMock(t *testing.T) {
	var data map[string]interface{}
	ns := "some-ns"
	key := "key"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("RemoveIf", ns, key, data).Return(true, nil)
	res, err := sdlSyncStorageMockTest.RemoveIf(ns, key, data)
	assert.Nil(t, err)
	assert.NotNil(t, res)

}

func TestRemoveGroupMock(t *testing.T) {
	ns := "some-ns"
	group := "group"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("RemoveGroup", ns, group).Return(nil)
	err := sdlSyncStorageMockTest.RemoveGroup(ns, group)
	assert.Nil(t, err)

}

func TestRemoveIfAndPublishMock(t *testing.T) {
	var data map[string]interface{}
	var channelsAndEvents []string
	ns := "some-ns"
	key := "key"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("RemoveIfAndPublish", ns, channelsAndEvents, key, data).Return(true, nil)
	res, err := sdlSyncStorageMockTest.RemoveIfAndPublish(ns, channelsAndEvents, key, data)
	assert.Nil(t, err)
	assert.NotNil(t, res)

}

func TestRemoveAndPublishMock(t *testing.T) {
	var channelsAndEvents []string
	var keys []string
	ns := "some-ns"

	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("RemoveAndPublish", ns, []string(channelsAndEvents), []string(keys)).Return(nil)
	err := sdlSyncStorageMockTest.RemoveAndPublish(ns, channelsAndEvents, keys)
	assert.Nil(t, err)
}

func TestRemoveAllAndPublishMock(t *testing.T) {
	var channelsAndEvents []string
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("RemoveAllAndPublish", ns, []string(channelsAndEvents)).Return(nil)
	err := sdlSyncStorageMockTest.RemoveAllAndPublish(ns, channelsAndEvents)
	assert.Nil(t, err)
}

func TestIsMemberMock(t *testing.T) {
	var ret map[string]interface{}
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("IsMember", ns, "group", ret).Return(true, nil)
	res, err := sdlSyncStorageMockTest.IsMember(ns, "group", ret)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestCloseMock(t *testing.T) {
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("Close").Return(nil)
	err := sdlSyncStorageMockTest.Close()
	assert.Nil(t, err)
}

func TestSetIfNotExistsMock(t *testing.T) {
	var data map[string]interface{}
	ns := "some-ns"
	key := "key"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("SetIfNotExists", ns, key, data).Return(true, nil)
	res, err := sdlSyncStorageMockTest.SetIfNotExists(ns, key, data)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestAddMemberMock(t *testing.T) {
	var ret []interface{}
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("AddMember", ns, "group", []interface{}{ret}).Return(nil)
	err := sdlSyncStorageMockTest.AddMember(ns, "group", ret)
	assert.Nil(t, err)
}

func TestRemoveMemberMock(t *testing.T) {
	var ret []interface{}
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("RemoveMember", ns, "group", []interface{}{ret}).Return(nil)
	err := sdlSyncStorageMockTest.RemoveMember(ns, "group", ret)
	assert.Nil(t, err)
}

func TestSetAndPublishMock(t *testing.T) {
	var pairs []interface{}
	var channelsAndEvents []string
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("SetAndPublish", ns, channelsAndEvents, []interface{}{pairs}).Return(nil)
	err := sdlSyncStorageMockTest.SetAndPublish(ns, channelsAndEvents, pairs)
	assert.Nil(t, err)
}

func TestSetIfAndPublishMock(t *testing.T) {
	var newData map[string]interface{}
	var oldData map[string]interface{}
	var group []string
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("SetIfAndPublish", ns, group, "key", oldData, newData).Return(true, nil)
	res, err := sdlSyncStorageMockTest.SetIfAndPublish(ns, group, "key", oldData, newData)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestSetMock(t *testing.T) {
	var pairs []interface{}
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("Set", ns, []interface{}{pairs}).Return(nil)
	err := sdlSyncStorageMockTest.Set(ns, pairs)
	assert.Nil(t, err)
}

func TestSetIfMock(t *testing.T) {
	var newData map[string]interface{}
	var oldData map[string]interface{}
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("SetIf", ns, "key", newData, oldData).Return(true, nil)
	res, err := sdlSyncStorageMockTest.SetIf(ns, "key", newData, oldData)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestGetAllMock(t *testing.T) {
	var data []string
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("GetAll", ns).Return(data, nil)
	res, err := sdlSyncStorageMockTest.GetAll(ns)
	assert.Nil(t, err)
	assert.Nil(t, res)
}

func TestSetIfNotExistsAndPublishMock(t *testing.T) {
	var data map[string]interface{}
	var channelsAndEvents []string
	ns := "some-ns"
	sdlSyncStorageMockTest := initSdlSyncStorageMockTest()
	sdlSyncStorageMockTest.On("SetIfNotExistsAndPublish", ns, channelsAndEvents, "key", data).Return(true, nil)
	res, err := sdlSyncStorageMockTest.SetIfNotExistsAndPublish(ns, channelsAndEvents, "key", data)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
