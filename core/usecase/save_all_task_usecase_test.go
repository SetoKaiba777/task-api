package usecase

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/usecase/input"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveAll(t *testing.T) {
	tt := []tableTest{
		{
			name: "saved successfully",
			input: []input.TaskInput{
				{
					Name: "Fazer altos nadas",
					Status: "Doing",
				},
				{
					Name: "Fazer altos nadas 2- 'Pesquisa científica'",
					Status: "Doing",
				},
			},
			mockedError: nil,
			expectedError: nil,
		},
		{
			name: "connection error",
			input: []input.TaskInput{
				{
					Name: "Fazer altos nadas",
					Status: "Doing",
				},
				{
					Name: "Fazer altos nadas 2- 'Pesquisa científica'",
					Status: "Doing",
				},
			},
			mockedError: errGenericConnection,
			expectedError: errGenericConnection,
		},
		{
			name: "invalid status error",
			input: []input.TaskInput{
				{
					Name: "Fazer altos nadas",
					Status: "Doing",
				},
				{
					Name: "Fazer altos nadas 2- 'Pesquisa científica'",
					Status: "Banana",
				},
			},
			mockedError: nil,
			expectedError: domain.ErrInvalidStatus,
		},
	}
	for _, scenario := range tt{
		t.Run(scenario.name, func (t *testing.T)  {
			rMock := &mockedRepository{}
			rMock.On("SaveAll",mock.Anything).Return(scenario.mockedError)
			uc := NewSaveAllUseCase(rMock)
			err:= uc.Execute(context.TODO(), scenario.input.([]input.TaskInput))

			assert.Equal(t, scenario.expectedError,err)

		})
	}
}