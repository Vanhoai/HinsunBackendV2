package account

import (
	"context"
	"hinsun-backend/internal/domain/values"
)

type AccountRepository interface {
	Create(ctx context.Context, account *AccountEntity) error
	FindByEmail(ctx context.Context, email *values.Email) (*AccountEntity, error)
	Update(ctx context.Context, account *AccountEntity) (int, error)
	Delete(ctx context.Context, id string) (int, error)
	FindByID(ctx context.Context, id string) (*AccountEntity, error)
}
