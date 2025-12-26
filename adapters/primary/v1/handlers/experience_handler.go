package handlers

import (
	"encoding/json"
	"hinsun-backend/adapters/shared/https"
	"hinsun-backend/internal/domain/applications"
	"hinsun-backend/internal/domain/usecases"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type ExperienceHandler struct {
	globalAppService applications.GlobalAppService
}

func NewExperienceHandler(globalAppService applications.GlobalAppService) *ExperienceHandler {
	return &ExperienceHandler{
		globalAppService: globalAppService,
	}
}

func (h *ExperienceHandler) Handler() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.findAllExperiencesRoute)
	r.Post("/", h.createExperienceRoute)

	return r
}

func (h *ExperienceHandler) findAllExperiencesRoute(w http.ResponseWriter, r *http.Request) {
	experiences, err := h.globalAppService.FindExperiences(r.Context())
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Experiences retrieved successfully", experiences)
}

func (h *ExperienceHandler) createExperienceRoute(w http.ResponseWriter, r *http.Request) {
	var params usecases.CreateExperienceParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	experience, err := h.globalAppService.CreateExperience(r.Context(), &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusCreated, "Experience created successfully", experience)
}
