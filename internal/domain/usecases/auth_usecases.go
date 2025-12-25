package usecases

import "hinsun-backend/internal/domain/values"

type AuthEmailParams struct {
	Email    string `json:"email" example:"example@example.com"`
	Password string `json:"password" example:"strongpassword123"`
}

type OAuth2Params struct {
	Provider values.OAuthProvider
}

type OAuth2Response struct {
	AuthorizationUrl string
	State            string
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string `json:"refreshToken" example:"dGhpcy1pcz1hLXJlZnJlc2gtdG9rZW4tZXhhbXBsZQ..."`
}

type OAuth2CallbackParams struct {
	Code  string
	State string
}

type ManageSessionAuthUseCase interface {
	authWithEmail(params *AuthEmailParams) (*AuthResponse, error)
	oauth2(params *OAuth2Params) (*OAuth2Response, error)
	oauth2Callback(params *OAuth2CallbackParams) (*AuthResponse, error)
}
