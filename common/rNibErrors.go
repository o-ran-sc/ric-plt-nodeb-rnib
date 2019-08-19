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

type ResourceNotFoundError struct{
	err error
}

type InternalError struct{
	err error
}

type ValidationError struct{
	err error
}

func NewResourceNotFoundError(error error) error {
	return &ResourceNotFoundError{err:error}
}

func NewInternalError(error error) error {
	return &InternalError{err:error}
}

func NewValidationError(error error) error {
	return &ValidationError{err:error}
}

func (e ResourceNotFoundError) Error() string {
	return e.err.Error()
}

func (e InternalError) Error() string {
	return e.err.Error()
}

func (e ValidationError) Error() string {
	return e.err.Error()
}