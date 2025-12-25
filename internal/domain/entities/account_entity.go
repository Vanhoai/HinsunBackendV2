package entities

import (
	"hinsun-backend/internal/domain/values"
	"time"

	"github.com/google/uuid"
)

type AccountEntity struct {
	ID            uuid.UUID
	Name          string
	Email         values.Email
	Password      string
	EmailVerified bool
	IsActive      bool
	Avatar        string
	Bio           string
	CreatedAt     int64
	UpdatedAt     int64
	DeletedAt     *int64
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
