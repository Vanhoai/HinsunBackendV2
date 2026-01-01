package handlers

import (
	"encoding/json"
	"errors"
	"hinsun-backend/adapters/shared/https"
	"hinsun-backend/internal/domain/applications"
	"hinsun-backend/internal/domain/usecases"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type AccountHandler struct {
	app       applications.AccountAppService
	validator *validator.Validate
}

func NewAccountHandler(app applications.AccountAppService, validator *validator.Validate) *AccountHandler {
	return &AccountHandler{
		app:       app,
		validator: validator,
	}
}

func (h *AccountHandler) Handler() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.findAllAccounts)
	r.Post("/", h.createAccount)
	r.Delete("/", h.deleteMultipleAccounts)

	r.Get("/search", h.searchAccounts)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.findAccountByID)
		r.Put("/", h.updateAccount)
		r.Delete("/", h.deleteAccount)
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

	https.ResponseSuccess(w, http.StatusOK, "Accounts retrieved successfully", nil)
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
