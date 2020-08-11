package model

import "encoding/json"

//Errors ...
type Errors struct {
	Code int `json:"code"`
	Debug string `json:"debug"`
	Details struct{} `json:"details,omitempty"`
	Message string `json:"message,omitempty"`
	Reason string `json:"reason,omitempty"`
	Request string `json:"request"`
	Status string `json:"status"`
}

func (e *Errors) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}
