package model

import (
    "bytes"
    "github.com/chatpuppy/puppychat/pkg/tea"
    jsoniter "github.com/json-iterator/go"
)

const (
    SyncStatusPending uint8 = iota // not synced
)

func SyncMarshal(data any) []byte {
    var buffer bytes.Buffer
    b, _ := jsoniter.Marshal(data)
    buffer.Write(b)
    buffer.WriteRune('\n')
    return buffer.Bytes()
}

type SyncRequest struct {
    NodeID      int64     `json:"nodeId"`
    Signature   []byte    `json:"signature"`
    Nonce       []byte    `json:"nonce"`
    Data        *SyncData `json:"data,omitempty"`
    SyncedStart *uint64   `json:"syncedStart,omitempty"`
    ApiUrl      *string   `json:"apiUrl,omitempty"`
}

func (r *SyncRequest) Marshal() []byte {
    return SyncMarshal(r)
}

type SyncResponse struct {
    Data     *SyncData `json:"data,omitempty"`
    Error    *string   `json:"error,omitempty"`
    Current  *uint64   `json:"current,omitempty"`  // current synced index
    Progress *float64  `json:"progress,omitempty"` // sync progress precent
}

func (r *SyncResponse) Marshal() []byte {
    return SyncMarshal(r)
}

func SyncResError(err error) []byte {
    res := &SyncResponse{
        Data:  nil,
        Error: tea.String(err.Error()),
    }
    return res.Marshal()
}

type SyncData struct {
    Table  string `json:"table"`
    Value  []byte `json:"value"`
    Op     uint   `json:"op"`
    Index  uint64 `json:"index"`
    NodeID int64  `json:"nodeId"`
}

func (d *SyncData) MarshalBinary() (data []byte, err error) {
    return jsoniter.Marshal(d)
}

func (d *SyncData) UnmarshalBinary(data []byte) error {
    return jsoniter.Unmarshal(data, d)
}
