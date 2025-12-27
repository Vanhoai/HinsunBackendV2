package account

import (
	"fmt"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/values"
	"time"

	"github.com/google/uuid"
)

const (
	MaxNameLength = 50
	MaxBioLength  = 160
)

type AccountEntity struct {
	ID            uuid.UUID     `json:"id"`
	Name          string        `json:"name"`
	Email         *values.Email `json:"email"`
	Password      string        `json:"password,omitempty"`
	EmailVerified bool          `json:"emailVerified"`
	IsActive      bool          `json:"isActive"`
	Avatar        string        `json:"avatar"`
	Bio           string        `json:"bio"`
	CreatedAt     int64         `json:"createdAt"`
	UpdatedAt     int64         `json:"updatedAt"`
	DeletedAt     *int64        `json:"deletedAt,omitempty"`
}

func NewAccount(name string, email *values.Email, password, avatar, bio string) (*AccountEntity, error) {
	if err := ValidateName(name); err != nil {
		return nil, err
	}

	if err := ValidateBio(bio); err != nil {
		return nil, err
	}

	now := time.Now()
	return &AccountEntity{
		ID:            uuid.New(),
		Name:          name,
		Email:         email,
		Password:      password,
		EmailVerified: false,
		IsActive:      true,
		Avatar:        avatar,
		Bio:           bio,
		CreatedAt:     now.Unix(),
		UpdatedAt:     now.Unix(),
		DeletedAt:     nil,
	}, nil
}

func ValidateName(name string) error {
	if len(name) == 0 {
		return failure.NewValidationFailure("name cannot be empty")
	}

	if len(name) > MaxNameLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("name exceeds maximum length of %d", MaxNameLength),
		)
	}

	return nil
}

func ValidateBio(bio string) error {
	if len(bio) > MaxBioLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("bio exceeds maximum length of %d", MaxBioLength),
		)
	}

	return nil
}
