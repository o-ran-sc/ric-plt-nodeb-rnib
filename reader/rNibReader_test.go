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

package reader

import (
	"errors"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/common"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

var namespace = "namespace"

func TestInit(t *testing.T) {
	readerPool = nil
	Init("", 1)
	assert.NotNil(t, readerPool)
	assert.NotNil(t, readerPool.New)
	assert.NotNil(t, readerPool.Destroy)
	available, created := readerPool.Stats()
	assert.Equal(t, 0, available, "number of available objects in the readerPool should be 0")
	assert.Equal(t, 0, created, "number of created objects in the readerPool should be 0")
}

func TestInitPool(t *testing.T) {
	readerPool = nil
	sdlInstanceMock := new(MockSdlInstance)
	initPool(1, func() interface{} {
		sdlI := common.ISdlInstance(sdlInstanceMock)
		return &rNibReaderInstance{sdl: &sdlI, namespace: namespace}
	},
		func(obj interface{}) {
		},
	)
	assert.NotNil(t, readerPool)
	assert.NotNil(t, readerPool.New)
	assert.NotNil(t, readerPool.Destroy)
	available, created := readerPool.Stats()
	assert.Equal(t, 0, available, "number of available objects in the readerPool should be 0")
	assert.Equal(t, 0, created, "number of created objects in the readerPool should be 0")
}

func initSdlInstanceMock(namespace string, poolSize int) *MockSdlInstance {
	sdlInstanceMock := new(MockSdlInstance)
	initPool(poolSize, func() interface{} {
		sdlI := common.ISdlInstance(sdlInstanceMock)
		return &rNibReaderInstance{sdl: &sdlI, namespace: namespace}
	},
		func(obj interface{}) {
		},
	)
	return sdlInstanceMock
}

func TestGetNodeB(t *testing.T) {
	name := "name"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	nb := entities.NodebInfo{}
	nb.ConnectionStatus = 1
	nb.Ip = "localhost"
	nb.Port = 5656
	enb := entities.Enb{}
	cell := entities.ServedCellInfo{Tac: "tac"}
	enb.ServedCells = []*entities.ServedCellInfo{&cell}
	nb.Configuration = &entities.NodebInfo_Enb{Enb: &enb}
	var e error
	data, err := proto.Marshal(&nb)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetNb - Failed to marshal ENB instance. Error: %v", err)
	}
	ret := map[string]interface{}{"RAN:" + name: string(data)}
	sdlInstanceMock.On("Get", []string{"RAN:" + name}).Return(ret, e)
	getNb, er := w.GetNodeb(name)
	assert.Nil(t, er)
	assert.Equal(t, getNb.Ip, nb.Ip)
	assert.Equal(t, getNb.Port, nb.Port)
	assert.Equal(t, getNb.ConnectionStatus, nb.ConnectionStatus)
	assert.Len(t, getNb.GetEnb().GetServedCells(), 1)
	assert.Equal(t, getNb.GetEnb().GetServedCells()[0].Tac, nb.GetEnb().GetServedCells()[0].Tac)
}

func TestGetNodeBNotFoundFailure(t *testing.T) {
	name := "name"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	var ret map[string]interface{}
	sdlInstanceMock.On("Get", []string{"RAN:" + name}).Return(ret, e)
	getNb, er := w.GetNodeb(name)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.Equal(t, 1, er.GetCode())
	assert.EqualValues(t, "1 RESOURCE_NOT_FOUND - #rNibReader.getNodeb - responding node not found. Key: RAN:name", er.Error())
}

func TestGetNodeBUnmarshalFailure(t *testing.T) {
	name := "name"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	ret := make(map[string]interface{}, 1)
	ret["RAN:"+name] = "data"
	sdlInstanceMock.On("Get", []string{"RAN:" + name}).Return(ret, e)
	getNb, er := w.GetNodeb(name)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, "2 INTERNAL_ERROR - proto: can't skip unknown wire type 4", er.Error())
}

