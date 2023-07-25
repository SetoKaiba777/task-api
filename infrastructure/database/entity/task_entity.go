package entity

import (
	"go-challenger/core/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskEntity struct {
	Id     string `gorm:"primaryKey;type:char(36)"`
	Name   string
	Status string
}

func NewTaskEntity(t domain.Task) *TaskEntity{
	return &TaskEntity{Id: t.Id, Name: t.Name, Status: t.Status}
}

func (task *TaskEntity) BeforeCreate(tx *gorm.DB) error {
	// Gera um UUID para o campo ID antes de criar o registro
	uuid := uuid.New().String()
	task.Id = uuid
	return nil
}

func (e *TaskEntity) ToDomain() domain.Task{
	return domain.Task{Id: e.Id,Name: e.Name, Status: e.Status }
}