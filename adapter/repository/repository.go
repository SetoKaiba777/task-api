package repository

import (
	"context"
	"go-challenger/core/domain"
	"go-challenger/core/repository"
)

type TaskRepository struct {
	db TaskRepositoryDb
}

var _ repository.TaskRepository = (*TaskRepository)(nil)

func NewTaskRepository(db TaskRepositoryDb) *TaskRepository{
	return &TaskRepository{db: db}
}


func (t * TaskRepository) Save(ctx context.Context,task domain.Task) (domain.Task, error){
	tr, err := t.db.Save(ctx, task)
	
	if err != nil{
		return domain.Task{},err
	}

	return tr, nil
}

func (t * TaskRepository) Update(ctx context.Context,task domain.Task) (domain.Task, error){
	tr, err := t.db.Update(ctx, task)
	
	if err != nil{
		return domain.Task{},err
	}

	return tr, nil
}

func (t * TaskRepository) Delete(ctx context.Context,id string) error{
	if err:= t.db.Delete(ctx, id); err !=nil{
		return err
	}

	return nil
}

func (t * TaskRepository) FindById(ctx context.Context,id string) (domain.Task, error){
	tr, err := t.db.FindById(ctx,id)

	if err != nil{
		return domain.Task{},err
	}

	return tr,nil
}

func (t * TaskRepository) SaveAll(ctx context.Context,tasks []domain.Task) error{
	err := t.db.SaveAll(ctx,tasks)

	if err != nil{
		return err
	}

	return nil
}