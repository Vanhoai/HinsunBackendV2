package project

import (
	"context"
	"hinsun-backend/internal/core/failure"

	"github.com/google/uuid"
)

type ProjectService interface {
	FindAllProjects(ctx context.Context) ([]*ProjectEntity, error)
	FindProjectByID(ctx context.Context, id string) (*ProjectEntity, error)
	CreateProject(ctx context.Context, name, description, github, cover string, tags []string, markdown string) (*ProjectEntity, error)
	UpdateProject(ctx context.Context, id string, name, description, github, cover string, tags []string, markdown string) (*ProjectEntity, error)
	DeleteProject(ctx context.Context, id string) (int, error)
	DeleteMultipleProjects(ctx context.Context, ids []string) (int, error)
}

type projectService struct {
	repository ProjectRepository
}

// NewProjectService creates a new instance of ProjectService
func NewProjectService(repository ProjectRepository) ProjectService {
	return &projectService{
		repository: repository,
	}
}

// FindAllProjects retrieves all project entities from the repository
func (s *projectService) FindAllProjects(ctx context.Context) ([]*ProjectEntity, error) {
	return s.repository.FindAll(ctx)
}

// FindProjectByID retrieves a specific project entity by its ID from the repository
func (s *projectService) FindProjectByID(ctx context.Context, id string) (*ProjectEntity, error) {
	return s.repository.FindByID(ctx, id)
}

// CreateProject creates a new project entity and saves it to the repository
func (s *projectService) CreateProject(
	ctx context.Context,
	name, description, github, cover string,
	tags []string,
	markdown string,
) (*ProjectEntity, error) {
	// Validate project name
	if err := ValidateProjectName(name); err != nil {
		return nil, err
	}

	// Validate project description
	if err := ValidateProjectDescription(description); err != nil {
		return nil, err
	}

	// Validate project tags
	if err := ValidateProjectTags(tags); err != nil {
		return nil, err
	}

	// Check if a project with the same name already exists
	existingProject, err := s.repository.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if existingProject != nil {
		return nil, failure.NewConflictFailure("Project with the same name already exists")
	}

	// Create new project entity
	newProject := NewProjectEntity(
		uuid.New(),
		name,
		description,
		github,
		cover,
		tags,
		markdown,
	)

	// Save to repository
	err = s.repository.Create(ctx, newProject)
	if err != nil {
		return nil, err
	}

	return newProject, nil
}

// UpdateProject updates an existing project entity
func (s *projectService) UpdateProject(
	ctx context.Context,
	id string,
	name, description, github, cover string,
	tags []string,
	markdown string,
) (*ProjectEntity, error) {
	// 1. Retrieve existing project
	existingProject, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if existingProject == nil {
		return nil, failure.NewNotFoundFailure("Project with the given ID does not exist")
	}

	// 2. Check for name conflict if name is being changed
	if existingProject.Name != name {
		conflictProject, err := s.repository.FindByName(ctx, name)
		if err != nil {
			return nil, err
		}

		if conflictProject != nil {
			return nil, failure.NewConflictFailure("Project with the same name already exists")
		}
	}

	// 3. Update fields
	err = existingProject.Update(
		name,
		description,
		github,
		cover,
		tags,
		markdown,
	)

	if err != nil {
		return nil, err
	}

	// 4. Save updated project
	rowsAffected, err := s.repository.Update(ctx, existingProject)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, failure.NewNotFoundFailure("Project with the given ID does not exist")
	}

	return existingProject, nil
}

// DeleteProject deletes a project by its ID
func (s *projectService) DeleteProject(ctx context.Context, id string) (int, error) {
	return s.repository.Delete(ctx, id)
}

// DeleteMultipleProjects deletes multiple projects by their IDs
func (s *projectService) DeleteMultipleProjects(ctx context.Context, ids []string) (int, error) {
	return s.repository.DeleteMany(ctx, ids)
}
