package repositories

import (
	"context"
	"hinsun-backend/internal/domain/notification"

	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *notificationRepository {
	return &notificationRepository{
		db: db,
	}
}

// Implement NotificationRepository methods here
func (r *notificationRepository) Create(ctx context.Context, notification *notification.NotificationEntity) error {
	// Implementation goes here
	return nil
}
