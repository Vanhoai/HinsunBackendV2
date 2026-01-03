package applications

import (
	"context"
	"fmt"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/blog"
	"hinsun-backend/internal/domain/comment"
	"hinsun-backend/internal/domain/usecases"
)

type CommentAppService interface {
	usecases.ManageCommentUseCase
}

type commentAppService struct {
	commentService comment.CommentService
	blogService    blog.BlogService
	accountService account.AccountService
}

func NewCommentAppService(
	commentService comment.CommentService,
	blogService blog.BlogService,
	accountService account.AccountService,
) CommentAppService {
	return &commentAppService{
		commentService: commentService,
		blogService:    blogService,
		accountService: accountService,
	}
}

// ================================== ManageCommentUseCase =================================
func (s *commentAppService) CreateComment(ctx context.Context, params *usecases.CreateCommentParams) (*comment.CommentEntity, error) {
	// Validate author exists
	author, err := s.accountService.FindAccountByID(ctx, params.AuthorID.String())
	if err != nil {
		return nil, err
	}

	if author == nil {
		return nil, failure.NewNotFoundFailure(fmt.Sprintf("author with ID %s not found", params.AuthorID.String()))
	}

	// Validate blog exists
	blog, err := s.blogService.FindBlog(ctx, params.BlogID.String())
	if err != nil {
		return nil, err
	}
	if blog == nil {
		return nil, failure.NewNotFoundFailure(fmt.Sprintf("blog with ID %s not found", params.BlogID.String()))
	}

	// Validate parent comment if provided
	if params.ParentID != nil {
		parentComment, err := s.commentService.FindComment(ctx, params.ParentID.String())
		if err != nil {
			return nil, err
		}

		if parentComment == nil {
			return nil, failure.NewNotFoundFailure(fmt.Sprintf("parent comment with ID %s not found", params.ParentID.String()))
		}

		// Ensure the parent comment belongs to the same blog
		if parentComment.BlogID != params.BlogID {
			return nil, failure.NewValidationFailure("parent comment does not belong to the specified blog")
		}
	}

	return s.commentService.CreateComment(ctx, params.AuthorID, params.BlogID, params.ParentID, params.Content)
}

func (s *commentAppService) UpdateComment(ctx context.Context, id string, params *usecases.UpdateCommentParams) (*comment.CommentEntity, error) {
	return s.commentService.UpdateComment(ctx, id, params.Content)
}

func (s *commentAppService) DeleteComment(ctx context.Context, id string) (*types.DeletedResult, error) {
	rowsAffected, err := s.commentService.DeleteComment(ctx, id)
	if err != nil {
		return nil, err
	}

	return &types.DeletedResult{RowsAffected: rowsAffected, Payload: id}, nil
}

func (s *commentAppService) DeleteMultipleComments(ctx context.Context, query *usecases.DeleteCommentsQuery) (*types.DeletedResult, error) {
	rowsAffected, err := s.commentService.DeleteMultipleComments(ctx, query.IDs)
	if err != nil {
		return nil, err
	}

	return &types.DeletedResult{RowsAffected: rowsAffected, Payload: query.IDs}, nil
}

func (s *commentAppService) FindComment(ctx context.Context, id string) (*comment.CommentEntity, error) {
	return s.commentService.FindComment(ctx, id)
}

func (s *commentAppService) FindComments(ctx context.Context) ([]*comment.CommentEntity, error) {
	return s.commentService.FindComments(ctx)
}

// ================================== ManageCommentUseCase =================================
