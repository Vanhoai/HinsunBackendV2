package comment

import (
	"context"
	"hinsun-backend/internal/core/failure"

	"github.com/google/uuid"
)

type CommentService interface {
	CreateComment(ctx context.Context, authorID, blogID uuid.UUID, parentID *uuid.UUID, content string) (*CommentEntity, error)
	DeleteMultipleComments(ctx context.Context, ids []string) (int, error)
	FindComments(ctx context.Context) ([]*CommentEntity, error)

	FindComment(ctx context.Context, id string) (*CommentEntity, error)
	UpdateComment(ctx context.Context, id string, content string) (*CommentEntity, error)
	DeleteComment(ctx context.Context, id string) (int, error)
}

type commentService struct {
	repository CommentRepository
}

func NewCommentService(repository CommentRepository) CommentService {
	return &commentService{
		repository: repository,
	}
}

func (s *commentService) CreateComment(
	ctx context.Context,
	authorID, blogID uuid.UUID,
	parentID *uuid.UUID,
	content string,
) (*CommentEntity, error) {
	// Create new comment entity with validation
	newComment, err := NewCommentEntity(authorID, blogID, parentID, content)
	if err != nil {
		return nil, err
	}

	// Save to repository
	err = s.repository.Create(ctx, newComment)
	if err != nil {
		return nil, err
	}

	return newComment, nil
}

func (s *commentService) DeleteMultipleComments(ctx context.Context, ids []string) (int, error) {
	return s.repository.DeleteMany(ctx, ids)
}

func (s *commentService) FindComments(ctx context.Context) ([]*CommentEntity, error) {
	return s.repository.Finds(ctx)
}

func (s *commentService) FindComment(ctx context.Context, id string) (*CommentEntity, error) {
	return s.repository.Find(ctx, id)
}

func (s *commentService) UpdateComment(ctx context.Context, id string, content string) (*CommentEntity, error) {
	// Find existing comment
	existingComment, err := s.repository.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	if existingComment == nil {
		return nil, failure.NewNotFoundFailure("Comment with the given ID does not exist")
	}

	// Update comment
	err = existingComment.Update(content)
	if err != nil {
		return nil, err
	}

	// Save updated comment
	rowsAffected, err := s.repository.Update(ctx, existingComment)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, failure.NewNotFoundFailure("Comment with the given ID does not exist")
	}

	return existingComment, nil
}

func (s *commentService) DeleteComment(ctx context.Context, id string) (int, error) {
	return s.repository.Delete(ctx, id)
}
