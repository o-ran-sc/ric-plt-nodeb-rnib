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

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResourceNotFoundError(t *testing.T){
	e := errors.New("Expected error")
	expectedErr := NewResourceNotFoundError(e)
	assert.NotNil(t, expectedErr)
	assert.Equal(t, RESOURCE_NOT_FOUND, expectedErr.GetCode())
	assert.Contains(t, expectedErr.Error(), fmt.Sprintf("%d",expectedErr.GetCode()))
	assert.Contains(t, expectedErr.Error(), rNibError_names[expectedErr.GetCode()])
	assert.Contains(t, expectedErr.Error(), e.Error())
	assert.Equal(t, RESOURCE_NOT_FOUND, expectedErr.GetCode())
}

func TestNewInternalError(t *testing.T){
	e := errors.New("Expected error")
	expectedErr := NewInternalError(e)
	assert.NotNil(t, expectedErr)
	assert.Equal(t, INTERNAL_ERROR, expectedErr.GetCode())
	assert.Contains(t, expectedErr.Error(), fmt.Sprintf("%d",expectedErr.GetCode()))
	assert.Contains(t, expectedErr.Error(), rNibError_names[expectedErr.GetCode()])
	assert.Contains(t, expectedErr.Error(), e.Error())
	assert.Equal(t, INTERNAL_ERROR, expectedErr.GetCode())
}

func TestNewValidationError(t *testing.T){
	e := errors.New("Expected error")
	expectedErr := NewValidationError(e)
	assert.NotNil(t, expectedErr)
	assert.Equal(t, VALIDATION_ERROR, expectedErr.GetCode())
	assert.Contains(t, expectedErr.Error(), fmt.Sprintf("%d",expectedErr.GetCode()))
	assert.Contains(t, expectedErr.Error(), rNibError_names[expectedErr.GetCode()])
	assert.Contains(t, expectedErr.Error(), e.Error())
	assert.Equal(t, VALIDATION_ERROR, expectedErr.GetCode())
}

func TestNewRNibErrorWithEmptyError(t *testing.T){
	var e error
	expectedErr := NewResourceNotFoundError(e)
	assert.NotNil(t, expectedErr)
	assert.Equal(t, RESOURCE_NOT_FOUND, expectedErr.GetCode())
	assert.Contains(t, expectedErr.Error(), fmt.Sprintf("%d",expectedErr.GetCode()))
	assert.Contains(t, expectedErr.Error(), rNibError_names[expectedErr.GetCode()])
	assert.Equal(t, 1, expectedErr.GetCode())
}

func TestGetError(t *testing.T){
	e := errors.New("Expected error")
	expectedErr := NewInternalError(e)
	assert.NotNil(t, expectedErr)
	assert.NotNil(t, expectedErr.GetError())
	assert.Equal(t, expectedErr.GetError(), e)
}

func TestGetCodeInternalError(t *testing.T){
	e := errors.New("Expected error")
	expectedErr := NewInternalError(e)
	assert.NotNil(t, expectedErr)
	assert.NotNil(t, expectedErr.GetError())
	assert.Equal(t, INTERNAL_ERROR, expectedErr.GetCode())
}

func TestGetCodeNotFoundError(t *testing.T){
	e := errors.New("Expected error")
	expectedErr := NewResourceNotFoundError(e)
	assert.NotNil(t, expectedErr)
	assert.NotNil(t, expectedErr.GetError())
	assert.Equal(t, RESOURCE_NOT_FOUND, expectedErr.GetCode())
}

func TestGetCodeValidationError(t *testing.T){
	e := errors.New("Expected error")
	expectedErr := NewValidationError(e)
	assert.NotNil(t, expectedErr)
	assert.NotNil(t, expectedErr.GetError())
	assert.Equal(t, VALIDATION_ERROR, expectedErr.GetCode())
}