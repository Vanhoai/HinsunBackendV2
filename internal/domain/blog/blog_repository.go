package blog

import (
	"context"
)

type BlogRepository interface {
	Create(ctx context.Context, blog *BlogEntity) error
	Update(ctx context.Context, blog *BlogEntity) (int, error)
	Delete(ctx context.Context, id string) (int, error)
	DeleteMany(ctx context.Context, ids []string) (int, error)
	FindByID(ctx context.Context, id string) (*BlogEntity, error)
	FindAll(ctx context.Context) ([]*BlogEntity, error)
	FindByAuthorID(ctx context.Context, authorID string) ([]*BlogEntity, error)
	FindByName(ctx context.Context, name string) (*BlogEntity, error)
}
