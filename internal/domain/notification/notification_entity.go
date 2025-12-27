package notification

import (
	"time"

	"github.com/google/uuid"
)

type NotificationEntity struct {
	ID        uuid.UUID
	AccountID string
	Title     string
	Content   string
	IsRead    bool
	CreatedAt int64
	UpdatedAt int64
	DeletedAt *int64
}

func NewNotification(
	AccountID string,
	Title string,
	Content string,
) (*NotificationEntity, error) {
	now := time.Now()

	return &NotificationEntity{
		ID:        uuid.New(),
		AccountID: AccountID,
		Title:     Title,
		Content:   Content,
		IsRead:    false,
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
		DeletedAt: nil,
	}, nil
}
