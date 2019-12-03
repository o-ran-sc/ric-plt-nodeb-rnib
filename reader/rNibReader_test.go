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
	"encoding/json"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/common"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var namespace = "namespace"

func initSdlInstanceMock() (w RNibReader, sdlInstanceMock *MockSdlInstance) {
	sdlInstanceMock = new(MockSdlInstance)
	w = GetRNibReader(sdlInstanceMock)
	return
}

func TestGetNodeB(t *testing.T) {
	name := "name"
	w, sdlInstanceMock := initSdlInstanceMock()
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
	redisKey, rNibErr := common.ValidateAndBuildNodeBNameKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetNodeB - failed to validate key parameter")
	}
	ret := map[string]interface{}{redisKey: string(data)}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
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
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	var ret map[string]interface{}
	redisKey, rNibErr := common.ValidateAndBuildNodeBNameKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetNodeBNotFoundFailure - failed to validate key parameter")
	}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	getNb, er := w.GetNodeb(name)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.IsType(t, &common.ResourceNotFoundError{}, er)
	assert.EqualValues(t, "#rNibReader.getByKeyAndUnmarshal - entity of type *entities.NodebInfo not found. Key: RAN:name", er.Error())
}

func TestGetNodeBUnmarshalFailure(t *testing.T) {
	name := "name"
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	ret := make(map[string]interface{}, 1)
	redisKey, rNibErr := common.ValidateAndBuildNodeBNameKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetNodeBUnmarshalFailure - failed to validate key parameter")
	}
	ret[redisKey] = "data"
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	getNb, er := w.GetNodeb(name)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, "proto: can't skip unknown wire type 4", er.Error())
}

func TestGetNodeBSdlgoFailure(t *testing.T) {
	name := "name"
	errMsg := "expected Sdlgo error"
	errMsgExpected := "expected Sdlgo error"
	w, sdlInstanceMock := initSdlInstanceMock()
	e := errors.New(errMsg)
	var ret map[string]interface{}
	redisKey, rNibErr := common.ValidateAndBuildNodeBNameKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetNodeBSdlgoFailure - failed to validate key parameter")
	}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	getNb, er := w.GetNodeb(name)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetNodeBCellsListEnb(t *testing.T) {
	name := "name"
	w, sdlInstanceMock := initSdlInstanceMock()
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
	redisKey, rNibErr := common.ValidateAndBuildNodeBNameKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetNodeBCellsListEnb - failed to validate key parameter")
	}
	ret := map[string]interface{}{redisKey: string(data)}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	cells, er := w.GetCellList(name)
	assert.Nil(t, er)
	assert.NotNil(t, cells)
	assert.Len(t, cells.GetServedCellInfos().GetServedCells(), 1)
	retCell := cells.GetServedCellInfos().GetServedCells()[0]
	assert.Equal(t, retCell.Tac, "tac")
}

func TestGetNodeBCellsListGnb(t *testing.T) {
	name := "name"
	w, sdlInstanceMock := initSdlInstanceMock()
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
	redisKey, rNibErr := common.ValidateAndBuildNodeBNameKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetNodeBCellsListGnb - failed to validate key parameter")
	}
	ret := map[string]interface{}{redisKey: string(data)}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	cells, er := w.GetCellList(name)
	assert.Nil(t, er)
	assert.NotNil(t, cells)
	assert.Len(t, cells.GetServedNrCells().GetServedCells(), 1)
	retCell := cells.GetServedNrCells().GetServedCells()[0]
	assert.Equal(t, retCell.GetServedNrCellInformation().GetNrPci(), uint32(10))
}

func TestGetNodeBCellsListNodeUnmarshalFailure(t *testing.T) {
	name := "name"
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	ret := make(map[string]interface{}, 1)
	redisKey, rNibErr := common.ValidateAndBuildNodeBNameKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetNodeBCellsListNodeUnmarshalFailure - failed to validate key parameter")
	}
	ret[redisKey] = "data"
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	cells, er := w.GetCellList(name)
	assert.NotNil(t, er)
	assert.Nil(t, cells)
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, "proto: can't skip unknown wire type 4", er.Error())
}

