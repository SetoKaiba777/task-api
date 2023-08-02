package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/usecase/input"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUseCasae(t *testing.T) {
	tt := []tableTest{
		{
			name: "update successfully",
			input: input.TaskInput{
				Id: "1",
				Name: "Fazer altos nadas",
				Status: "Doing",
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
			input: input.TaskInput{
				Id: "1",
				Name: "Fazer altos nadas",
				Status: "Doing",
			},
			output: domain.Task{},
			mockedError: errGenericConnection,
			expectedError: errGenericConnection,
		},
	}
	for _, scenario := range tt{
		t.Run(scenario.name, func (t *testing.T)  {
			rMock := &mockedRepository{}
			rMock.On("Update", mock.Anything).Return(scenario.output,scenario.mockedError)
			uc := NewUpdateUseCase(rMock)
			err:= uc.Execute(context.TODO(), scenario.input.(input.TaskInput))
			assert.Equal(t, scenario.expectedError,err)
		})
	}
}