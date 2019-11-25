package entities

type E2TInstanceInfo struct {
	Address string `json:"address"`
	AssociatedRanCount int `json:"associatedRanCount"`
}

func NewE2TInstanceInfo(address string) *E2TInstanceInfo {
	return &E2TInstanceInfo{
		Address: address,
	}
}