package comment

import (
	"context"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *CommentEntity) error
	Update(ctx context.Context, comment *CommentEntity) (int, error)
	Delete(ctx context.Context, id string) (int, error)
	DeleteMany(ctx context.Context, ids []string) (int, error)
	Find(ctx context.Context, id string) (*CommentEntity, error)
	Finds(ctx context.Context) ([]*CommentEntity, error)
	FindByBlogID(ctx context.Context, blogId string) ([]*CommentEntity, error)
}
