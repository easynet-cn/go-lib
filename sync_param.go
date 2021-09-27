package golib

type SyncParam struct {
	PageParam
	ReturnTotalCount int    `json:"returnTotalCount" url:"returnTotalCount"`
	ReturnResult     int    `json:"returnResult" url:"returnResult"`
	LastId           *int64 `json:"lastId" url:"lastId"`
	SyncTime         string `json:"syncTime" url:"syncTime"`
}
