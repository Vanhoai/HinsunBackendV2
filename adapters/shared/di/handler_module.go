package di

import (
	v1 "hinsun-backend/adapters/primary/v1"
	"hinsun-backend/adapters/primary/v1/handlers"
	v2 "hinsun-backend/adapters/primary/v2"
	"hinsun-backend/internal/domain/applications"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var HandlerModule = fx.Module("v1_handlers",
	fx.Provide(
		ProvideValidator,
		ProvideAuthHandler,
		ProvideExperienceHandler,
	),
)

func ProvideValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func ProvideAuthHandler(app applications.AuthAppService, validator *validator.Validate) *handlers.AuthHandler {
	return handlers.NewAuthHandler(app, validator)
}

func ProvideExperienceHandler(app applications.GlobalAppService, validator *validator.Validate) *handlers.ExperienceHandler {
	return handlers.NewExperienceHandler(app, validator)
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
