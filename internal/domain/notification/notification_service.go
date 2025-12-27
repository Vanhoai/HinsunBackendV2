package notification

import "context"

type NotificationService interface {
	CreateNotification(ctx context.Context, accountID, title, content string) (*NotificationEntity, error)
}

type notificationService struct {
	repository NotificationRepository
}

// NewNotificationService creates a new notification service
func NewNotificationService(repository NotificationRepository) NotificationService {
	return &notificationService{
		repository: repository,
	}
}

func (s *notificationService) CreateNotification(ctx context.Context, accountID, title, content string) (*NotificationEntity, error) {
	notification, err := NewNotification(accountID, title, content)
	if err != nil {
		// return validation error
		return nil, err
	}

	err = s.repository.Create(ctx, notification)
	if err != nil {
		return nil, err
	}

	return notification, nil
}
