package notification

import "hinsun-backend/internal/core/events"

// Event names constants
const (
	EventNotificationCreated = "notification.created"
	EventBlogPublished       = "notification.blog_published"
)

type NotificationCreatedEvent struct {
	events.BaseEvent
}

type NotificationCreatedPayload struct {
	NotificationID string `json:"notification_id"`
}

func NewNotificationCreatedEvent(payload NotificationCreatedPayload) *NotificationCreatedEvent {
	return &NotificationCreatedEvent{
		BaseEvent: events.NewBaseEvent(
			EventNotificationCreated,
			payload.NotificationID, // Aggregate ID
			payload,
		),
	}
}

type BlogPublishedEvent struct {
	events.BaseEvent
}

type BlogPublishedPayload struct {
	BlogID        string   `json:"blog_id"`
	BlogTitle     string   `json:"blog_title"`
	AuthorID      string   `json:"author_id"`
	AuthorName    string   `json:"author_name"`
	SubscriberIDs []string `json:"subscriber_ids"`
}

// NewBlogPublishedEvent creates a new blog published event
func NewBlogPublishedEvent(payload BlogPublishedPayload) *BlogPublishedEvent {
	return &BlogPublishedEvent{
		BaseEvent: events.NewBaseEvent(
			EventBlogPublished,
			payload.BlogID,
			payload,
		),
	}
}
