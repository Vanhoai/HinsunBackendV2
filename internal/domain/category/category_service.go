package category

import (
	"context"
	"hinsun-backend/internal/core/failure"
)

type CategoryService interface {
	FindAllCategories(ctx context.Context) ([]*CategoryEntity, error)
	FindCategoryByID(ctx context.Context, id string) (*CategoryEntity, error)
	FindCategoryByName(ctx context.Context, name string) (*CategoryEntity, error)
	CreateCategory(ctx context.Context, name string) (*CategoryEntity, error)
	UpdateCategory(ctx context.Context, id string, name string) (*CategoryEntity, error)
	DeleteCategory(ctx context.Context, id string) (int, error)
	DeleteMultipleCategories(ctx context.Context, ids []string) (int, error)
	IncrementBlogCount(ctx context.Context, categoryID string) error
	DecrementBlogCount(ctx context.Context, categoryID string) error
}

type categoryService struct {
	repository CategoryRepository
}

func NewCategoryService(repository CategoryRepository) CategoryService {
	return &categoryService{
		repository: repository,
	}
}

func (s *categoryService) FindAllCategories(ctx context.Context) ([]*CategoryEntity, error) {
	return s.repository.FindAll(ctx)
}

func (s *categoryService) FindCategoryByID(ctx context.Context, id string) (*CategoryEntity, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *categoryService) FindCategoryByName(ctx context.Context, name string) (*CategoryEntity, error) {
	return s.repository.FindByName(ctx, name)
}

func (s *categoryService) CreateCategory(ctx context.Context, name string) (*CategoryEntity, error) {
	// Check if category with the same name already exists
	existingCategory, err := s.repository.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if existingCategory != nil {
		return nil, failure.NewConflictFailure("Category with the same name already exists")
	}

	// Create new category entity
	newCategory, err := NewCategory(name)
	if err != nil {
		return nil, err
	}

	// Save to repository
	err = s.repository.Create(ctx, newCategory)
	if err != nil {
		return nil, err
	}

	return newCategory, nil
}

func (s *categoryService) UpdateCategory(ctx context.Context, id string, name string) (*CategoryEntity, error) {
	// Find existing category
	existingCategory, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if existingCategory == nil {
		return nil, failure.NewNotFoundFailure("Category with the given ID does not exist")
	}

	// Check if another category with the same name exists
	categoryWithSameName, err := s.repository.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if categoryWithSameName != nil && categoryWithSameName.ID != existingCategory.ID {
		return nil, failure.NewConflictFailure("Another category with the same name already exists")
	}

	// Update category
	err = existingCategory.Update(name)
	if err != nil {
		return nil, err
	}

	// Save to repository
	rowsAffected, err := s.repository.Update(ctx, existingCategory)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, failure.NewNotFoundFailure("Category with the given ID does not exist")
	}

	return existingCategory, nil
}

func (s *categoryService) DeleteCategory(ctx context.Context, id string) (int, error) {
	return s.repository.Delete(ctx, id)
}

func (s *categoryService) DeleteMultipleCategories(ctx context.Context, ids []string) (int, error) {
	return s.repository.DeleteMany(ctx, ids)
}

func (s *categoryService) IncrementBlogCount(ctx context.Context, categoryID string) error {
	category, err := s.repository.FindByID(ctx, categoryID)
	if err != nil {
		return err
	}

	if category == nil {
		return failure.NewNotFoundFailure("Category not found")
	}

	category.IncrementBlogCount()
	_, err = s.repository.Update(ctx, category)
	return err
}

func (s *categoryService) DecrementBlogCount(ctx context.Context, categoryID string) error {
	category, err := s.repository.FindByID(ctx, categoryID)
	if err != nil {
		return err
	}

	if category == nil {
		return failure.NewNotFoundFailure("Category not found")
	}

	category.DecrementBlogCount()
	_, err = s.repository.Update(ctx, category)
	return err
}
