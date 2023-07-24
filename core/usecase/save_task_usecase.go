package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/repository"
	"go-challenger/core/usecase/input"
)

type(
	SaveUseCaseInterface interface {
	Execute(ctx context.Context, i input.TaskInput )  error
}
	saveUseCase struct{
		repository repository.TaskRepository
	}
)

func NewSaveUseCase(repository repository.TaskRepository) SaveUseCaseInterface{
	return &saveUseCase{repository: repository}
}

func (uc *saveUseCase) Execute(ctx context.Context, i input.TaskInput ) error{
	task, err := domain.NewTask().
						WithId(i.Id).
						WithName(i.Name).
						WithStatus(i.Status).
						Build()
	if err != nil{
		return err
	}

	_, err = uc.repository.Save(ctx,*task)
	if err != nil{
		return err
	}
	
	return nil
}