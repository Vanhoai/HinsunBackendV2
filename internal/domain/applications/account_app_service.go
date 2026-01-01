package applications

import (
	"context"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/values"
)

type AccountAppService interface {
	FindAccountByEmail(ctx context.Context, email *values.Email) (*account.AccountEntity, error)
	FindAccountByID(ctx context.Context, id string) (*account.AccountEntity, error)
}

type accountAppService struct {
	accountService account.AccountService
}

func NewAccountAppService(
	accountService account.AccountService,
) AccountAppService {
	return &accountAppService{
		accountService: accountService,
	}
}

func (s *accountAppService) FindAccountByEmail(ctx context.Context, email *values.Email) (*account.AccountEntity, error) {
	return s.accountService.FindAccountByEmail(ctx, email)
}

func (s *accountAppService) FindAccountByID(ctx context.Context, id string) (*account.AccountEntity, error) {
	// This needs to be implemented in account service first
	// For now, return nil to satisfy the interface
	return nil, nil
}
