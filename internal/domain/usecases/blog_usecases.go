package usecases

import (
	"context"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/blog"
	"hinsun-backend/internal/domain/comment"
	"hinsun-backend/internal/domain/values"

	"github.com/google/uuid"
)

// ================================== ManageBlogUseCase =================================

type CreateBlogParams struct {
	AuthorID                 uuid.UUID            `json:"authorId" validate:"required"`
	Names                    values.MultiLangText `json:"names"`
	Descriptions             values.MultiLangText `json:"descriptions"`
	Categories               []string             `json:"categories" validate:"required,min=1,max=5,dive,min=2,max=50"`
	Languages                []string             `json:"languages" validate:"required,min=1,max=10,dive,min=2,max=3"`
	Markdowns                values.MultiLangText `json:"markdowns"`
	IsPublished              bool                 `json:"isPublished"`
	EstimatedReadTimeSeconds int64                `json:"estimatedReadTimeSeconds" validate:"min=0"`
}

type UpdateBlogParams struct {
	AuthorID                 uuid.UUID            `json:"authorId" validate:"required"`
	Names                    values.MultiLangText `json:"names"`
	Descriptions             values.MultiLangText `json:"descriptions"`
	Categories               []string             `json:"categories" validate:"required,min=1,max=5,dive,min=2,max=50"`
	Languages                []string             `json:"languages" validate:"required,min=1,max=10,dive,min=2,max=3"`
	Markdowns                values.MultiLangText `json:"markdowns"`
	IsPublished              bool                 `json:"isPublished"`
	EstimatedReadTimeSeconds int64                `json:"estimatedReadTimeSeconds" validate:"min=0"`
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

// ================================== ManageBlogUseCase =================================

// ================================== CommentBlogUseCase =================================

type AddCommentToBlogParams struct {
	Content string `json:"content" validate:"required,min=1,max=1000"`
}

type UpdateCommentOnBlogParams struct {
	Content string `json:"content" validate:"required,min=1,max=1000"`
}

type CommentBlogUseCase interface {
	AddCommentToBlog(ctx context.Context, blogId string, accountId string, params *AddCommentToBlogParams) (*comment.CommentEntity, error)
	UpdateCommentOnBlog(ctx context.Context, blogId string, accountId string, commentId string, params *UpdateCommentOnBlogParams) (*comment.CommentEntity, error)
	DeleteCommentOnBlog(ctx context.Context, blogId string, accountId string, commentId string) (*types.DeletedResult, error)
}

// ================================== CommentBlogUseCase =================================
