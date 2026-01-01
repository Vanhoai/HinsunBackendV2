package v1

import (
	"hinsun-backend/adapters/primary/v1/handlers"

	"github.com/go-chi/chi/v5"
)

type V1Routes struct {
	authHandler       *handlers.AuthHandler
	experienceHandler *handlers.ExperienceHandler
	blogHandler       *handlers.BlogHandler
	projectHandler    *handlers.ProjectHandler
	accountHandler    *handlers.AccountHandler
}

func NewV1Routes(
	authHandler *handlers.AuthHandler,
	experienceHandler *handlers.ExperienceHandler,
	blogHandler *handlers.BlogHandler,
	projectHandler *handlers.ProjectHandler,
	accountHandler *handlers.AccountHandler,
) *V1Routes {
	return &V1Routes{
		authHandler:       authHandler,
		experienceHandler: experienceHandler,
		blogHandler:       blogHandler,
		projectHandler:    projectHandler,
		accountHandler:    accountHandler,
	}
}

func (vr *V1Routes) RegisterRoutes() chi.Router {
	r := chi.NewRouter()

	r.Mount("/auth", vr.authHandler.Handler())
	r.Mount("/experiences", vr.experienceHandler.Handler())
	r.Mount("/blogs", vr.blogHandler.Handler())
	r.Mount("/projects", vr.projectHandler.Handler())
	r.Mount("/accounts", vr.accountHandler.Handler())

	return r
}
