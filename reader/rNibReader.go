//
// Copyright 2019 AT&T Intellectual Property
// Copyright 2019 Nokia
// Copyright 2023 Capgemini
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
	"reflect"
)

const E2TAddressesKey = "E2TAddresses"

type rNibReaderInstance struct {
	sdl        common.ISdlInstance //Deprecated: Will be removed in a future release and replaced by sdlStorage
	sdlStorage common.ISdlSyncStorage
	ns         string
}

/*
RNibReader interface allows retrieving data from redis BD by various keys
*/
type RNibReader interface {
	// GetNodeb retrieves responding nodeb entity from redis DB by nodeb inventory name
	GetNodeb(inventoryName string) (*entities.NodebInfo, error)
	// GetNodebByGlobalNbId retrieves responding nodeb entity from redis DB by nodeb global Id
	GetNodebByGlobalNbId(nodeType entities.Node_Type, globalNbId *entities.GlobalNbId, cuupId string,duid string) (*entities.NodebInfo, error)
	// GetCellList retrieves the list of cell entities belonging to responding nodeb entity from redis DB by nodeb inventory name
	GetCellList(inventoryName string) (*entities.Cells, error)
	// GetListGnbIds retrieves the list of gNodeb identity entities
	GetListGnbIds() ([]*entities.NbIdentity, error)
	// GetListEnbIds retrieves the list of eNodeb identity entities
	GetListEnbIds() ([]*entities.NbIdentity, error)
	// Close closes reader's pool
	GetCountGnbList() (int, error)
	// GetCell retrieves the cell entity belonging to responding nodeb from redis DB by nodeb inventory name and cell pci
	GetCell(inventoryName string, pci uint32) (*entities.Cell, error)
	// GetCellById retrieves the cell entity from redis DB by cell type and cell Id
	GetCellById(cellType entities.Cell_Type, cellId string) (*entities.Cell, error)
	// GetListNodebIds returns the full list of Nodeb identity entities
	GetListNodebIds() ([]*entities.NbIdentity, error)
	// GetRanLoadInformation retrieves nodeb load information entity from redis DB by nodeb inventory name
	GetRanLoadInformation(inventoryName string) (*entities.RanLoadInformation, error)

	GetE2TInstance(address string) (*entities.E2TInstance, error)

	GetE2TInstances(addresses []string) ([]*entities.E2TInstance, error)

	GetE2TAddresses() ([]string, error)

	GetGeneralConfiguration() (*entities.GeneralConfiguration, error)

        GetRanFunctionDefinition(inventoryName string, Oid string) ([]string, error)
}

//GetNewRNibReader returns reference to RNibReader
func GetNewRNibReader(storage common.ISdlSyncStorage) RNibReader {
	return &rNibReaderInstance{
		sdl: nil,
		sdlStorage: storage,
		ns:         common.GetRNibNamespace(),
	}
}

//GetRanFunctionDefinition from the OID
func (w *rNibReaderInstance) GetRanFunctionDefinition(inventoryName string, oid string) ([]string, error){
    nb, err := w.GetNodeb (inventoryName)
    if (nb.GetGnb() != nil) {
        ranFunction := nb.GetGnb().RanFunctions
        functionDefinitionList := make([]string, 0)
        for _, ranFunction := range ranFunction {
            if (oid == ranFunction.RanFunctionOid) {
                functionDefinitionList = append(functionDefinitionList ,ranFunction.RanFunctionDefinition)
                }
        }
        return functionDefinitionList, err
    }
    return nil, common.NewResourceNotFoundErrorf("#rNibReader.GetCellList - served cells not found. Responding node RAN name: %    s.", inventoryName)
}

//GetRNibReader returns reference to RNibReader
//Deprecated: Will be removed in a future release, please use GetNewRNibReader instead.
func GetRNibReader(sdl common.ISdlInstance) RNibReader {
	return &rNibReaderInstance{
		sdl:        sdl,
		sdlStorage: nil,
		ns:         "",
	}
}

func (w *rNibReaderInstance) GetNodeb(inventoryName string) (*entities.NodebInfo, error) {
	key, rNibErr := common.ValidateAndBuildNodeBNameKey(inventoryName)
	if rNibErr != nil {
		return nil, rNibErr
	}
	nbInfo := &entities.NodebInfo{}
	err := w.getByKeyAndUnmarshal(key, nbInfo)
	if err != nil {
		return nil, err
	}
	return nbInfo, nil
}

