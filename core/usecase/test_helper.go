package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/repository"

	"github.com/stretchr/testify/mock"
)

type (
	tableTest struct {
		name          string
		input         any
		output        any
		mockedError   error
		expectedError error
	}
	mockedRepository struct {
		mock.Mock
	}
)

var _ repository.TaskRepository = (*mockedRepository)(nil)


func (r*  mockedRepository) Save(ctx context.Context,task domain.Task) (domain.Task, error){
	args := r.Called()
	return args.Get(0).(domain.Task),args.Error(1)
}

func (r * mockedRepository) Update(ctx context.Context,task domain.Task) (domain.Task, error){
	args := r.Called()
	return args.Get(0).(domain.Task),args.Error(1)
}

func (r  * mockedRepository) Delete(ctx context.Context,id string) error{
	args := r.Called()
	return args.Error(0)
}

func (r * mockedRepository) FindById(ctx context.Context,id string) (domain.Task, error){
	args := r.Called()
	return args.Get(0).(domain.Task),args.Error(1)
}

func (r*  mockedRepository) SaveAll(ctx context.Context,task []domain.Task)  error{
	args := r.Called()
	return args.Error(0)
}
