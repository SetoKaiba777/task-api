package controller

import (
	"go-challenger/adapter/api/handler"
	"go-challenger/adapter/api/response"
	"go-challenger/core/usecase"
	"go-challenger/core/usecase/input"
	"net/http"
)

type DeleteTaskController struct {
	uc usecase.DeleteByIdUseCaseInterface
}

var _ Controller = (*DeleteTaskController)(nil)

func NewDeleteTaskController(uc usecase.DeleteByIdUseCaseInterface) DeleteTaskController{
	return DeleteTaskController{
		uc: uc,
	}
}

func (c  DeleteTaskController) Execute(w http.ResponseWriter, r *http.Request){
	var taskId = r.URL.Query().Get("taskId")

	i := input.TaskIdInput{Id: taskId}

	if err :=c.uc.Execute(r.Context(),i); err != nil{
		handler.HandleError(w,err)
	}

	response.NewSuccess(http.StatusNoContent,"").Send(w)

}
