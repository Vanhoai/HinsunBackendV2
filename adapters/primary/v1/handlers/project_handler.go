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

type ProjectHandler struct {
	globalAppService applications.GlobalAppService
	validator        *validator.Validate
}

func NewProjectHandler(globalAppService applications.GlobalAppService, validator *validator.Validate) *ProjectHandler {
	return &ProjectHandler{
		globalAppService: globalAppService,
		validator:        validator,
	}
}

func (h *ProjectHandler) Handler() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.findAllProjects)
	r.Post("/", h.createProject)
	r.Delete("/", h.deleteMultipleProjects)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.findProjectByID)
		r.Delete("/", h.deleteProject)
		r.Put("/", h.updateProject)
	})

	return r
}

func (h *ProjectHandler) findAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.globalAppService.FindProjects(r.Context())
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Projects retrieved successfully", projects)
}

func (h *ProjectHandler) findProjectByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	project, err := h.globalAppService.FindProject(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Project retrieved successfully", project)
}

func (h *ProjectHandler) createProject(w http.ResponseWriter, r *http.Request) {
	var params usecases.CreateProjectParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	project, err := h.globalAppService.CreateProject(r.Context(), &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusCreated, "Project created successfully", project)
}

func (h *ProjectHandler) updateProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var params usecases.UpdateProjectParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	updatedProject, err := h.globalAppService.UpdateProject(r.Context(), id, &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Project updated successfully", updatedProject)
}

func (h *ProjectHandler) deleteProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	deletedResult, err := h.globalAppService.DeleteProject(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Project deleted successfully", deletedResult)
}

func (h *ProjectHandler) deleteMultipleProjects(w http.ResponseWriter, r *http.Request) {
	var query usecases.DeleteProjectsQuery
	if err := https.BindQuery(r, &query); err != nil {
		https.BadRequest(w, err)
		return
	}

	// Validate that at least one ID is provided
	if len(query.IDs) == 0 {
		https.BadRequest(w, errors.New("at least one id must be provided in ids parameter"))
		return
	}

	deletedResult, err := h.globalAppService.DeleteMultipleProjects(r.Context(), &query)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Projects deleted successfully", deletedResult)
}
