package controller

import (
	"go-challenger/adapter/api/handler"
	"go-challenger/adapter/api/response"
	"go-challenger/core/usecase"
	"go-challenger/core/usecase/input"
	"net/http"
)

type GetTaskController struct {
	uc usecase.FindByIdUseCaseInterface
}

var _ Controller = (*GetTaskController)(nil)


func NewGetTaskController(uc usecase.FindByIdUseCaseInterface) *GetTaskController{
	return &GetTaskController{
		uc: uc,
	}
}

func (c GetTaskController) Execute(w http.ResponseWriter, r *http.Request){
	var taskId = r.URL.Query().Get("taskId")
	
	i := input.TaskIdInput{Id: taskId}
	
	task, err :=c.uc.Execute(r.Context(),i)
	if err != nil{
		handler.HandleError(w,err)
		return
	}

	response.NewSuccess(http.StatusOK,task).Send(w)
}
