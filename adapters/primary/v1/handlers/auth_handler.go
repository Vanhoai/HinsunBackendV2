package handlers

import (
	"hinsun-backend/adapters/shared/https"
	"hinsun-backend/internal/domain/applications"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	app applications.AuthAppService
}

func NewAuthHandler(app applications.AuthAppService) *AuthHandler {
	return &AuthHandler{
		app: app,
	}
}

func (ah *AuthHandler) Handler() chi.Router {
	r := chi.NewRouter()

	r.Post("/auth-email", authEmail)
	r.Get("/oauth2", oauth2)

	return r
}

// authEmail godoc
// @Summary      Sign in with email and password
// @Description  Authenticate an account using email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      AuthEmailParams  true  "Login credentials"
// @Success      200  {object}  AuthResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      401  {object}  ErrorResponse
// @Router       /auth [post]
func authEmail(w http.ResponseWriter, r *http.Request) {
	https.ResponseSuccess(w, http.StatusOK, "Auth email endpoint", nil)
}

func oauth2(w http.ResponseWriter, r *http.Request) {
	https.ResponseSuccess(w, http.StatusOK, "OAuth2 endpoint", nil)
}
