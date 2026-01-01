package usecases

import (
	"context"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/category"
)

type CreateCategoryParams struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
}

type UpdateCategoryParams struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
}

type DeleteCategoriesQuery struct {
	IDs []string `query:"ids"`
}

type ManageCategoryUseCase interface {
	FindCategories(ctx context.Context) ([]*category.CategoryEntity, error)
	FindCategory(ctx context.Context, id string) (*category.CategoryEntity, error)
	CreateCategory(ctx context.Context, params *CreateCategoryParams) (*category.CategoryEntity, error)
	UpdateCategory(ctx context.Context, id string, params *UpdateCategoryParams) (*category.CategoryEntity, error)
	DeleteCategory(ctx context.Context, id string) (*types.DeletedResult, error)
	DeleteMultipleCategories(ctx context.Context, query *DeleteCategoriesQuery) (*types.DeletedResult, error)
}
