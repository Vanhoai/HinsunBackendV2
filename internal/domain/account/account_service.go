package account

import (
	"context"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/values"
)

type AccountService interface {
	FindAccountByEmail(ctx context.Context, email *values.Email) (*AccountEntity, error)
	CreateNewAccount(ctx context.Context, name string, email *values.Email, hashedPassword, avatar, bio string) (*AccountEntity, error)
}

type accountService struct {
	repository AccountRepository
}

func NewAccountService(repository AccountRepository) AccountService {
	return &accountService{
		repository: repository,
	}
}

// FindAccountByEmail retrieves an account entity by its email.
func (s *accountService) FindAccountByEmail(ctx context.Context, email *values.Email) (*AccountEntity, error) {
	return s.repository.FindByEmail(ctx, email)
}

// CreateNewAccount creates a new account with the provided email and password.
func (s *accountService) CreateNewAccount(ctx context.Context, name string, email *values.Email, hashedPassword, avatar, bio string) (*AccountEntity, error) {
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
