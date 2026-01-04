package di

import (
	"hinsun-backend/adapters/secondary/repositories"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/blog"
	"hinsun-backend/internal/domain/category"
	"hinsun-backend/internal/domain/comment"
	"hinsun-backend/internal/domain/experience"
	"hinsun-backend/internal/domain/notification"
	"hinsun-backend/internal/domain/project"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var RepositoryModule = fx.Module("repositories",
	fx.Provide(
		ProvideExperienceRepository,
		ProvideNotificationRepository,
		ProvideAccountRepository,
		ProvideBlogRepository,
		ProvideProjectRepository,
		ProvideCategoryRepository,
		ProvideCommentRepository,
	),
)

func ProvideExperienceRepository(db *gorm.DB) experience.ExperienceRepository {
	return repositories.NewExperienceRepository(db)
}

func ProvideNotificationRepository(db *gorm.DB) notification.NotificationRepository {
	return repositories.NewNotificationRepository(db)
}

func ProvideAccountRepository(db *gorm.DB) account.AccountRepository {
	return repositories.NewAccountRepository(db)
}

func ProvideBlogRepository(db *gorm.DB) blog.BlogRepository {
	return repositories.NewBlogRepository(db)
}

func ProvideProjectRepository(db *gorm.DB) project.ProjectRepository {
	return repositories.NewProjectRepository(db)
}

func ProvideCategoryRepository(db *gorm.DB) category.CategoryRepository {
	return repositories.NewCategoryRepository(db)
}

func ProvideCommentRepository(db *gorm.DB) comment.CommentRepository {
	return repositories.NewCommentRepository(db)
}
