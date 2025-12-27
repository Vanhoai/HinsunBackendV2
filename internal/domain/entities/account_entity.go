package entities

import (
	"hinsun-backend/internal/domain/values"
	"time"

	"github.com/google/uuid"
)

type AccountEntity struct {
	ID            uuid.UUID    `json:"id"`
	Name          string       `json:"name"`
	Email         values.Email `json:"email"`
	Password      string       `json:"password,omitempty"`
	EmailVerified bool         `json:"emailVerified"`
	IsActive      bool         `json:"isActive"`
	Avatar        string       `json:"avatar"`
	Bio           string       `json:"bio"`
	CreatedAt     int64        `json:"createdAt"`
	UpdatedAt     int64        `json:"updatedAt"`
	DeletedAt     *int64       `json:"deletedAt,omitempty"`
}

func NewAccount(name string, email values.Email, password string, avatar string, bio string) *AccountEntity {
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
	}
}

func (a *AccountEntity) VerifyEmail() {
	a.EmailVerified = true
	a.UpdatedAt = time.Now().Unix()
}

func (a *AccountEntity) Deactivate() {
	a.IsActive = false
	a.UpdatedAt = time.Now().Unix()
}