func TestGetNodeBCellsListNodeNotFoundFailure(t *testing.T) {
	name := "name"
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	var ret map[string]interface{}
	redisKey, rNibErr := common.ValidateAndBuildNodeBNameKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetNodeBCellsListNodeNotFoundFailure - failed to validate key parameter")
	}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	cells, er := w.GetCellList(name)
	assert.NotNil(t, er)
	assert.Nil(t, cells)
	assert.IsType(t, &common.ResourceNotFoundError{}, er)
	assert.EqualValues(t, "#rNibReader.getByKeyAndUnmarshal - entity of type *entities.NodebInfo not found. Key: RAN:name", er.Error())
}

func TestGetNodeBCellsListNotFoundFailureEnb(t *testing.T) {
	name := "name"
	w, sdlInstanceMock := initSdlInstanceMock()
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
	redisKey, rNibErr := common.ValidateAndBuildNodeBNameKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetNodeBCellsListNotFoundFailureEnb - failed to validate key parameter")
	}
	ret := map[string]interface{}{redisKey: string(data)}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	_, er := w.GetCellList(name)
	assert.NotNil(t, er)
	assert.EqualValues(t, "#rNibReader.GetCellList - served cells not found. Responding node RAN name: name.", er.Error())
}

func TestGetNodeBCellsListNotFoundFailureGnb(t *testing.T) {
	name := "name"
	w, sdlInstanceMock := initSdlInstanceMock()
	nb := entities.NodebInfo{}
	nb.ConnectionStatus = 1
	nb.Ip = "localhost"
	nb.Port = 5656
	gnb := entities.Gnb{}
	nb.Configuration = &entities.NodebInfo_Gnb{Gnb: &gnb}
	var e error
	data, err := proto.Marshal(&nb)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetNodeBCellsListNotFoundFailureGnb - Failed to marshal ENB instance. Error: %v", err)
	}
	redisKey, rNibErr := common.ValidateAndBuildNodeBNameKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetNodeBCellsListNotFoundFailureGnb - failed to validate key parameter")
	}
	ret := map[string]interface{}{redisKey: string(data)}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	_, er := w.GetCellList(name)
	assert.NotNil(t, er)
	assert.EqualValues(t, "#rNibReader.GetCellList - served cells not found. Responding node RAN name: name.", er.Error())
}

func TestGetListGnbIdsUnmarshalFailure(t *testing.T) {
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	sdlInstanceMock.On("GetMembers", entities.Node_GNB.String()).Return([]string{"data"}, e)
	ids, er := w.GetListGnbIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.IsType(t, &common.InternalError{}, er)
	assert.Equal(t, "proto: can't skip unknown wire type 4", er.Error())
}

