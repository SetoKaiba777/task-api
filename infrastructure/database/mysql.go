package database

import (
	"context"
	"go-challenger/adapter/repository"
	"go-challenger/core/domain"
	"go-challenger/infrastructure/database/entity"

	"golang.org/x/sync/errgroup"
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

func (m *MySQLConnection) SaveAll(ctx context.Context,tasks []domain.Task) error{
	tx := m.db.Begin()

	eg := &errgroup.Group{}

	for _, task := range tasks{
		eg.Go(func() error{
			if err:= tx.Create(entity.NewTaskEntity(task)).Error; err != nil{
				return err
			}
			return nil
		})
	}

	if err:= eg.Wait(); err != nil{
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
