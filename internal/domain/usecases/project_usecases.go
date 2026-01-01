package usecases

import (
	"context"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/project"
)

type CreateProjectParams struct {
	Name        string   `json:"name" validate:"required,min=2,max=100"`
	Description string   `json:"description" validate:"required,min=10,max=500"`
	Github      string   `json:"github" validate:"required,url"`
	Cover       string   `json:"cover" validate:"omitempty,url"`
	Tags        []string `json:"tags" validate:"required,min=1,max=5,dive,min=2,max=50"`
	Markdown    string   `json:"markdown" validate:"required,min=10"`
}

type UpdateProjectParams struct {
	Name        string   `json:"name" validate:"required,min=2,max=100"`
	Description string   `json:"description" validate:"required,min=10,max=500"`
	Github      string   `json:"github" validate:"required,url"`
	Cover       string   `json:"cover" validate:"omitempty,url"`
	Tags        []string `json:"tags" validate:"required,min=1,max=5,dive,min=2,max=50"`
	Markdown    string   `json:"markdown" validate:"required,min=10"`
}

type DeleteProjectsQuery struct {
	IDs []string `query:"ids"`
}

type ManageProjectUseCase interface {
	FindProject(ctx context.Context, id string) (*project.ProjectEntity, error)
	FindProjects(ctx context.Context) ([]*project.ProjectEntity, error)
	CreateProject(ctx context.Context, params *CreateProjectParams) (*project.ProjectEntity, error)
	UpdateProject(ctx context.Context, id string, params *UpdateProjectParams) (*project.ProjectEntity, error)
	DeleteProject(ctx context.Context, id string) (*types.DeletedResult, error)
	DeleteMultipleProjects(ctx context.Context, query *DeleteProjectsQuery) (*types.DeletedResult, error)
}
