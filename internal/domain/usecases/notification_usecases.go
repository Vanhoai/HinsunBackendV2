package usecases

import "hinsun-backend/internal/domain/notification"

type ManageNotificationUseCase interface {
	FindNotificationsByAccountID(accountID string) ([]*notification.NotificationEntity, error)
}
