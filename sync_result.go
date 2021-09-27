package golib

type SyncResult struct {
	SyncParam  SyncParam     `json:"syncParam"`
	Table      string        `json:"table"`
	TotalCount int64         `json:"totalCount"`
	Result     []interface{} `json:"result"`
}

func NewSyncResult(syncParam SyncParam) *SyncResult {
	return &SyncResult{SyncParam: syncParam, Result: make([]interface{}, 0)}
}
