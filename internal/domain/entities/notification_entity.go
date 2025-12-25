package entities

import "github.com/google/uuid"

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
