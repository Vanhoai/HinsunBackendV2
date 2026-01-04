package repositories

import (
	"context"
	"fmt"
	"hinsun-backend/adapters/shared/models"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/comment"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) comment.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(ctx context.Context, comment *comment.CommentEntity) error {
	model := models.FromCommentEntity(comment)
	err := gorm.G[models.CommentModel](r.db).Create(ctx, &model)
	if err != nil {
		return failure.NewDatabaseFailure("Failed to create comment in database").WithCause(err)
	}

	return nil
}

func (r *commentRepository) Update(ctx context.Context, comment *comment.CommentEntity) (int, error) {
	rowsAffected, err := gorm.G[models.CommentModel](r.db).Where("id = ?", comment.ID).Updates(ctx, models.FromCommentEntity(comment))
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to update comment in database").WithCause(err)
	}

	return rowsAffected, nil
}

func (r *commentRepository) Delete(ctx context.Context, id string) (int, error) {
	rowAffected, err := gorm.G[models.CommentModel](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete comment from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *commentRepository) DeleteMany(ctx context.Context, ids []string) (int, error) {
	rowAffected, err := gorm.G[models.CommentModel](r.db).Where("id IN ?", ids).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete comments from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *commentRepository) Find(ctx context.Context, id string) (*comment.CommentEntity, error) {
	comment, err := gorm.G[models.CommentModel](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, failure.NewDatabaseFailure("Failed to retrieve comment from database").WithCause(err)
	}

	return comment.ToEntity(), nil
}

func (r *commentRepository) Finds(ctx context.Context) ([]*comment.CommentEntity, error) {
	comments, err := gorm.G[models.CommentModel](r.db).Find(ctx)
	if err != nil {
		return nil, failure.NewDatabaseFailure("Failed to retrieve comments from database").WithCause(err)
	}

	var commentEntities []*comment.CommentEntity
	for _, commentModel := range comments {
		commentEntities = append(commentEntities, commentModel.ToEntity())
	}

	return commentEntities, nil
}

func (r *commentRepository) FindByBlogID(ctx context.Context, blogId string) ([]*comment.CommentEntity, error) {
	var comments []models.CommentModel
	err := r.db.Model(&models.CommentModel{}).Where("blog_id = ?", blogId).Find(&comments).Error

	if err != nil {
		return nil, failure.NewDatabaseFailure("Failed to retrieve comments by blog ID from database").WithCause(err)
	}

	fmt.Printf("Retrieved %d comments for blog ID %s\n", len(comments), blogId)

	var commentEntities []*comment.CommentEntity
	for _, commentModel := range comments {
		commentEntities = append(commentEntities, commentModel.ToEntity())
	}

	return commentEntities, nil
}
