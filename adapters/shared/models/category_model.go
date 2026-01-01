package models

import (
	"hinsun-backend/internal/domain/category"

	"github.com/google/uuid"
)

type CategoryModel struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	Name      string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	NumBlogs  int64     `gorm:"type:bigint;default:0;not null"`
	CreatedAt int64     `gorm:"autoCreateTime:nano"`
	UpdatedAt int64     `gorm:"autoUpdateTime:nano"`
	DeletedAt *int64    `gorm:"index"`
}

func (CategoryModel) TableName() string { return "categories" }

func (c *CategoryModel) ToEntity() *category.CategoryEntity {
	return &category.CategoryEntity{
		ID:        c.ID,
		Name:      c.Name,
		NumBlogs:  c.NumBlogs,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt,
	}
}

func FromCategoryEntity(entity *category.CategoryEntity) CategoryModel {
	return CategoryModel{
		ID:        entity.ID,
		Name:      entity.Name,
		NumBlogs:  entity.NumBlogs,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}
