package auth

import (
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/pkg/jwt"
	"hinsun-backend/pkg/security"
)

type AuthService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) error
	GenerateTokenPair(accountID, email string, role int) (*jwt.TokenPair, error)
}

type authService struct {
	passwordHasher security.PasswordHasher
	jwtService     jwt.JwtService
}

func NewAuthService(passwordHasher security.PasswordHasher, jwtService jwt.JwtService) AuthService {
	return &authService{
		passwordHasher: passwordHasher,
		jwtService:     jwtService,
	}
}

// Hash hashes the provided password.
func (s *authService) HashPassword(password string) (string, error) {
	hashed, err := s.passwordHasher.Hash(password)
	if err != nil {
		return "", failure.NewInternalFailure("failed to hash password", err)
	}

	return hashed, nil
}

// Verify checks if the provided password matches the hash.
func (s *authService) VerifyPassword(password, hash string) error {
	isValid, err := s.passwordHasher.Verify(password, hash)
	if err != nil {
		return failure.NewValidationFailure("failed to verify password").WithCause(err)
	}

	if !isValid {
		return failure.NewValidationFailure("invalid password")
	}

	return nil
}

// GenerateTokenPair generates a new JWT token pair for the given account ID and email.
func (s *authService) GenerateTokenPair(accountID, email string, role int) (*jwt.TokenPair, error) {
	tokenPair, err := s.jwtService.GenerateTokenPair(accountID, email, role)
	if err != nil {
		return nil, failure.NewInternalFailure("failed to generate token pair", err)
	}

	return tokenPair, nil
}
