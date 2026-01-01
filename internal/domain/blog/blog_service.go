package blog

import (
	"context"
	"hinsun-backend/internal/core/failure"

	"github.com/google/uuid"
)

type BlogService interface {
	CreateBlog(ctx context.Context, authorID uuid.UUID, languages, categories []string, name, description, markdown MultiLangText, isPublished bool, estimatedReadTimeSeconds int64) (*BlogEntity, error)
	UpdateBlog(ctx context.Context, id string, languages, categories []string, name, description, markdown MultiLangText, isPublished bool, estimatedReadTimeSeconds int64) (*BlogEntity, error)
}

type blogService struct {
	repository BlogRepository
}

// NewBlogService creates a new instance of BlogService
func NewBlogService(repository BlogRepository) BlogService {
	return &blogService{
		repository: repository,
	}
}

func (s *blogService) CreateBlog(
	ctx context.Context,
	authorID uuid.UUID,
	languages []string,
	categories []string,
	name, description, markdown MultiLangText,
	isPublished bool,
	estimatedReadTimeSeconds int64,
) (*BlogEntity, error) {
	// Validate and create new blog entity
	newBlog, err := NewBlogEntity(
		uuid.New(),
		authorID,
		languages,
		categories,
		name,
		description,
		markdown,
		isPublished,
		estimatedReadTimeSeconds,
	)

	if err != nil {
		return nil, err
	}

	// Check for duplicate names in any language
	// (You can add more sophisticated duplicate checking if needed)

	// Save to repository
	err = s.repository.Create(ctx, newBlog)
	if err != nil {
		return nil, err
	}

	return newBlog, nil
}

func (s *blogService) UpdateBlog(
	ctx context.Context,
	id string,
	languages []string,
	categories []string,
	name, description, markdown MultiLangText,
	isPublished bool,
	estimatedReadTimeSeconds int64,
) (*BlogEntity, error) {
	// 1. Retrieve existing blog
	existingBlog, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if existingBlog == nil {
		return nil, failure.NewNotFoundFailure("Blog with the given ID does not exist")
	}

	// 2. Update fields
	err = existingBlog.Update(
		languages,
		categories,
		name,
		description,
		markdown,
		isPublished,
		estimatedReadTimeSeconds,
	)

	if err != nil {
		return nil, err
	}

	// 3. Save updated blog
	rowsAffected, err := s.repository.Update(ctx, existingBlog)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, failure.NewNotFoundFailure("Blog with the given ID does not exist")
	}

	return existingBlog, nil
}
