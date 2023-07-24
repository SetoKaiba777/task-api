package controller

import (
	"context"
	"errors"
	"go-challenger/core/domain"
	"go-challenger/core/usecase"
	"go-challenger/core/usecase/input"

	"github.com/stretchr/testify/mock"
)

var errGeneric = errors.New("generic error")

type (
	TableTestController struct {
		name           string
		input          any
		ucMock         any
		expectedStatus int
		expectedError  error
	}

	TableTestUpdateController struct {
		name           string
		input       string
		body		   input.TaskInput
		expectedBodyReturn   any       
		expectedStatus int
		expectedError  error
	}

	saveUseCaseMock struct {
		mock.Mock
	}

	findByIdUseCaseMock struct {
		mock.Mock
	}

	deleteByIdUseCaseMock struct {
		mock.Mock
	}

	updateUseCaseMock struct {
		mock.Mock
	}
)

var _ usecase.SaveUseCaseInterface = (*saveUseCaseMock)(nil)
var _ usecase.FindByIdUseCaseInterface = (*findByIdUseCaseMock)(nil)
var _ usecase.DeleteByIdUseCaseInterface = (*deleteByIdUseCaseMock)(nil)
var _ usecase.UpdateUseCaseInterface = (*updateUseCaseMock)(nil)

func (mSave *saveUseCaseMock) Execute(ctx context.Context, i input.TaskInput ) error{
	args:= mSave.Called()
	return args.Error(0)
}

func (mFind *findByIdUseCaseMock) Execute(ctx context.Context, id input.TaskIdInput) (domain.Task, error){
	args:= mFind.Called()
	return args.Get(0).(domain.Task),args.Error(1)
}

func (mDel *deleteByIdUseCaseMock) Execute(ctx context.Context, id input.TaskIdInput) error{
	args:= mDel.Called()
	return args.Error(0)
}

func (mUp *updateUseCaseMock) Execute(ctx context.Context, i input.TaskInput ) error{
	args:= mUp.Called()
	return args.Error(0)
}