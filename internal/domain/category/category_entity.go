package category

import (
	"fmt"
	"hinsun-backend/internal/core/failure"
	"time"

	"github.com/google/uuid"
)

const (
	MaxCategoryNameLength = 50
	MinCategoryNameLength = 2
)

type CategoryEntity struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	NumBlogs  int64     `json:"numBlogs"`
	CreatedAt int64     `json:"createdAt"`
	UpdatedAt int64     `json:"updatedAt"`
	DeletedAt *int64    `json:"deletedAt,omitempty"`
}

func NewCategory(name string) (*CategoryEntity, error) {
	if err := ValidateCategoryName(name); err != nil {
		return nil, err
	}

	now := time.Now()
	return &CategoryEntity{
		ID:        uuid.New(),
		Name:      name,
		NumBlogs:  0,
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
		DeletedAt: nil,
	}, nil
}

func (c *CategoryEntity) Update(name string) error {
	if err := ValidateCategoryName(name); err != nil {
		return err
	}

	c.Name = name
	c.UpdatedAt = time.Now().Unix()
	return nil
}

func (c *CategoryEntity) IncrementBlogCount() {
	c.NumBlogs++
	c.UpdatedAt = time.Now().Unix()
}

func (c *CategoryEntity) DecrementBlogCount() {
	if c.NumBlogs > 0 {
		c.NumBlogs--
		c.UpdatedAt = time.Now().Unix()
	}
}

func ValidateCategoryName(name string) error {
	if len(name) < MinCategoryNameLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("category name must be at least %d characters", MinCategoryNameLength),
		)
	}

	if len(name) > MaxCategoryNameLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("category name exceeds maximum of %d characters", MaxCategoryNameLength),
		)
	}

	return nil
}
