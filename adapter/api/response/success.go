package response

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	StatusCode int
	Result     interface{}
}

func NewSuccess(status int, res interface{}) *Success {
	return &Success{
		StatusCode: status,
		Result:     res,
	}
}

func (s Success) Send(w http.ResponseWriter){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(s.StatusCode)
	if errEncoder := json.NewEncoder(w).Encode(s.Result); errEncoder != nil{
		return
	}
}