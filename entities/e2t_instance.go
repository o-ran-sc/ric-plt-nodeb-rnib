package entities

import "time"

type E2TInstance struct {
	Address string `json:"address"`
	AssociatedRanList   []string `json:"associatedRanList"`
	KeepAliveTimestamp int64 `json:"keepAliveTimestamp"`
	State E2TInstanceState `json:"state"`
}

func NewE2TInstance(address string) *E2TInstance {
	return &E2TInstance{
		Address: address,
		KeepAliveTimestamp:time.Now().UnixNano(),
		State: Active,
	}
}