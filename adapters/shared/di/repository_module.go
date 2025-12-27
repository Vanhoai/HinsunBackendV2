package di

import (
	"hinsun-backend/adapters/secondary/repositories"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/experience"
	"hinsun-backend/internal/domain/notification"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var RepositioryModule = fx.Module("repositories",
	fx.Provide(
		ProvideExperienceRepository,
		ProvideNotificationRepository,
		PriovideAccountRepository,
	),
)

func ProvideExperienceRepository(db *gorm.DB) experience.ExperienceRepository {
	return repositories.NewExperienceRepostory(db)
}

func ProvideNotificationRepository(db *gorm.DB) notification.NotificationRepository {
	return repositories.NewNotificationRepository(db)
}

func PriovideAccountRepository(db *gorm.DB) account.AccountRepository {
	return repositories.NewAccountRepository(db)
}
