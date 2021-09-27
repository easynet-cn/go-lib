package golib

type SyncResult struct {
	SyncParam  interface{}   `json:"syncParam"`
	Table      string        `json:"table"`
	TotalCount int64         `json:"totalCount"`
	Result     []interface{} `json:"result"`
}

func NewSyncResult(syncParam interface{}) *SyncResult {
	return &SyncResult{SyncParam: syncParam, Result: make([]interface{}, 0)}
}
