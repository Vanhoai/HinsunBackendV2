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

	entity, err := account.ToEntity()
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *accountRepository) Update(ctx context.Context, account *account.AccountEntity) (int, error) {
	rowsAffected, err := gorm.G[models.AccountModel](r.db).Where("id = ?", account.ID).Updates(ctx, models.FromAccountEntity(account))
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to update account in database").WithCause(err)
	}

	return rowsAffected, nil
}

func (r *accountRepository) Delete(ctx context.Context, id string) (int, error) {
	rowAffected, err := gorm.G[models.AccountModel](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete account from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *accountRepository) DeleteMany(ctx context.Context, ids []string) (int, error) {
	rowAffected, err := gorm.G[models.AccountModel](r.db).Where("id IN ?", ids).Delete(ctx)
	if err != nil {
		return 0, failure.NewDatabaseFailure("Failed to delete accounts from database").WithCause(err)
	}

	return rowAffected, nil
}

func (r *accountRepository) FindByID(ctx context.Context, id string) (*account.AccountEntity, error) {
	accountModel, err := gorm.G[models.AccountModel](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, failure.NewDatabaseFailure("Failed to retrieve account from database").WithCause(err)
	}

	entity, err := accountModel.ToEntity()
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *accountRepository) FindAll(ctx context.Context) ([]*account.AccountEntity, error) {
	blogs, err := gorm.G[models.AccountModel](r.db).Find(ctx)
	if err != nil {
		return nil, failure.NewDatabaseFailure("Failed to retrieve blogs from database").WithCause(err)
	}

	var accountEntities []*account.AccountEntity
	for _, accountModel := range blogs {
		entity, err := accountModel.ToEntity()
		if err != nil {
			return nil, err
		}

		accountEntities = append(accountEntities, entity)
	}

	return accountEntities, nil
}

func (r *accountRepository) SearchByNameAndEmail(ctx context.Context, name, email string) ([]*account.AccountEntity, error) {
	var accountModels []models.AccountModel
	accountModels, err := gorm.G[models.AccountModel](r.db).Where("email LIKE ?", "%"+email+"%").Where("name LIKE ?", "%"+name+"%").Find(ctx)
	if err != nil {
		return nil, failure.NewDatabaseFailure("Failed to search accounts in database").WithCause(err)
	}

	var accountEntities []*account.AccountEntity
	for _, accountModel := range accountModels {
		entity, err := accountModel.ToEntity()
		if err != nil {
			return nil, err
		}
		accountEntities = append(accountEntities, entity)
	}

	return accountEntities, nil
}
