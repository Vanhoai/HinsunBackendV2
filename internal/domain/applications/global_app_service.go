package applications

import (
	"context"
	"hinsun-backend/internal/domain/entities"
	"hinsun-backend/internal/domain/services"
	"hinsun-backend/internal/domain/usecases"
)

// Application Service layer orchestrates multiple domain services to fulfill use cases.
// at here, we define functions that use the params and responses defined in usecases package,
// and call the appropriate methods from the domain services.
//
// Notice: normally, responses from domain services are entities and we need to convert them to
// DTOs defined in usecases package, but for simplicity, we directly return entities here.

type GlobalAppService interface {
	usecases.ManageExperienceUseCase
}

type globalAppService struct {
	experienceService services.ExperienceService
}

// NewGlobalAppService creates a new instance of GlobalAppService
func NewGlobalAppService(experienceService services.ExperienceService) GlobalAppService {
	return &globalAppService{
		experienceService: experienceService,
	}
}

func (application *globalAppService) FindExperience(ctx context.Context, id string) (*entities.ExperienceEntity, error) {
	return application.experienceService.FindExperienceByID(ctx, id)
}

func (application *globalAppService) FindExperiences(ctx context.Context) ([]*entities.ExperienceEntity, error) {
	return application.experienceService.FindAllExperiences(ctx)
}

func (application *globalAppService) CreateExperience(ctx context.Context, params *usecases.CreateExperienceParams) (*entities.ExperienceEntity, error) {
	return application.experienceService.CreateExperience(
		ctx,
		params.OrderIdx,
		params.Position,
		params.Company,
		params.Location,
		params.Technologies,
		params.Responsibilities,
		params.Period,
	)
}

func (application *globalAppService) UpdateExperience(ctx context.Context, id string, params *usecases.UpdateExperienceParams) (*entities.ExperienceEntity, error) {
	return nil, nil
}

func (application *globalAppService) DeleteExperience(ctx context.Context, id string) (*int, error) {
	return nil, nil
}

func (application *globalAppService) DeleteMultipleExperiences(ctx context.Context, ids []string) (*int, error) {
	return nil, nil
}
