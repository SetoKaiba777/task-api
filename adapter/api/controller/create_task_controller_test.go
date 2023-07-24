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

func TestCreateTaskController(t *testing.T) {
	tt := []TableTestController{
		{
			name: "Successful Task",
			input: input.TaskInput{
				Id: "1",
				Name: "Fazer altos nadas",
				Status: "Doing",
			},
			ucMock: nil,
			expectedStatus: http.StatusCreated,
			expectedError: nil,
		},
		{
			name: "Invalid Task",
			input: input.TaskInput{
				Id: "1",
				Name: "Fazer altos nadas",
				Status: "Batata",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError: domain.ErrInvalidStatus,
		},
	}

	for _,scenario := range tt {
		t.Run(scenario.name,func(t * testing.T){
			jsonBody, _ := json.Marshal(scenario.input)

			r, _ := http.NewRequest("POST", "v1/tasks", bytes.NewBuffer(jsonBody))
			w := httptest.NewRecorder()

			ucMock := &saveUseCaseMock{}
			ucMock.On("Execute",mock.Anything).Return(scenario.expectedError)

			c := NewCreateTaskController(ucMock)
			c.Execute(w,r)
		
			assert.Equal(t, scenario.expectedStatus, w.Code)
		})
	}
}