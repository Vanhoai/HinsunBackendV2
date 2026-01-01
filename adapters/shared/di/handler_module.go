package di

import (
	v1 "hinsun-backend/adapters/primary/v1"
	"hinsun-backend/adapters/primary/v1/handlers"
	v2 "hinsun-backend/adapters/primary/v2"
	"hinsun-backend/adapters/shared/middlewares"
	"hinsun-backend/internal/domain/applications"
	"hinsun-backend/pkg/jwt"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var HandlerModule = fx.Module("v1_handlers",
	fx.Provide(
		ProvideValidator,
		ProvideAuthMiddleware,
		ProvideRoleMiddleware,
		ProvideAuthHandler,
		ProvideExperienceHandler,
		ProvideBlogHandler,
		ProvideProjectHandler,
		ProvideAccountHandler,
		ProvideCategoryHandler,
	),
)

func ProvideAuthMiddleware(jwtService jwt.JwtService) *middlewares.AuthMiddleware {
	return middlewares.NewAuthMiddleware(jwtService)
}

func ProvideRoleMiddleware() *middlewares.RoleMiddleware {
	return middlewares.NewRoleMiddleware()
}

func ProvideValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func ProvideAuthHandler(
	app applications.AuthAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *handlers.AuthHandler {
	return handlers.NewAuthHandler(app, validator, authMiddleware, roleMiddleware)
}

func ProvideExperienceHandler(
	app applications.GlobalAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *handlers.ExperienceHandler {
	return handlers.NewExperienceHandler(app, validator, authMiddleware, roleMiddleware)
}

func ProvideBlogHandler(
	app applications.BlogAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *handlers.BlogHandler {
	return handlers.NewBlogHandler(app, validator, authMiddleware, roleMiddleware)
}

func ProvideProjectHandler(
	app applications.GlobalAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *handlers.ProjectHandler {
	return handlers.NewProjectHandler(app, validator, authMiddleware, roleMiddleware)
}

func ProvideAccountHandler(
	app applications.AccountAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *handlers.AccountHandler {
	return handlers.NewAccountHandler(app, validator, authMiddleware, roleMiddleware)
}

func ProvideCategoryHandler(
	app applications.CategoryAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *handlers.CategoryHandler {
	return handlers.NewCategoryHandler(app, validator, authMiddleware, roleMiddleware)
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
	blogHandler *handlers.BlogHandler,
	projectHandler *handlers.ProjectHandler,
	accountHandler *handlers.AccountHandler,
	categoryHandler *handlers.CategoryHandler,
) *v1.V1Routes {
	return v1.NewV1Routes(
		authHandler,
		experienceHandler,
		blogHandler,
		projectHandler,
		accountHandler,
		categoryHandler,
	)
}

func ProvideV2Route() *v2.V2Routes {
	return v2.NewV2Routes()
}
