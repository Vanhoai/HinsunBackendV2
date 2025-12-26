package di

import (
	adapterRepositories "hinsun-backend/adapters/secondary/repositories"
	domainRepositories "hinsun-backend/internal/domain/repositories"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var RepositioryModule = fx.Module("repositories",
	fx.Provide(
		ProvideExperienceRepository,
	),
)

func ProvideExperienceRepository(db *gorm.DB) domainRepositories.ExperienceRepository {
	return adapterRepositories.NewExperienceRepostory(db)
}