func TestGetNodeBSdlgoFailure(t *testing.T) {
	name := "name"
	errMsg := "expected Sdlgo error"
	errMsgExpected := "2 INTERNAL_ERROR - expected Sdlgo error"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	e := errors.New(errMsg)
	var ret map[string]interface{}
	sdlInstanceMock.On("Get", []string{"RAN:" + name}).Return(ret, e)
	getNb, er := w.GetNodeb(name)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetNodeBCellsListEnb(t *testing.T) {
	name := "name"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	nb := entities.NodebInfo{}
	nb.ConnectionStatus = 1
	nb.Ip = "localhost"
	nb.Port = 5656
	enb := entities.Enb{}
	cell := entities.ServedCellInfo{Tac: "tac"}
	enb.ServedCells = []*entities.ServedCellInfo{&cell}
	nb.Configuration = &entities.NodebInfo_Enb{Enb: &enb}
	var e error
	data, err := proto.Marshal(&nb)
	if err != nil {
		t.Errorf("#rNibReader_test.GetNodeBCellsList - Failed to marshal ENB instance. Error: %v", err)
	}
	ret := map[string]interface{}{"RAN:" + name: string(data)}
	sdlInstanceMock.On("Get", []string{"RAN:" + name}).Return(ret, e)
	cells, er := w.GetCellList(name)
	assert.Nil(t, er)
	assert.NotNil(t, cells)
	assert.Len(t, cells.GetServedCellInfos().GetServedCells(), 1)
	retCell := cells.GetServedCellInfos().GetServedCells()[0]
	assert.Equal(t, retCell.Tac, "tac")
}

func TestGetNodeBCellsListGnb(t *testing.T) {
	name := "name"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	nb := entities.NodebInfo{}
	nb.ConnectionStatus = 1
	nb.Ip = "localhost"
	nb.Port = 5656
	nb.NodeType = entities.Node_GNB
	gnb := entities.Gnb{}
	cell := entities.ServedNRCell{ServedNrCellInformation: &entities.ServedNRCellInformation{NrPci: 10}}
	gnb.ServedNrCells = []*entities.ServedNRCell{&cell}
	nb.Configuration = &entities.NodebInfo_Gnb{Gnb: &gnb}
	var e error
	data, err := proto.Marshal(&nb)
	if err != nil {
		t.Errorf("#rNibReader_test.GetNodeBCellsList - Failed to marshal GNB instance. Error: %v", err)
	}
	ret := map[string]interface{}{"RAN:" + name: string(data)}
	sdlInstanceMock.On("Get", []string{"RAN:" + name}).Return(ret, e)
	cells, er := w.GetCellList(name)
	assert.Nil(t, er)
	assert.NotNil(t, cells)
	assert.Len(t, cells.GetServedNrCells().GetServedCells(), 1)
	retCell := cells.GetServedNrCells().GetServedCells()[0]
	assert.Equal(t, retCell.GetServedNrCellInformation().GetNrPci(), uint32(10))
}

func TestGetNodeBCellsListNodeUnmarshalFailure(t *testing.T) {
	name := "name"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	ret := make(map[string]interface{}, 1)
	ret["RAN:"+name] = "data"
	sdlInstanceMock.On("Get", []string{"RAN:" + name}).Return(ret, e)
	cells, er := w.GetCellList(name)
	assert.NotNil(t, er)
	assert.Nil(t, cells)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, "2 INTERNAL_ERROR - proto: can't skip unknown wire type 4", er.Error())
}

func TestGetNodeBCellsListNodeNotFoundFailure(t *testing.T) {
	name := "name"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	var ret map[string]interface{}
	sdlInstanceMock.On("Get", []string{"RAN:" + name}).Return(ret, e)
	cells, er := w.GetCellList(name)
	assert.NotNil(t, er)
	assert.Nil(t, cells)
	assert.Equal(t, 1, er.GetCode())
	assert.EqualValues(t, "1 RESOURCE_NOT_FOUND - #rNibReader.getNodeb - responding node not found. Key: RAN:name", er.Error())
}

func TestGetNodeBCellsListNotFoundFailureEnb(t *testing.T) {
	name := "name"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	nb := entities.NodebInfo{}
	nb.ConnectionStatus = 1
	nb.Ip = "localhost"
	nb.Port = 5656
	enb := entities.Enb{}
	nb.Configuration = &entities.NodebInfo_Enb{Enb: &enb}
	var e error
	data, err := proto.Marshal(&nb)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetNbCellsListNotFoundFailure - Failed to marshal ENB instance. Error: %v", err)
	}
	ret := map[string]interface{}{"RAN:" + name: string(data)}
	sdlInstanceMock.On("Get", []string{"RAN:" + name}).Return(ret, e)
	_, er := w.GetCellList(name)
	assert.NotNil(t, er)
	assert.EqualValues(t, "1 RESOURCE_NOT_FOUND - #rNibReader.GetCellList - served cells not found. Responding node RAN name: name.", er.Error())
}

