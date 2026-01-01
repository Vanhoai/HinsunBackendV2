package applications

import (
	"context"
	"hinsun-backend/internal/core/types"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/auth"
	"hinsun-backend/internal/domain/usecases"
	"hinsun-backend/internal/domain/values"
)

type AccountAppService interface {
	usecases.ManageAccountUseCase
}

type accountAppService struct {
	accountService account.AccountService
	authService    auth.AuthService
}

func NewAccountAppService(
	accountService account.AccountService,
	authService auth.AuthService,
) AccountAppService {
	return &accountAppService{
		accountService: accountService,
		authService:    authService,
	}
}

func (s *accountAppService) SearchAccounts(ctx context.Context, query *usecases.SearchAccountsQuery) ([]*account.AccountEntity, error) {
	return s.accountService.SearchAccountsByNameAndEmail(ctx, query.Name, query.Email)
}

func (s *accountAppService) FindAccountByEmail(ctx context.Context, email string) (*account.AccountEntity, error) {
	emailValidated, err := values.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return s.accountService.FindAccountByEmail(ctx, emailValidated)
}

func (s *accountAppService) FindAllAccounts(ctx context.Context) ([]*account.AccountEntity, error) {
	return s.accountService.FindAllAccounts(ctx)
}

func (s *accountAppService) CreateNewAccount(ctx context.Context, params *usecases.CreateAccountParams) (*account.AccountEntity, error) {
	email, err := values.NewEmail(params.Email)
	if err != nil {
		return nil, err
	}

	role, err := values.RoleFromInt(params.Role)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := s.authService.HashPassword(params.Password)
	if err != nil {
		return nil, err
	}

	return s.accountService.CreateNewAccount(
		ctx,
		params.Name,
		email,
		hashedPassword,
		params.Avatar,
		params.Bio,
		role,
	)
}

func (s *accountAppService) DeleteMultipleAccounts(ctx context.Context, query *usecases.DeleteAccountsQuery) (*types.DeletedResult, error) {
	rowsAffected, err := s.accountService.DeleteMultipleAccounts(ctx, query.IDs)
	if err != nil {
		return nil, err
	}

	deletedResult := &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      query.IDs,
	}

	return deletedResult, nil
}

func (s *accountAppService) FindAccountByID(ctx context.Context, id string) (*account.AccountEntity, error) {
	return s.accountService.FindAccountByID(ctx, id)
}

func (s *accountAppService) UpdateAccount(ctx context.Context, id string, params *usecases.UpdateAccountParams) (*account.AccountEntity, error) {
	email, err := values.NewEmail(params.Email)
	if err != nil {
		return nil, err
	}

	role, err := values.RoleFromInt(params.Role)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := s.authService.HashPassword(params.Password)
	if err != nil {
		return nil, err
	}

	return s.accountService.UpdateAccount(
		ctx,
		id,
		params.Name,
		email,
		hashedPassword,
		params.Avatar,
		params.Bio,
		role,
	)
}

func (s *accountAppService) DeleteAccount(ctx context.Context, id string) (*types.DeletedResult, error) {
	rowsAffected, err := s.accountService.DeleteAccount(ctx, id)
	if err != nil {
		return nil, err
	}

	deletedResult := &types.DeletedResult{
		RowsAffected: rowsAffected,
		Payload:      id,
	}

	return deletedResult, nil
}
