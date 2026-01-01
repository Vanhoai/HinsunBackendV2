package models

import (
	"encoding/json"
	"hinsun-backend/internal/domain/blog"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type BlogModel struct {
	ID                       uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	AuthorID                 uuid.UUID      `gorm:"type:uuid;not null;index"`
	Languages                pq.StringArray `gorm:"type:text[];not null"`
	Categories               pq.StringArray `gorm:"type:text[];not null"`
	Name                     datatypes.JSON `gorm:"type:jsonb;not null"` // Map of language code -> name
	Description              datatypes.JSON `gorm:"type:jsonb;not null"` // Map of language code -> description
	IsPublished              bool           `gorm:"type:boolean;default:false;not null"`
	Markdown                 datatypes.JSON `gorm:"type:jsonb;not null"` // Map of language code -> markdown
	Favorites                int64          `gorm:"type:bigint;default:0;not null"`
	Views                    int64          `gorm:"type:bigint;default:0;not null"`
	EstimatedReadTimeSeconds int64          `gorm:"type:bigint;not null"`
	CreatedAt                int64          `gorm:"autoCreateTime:nano"`
	UpdatedAt                int64          `gorm:"autoUpdateTime:nano"`
	DeletedAt                *int64         `gorm:"index"`

	// Relationship: Many Blogs belong to One Account
	Author *AccountModel `gorm:"foreignKey:AuthorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (BlogModel) TableName() string { return "blogs" }

func (b *BlogModel) ToEntity() *blog.BlogEntity {
	// Convert JSON to MultiLangText maps
	name := make(blog.MultiLangText)
	description := make(blog.MultiLangText)
	markdown := make(blog.MultiLangText)

	if len(b.Name) > 0 {
		json.Unmarshal(b.Name, &name)
	}
	if len(b.Description) > 0 {
		json.Unmarshal(b.Description, &description)
	}
	if len(b.Markdown) > 0 {
		json.Unmarshal(b.Markdown, &markdown)
	}

	return &blog.BlogEntity{
		ID:                       b.ID,
		AuthorID:                 b.AuthorID,
		Languages:                b.Languages,
		Categories:               b.Categories,
		Name:                     name,
		Description:              description,
		IsPublished:              b.IsPublished,
		Markdown:                 markdown,
		Favorites:                b.Favorites,
		Views:                    b.Views,
		EstimatedReadTimeSeconds: b.EstimatedReadTimeSeconds,
		CreatedAt:                b.CreatedAt,
		UpdatedAt:                b.UpdatedAt,
		DeletedAt:                b.DeletedAt,
	}
}

func FromBlogEntity(b *blog.BlogEntity) BlogModel {
	// Convert MultiLangText maps to JSON
	nameJSON, _ := json.Marshal(b.Name)
	descJSON, _ := json.Marshal(b.Description)
	markdownJSON, _ := json.Marshal(b.Markdown)

	return BlogModel{
		ID:                       b.ID,
		AuthorID:                 b.AuthorID,
		Languages:                b.Languages,
		Categories:               b.Categories,
		Name:                     nameJSON,
		Description:              descJSON,
		IsPublished:              b.IsPublished,
		Markdown:                 markdownJSON,
		Favorites:                b.Favorites,
		Views:                    b.Views,
		EstimatedReadTimeSeconds: b.EstimatedReadTimeSeconds,
		CreatedAt:                b.CreatedAt,
		UpdatedAt:                b.UpdatedAt,
		DeletedAt:                b.DeletedAt,
	}
}
