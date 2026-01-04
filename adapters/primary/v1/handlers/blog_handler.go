package handlers

import (
	"encoding/json"
	"errors"
	"hinsun-backend/adapters/shared/https"
	"hinsun-backend/adapters/shared/middlewares"
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/applications"
	"hinsun-backend/internal/domain/usecases"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type BlogHandler struct {
	app            applications.BlogAppService
	validator      *validator.Validate
	authMiddleware *middlewares.AuthMiddleware
	roleMiddleware *middlewares.RoleMiddleware
}

func NewBlogHandler(
	app applications.BlogAppService,
	validator *validator.Validate,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) *BlogHandler {
	return &BlogHandler{
		app:            app,
		validator:      validator,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
	}
}

func (h *BlogHandler) Handler() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.findBlogs)
	r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Post("/", h.createBlog)
	r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Delete("/", h.deleteMultipleBlogs)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.findBlog)
		r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Delete("/", h.deleteBlog)
		r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Put("/", h.updateBlog)

		r.Route("/comments", func(r chi.Router) {
			r.Get("/", h.findBlogComments)
			r.With(h.authMiddleware.RequireAuth).Post("/", h.addCommentToBlog)

			r.With(h.authMiddleware.RequireAuth).Put("/{commentId}", h.updateCommentOnBlog)
			r.With(h.authMiddleware.RequireAuth).Delete("/{commentId}", h.deleteCommentOnBlog)
		})
	})

	return r
}

func (h *BlogHandler) findBlogs(w http.ResponseWriter, r *http.Request) {
	var query usecases.FindBlogsQuery
	if err := https.BindQuery(r, &query); err != nil {
		https.BadRequest(w, err)
		return
	}

	blogs, err := h.app.FindBlogs(r.Context(), &query)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blogs retrieved successfully", blogs)
}

func (h *BlogHandler) findBlog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	blog, err := h.app.FindBlog(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blog retrieved successfully", blog)
}

func (h *BlogHandler) createBlog(w http.ResponseWriter, r *http.Request) {
	var params usecases.CreateBlogParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	blog, err := h.app.CreateBlog(r.Context(), &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusCreated, "Blog created successfully", blog)
}

func (h *BlogHandler) updateBlog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var params usecases.UpdateBlogParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	updatedBlog, err := h.app.UpdateBlog(r.Context(), id, &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blog updated successfully", updatedBlog)
}

func (h *BlogHandler) deleteBlog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	deletedResult, err := h.app.DeleteBlog(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blog deleted successfully", deletedResult)
}

func (h *BlogHandler) deleteMultipleBlogs(w http.ResponseWriter, r *http.Request) {
	var query usecases.DeleteBlogsQuery
	if err := https.BindQuery(r, &query); err != nil {
		https.BadRequest(w, err)
		return
	}

	// Validate that at least one ID is provided
	if len(query.IDs) == 0 {
		https.BadRequest(w, errors.New("at least one id must be provided in ids parameter"))
		return
	}

	deletedResult, err := h.app.DeleteMultipleBlogs(r.Context(), &query)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blogs deleted successfully", deletedResult)
}

// ================================== Comment Management Handlers =================================

func (h *BlogHandler) findBlogComments(w http.ResponseWriter, r *http.Request) {
	blogId := chi.URLParam(r, "id")
	comments, err := h.app.FindAllCommentsOnBlog(r.Context(), blogId)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Comments retrieved successfully", comments)
}

func (h *BlogHandler) addCommentToBlog(w http.ResponseWriter, r *http.Request) {
	blogId := chi.URLParam(r, "id")

	// Get claims from context (set by AuthMiddleware)
	claims, ok := middlewares.GetClaimsFromContext(r.Context())
	if !ok {
		https.RespondWithFailure(w, failure.NewAuthenticationFailure("authentication required"))
		return
	}

	var params usecases.AddCommentToBlogParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	comment, err := h.app.AddCommentToBlog(r.Context(), blogId, claims.AccountID, &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusCreated, "Comment created successfully", comment)
}

func (h *BlogHandler) updateCommentOnBlog(w http.ResponseWriter, r *http.Request) {
	blogId := chi.URLParam(r, "id")
	commentId := chi.URLParam(r, "commentId")

	// Get claims from context (set by AuthMiddleware)
	claims, ok := middlewares.GetClaimsFromContext(r.Context())
	if !ok {
		https.RespondWithFailure(w, failure.NewAuthenticationFailure("authentication required"))
		return
	}

	var params usecases.UpdateCommentOnBlogParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		https.BadRequest(w, err)
		return
	}

	if err := h.validator.Struct(params); err != nil {
		https.ValidationFailed(w, err)
		return
	}

	updatedComment, err := h.app.UpdateCommentOnBlog(r.Context(), blogId, claims.AccountID, commentId, &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Comment updated successfully", updatedComment)
}

func (h *BlogHandler) deleteCommentOnBlog(w http.ResponseWriter, r *http.Request) {
	blogId := chi.URLParam(r, "id")
	commentId := chi.URLParam(r, "commentId")

	// Get claims from context (set by AuthMiddleware)
	claims, ok := middlewares.GetClaimsFromContext(r.Context())
	if !ok {
		https.RespondWithFailure(w, failure.NewAuthenticationFailure("authentication required"))
		return
	}

	deletedResult, err := h.app.DeleteCommentOnBlog(r.Context(), blogId, claims.AccountID, commentId)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Comment deleted successfully", deletedResult)
}
