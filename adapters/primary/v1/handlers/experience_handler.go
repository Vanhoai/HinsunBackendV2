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

type ExperienceHandler struct {
	app            applications.GlobalAppService
	validator      *validator.Validate
	authMiddleware *middlewares.AuthMiddleware
	roleMiddleware *middlewares.RoleMiddleware
}

func NewExperienceHandler(
	app applications.GlobalAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *ExperienceHandler {
	return &ExperienceHandler{
		app:            app,
		validator:      validator,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
	}
}

func (h *ExperienceHandler) Handler() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.findAllExperiences)

	r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Post("/", h.createExperience)
	r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Delete("/", h.deleteMultipleExperiences)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.findExperienceByID)

		r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Delete("/", h.deleteExperience)
		r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Put("/", h.updateExperience)
	})

	return r
}

func (h *ExperienceHandler) findAllExperiences(w http.ResponseWriter, r *http.Request) {
	experiences, err := h.app.FindExperiences(r.Context())
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Experiences retrieved successfully", experiences)
}

func (h *ExperienceHandler) findExperienceByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	experience, err := h.app.FindExperience(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Experience retrieved successfully", experience)
}

func (h *ExperienceHandler) createExperience(w http.ResponseWriter, r *http.Request) {
	var params usecases.CreateExperienceParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	experience, err := h.app.CreateExperience(r.Context(), &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusCreated, "Experience created successfully", experience)
}

func (h *ExperienceHandler) updateExperience(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var params usecases.UpdateExperienceParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	updatedExperience, err := h.app.UpdateExperience(r.Context(), id, &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Experience updated successfully", updatedExperience)
}

func (h *ExperienceHandler) deleteExperience(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	deletedResult, err := h.app.DeleteExperience(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Experience deleted successfully", deletedResult)
}

func (h *ExperienceHandler) deleteMultipleExperiences(w http.ResponseWriter, r *http.Request) {
	var query usecases.DeleteExperiencesQuery
	if err := https.BindQuery(r, &query); err != nil {
		https.BadRequest(w, err)
		return
	}

	// Validate that at least one ID is provided
	if len(query.IDs) == 0 {
		https.BadRequest(w, errors.New("at least one id must be provided in ids parameter"))
		return
	}

	deletedResult, err := h.app.DeleteMultipleExperiences(r.Context(), &query)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Experiences deleted successfully", deletedResult)
}
