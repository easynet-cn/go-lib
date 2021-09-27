package golib

type RestResult struct {
	Success int         `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
