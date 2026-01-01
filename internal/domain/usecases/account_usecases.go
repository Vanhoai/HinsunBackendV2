package usecases

import (
	"context"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/account"
)

type CreateAccountParams struct {
	Name          string `json:"name" validate:"required,min=2,max=50"`
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required,min=8"`
	EmailVerified bool   `json:"emailVerified"`
	Avatar        string `json:"avatar" validate:"omitempty,url"`
	Bio           string `json:"bio" validate:"omitempty,max=300"`
	IsActive      bool   `json:"isActive"`
	Role          int    `json:"role"`
}

type UpdateAccountParams struct {
	Name          string `json:"name" validate:"required,min=2,max=50"`
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required,min=8"`
	EmailVerified bool   `json:"emailVerified"`
	Avatar        string `json:"avatar" validate:"omitempty,url"`
	Bio           string `json:"bio" validate:"omitempty,max=300"`
	IsActive      bool   `json:"isActive"`
	Role          int    `json:"role"`
}

type DeleteAccountsQuery struct {
	IDs []string `query:"ids"`
}

type SearchAccountsQuery struct {
	Email string
	Name  string
}

type ManageAccountUseCase interface {
	SearchAccounts(ctx context.Context, query *SearchAccountsQuery) ([]*account.AccountEntity, error)
	FindAccountByEmail(ctx context.Context, email string) (*account.AccountEntity, error)
	FindAllAccounts(ctx context.Context) ([]*account.AccountEntity, error)
	CreateNewAccount(ctx context.Context, params *CreateAccountParams) (*account.AccountEntity, error)
	DeleteMultipleAccounts(ctx context.Context, query *DeleteAccountsQuery) (*types.DeletedResult, error)

	FindAccountByID(ctx context.Context, id string) (*account.AccountEntity, error)
	UpdateAccount(ctx context.Context, id string, params *UpdateAccountParams) (*account.AccountEntity, error)
	DeleteAccount(ctx context.Context, id string) (*types.DeletedResult, error)
}
