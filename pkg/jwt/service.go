package jwt

import (
	"errors"
	"fmt"
	"hinsun-backend/pkg/security"
	"time"

	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type JwtService interface {
	GenerateTokenPair(accountID, email string) (*TokenPair, error)
	ValidateAccessToken(tokenString string) (*Claims, error)
	ValidateRefreshToken(tokenString string) (*Claims, error)
}

type jwtService struct {
	algorithm          security.Algorithm
	keyManager         *KeyManager
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}

func NewJwtService(keyManager *KeyManager, accessTokenExpiry, refreshTokenExpiry time.Duration) JwtService {
	return &jwtService{
		keyManager:         keyManager,
		algorithm:          keyManager.algorithm,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}

func (j *jwtService) retrieveSigningMethod() jwtv4.SigningMethod {
	switch j.algorithm {
	case security.RS256:
		return jwtv4.SigningMethodRS256
	case security.RS384:
		return jwtv4.SigningMethodRS384
	case security.RS512:
		return jwtv4.SigningMethodRS512
	case security.ES256:
		return jwtv4.SigningMethodES256
	case security.ES384:
		return jwtv4.SigningMethodES384
	case security.ES521:
		return jwtv4.SigningMethodES512
	case security.HS256:
		return jwtv4.SigningMethodHS256
	default:
		return jwtv4.SigningMethodHS256
	}
}

func (j *jwtService) GenerateTokenPair(accountID, email string) (*TokenPair, error) {
	jti := uuid.New().String()
	now := time.Now()

	accessExpiresAt := now.Add(j.accessTokenExpiry)
	refreshExpiresAt := now.Add(j.refreshTokenExpiry)
	signingMethod := j.retrieveSigningMethod()

	accessClaims := Claims{
		AccountID: accountID,
		Email:     email,
		JTI:       jti,
		RegisteredClaims: jwtv4.RegisteredClaims{
			ExpiresAt: jwtv4.NewNumericDate(accessExpiresAt),
			IssuedAt:  jwtv4.NewNumericDate(now),
			NotBefore: jwtv4.NewNumericDate(now),
		},
	}

	accessToken := jwtv4.NewWithClaims(signingMethod, accessClaims)
	accessTokenString, err := accessToken.SignedString(j.keyManager.accessKey.SigningKey)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshClaims := Claims{
		AccountID: accountID,
		Email:     email,
		JTI:       jti,
		RegisteredClaims: jwtv4.RegisteredClaims{
			ExpiresAt: jwtv4.NewNumericDate(refreshExpiresAt),
			IssuedAt:  jwtv4.NewNumericDate(now),
			NotBefore: jwtv4.NewNumericDate(now),
		},
	}

	refreshToken := jwtv4.NewWithClaims(signingMethod, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(j.keyManager.refreshKey.SigningKey)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (j *jwtService) ValidateAccessToken(token string) (*Claims, error) {
	return j.validateToken(token, j.keyManager.accessKey.VerificationKey)
}

func (j *jwtService) ValidateRefreshToken(token string) (*Claims, error) {
	return j.validateToken(token, j.keyManager.refreshKey.VerificationKey)
}

func (j *jwtService) validateToken(token string, verificationKey any) (*Claims, error) {
	parsedToken, err := jwtv4.ParseWithClaims(token, &Claims{}, func(t *jwtv4.Token) (any, error) {
		return verificationKey, nil
	})

	if err != nil {
		if errors.Is(err, jwtv4.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}

		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok || !parsedToken.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
