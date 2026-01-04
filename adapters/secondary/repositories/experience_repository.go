package repositories

import (
	"context"
	"hinsun-backend/adapters/shared/models"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/experience"

	"gorm.io/gorm"
)

// Repository Layer in adapters is an implementation of the repository interface defined in the
// domain layer. It interacts with the database using GORM to perform CRUD operations.
//
// Note: In adapters, we use GORM models or any other model suitable for the database operations.
// So, we need to convert between domain entities and GORM models when implementing the methods.

type experienceRepository struct {
	db *gorm.DB
}

// NewExperienceRepository creates a new instance of ExperienceRepository
func NewExperienceRepository(db *gorm.DB) experience.ExperienceRepository {
	return &experienceRepository{
		db: db,
	}
}

func (r *experienceRepository) Create(ctx context.Context, experience *experience.ExperienceEntity) error {
	model := models.FromExperienceEntity(experience)
	err := gorm.G[models.ExperienceModel](r.db).Create(ctx, &model)
	if err != nil {
		return failure.NewDatabaseFailure("Failed to create experience in database").WithCause(err)
	}

	return nil
}

func (r *experienceRepository) Update(ctx context.Context, experience *experience.ExperienceEntity) (int, error) {
	rowsAffected, err := gorm.G[models.ExperienceModel](r.db).Where("id = ?", experience.ID).Updates(ctx, models.FromExperienceEntity(experience))
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to update experience in database").WithCause(err)
	}

	return rowsAffected, nil
}

func (r *experienceRepository) Delete(ctx context.Context, id string) (int, error) {
	rowAffected, err := gorm.G[models.ExperienceModel](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete experience from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *experienceRepository) DeleteMany(ctx context.Context, ids []string) (int, error) {
	rowAffected, err := gorm.G[models.ExperienceModel](r.db).Where("id IN ?", ids).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete experiences from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *experienceRepository) FindByID(ctx context.Context, id string) (*experience.ExperienceEntity, error) {
	experience, err := gorm.G[models.ExperienceModel](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, failure.NewDatabaseFailure("Failed to retrieve experience from database").WithCause(err)
	}

	return experience.ToEntity(), nil
}

func (r *experienceRepository) FindAll(ctx context.Context) ([]*experience.ExperienceEntity, error) {
	experiences, err := gorm.G[models.ExperienceModel](r.db).Find(ctx)
	if err != nil {
		return nil, failure.NewDatabaseFailure("Failed to retrieve experiences from database").WithCause(err)
	}

	var experienceEntities []*experience.ExperienceEntity
	for _, experienceEntity := range experiences {
		experienceEntities = append(experienceEntities, experienceEntity.ToEntity())
	}

	return experienceEntities, nil
}

func (r *experienceRepository) FindByOrderIdx(ctx context.Context, orderIdx int8) (*experience.ExperienceEntity, error) {
	experience, err := gorm.G[models.ExperienceModel](r.db).Where("order_idx = ?", orderIdx).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, failure.NewDatabaseFailure("Failed to retrieve experience by order index from database").WithCause(err)
	}

	return experience.ToEntity(), nil
}

func (r *experienceRepository) FindByCompany(ctx context.Context, company string) (*experience.ExperienceEntity, error) {
	experience, err := gorm.G[models.ExperienceModel](r.db).Where("company = ?", company).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, failure.NewDatabaseFailure("Failed to retrieve experience by company from database").WithCause(err)
	}

	return experience.ToEntity(), nil
}
