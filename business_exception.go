package golib

type BusinessException struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Trace   string `json:"trace"`
}

func (e *BusinessException) Error() string {
	return e.Trace
}