func (w *rNibReaderInstance) GetNodebByGlobalNbId(nodeType entities.Node_Type, globalNbId *entities.GlobalNbId, cuupid string, duid string) (*entities.NodebInfo, error) {
	key, rNibErr := common.ValidateAndBuildNodeBIdKey(nodeType.String(), globalNbId.GetPlmnId(), globalNbId.GetNbId(), cuupid, duid)
	if rNibErr != nil {
		return nil, rNibErr
	}
	nbInfo := &entities.NodebInfo{}
	err := w.getByKeyAndUnmarshal(key, nbInfo)
	if err != nil {
		return nil, err
	}
	return nbInfo, nil
}

func (w *rNibReaderInstance) GetCellList(inventoryName string) (*entities.Cells, error) {
	cells := &entities.Cells{}
	nb, err := w.GetNodeb(inventoryName)
	if err != nil {
		return nil, err
	}
	if nb.GetEnb() != nil && len(nb.GetEnb().GetServedCells()) > 0 {
		cells.Type = entities.Cell_LTE_CELL
		cells.List = &entities.Cells_ServedCellInfos{ServedCellInfos: &entities.ServedCellInfoList{ServedCells: nb.GetEnb().GetServedCells()}}
		return cells, nil
	}
	if nb.GetGnb() != nil && len(nb.GetGnb().GetServedNrCells()) > 0 {
		cells.Type = entities.Cell_NR_CELL
		cells.List = &entities.Cells_ServedNrCells{ServedNrCells: &entities.ServedNRCellList{ServedCells: nb.GetGnb().GetServedNrCells()}}
		return cells, nil
	}
	return nil, common.NewResourceNotFoundErrorf("#rNibReader.GetCellList - served cells not found. Responding node RAN name: %s.", inventoryName)
}

func (w *rNibReaderInstance) GetListGnbIds() ([]*entities.NbIdentity, error) {
	return w.getListNodebIdsByType(entities.Node_GNB.String())
}

func (w *rNibReaderInstance) GetListEnbIds() ([]*entities.NbIdentity, error) {
	return w.getListNodebIdsByType(entities.Node_ENB.String())
}

func (w *rNibReaderInstance) GetCountGnbList() (int, error) {
	var size int64
	var err error
	if w.sdlStorage != nil {
		size, err = w.sdlStorage.GroupSize(w.ns, entities.Node_GNB.String())
	} else {
		size, err = w.sdl.GroupSize(entities.Node_GNB.String())
	}
	if err != nil {
		return 0, common.NewInternalError(err)
	}
	return int(size), nil
}

func (w *rNibReaderInstance) GetCell(inventoryName string, pci uint32) (*entities.Cell, error) {
	key, rNibErr := common.ValidateAndBuildCellNamePciKey(inventoryName, pci)
	if rNibErr != nil {
		return nil, rNibErr
	}
	cell := &entities.Cell{}
	err := w.getByKeyAndUnmarshal(key, cell)
	if err != nil {
		return nil, err
	}
	return cell, err
}

func (w *rNibReaderInstance) GetCellById(cellType entities.Cell_Type, cellId string) (*entities.Cell, error) {
	var key string
	var rNibErr error
	if cellType == entities.Cell_LTE_CELL {
		key, rNibErr = common.ValidateAndBuildCellIdKey(cellId)
	} else if cellType == entities.Cell_NR_CELL {
		key, rNibErr = common.ValidateAndBuildNrCellIdKey(cellId)
	} else {
		return nil, common.NewValidationErrorf("#rNibReader.GetCellById - invalid cell type: %v", cellType)
	}
	if rNibErr != nil {
		return nil, rNibErr
	}
	cell := &entities.Cell{}
	err := w.getByKeyAndUnmarshal(key, cell)
	if err != nil {
		return nil, err
	}
	return cell, err
}

func (w *rNibReaderInstance) GetListNodebIds() ([]*entities.NbIdentity, error) {
	var dataEnb, dataGnb []string
	var err error
	if w.sdlStorage != nil {
		dataEnb, err = w.sdlStorage.GetMembers(w.ns, entities.Node_ENB.String())
	} else {
		dataEnb, err = w.sdl.GetMembers(entities.Node_ENB.String())
	}
	if err != nil {
		return nil, common.NewInternalError(err)
	}
	if w.sdlStorage != nil {
		dataGnb, err = w.sdlStorage.GetMembers(w.ns, entities.Node_GNB.String())
	} else {
		dataGnb, err = w.sdl.GetMembers(entities.Node_GNB.String())
	}
	if err != nil {
		return nil, common.NewInternalError(err)
	}
	allIds := append(dataEnb, dataGnb...)
	data, rnibErr := w.unmarshalIdentityList(allIds)
	return data, rnibErr
}

func (w *rNibReaderInstance) GetRanLoadInformation(inventoryName string) (*entities.RanLoadInformation, error) {
	key, rNibErr := common.ValidateAndBuildRanLoadInformationKey(inventoryName)
	if rNibErr != nil {
		return nil, rNibErr
	}
	loadInfo := &entities.RanLoadInformation{}
	err := w.getByKeyAndUnmarshal(key, loadInfo)
	if err != nil {
		return nil, err
	}
	return loadInfo, err
}

