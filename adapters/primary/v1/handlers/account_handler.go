package handlers

import (
	"hinsun-backend/adapters/shared/https"
	"hinsun-backend/internal/domain/applications"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type AccountHandler struct {
	globalAppService applications.GlobalAppService
	validator        *validator.Validate
}

func NewAccountHandler(globalAppService applications.GlobalAppService, validator *validator.Validate) *AccountHandler {
	return &AccountHandler{
		globalAppService: globalAppService,
		validator:        validator,
	}
}

func (h *AccountHandler) Handler() chi.Router {
	r := chi.NewRouter()

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.findAccountByID)
	})

	return r
}

func (h *AccountHandler) findAccountByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	account, err := h.globalAppService.FindAccountByID(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Account retrieved successfully", account)
}
