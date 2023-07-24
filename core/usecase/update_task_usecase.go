package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/repository"
	"go-challenger/core/usecase/input"
)

type(
	UpdateUseCaseInterface interface {
	Execute(ctx context.Context, i input.TaskInput)  error
}
	updateUseCase struct{
		repository repository.TaskRepository
	}
)

func NewUpdateUseCase(repository repository.TaskRepository) UpdateUseCaseInterface{
	return &updateUseCase{repository: repository}
}

func (uc *updateUseCase) Execute(ctx context.Context, i input.TaskInput) error{
	task, err := domain.NewTask().
						WithId(i.Id).
						WithName(i.Name).
						WithStatus(i.Status).
						Build()
	if err != nil{
		return err
	}

	_, err = uc.repository.Update(ctx,*task)
	if err != nil{
		return err
	}
	
	return nil
}