func (w *rNibReaderInstance) GetE2TInstance(address string) (*entities.E2TInstance, error) {
	key, rNibErr := common.ValidateAndBuildE2TInstanceKey(address)
	if rNibErr != nil {
		return nil, rNibErr
	}
	e2tInstance := &entities.E2TInstance{}
	err := w.getByKeyAndUnmarshalJson(key, e2tInstance)
	if err != nil {
		return nil, err
	}
	return e2tInstance, err
}

func (w *rNibReaderInstance) GetE2TInstances(addresses []string) ([]*entities.E2TInstance, error) {
	var data map[string]interface{}
	var err error

	keys := common.MapE2TAddressesToKeys(addresses)

	e2tInstances := []*entities.E2TInstance{}

	if w.sdlStorage != nil {
		data, err = w.sdlStorage.Get(w.ns, keys)
	} else {
		data, err = w.sdl.Get(keys)
	}

	if err != nil {
		return []*entities.E2TInstance{}, common.NewInternalError(err)
	}

	if len(data) == 0 {
		return []*entities.E2TInstance{}, common.NewResourceNotFoundErrorf("#rNibReader.GetE2TInstances - e2t instances not found")
	}

	for _, v := range keys {

		if data[v] != nil {
			var e2tInstance entities.E2TInstance
			err = json.Unmarshal([]byte(data[v].(string)), &e2tInstance)
			if err != nil {
				continue
			}

			e2tInstances = append(e2tInstances, &e2tInstance)
		}
	}

	return e2tInstances, nil
}

func (w *rNibReaderInstance) GetE2TAddresses() ([]string, error) {
	var e2tAddresses []string
	err := w.getByKeyAndUnmarshalJson(E2TAddressesKey, &e2tAddresses)
	if err != nil {
		return nil, err
	}
	return e2tAddresses, err
}

func (w *rNibReaderInstance) GetGeneralConfiguration() (*entities.GeneralConfiguration, error) {
	config := &entities.GeneralConfiguration{}
	key := common.BuildGeneralConfigurationKey()

	err := w.getByKeyAndUnmarshalJson(key, config)

	return config, err
}

func (w *rNibReaderInstance) getByKeyAndUnmarshalJson(key string, entity interface{}) error {
	var data map[string]interface{}
	var err error
	if w.sdlStorage != nil {
		data, err = w.sdlStorage.Get(w.ns, []string{key})
	} else {
		data, err = w.sdl.Get([]string{key})
	}

	if err != nil {
		return common.NewInternalError(err)
	}

	if data != nil && data[key] != nil {
		err = json.Unmarshal([]byte(data[key].(string)), entity)
		if err != nil {
			return common.NewInternalError(err)
		}
		return nil
	}
	return common.NewResourceNotFoundErrorf("#rNibReader.getByKeyAndUnmarshalJson - entity of type %s not found. Key: %s", reflect.TypeOf(entity).String(), key)
}

func (w *rNibReaderInstance) getByKeyAndUnmarshal(key string, entity proto.Message) error {
	var data map[string]interface{}
	var err error
	if w.sdlStorage != nil {
		data, err = w.sdlStorage.Get(w.ns, []string{key})
	} else {
		data, err = w.sdl.Get([]string{key})
	}

	if err != nil {
		return common.NewInternalError(err)
	}
	if data != nil && data[key] != nil {
		err = proto.Unmarshal([]byte(data[key].(string)), entity)
		if err != nil {
			return common.NewInternalError(err)
		}
		return nil
	}
	return common.NewResourceNotFoundErrorf("#rNibReader.getByKeyAndUnmarshal - entity of type %s not found. Key: %s", reflect.TypeOf(entity).String(), key)
}

func (w *rNibReaderInstance) getListNodebIdsByType(nbType string) ([]*entities.NbIdentity, error) {
	var data []string
	var err error
	if w.sdlStorage != nil {
		data, err = w.sdlStorage.GetMembers(w.ns, nbType)
	} else {
		data, err = w.sdl.GetMembers(nbType)
	}
	if err != nil {
		return nil, common.NewInternalError(err)
	}
	return w.unmarshalIdentityList(data)
}

func (w *rNibReaderInstance) unmarshalIdentityList(data []string) ([]*entities.NbIdentity, error) {
	var members []*entities.NbIdentity
	for _, d := range data {
		member := entities.NbIdentity{}
		err := proto.Unmarshal([]byte(d), &member)
		if err != nil {
			return nil, common.NewInternalError(err)
		}
		members = append(members, &member)
	}
	return members, nil
}

//Close the reader
func Close() {
	// Nothing to do
}
