package di

import (
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/auth"
	"hinsun-backend/internal/domain/blog"
	"hinsun-backend/internal/domain/category"
	"hinsun-backend/internal/domain/experience"
	"hinsun-backend/internal/domain/notification"
	"hinsun-backend/internal/domain/project"
	"hinsun-backend/pkg/jwt"
	"hinsun-backend/pkg/security"

	"go.uber.org/fx"
)

var ServiceModule = fx.Module("services",
	fx.Provide(
		ProvideExperienceService,
		ProvideNotificationService,
		ProvideAccountService,
		ProvideAuthService,
		ProvideBlogService,
		ProvideProjectService,
		ProvideCategoryService,
	),
)

func ProvideExperienceService(repository experience.ExperienceRepository) experience.ExperienceService {
	return experience.NewExperienceService(repository)
}

func ProvideNotificationService(repository notification.NotificationRepository) notification.NotificationService {
	return notification.NewNotificationService(repository)
}

func ProvideAccountService(repository account.AccountRepository) account.AccountService {
	return account.NewAccountService(repository)
}

func ProvideAuthService(passwordHasher security.PasswordHasher, jwtService jwt.JwtService) auth.AuthService {
	return auth.NewAuthService(passwordHasher, jwtService)
}

func ProvideBlogService(repository blog.BlogRepository) blog.BlogService {
	return blog.NewBlogService(repository)
}

func ProvideProjectService(repository project.ProjectRepository) project.ProjectService {
	return project.NewProjectService(repository)
}

func ProvideCategoryService(repository category.CategoryRepository) category.CategoryService {
	return category.NewCategoryService(repository)
}
