package account

import (
	"encoding/json"
	"fmt"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/values"
	"time"

	"github.com/google/uuid"
)

const (
	MaxNameLength = 50
	MaxBioLength  = 300
)

type AccountEntity struct {
	ID            uuid.UUID          `json:"id"`
	Name          string             `json:"name"`
	Email         *values.Email      `json:"email"`
	EmailVerified bool               `json:"emailVerified"`
	Password      string             `json:"-"`
	Role          values.AccountRole `json:"role"`
	IsActive      bool               `json:"isActive"`
	Avatar        string             `json:"avatar"`
	Bio           string             `json:"bio"`
	CreatedAt     int64              `json:"createdAt"`
	UpdatedAt     int64              `json:"updatedAt"`
	DeletedAt     *int64             `json:"deletedAt,omitempty"`
}

type PublicJSON struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"emailVerified"`
	IsActive      bool      `json:"isActive"`
	Avatar        string    `json:"avatar,omitempty"`
	Bio           string    `json:"bio,omitempty"`
	CreatedAt     int64     `json:"createdAt"`
	UpdatedAt     int64     `json:"updatedAt"`
}

// MarshalJSON customizes JSON serialization to exclude password
func (a AccountEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(&PublicJSON{
		ID:            a.ID,
		Name:          a.Name,
		Email:         a.Email.Value(),
		EmailVerified: a.EmailVerified,
		IsActive:      a.IsActive,
		Avatar:        a.Avatar,
		Bio:           a.Bio,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	})
}

// ToPublicJSON converts AccountEntity to public representation
func (a *AccountEntity) ToPublicJSON() *PublicJSON {
	return &PublicJSON{
		ID:            a.ID,
		Name:          a.Name,
		Email:         a.Email.Value(),
		EmailVerified: a.EmailVerified,
		IsActive:      a.IsActive,
		Avatar:        a.Avatar,
		Bio:           a.Bio,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func NewAccount(
	name string,
	email *values.Email,
	password, avatar, bio string,
	role values.AccountRole,
) (*AccountEntity, error) {
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
		Role:          role,
		IsActive:      true,
		Avatar:        avatar,
		Bio:           bio,
		CreatedAt:     now.Unix(),
		UpdatedAt:     now.Unix(),
		DeletedAt:     nil,
	}, nil
}

func (a *AccountEntity) Update(
	name string,
	email *values.Email,
	password, avatar, bio string,
	role values.AccountRole,
) error {
	if err := ValidateName(name); err != nil {
		return err
	}

	if err := ValidateBio(bio); err != nil {
		return err
	}

	a.Name = name
	a.Email = email
	a.Password = password
	a.Avatar = avatar
	a.Bio = bio
	a.Role = role
	a.UpdatedAt = time.Now().Unix()

	return nil
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
