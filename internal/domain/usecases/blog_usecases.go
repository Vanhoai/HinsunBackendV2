package usecases

import (
	"context"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/blog"

	"github.com/google/uuid"
)

type CreateBlogParams struct {
	AuthorID                 uuid.UUID `json:"authorId" validate:"required"`
	Categories               []string  `json:"categories" validate:"required,min=1,max=5,dive,min=2,max=50"`
	Name                     string    `json:"name" validate:"required,min=2,max=100"`
	Description              string    `json:"description" validate:"required,min=10,max=300"`
	Markdown                 string    `json:"markdown" validate:"required,min=10"`
	IsPublished              bool      `json:"isPublished"`
	EstimatedReadTimeSeconds int64     `json:"estimatedReadTimeSeconds" validate:"required,min=0"`
}

type UpdateBlogParams struct {
	Categories               []string `json:"categories" validate:"required,min=1,max=5,dive,min=2,max=50"`
	Name                     string   `json:"name" validate:"required,min=2,max=100"`
	Description              string   `json:"description" validate:"required,min=10,max=300"`
	Markdown                 string   `json:"markdown" validate:"required,min=10"`
	IsPublished              bool     `json:"isPublished"`
	EstimatedReadTimeSeconds int64    `json:"estimatedReadTimeSeconds" validate:"required,min=0"`
}

type DeleteBlogsQuery struct {
	IDs []string `query:"ids"`
}

type ManageBlogUseCase interface {
	FindBlog(ctx context.Context, id string) (*blog.BlogEntity, error)
	FindBlogs(ctx context.Context) ([]*blog.BlogEntity, error)
	FindBlogsByAuthor(ctx context.Context, authorID string) ([]*blog.BlogEntity, error)
	CreateBlog(ctx context.Context, params *CreateBlogParams) (*blog.BlogEntity, error)
	UpdateBlog(ctx context.Context, id string, params *UpdateBlogParams) (*blog.BlogEntity, error)
	DeleteBlog(ctx context.Context, id string) (*types.DeletedResult, error)
	DeleteMultipleBlogs(ctx context.Context, query *DeleteBlogsQuery) (*types.DeletedResult, error)
	IncrementBlogViews(ctx context.Context, id string) error
	IncrementBlogFavorites(ctx context.Context, id string) error
	DecrementBlogFavorites(ctx context.Context, id string) error
}
