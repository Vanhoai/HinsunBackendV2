package comment

import (
	"hinsun-backend/internal/core/failure"
	"time"

	"github.com/google/uuid"
)

const (
	MaxCommentLength = 1000
)

type CommentEntity struct {
	ID        uuid.UUID  `json:"id"`
	AuthorID  uuid.UUID  `json:"authorId"`
	BlogID    uuid.UUID  `json:"blogId"`
	ParentID  *uuid.UUID `json:"parentId"`
	Content   string     `json:"content"`
	Favorites int64      `json:"favorites"`
	CreatedAt int64      `json:"createdAt"`
	UpdatedAt int64      `json:"updatedAt"`
	DeletedAt *int64     `json:"deletedAt,omitempty"`
}

func NewCommentEntity(authorID, blogID uuid.UUID, parentID *uuid.UUID, content string) (*CommentEntity, error) {
	// Validate content
	if err := ValidateCommentContent(content); err != nil {
		return nil, err
	}

	now := time.Now()
	return &CommentEntity{
		ID:        uuid.New(),
		AuthorID:  authorID,
		BlogID:    blogID,
		ParentID:  parentID,
		Content:   content,
		Favorites: 0,
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
		DeletedAt: nil,
	}, nil
}

func (c *CommentEntity) Update(content string) error {
	// Validate content
	if err := ValidateCommentContent(content); err != nil {
		return err
	}

	c.Content = content
	c.UpdatedAt = time.Now().Unix()

	return nil
}

func ValidateCommentContent(content string) error {
	if len(content) == 0 {
		return failure.NewValidationFailure("Comment content cannot be empty")
	}

	if len(content) > MaxCommentLength {
		return failure.NewValidationFailure("Comment content exceeds maximum length")
	}

	return nil
}