func TestGetNodeBCellsListNotFoundFailureGnb(t *testing.T) {
	name := "name"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	nb := entities.NodebInfo{}
	nb.ConnectionStatus = 1
	nb.Ip = "localhost"
	nb.Port = 5656
	gnb := entities.Gnb{}
	nb.Configuration = &entities.NodebInfo_Gnb{Gnb: &gnb}
	var e error
	data, err := proto.Marshal(&nb)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetNbCellsListNotFoundFailure - Failed to marshal ENB instance. Error: %v", err)
	}
	ret := map[string]interface{}{"RAN:" + name: string(data)}
	sdlInstanceMock.On("Get", []string{"RAN:" + name}).Return(ret, e)
	_, er := w.GetCellList(name)
	assert.NotNil(t, er)
	assert.EqualValues(t, "1 RESOURCE_NOT_FOUND - #rNibReader.GetCellList - served cells not found. Responding node RAN name: name.", er.Error())
}

func TestCloseOnClosedPoolFailure(t *testing.T) {
	readerPool = nil
	instanceMock := initSdlInstanceMock(namespace, 1)
	w1 := GetRNibReader()
	readerPool.Put(w1)
	available, created := readerPool.Stats()
	assert.Equal(t, 1, available, "number of available objects in the readerPool should be 1")
	assert.Equal(t, 1, created, "number of created objects in the readerPool should be 1")
	var e error
	instanceMock.On("Close").Return(e)
	Close()
	assert.Panics(t, func() { Close() })
}

func TestCloseFailure(t *testing.T) {
	readerPool = nil
	instanceMock := initSdlInstanceMock(namespace, 2)
	w1 := GetRNibReader()
	readerPool.Put(w1)
	available, created := readerPool.Stats()
	assert.Equal(t, 1, available, "number of available objects in the readerPool should be 1")
	assert.Equal(t, 1, created, "number of created objects in the readerPool should be 1")
	e := errors.New("expected error")
	instanceMock.On("Close").Return(e)
	Close()
	available, created = readerPool.Stats()
	assert.Equal(t, 0, available, "number of available objects in the readerPool should be 0")
	assert.Equal(t, 0, created, "number of created objects in the readerPool should be 0")
}

func TestGetListGnbIdsUnmarshalFailure(t *testing.T) {
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	sdlInstanceMock.On("GetMembers", GnbType).Return([]string{"data"}, e)
	ids, er := w.GetListGnbIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.Equal(t, 2, er.GetCode())
	assert.Equal(t, "2 INTERNAL_ERROR - proto: can't skip unknown wire type 4", er.Error())
}

func TestGetListGnbIdsSdlgoFailure(t *testing.T) {
	errMsg := "expected Sdlgo error"
	errMsgExpected := "2 INTERNAL_ERROR - expected Sdlgo error"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	e := errors.New(errMsg)
	var data []string
	sdlInstanceMock.On("GetMembers", GnbType).Return(data, e)
	ids, er := w.GetListGnbIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetListNodesIdsGnbSdlgoFailure(t *testing.T) {

	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()

	name := "name"
	plmnId := "02f829"
	nbId := "4a952a0a"
	nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}}
	var nilError error
	data, err := proto.Marshal(nbIdentity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListNodesIdsGnbSdlgoFailure - Failed to marshal nodeb identity entity. Error: %v", err)
	}
	sdlInstanceMock.On("GetMembers", EnbType).Return([]string{string(data)}, nilError)

	errMsg := "expected Sdlgo error"
	errMsgExpected := "2 INTERNAL_ERROR - expected Sdlgo error"
	expectedError := errors.New(errMsg)
	var nilData []string
	sdlInstanceMock.On("GetMembers", GnbType).Return(nilData, expectedError)

	ids, er := w.GetListNodebIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetListNodesIdsEnbSdlgoFailure(t *testing.T) {

	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()

	name := "name"
	plmnId := "02f829"
	nbId := "4a952a0a"
	nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}}
	var nilError error
	data, err := proto.Marshal(nbIdentity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListNodesIdsEnbSdlgoFailure - Failed to marshal nodeb identity entity. Error: %v", err)
	}
	sdlInstanceMock.On("GetMembers", GnbType).Return([]string{string(data)}, nilError)

	errMsg := "expected Sdlgo error"
	errMsgExpected := "2 INTERNAL_ERROR - expected Sdlgo error"
	expectedError := errors.New(errMsg)
	var nilData []string
	sdlInstanceMock.On("GetMembers", EnbType).Return(nilData, expectedError)

	ids, er := w.GetListNodebIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetListNodesIdsEnbSdlgoSuccess(t *testing.T) {

	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()

	name := "name"
	plmnId := "02f829"
	nbId := "4a952a0a"
	nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}}
	var nilError error
	data, err := proto.Marshal(nbIdentity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListNodesIdsEnbSdlgoFailure - Failed to marshal nodeb identity entity. Error: %v", err)
	}
	sdlInstanceMock.On("GetMembers", GnbType).Return([]string{string(data)}, nilError)
	sdlInstanceMock.On("GetMembers", EnbType).Return([]string{string(data)}, nilError)

	ids, er := w.GetListNodebIds()
	assert.Nil(t, er)
	assert.NotNil(t, ids)
	assert.Len(t, ids, 2)
}

