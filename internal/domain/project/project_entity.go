package project

import (
	"fmt"
	"hinsun-backend/internal/core/failure"
	"time"

	"github.com/google/uuid"
)

const (
	MaxProjectNameLength        = 100
	MaxProjectDescriptionLength = 500
	MaxProjectTags              = 5
)

type ProjectEntity struct {
	ID          uuid.UUID
	Name        string
	Description string
	Cover       string
	Github      string
	Tags        []string
	Markdown    string
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   *int64
}

func NewProjectEntity(id uuid.UUID, name, description, github, cover string, tags []string, markdown string) *ProjectEntity {
	now := time.Now()
	return &ProjectEntity{
		ID:          id,
		Name:        name,
		Description: description,
		Cover:       cover,
		Github:      github,
		Tags:        tags,
		Markdown:    markdown,
		CreatedAt:   now.Unix(),
		UpdatedAt:   now.Unix(),
		DeletedAt:   nil,
	}
}

func ValidateProjectName(name string) error {
	if len(name) > MaxProjectNameLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("project name exceeds maximum of %d characters", MaxProjectNameLength),
		)
	}

	return nil
}

func ValidateProjectDescription(description string) error {
	if len(description) > MaxProjectDescriptionLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("project description exceeds maximum of %d characters", MaxProjectDescriptionLength),
		)
	}

	return nil
}

func ValidateProjectTags(tags []string) error {
	if len(tags) == 0 {
		return failure.NewValidationFailure(
			"project must have at least one tag",
		)
	}

	if len(tags) > MaxProjectTags {
		return failure.NewValidationFailure(
			fmt.Sprintf("number of project tags exceeds maximum of %d", MaxProjectTags),
		)
	}

	return nil
}

func (p *ProjectEntity) Update(
	name, description, github, cover string,
	tags []string,
	markdown string,
) error {
	if err := ValidateProjectName(name); err != nil {
		return err
	}

	if err := ValidateProjectDescription(description); err != nil {
		return err
	}

	if err := ValidateProjectTags(tags); err != nil {
		return err
	}

	p.Name = name
	p.Description = description
	p.Cover = cover
	p.Github = github
	p.Tags = tags
	p.Markdown = markdown
	p.UpdatedAt = time.Now().Unix()

	return nil
}
