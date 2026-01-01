package usecases

import (
	"context"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/values"
)

type FindAccountByEmailParams struct {
	Email string `json:"email" validate:"required,email"`
}

type ManageAccountUseCase interface {
	FindAccountByEmail(ctx context.Context, email *values.Email) (*account.AccountEntity, error)
	FindAccountByID(ctx context.Context, id string) (*account.AccountEntity, error)
}
