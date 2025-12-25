package entities

import (
	"time"

	"github.com/google/uuid"
)

type TopicEntity struct {
	ID        uuid.UUID
	Name      string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt *int64
}

func NewTopicEntity(id uuid.UUID, name string) *TopicEntity {
	now := time.Now()
	return &TopicEntity{
		ID:        id,
		Name:      name,
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
		DeletedAt: nil,
	}
}