func TestGetListEnbIdsUnmarshalFailure(t *testing.T) {
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	sdlInstanceMock.On("GetMembers", EnbType).Return([]string{"data"}, e)
	ids, er := w.GetListEnbIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.Equal(t, 2, er.GetCode())
	assert.Equal(t, "2 INTERNAL_ERROR - proto: can't skip unknown wire type 4", er.Error())
}

func TestGetListEnbIdsOneId(t *testing.T) {
	name := "name"
	plmnId := "02f829"
	nbId := "4a952a0a"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}}
	var e error
	data, err := proto.Marshal(nbIdentity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListEnbIds - Failed to marshal nodeb identity entity. Error: %v", err)
	}
	sdlInstanceMock.On("GetMembers", EnbType).Return([]string{string(data)}, e)
	ids, er := w.GetListEnbIds()
	assert.Nil(t, er)
	assert.Len(t, *ids, 1)
	assert.Equal(t, (*ids)[0].GetInventoryName(), name)
	assert.Equal(t, (*ids)[0].GetGlobalNbId().GetPlmnId(), nbIdentity.GetGlobalNbId().GetPlmnId())
	assert.Equal(t, (*ids)[0].GetGlobalNbId().GetNbId(), nbIdentity.GetGlobalNbId().GetNbId())
}

func TestGetListEnbIdsNoIds(t *testing.T) {
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	sdlInstanceMock.On("GetMembers", EnbType).Return([]string{}, e)
	ids, er := w.GetListEnbIds()
	assert.Nil(t, er)
	assert.Len(t, *ids, 0)
}

func TestGetListEnbIds(t *testing.T) {
	name := "name"
	plmnId := 0x02f829
	nbId := 0x4a952a0a
	listSize := 3
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	idsData := make([]string, listSize)
	idsEntities := make([]*entities.NbIdentity, listSize)
	for i := 0; i < listSize; i++ {
		nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: string(plmnId + i), NbId: string(nbId + i)}}
		data, err := proto.Marshal(nbIdentity)
		if err != nil {
			t.Errorf("#rNibReader_test.TestGetListEnbIds - Failed to marshal nodeb identity entity. Error: %v", err)
		}
		idsData[i] = string(data)
		idsEntities[i] = nbIdentity
	}
	var e error
	sdlInstanceMock.On("GetMembers", EnbType).Return(idsData, e)
	ids, er := w.GetListEnbIds()
	assert.Nil(t, er)
	assert.Len(t, *ids, listSize)
	for i, id := range *ids {
		assert.Equal(t, id.GetInventoryName(), name)
		assert.Equal(t, id.GetGlobalNbId().GetPlmnId(), idsEntities[i].GetGlobalNbId().GetPlmnId())
		assert.Equal(t, id.GetGlobalNbId().GetNbId(), idsEntities[i].GetGlobalNbId().GetNbId())
	}
}

func TestGetListGnbIdsOneId(t *testing.T) {
	name := "name"
	plmnId := "02f829"
	nbId := "4a952a0a"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}}
	var e error
	data, err := proto.Marshal(nbIdentity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListGnbIds - Failed to marshal nodeb identity entity. Error: %v", err)
	}
	sdlInstanceMock.On("GetMembers", GnbType).Return([]string{string(data)}, e)
	ids, er := w.GetListGnbIds()
	assert.Nil(t, er)
	assert.Len(t, *ids, 1)
	assert.Equal(t, (*ids)[0].GetInventoryName(), name)
	assert.Equal(t, (*ids)[0].GetGlobalNbId().GetPlmnId(), nbIdentity.GetGlobalNbId().GetPlmnId())
	assert.Equal(t, (*ids)[0].GetGlobalNbId().GetNbId(), nbIdentity.GetGlobalNbId().GetNbId())
}

func TestGetListGnbIdsNoIds(t *testing.T) {
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	sdlInstanceMock.On("GetMembers", GnbType).Return([]string{}, e)
	ids, er := w.GetListGnbIds()
	assert.Nil(t, er)
	assert.Len(t, *ids, 0)
}

