package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	statusCode int
	erros      []string
}

func NewError(err error, status int) *Error {
	return &Error{
		statusCode: status,
		erros:      []string{err.Error()},
	}
}

func (e Error) Send(w http.ResponseWriter){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(e.statusCode)
	if errEncoder := json.NewEncoder(w).Encode(e.erros); errEncoder != nil {
		return
	}
}