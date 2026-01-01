package account

import (
	"context"
	"fmt"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/values"
)

type AccountService interface {
	FindAccountByEmail(ctx context.Context, email *values.Email) (*AccountEntity, error)
	CreateNewAccount(ctx context.Context, name string, email *values.Email, hashedPassword, avatar, bio string, role values.AccountRole) (*AccountEntity, error)
	FindAllAccounts(ctx context.Context) ([]*AccountEntity, error)
	DeleteMultipleAccounts(ctx context.Context, ids []string) (int, error)

	FindAccountByID(ctx context.Context, id string) (*AccountEntity, error)
	DeleteAccount(ctx context.Context, id string) (int, error)
	UpdateAccount(ctx context.Context, id string, name string, email *values.Email, hashedPassword, avatar, bio string, role values.AccountRole) (*AccountEntity, error)
}

type accountService struct {
	repository AccountRepository
}

func NewAccountService(repository AccountRepository) AccountService {
	return &accountService{
		repository: repository,
	}
}

func (s *accountService) SearchAccounts(ctx context.Context) ([]*AccountEntity, error) {
	return nil, nil
}

// FindAccountByEmail retrieves an account entity by its email.
func (s *accountService) FindAccountByEmail(ctx context.Context, email *values.Email) (*AccountEntity, error) {
	return s.repository.FindByEmail(ctx, email)
}

func (s *accountService) FindAllAccounts(ctx context.Context) ([]*AccountEntity, error) {
	return s.repository.FindAll(ctx)
}

// CreateNewAccount creates a new account with the provided email and password.
func (s *accountService) CreateNewAccount(ctx context.Context, name string, email *values.Email, hashedPassword, avatar, bio string, role values.AccountRole) (*AccountEntity, error) {
	// 1. Check if account with the email already exists
	accountEntity, err := s.repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if accountEntity != nil {
		// Account already exists
		return nil, failure.NewConflictFailure("account with the provided email already exists")
	}

	// 2. Create new AccountEntity
	accountEntity, err = NewAccount(
		name,
		email,
		hashedPassword,
		avatar,
		bio,
		role,
	)

	if err != nil {
		// Return validation error from entity creation
		return nil, err
	}

	// 3. Save to repository
	if err := s.repository.Create(ctx, accountEntity); err != nil {
		return nil, err
	}

	// 4. Return the created account entity
	return accountEntity, nil
}

func (s *accountService) FindAccountByID(ctx context.Context, id string) (*AccountEntity, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *accountService) DeleteAccount(ctx context.Context, id string) (int, error) {
	return s.repository.Delete(ctx, id)
}

func (s *accountService) DeleteMultipleAccounts(ctx context.Context, ids []string) (int, error) {
	return s.repository.DeleteMany(ctx, ids)
}

func (s *accountService) UpdateAccount(ctx context.Context, id string, name string, email *values.Email, hashedPassword, avatar, bio string, role values.AccountRole) (*AccountEntity, error) {
	// 1. Retrieve existing account
	existingAccount, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if existingAccount == nil {
		return nil, failure.NewNotFoundFailure("Account with the given ID does not exist")
	}

	fmt.Println("Existing Account Email:", existingAccount.Email.Value())

	// 2. Check for email conflict if email is being changed
	if existingAccount.Email.Value() != email.Value() {
		conflictAccount, err := s.repository.FindByEmail(ctx, email)
		if err != nil {
			return nil, err
		}

		if conflictAccount != nil {
			return nil, failure.NewConflictFailure("Account with the same email already exists")
		}
	}

	// 3. Update fields
	err = existingAccount.Update(
		name,
		email,
		hashedPassword,
		avatar,
		bio,
		role,
	)

	if err != nil {
		return nil, err
	}

	// 4. Save updated account
	rowsAffected, err := s.repository.Update(ctx, existingAccount)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, failure.NewNotFoundFailure("Account with the given ID does not exist")
	}

	return existingAccount, nil
}
