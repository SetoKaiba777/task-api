package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/repository"
	"go-challenger/core/usecase/input"
)

type (
	FindByIdUseCaseInterface interface {
		Execute(ctx context.Context, id input.TaskIdInput) (domain.Task, error)
	}
	
	findByIdUseCase struct {
		repository repository.TaskRepository
	}
)

func NewFindByIdUseCase(repository repository.TaskRepository) FindByIdUseCaseInterface{
	return &findByIdUseCase{
		repository: repository,
	}
}

func (uc * findByIdUseCase) Execute(ctx context.Context, id input.TaskIdInput) (domain.Task, error){
	task, err:= uc.repository.FindById(ctx,id.Id)
	if err!= nil {
		return domain.Task{}, err
	}
	
	return task, nil
}
