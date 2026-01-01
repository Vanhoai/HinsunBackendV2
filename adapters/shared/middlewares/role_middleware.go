package middlewares

import (
	"hinsun-backend/adapters/shared/https"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/values"
	"net/http"
	"strconv"
)

type RoleMiddleware struct{}

func NewRoleMiddleware() *RoleMiddleware {
	return &RoleMiddleware{}
}

// RequireRole checks if user has required role or higher
func (m *RoleMiddleware) RequireRole(requiredRole values.AccountRole) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get claims from context (set by AuthMiddleware)
			claims, ok := GetClaimsFromContext(r.Context())
			if !ok {
				https.RespondWithFailure(w, failure.NewAuthenticationFailure("authentication required"))
				return
			}

			// Convert role from string to AccountRole
			roleInt, err := strconv.Atoi(claims.Role)
			if err != nil {
				https.RespondWithFailure(w, failure.NewAuthenticationFailure("invalid role format"))
				return
			}

			// Check role permission
			if roleInt < int(requiredRole) {
				https.RespondWithFailure(w, failure.NewAuthenticationFailure("insufficient permissions"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequireAdmin is a helper for admin role
func (m *RoleMiddleware) RequireAdmin(next http.Handler) http.Handler {
	return m.RequireRole(values.AdminRole)(next)
}

// RequireGod is a helper for god role
func (m *RoleMiddleware) RequireGod(next http.Handler) http.Handler {
	return m.RequireRole(values.GodRole)(next)
}