func TestGetListGnbIds(t *testing.T) {
	name := "name"
	plmnId := 0x02f829
	nbId := 0x4a952a0a
	listSize := 3
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	idsData := make([]string, listSize)
	idsEntities := make([]*entities.NbIdentity, listSize)
	for i := 0; i < listSize; i++ {
		nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: string(plmnId + i), NbId: string(nbId + i)}}
		data, err := proto.Marshal(nbIdentity)
		if err != nil {
			t.Errorf("#rNibReader_test.TestGetListGnbIds - Failed to marshal nodeb identity entity. Error: %v", err)
		}
		idsData[i] = string(data)
		idsEntities[i] = nbIdentity
	}
	var e error
	sdlInstanceMock.On("GetMembers", GnbType).Return(idsData, e)
	ids, er := w.GetListGnbIds()
	assert.Nil(t, er)
	assert.Len(t, *ids, listSize)
	for i, id := range *ids {
		assert.Equal(t, id.GetInventoryName(), name)
		assert.Equal(t, id.GetGlobalNbId().GetPlmnId(), idsEntities[i].GetGlobalNbId().GetPlmnId())
		assert.Equal(t, id.GetGlobalNbId().GetNbId(), idsEntities[i].GetGlobalNbId().GetNbId())
	}
}

func TestGetListEnbIdsSdlgoFailure(t *testing.T) {
	errMsg := "expected Sdlgo error"
	errMsgExpected := "2 INTERNAL_ERROR - expected Sdlgo error"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	e := errors.New(errMsg)
	var data []string
	sdlInstanceMock.On("GetMembers", EnbType).Return(data, e)
	ids, er := w.GetListEnbIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetCountGnbListOneId(t *testing.T) {
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	sdlInstanceMock.On("GroupSize", GnbType).Return(1, e)
	count, er := w.GetCountGnbList()
	assert.Nil(t, er)
	assert.Equal(t, count, 1)
}

func TestGetCountGnbList(t *testing.T) {
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	sdlInstanceMock.On("GroupSize", GnbType).Return(3, e)
	count, er := w.GetCountGnbList()
	assert.Nil(t, er)
	assert.Equal(t, count, 3)
}

func TestGetCountGnbListSdlgoFailure(t *testing.T) {
	errMsg := "expected Sdlgo error"
	errMsgExpected := "2 INTERNAL_ERROR - expected Sdlgo error"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	e := errors.New(errMsg)
	var count int
	sdlInstanceMock.On("GroupSize", GnbType).Return(count, e)
	count, er := w.GetCountGnbList()
	assert.NotNil(t, er)
	assert.Equal(t, 0, count)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetCell(t *testing.T) {
	name := "name"
	var pci uint32 = 10
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	cellEntity := entities.Cell{Type: entities.Cell_LTE_CELL, Cell: &entities.Cell_ServedCellInfo{ServedCellInfo: &entities.ServedCellInfo{Pci: pci}}}
	cellData, err := proto.Marshal(&cellEntity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetCell - Failed to marshal Cell entity. Error: %v", err)
	}
	var e error
	key, rNibErr := common.ValidateAndBuildCellNamePciKey(name, pci)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetCell - failed to validate key parameter")
	}
	ret := map[string]interface{}{key: string(cellData)}
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	cell, er := w.GetCell(name, pci)
	assert.Nil(t, er)
	assert.NotNil(t, cell)
	assert.Equal(t, cell.Type, entities.Cell_LTE_CELL)
	assert.NotNil(t, cell.GetServedCellInfo())
	assert.Equal(t, cell.GetServedCellInfo().GetPci(), pci)
}

func TestGetCellNotFoundFailure(t *testing.T) {
	name := "name"
	var pci uint32
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	var ret map[string]interface{}
	key, rNibErr := common.ValidateAndBuildCellNamePciKey(name, pci)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetCellNotFoundFailure - failed to validate key parameter")
	}
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	cell, er := w.GetCell(name, pci)
	assert.NotNil(t, er)
	assert.Nil(t, cell)
	assert.Equal(t, 1, er.GetCode())
	assert.EqualValues(t, "1 RESOURCE_NOT_FOUND - #rNibReader.getCellByKey - cell not found, key: PCI:name:00", er.Error())
}

func TestGetCellUnmarshalFailure(t *testing.T) {
	name := "name"
	var pci uint32
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	ret := make(map[string]interface{}, 1)
	key, rNibErr := common.ValidateAndBuildCellNamePciKey(name, pci)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetCellUnmarshalFailure - failed to validate key parameter")
	}
	ret[key] = "data"
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	cell, er := w.GetCell(name, pci)
	assert.NotNil(t, er)
	assert.Nil(t, cell)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, "2 INTERNAL_ERROR - proto: can't skip unknown wire type 4", er.Error())
}

