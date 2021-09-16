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

import (
        "testing"
        "github.com/stretchr/testify/assert"
)

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func initSdlInstanceMockTest() (sdlInstanceMockTest *MockSdlInstance) {
        sdlInstanceMockTest = new(MockSdlInstance)
        return
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestRemoveAll(t *testing.T){
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("RemoveAll").Return(nil)
        err := sdlInstanceMockTest.RemoveAll()
        assert.Nil(t, err)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestRemove(t *testing.T){
        var data []string
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("Remove", []string(data)).Return(nil)
        err := sdlInstanceMockTest.Remove(data)
        assert.Nil(t, err)

}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestRemoveIf(t *testing.T){
        var data map[string]interface{}
	key := "key"
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("RemoveIf", key, data).Return(true,nil)
        res, err := sdlInstanceMockTest.RemoveIf(key, data)
        assert.Nil(t, err)
        assert.NotNil(t, res)

}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestRemoveGroup(t *testing.T){
	group := "group"
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("RemoveGroup", group).Return(nil)
        err := sdlInstanceMockTest.RemoveGroup(group)
        assert.Nil(t, err)

}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestRemoveIfAndPublish(t *testing.T){
        var data map[string]interface{}
        var channelsAndEvents []string
	key := "key"
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("RemoveIfAndPublish", channelsAndEvents, key, data).Return(true,nil)
        res, err := sdlInstanceMockTest.RemoveIfAndPublish(channelsAndEvents, key, data)
        assert.Nil(t, err)
        assert.NotNil(t, res)

}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestRemoveAndPublish(t *testing.T){
        var channelsAndEvents []string
        var keys []string
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("RemoveAndPublish", []string(channelsAndEvents), []string(keys)).Return(nil)
        err := sdlInstanceMockTest.RemoveAndPublish(channelsAndEvents, keys)
        assert.Nil(t, err)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestRemoveAllAndPublish(t *testing.T){
        var channelsAndEvents []string
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("RemoveAllAndPublish", []string(channelsAndEvents)).Return(nil)
        err := sdlInstanceMockTest.RemoveAllAndPublish(channelsAndEvents)
        assert.Nil(t, err)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestIsMember(t *testing.T){
        var ret map[string]interface{}
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("IsMember", "group", ret).Return(true,nil)
        res, err := sdlInstanceMockTest.IsMember("group", ret)
        assert.Nil(t, err)
        assert.NotNil(t, res)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestClose(t *testing.T){
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("Close").Return(nil)
        err := sdlInstanceMockTest.Close()
        assert.Nil(t, err)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestSetIfNotExists(t *testing.T){
	var data map[string]interface{}
        key := "key"
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("SetIfNotExists", key, data).Return(true,nil)
        res, err := sdlInstanceMockTest.SetIfNotExists(key, data)
        assert.Nil(t, err)
        assert.NotNil(t, res)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestAddMember(t *testing.T){
        var ret []interface{}
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("AddMember", "group", []interface{}{ret}).Return(nil)
        err := sdlInstanceMockTest.AddMember("group", ret)
        assert.Nil(t, err)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestRemoveMember(t *testing.T){
        var ret []interface{}
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("RemoveMember", "group", []interface{}{ret}).Return(nil)
        err := sdlInstanceMockTest.RemoveMember("group", ret)
        assert.Nil(t, err)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestSetAndPublish(t *testing.T){
        var pairs []interface{}
	var channelsAndEvents []string
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("SetAndPublish", channelsAndEvents, []interface{}{pairs}).Return(nil)
        err := sdlInstanceMockTest.SetAndPublish(channelsAndEvents, pairs)
        assert.Nil(t, err)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestSetIfAndPublish(t *testing.T){
        var newData map[string]interface{}
        var oldData map[string]interface{}
        var group []string
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("SetIfAndPublish", group, "key", oldData, newData).Return(true, nil)
        res, err := sdlInstanceMockTest.SetIfAndPublish(group, "key", oldData, newData)
        assert.Nil(t, err)
        assert.NotNil(t, res)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestSet(t *testing.T){
        var pairs []interface{}
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("Set", []interface{}{pairs}).Return(nil)
        err := sdlInstanceMockTest.Set(pairs)
        assert.Nil(t, err)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestSetIf(t *testing.T){
        var newData map[string]interface{}
        var oldData map[string]interface{}
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("SetIf", "key", newData,  oldData).Return(true, nil)
        res, err := sdlInstanceMockTest.SetIf("key", newData,  oldData)
        assert.Nil(t, err)
        assert.NotNil(t, res)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestGetAll(t *testing.T){
        var data []string
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("GetAll").Return(data, nil)
        res, err := sdlInstanceMockTest.GetAll()
        assert.Nil(t, err)
        assert.Nil(t, res)
}

//Deprecated: Will be removed in a future release and tests in sdlSyncStorageMock_test.go
//should be used instead.
func TestSetIfNotExistsAndPublish(t *testing.T){
        var data map[string]interface{}
        var channelsAndEvents []string
        sdlInstanceMockTest := initSdlInstanceMockTest()
        sdlInstanceMockTest.On("SetIfNotExistsAndPublish", channelsAndEvents, "key", data).Return(true, nil)
        res, err := sdlInstanceMockTest.SetIfNotExistsAndPublish(channelsAndEvents, "key", data)
        assert.Nil(t, err)
        assert.NotNil(t, res)
}
