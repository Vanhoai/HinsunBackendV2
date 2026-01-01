package di

import (
	"hinsun-backend/internal/core/events"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/applications"
	"hinsun-backend/internal/domain/auth"
	"hinsun-backend/internal/domain/blog"
	"hinsun-backend/internal/domain/experience"
	"hinsun-backend/internal/domain/notification"
	"hinsun-backend/internal/domain/project"

	"go.uber.org/fx"
)

var ApplicationModule = fx.Module("applications",
	fx.Provide(
		ProvideNotificationAppService,
		ProvideAsyncEventBus,
		ProvideAuthAppService,
		ProvideGlobalAppService,
		ProvideBlogAppService,
		ProvideAccountAppService,
	),
)

func ProvideNotificationAppService(notificationService notification.NotificationService) applications.NotificationAppService {
	return applications.NewNotificationAppService(notificationService)
}

func ProvideAuthAppService(authService auth.AuthService, accountService account.AccountService) applications.AuthAppService {
	return applications.NewAuthAppService(authService, accountService)
}

func ProvideGlobalAppService(
	experienceService experience.ExperienceService,
	projectService project.ProjectService,
	asyncEventBus *events.AsyncEventBus,
) applications.GlobalAppService {
	return applications.NewGlobalAppService(experienceService, projectService, asyncEventBus)
}

func ProvideBlogAppService(blogService blog.BlogService, accountService account.AccountService) applications.BlogAppService {
	return applications.NewBlogAppService(blogService, accountService)
}

func ProvideAccountAppService(accountService account.AccountService, authService auth.AuthService) applications.AccountAppService {
	return applications.NewAccountAppService(accountService, authService)
}

// ProvideAsyncEventBus provides an asynchronous event bus
func ProvideAsyncEventBus(notificationAppService applications.NotificationAppService) *events.AsyncEventBus {
	eventBus := events.NewAsyncEventBus()

	// Here you can subscribe your event handlers
	// e.g., eventBus.Subscribe(yourEventHandler)

	eventBus.Subscribe(notificationAppService)
	return eventBus
}
