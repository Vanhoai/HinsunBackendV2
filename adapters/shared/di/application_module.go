package di

import (
	"hinsun-backend/internal/core/events"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/applications"
	"hinsun-backend/internal/domain/auth"
	"hinsun-backend/internal/domain/experience"
	"hinsun-backend/internal/domain/notification"

	"go.uber.org/fx"
)

var ApplicationModule = fx.Module("applications",
	fx.Provide(
		ProvideNotificationAppService,
		ProvideAsyncEventBus,
		ProvideAuthAppService,
		ProvideGlobalAppService,
	),
)

func ProvideNotificationAppService(notificationService notification.NotificationService) applications.NotificationAppService {
	return applications.NewNotificationAppService(notificationService)
}

func ProvideAuthAppService(authService auth.AuthService, accountService account.AccountService) applications.AuthAppService {
	return applications.NewAuthAppService(authService, accountService)
}

func ProvideGlobalAppService(experienceService experience.ExperienceService, asyncEventBus *events.AsyncEventBus) applications.GlobalAppService {
	return applications.NewGlobalAppService(experienceService, asyncEventBus)
}

// ProvideAsyncEventBus provides an asynchronous event bus
func ProvideAsyncEventBus(notificationAppService applications.NotificationAppService) *events.AsyncEventBus {
	eventBus := events.NewAsyncEventBus()

	// Here you can subscribe your event handlers
	// e.g., eventBus.Subscribe(yourEventHandler)

	eventBus.Subscribe(notificationAppService)
	return eventBus
}
