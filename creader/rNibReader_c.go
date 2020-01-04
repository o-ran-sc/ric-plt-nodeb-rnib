package main

import (
	"encoding/json"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/common"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/reader"
	"gerrit.o-ran-sc.org/r/ric-plt/sdlgo"
	"unsafe"
)

//#include <string.h>
import "C"

var sdl common.ISdlInstance
var instance reader.RNibReader

type response struct {
	GnbList  []string `json:"gnb_list"`
	ErrorMsg string   `json:"error_msg,omitempty"`
}

//export open
func open() {
	sdl = sdlgo.NewSdlInstance("e2Manager", sdlgo.NewDatabase())
	instance = reader.GetRNibReader(sdl)
}

//export close
func close() {
	_ = sdl.Close()
}

//export getListGnbIds
func getListGnbIds() unsafe.Pointer {
	listGnbIds, err := instance.GetListGnbIds()
	res := &response{
		GnbList: []string{},
	}

	if err != nil {
		res.ErrorMsg = err.Error()

		return createCBytesResponse(res)
	}

	for _, value := range listGnbIds {
		res.GnbList = append(res.GnbList, value.InventoryName)
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

func main() {

}
