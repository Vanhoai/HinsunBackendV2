package models

import (
	"encoding/json"
	"hinsun-backend/internal/domain/blog"
	"hinsun-backend/internal/domain/values"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type BlogModel struct {
	ID                       uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	AuthorID                 uuid.UUID      `gorm:"type:uuid;not null;index"`
	Languages                pq.StringArray `gorm:"type:text[];not null"`
	Categories               pq.StringArray `gorm:"type:text[];not null"`
	Names                    datatypes.JSON `gorm:"type:jsonb;not null"` // Map of language code -> name
	Descriptions             datatypes.JSON `gorm:"type:jsonb;not null"` // Map of language code -> description
	IsPublished              bool           `gorm:"type:boolean;default:false;not null"`
	Markdowns                datatypes.JSON `gorm:"type:jsonb;not null"` // Map of language code -> markdown
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
	names := make(values.MultiLangText)
	descriptions := make(values.MultiLangText)
	markdowns := make(values.MultiLangText)

	if len(b.Names) > 0 {
		json.Unmarshal(b.Names, &names)
	}

	if len(b.Descriptions) > 0 {
		json.Unmarshal(b.Descriptions, &descriptions)
	}

	if len(b.Markdowns) > 0 {
		json.Unmarshal(b.Markdowns, &markdowns)
	}

	languages, err := values.ConvertStringArrayToMarkdownLanguageCodes(b.Languages)
	if err != nil {
		return nil
	}

	return &blog.BlogEntity{
		ID:                       b.ID,
		AuthorID:                 b.AuthorID,
		Languages:                languages,
		Categories:               b.Categories,
		Names:                    names,
		Descriptions:             descriptions,
		IsPublished:              b.IsPublished,
		Markdowns:                markdowns,
		Favorites:                b.Favorites,
		Views:                    b.Views,
		EstimatedReadTimeSeconds: b.EstimatedReadTimeSeconds,
		CreatedAt:                b.CreatedAt,
		UpdatedAt:                b.UpdatedAt,
		DeletedAt:                b.DeletedAt,
	}
}

func FromBlogEntity(b *blog.BlogEntity) BlogModel {
	namesJSON, _ := json.Marshal(b.Names)
	descsJSON, _ := json.Marshal(b.Descriptions)
	markdownsJSON, _ := json.Marshal(b.Markdowns)

	return BlogModel{
		ID:                       b.ID,
		AuthorID:                 b.AuthorID,
		Languages:                values.ConvertMarkdownLanguageCodesToStringArray(b.Languages),
		Categories:               b.Categories,
		Names:                    namesJSON,
		Descriptions:             descsJSON,
		IsPublished:              b.IsPublished,
		Markdowns:                markdownsJSON,
		Favorites:                b.Favorites,
		Views:                    b.Views,
		EstimatedReadTimeSeconds: b.EstimatedReadTimeSeconds,
		CreatedAt:                b.CreatedAt,
		UpdatedAt:                b.UpdatedAt,
		DeletedAt:                b.DeletedAt,
	}
}
