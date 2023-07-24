
package domain

import "errors"

var(
	 ErrInvalidStatus = errors.New("invalid status")
	 ErrNotFoundTask = errors.New("not found task")
)

type Task struct {
	Id, Name, Status string
}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) WithId(id string) *Task {
	t.Id = id
	return t
}

func (t *Task) WithName(name string) *Task {
	t.Name = name
	return t
}

func (t *Task) WithStatus(status string) *Task {
	t.Status = status
	return t
}

func (t *Task) Build() (*Task, error) {
	if err := Status(t.Status).fromString(); err != nil {
		return &Task{}, err
	}
	return t, nil
}