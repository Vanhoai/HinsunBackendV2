package middlewares

import (
	"context"
	"hinsun-backend/adapters/shared/https"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/pkg/jwt"
	"net/http"
	"strings"
)

type contextKey string

const (
	ClaimsContextKey contextKey = "claims"
)

type AuthMiddleware struct {
	jwtService jwt.JwtService
}

func NewAuthMiddleware(jwtService jwt.JwtService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

// RequireAuth validates JWT access token
func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			https.RespondWithFailure(w, failure.NewAuthenticationFailure("missing authorization header"))
			return
		}

		// Check Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			https.RespondWithFailure(w, failure.NewAuthenticationFailure("invalid authorization header format"))
			return
		}

		token := parts[1]

		// Validate token
		claims, err := m.jwtService.ValidateAccessToken(token)
		if err != nil {
			https.RespondWithFailure(w, failure.NewAuthenticationFailure("invalid or expired token").WithCause(err))
			return
		}

		// Add claims to context
		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetClaimsFromContext retrieves claims from request context
func GetClaimsFromContext(ctx context.Context) (*jwt.Claims, bool) {
	claims, ok := ctx.Value(ClaimsContextKey).(*jwt.Claims)
	return claims, ok
}
