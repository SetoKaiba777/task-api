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

func TestDeleteTaskController(t *testing.T){
	tt := []TableTestController{
		{
			name: "successful delete task",
			input: "v1/tasks?taskId=1",
			expectedStatus: http.StatusNoContent,
			expectedError: nil,
		},
		{
			name: "invalid task id",
			input: "v1/tasks?taskId=1",
			expectedStatus: http.StatusNotFound,
			expectedError: domain.ErrNotFoundTask,
		},
		{
			name: "server error",
			input: "v1/tasks?taskId=1",
			expectedStatus: http.StatusInternalServerError,
			expectedError: errGeneric,
		},
	}

	for _,scenario := range tt {
		t.Run(scenario.name,func(t * testing.T){
			r, _ := http.NewRequest("DELETE", scenario.input.(string), &bytes.Reader{})
			w := httptest.NewRecorder()

			ucMock := &deleteByIdUseCaseMock{}
			ucMock.On("Execute",mock.Anything).Return(scenario.expectedError)

			c := NewDeleteTaskController(ucMock)
			c.Execute(w,r)
		
			assert.Equal(t, scenario.expectedStatus, w.Code)
		})
	}
}