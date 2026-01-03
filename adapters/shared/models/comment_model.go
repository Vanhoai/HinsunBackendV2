package models

import (
	"hinsun-backend/internal/domain/comment"

	"github.com/google/uuid"
)

type CommentModel struct {
	ID        uuid.UUID     `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	AuthorID  uuid.UUID     `gorm:"type:uuid;not null;index"`
	BlogID    uuid.UUID     `gorm:"type:uuid;not null;index"`
	ParentID  *uuid.UUID    `gorm:"type:uuid;index"`
	Content   string        `gorm:"type:text;not null"`
	Favorites int64         `gorm:"type:bigint;not null;default:0"`
	CreatedAt int64         `gorm:"autoCreateTime"`
	UpdatedAt int64         `gorm:"autoUpdateTime"`
	DeletedAt *int64        `gorm:"index"`
	Author    *AccountModel `gorm:"foreignKey:AuthorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Blog      *BlogModel    `gorm:"foreignKey:BlogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (CommentModel) TableName() string { return "comments" }

func (c *CommentModel) ToEntity() *comment.CommentEntity {
	return &comment.CommentEntity{
		ID:        c.ID,
		AuthorID:  c.AuthorID,
		BlogID:    c.BlogID,
		ParentID:  c.ParentID,
		Content:   c.Content,
		Favorites: c.Favorites,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt,
	}
}

func FromCommentEntity(entity *comment.CommentEntity) CommentModel {
	return CommentModel{
		ID:        entity.ID,
		AuthorID:  entity.AuthorID,
		BlogID:    entity.BlogID,
		ParentID:  entity.ParentID,
		Content:   entity.Content,
		Favorites: entity.Favorites,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}
