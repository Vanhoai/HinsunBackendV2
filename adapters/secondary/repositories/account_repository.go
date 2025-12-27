package repositories

import (
	"context"
	"hinsun-backend/adapters/shared/models"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/values"

	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) account.AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) Create(ctx context.Context, account *account.AccountEntity) error {
	model := models.FromAccountEntity(account)
	err := gorm.G[models.AccountModel](r.db).Create(ctx, &model)
	if err != nil {
		return failure.NewDatabaseFailure("Failed to create account in database").WithCause(err)
	}

	return nil
}

func (r *accountRepository) FindByEmail(ctx context.Context, email *values.Email) (*account.AccountEntity, error) {
	account, err := gorm.G[models.AccountModel](r.db).Where("email = ?", email.Value()).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, failure.NewDatabaseFailure("Failed to retrieve account from database").WithCause(err)
	}

	return models.ToAccountEntity(&account), nil
}

func (r *accountRepository) Update(ctx context.Context, account *account.AccountEntity) (int, error) {
	return 0, nil
}

func (r *accountRepository) Delete(ctx context.Context, id string) (int, error) {
	return 0, nil
}

func (r *accountRepository) FindByID(ctx context.Context, id string) (*account.AccountEntity, error) {
	return nil, nil
}