func TestGetCellSdlgoFailure(t *testing.T) {
	name := "name"
	var pci uint32
	errMsg := "expected Sdlgo error"
	errMsgExpected := "2 INTERNAL_ERROR - expected Sdlgo error"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	e := errors.New(errMsg)
	var ret map[string]interface{}
	key, rNibErr := common.ValidateAndBuildCellNamePciKey(name, pci)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetCellSdlgoFailure - failed to validate key parameter")
	}
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	cell, er := w.GetCell(name, pci)
	assert.NotNil(t, er)
	assert.Nil(t, cell)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetNodebById(t *testing.T) {
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	nb := entities.NodebInfo{NodeType: entities.Node_ENB}
	nb.ConnectionStatus = 1
	nb.Ip = "localhost"
	nb.Port = 5656
	enb := entities.Enb{}
	cell := entities.ServedCellInfo{Tac: "tac"}
	enb.ServedCells = []*entities.ServedCellInfo{&cell}
	nb.Configuration = &entities.NodebInfo_Enb{Enb: &enb}
	var e error
	data, err := proto.Marshal(&nb)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetNodebById - Failed to marshal ENB instance. Error: %v", err)
	}

	plmnId := "02f829"
	nbId := "4a952a0a"
	key, rNibErr := common.ValidateAndBuildNodeBIdKey(entities.Node_ENB.String(), plmnId, nbId)
	if rNibErr != nil {
		t.Errorf("Failed to validate nodeb identity, plmnId: %s, nbId: %s", plmnId, nbId)
	}
	ret := map[string]interface{}{key: string(data)}
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	globalNbId := &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}
	getNb, er := w.GetNodebByGlobalNbId(entities.Node_ENB, globalNbId)
	assert.Nil(t, er)
	assert.Equal(t, getNb.Ip, nb.Ip)
	assert.Equal(t, getNb.Port, nb.Port)
	assert.Equal(t, getNb.ConnectionStatus, nb.ConnectionStatus)
	assert.Len(t, getNb.GetEnb().GetServedCells(), 1)
	assert.Equal(t, getNb.GetEnb().GetServedCells()[0].Tac, nb.GetEnb().GetServedCells()[0].Tac)
}

func TestGetNodebByIdNotFoundFailureEnb(t *testing.T) {
	plmnId := "02f829"
	nbId := "4a952a0a"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	key, rNibErr := common.ValidateAndBuildNodeBIdKey(entities.Node_ENB.String(), plmnId, nbId)
	if rNibErr != nil {
		t.Errorf("Failed to validate nodeb identity, plmnId: %s, nbId: %s", plmnId, nbId)
	}
	var ret map[string]interface{}
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	globalNbId := &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}
	getNb, er := w.GetNodebByGlobalNbId(entities.Node_ENB, globalNbId)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.Equal(t, 1, er.GetCode())
	assert.EqualValues(t, "1 RESOURCE_NOT_FOUND - #rNibReader.getNodeb - responding node not found. Key: ENB:02f829:4a952a0a", er.Error())
}

func TestGetNodebByIdNotFoundFailureGnb(t *testing.T) {
	plmnId := "02f829"
	nbId := "4a952a0a"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	key, rNibErr := common.ValidateAndBuildNodeBIdKey(entities.Node_GNB.String(), plmnId, nbId)
	if rNibErr != nil {
		t.Errorf("Failed to validate nodeb identity, plmnId: %s, nbId: %s", plmnId, nbId)
	}
	var ret map[string]interface{}
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	globalNbId := &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}
	getNb, er := w.GetNodebByGlobalNbId(entities.Node_GNB, globalNbId)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.Equal(t, 1, er.GetCode())
	assert.EqualValues(t, "1 RESOURCE_NOT_FOUND - #rNibReader.getNodeb - responding node not found. Key: GNB:02f829:4a952a0a", er.Error())
}

func TestGetNodeByIdUnmarshalFailure(t *testing.T) {
	plmnId := "02f829"
	nbId := "4a952a0a"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	key, rNibErr := common.ValidateAndBuildNodeBIdKey(entities.Node_ENB.String(), plmnId, nbId)
	if rNibErr != nil {
		t.Errorf("Failed to validate nodeb identity, plmnId: %s, nbId: %s", plmnId, nbId)
	}
	var e error
	ret := make(map[string]interface{}, 1)
	ret[key] = "data"
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	globalNbId := &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}
	getNb, er := w.GetNodebByGlobalNbId(entities.Node_ENB, globalNbId)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, "2 INTERNAL_ERROR - proto: can't skip unknown wire type 4", er.Error())
}

