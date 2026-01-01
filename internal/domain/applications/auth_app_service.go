package applications

import (
	"context"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/auth"
	"hinsun-backend/internal/domain/usecases"
	"hinsun-backend/internal/domain/values"
)

type AuthAppService interface {
	usecases.ManageSessionAuthUseCase
}

type authAppService struct {
	authService    auth.AuthService
	accountService account.AccountService
}

func NewAuthAppService(
	authService auth.AuthService,
	accountService account.AccountService,
) AuthAppService {
	return &authAppService{
		authService:    authService,
		accountService: accountService,
	}
}

func (a *authAppService) AuthWithEmail(ctx context.Context, params *usecases.AuthEmailParams) (*usecases.AuthResponse, error) {
	accountEntity, err := a.EnsureAccountExists(ctx, params)
	if err != nil {
		return nil, err
	}

	// At here, accountEntity must be not nil
	tokenPair, err := a.authService.GenerateTokenPair(accountEntity.ID.String(), accountEntity.Email.Value())
	if err != nil {
		return nil, err
	}

	return &usecases.AuthResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

func (a *authAppService) EnsureAccountExists(ctx context.Context, params *usecases.AuthEmailParams) (*account.AccountEntity, error) {
	email, err := values.NewEmail(params.Email)
	if err != nil {
		return nil, err
	}

	// accountEntity is global to handle both cases
	var accountEntity *account.AccountEntity

	accountEntity, err = a.accountService.FindAccountByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// At here, there are two cases we need to handle:
	// 1. accountEntity is nil -> account not found -> need to create account first
	// 2. accountEntity is not nil -> account found -> proceed to password verification
	if accountEntity == nil {
		// create a new account
		name := email.LocalPart()
		hashedPassword, err := a.authService.HashPassword(params.Password)
		if err != nil {
			return nil, err
		}

		accountEntity, err = a.accountService.CreateNewAccount(
			ctx,
			name,
			email,
			hashedPassword,
			"https://i.pinimg.com/1200x/1d/b1/e6/1db1e66e7b8da532a1271b796621607d.jpg",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			values.NormalRole,
		)

		if err != nil {
			return nil, err
		}

	} else {
		// verify password
		if err := a.authService.VerifyPassword(params.Password, accountEntity.Password); err != nil {
			return nil, err
		}
	}

	return accountEntity, nil
}
