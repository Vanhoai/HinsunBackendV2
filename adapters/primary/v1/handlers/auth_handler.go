package handlers

import (
	"encoding/json"
	"hinsun-backend/adapters/shared/https"
	"hinsun-backend/adapters/shared/middlewares"
	"hinsun-backend/internal/domain/applications"
	"hinsun-backend/internal/domain/usecases"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	app            applications.AuthAppService
	validator      *validator.Validate
	authMiddleware *middlewares.AuthMiddleware
	roleMiddleware *middlewares.RoleMiddleware
}

func NewAuthHandler(
	app applications.AuthAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *AuthHandler {
	return &AuthHandler{
		app:            app,
		validator:      validator,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
	}
}

func (h *AuthHandler) Handler() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.authEmail)
	r.Post("/refresh", h.refreshTokens)
	r.With(h.authMiddleware.RequireAuth).Post("/signout", h.signOut)

	return r
}

func (h *AuthHandler) authEmail(w http.ResponseWriter, r *http.Request) {
	var params usecases.AuthEmailParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	response, err := h.app.AuthWithEmail(r.Context(), &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Authentication successful", response)
}

func (h *AuthHandler) refreshTokens(w http.ResponseWriter, r *http.Request) {
	var params usecases.RefreshTokensParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	response, err := h.app.RefreshTokens(r.Context(), &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Tokens refreshed successfully", response)
}

func (h *AuthHandler) signOut(w http.ResponseWriter, r *http.Request) {
	err := h.app.SignOut(r.Context())
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Signed out successfully", nil)
}