func TestGetNodeByIdSdlgoFailure(t *testing.T) {
	plmnId := "02f829"
	nbId := "4a952a0a"
	errMsg := "expected Sdlgo error"
	errMsgExpected := "2 INTERNAL_ERROR - expected Sdlgo error"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	key, rNibErr := common.ValidateAndBuildNodeBIdKey(entities.Node_GNB.String(), plmnId, nbId)
	if rNibErr != nil {
		t.Errorf("Failed to validate nodeb identity, plmnId: %s, nbId: %s", plmnId, nbId)
	}
	e := errors.New(errMsg)
	var ret map[string]interface{}
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	globalNbId := &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}
	getNb, er := w.GetNodebByGlobalNbId(entities.Node_GNB, globalNbId)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.Equal(t, 2, er.GetCode())
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetCellById(t *testing.T) {
	cellId := "aaaa"
	var pci uint32 = 10
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	cellEntity := entities.Cell{Type: entities.Cell_LTE_CELL, Cell: &entities.Cell_ServedCellInfo{ServedCellInfo: &entities.ServedCellInfo{Pci: pci}}}
	cellData, err := proto.Marshal(&cellEntity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetCellById - Failed to marshal Cell entity. Error: %v", err)
	}
	var e error
	key, rNibErr := common.ValidateAndBuildCellIdKey(cellId)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetCellById - failed to validate key parameter")
	}
	ret := map[string]interface{}{key: string(cellData)}
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	cell, er := w.GetCellById(entities.Cell_LTE_CELL, cellId)
	assert.Nil(t, er)
	assert.NotNil(t, cell)
	assert.Equal(t, cell.Type, entities.Cell_LTE_CELL)
	assert.NotNil(t, cell.GetServedCellInfo())
	assert.Equal(t, cell.GetServedCellInfo().GetPci(), pci)
}

func TestGetCellByIdNotFoundFailureEnb(t *testing.T) {
	cellId := "bbbb"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	var ret map[string]interface{}
	key, rNibErr := common.ValidateAndBuildCellIdKey(cellId)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetCellByIdNotFoundFailure - failed to validate key parameter")
	}
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	cell, er := w.GetCellById(entities.Cell_LTE_CELL, cellId)
	assert.NotNil(t, er)
	assert.Nil(t, cell)
	assert.Equal(t, 1, er.GetCode())
	assert.EqualValues(t, "1 RESOURCE_NOT_FOUND - #rNibReader.getCellByKey - cell not found, key: CELL:bbbb", er.Error())
}

func TestGetCellByIdNotFoundFailureGnb(t *testing.T) {
	cellId := "bbbb"
	readerPool = nil
	sdlInstanceMock := initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	var e error
	var ret map[string]interface{}
	key, rNibErr := common.ValidateAndBuildNrCellIdKey(cellId)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetCellByIdNotFoundFailure - failed to validate key parameter")
	}
	sdlInstanceMock.On("Get", []string{key}).Return(ret, e)
	cell, er := w.GetCellById(entities.Cell_NR_CELL, cellId)
	assert.NotNil(t, er)
	assert.Nil(t, cell)
	assert.Equal(t, 1, er.GetCode())
	assert.EqualValues(t, "1 RESOURCE_NOT_FOUND - #rNibReader.getCellByKey - cell not found, key: NRCELL:bbbb", er.Error())
}

func TestGetCellByIdTypeValidationFailure(t *testing.T) {
	cellId := "dddd"
	readerPool = nil
	initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	cell, er := w.GetCellById(5, cellId)
	assert.NotNil(t, er)
	assert.Nil(t, cell)
	assert.Equal(t, 3, er.GetCode())
	assert.EqualValues(t, "3 VALIDATION_ERROR - #rNibReader.GetCellById - invalid cell type: 5", er.Error())
}

func TestGetCellByIdValidationFailureGnb(t *testing.T) {
	cellId := ""
	readerPool = nil
	initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	cell, er := w.GetCellById(entities.Cell_NR_CELL, cellId)
	assert.NotNil(t, er)
	assert.Nil(t, cell)
	assert.Equal(t, 3, er.GetCode())
	assert.EqualValues(t, "3 VALIDATION_ERROR - #utils.ValidateAndBuildNrCellIdKey - an empty cell id received", er.Error())
}

