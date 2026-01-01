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
	MaxLanguages      = 10
	MinLanguages      = 1
)

type BlogEntity struct {
	ID                       uuid.UUID     `json:"id"`
	AuthorID                 uuid.UUID     `json:"authorId"`
	Languages                []string      `json:"languages"` // Array of supported language codes
	Categories               []string      `json:"categories"`
	Name                     MultiLangText `json:"name"`        // Map of language code -> name
	Description              MultiLangText `json:"description"` // Map of language code -> description
	IsPublished              bool          `json:"isPublished"`
	Markdown                 MultiLangText `json:"markdown"` // Map of language code -> markdown content
	Favorites                int64         `json:"favorites"`
	Views                    int64         `json:"views"`
	EstimatedReadTimeSeconds int64         `json:"estimatedReadTimeSeconds"`
	CreatedAt                int64         `json:"createdAt"`
	UpdatedAt                int64         `json:"updatedAt"`
	DeletedAt                *int64        `json:"deletedAt,omitempty"`
}

func NewBlogEntity(
	id,
	authorID uuid.UUID,
	languages, categories []string,
	name, description, markdown MultiLangText,
	isPublished bool,
	estimatedReadTimeSeconds int64,
) (*BlogEntity, error) {
	// Validate languages
	if err := ValidateLanguages(languages); err != nil {
		return nil, err
	}

	// Validate that all languages have corresponding content
	if err := ValidateMultiLangContent(languages, name, "name"); err != nil {
		return nil, err
	}
	if err := ValidateMultiLangContent(languages, description, "description"); err != nil {
		return nil, err
	}
	if err := ValidateMultiLangContent(languages, markdown, "markdown"); err != nil {
		return nil, err
	}

	// Validate each language's content
	for lang, text := range name {
		if err := ValidateBlogName(text); err != nil {
			return nil, failure.NewValidationFailure(fmt.Sprintf("name for language '%s': %s", lang, err.Error()))
		}
	}

	for lang, text := range description {
		if err := ValidateBlogDescription(text); err != nil {
			return nil, failure.NewValidationFailure(fmt.Sprintf("description for language '%s': %s", lang, err.Error()))
		}
	}

	if err := ValidateBlogCategories(categories); err != nil {
		return nil, err
	}

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
	}, nil
}

func ValidateLanguages(languages []string) error {
	if len(languages) < MinLanguages {
		return failure.NewValidationFailure(
			fmt.Sprintf("at least %d language is required", MinLanguages),
		)
	}

	if len(languages) > MaxLanguages {
		return failure.NewValidationFailure(
			fmt.Sprintf("number of languages exceeds maximum of %d", MaxLanguages),
		)
	}

	// Check for duplicates
	seen := make(map[string]bool)
	for _, lang := range languages {
		if seen[lang] {
			return failure.NewValidationFailure(fmt.Sprintf("duplicate language code: %s", lang))
		}
		seen[lang] = true

		// Validate language code format (should be 2-3 lowercase letters)
		if len(lang) < 2 || len(lang) > 3 {
			return failure.NewValidationFailure(
				fmt.Sprintf("invalid language code '%s': must be 2-3 characters", lang),
			)
		}
	}

	return nil
}

func ValidateMultiLangContent(languages []string, content MultiLangText, fieldName string) error {
	if content == nil {
		return failure.NewValidationFailure(fmt.Sprintf("%s cannot be nil", fieldName))
	}

	// Check that all languages have content
	for _, lang := range languages {
		if text, exists := content[lang]; !exists || text == "" {
			return failure.NewValidationFailure(
				fmt.Sprintf("%s is missing content for language: %s", fieldName, lang),
			)
		}
	}

	return nil
}

func ValidateBlogName(name string) error {
	if len(name) == 0 {
		return failure.NewValidationFailure("blog name cannot be empty")
	}

	if len(name) > MaxBlogNameLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("blog name exceeds maximum of %d characters", MaxBlogNameLength),
		)
	}

	return nil
}

func ValidateBlogDescription(description string) error {
	if len(description) == 0 {
		return failure.NewValidationFailure("blog description cannot be empty")
	}

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

func (b *BlogEntity) Update(
	languages []string,
	categories []string,
	name, description, markdown MultiLangText,
	isPublished bool,
	estimatedReadTimeSeconds int64,
) error {
	// Validate languages
	if err := ValidateLanguages(languages); err != nil {
		return err
	}

	// Validate that all languages have corresponding content
	if err := ValidateMultiLangContent(languages, name, "name"); err != nil {
		return err
	}
	if err := ValidateMultiLangContent(languages, description, "description"); err != nil {
		return err
	}
	if err := ValidateMultiLangContent(languages, markdown, "markdown"); err != nil {
		return err
	}

	// Validate each language's content
	for lang, text := range name {
		if err := ValidateBlogName(text); err != nil {
			return failure.NewValidationFailure(fmt.Sprintf("name for language '%s': %s", lang, err.Error()))
		}
	}

	for lang, text := range description {
		if err := ValidateBlogDescription(text); err != nil {
			return failure.NewValidationFailure(fmt.Sprintf("description for language '%s': %s", lang, err.Error()))
		}
	}

	if err := ValidateBlogCategories(categories); err != nil {
		return err
	}

	b.Languages = languages
	b.Categories = categories
	b.Name = name
	b.Description = description
	b.Markdown = markdown
	b.IsPublished = isPublished
	b.EstimatedReadTimeSeconds = estimatedReadTimeSeconds
	b.UpdatedAt = time.Now().Unix()

	return nil
}
