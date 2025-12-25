package usecases

import (
	"context"
	"hinsun-backend/internal/domain/entities"
)

// #[validate(length(min = 1, message = "At least one technology must be provided"))]
//    pub technologies: Vec<String>,

//    #[validate(length(min = 1, message = "Position must not be empty"))]
//    pub position: String,

//    #[validate(length(min = 1, message = "At least one responsibility must be provided"))]
//    pub responsibilities: Vec<String>,

//    #[validate(length(min = 1, message = "Company must not be empty"))]
//    pub company: String,

//    #[validate(length(min = 1, message = "Location must not be empty"))]
//    pub location: String,

//    #[validate(length(min = 1, message = "Period must be not be empty"))]
//    pub period: String,

// #[validate(range(min = 0, message = "Order index must be non-negative"))]
// pub order_idx: i32,

type CreateExperienceParams struct {
	OrderIdx         int64
	Position         string
	Company          string
	Location         string
	Period           string
	Technologies     []string
	Responsibilities []string
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
	FindExperience(ctx context.Context, experienceId int64) (*entities.ExperienceEntity, error)
	FindExperiences(ctx context.Context) ([]*entities.ExperienceEntity, error)
	CreateExperience(ctx context.Context, params *CreateExperienceParams) (*entities.ExperienceEntity, error)
	UpdateExperience(ctx context.Context, experienceId int64, params *UpdateExperienceParams) (*entities.ExperienceEntity, error)
	DeleteExperience(ctx context.Context, experienceId int64) error
}
