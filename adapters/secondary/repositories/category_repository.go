package repositories

import (
	"context"
	"hinsun-backend/adapters/shared/models"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/category"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) category.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *category.CategoryEntity) error {
	model := models.FromCategoryEntity(category)
	err := gorm.G[models.CategoryModel](r.db).Create(ctx, &model)
	if err != nil {
		return failure.NewDatabaseFailure("Failed to create category in database").WithCause(err)
	}

	return nil
}

func (r *categoryRepository) Update(ctx context.Context, category *category.CategoryEntity) (int, error) {
	rowsAffected, err := gorm.G[models.CategoryModel](r.db).
		Where("id = ?", category.ID).
		Updates(ctx, models.FromCategoryEntity(category))
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to update category in database").WithCause(err)
	}

	return rowsAffected, nil
}

func (r *categoryRepository) Delete(ctx context.Context, id string) (int, error) {
	rowAffected, err := gorm.G[models.CategoryModel](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete category from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *categoryRepository) DeleteMany(ctx context.Context, ids []string) (int, error) {
	rowAffected, err := gorm.G[models.CategoryModel](r.db).Where("id IN ?", ids).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete categories from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *categoryRepository) FindByID(ctx context.Context, id string) (*category.CategoryEntity, error) {
	categoryModel, err := gorm.G[models.CategoryModel](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, failure.NewDatabaseFailure("Failed to retrieve category from database").WithCause(err)
	}

	return categoryModel.ToEntity(), nil
}

func (r *categoryRepository) FindByName(ctx context.Context, name string) (*category.CategoryEntity, error) {
	categoryModel, err := gorm.G[models.CategoryModel](r.db).Where("name = ?", name).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, failure.NewDatabaseFailure("Failed to retrieve category by name from database").WithCause(err)
	}

	return categoryModel.ToEntity(), nil
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]*category.CategoryEntity, error) {
	categories, err := gorm.G[models.CategoryModel](r.db).Find(ctx)
	if err != nil {
		return nil, failure.NewDatabaseFailure("Failed to retrieve categories from database").WithCause(err)
	}

	var categoryEntities []*category.CategoryEntity
	for _, categoryModel := range categories {
		categoryEntities = append(categoryEntities, categoryModel.ToEntity())
	}

	return categoryEntities, nil
}
