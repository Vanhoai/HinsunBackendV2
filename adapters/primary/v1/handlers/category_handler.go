package handlers

import (
	"encoding/json"
	"errors"
	"hinsun-backend/adapters/shared/https"
	"hinsun-backend/adapters/shared/middlewares"
	"hinsun-backend/internal/domain/applications"
	"hinsun-backend/internal/domain/usecases"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type CategoryHandler struct {
	app            applications.CategoryAppService
	validator      *validator.Validate
	authMiddleware *middlewares.AuthMiddleware
	roleMiddleware *middlewares.RoleMiddleware
}

func NewCategoryHandler(
	app applications.CategoryAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *CategoryHandler {
	return &CategoryHandler{
		app:            app,
		validator:      validator,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
	}
}

func (h *CategoryHandler) Handler() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.findAllCategories)
	r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Post("/", h.createCategory)
	r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Delete("/", h.deleteMultipleCategories)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.findCategoryByID)
		r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Put("/", h.updateCategory)
		r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Delete("/", h.deleteCategory)
	})

	return r
}

func (h *CategoryHandler) findAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.app.FindCategories(r.Context())
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Categories retrieved successfully", categories)
}

func (h *CategoryHandler) findCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	category, err := h.app.FindCategory(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Category retrieved successfully", category)
}

func (h *CategoryHandler) createCategory(w http.ResponseWriter, r *http.Request) {
	var params usecases.CreateCategoryParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	category, err := h.app.CreateCategory(r.Context(), &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusCreated, "Category created successfully", category)
}

func (h *CategoryHandler) updateCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var params usecases.UpdateCategoryParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	updatedCategory, err := h.app.UpdateCategory(r.Context(), id, &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Category updated successfully", updatedCategory)
}

func (h *CategoryHandler) deleteCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	deletedResult, err := h.app.DeleteCategory(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Category deleted successfully", deletedResult)
}

func (h *CategoryHandler) deleteMultipleCategories(w http.ResponseWriter, r *http.Request) {
	var query usecases.DeleteCategoriesQuery
	if err := https.BindQuery(r, &query); err != nil {
		https.BadRequest(w, err)
		return
	}

	// Validate that at least one ID is provided
	if len(query.IDs) == 0 {
		https.BadRequest(w, errors.New("at least one id must be provided in ids parameter"))
		return
	}

	deletedResult, err := h.app.DeleteMultipleCategories(r.Context(), &query)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Categories deleted successfully", deletedResult)
}
