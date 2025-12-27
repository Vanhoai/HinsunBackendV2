package applications

import (
	"context"
	"fmt"
	"hinsun-backend/internal/core/events"
	"hinsun-backend/internal/domain/notification"
	"hinsun-backend/internal/domain/usecases"
)

type NotificationAppService interface {
	usecases.ManageNotificationUseCase

	// EventHandler methods
	HandleEvent(ctx context.Context, event events.Event) error
	InterestedIn(eventName string) bool
}

type notificationAppService struct {
	notificationService notification.NotificationService
	interestedEvents    map[string]bool
}

func NewNotificationAppService(notificationService notification.NotificationService) NotificationAppService {
	service := &notificationAppService{
		notificationService: notificationService,
		interestedEvents: map[string]bool{
			notification.EventBlogPublished:       true,
			notification.EventNotificationCreated: true,
		},
	}

	return service
}

func (s *notificationAppService) InterestedIn(eventName string) bool {
	return s.interestedEvents[eventName]
}

func (s *notificationAppService) HandleEvent(ctx context.Context, event events.Event) error {
	switch event.EventName() {
	case notification.EventBlogPublished:
		return nil
	case notification.EventNotificationCreated:
		return nil
	default:
		return fmt.Errorf("unhandled event: %s", event.EventName())
	}
}

func (s *notificationAppService) FindNotificationsByAccountID(accountID string) ([]*notification.NotificationEntity, error) {
	return nil, nil
}
