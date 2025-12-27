package applications

import (
	"context"
	"hinsun-backend/internal/core/events"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/experience"
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
	experienceService experience.ExperienceService
	asyncEventBus     *events.AsyncEventBus
}

// NewGlobalAppService creates a new instance of GlobalAppService
func NewGlobalAppService(
	experienceService experience.ExperienceService,
	asyncEventBus *events.AsyncEventBus,
) GlobalAppService {
	return &globalAppService{
		experienceService: experienceService,
		asyncEventBus:     asyncEventBus,
	}
}

func (g *globalAppService) FindExperience(ctx context.Context, id string) (*experience.ExperienceEntity, error) {
	experience, err := g.experienceService.FindExperienceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if experience == nil {
		return nil, failure.NewNotFoundFailure("Experience with the given ID does not exist")
	}

	return experience, nil
}

func (g *globalAppService) FindExperiences(ctx context.Context) ([]*experience.ExperienceEntity, error) {
	return g.experienceService.FindAllExperiences(ctx)
}

func (g *globalAppService) CreateExperience(ctx context.Context, params *usecases.CreateExperienceParams) (*experience.ExperienceEntity, error) {
	return g.experienceService.CreateExperience(
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

func (g *globalAppService) UpdateExperience(ctx context.Context, id string, params *usecases.UpdateExperienceParams) (*experience.ExperienceEntity, error) {
	return g.experienceService.UpdateExperience(
		ctx,
		id,
		params.OrderIdx,
		params.Position,
		params.Company,
		params.Location,
		params.Technologies,
		params.Responsibilities,
		params.Period,
	)
}

func (g *globalAppService) DeleteExperience(ctx context.Context, id string) (*types.DeletedResult, error) {
	rowsAffected, err := g.experienceService.DeleteExperience(ctx, id)
	if err != nil {
		return nil, err
	}

	deletedResult := &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      id,
	}

	return deletedResult, nil
}

func (g *globalAppService) DeleteMultipleExperiences(ctx context.Context, query *usecases.DeleteExperiencesQuery) (*types.DeletedResult, error) {
	rowsAffected, err := g.experienceService.DeleteMultipleExperiences(ctx, query.IDs)
	if err != nil {
		return nil, err
	}

	deletedResult := &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      query.IDs,
	}

	return deletedResult, nil
}
