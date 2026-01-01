package usecases

import (
	"context"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/blog"

	"github.com/google/uuid"
)

type CreateBlogParams struct {
	AuthorID                 uuid.UUID          `json:"authorId" validate:"required"`
	Name                     blog.MultiLangText `json:"name" validate:"required"`
	Description              blog.MultiLangText `json:"description" validate:"required"`
	Categories               []string           `json:"categories" validate:"required,min=1,max=5,dive,min=2,max=50"`
	Languages                []string           `json:"languages" validate:"required,min=1,max=10,dive,min=2,max=3"`
	Markdown                 blog.MultiLangText `json:"markdown" validate:"required"`
	IsPublished              bool               `json:"isPublished"`
	EstimatedReadTimeSeconds int64              `json:"estimatedReadTimeSeconds" validate:"required,min=0"`
}

type UpdateBlogParams struct {
	Languages                []string           `json:"languages" validate:"required,min=1,max=10,dive,min=2,max=3"`
	Categories               []string           `json:"categories" validate:"required,min=1,max=5,dive,min=2,max=50"`
	Name                     blog.MultiLangText `json:"name" validate:"required"`
	Description              blog.MultiLangText `json:"description" validate:"required"`
	Markdown                 blog.MultiLangText `json:"markdown" validate:"required"`
	IsPublished              bool               `json:"isPublished"`
	EstimatedReadTimeSeconds int64              `json:"estimatedReadTimeSeconds" validate:"required,min=0"`
}

type DeleteBlogsQuery struct {
	IDs []string `query:"ids"`
}

type FindBlogsQuery struct {
	Categories []string `query:"categories"`
}

type ManageBlogUseCase interface {
	FindBlogs(ctx context.Context, query *FindBlogsQuery) ([]*blog.BlogEntity, error)
	CreateBlog(ctx context.Context, params *CreateBlogParams) (*blog.BlogEntity, error)
	DeleteMultipleBlogs(ctx context.Context, query *DeleteBlogsQuery) (*types.DeletedResult, error)

	FindBlog(ctx context.Context, id string) (*blog.BlogEntity, error)
	UpdateBlog(ctx context.Context, id string, params *UpdateBlogParams) (*blog.BlogEntity, error)
	DeleteBlog(ctx context.Context, id string) (*types.DeletedResult, error)
}
