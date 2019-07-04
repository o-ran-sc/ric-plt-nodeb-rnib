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
	"gerrit.ranco-dev-tools.eastus.cloudapp.azure.com/ric-plt/nodeb-rnib.git/entities"
	"gerrit.ranco-dev-tools.eastus.cloudapp.azure.com/ric-plt/nodeb-rnib.git/common"
	"gerrit.o-ran-sc.org/r/ric-plt/sdlgo"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
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
	GetNodeb(inventoryName string) (*entities.NodebInfo, common.IRNibError)
	// GetNodebByGlobalNbId retrieves responding nodeb entity from redis DB by nodeb global Id
	GetNodebByGlobalNbId(nodeType entities.Node_Type, globalNbId *entities.GlobalNbId) (*entities.NodebInfo, common.IRNibError)
	// GetCellList retrieves the list of cell entities belonging to responding nodeb entity from redis DB by nodeb inventory name
	GetCellList(inventoryName string) (*entities.Cells, common.IRNibError)
	// GetListGnbIds retrieves the list of gNodeb identity entities
	GetListGnbIds()(*[]*entities.NbIdentity, common.IRNibError)
	// GetListEnbIds retrieves the list of eNodeb identity entities
	GetListEnbIds()(*[]*entities.NbIdentity, common.IRNibError)
	// Close closes reader's pool
	GetCountGnbList()(int, common.IRNibError)
	// GetCell retrieves the cell entity belonging to responding nodeb from redis DB by nodeb inventory name and cell pci
	GetCell(inventoryName string, pci uint32) (*entities.Cell, common.IRNibError)
	// GetCellById retrieves the cell entity from redis DB by cell type and cell Id
	GetCellById(cellType entities.Cell_Type, cellId string) (*entities.Cell, common.IRNibError)
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
		(*obj.(*rNibReaderInstance).sdl).Close()
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

func (w *rNibReaderInstance) GetNodeb(inventoryName string) (*entities.NodebInfo, common.IRNibError) {
	name, rNibErr := common.ValidateAndBuildNodeBNameKey(inventoryName)
	if rNibErr != nil{
		return nil, rNibErr
	}
	defer readerPool.Put(w)
	data, err := (*w.sdl).Get([]string{name})
	if err != nil{
		return nil, common.NewInternalError(err)
	}
	nb := entities.NodebInfo{}
	if data != nil && data[name] != nil{
		err = proto.Unmarshal( []byte(data[name].(string)), &nb)
		if err != nil{
			return nil, common.NewInternalError(err)
		}
		return &nb, nil
	}
	return nil, common.NewResourceNotFoundError(errors.Errorf("#rNibReader.GetNodeb - responding node %s not found", inventoryName))
}

func (w *rNibReaderInstance) GetNodebByGlobalNbId(nodeType entities.Node_Type, globalNbId *entities.GlobalNbId) (*entities.NodebInfo, common.IRNibError) {
	key, rNibErr := common.ValidateAndBuildNodeBIdKey(nodeType.String(), globalNbId.GetPlmnId(), globalNbId.GetNbId())
	if rNibErr != nil{
		return nil, rNibErr
	}
	defer readerPool.Put(w)
	data, err := (*w.sdl).Get([]string{key})
	if err != nil{
		return nil, common.NewInternalError(err)
	}
	nb := entities.NodebInfo{}
	if data != nil && data[key] != nil{
		err = proto.Unmarshal( []byte(data[key].(string)), &nb)
		if err != nil{
			return nil, common.NewInternalError(err)
		}
		return &nb, nil
	}
	return nil, common.NewResourceNotFoundError(errors.Errorf("#rNibReader.GetNodebByGlobalNbId - responding node not found, global nodeb Id: %s", key))
}

