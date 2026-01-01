package applications

import (
	"context"
	"hinsun-backend/internal/core/events"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/blog"
	"hinsun-backend/internal/domain/experience"
	"hinsun-backend/internal/domain/project"
	"hinsun-backend/internal/domain/usecases"
	"hinsun-backend/internal/domain/values"
)

// Application Service layer orchestrates multiple domain services to fulfill use cases.
// at here, we define functions that use the params and responses defined in usecases package,
// and call the appropriate methods from the domain services.
//
// Notice: normally, responses from domain services are entities and we need to convert them to
// DTOs defined in usecases package, but for simplicity, we directly return entities here.

type GlobalAppService interface {
	usecases.ManageExperienceUseCase
	usecases.ManageBlogUseCase
	usecases.ManageProjectUseCase
	usecases.ManageAccountUseCase
}

type globalAppService struct {
	experienceService experience.ExperienceService
	blogService       blog.BlogService
	projectService    project.ProjectService
	accountService    account.AccountService
	asyncEventBus     *events.AsyncEventBus
}

// NewGlobalAppService creates a new instance of GlobalAppService
func NewGlobalAppService(
	experienceService experience.ExperienceService,
	blogService blog.BlogService,
	projectService project.ProjectService,
	accountService account.AccountService,
	asyncEventBus *events.AsyncEventBus,
) GlobalAppService {
	return &globalAppService{
		experienceService: experienceService,
		blogService:       blogService,
		projectService:    projectService,
		accountService:    accountService,
		asyncEventBus:     asyncEventBus,
	}
}

// ============================== EXPERIENCE USE CASES ==============================

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

// ============================== BLOG USE CASES ==============================

func (g *globalAppService) FindBlog(ctx context.Context, id string) (*blog.BlogEntity, error) {
	blog, err := g.blogService.FindBlogByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if blog == nil {
		return nil, failure.NewNotFoundFailure("Blog with the given ID does not exist")
	}

	return blog, nil
}

func (g *globalAppService) FindBlogs(ctx context.Context) ([]*blog.BlogEntity, error) {
	return g.blogService.FindAllBlogs(ctx)
}

func (g *globalAppService) FindBlogsByAuthor(ctx context.Context, authorID string) ([]*blog.BlogEntity, error) {
	return g.blogService.FindBlogsByAuthorID(ctx, authorID)
}

func (g *globalAppService) CreateBlog(ctx context.Context, params *usecases.CreateBlogParams) (*blog.BlogEntity, error) {
	return g.blogService.CreateBlog(
		ctx,
		params.AuthorID,
		params.Categories,
		params.Name,
		params.Description,
		params.Markdown,
		params.IsPublished,
		params.EstimatedReadTimeSeconds,
	)
}

func (g *globalAppService) UpdateBlog(ctx context.Context, id string, params *usecases.UpdateBlogParams) (*blog.BlogEntity, error) {
	return g.blogService.UpdateBlog(
		ctx,
		id,
		params.Categories,
		params.Name,
		params.Description,
		params.Markdown,
		params.IsPublished,
		params.EstimatedReadTimeSeconds,
	)
}

func (g *globalAppService) DeleteBlog(ctx context.Context, id string) (*types.DeletedResult, error) {
	rowsAffected, err := g.blogService.DeleteBlog(ctx, id)
	if err != nil {
		return nil, err
	}

	deletedResult := &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      id,
	}

	return deletedResult, nil
}

func (g *globalAppService) DeleteMultipleBlogs(ctx context.Context, query *usecases.DeleteBlogsQuery) (*types.DeletedResult, error) {
	rowsAffected, err := g.blogService.DeleteMultipleBlogs(ctx, query.IDs)
	if err != nil {
		return nil, err
	}

	deletedResult := &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      query.IDs,
	}

	return deletedResult, nil
}

func (g *globalAppService) IncrementBlogViews(ctx context.Context, id string) error {
	return g.blogService.IncrementBlogViews(ctx, id)
}

func (g *globalAppService) IncrementBlogFavorites(ctx context.Context, id string) error {
	return g.blogService.IncrementBlogFavorites(ctx, id)
}

func (g *globalAppService) DecrementBlogFavorites(ctx context.Context, id string) error {
	return g.blogService.DecrementBlogFavorites(ctx, id)
}

// ============================== PROJECT USE CASES ==============================

func (g *globalAppService) FindProject(ctx context.Context, id string) (*project.ProjectEntity, error) {
	project, err := g.projectService.FindProjectByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if project == nil {
		return nil, failure.NewNotFoundFailure("Project with the given ID does not exist")
	}

	return project, nil
}

func (g *globalAppService) FindProjects(ctx context.Context) ([]*project.ProjectEntity, error) {
	return g.projectService.FindAllProjects(ctx)
}

func (g *globalAppService) CreateProject(ctx context.Context, params *usecases.CreateProjectParams) (*project.ProjectEntity, error) {
	return g.projectService.CreateProject(
		ctx,
		params.Name,
		params.Description,
		params.Github,
		params.Cover,
		params.Tags,
		params.Markdown,
	)
}

func (g *globalAppService) UpdateProject(ctx context.Context, id string, params *usecases.UpdateProjectParams) (*project.ProjectEntity, error) {
	return g.projectService.UpdateProject(
		ctx,
		id,
		params.Name,
		params.Description,
		params.Github,
		params.Cover,
		params.Tags,
		params.Markdown,
	)
}

func (g *globalAppService) DeleteProject(ctx context.Context, id string) (*types.DeletedResult, error) {
	rowsAffected, err := g.projectService.DeleteProject(ctx, id)
	if err != nil {
		return nil, err
	}

	deletedResult := &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      id,
	}

	return deletedResult, nil
}

func (g *globalAppService) DeleteMultipleProjects(ctx context.Context, query *usecases.DeleteProjectsQuery) (*types.DeletedResult, error) {
	rowsAffected, err := g.projectService.DeleteMultipleProjects(ctx, query.IDs)
	if err != nil {
		return nil, err
	}

	deletedResult := &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      query.IDs,
	}

	return deletedResult, nil
}

// ============================== ACCOUNT USE CASES ==============================

func (g *globalAppService) FindAccountByEmail(ctx context.Context, email *values.Email) (*account.AccountEntity, error) {
	account, err := g.accountService.FindAccountByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, failure.NewNotFoundFailure("Account with the given email does not exist")
	}

	return account, nil
}

func (g *globalAppService) FindAccountByID(ctx context.Context, id string) (*account.AccountEntity, error) {
	// This method needs to be implemented in account service
	// For now, return a not found error as a placeholder
	return nil, failure.NewNotFoundFailure("FindAccountByID not yet implemented")
}
