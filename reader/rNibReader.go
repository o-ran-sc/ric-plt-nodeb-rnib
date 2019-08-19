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
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/common"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities"
	"gerrit.o-ran-sc.org/r/ric-plt/sdlgo"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"reflect"
)

var readerPool *common.Pool

type rNibReaderInstance struct {
	sdl       *common.ISdlInstance
	namespace string
}

/*
RNibReader interface allows retrieving data from redis BD by various keys
*/
type RNibReader interface {
	// GetNodeb retrieves responding nodeb entity from redis DB by nodeb inventory name
	GetNodeb(inventoryName string) (*entities.NodebInfo, error)
	// GetNodebByGlobalNbId retrieves responding nodeb entity from redis DB by nodeb global Id
	GetNodebByGlobalNbId(nodeType entities.Node_Type, globalNbId *entities.GlobalNbId) (*entities.NodebInfo, error)
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
	GetListNodebIds()([]*entities.NbIdentity, error)
	// GetRanLoadInformation retrieves nodeb load information entity from redis DB by nodeb inventory name
	GetRanLoadInformation(inventoryName string) (*entities.RanLoadInformation, error)
}

/*
 Init initializes the infrastructure required for the RNibReader instance
*/
func Init(namespace string, poolSize int) {
	initPool(poolSize,
		func() interface{} {
			var sdlI common.ISdlInstance = sdlgo.NewSdlInstance(namespace, sdlgo.NewDatabase())
			return &rNibReaderInstance{sdl: &sdlI, namespace: namespace}
		},
		func(obj interface{}) {
			i, ok := obj.(*rNibReaderInstance)
			if ok{
				(*i.sdl).Close()
			}
		})
}

func initPool(poolSize int, newObj func() interface{}, destroyObj func(interface{})) {
	readerPool = common.NewPool(poolSize, newObj, destroyObj)
}

/*
GetRNibReader returns RNibReader instance from the pool
*/
func GetRNibReader() RNibReader {
	return readerPool.Get().(RNibReader)
}

func (w *rNibReaderInstance) GetNodeb(inventoryName string) (*entities.NodebInfo, error) {
	defer readerPool.Put(w)
	key, rNibErr := common.ValidateAndBuildNodeBNameKey(inventoryName)
	if rNibErr != nil {
		return nil, rNibErr
	}
	nbInfo := &entities.NodebInfo{}
	err := w.getByKeyAndUnmarshal(key, nbInfo)
	if err!= nil{
		return nil, err
	}
	return nbInfo, nil
}

func (w *rNibReaderInstance) GetNodebByGlobalNbId(nodeType entities.Node_Type, globalNbId *entities.GlobalNbId) (*entities.NodebInfo, error) {
	defer readerPool.Put(w)
	key, rNibErr := common.ValidateAndBuildNodeBIdKey(nodeType.String(), globalNbId.GetPlmnId(), globalNbId.GetNbId())
	if rNibErr != nil {
		return nil, rNibErr
	}
	nbInfo := &entities.NodebInfo{}
	err := w.getByKeyAndUnmarshal(key, nbInfo)
	if err!= nil{
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
	} else if nb.GetGnb() != nil && len(nb.GetGnb().GetServedNrCells()) > 0 {
		cells.Type = entities.Cell_NR_CELL
		cells.List = &entities.Cells_ServedNrCells{ServedNrCells: &entities.ServedNRCellList{ServedCells: nb.GetGnb().GetServedNrCells()}}
		return cells, nil
	}
	return nil, common.NewResourceNotFoundError(errors.Errorf("#rNibReader.GetCellList - served cells not found. Responding node RAN name: %s.", inventoryName))
}

func (w *rNibReaderInstance) GetListGnbIds() ([]*entities.NbIdentity, error) {
	defer readerPool.Put(w)
	return w.getListNodebIdsByType(entities.Node_GNB.String())
}

