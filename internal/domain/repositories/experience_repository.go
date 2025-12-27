package repositories

import (
	"context"
	"hinsun-backend/internal/domain/entities"
)

type ExperienceRepository interface {
	Create(ctx context.Context, experience *entities.ExperienceEntity) error
	Update(ctx context.Context, experience *entities.ExperienceEntity) (int, error)
	Delete(ctx context.Context, id string) (int, error)
	DeleteMany(ctx context.Context, ids []string) (int, error)
	FindByID(ctx context.Context, id string) (*entities.ExperienceEntity, error)
	FindAll(ctx context.Context) ([]*entities.ExperienceEntity, error)
	FindByOrderIdx(ctx context.Context, orderIdx int8) (*entities.ExperienceEntity, error)
	FindByCompany(ctx context.Context, company string) (*entities.ExperienceEntity, error)
}
