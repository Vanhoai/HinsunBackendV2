package di

import (
	"hinsun-backend/internal/domain/repositories"
	"hinsun-backend/internal/domain/services"

	"go.uber.org/fx"
)

var ServiceModule = fx.Module("services",
	fx.Provide(
		ProvideExperienceService,
	),
)

func ProvideExperienceService(repository repositories.ExperienceRepository) services.ExperienceService {
	return services.NewExperienceService(repository)
}
