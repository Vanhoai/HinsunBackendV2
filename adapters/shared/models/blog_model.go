package models

import (
	"hinsun-backend/internal/domain/blog"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type BlogModel struct {
	ID                       uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	AuthorID                 uuid.UUID      `gorm:"type:uuid;not null;index"`
	Categories               pq.StringArray `gorm:"type:text[];not null"`
	Name                     string         `gorm:"type:varchar(100);not null;uniqueIndex"`
	Description              string         `gorm:"type:varchar(300);not null"`
	IsPublished              bool           `gorm:"type:boolean;default:false;not null"`
	Markdown                 string         `gorm:"type:text;not null"`
	Favorites                int64          `gorm:"type:bigint;default:0;not null"`
	Views                    int64          `gorm:"type:bigint;default:0;not null"`
	EstimatedReadTimeSeconds int64          `gorm:"type:bigint;not null"`
	CreatedAt                int64          `gorm:"autoCreateTime:nano"`
	UpdatedAt                int64          `gorm:"autoUpdateTime:nano"`
	DeletedAt                *int64         `gorm:"index"`
}

func (BlogModel) TableName() string { return "blogs" }

// =================================== CONVERTERS ===================================
func (b *BlogModel) ToEntity() *blog.BlogEntity {
	return &blog.BlogEntity{
		ID:                       b.ID,
		AuthorID:                 b.AuthorID,
		Categories:               b.Categories,
		Name:                     b.Name,
		Description:              b.Description,
		IsPublished:              b.IsPublished,
		Markdown:                 b.Markdown,
		Favorites:                b.Favorites,
		Views:                    b.Views,
		EstimatedReadTimeSeconds: b.EstimatedReadTimeSeconds,
		CreatedAt:                b.CreatedAt,
		UpdatedAt:                b.UpdatedAt,
		DeletedAt:                b.DeletedAt,
	}
}

func FromBlogEntity(b *blog.BlogEntity) BlogModel {
	return BlogModel{
		ID:                       b.ID,
		AuthorID:                 b.AuthorID,
		Categories:               b.Categories,
		Name:                     b.Name,
		Description:              b.Description,
		IsPublished:              b.IsPublished,
		Markdown:                 b.Markdown,
		Favorites:                b.Favorites,
		Views:                    b.Views,
		EstimatedReadTimeSeconds: b.EstimatedReadTimeSeconds,
		CreatedAt:                b.CreatedAt,
		UpdatedAt:                b.UpdatedAt,
		DeletedAt:                b.DeletedAt,
	}
}
