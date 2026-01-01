package blog

import (
	"context"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/values"

	"github.com/google/uuid"
)

type BlogService interface {
	CreateBlog(
		ctx context.Context,
		authorID uuid.UUID,
		languages []values.MarkdownLanguageCode,
		categories []string,
		names, descriptions, markdowns values.MultiLangText,
		isPublished bool,
		estimatedReadTimeSeconds int64,
	) (*BlogEntity, error)

	UpdateBlog(
		ctx context.Context,
		id string,
		authorId uuid.UUID,
		languages []values.MarkdownLanguageCode,
		categories []string,
		name, description, markdown values.MultiLangText,
		isPublished bool,
		estimatedReadTimeSeconds int64,
	) (*BlogEntity, error)

	FindBlogs(ctx context.Context) ([]*BlogEntity, error)
	FindBlog(ctx context.Context, id string) (*BlogEntity, error)
	DeleteBlog(ctx context.Context, id string) (int, error)
	DeleteMultipleBlogs(ctx context.Context, ids []string) (int, error)
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
	languages []values.MarkdownLanguageCode,
	categories []string,
	names, descriptions, markdowns values.MultiLangText,
	isPublished bool,
	estimatedReadTimeSeconds int64,
) (*BlogEntity, error) {
	// Validate and create new blog entity
	newBlog, err := NewBlogEntity(
		authorID,
		languages,
		categories,
		names,
		descriptions,
		markdowns,
		isPublished,
		estimatedReadTimeSeconds,
	)

	if err != nil {
		return nil, err
	}

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
	authorID uuid.UUID,
	languages []values.MarkdownLanguageCode,
	categories []string,
	name, description, markdown values.MultiLangText,
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

func (s *blogService) FindBlogs(ctx context.Context) ([]*BlogEntity, error) {
	return s.repository.FindAll(ctx)
}

func (s *blogService) FindBlog(ctx context.Context, id string) (*BlogEntity, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *blogService) DeleteBlog(ctx context.Context, id string) (int, error) {
	return s.repository.Delete(ctx, id)
}

func (s *blogService) DeleteMultipleBlogs(ctx context.Context, ids []string) (int, error) {
	return s.repository.DeleteMany(ctx, ids)
}
