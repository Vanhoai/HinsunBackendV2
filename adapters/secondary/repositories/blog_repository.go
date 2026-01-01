package repositories

import (
	"context"
	"hinsun-backend/adapters/shared/models"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/blog"

	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

// NewBlogRepository creates a new instance of BlogRepository
func NewBlogRepository(db *gorm.DB) blog.BlogRepository {
	return &blogRepository{
		db: db,
	}
}

func (r *blogRepository) Create(ctx context.Context, blog *blog.BlogEntity) error {
	model := models.FromBlogEntity(blog)
	err := gorm.G[models.BlogModel](r.db).Create(ctx, &model)
	if err != nil {
		return failure.NewDatabaseFailure("Failed to create blog in database").WithCause(err)
	}

	return nil
}

func (r *blogRepository) Update(ctx context.Context, blog *blog.BlogEntity) (int, error) {
	rowsAffected, err := gorm.G[models.BlogModel](r.db).Where("id = ?", blog.ID).Updates(ctx, models.FromBlogEntity(blog))
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to update blog in database").WithCause(err)
	}

	return rowsAffected, nil
}

func (r *blogRepository) Delete(ctx context.Context, id string) (int, error) {
	rowAffected, err := gorm.G[models.BlogModel](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete blog from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *blogRepository) DeleteMany(ctx context.Context, ids []string) (int, error) {
	rowAffected, err := gorm.G[models.BlogModel](r.db).Where("id IN ?", ids).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete blogs from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *blogRepository) FindByID(ctx context.Context, id string) (*blog.BlogEntity, error) {
	blog, err := gorm.G[models.BlogModel](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, failure.NewDatabaseFailure("Failed to retrieve blog from database").WithCause(err)
	}

	return blog.ToEntity(), nil
}

func (r *blogRepository) FindAll(ctx context.Context) ([]*blog.BlogEntity, error) {
	blogs, err := gorm.G[models.BlogModel](r.db).Find(ctx)
	if err != nil {
		return nil, failure.NewDatabaseFailure("Failed to retrieve blogs from database").WithCause(err)
	}

	var blogEntities []*blog.BlogEntity
	for _, blogEntity := range blogs {
		blogEntities = append(blogEntities, blogEntity.ToEntity())
	}

	return blogEntities, nil
}

func (r *blogRepository) FindByAuthorID(ctx context.Context, authorID string) ([]*blog.BlogEntity, error) {
	blogs, err := gorm.G[models.BlogModel](r.db).Where("author_id = ?", authorID).Find(ctx)
	if err != nil {
		return nil, failure.NewDatabaseFailure("Failed to retrieve blogs by author from database").WithCause(err)
	}

	var blogEntities []*blog.BlogEntity
	for _, blogEntity := range blogs {
		blogEntities = append(blogEntities, blogEntity.ToEntity())
	}

	return blogEntities, nil
}

func (r *blogRepository) FindByName(ctx context.Context, name string) (*blog.BlogEntity, error) {
	blog, err := gorm.G[models.BlogModel](r.db).Where("name = ?", name).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, failure.NewDatabaseFailure("Failed to retrieve blog by name from database").WithCause(err)
	}

	return blog.ToEntity(), nil
}
