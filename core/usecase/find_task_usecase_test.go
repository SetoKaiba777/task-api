package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/usecase/input"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindByIdUseCase(t *testing.T){
	tt := []tableTest{
		{
			name: "found task",
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
			name: "not found task",
			input: input.TaskIdInput{
				Id:"1",
			},
			output: domain.Task{},
			mockedError: nil,
			expectedError: domain.ErrNotFoundTask,
		},
	}
	for _, scenario := range tt{
		t.Run(scenario.name, func(t *testing.T){
			rMock := &mockedRepository{}
			rMock.On("FindById",mock.Anything).Return(scenario.output,scenario.mockedError)

			uc := NewFindByIdUseCase(rMock)
			resp, err := uc.Execute(context.TODO(),scenario.input.(input.TaskIdInput))

			assert.Equal(t,scenario.expectedError, err)
			assert.Equal(t, scenario.output.(domain.Task),resp)
		})
	}
}