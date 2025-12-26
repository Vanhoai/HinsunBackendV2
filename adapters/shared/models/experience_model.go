package models

import (
	"hinsun-backend/internal/domain/entities"

	"github.com/google/uuid"
)

type ExperienceModel struct {
	ID               uuid.UUID `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	OrderIdx         int8      `gorm:"type:int;not null"`
	Position         string    `gorm:"type:varchar(100);not null"`
	Company          string    `gorm:"type:varchar(100);not null"`
	Location         string    `gorm:"type:varchar(100);not null"`
	Technologies     []string  `gorm:"type:text[];not null"`
	Responsibilities []string  `gorm:"type:text[];not null"`
	Period           string    `gorm:"type:varchar(100);not null"`
	CreatedAt        int64     `gorm:"autoCreateTime:nano"`
	UpdatedAt        int64     `gorm:"autoUpdateTime:nano"`
	DeletedAt        *int64    `gorm:"index"`
}

func (ExperienceModel) TableName() string { return "experiences" }

func (e *ExperienceModel) ToEntity() *entities.ExperienceEntity {
	return &entities.ExperienceEntity{
		ID:               e.ID,
		OrderIdx:         e.OrderIdx,
		Position:         e.Position,
		Company:          e.Company,
		Location:         e.Location,
		Technologies:     e.Technologies,
		Responsibilities: e.Responsibilities,
		Period:           e.Period,
		CreatedAt:        e.CreatedAt,
		UpdatedAt:        e.UpdatedAt,
		DeletedAt:        e.DeletedAt,
	}
}

func FromExperienceEntity(e *entities.ExperienceEntity) ExperienceModel {
	return ExperienceModel{
		ID:               e.ID,
		OrderIdx:         e.OrderIdx,
		Position:         e.Position,
		Company:          e.Company,
		Location:         e.Location,
		Technologies:     e.Technologies,
		Responsibilities: e.Responsibilities,
		Period:           e.Period,
		CreatedAt:        e.CreatedAt,
		UpdatedAt:        e.UpdatedAt,
		DeletedAt:        e.DeletedAt,
	}
}
