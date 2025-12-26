package usecases

import (
	"context"
	"hinsun-backend/internal/domain/entities"
)

type CreateExperienceParams struct {
	OrderIdx         int8     `json:"orderIdx" validate:"required,min=0,max=100"`
	Position         string   `json:"position" validate:"required,min=2,max=100"`
	Company          string   `json:"company" validate:"required,min=2,max=100"`
	Location         string   `json:"location" validate:"required,min=2,max=100"`
	Period           string   `json:"period" validate:"required,min=2,max=100"`
	Technologies     []string `json:"technologies" validate:"required,dive,min=1,max=50"`
	Responsibilities []string `json:"responsibilities" validate:"required,dive,min=5,max=500"`
}

type UpdateExperienceParams struct {
	OrderIdx         int64
	Position         string
	Company          string
	Location         string
	Period           string
	Technologies     []string
	Responsibilities []string
}

type ManageExperienceUseCase interface {
	FindExperience(ctx context.Context, id string) (*entities.ExperienceEntity, error)
	FindExperiences(ctx context.Context) ([]*entities.ExperienceEntity, error)
	CreateExperience(ctx context.Context, params *CreateExperienceParams) (*entities.ExperienceEntity, error)
	UpdateExperience(ctx context.Context, id string, params *UpdateExperienceParams) (*entities.ExperienceEntity, error)
	DeleteExperience(ctx context.Context, id string) (*int, error)
	DeleteMultipleExperiences(ctx context.Context, ids []string) (*int, error)
}
