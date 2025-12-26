package di

import (
	v1 "hinsun-backend/adapters/primary/v1"
	"hinsun-backend/adapters/primary/v1/handlers"
	v2 "hinsun-backend/adapters/primary/v2"
	"hinsun-backend/internal/domain/applications"

	"go.uber.org/fx"
)

var HandlerModule = fx.Module("v1_handlers",
	fx.Provide(
		ProvideAuthHandler,
		ProvideExperienceHandler,
	),
)

func ProvideAuthHandler(app applications.AuthAppService) *handlers.AuthHandler {
	return handlers.NewAuthHandler(app)
}

func ProvideExperienceHandler(app applications.GlobalAppService) *handlers.ExperienceHandler {
	return handlers.NewExperienceHandler(app)
}

var RouterVersionModule = fx.Module("routers",
	fx.Provide(
		ProvideV1Route,
		ProvideV2Route,
	),
)

func ProvideV1Route(
	authHandler *handlers.AuthHandler,
	experienceHandler *handlers.ExperienceHandler,
) *v1.V1Routes {
	return v1.NewV1Routes(authHandler, experienceHandler)
}

func ProvideV2Route() *v2.V2Routes {
	return v2.NewV2Routes()
}
