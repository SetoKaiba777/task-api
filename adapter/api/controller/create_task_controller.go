package controller

import (
	"encoding/json"
	"go-challenger/adapter/api/handler"
	"go-challenger/adapter/api/response"
	"go-challenger/core/usecase"
	"go-challenger/core/usecase/input"
	"io"
	"net/http"
)

type CreateTaskController struct {
	uc usecase.SaveUseCaseInterface
}

var _ Controller = (*CreateTaskController)(nil)

func NewCreateTaskController(uc usecase.SaveUseCaseInterface) CreateTaskController{
	return CreateTaskController{
		uc : uc,
	}
}

func (c CreateTaskController) Execute(w http.ResponseWriter, r *http.Request){
	jsonBody, err := io.ReadAll(r.Body)
	if err != nil{
		handler.HandleError(w,err)
		return
	}

	var input input.TaskInput
	if err:= json.Unmarshal(jsonBody,&input); err != nil{
		handler.HandleError(w,err)
		return
	}	

	if err := c.uc.Execute(r.Context(),input); err != nil{
		handler.HandleError(w,err)
		return
	}

	response.NewSuccess(http.StatusCreated,"").Send(w)
}