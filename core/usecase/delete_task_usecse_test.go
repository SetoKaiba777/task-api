package usecase

import (
	"context"
	"errors"
	"go-challenger/core/domain"
	"go-challenger/core/usecase/input"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	errGenericConnection = errors.New("genereic connection error")
)

func TestDeleteUseCase(t *testing.T) {
	tt := []tableTest{
		{
			name: "deleted user",
			input: input.TaskIdInput{
				Id:"1",
			},
			output: domain.Task{
				Id: "1",
				Name: "Fazer altos nadas",
				Status: "Doing",
			},
			mockedError: nil,
			expectedError: nil,
		},
		{
			name: "connection error",
			input: input.TaskIdInput{
				Id:"1",
			},
			output: domain.Task{},
			mockedError: errGenericConnection,
			expectedError: errGenericConnection,
		},
		{
			name: "not found user to delete",
			input: input.TaskIdInput{
				Id:"1",
			},
			output: domain.Task{},
			mockedError:nil,
			expectedError: domain.ErrNotFoundTask,
		},
	}
	for _, scenario := range tt{
		t.Run(scenario.name, func(t * testing.T){
			rMock := &mockedRepository{}
			rMock.On("FindById",mock.Anything).Return(scenario.output.(domain.Task),scenario.mockedError)
			rMock.On("Delete",mock.Anything).Return(nil)

			uc := NewDeleteByIdUseCase(rMock)
			err := uc.Execute(context.TODO(), scenario.input.(input.TaskIdInput))

			assert.Equal(t,scenario.expectedError,err)
		})
	}
}