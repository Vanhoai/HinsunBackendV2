package services

import (
	"context"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/entities"
	"hinsun-backend/internal/domain/repositories"
)

type ExperienceService interface {
	FindAllExperiences(ctx context.Context) ([]*entities.ExperienceEntity, error)
	FindExperienceByID(ctx context.Context, id string) (*entities.ExperienceEntity, error)
	CreateExperience(ctx context.Context, orderIdx int8, position string, company string, location string, technologies []string, responsibilities []string, period string) (*entities.ExperienceEntity, error)
}

type experienceService struct {
	repository repositories.ExperienceRepository
}

// NewExperienceService creates a new instance of ExperienceService
func NewExperienceService(repository repositories.ExperienceRepository) ExperienceService {
	return &experienceService{
		repository: repository,
	}
}

// FindAllExperiences retrieves all experience entities from the repository
func (s *experienceService) FindAllExperiences(ctx context.Context) ([]*entities.ExperienceEntity, error) {
	return s.repository.FindAll(ctx)
}

// FindExperienceByID retrieves a specific experience entity by its ID from the repository
func (s *experienceService) FindExperienceByID(ctx context.Context, id string) (*entities.ExperienceEntity, error) {
	return s.repository.FindByID(ctx, id)
}

// CreateExperience creates a new experience entity and saves it to the repository
func (s *experienceService) CreateExperience(
	ctx context.Context,
	orderIdx int8,
	position string,
	company string,
	location string,
	technologies []string,
	responsibilities []string,
	period string,
) (*entities.ExperienceEntity, error) {
	// Validate orderIdx that don't conflict with existing experiences
	existingExperience, err := s.repository.FindByOrderIdx(ctx, orderIdx)
	if err != nil {
		return nil, err
	}

	// If an experience with the same orderIdx exists, return conflict error
	if existingExperience != nil {
		return nil, failure.NewConflictFailure("Experience with the same orderIdx already exists")
	}

	// If an experience with the same company exists, return conflict error
	existingExperience, err = s.repository.FindByCompany(ctx, company)
	if err != nil {
		return nil, err
	}

	if existingExperience != nil {
		return nil, failure.NewConflictFailure("Experience with the same company already exists")
	}

	newExperience, err := entities.NewExperience(
		orderIdx,
		position,
		company,
		location,
		technologies,
		responsibilities,
		period,
	)

	if err != nil {
		// Return validation errors
		return nil, err
	}

	err = s.repository.Create(ctx, newExperience)
	if err != nil {
		// Return database errors
		return nil, err
	}

	return newExperience, nil
}
