package database

import (
	"context"
	"go-challenger/adapter/repository"
	"go-challenger/core/domain"
	"go-challenger/infrastructure/database/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConnection struct{
	db *gorm.DB
}

var _ repository.TaskRepositoryDb = (*MySQLConnection)(nil)

func NewMySQLConnection(dns string) (*MySQLConnection, error){
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	db.AutoMigrate(&entity.TaskEntity{})
    
	if err != nil {
        return &MySQLConnection{}, err
    }
	return &MySQLConnection{db: db}, nil
}

func (m *MySQLConnection) Save(ctx context.Context, task domain.Task) (domain.Task, error){
	result:=m.db.Create(entity.NewTaskEntity(task))
	if result.Error != nil{
		return domain.Task{}, result.Error
	}
	return task, nil		
}
func (m *MySQLConnection) Update(ctx context.Context, task domain.Task) (domain.Task, error){
	result := m.db.Where("id =?", task.Id).Updates(entity.NewTaskEntity(task))
	if result.Error != nil{
		return domain.Task{}, result.Error
	}
	return task, nil		
}
func (m *MySQLConnection) Delete(ctx context.Context, id string) error{
	result := m.db.Where("id =?",id).Delete(&entity.TaskEntity{})
	if result.Error != nil{
		return result.Error
	}

	return nil		
}

func (m *MySQLConnection) FindById(ctx context.Context, id string) (domain.Task, error){
	var task entity.TaskEntity
	result := m.db.Where("id =?",id).First(&task)
	if result.Error != nil{
		switch result.Error.Error(){
		case "record not found":
			return domain.Task{},domain.ErrNotFoundTask
		default:
			return domain.Task{}, result.Error				
		}
	}
	return task.ToDomain(), nil
}

