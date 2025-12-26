package v1

import (
	"hinsun-backend/adapters/primary/v1/handlers"

	"github.com/go-chi/chi/v5"
)

type V1Routes struct {
	authHandler       *handlers.AuthHandler
	experienceHandler *handlers.ExperienceHandler
}

func NewV1Routes(
	authHandler *handlers.AuthHandler,
	experienceHandler *handlers.ExperienceHandler,
) *V1Routes {
	return &V1Routes{
		authHandler:       authHandler,
		experienceHandler: experienceHandler,
	}
}

func (vr *V1Routes) RegisterRoutes() chi.Router {
	r := chi.NewRouter()

	r.Mount("/auth", vr.authHandler.Handler())
	r.Mount("/experiences", vr.experienceHandler.Handler())

	return r
}
