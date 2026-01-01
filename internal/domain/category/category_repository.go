package category

import (
	"context"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *CategoryEntity) error
	Update(ctx context.Context, category *CategoryEntity) (int, error)
	Delete(ctx context.Context, id string) (int, error)
	DeleteMany(ctx context.Context, ids []string) (int, error)
	FindByID(ctx context.Context, id string) (*CategoryEntity, error)
	FindByName(ctx context.Context, name string) (*CategoryEntity, error)
	FindAll(ctx context.Context) ([]*CategoryEntity, error)
}
