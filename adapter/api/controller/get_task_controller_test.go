package controller

import (
	"bytes"
	"go-challenger/core/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTaskController(t *testing.T) {
	tt := []TableTestController{
		{
			name: "successful get task",
			input: "v1/tasks?taskId=1",
			ucMock: domain.Task{
				Id: "1",
				Name: "Fazer altos nadas",
				Status: "Doing",
			},
			expectedStatus: http.StatusOK,
			expectedError: nil,
		},
		{
			name: "invalid task id",
			input: "v1/tasks?taskId=1",
			ucMock: domain.Task{},
			expectedStatus: http.StatusNotFound,
			expectedError: domain.ErrNotFoundTask,
		},
		{
			name: "server error",
			input: "v1/tasks?taskId=1",
			ucMock: domain.Task{},
			expectedStatus: http.StatusInternalServerError,
			expectedError: errGeneric,
		},
	}

	for _,scenario := range tt {
		t.Run(scenario.name,func(t * testing.T){
			r, _ := http.NewRequest("GET", scenario.input.(string), &bytes.Reader{})
			w := httptest.NewRecorder()

			ucMock := &findByIdUseCaseMock{}
			ucMock.On("Execute",mock.Anything).Return(scenario.ucMock,scenario.expectedError)

			c := NewGetTaskController(ucMock)
			c.Execute(w,r)
			
			assert.Equal(t, scenario.expectedStatus, w.Code)
		})
	}
}