package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type tableTest struct {
	name          string
	input         Task
	expectedError error
}

func TestBuild(t *testing.T) {
	tt := []tableTest{
		{
			name: "build successfully",
			input: Task{
				Id: "1",
				Name:   "Fazer altos nadas",
				Status: "Doing",
			},
			expectedError: nil,
		},
		{
			name: "invalid status error",
			input: Task{
				Id: "1",
				Name:   "Fazer altos nadas",
				Status: "Batata",
			},
			expectedError: ErrInvalidStatus,
		},
	}
	for _, scenario := range tt {
		t.Run(scenario.name, func(t *testing.T) {
			_, err :=scenario.input.Build()
			assert.Equal(t,scenario.expectedError,err)
		})
	}
}