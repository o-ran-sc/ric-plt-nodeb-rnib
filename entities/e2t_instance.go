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

package entities

import "time"

type E2TInstance struct {
	Address            string           `json:"address"`
	PodName            string           `json:"podName"`
	AssociatedRanList  []string         `json:"associatedRanList"`
	KeepAliveTimestamp int64            `json:"keepAliveTimestamp"`
	State              E2TInstanceState `json:"state"`
	DeletionTimestamp  int64            `json:"deletionTimeStamp"`
}

func NewE2TInstance(address string, podName string) *E2TInstance {
		return &E2TInstance{
		Address:            address,
		KeepAliveTimestamp: time.Now().UnixNano(),
		State:              Active,
		AssociatedRanList:  []string{},
		PodName:            podName,
	}
}
