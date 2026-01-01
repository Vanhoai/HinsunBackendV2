package blog

import (
	"context"
	"hinsun-backend/internal/core/failure"

	"github.com/google/uuid"
)

type BlogService interface {
	FindAllBlogs(ctx context.Context) ([]*BlogEntity, error)
	FindBlogByID(ctx context.Context, id string) (*BlogEntity, error)
	FindBlogsByAuthorID(ctx context.Context, authorID string) ([]*BlogEntity, error)
	CreateBlog(ctx context.Context, authorID uuid.UUID, categories []string, name, description, markdown string, isPublished bool, estimatedReadTimeSeconds int64) (*BlogEntity, error)
	UpdateBlog(ctx context.Context, id string, categories []string, name, description, markdown string, isPublished bool, estimatedReadTimeSeconds int64) (*BlogEntity, error)
	DeleteBlog(ctx context.Context, id string) (int, error)
	DeleteMultipleBlogs(ctx context.Context, ids []string) (int, error)
	IncrementBlogViews(ctx context.Context, id string) error
	IncrementBlogFavorites(ctx context.Context, id string) error
	DecrementBlogFavorites(ctx context.Context, id string) error
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

// FindAllBlogs retrieves all blog entities from the repository
func (s *blogService) FindAllBlogs(ctx context.Context) ([]*BlogEntity, error) {
	return s.repository.FindAll(ctx)
}

// FindBlogByID retrieves a specific blog entity by its ID from the repository
func (s *blogService) FindBlogByID(ctx context.Context, id string) (*BlogEntity, error) {
	return s.repository.FindByID(ctx, id)
}

// FindBlogsByAuthorID retrieves all blog entities by a specific author
func (s *blogService) FindBlogsByAuthorID(ctx context.Context, authorID string) ([]*BlogEntity, error) {
	return s.repository.FindByAuthorID(ctx, authorID)
}

// CreateBlog creates a new blog entity and saves it to the repository
func (s *blogService) CreateBlog(
	ctx context.Context,
	authorID uuid.UUID,
	categories []string,
	name, description, markdown string,
	isPublished bool,
	estimatedReadTimeSeconds int64,
) (*BlogEntity, error) {
	// Validate blog name
	if err := ValidateBlogName(name); err != nil {
		return nil, err
	}

	// Validate blog description
	if err := ValidateBlogDescription(description); err != nil {
		return nil, err
	}

	// Validate blog categories
	if err := ValidateBlogCategories(categories); err != nil {
		return nil, err
	}

	// Check if a blog with the same name already exists
	existingBlog, err := s.repository.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if existingBlog != nil {
		return nil, failure.NewConflictFailure("Blog with the same name already exists")
	}

	// Create new blog entity
	newBlog := NewBlogEntity(
		uuid.New(),
		authorID,
		categories,
		name,
		description,
		markdown,
		isPublished,
		estimatedReadTimeSeconds,
	)

	// Save to repository
	err = s.repository.Create(ctx, newBlog)
	if err != nil {
		return nil, err
	}

	return newBlog, nil
}

// UpdateBlog updates an existing blog entity
func (s *blogService) UpdateBlog(
	ctx context.Context,
	id string,
	categories []string,
	name, description, markdown string,
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

	// 2. Check for name conflict if name is being changed
	if existingBlog.Name != name {
		conflictBlog, err := s.repository.FindByName(ctx, name)
		if err != nil {
			return nil, err
		}

		if conflictBlog != nil {
			return nil, failure.NewConflictFailure("Blog with the same name already exists")
		}
	}

	// 3. Update fields
	err = existingBlog.Update(
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

	// 4. Save updated blog
	rowsAffected, err := s.repository.Update(ctx, existingBlog)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, failure.NewNotFoundFailure("Blog with the given ID does not exist")
	}

	return existingBlog, nil
}

// DeleteBlog deletes a blog by its ID
func (s *blogService) DeleteBlog(ctx context.Context, id string) (int, error) {
	return s.repository.Delete(ctx, id)
}

// DeleteMultipleBlogs deletes multiple blogs by their IDs
func (s *blogService) DeleteMultipleBlogs(ctx context.Context, ids []string) (int, error) {
	return s.repository.DeleteMany(ctx, ids)
}

// IncrementBlogViews increments the view count of a blog
func (s *blogService) IncrementBlogViews(ctx context.Context, id string) error {
	blog, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if blog == nil {
		return failure.NewNotFoundFailure("Blog with the given ID does not exist")
	}

	blog.IncrementViews()

	_, err = s.repository.Update(ctx, blog)
	return err
}

// IncrementBlogFavorites increments the favorite count of a blog
func (s *blogService) IncrementBlogFavorites(ctx context.Context, id string) error {
	blog, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if blog == nil {
		return failure.NewNotFoundFailure("Blog with the given ID does not exist")
	}

	blog.IncrementFavorites()

	_, err = s.repository.Update(ctx, blog)
	return err
}

// DecrementBlogFavorites decrements the favorite count of a blog
func (s *blogService) DecrementBlogFavorites(ctx context.Context, id string) error {
	blog, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if blog == nil {
		return failure.NewNotFoundFailure("Blog with the given ID does not exist")
	}

	blog.DecrementFavorites()

	_, err = s.repository.Update(ctx, blog)
	return err
}
