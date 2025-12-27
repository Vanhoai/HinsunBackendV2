package models

import (
	"hinsun-backend/internal/domain/entities"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type ExperienceModel struct {
	ID               uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	OrderIdx         int8           `gorm:"type:int;not null"`
	Position         string         `gorm:"type:varchar(100);not null"`
	Company          string         `gorm:"type:varchar(100);not null"`
	Location         string         `gorm:"type:varchar(100);not null"`
	Technologies     pq.StringArray `gorm:"type:text[];not null"`
	Responsibilities pq.StringArray `gorm:"type:text[];not null"`
	Period           string         `gorm:"type:varchar(100);not null"`
	Extra            datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt        int64          `gorm:"autoCreateTime:nano"`
	UpdatedAt        int64          `gorm:"autoUpdateTime:nano"`
	DeletedAt        *int64         `gorm:"index"`
}

func (ExperienceModel) TableName() string { return "experiences" }

// =================================== CONVERTERS ===================================
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
		Extra:            e.Extra,
		CreatedAt:        e.CreatedAt,
		UpdatedAt:        e.UpdatedAt,
		DeletedAt:        e.DeletedAt,
	}
}

func FromExperienceEntity(e *entities.ExperienceEntity) ExperienceModel {
	Extra := datatypes.JSON{}
	if e.Extra != nil {
		Extra = datatypes.JSON(e.Extra.(datatypes.JSON))
	}

	return ExperienceModel{
		ID:               e.ID,
		OrderIdx:         e.OrderIdx,
		Position:         e.Position,
		Company:          e.Company,
		Location:         e.Location,
		Technologies:     e.Technologies,
		Responsibilities: e.Responsibilities,
		Period:           e.Period,
		Extra:            Extra,
		CreatedAt:        e.CreatedAt,
		UpdatedAt:        e.UpdatedAt,
		DeletedAt:        e.DeletedAt,
	}
}
