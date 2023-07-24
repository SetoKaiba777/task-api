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

type UpdateTaskController struct {
	uc usecase.UpdateUseCaseInterface
}

var _ Controller = (*UpdateTaskController)(nil)

func NewUpdateTaskController(uc usecase.UpdateUseCaseInterface) UpdateTaskController{
	return UpdateTaskController{
		uc : uc,
	}
}

func (c UpdateTaskController) Execute(w http.ResponseWriter, r *http.Request){

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

	var taskId = r.URL.Query().Get("taskId")	
	input.Id = taskId

	if err := c.uc.Execute(r.Context(),input); err != nil{
		handler.HandleError(w,err)
	}

	response.NewSuccess(http.StatusOK,input).Send(w)
}