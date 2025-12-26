package di

import (
	"hinsun-backend/internal/domain/applications"
	"hinsun-backend/internal/domain/services"

	"go.uber.org/fx"
)

var ApplicationModule = fx.Module("applications",
	fx.Provide(
		ProvideAuthAppService,
		ProvideGlobalAppService,
	),
)

func ProvideAuthAppService() applications.AuthAppService {
	return applications.NewAuthAppService()
}

func ProvideGlobalAppService(experienceService services.ExperienceService) applications.GlobalAppService {
	return applications.NewGlobalAppService(experienceService)
}
