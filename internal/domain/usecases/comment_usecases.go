package usecases

import (
	"context"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/comment"

	"github.com/google/uuid"
)

type CreateCommentParams struct {
	AuthorID uuid.UUID  `json:"authorId" validate:"required"`
	BlogID   uuid.UUID  `json:"blogId" validate:"required"`
	ParentID *uuid.UUID `json:"parentId,omitempty"`
	Content  string     `json:"content" validate:"required,min=1,max=1000"`
}

type UpdateCommentParams struct {
	Content string `json:"content" validate:"required,min=1,max=1000"`
}

type DeleteCommentsQuery struct {
	IDs []string `query:"ids"`
}

// ================================== ManageCommentUseCase =================================

type ManageCommentUseCase interface {
	CreateComment(ctx context.Context, params *CreateCommentParams) (*comment.CommentEntity, error)
	DeleteMultipleComments(ctx context.Context, query *DeleteCommentsQuery) (*types.DeletedResult, error)
	FindComments(ctx context.Context) ([]*comment.CommentEntity, error)

	DeleteComment(ctx context.Context, id string) (*types.DeletedResult, error)
	UpdateComment(ctx context.Context, id string, params *UpdateCommentParams) (*comment.CommentEntity, error)
	FindComment(ctx context.Context, id string) (*comment.CommentEntity, error)
}

// ================================== ManageCommentUseCase =================================
