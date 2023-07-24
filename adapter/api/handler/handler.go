package handler

import (
	"go-challenger/adapter/api/response"
	"go-challenger/core/domain"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error){
	var status int

	switch err{
	case domain.ErrInvalidStatus:
		status = http.StatusBadRequest
	case domain.ErrNotFoundTask:
		status = http.StatusNotFound
	default:
		status = http.StatusInternalServerError
	}

	response.NewError(err,status).Send(w)
}