func TestGetListGnbIdsSdlgoFailure(t *testing.T) {
	errMsg := "expected Sdlgo error"
	errMsgExpected := "expected Sdlgo error"
	w, sdlInstanceMock := initSdlInstanceMock()
	e := errors.New(errMsg)
	var data []string
	sdlInstanceMock.On("GetMembers", entities.Node_GNB.String()).Return(data, e)
	ids, er := w.GetListGnbIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetListNodesIdsGnbSdlgoFailure(t *testing.T) {
	w, sdlInstanceMock := initSdlInstanceMock()

	name := "name"
	plmnId := "02f829"
	nbId := "4a952a0a"
	nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}}
	var nilError error
	data, err := proto.Marshal(nbIdentity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListNodesIdsGnbSdlgoFailure - Failed to marshal nodeb identity entity. Error: %v", err)
	}
	sdlInstanceMock.On("GetMembers", entities.Node_ENB.String()).Return([]string{string(data)}, nilError)

	errMsg := "expected Sdlgo error"
	errMsgExpected := "expected Sdlgo error"
	expectedError := errors.New(errMsg)
	var nilData []string
	sdlInstanceMock.On("GetMembers", entities.Node_GNB.String()).Return(nilData, expectedError)

	ids, er := w.GetListNodebIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetListNodesIdsEnbSdlgoFailure(t *testing.T) {
	w, sdlInstanceMock := initSdlInstanceMock()

	name := "name"
	plmnId := "02f829"
	nbId := "4a952a0a"
	nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}}
	var nilError error
	data, err := proto.Marshal(nbIdentity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListNodesIdsEnbSdlgoFailure - Failed to marshal nodeb identity entity. Error: %v", err)
	}
	sdlInstanceMock.On("GetMembers", entities.Node_GNB.String()).Return([]string{string(data)}, nilError)

	errMsg := "expected Sdlgo error"
	errMsgExpected := "expected Sdlgo error"
	expectedError := errors.New(errMsg)
	var nilData []string
	sdlInstanceMock.On("GetMembers", entities.Node_ENB.String()).Return(nilData, expectedError)

	ids, er := w.GetListNodebIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetListNodesIdsSuccess(t *testing.T) {
	w, sdlInstanceMock := initSdlInstanceMock()
	var nilError error

	name := "name"
	plmnId := "02f829"
	nbId := "4a952a0a"
	nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}}
	data, err := proto.Marshal(nbIdentity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListNodesIdsSuccess - Failed to marshal nodeb identity entity. Error: %v", err)
	}

	name1 := "name1"
	plmnId1 := "02f845"
	nbId1 := "4a952a75"
	nbIdentity1 := &entities.NbIdentity{InventoryName: name1, GlobalNbId: &entities.GlobalNbId{PlmnId: plmnId1, NbId: nbId1}}
	data1, err := proto.Marshal(nbIdentity1)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListNodesIdsSuccess - Failed to marshal nodeb identity entity. Error: %v", err)
	}

	name2 := "name2"
	nbIdentity2 := &entities.NbIdentity{InventoryName: name2}
	data2, err := proto.Marshal(nbIdentity2)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListNodesIdsSuccess - Failed to marshal nodeb identity entity. Error: %v", err)
	}

	sdlInstanceMock.On("GetMembers", entities.Node_GNB.String()).Return([]string{string(data)}, nilError)
	sdlInstanceMock.On("GetMembers", entities.Node_ENB.String()).Return([]string{string(data1)}, nilError)
	sdlInstanceMock.On("GetMembers", entities.Node_UNKNOWN.String()).Return([]string{string(data2)}, nilError)

	ids, er := w.GetListNodebIds()
	assert.Nil(t, er)
	assert.NotNil(t, ids)
	assert.Len(t, ids, 3)
}

func TestGetListEnbIdsUnmarshalFailure(t *testing.T) {
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	sdlInstanceMock.On("GetMembers", entities.Node_ENB.String()).Return([]string{"data"}, e)
	ids, er := w.GetListEnbIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.IsType(t, &common.InternalError{}, er)
	assert.Equal(t, "proto: can't skip unknown wire type 4", er.Error())
}

func TestGetListEnbIdsOneId(t *testing.T) {
	name := "name"
	plmnId := "02f829"
	nbId := "4a952a0a"
	w, sdlInstanceMock := initSdlInstanceMock()
	nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}}
	var e error
	data, err := proto.Marshal(nbIdentity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListEnbIds - Failed to marshal nodeb identity entity. Error: %v", err)
	}
	sdlInstanceMock.On("GetMembers", entities.Node_ENB.String()).Return([]string{string(data)}, e)
	ids, er := w.GetListEnbIds()
	assert.Nil(t, er)
	assert.Len(t, ids, 1)
	assert.Equal(t, (ids)[0].GetInventoryName(), name)
	assert.Equal(t, (ids)[0].GetGlobalNbId().GetPlmnId(), nbIdentity.GetGlobalNbId().GetPlmnId())
	assert.Equal(t, (ids)[0].GetGlobalNbId().GetNbId(), nbIdentity.GetGlobalNbId().GetNbId())
}

func TestGetListEnbIdsNoIds(t *testing.T) {
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	sdlInstanceMock.On("GetMembers", entities.Node_ENB.String()).Return([]string{}, e)
	ids, er := w.GetListEnbIds()
	assert.Nil(t, er)
	assert.Len(t, ids, 0)
}

func TestGetListEnbIds(t *testing.T) {
	name := "name"
	plmnId := 0x02f829
	nbId := 0x4a952a0a
	listSize := 3
	w, sdlInstanceMock := initSdlInstanceMock()
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
	sdlInstanceMock.On("GetMembers", entities.Node_ENB.String()).Return(idsData, e)
	ids, er := w.GetListEnbIds()
	assert.Nil(t, er)
	assert.Len(t, ids, listSize)
	for i, id := range ids {
		assert.Equal(t, id.GetInventoryName(), name)
		assert.Equal(t, id.GetGlobalNbId().GetPlmnId(), idsEntities[i].GetGlobalNbId().GetPlmnId())
		assert.Equal(t, id.GetGlobalNbId().GetNbId(), idsEntities[i].GetGlobalNbId().GetNbId())
	}
}