func (w *rNibReaderInstance) GetCellList(inventoryName string) (*entities.Cells, common.IRNibError){
	cells := &entities.Cells{}
	nb, err := w.GetNodeb(inventoryName)
	if err != nil{
		return nil, err
	}
	if nb.GetEnb() != nil && len(nb.GetEnb().GetServedCells()) > 0{
		cells.Type = entities.Cell_LTE_CELL
		cells.List = &entities.Cells_ServedCellInfos{ServedCellInfos:&entities.ServedCellInfoList{ServedCells:nb.GetEnb().GetServedCells()}}
		return cells, nil
	} else if nb.GetGnb() != nil && len(nb.GetGnb().GetServedNrCells()) > 0{
		cells.Type = entities.Cell_NR_CELL
		cells.List = &entities.Cells_ServedNrCells{ServedNrCells:&entities.ServedNRCellList{ServedCells:nb.GetGnb().GetServedNrCells()}}
		return cells, nil
	}
	return nil, common.NewResourceNotFoundError(errors.Errorf("#rNibReader.GetCellList - served cells not found. Responding node RAN name: %s.", inventoryName))
}

func (w *rNibReaderInstance) GetListGnbIds()(*[]*entities.NbIdentity, common.IRNibError){
	defer readerPool.Put(w)
	data, err := (*w.sdl).GetMembers("GNB")
	if err != nil{
		return nil, common.NewInternalError(err)
	}
	return unmarshalIdentityList(data)
}

func (w *rNibReaderInstance) GetListEnbIds()(*[]*entities.NbIdentity, common.IRNibError){
	defer readerPool.Put(w)
	data, err := (*w.sdl).GetMembers("ENB")
	if err != nil{
		return nil, common.NewInternalError(err)
	}
	return unmarshalIdentityList(data)
}

func (w *rNibReaderInstance) GetCountGnbList()(int, common.IRNibError){
	defer readerPool.Put(w)
	data, err := (*w.sdl).GetMembers("GNB")
	if err != nil{
		return 0, common.NewInternalError(err)
	}
	return len(data), nil
}

func (w *rNibReaderInstance) GetCell(inventoryName string, pci uint32) (*entities.Cell, common.IRNibError){
	key, rNibErr := common.ValidateAndBuildCellNamePciKey(inventoryName, pci)
	if rNibErr != nil{
		return nil, rNibErr
	}
	return (*w).getCellByKey(key)
}

func (w *rNibReaderInstance) GetCellById(cellType entities.Cell_Type, cellId string) (*entities.Cell, common.IRNibError){
	var key string
	var rNibErr common.IRNibError
	if cellType == entities.Cell_LTE_CELL {
		key, rNibErr = common.ValidateAndBuildCellIdKey(cellId)
	} else if cellType == entities.Cell_NR_CELL {
		key, rNibErr = common.ValidateAndBuildNrCellIdKey(cellId)
	} else {
		return nil, common.NewValidationError(errors.Errorf("#rNibReader.GetCellById - invalid cell type: %v", cellType))
	}
	if rNibErr != nil{
		return nil, rNibErr
	}
	return (*w).getCellByKey(key)
}

func (w *rNibReaderInstance) getCellByKey(key string) (*entities.Cell, common.IRNibError){
	defer readerPool.Put(w)
	data, err := (*w.sdl).Get([]string{key})
	if err != nil{
		return nil, common.NewInternalError(err)
	}
	cell := entities.Cell{}
	if data != nil && data[key] != nil{
		err = proto.Unmarshal( []byte(data[key].(string)), &cell)
		if err != nil{
			return nil, common.NewInternalError(err)
		}
		return &cell, nil
	}
	return nil, common.NewResourceNotFoundError(errors.Errorf("#rNibReader.getCellByKey - cell not found, key: %s", key))
}

func unmarshalIdentityList(data []string ) (*[]*entities.NbIdentity, common.IRNibError){
	var members []*entities.NbIdentity
	for _, d := range data{
		member := entities.NbIdentity{}
		err := proto.Unmarshal([]byte(d), &member)
		if err != nil{
			return nil, common.NewInternalError(err)
		}
		members = append(members, &member)
	}
	return &members, nil
}

func Close(){
	readerPool.Close()
}