func TestGetCellByIdValidationFailureEnb(t *testing.T) {
	cellId := ""
	readerPool = nil
	initSdlInstanceMock(namespace, 1)
	w := GetRNibReader()
	cell, er := w.GetCellById(entities.Cell_LTE_CELL, cellId)
	assert.NotNil(t, er)
	assert.Nil(t, cell)
	assert.Equal(t, 3, er.GetCode())
	assert.EqualValues(t, "3 VALIDATION_ERROR - #utils.ValidateAndBuildCellIdKey - an empty cell id received", er.Error())
}

//integration tests

//func TestGetEnbInteg(t *testing.T){
//	name := "nameEnb1"
//	Init("namespace", 1)
//	w := GetRNibReader()
//	nb, err := w.GetNodeb(name)
//	if err != nil{
//		fmt.Println(err)
//	} else {
//		fmt.Printf("#TestGetEnbInteg - responding node type: %v\n", nb)
//	}
//}
//
//func TestGetEnbCellsInteg(t *testing.T){
//	name := "nameEnb1"
//	Init("namespace", 1)
//	w := GetRNibReader()
//	cells, err := w.GetCellList(name)
//	if err != nil{
//		fmt.Println(err)
//	} else if cells != nil{
//		for _, cell := range cells.GetServedCellInfos().ServedCells{
//			fmt.Printf("responding node type Cell: %v\n", *cell)
//		}
//	}
//}
//
//func TestGetGnbInteg(t *testing.T){
//	name := "nameGnb1"
//	Init("namespace", 1)
//	w := GetRNibReader()
//	nb, err := w.GetNodeb(name)
//	if err != nil{
//		fmt.Println(err)
//	} else {
//		fmt.Printf("#TestGetGnbInteg - responding node type: %v\n", nb)
//	}
//}
//
//func TestGetGnbCellsInteg(t *testing.T){
//	name := "nameGnb1"
//	Init("namespace", 1)
//	w := GetRNibReader()
//	cells, err := w.GetCellList(name)
//	if err != nil{
//		fmt.Println(err)
//	} else if cells != nil{
//		for _, cell := range cells.GetServedNrCells().ServedCells{
//			fmt.Printf("responding node type NR Cell: %v\n", *cell)
//		}
//	}
//}
//
//func TestGetListEnbIdsInteg(t *testing.T) {
//	Init("namespace", 1)
//	w := GetRNibReader()
//	ids, err := w.GetListEnbIds()
//	if err != nil{
//		fmt.Println(err)
//	} else {
//		for _, id := range *ids{
//			fmt.Printf("#TestGetListEnbIdsInteg - ENB ID: %s\n", id)
//		}
//	}
//}
//
//func TestGetListGnbIdsInteg(t *testing.T) {
//	Init("namespace", 1)
//	w := GetRNibReader()
//	ids, err := w.GetListGnbIds()
//	if err != nil{
//		fmt.Println(err)
//	} else {
//		for _, id := range *ids{
//			fmt.Printf("#TestGetListGnbIdsInteg - GNB ID: %s\n", id)
//		}
//	}
//}
//
//func TestGetCountGnbListInteg(t *testing.T) {
//	Init("namespace", 1)
//	w := GetRNibReader()
//	count, err := w.GetCountGnbList()
//	if err != nil{
//		fmt.Println(err)
//	} else {
//		fmt.Printf("#TestGetCountGnbListInteg - count Gnb list: %d\n", count)
//	}
//}
//
//func TestGetGnbCellInteg(t *testing.T){
//	name := "nameGnb7"
//	pci := 0x0a
//	Init("namespace", 1)
//	w := GetRNibReader()
//	cell, err := w.GetCell(name, uint32(pci))
//	if err != nil{
//		fmt.Println(err)
//	} else if cell != nil{
//		fmt.Printf("responding node type NR Cell: %v\n", cell.GetServedNrCell())
//	}
//}
//
//func TestGetEnbCellInteg(t *testing.T) {
//	name := "nameEnb1"
//	pci := 0x22
//	Init("namespace", 1)
//	w := GetRNibReader()
//	cell, err := w.GetCell(name, uint32(pci))
//	if err != nil {
//		fmt.Println(err)
//	} else if cell != nil {
//		fmt.Printf("responding node type LTE Cell: %v\n", cell.GetServedCellInfo())
//	}
//}
//
//func TestGetEnbCellByIdInteg(t *testing.T){
//	Init("namespace", 1)
//	w := GetRNibReader()
//	cell, err := w.GetCellById(entities.Cell_NR_CELL, "45d")
//	if err != nil{
//		fmt.Println(err)
//	} else if cell != nil{
//		fmt.Printf("responding node type NR Cell: %v\n", cell.GetServedNrCell())
//	}
//}
