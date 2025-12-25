package v1

import (
	"hinsun-backend/adapters/shared/https"
	"net/http"
)

func ExperienceHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", https.MethodHandler{
		http.MethodGet:  FindAllExperiences,
		http.MethodPost: CreateExperience,
	})

	return mux
}

// GET /api/v1/experiences - Find all experiences
func FindAllExperiences(w http.ResponseWriter, r *http.Request) {
	https.JsonResponse(w, http.StatusOK, `{"accounts": [], "message": "List all accounts 📋"}`)
}

// POST /api/v1/experiences - Create a new experience
func CreateExperience(w http.ResponseWriter, r *http.Request) {
	https.JsonResponse(w, http.StatusCreated, `{"account": {}, "message": "Account created successfully 🎉"}`)
}
