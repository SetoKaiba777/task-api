package controller

import (
	"bytes"
	"encoding/json"
	"go-challenger/core/domain"
	"go-challenger/core/usecase/input"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAllTaskController(t *testing.T) {
	tt := []TableTestController{
		{
			name: "successful Task",
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
			expectedStatus: http.StatusCreated,
			expectedError: nil,
		},
		{
			name: "invalid Task",
			input: []input.TaskInput{
				{
					Name: "Fazer altos nadas",
					Status: "Doing",
				},
				{
					Name: "Fazer altos nadas 2- Pesquisa científica",
					Status: "Batata",
				},
			},
			expectedStatus: http.StatusBadRequest,
			expectedError: domain.ErrInvalidStatus,
		},
		{
			name: "server error",
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
			expectedStatus: http.StatusInternalServerError,
			expectedError: errGeneric,
		},
	}

	for _,scenario := range tt {
		t.Run(scenario.name,func(t * testing.T){
			jsonBody, _ := json.Marshal(scenario.input)

			r, _ := http.NewRequest("POST", "v1/tasks/all", bytes.NewBuffer(jsonBody))
			w := httptest.NewRecorder()

			ucMock := &saveAllUseCaseMock{}
			ucMock.On("Execute",mock.Anything).Return(scenario.expectedError)

			c := NewCreateAllTaskController(ucMock)
			c.Execute(w,r)
		
			assert.Equal(t, scenario.expectedStatus, w.Code)
		})
	}
}