func TestGetListGnbIdsOneId(t *testing.T) {
	name := "name"
	plmnId := "02f829"
	nbId := "4a952a0a"
	w, sdlInstanceMock := initSdlInstanceMock()
	nbIdentity := &entities.NbIdentity{InventoryName: name, GlobalNbId: &entities.GlobalNbId{PlmnId: plmnId, NbId: nbId}}
	var e error
	data, err := proto.Marshal(nbIdentity)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetListGnbIds - Failed to marshal nodeb identity entity. Error: %v", err)
	}
	sdlInstanceMock.On("GetMembers", entities.Node_GNB.String()).Return([]string{string(data)}, e)
	ids, er := w.GetListGnbIds()
	assert.Nil(t, er)
	assert.Len(t, ids, 1)
	assert.Equal(t, (ids)[0].GetInventoryName(), name)
	assert.Equal(t, (ids)[0].GetGlobalNbId().GetPlmnId(), nbIdentity.GetGlobalNbId().GetPlmnId())
	assert.Equal(t, (ids)[0].GetGlobalNbId().GetNbId(), nbIdentity.GetGlobalNbId().GetNbId())
}

func TestGetListGnbIdsNoIds(t *testing.T) {
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	sdlInstanceMock.On("GetMembers", entities.Node_GNB.String()).Return([]string{}, e)
	ids, er := w.GetListGnbIds()
	assert.Nil(t, er)
	assert.Len(t, ids, 0)
}

func TestGetListGnbIds(t *testing.T) {
	name := "name"
	plmnId := 0x02f829
	nbId := 0x4a952a0a
	listSize := 3
	w, sdlInstanceMock := initSdlInstanceMock()
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
	sdlInstanceMock.On("GetMembers", entities.Node_GNB.String()).Return(idsData, e)
	ids, er := w.GetListGnbIds()
	assert.Nil(t, er)
	assert.Len(t, ids, listSize)
	for i, id := range ids {
		assert.Equal(t, id.GetInventoryName(), name)
		assert.Equal(t, id.GetGlobalNbId().GetPlmnId(), idsEntities[i].GetGlobalNbId().GetPlmnId())
		assert.Equal(t, id.GetGlobalNbId().GetNbId(), idsEntities[i].GetGlobalNbId().GetNbId())
	}
}

