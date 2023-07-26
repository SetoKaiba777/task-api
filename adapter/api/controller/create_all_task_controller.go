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

type CreateAllTaskController struct {
	uc usecase.SaveAllUseCaseInterface
}

var _ Controller = (*CreateTaskController)(nil)

func NewCreateAllTaskController(uc usecase.SaveAllUseCaseInterface) CreateAllTaskController{
	return CreateAllTaskController{
		uc : uc,
	}
}

func (c CreateAllTaskController) Execute(w http.ResponseWriter, r *http.Request){
	jsonBody, err := io.ReadAll(r.Body)
	if err != nil{
		handler.HandleError(w,err)
		return
	}

	var inputs []input.TaskInput
	if err := json.Unmarshal([]byte(jsonBody), &inputs); err != nil{
		handler.HandleError(w,err)
		return
	}	

	if err := c.uc.Execute(r.Context(),inputs); err != nil{
		handler.HandleError(w,err)
		return
	}

	response.NewSuccess(http.StatusCreated,"Create All successfuly").Send(w)
}