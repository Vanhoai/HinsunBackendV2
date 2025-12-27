package experience

import (
	"context"
)

type ExperienceRepository interface {
	Create(ctx context.Context, experience *ExperienceEntity) error
	Update(ctx context.Context, experience *ExperienceEntity) (int, error)
	Delete(ctx context.Context, id string) (int, error)
	DeleteMany(ctx context.Context, ids []string) (int, error)
	FindByID(ctx context.Context, id string) (*ExperienceEntity, error)
	FindAll(ctx context.Context) ([]*ExperienceEntity, error)
	FindByOrderIdx(ctx context.Context, orderIdx int8) (*ExperienceEntity, error)
	FindByCompany(ctx context.Context, company string) (*ExperienceEntity, error)
}
