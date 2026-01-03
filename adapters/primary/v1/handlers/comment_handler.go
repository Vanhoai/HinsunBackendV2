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

type CommentHandler struct {
	app            applications.CommentAppService
	validator      *validator.Validate
	authMiddleware *middlewares.AuthMiddleware
	roleMiddleware *middlewares.RoleMiddleware
}

func NewCommentHandler(
	app applications.CommentAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *CommentHandler {
	return &CommentHandler{
		app:            app,
		validator:      validator,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
	}
}

func (h *CommentHandler) Handler() chi.Router {
	r := chi.NewRouter()
	r.Use(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin)

	r.Get("/", h.findComments)
	r.Post("/", h.createComment)
	r.Delete("/", h.deleteMultipleComments)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.findComment)
		r.Put("/", h.updateComment)
		r.Delete("/", h.deleteComment)
	})

	return r
}

func (h *CommentHandler) findComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.app.FindComments(r.Context())
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Comments retrieved successfully", comments)
}

func (h *CommentHandler) findComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	comment, err := h.app.FindComment(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Comment retrieved successfully", comment)
}

func (h *CommentHandler) createComment(w http.ResponseWriter, r *http.Request) {
	var params usecases.CreateCommentParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	comment, err := h.app.CreateComment(r.Context(), &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusCreated, "Comment created successfully", comment)
}

func (h *CommentHandler) updateComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var params usecases.UpdateCommentParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	updatedComment, err := h.app.UpdateComment(r.Context(), id, &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Comment updated successfully", updatedComment)
}

func (h *CommentHandler) deleteComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	deletedResult, err := h.app.DeleteComment(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Comment deleted successfully", deletedResult)
}

func (h *CommentHandler) deleteMultipleComments(w http.ResponseWriter, r *http.Request) {
	var query usecases.DeleteCommentsQuery
	if err := https.BindQuery(r, &query); err != nil {
		https.BadRequest(w, err)
		return
	}

	// Validate that at least one ID is provided
	if len(query.IDs) == 0 {
		https.BadRequest(w, errors.New("at least one id must be provided in ids parameter"))
		return
	}

	deletedResult, err := h.app.DeleteMultipleComments(r.Context(), &query)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Comments deleted successfully", deletedResult)
}
