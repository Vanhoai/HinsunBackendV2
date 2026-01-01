package repositories

import (
	"context"
	"hinsun-backend/adapters/shared/models"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/project"

	"gorm.io/gorm"
)

type projectRepository struct {
	db *gorm.DB
}

// NewProjectRepository creates a new instance of ProjectRepository
func NewProjectRepository(db *gorm.DB) project.ProjectRepository {
	return &projectRepository{
		db: db,
	}
}

func (r *projectRepository) Create(ctx context.Context, project *project.ProjectEntity) error {
	model := models.FromProjectEntity(project)
	err := gorm.G[models.ProjectModel](r.db).Create(ctx, &model)
	if err != nil {
		return failure.NewDatabaseFailure("Failed to create project in database").WithCause(err)
	}

	return nil
}

func (r *projectRepository) Update(ctx context.Context, project *project.ProjectEntity) (int, error) {
	rowsAffected, err := gorm.G[models.ProjectModel](r.db).Where("id = ?", project.ID).Updates(ctx, models.FromProjectEntity(project))
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to update project in database").WithCause(err)
	}

	return rowsAffected, nil
}

func (r *projectRepository) Delete(ctx context.Context, id string) (int, error) {
	rowAffected, err := gorm.G[models.ProjectModel](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete project from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *projectRepository) DeleteMany(ctx context.Context, ids []string) (int, error) {
	rowAffected, err := gorm.G[models.ProjectModel](r.db).Where("id IN ?", ids).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete projects from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *projectRepository) FindByID(ctx context.Context, id string) (*project.ProjectEntity, error) {
	project, err := gorm.G[models.ProjectModel](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, failure.NewDatabaseFailure("Failed to retrieve project from database").WithCause(err)
	}

	return project.ToEntity(), nil
}

func (r *projectRepository) FindAll(ctx context.Context) ([]*project.ProjectEntity, error) {
	projects, err := gorm.G[models.ProjectModel](r.db).Find(ctx)
	if err != nil {
		return nil, failure.NewDatabaseFailure("Failed to retrieve projects from database").WithCause(err)
	}

	var projectEntities []*project.ProjectEntity
	for _, projectEntity := range projects {
		projectEntities = append(projectEntities, projectEntity.ToEntity())
	}

	return projectEntities, nil
}

func (r *projectRepository) FindByName(ctx context.Context, name string) (*project.ProjectEntity, error) {
	project, err := gorm.G[models.ProjectModel](r.db).Where("name = ?", name).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, failure.NewDatabaseFailure("Failed to retrieve project by name from database").WithCause(err)
	}

	return project.ToEntity(), nil
}
