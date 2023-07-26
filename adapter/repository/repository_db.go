package repository

import (
	"context"
	"go-challenger/core/domain"
)

type TaskRepositoryDb interface {
	Save(ctx context.Context, task domain.Task) (domain.Task, error)
	Update(ctx context.Context, task domain.Task) (domain.Task, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (domain.Task, error)
	SaveAll(ctx context.Context, tasks []domain.Task) error
}