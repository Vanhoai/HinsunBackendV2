package project

import (
	"context"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *ProjectEntity) error
	Update(ctx context.Context, project *ProjectEntity) (int, error)
	Delete(ctx context.Context, id string) (int, error)
	DeleteMany(ctx context.Context, ids []string) (int, error)
	FindByID(ctx context.Context, id string) (*ProjectEntity, error)
	FindAll(ctx context.Context) ([]*ProjectEntity, error)
	FindByName(ctx context.Context, name string) (*ProjectEntity, error)
}
