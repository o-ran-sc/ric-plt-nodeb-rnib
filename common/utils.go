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
	"fmt"
)

/*
ValidateAndBuildCellIdKey builds key according to the specified format returns the resulting string
*/
func ValidateAndBuildCellIdKey(cellId string) (string, error) {
	if cellId == "" {
		return "", NewValidationError("#utils.ValidateAndBuildCellIdKey - an empty cell id received")
	}
	return fmt.Sprintf("CELL:%s", cellId), nil
}

/*
ValidateAndBuildNrCellIdKey builds key according to the specified format returns the resulting string
*/
func ValidateAndBuildNrCellIdKey(cellId string) (string, error) {
	if cellId == "" {
		return "", NewValidationError("#utils.ValidateAndBuildNrCellIdKey - an empty cell id received")
	}
	return fmt.Sprintf("NRCELL:%s", cellId), nil
}

/*
ValidateAndBuildNodeBNameKey builds key according to the specified format returns the resulting string
*/
func ValidateAndBuildNodeBNameKey(inventoryName string) (string, error) {
	if inventoryName == "" {
		return "", NewValidationError("#utils.ValidateAndBuildNodeBNameKey - an empty inventory name received")
	}
	return fmt.Sprintf("RAN:%s", inventoryName), nil
}

/*
ValidateAndBuildNodeBIdKey builds key according to the specified format returns the resulting string
*/
func ValidateAndBuildNodeBIdKey(nodeType string, plmnId string, nbId string) (string, error) {
	if nodeType == "" {
		return "", NewValidationError("#utils.ValidateAndBuildNodeBIdKey - an empty node type received")
	}
	if plmnId == "" {
		return "", NewValidationError("#utils.ValidateAndBuildNodeBIdKey - an empty plmnId received")
	}
	if nbId == "" {
		return "", NewValidationError("#utils.ValidateAndBuildNodeBIdKey - an empty nbId received")
	}
	return fmt.Sprintf("%s:%s:%s", nodeType, plmnId, nbId), nil
}

/*
ValidateAndBuildCellNamePciKey builds key according to the specified format returns the resulting string
*/
func ValidateAndBuildCellNamePciKey(inventoryName string, pci uint32) (string, error) {
	if inventoryName == "" {
		return "", NewValidationError("#utils.ValidateAndBuildCellNamePciKey - an empty inventory name received")
	}
	return fmt.Sprintf("PCI:%s:%02x", inventoryName, pci), nil
}

func ValidateAndBuildRanLoadInformationKey(inventoryName string) (string, error) {

	if inventoryName == "" {
		return "", NewValidationError("#utils.ValidateAndBuildRanLoadInformationKey - an empty inventory name received")
	}

	return fmt.Sprintf("LOAD:%s", inventoryName), nil
}