func TestGetListEnbIdsSdlgoFailure(t *testing.T) {
	errMsg := "expected Sdlgo error"
	errMsgExpected := "expected Sdlgo error"
	w, sdlInstanceMock := initSdlInstanceMock()
	e := errors.New(errMsg)
	var data []string
	sdlInstanceMock.On("GetMembers", entities.Node_ENB.String()).Return(data, e)
	ids, er := w.GetListEnbIds()
	assert.NotNil(t, er)
	assert.Nil(t, ids)
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetCountGnbListOneId(t *testing.T) {
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	sdlInstanceMock.On("GroupSize", entities.Node_GNB.String()).Return(1, e)
	count, er := w.GetCountGnbList()
	assert.Nil(t, er)
	assert.Equal(t, count, 1)
}

func TestGetCountGnbList(t *testing.T) {
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	sdlInstanceMock.On("GroupSize", entities.Node_GNB.String()).Return(3, e)
	count, er := w.GetCountGnbList()
	assert.Nil(t, er)
	assert.Equal(t, count, 3)
}

func TestGetCountGnbListSdlgoFailure(t *testing.T) {
	errMsg := "expected Sdlgo error"
	errMsgExpected := "expected Sdlgo error"
	w, sdlInstanceMock := initSdlInstanceMock()
	e := errors.New(errMsg)
	var count int
	sdlInstanceMock.On("GroupSize", entities.Node_GNB.String()).Return(count, e)
	count, er := w.GetCountGnbList()
	assert.NotNil(t, er)
	assert.Equal(t, 0, count)
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetCell(t *testing.T) {
	name := "name"
	var pci uint32 = 10
	w, sdlInstanceMock := initSdlInstanceMock()
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
	w, sdlInstanceMock := initSdlInstanceMock()
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
	assert.IsType(t, &common.ResourceNotFoundError{}, er)
	assert.EqualValues(t, "#rNibReader.getByKeyAndUnmarshal - entity of type *entities.Cell not found. Key: PCI:name:00", er.Error())
}

func TestGetCellUnmarshalFailure(t *testing.T) {
	name := "name"
	var pci uint32
	w, sdlInstanceMock := initSdlInstanceMock()
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
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, "proto: can't skip unknown wire type 4", er.Error())
}

func TestGetCellSdlgoFailure(t *testing.T) {
	name := "name"
	var pci uint32
	errMsg := "expected Sdlgo error"
	errMsgExpected := "expected Sdlgo error"
	w, sdlInstanceMock := initSdlInstanceMock()
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
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetNodebById(t *testing.T) {
	w, sdlInstanceMock := initSdlInstanceMock()
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
	w, sdlInstanceMock := initSdlInstanceMock()
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
	assert.IsType(t, &common.ResourceNotFoundError{}, er)
	assert.EqualValues(t, "#rNibReader.getByKeyAndUnmarshal - entity of type *entities.NodebInfo not found. Key: ENB:02f829:4a952a0a", er.Error())
}

func TestGetNodebByIdNotFoundFailureGnb(t *testing.T) {
	plmnId := "02f829"
	nbId := "4a952a0a"
	w, sdlInstanceMock := initSdlInstanceMock()
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
	assert.IsType(t, &common.ResourceNotFoundError{}, er)
	assert.EqualValues(t, "#rNibReader.getByKeyAndUnmarshal - entity of type *entities.NodebInfo not found. Key: GNB:02f829:4a952a0a", er.Error())
}

func TestGetNodeByIdUnmarshalFailure(t *testing.T) {
	plmnId := "02f829"
	nbId := "4a952a0a"
	w, sdlInstanceMock := initSdlInstanceMock()
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
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, "proto: can't skip unknown wire type 4", er.Error())
}

func TestGetNodeByIdSdlgoFailure(t *testing.T) {
	plmnId := "02f829"
	nbId := "4a952a0a"
	errMsg := "expected Sdlgo error"
	errMsgExpected := "expected Sdlgo error"
	w, sdlInstanceMock := initSdlInstanceMock()
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
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func TestGetCellById(t *testing.T) {
	cellId := "aaaa"
	var pci uint32 = 10
	w, sdlInstanceMock := initSdlInstanceMock()
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
	w, sdlInstanceMock := initSdlInstanceMock()
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
	assert.IsType(t, &common.ResourceNotFoundError{}, er)
	assert.EqualValues(t, "#rNibReader.getByKeyAndUnmarshal - entity of type *entities.Cell not found. Key: CELL:bbbb", er.Error())
}

func TestGetCellByIdNotFoundFailureGnb(t *testing.T) {
	cellId := "bbbb"
	w, sdlInstanceMock := initSdlInstanceMock()
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
	assert.IsType(t, &common.ResourceNotFoundError{}, er)
	assert.EqualValues(t, "#rNibReader.getByKeyAndUnmarshal - entity of type *entities.Cell not found. Key: NRCELL:bbbb", er.Error())
}

func TestGetCellByIdTypeValidationFailure(t *testing.T) {
	cellId := "dddd"
	w, _ := initSdlInstanceMock()
	cell, er := w.GetCellById(5, cellId)
	assert.NotNil(t, er)
	assert.Nil(t, cell)
	assert.IsType(t, &common.ValidationError{}, er)
	assert.EqualValues(t, "#rNibReader.GetCellById - invalid cell type: 5", er.Error())
}

func TestGetCellByIdValidationFailureGnb(t *testing.T) {
	cellId := ""
	w, _ := initSdlInstanceMock()
	cell, er := w.GetCellById(entities.Cell_NR_CELL, cellId)
	assert.NotNil(t, er)
	assert.Nil(t, cell)
	assert.IsType(t, &common.ValidationError{}, er)
	assert.EqualValues(t, "#utils.ValidateAndBuildNrCellIdKey - an empty cell id received", er.Error())
}

func TestGetCellByIdValidationFailureEnb(t *testing.T) {
	cellId := ""
	w, _ := initSdlInstanceMock()
	cell, er := w.GetCellById(entities.Cell_LTE_CELL, cellId)
	assert.NotNil(t, er)
	assert.Nil(t, cell)
	assert.IsType(t, &common.ValidationError{}, er)
	assert.EqualValues(t, "#utils.ValidateAndBuildCellIdKey - an empty cell id received", er.Error())
}

func TestGetRanLoadInformation(t *testing.T) {
	name := "name"
	w, sdlInstanceMock := initSdlInstanceMock()
	loadInfo := generateRanLoadInformation()
	var e error
	data, err := proto.Marshal(loadInfo)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetRanLoadInformation - Failed to marshal RanLoadInformation entity. Error: %v", err)
	}
	redisKey, rNibErr := common.ValidateAndBuildRanLoadInformationKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetRanLoadInformationNotFoundFailure - failed to validate key parameter")
	}
	ret := map[string]interface{}{redisKey: string(data)}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	getLoadInfo, er := w.GetRanLoadInformation(name)
	assert.Nil(t, er)
	assert.NotNil(t, getLoadInfo)
	expected, err := json.Marshal(loadInfo)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetRanLoadInformation - Failed to marshal RanLoadInformation entity. Error: %v", err)
	}
	actual, err := json.Marshal(getLoadInfo)
	if err != nil {
		t.Errorf("#rNibReader_test.TestGetRanLoadInformation - Failed to marshal RanLoadInformation entity. Error: %v", err)
	}
	assert.EqualValues(t, expected, actual)
}

