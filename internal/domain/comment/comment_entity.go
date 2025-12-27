package comment

import (
	"time"

	"github.com/google/uuid"
)

type CommentEntity struct {
	ID        uuid.UUID
	AuthorID  uuid.UUID
	BlogID    uuid.UUID
	ParentID  *uuid.UUID
	Content   string
	Favorites int64
	CreatedAt int64
	UpdatedAt int64
	DeletedAt *int64
}

func NewCommentEntity(id, authorID, blogID uuid.UUID, parentID *uuid.UUID, content string) *CommentEntity {
	now := time.Now()
	return &CommentEntity{
		ID:        id,
		AuthorID:  authorID,
		BlogID:    blogID,
		ParentID:  parentID,
		Content:   content,
		Favorites: 0,
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
		DeletedAt: nil,
	}
}
