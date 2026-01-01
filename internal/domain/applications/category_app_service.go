package applications

import (
	"context"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/category"
	"hinsun-backend/internal/domain/usecases"
)

type CategoryAppService interface {
	usecases.ManageCategoryUseCase
}

type categoryAppService struct {
	categoryService category.CategoryService
}

func NewCategoryAppService(categoryService category.CategoryService) CategoryAppService {
	return &categoryAppService{
		categoryService: categoryService,
	}
}

func (s *categoryAppService) FindCategories(ctx context.Context) ([]*category.CategoryEntity, error) {
	return s.categoryService.FindAllCategories(ctx)
}

func (s *categoryAppService) FindCategory(ctx context.Context, id string) (*category.CategoryEntity, error) {
	category, err := s.categoryService.FindCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, failure.NewNotFoundFailure("Category with the given ID does not exist")
	}

	return category, nil
}

func (s *categoryAppService) CreateCategory(ctx context.Context, params *usecases.CreateCategoryParams) (*category.CategoryEntity, error) {
	return s.categoryService.CreateCategory(ctx, params.Name)
}

func (s *categoryAppService) UpdateCategory(ctx context.Context, id string, params *usecases.UpdateCategoryParams) (*category.CategoryEntity, error) {
	return s.categoryService.UpdateCategory(ctx, id, params.Name)
}

func (s *categoryAppService) DeleteCategory(ctx context.Context, id string) (*types.DeletedResult, error) {
	rowsAffected, err := s.categoryService.DeleteCategory(ctx, id)
	if err != nil {
		return nil, err
	}

	return &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      id,
	}, nil
}

func (s *categoryAppService) DeleteMultipleCategories(ctx context.Context, query *usecases.DeleteCategoriesQuery) (*types.DeletedResult, error) {
	rowsAffected, err := s.categoryService.DeleteMultipleCategories(ctx, query.IDs)
	if err != nil {
		return nil, err
	}

	return &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      query.IDs,
	}, nil
}