func TestGetRanLoadInformationValidationFailure(t *testing.T) {
	name := ""
	w, _ := initSdlInstanceMock()
	getNb, er := w.GetRanLoadInformation(name)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.IsType(t, &common.ValidationError{}, er)
	assert.EqualValues(t, "#utils.ValidateAndBuildRanLoadInformationKey - an empty inventory name received", er.Error())
}

func TestGetRanLoadInformationNotFoundFailure(t *testing.T) {
	name := "name"
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	var ret map[string]interface{}
	redisKey, rNibErr := common.ValidateAndBuildRanLoadInformationKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetRanLoadInformationNotFoundFailure - failed to validate key parameter")
	}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	getNb, er := w.GetRanLoadInformation(name)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.IsType(t, &common.ResourceNotFoundError{}, er)
	assert.EqualValues(t, "#rNibReader.getByKeyAndUnmarshal - entity of type *entities.RanLoadInformation not found. Key: LOAD:name", er.Error())
}

func TestGetRanLoadInformationUnmarshalFailure(t *testing.T) {
	name := "name"
	w, sdlInstanceMock := initSdlInstanceMock()
	var e error
	ret := make(map[string]interface{}, 1)
	redisKey, rNibErr := common.ValidateAndBuildRanLoadInformationKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetRanLoadInformationUnmarshalFailure - failed to validate key parameter")
	}
	ret[redisKey] = "data"
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	getNb, er := w.GetRanLoadInformation(name)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, "proto: can't skip unknown wire type 4", er.Error())
}

func TestGetRanLoadInformationSdlgoFailure(t *testing.T) {
	name := "name"
	errMsg := "expected Sdlgo error"
	errMsgExpected := "expected Sdlgo error"
	w, sdlInstanceMock := initSdlInstanceMock()
	e := errors.New(errMsg)
	var ret map[string]interface{}
	redisKey, rNibErr := common.ValidateAndBuildRanLoadInformationKey(name)
	if rNibErr != nil {
		t.Errorf("#rNibReader_test.TestGetRanLoadInformationSdlgoFailure - failed to validate key parameter")
	}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)
	getNb, er := w.GetRanLoadInformation(name)
	assert.NotNil(t, er)
	assert.Nil(t, getNb)
	assert.IsType(t, &common.InternalError{}, er)
	assert.EqualValues(t, errMsgExpected, er.Error())
}

