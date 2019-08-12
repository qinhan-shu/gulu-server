package req

import (
	"encoding/json"
)

// PvpInfo
type PvpInfo struct {
	ID    int `json:"id"`
	Arena int `json:"arena"`
	Cup   int `json:"cup"`
}

func PvpInfoReq(data []byte) *PvpInfo {
	p := &PvpInfo{}
	if err := json.Unmarshal(data, &p); err != nil {
		return &PvpInfo{}
	}
	return p
}
