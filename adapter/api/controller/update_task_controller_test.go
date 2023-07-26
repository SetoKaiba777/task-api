package controller

import (
	"bytes"
	"encoding/json"
	"go-challenger/adapter/api/response"
	"go-challenger/core/usecase/input"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateTaskController(t *testing.T) {
	tt := []TableTestUpdateController{
		{
			name: "successful update task",
			input: "v1/tasks?taskId=1",
			body: input.TaskInput{
				Id: "2",
				Name: "Fazer altos nadas",
				Status: "Doing",
			},
			expectedBodyReturn: "Update successfuly",
			expectedStatus: http.StatusOK,
			expectedError: nil,
		},
		{
			name: "failed update task",
			input: "v1/tasks?taskId=1",
			body: input.TaskInput{
				Id: "1",
				Name: "Fazer altos nadas",
				Status: "Doing",
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError: errGeneric,
		},
	}

	for _,scenario := range tt {
		t.Run(scenario.name,func(t * testing.T){
			jsonBody, _ := json.Marshal(scenario.body)
			r, _ := http.NewRequest("PUT", scenario.input, bytes.NewBuffer(jsonBody))
			w := httptest.NewRecorder()

			ucMock := &updateUseCaseMock{}
			ucMock.On("Execute",mock.Anything).Return(scenario.expectedError)

			c := NewUpdateTaskController(ucMock)
			c.Execute(w,r)
		
			assert.Equal(t, scenario.expectedStatus, w.Code)

			if scenario.expectedError == nil{
				expectedResponse :=response.NewSuccess(200,scenario.expectedBodyReturn)
				
				var bodyDecoded string
				json.NewDecoder(w.Body).Decode(&bodyDecoded)
				actualResponse :=response.NewSuccess(200,bodyDecoded)
				
				assert.Equal(t, expectedResponse, actualResponse)
			
			}
		})
	}
}