func generateCellLoadInformation() *entities.CellLoadInformation {
	cellLoadInformation := entities.CellLoadInformation{}

	cellLoadInformation.CellId = "123"

	ulInterferenceOverloadIndication := entities.UlInterferenceOverloadIndication_HIGH_INTERFERENCE
	cellLoadInformation.UlInterferenceOverloadIndications = []entities.UlInterferenceOverloadIndication{ulInterferenceOverloadIndication}

	ulHighInterferenceInformation := entities.UlHighInterferenceInformation{
		TargetCellId:                 "456",
		UlHighInterferenceIndication: "xxx",
	}

	cellLoadInformation.UlHighInterferenceInfos = []*entities.UlHighInterferenceInformation{&ulHighInterferenceInformation}

	cellLoadInformation.RelativeNarrowbandTxPower = &entities.RelativeNarrowbandTxPower{
		RntpPerPrb:                       "xxx",
		RntpThreshold:                    entities.RntpThreshold_NEG_4,
		NumberOfCellSpecificAntennaPorts: entities.NumberOfCellSpecificAntennaPorts_V1_ANT_PRT,
		PB:                               1,
		PdcchInterferenceImpact:          2,
		EnhancedRntp: &entities.EnhancedRntp{
			EnhancedRntpBitmap:     "xxx",
			RntpHighPowerThreshold: entities.RntpThreshold_NEG_2,
			EnhancedRntpStartTime:  &entities.StartTime{StartSfn: 500, StartSubframeNumber: 5},
		},
	}

	cellLoadInformation.AbsInformation = &entities.AbsInformation{
		Mode:                             entities.AbsInformationMode_ABS_INFO_FDD,
		AbsPatternInfo:                   "xxx",
		NumberOfCellSpecificAntennaPorts: entities.NumberOfCellSpecificAntennaPorts_V2_ANT_PRT,
		MeasurementSubset:                "xxx",
	}

	cellLoadInformation.InvokeIndication = entities.InvokeIndication_ABS_INFORMATION

	cellLoadInformation.ExtendedUlInterferenceOverloadInfo = &entities.ExtendedUlInterferenceOverloadInfo{
		AssociatedSubframes:                       "xxx",
		ExtendedUlInterferenceOverloadIndications: cellLoadInformation.UlInterferenceOverloadIndications,
	}

	compInformationItem := &entities.CompInformationItem{
		CompHypothesisSets: []*entities.CompHypothesisSet{{CellId: "789", CompHypothesis: "xxx"}},
		BenefitMetric:      50,
	}

	cellLoadInformation.CompInformation = &entities.CompInformation{
		CompInformationItems:     []*entities.CompInformationItem{compInformationItem},
		CompInformationStartTime: &entities.StartTime{StartSfn: 123, StartSubframeNumber: 456},
	}

	cellLoadInformation.DynamicDlTransmissionInformation = &entities.DynamicDlTransmissionInformation{
		State:             entities.NaicsState_NAICS_ACTIVE,
		TransmissionModes: "xxx",
		PB:                2,
		PAList:            []entities.PA{entities.PA_DB_NEG_3},
	}

	return &cellLoadInformation
}

func generateRanLoadInformation() *entities.RanLoadInformation {
	ranLoadInformation := entities.RanLoadInformation{}

	ranLoadInformation.LoadTimestamp = uint64(time.Now().UnixNano())

	cellLoadInformation := generateCellLoadInformation()
	ranLoadInformation.CellLoadInfos = []*entities.CellLoadInformation{cellLoadInformation}

	return &ranLoadInformation
}

func TestGetE2TInstanceSuccess(t *testing.T) {
	address := "10.10.2.15:9800"
	redisKey, validationErr := common.ValidateAndBuildE2TInstanceKey(address)

	if validationErr != nil {
		t.Errorf("#rNibReader_test.TestGetE2TInstanceSuccess - Failed to build E2T Instance key. Error: %v", validationErr)
	}

	w, sdlInstanceMock := initSdlInstanceMock()

	e2tInstance := generateE2tInstance(address)
	data, err := json.Marshal(e2tInstance)

	if err != nil {
		t.Errorf("#rNibReader_test.TestGetE2TInstanceSuccess - Failed to marshal E2tInstance entity. Error: %v", err)
	}

	var e error
	ret := map[string]interface{}{redisKey: string(data)}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, e)

	res, rNibErr := w.GetE2TInstance(address)
	assert.Nil(t, rNibErr)
	assert.Equal(t, e2tInstance, res)
}

func TestUnmarshal(t *testing.T) {
	e2tInstance := generateE2tInstance("10.0.2.15:5555")
	marshaled, _ := json.Marshal(e2tInstance)
	m := map[string]interface{}{
		"whatever": string(marshaled),
	}
	var entity *entities.E2TInstance
	_ = json.Unmarshal([]byte(m["whatever"].(string)), entity)
}

func TestGetE2TInstanceEmptyAddressFailure(t *testing.T) {
	w, _ := initSdlInstanceMock()
	res, err := w.GetE2TInstance("")
	assert.NotNil(t, err)
	assert.IsType(t, &common.ValidationError{}, err)
	assert.Nil(t, res)
}

