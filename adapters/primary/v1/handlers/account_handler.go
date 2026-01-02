package handlers

import (
	"encoding/json"
	"errors"
	"hinsun-backend/adapters/shared/https"
	"hinsun-backend/adapters/shared/middlewares"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/applications"
	"hinsun-backend/internal/domain/usecases"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type AccountHandler struct {
	app            applications.AccountAppService
	validator      *validator.Validate
	authMiddleware *middlewares.AuthMiddleware
	roleMiddleware *middlewares.RoleMiddleware
}

func NewAccountHandler(
	app applications.AccountAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *AccountHandler {
	return &AccountHandler{
		app:            app,
		validator:      validator,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
	}
}

func (h *AccountHandler) Handler() chi.Router {
	r := chi.NewRouter()

	r.With(h.authMiddleware.RequireAuth).Get("/", h.findAllAccounts)
	r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireGod).Post("/", h.createAccount)
	r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireGod).Delete("/", h.deleteMultipleAccounts)

	r.Get("/search", h.searchAccounts)
	r.With(h.authMiddleware.RequireAuth).Get("/profile", h.findAccountProfile)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.findAccountByID)

		r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireGod).Put("/", h.updateAccount)
		r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireGod).Delete("/", h.deleteAccount)
	})

	return r
}

func (h *AccountHandler) findAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.app.FindAllAccounts(r.Context())
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Accounts retrieved successfully", accounts)
}

func (h *AccountHandler) searchAccounts(w http.ResponseWriter, r *http.Request) {
	var query usecases.SearchAccountsQuery
	if err := https.BindQuery(r, &query); err != nil {
		https.BadRequest(w, err)
		return
	}

	accounts, err := h.app.SearchAccounts(r.Context(), &query)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Accounts retrieved successfully", accounts)
}

func (h *AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	var params usecases.CreateAccountParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	account, err := h.app.CreateNewAccount(r.Context(), &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusCreated, "Account created successfully", account)
}

func (h *AccountHandler) deleteMultipleAccounts(w http.ResponseWriter, r *http.Request) {
	var query usecases.DeleteAccountsQuery
	if err := https.BindQuery(r, &query); err != nil {
		https.BadRequest(w, err)
		return
	}

	// Validate that at least one ID is provided
	if len(query.IDs) == 0 {
		https.BadRequest(w, errors.New("at least one id must be provided in ids parameter"))
		return
	}

	result, err := h.app.DeleteMultipleAccounts(r.Context(), &query)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Accounts deleted successfully", result)
}

func (h *AccountHandler) findAccountByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	account, err := h.app.FindAccountByID(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Account retrieved successfully", account)
}

func (h *AccountHandler) updateAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var params usecases.UpdateAccountParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	updatedAccount, err := h.app.UpdateAccount(r.Context(), id, &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Account updated successfully", updatedAccount)
}

func (h *AccountHandler) deleteAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	deletedResult, err := h.app.DeleteAccount(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Account deleted successfully", deletedResult)
}

func (h *AccountHandler) findAccountProfile(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.GetClaimsFromContext(r.Context())
	if !ok {
		https.RespondWithFailure(w, failure.NewAuthenticationFailure("authentication required"))
		return
	}

	account, err := h.app.FindAccountByID(r.Context(), claims.AccountID)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Account profile retrieved successfully", account)
}
