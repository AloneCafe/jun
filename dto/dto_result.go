package dto

type ResultStatus bool

type Result struct {
	Status  ResultStatus `json:"r_status"`
	Message *string      `json:"r_message"`
	Data    interface{}  `json:"r_data"`
}

func NewResult(status ResultStatus, message string, data interface{}) *Result {
	return &Result{status, &message, data}
}
