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

package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateAndBuildCellKeySuccess(t *testing.T) {
	cellId := "aaaa"
	prefix := "CELL"
	delimiter := ":"
	key, err := ValidateAndBuildCellIdKey(cellId)
	if err != nil{
		t.Errorf("#utils_test.TestValidateAndBuildCellKey - failed to validate key parameter")
	}
	assert.Contains(t, key, cellId)
	assert.Contains(t, key, delimiter)
	assert.Contains(t, key, prefix)
}

func TestValidateAndBuildNodeBNameKeySuccess(t *testing.T) {
	name := "name"
	prefix := "RAN"
	delimiter := ":"
	key, err := ValidateAndBuildNodeBNameKey(name)
	if err != nil{
		t.Errorf("#utils_test.TestValidateAndBuildNodeBNameKey - failed to validate key parameter")
	}
	assert.Contains(t, key, name)
	assert.Contains(t, key, delimiter)
	assert.Contains(t, key, prefix)
}

func TestValidateAndBuildNodeBIdKeySuccess(t *testing.T) {
	nodeType := "ENB"
	plmnId := "bbbb"
	nbId := "cccc"
	delimiter := ":"
	key, err := ValidateAndBuildNodeBIdKey(nodeType, plmnId, nbId)
	if err != nil{
		t.Errorf("#utils_test.TestValidateAndBuildNodeBIdKey - failed to validate key parameter")
	}
	assert.Contains(t, key, nodeType)
	assert.Contains(t, key, plmnId)
	assert.Contains(t, key, nbId)
	assert.Contains(t, key, delimiter)
}

func TestValidateAndBuildCellNamePciKeyNameValidationFailure(t *testing.T){
	pci := 9999
	_, err := ValidateAndBuildCellNamePciKey("", uint32(pci))
	assert.NotNil(t, err)
	assert.IsType(t, &ValidationError{}, err)
	assert.Equal(t, "#utils.ValidateAndBuildCellNamePciKey - an empty inventory name received", err.Error())
}

func TestValidateAndBuildCellKeyCellidValidationFailure(t *testing.T) {
	_, err := ValidateAndBuildCellIdKey("")
	assert.NotNil(t, err)
	assert.IsType(t, &ValidationError{}, err)
	assert.Equal(t, "#utils.ValidateAndBuildCellIdKey - an empty cell id received", err.Error())
}

func TestValidateAndBuildNodeBNameKeyNameValidationFailure(t *testing.T) {
	_, err := ValidateAndBuildNodeBNameKey("")
	assert.NotNil(t, err)
	assert.IsType(t, &ValidationError{}, err)
	assert.Equal(t, "#utils.ValidateAndBuildNodeBNameKey - an empty inventory name received", err.Error())
}

func TestValidateAndBuildNodeBIdKeyNodeTypeValidationFailure(t *testing.T) {
	plmnId := "dddd"
	nbId := "eeee"
	_, err := ValidateAndBuildNodeBIdKey("", plmnId, nbId)
	assert.NotNil(t, err)
	assert.IsType(t, &ValidationError{}, err)
	assert.Equal(t, "#utils.ValidateAndBuildNodeBIdKey - an empty node type received", err.Error())
}

func TestValidateAndBuildNodeBIdKeyPlmnIdValidationFailure(t *testing.T) {
	nodeType := "ffff"
	nbId := "aaaa"
	_, err := ValidateAndBuildNodeBIdKey(nodeType, "", nbId)
	assert.NotNil(t, err)
	assert.IsType(t, &ValidationError{}, err)
	assert.Equal(t, "#utils.ValidateAndBuildNodeBIdKey - an empty plmnId received", err.Error())
}

func TestValidateAndBuildNodeBIdKeyNbIdValidationFailure(t *testing.T) {
	nodeType := "bbbb"
	plmnId := "cccc"
	_, err := ValidateAndBuildNodeBIdKey(nodeType, plmnId, "")
	assert.NotNil(t, err)
	assert.IsType(t, &ValidationError{}, err)
	assert.Equal(t, "#utils.ValidateAndBuildNodeBIdKey - an empty nbId received", err.Error())
}

func TestValidateAndBuildCellNamePciKeySuccess(t *testing.T){
	name := "name"
	prefix := "PCI"
	pci := 9999
	delimiter := ":"
	key, err := ValidateAndBuildCellNamePciKey(name, uint32(pci))
	if err != nil{
		t.Errorf("#utils_test.TestValidateAndBuildCellNamePciKey - failed to validate key parameter")
	}
	assert.Contains(t, key, name)
	assert.Contains(t, key, delimiter)
	assert.Contains(t, key, prefix)
	assert.Contains(t, key, "270f")
}


func TestValidateAndBuildRanLoadInformationKeySuccess(t *testing.T) {
	name := "name"
	prefix := "LOAD"
	delimiter := ":"
	key, err := ValidateAndBuildRanLoadInformationKey(name)
	if err != nil{
		t.Errorf("#utils_test.TestValidateAndBuildRanLoadInformationKeySuccess - failed to validate key parameter")
	}
	assert.Contains(t, key, name)
	assert.Contains(t, key, delimiter)
	assert.Contains(t, key, prefix)
}

func TestValidateAndBuildRanLoadInformationKeyFailure(t *testing.T) {
	name := ""
	_, err := ValidateAndBuildRanLoadInformationKey(name)
	assert.NotNil(t, err)
	assert.IsType(t, &ValidationError{}, err)
}