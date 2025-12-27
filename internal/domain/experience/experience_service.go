package experience

import (
	"context"
	"hinsun-backend/internal/core/failure"
)

type ExperienceService interface {
	FindAllExperiences(ctx context.Context) ([]*ExperienceEntity, error)
	FindExperienceByID(ctx context.Context, id string) (*ExperienceEntity, error)
	CreateExperience(ctx context.Context, orderIdx int8, position, company, location string, technologies, responsibilities []string, period string) (*ExperienceEntity, error)
	UpdateExperience(ctx context.Context, id string, orderIdx int8, position, company, location string, technologies, responsibilities []string, period string) (*ExperienceEntity, error)
	DeleteExperience(ctx context.Context, id string) (int, error)
	DeleteMultipleExperiences(ctx context.Context, ids []string) (int, error)
}

type experienceService struct {
	repository ExperienceRepository
}

// NewExperienceService creates a new instance of ExperienceService
func NewExperienceService(repository ExperienceRepository) ExperienceService {
	return &experienceService{
		repository: repository,
	}
}

// FindAllExperiences retrieves all experience entities from the repository
func (s *experienceService) FindAllExperiences(ctx context.Context) ([]*ExperienceEntity, error) {
	return s.repository.FindAll(ctx)
}

// FindExperienceByID retrieves a specific experience entity by its ID from the repository
func (s *experienceService) FindExperienceByID(ctx context.Context, id string) (*ExperienceEntity, error) {
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
) (*ExperienceEntity, error) {
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

	newExperience, err := NewExperience(
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

func (s *experienceService) UpdateExperience(
	ctx context.Context,
	id string,
	orderIdx int8,
	position string,
	company string,
	location string,
	technologies []string,
	responsibilities []string,
	period string,
) (*ExperienceEntity, error) {
	// 1. Retrieve existing experience
	existingExperience, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if existingExperience == nil {
		return nil, failure.NewNotFoundFailure("Experience with the given ID does not exist")
	}

	// 2. Check for orderIdx conflict
	if existingExperience.OrderIdx != orderIdx {
		conflictExperience, err := s.repository.FindByOrderIdx(ctx, orderIdx)
		if err != nil {
			return nil, err
		}

		if conflictExperience != nil {
			return nil, failure.NewConflictFailure("Experience with the same orderIdx already exists")
		}
	}

	// 3. Check for company conflict
	if existingExperience.Company != company {
		conflictExperience, err := s.repository.FindByCompany(ctx, company)
		if err != nil {
			return nil, err
		}

		if conflictExperience != nil {
			return nil, failure.NewConflictFailure("Experience with the same company already exists")
		}
	}

	// 4. Update fields
	err = existingExperience.Update(
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

	// 5. Save updated experience
	rowsAffected, err := s.repository.Update(ctx, existingExperience)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, failure.NewNotFoundFailure("Experience with the given ID does not exist")
	}

	return existingExperience, nil
}

func (s *experienceService) DeleteExperience(ctx context.Context, id string) (int, error) {
	return s.repository.Delete(ctx, id)
}

func (s *experienceService) DeleteMultipleExperiences(ctx context.Context, ids []string) (int, error) {
	return s.repository.DeleteMany(ctx, ids)
}
