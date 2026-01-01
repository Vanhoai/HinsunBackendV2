package models

import (
	"hinsun-backend/internal/domain/project"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ProjectModel struct {
	ID          uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	Name        string         `gorm:"type:varchar(100);not null;uniqueIndex"`
	Description string         `gorm:"type:varchar(500);not null"`
	Github      string         `gorm:"type:varchar(255);not null"`
	Tags        pq.StringArray `gorm:"type:text[];not null"`
	Markdown    string         `gorm:"type:text;not null"`
	CreatedAt   int64          `gorm:"autoCreateTime:nano"`
	UpdatedAt   int64          `gorm:"autoUpdateTime:nano"`
	DeletedAt   *int64         `gorm:"index"`
}

func (ProjectModel) TableName() string { return "projects" }

// =================================== CONVERTERS ===================================
func (p *ProjectModel) ToEntity() *project.ProjectEntity {
	return &project.ProjectEntity{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Github:      p.Github,
		Tags:        p.Tags,
		Markdown:    p.Markdown,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   p.DeletedAt,
	}
}

func FromProjectEntity(p *project.ProjectEntity) ProjectModel {
	return ProjectModel{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Github:      p.Github,
		Tags:        p.Tags,
		Markdown:    p.Markdown,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   p.DeletedAt,
	}
}
