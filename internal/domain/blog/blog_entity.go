package blog

import (
	"fmt"
	"hinsun-backend/internal/core/failure"
	"time"

	"github.com/google/uuid"
)

const (
	MaxBlogNameLength = 100
	MaxBlogDescLength = 300
	MaxCategories     = 5
)

type BlogEntity struct {
	ID                       uuid.UUID
	AuthorID                 uuid.UUID
	Categories               []string
	Name                     string
	Description              string
	IsPublished              bool
	Markdown                 string
	Favorites                int64
	Views                    int64
	EstimatedReadTimeSeconds int64
	CreatedAt                int64
	UpdatedAt                int64
	DeletedAt                *int64
}

func NewBlogEntity(id, authorID uuid.UUID, categories []string, name, description, markdown string, isPublished bool, estimatedReadTimeSeconds int64) *BlogEntity {
	now := time.Now()
	return &BlogEntity{
		ID:                       id,
		AuthorID:                 authorID,
		Categories:               categories,
		Name:                     name,
		Description:              description,
		IsPublished:              isPublished,
		Markdown:                 markdown,
		Favorites:                0,
		Views:                    0,
		EstimatedReadTimeSeconds: estimatedReadTimeSeconds,
		CreatedAt:                now.Unix(),
		UpdatedAt:                now.Unix(),
		DeletedAt:                nil,
	}
}

func ValidateBlogName(name string) error {
	if len(name) > MaxBlogNameLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("blog name exceeds maximum of %d characters", MaxBlogNameLength),
		)
	}

	return nil
}

func ValidateBlogDescription(description string) error {
	if len(description) > MaxBlogDescLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("blog description exceeds maximum of %d characters", MaxBlogDescLength),
		)
	}

	return nil
}

func ValidateBlogCategories(categories []string) error {
	if len(categories) == 0 {
		return failure.NewValidationFailure("at least one category is required")
	}

	if len(categories) > MaxCategories {
		return failure.NewValidationFailure(
			fmt.Sprintf("number of categories exceeds maximum of %d", MaxCategories),
		)
	}

	return nil
}
