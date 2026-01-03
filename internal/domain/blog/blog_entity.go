package blog

import (
	"fmt"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/values"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

const (
	MaxBlogNameLength = 100
	MaxBlogDescLength = 300
	MaxCategories     = 5
	MaxLanguages      = 10
	MinLanguages      = 1
)

type BlogEntity struct {
	ID                       uuid.UUID                     `json:"id"`
	AuthorID                 uuid.UUID                     `json:"authorId"`
	Slug                     string                        `json:"slug"`
	Languages                []values.MarkdownLanguageCode `json:"languages"` // Array of supported language codes
	Categories               []string                      `json:"categories"`
	Names                    values.MultiLangText          `json:"names"`        // Map of language code -> name
	Descriptions             values.MultiLangText          `json:"descriptions"` // Map of language code -> description
	IsPublished              bool                          `json:"isPublished"`
	Markdowns                values.MultiLangText          `json:"markdowns"` // Map of language code -> markdown content
	Favorites                int64                         `json:"favorites"`
	Views                    int64                         `json:"views"`
	EstimatedReadTimeSeconds int64                         `json:"estimatedReadTimeSeconds"`
	CreatedAt                int64                         `json:"createdAt"`
	UpdatedAt                int64                         `json:"updatedAt"`
	DeletedAt                *int64                        `json:"deletedAt,omitempty"`
}

func NewBlogEntity(
	authorID uuid.UUID,
	languages []values.MarkdownLanguageCode,
	categories []string,
	names, descriptions, markdowns values.MultiLangText,
	isPublished bool,
	estimatedReadTimeSeconds int64,
) (*BlogEntity, error) {
	// Validate languages
	if err := ValidateLanguages(languages); err != nil {
		return nil, err
	}

	// Validate that all languages have corresponding content
	if err := ValidateMultiLangContent(languages, names, "name"); err != nil {
		return nil, err
	}

	if err := ValidateMultiLangContent(languages, descriptions, "description"); err != nil {
		return nil, err
	}

	if err := ValidateMultiLangContent(languages, markdowns, "markdown"); err != nil {
		return nil, err
	}

	// Validate each language's content
	for lang, text := range names {
		if err := ValidateBlogName(text); err != nil {
			return nil, failure.NewValidationFailure(fmt.Sprintf("name for language '%s': %s", lang, err.Error()))
		}
	}

	for lang, text := range descriptions {
		if err := ValidateBlogDescription(text); err != nil {
			return nil, failure.NewValidationFailure(fmt.Sprintf("description for language '%s': %s", lang, err.Error()))
		}
	}

	if err := ValidateBlogCategories(categories); err != nil {
		return nil, err
	}

	blogSlug := ""
	// prioritize English name for slug generation
	if enName, exists := names["en"]; exists {
		blogSlug = slug.Make(enName)
	}

	// fallback to first available name
	if blogSlug == "" {
		for _, name := range names {
			blogSlug = slug.Make(name)
			break
		}
	}

	now := time.Now()
	return &BlogEntity{
		ID:                       uuid.New(),
		AuthorID:                 authorID,
		Slug:                     blogSlug,
		Categories:               categories,
		Names:                    names,
		Languages:                languages,
		Descriptions:             descriptions,
		IsPublished:              isPublished,
		Markdowns:                markdowns,
		Favorites:                0,
		Views:                    0,
		EstimatedReadTimeSeconds: estimatedReadTimeSeconds,
		CreatedAt:                now.Unix(),
		UpdatedAt:                now.Unix(),
		DeletedAt:                nil,
	}, nil
}

func ValidateLanguages(languages []values.MarkdownLanguageCode) error {
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
	seen := make(map[values.MarkdownLanguageCode]bool)
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

func ValidateMultiLangContent(languages []values.MarkdownLanguageCode, content values.MultiLangText, fieldName string) error {
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
	authorID uuid.UUID,
	languages []values.MarkdownLanguageCode,
	categories []string,
	names, descriptions, markdowns values.MultiLangText,
	isPublished bool,
	estimatedReadTimeSeconds int64,
) error {
	// Validate languages
	if err := ValidateLanguages(languages); err != nil {
		return err
	}

	// Validate that all languages have corresponding content
	if err := ValidateMultiLangContent(languages, names, "name"); err != nil {
		return err
	}
	if err := ValidateMultiLangContent(languages, descriptions, "description"); err != nil {
		return err
	}
	if err := ValidateMultiLangContent(languages, markdowns, "markdown"); err != nil {
		return err
	}

	// Validate each language's content
	for lang, text := range names {
		if err := ValidateBlogName(text); err != nil {
			return failure.NewValidationFailure(fmt.Sprintf("name for language '%s': %s", lang, err.Error()))
		}
	}

	for lang, text := range descriptions {
		if err := ValidateBlogDescription(text); err != nil {
			return failure.NewValidationFailure(fmt.Sprintf("description for language '%s': %s", lang, err.Error()))
		}
	}

	if err := ValidateBlogCategories(categories); err != nil {
		return err
	}

	b.AuthorID = authorID
	b.Languages = languages
	b.Categories = categories
	b.Names = names
	b.Descriptions = descriptions
	b.Markdowns = markdowns
	b.IsPublished = isPublished
	b.EstimatedReadTimeSeconds = estimatedReadTimeSeconds

	return nil
}