func TestGetE2TInstanceSdlError(t *testing.T) {
	address := "10.10.2.15:9800"
	redisKey, validationErr := common.ValidateAndBuildE2TInstanceKey(address)

	if validationErr != nil {
		t.Errorf("#rNibReader_test.TestGetE2TInstanceSuccess - Failed to build E2T Instance key. Error: %v", validationErr)
	}

	w, sdlInstanceMock := initSdlInstanceMock()

	expectedErr := errors.New("expected error")
	var ret map[string]interface{}
	sdlInstanceMock.On("Get", []string{redisKey}).Return(ret, expectedErr)

	res, rNibErr := w.GetE2TInstance(address)
	assert.NotNil(t, rNibErr)
	assert.Nil(t, res)
}

func generateE2tInstance(address string) *entities.E2TInstance {
	e2tInstance := entities.NewE2TInstance(address)
	e2tInstance.AssociatedRanList = []string{"test1", "test2"}
	return e2tInstance
}

func TestGetE2TAddressesSuccess(t *testing.T) {
	address := "10.10.2.15:9800"
	w, sdlInstanceMock := initSdlInstanceMock()

	e2tAddresses := []string{address}
	data, err := json.Marshal(e2tAddresses)

	if err != nil {
		t.Errorf("#rNibReader_test.TestGetE2TInfoListSuccess - Failed to marshal E2TInfoList. Error: %v", err)
	}

	var e error
	ret := map[string]interface{}{E2TAddressesKey: string(data)}
	sdlInstanceMock.On("Get", []string{E2TAddressesKey}).Return(ret, e)

	res, rNibErr := w.GetE2TAddresses()
	assert.Nil(t, rNibErr)
	assert.Equal(t, e2tAddresses, res)
}

func TestGetE2TInstancesSuccess(t *testing.T) {
	address := "10.10.2.15:9800"
	address2 := "10.10.2.16:9800"
	redisKey, _ := common.ValidateAndBuildE2TInstanceKey(address)
	redisKey2, _ := common.ValidateAndBuildE2TInstanceKey(address2)

	w, sdlInstanceMock := initSdlInstanceMock()

	e2tInstance1 := generateE2tInstance(address)
	e2tInstance2 := generateE2tInstance(address2)

	data1, _ := json.Marshal(e2tInstance1)
	data2, _ := json.Marshal(e2tInstance2)

	var e error
	ret := map[string]interface{}{redisKey: string(data1), redisKey2: string(data2)}
	sdlInstanceMock.On("Get", []string{redisKey, redisKey2}).Return(ret, e)

	res, err := w.GetE2TInstances([]string{address, address2})
	assert.Nil(t, err)
	assert.Equal(t, []*entities.E2TInstance{e2tInstance1, e2tInstance2}, res)
}

func TestGetE2TInfoListSdlError(t *testing.T) {
	w, sdlInstanceMock := initSdlInstanceMock()

	expectedErr := errors.New("expected error")
	var ret map[string]interface{}
	sdlInstanceMock.On("Get", []string{E2TAddressesKey}).Return(ret, expectedErr)

	res, rNibErr := w.GetE2TAddresses()
	assert.NotNil(t, rNibErr)
	assert.Nil(t, res)
}

//integration tests
//
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
//		for _, id := range ids{
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
//		for _, id := range ids{
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
//
//func TestGetListNbIdsInteg(t *testing.T) {
//	Init("e2Manager", 1)
//	w := GetRNibReader()
//	ids, err := w.GetListNodebIds()
//	if err != nil{
//		fmt.Println(err)
//	} else {
//		for _, id := range ids{
//			fmt.Printf("#TestGetListGnbIdsInteg - NB ID: %s\n", id)
//		}
//	}
//}
//
//func TestGetRanLoadInformationInteg(t *testing.T){
//	Init("e2Manager", 1)
//	w := GetRNibReader()
//	ranLoadInformation, err := w.GetRanLoadInformation("ran_integ")
//	if err != nil{
//		t.Errorf("#rNibReader_test.TestGetRanLoadInformationInteg - Failed to get RanLoadInformation entity. Error: %v", err)
//	}
//	assert.NotNil(t, ranLoadInformation)
//	fmt.Printf("#rNibReader_test.TestGetRanLoadInformationInteg - GNB ID: %s\n", ranLoadInformation)
//}
