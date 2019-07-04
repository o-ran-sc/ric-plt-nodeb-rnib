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
//

package common

import "fmt"

const(
	RESOURCE_NOT_FOUND int = 1
	INTERNAL_ERROR int = 2
	VALIDATION_ERROR int = 3
)

var rNibError_names = map[int]string{
	RESOURCE_NOT_FOUND:"RESOURCE_NOT_FOUND",
	INTERNAL_ERROR:"INTERNAL_ERROR",
	VALIDATION_ERROR:"VALIDATION_ERROR",
}

type IRNibError interface{
	error
	GetCode() int
	GetError() error
}

type RNibError struct{
	err error
	errorCode int
}

func NewResourceNotFoundError(error error) IRNibError {
	return IRNibError(&RNibError{err:error, errorCode:RESOURCE_NOT_FOUND})
}

func NewInternalError(error error) IRNibError {
	return IRNibError(&RNibError{err:error, errorCode:INTERNAL_ERROR})
}

func NewValidationError(error error) IRNibError {
	return IRNibError(&RNibError{err:error, errorCode:VALIDATION_ERROR})
}

func (e RNibError) GetError() error {
	return e.err
}

func (e RNibError) GetCode() int {
	return e.errorCode
}

func (e RNibError) Error() string {
	return fmt.Sprintf("%d %s - %s", e.errorCode, rNibError_names[e.errorCode], e.err)
}