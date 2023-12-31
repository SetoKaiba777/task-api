package usecase

import (
	"context"
	"go-challenger/core/repository"
	"go-challenger/core/usecase/input"
)

type (
	DeleteByIdUseCaseInterface interface {
		Execute(ctx context.Context, id input.TaskIdInput) error
	}

	deleteByIdUseCase struct{
		repository repository.TaskRepository
	}
)

func NewDeleteByIdUseCase(repository repository.TaskRepository) DeleteByIdUseCaseInterface{
	return &deleteByIdUseCase{
		repository: repository,
	}
}

func (uc * deleteByIdUseCase) Execute(ctx context.Context,id input.TaskIdInput) error{
	_, err:= uc.repository.FindById(ctx,id.Id)
	if err!= nil {
		return err
	}

	if err := uc.repository.Delete(ctx,id.Id);err!=nil{
		return err
	}
	return nil	
}