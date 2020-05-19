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

package main

import (
	"encoding/json"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/common"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/reader"
	"gerrit.o-ran-sc.org/r/ric-plt/sdlgo"
	"unsafe"
)

//#include <string.h>
import "C"

var sdl common.ISdlInstance
var instance reader.RNibReader

type response struct {
	GnbList  []*entities.NbIdentity `json:"gnb_list"`
	ErrorMsg string                 `json:"error_msg,omitempty"`
}

//export openSdl
func openSdl() {
	sdl = sdlgo.NewSdlInstance("e2Manager", sdlgo.NewDatabase())
	instance = reader.GetRNibReader(sdl)
}

//export closeSdl
func closeSdl() {
	_ = sdl.Close()
}

//export getListGnbIds
func getListGnbIds() unsafe.Pointer {
	listGnbIds, err := instance.GetListGnbIds()
	res := &response{}

	if err != nil {
		res.ErrorMsg = err.Error()

		return createCBytesResponse(res)
	}

	if listGnbIds != nil {
		res.GnbList = listGnbIds
	}

	return createCBytesResponse(res)
}

func createCBytesResponse(res *response) unsafe.Pointer {
	byteResponse, err := json.Marshal(res)
	if err != nil {
		return nil
	}

	return C.CBytes(byteResponse)
}

func main() {}