func (w *rNibReaderInstance) GetListEnbIds() ([]*entities.NbIdentity, error) {
	defer readerPool.Put(w)
	return w.getListNodebIdsByType(entities.Node_ENB.String())
}

func (w *rNibReaderInstance) GetCountGnbList() (int, error) {
	defer readerPool.Put(w)
	size, err := (*w.sdl).GroupSize(entities.Node_GNB.String())
	if err != nil {
		return 0, common.NewInternalError(err)
	}
	return int(size), nil
}

func (w *rNibReaderInstance) GetCell(inventoryName string, pci uint32) (*entities.Cell, error) {
	defer readerPool.Put(w)
	key, rNibErr := common.ValidateAndBuildCellNamePciKey(inventoryName, pci)
	if rNibErr != nil {
		return nil, rNibErr
	}
	cell := &entities.Cell{}
	err := w.getByKeyAndUnmarshal(key, cell)
	if err!= nil{
		return nil, err
	}
	return cell, err
}

func (w *rNibReaderInstance) GetCellById(cellType entities.Cell_Type, cellId string) (*entities.Cell, error) {
	defer readerPool.Put(w)
	var key string
	var rNibErr error
	if cellType == entities.Cell_LTE_CELL {
		key, rNibErr = common.ValidateAndBuildCellIdKey(cellId)
	} else if cellType == entities.Cell_NR_CELL {
		key, rNibErr = common.ValidateAndBuildNrCellIdKey(cellId)
	} else {
		return nil, common.NewValidationError(errors.Errorf("#rNibReader.GetCellById - invalid cell type: %v", cellType))
	}
	if rNibErr != nil {
		return nil, rNibErr
	}
	cell := &entities.Cell{}
	err := w.getByKeyAndUnmarshal(key, cell)
	if err!= nil{
		return nil, err
	}
	return cell, err
}

func (w *rNibReaderInstance) GetListNodebIds()([]*entities.NbIdentity, error){
	defer readerPool.Put(w)
	dataEnb, err := (*w.sdl).GetMembers(entities.Node_ENB.String())
	if err != nil{
		return nil, common.NewInternalError(err)
	}
	dataGnb, err := (*w.sdl).GetMembers(entities.Node_GNB.String())
	if err != nil{
		return nil, common.NewInternalError(err)
	}
	dataUnknown, err := (*w.sdl).GetMembers(entities.Node_UNKNOWN.String())
	if err != nil{
		return nil, common.NewInternalError(err)
	}
	allIds := append(dataEnb, dataGnb...)
	allIds = append(allIds, dataUnknown...)
	data, rnibErr := unmarshalIdentityList(allIds)
	return data, rnibErr
}

func (w *rNibReaderInstance) GetRanLoadInformation(inventoryName string) (*entities.RanLoadInformation, error){
	key, rNibErr := common.ValidateAndBuildRanLoadInformationKey(inventoryName)
	if rNibErr != nil {
		return nil, rNibErr
	}
	loadInfo := &entities.RanLoadInformation{}
	err := w.getByKeyAndUnmarshal(key, loadInfo)
	if err!= nil{
		return nil, err
	}
	return loadInfo, err
}

func (w *rNibReaderInstance) getByKeyAndUnmarshal(key string, entity proto.Message)error{
	data, err := (*w.sdl).Get([]string{key})
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
	return common.NewResourceNotFoundError(errors.Errorf("#rNibReader.getByKeyAndUnmarshal - entity of type %s not found. Key: %s", reflect.TypeOf(entity).String(), key))
}

func (w *rNibReaderInstance) getListNodebIdsByType(nbType string) ([]*entities.NbIdentity, error) {
	data, err := (*w.sdl).GetMembers(nbType)
	if err != nil {
		return nil, common.NewInternalError(err)
	}
	return unmarshalIdentityList(data)
}

func unmarshalIdentityList(data []string) ([]*entities.NbIdentity, error) {
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

func Close() {
	readerPool.Close()
}