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
	})

	// Comment management routes
	// r.Route("/comments", func(r chi.Router) {
	// 	r.Get("/", h.findBlogComments)
	// 	r.With(h.authMiddleware.RequireAuth).Post("/", h.createBlogComment)
	// 	r.With(h.authMiddleware.RequireAuth, h.roleMiddleware.RequireAdmin).Delete("/", h.deleteMultipleBlogComments)

	// 	r.Route("/{commentId}", func(r chi.Router) {
	// 		r.Get("/", h.findBlogComment)
	// 		r.With(h.authMiddleware.RequireAuth).Put("/", h.updateBlogComment)
	// 		r.With(h.authMiddleware.RequireAuth).Delete("/", h.deleteBlogComment)
	// 	})
	// })

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
	// var query usecases.FindBlogCommentsQuery
	// if err := https.BindQuery(r, &query); err != nil {
	// 	https.BadRequest(w, err)
	// 	return
	// }

	// comments, err := h.app.FindBlogComments(r.Context(), &query)
	// if err != nil {
	// 	https.RespondWithFailure(w, err)
	// 	return
	// }

	// https.ResponseSuccess(w, http.StatusOK, "Comments retrieved successfully", comments)
}

func (h *BlogHandler) findBlogComment(w http.ResponseWriter, r *http.Request) {
	// commentID := chi.URLParam(r, "commentId")
	// comment, err := h.app.FindBlogComment(r.Context(), commentID)
	// if err != nil {
	// 	https.RespondWithFailure(w, err)
	// 	return
	// }

	// https.ResponseSuccess(w, http.StatusOK, "Comment retrieved successfully", comment)
}

func (h *BlogHandler) createBlogComment(w http.ResponseWriter, r *http.Request) {
	// var params usecases.CreateCommentParams
	// if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
	// 	https.BadRequest(w, err)
	// 	return
	// }

	// if err := h.validator.Struct(params); err != nil {
	// 	https.ValidationFailed(w, err)
	// 	return
	// }

	// comment, err := h.app.CreateBlogComment(r.Context(), &params)
	// if err != nil {
	// 	https.RespondWithFailure(w, err)
	// 	return
	// }

	// https.ResponseSuccess(w, http.StatusCreated, "Comment created successfully", comment)
}

func (h *BlogHandler) updateBlogComment(w http.ResponseWriter, r *http.Request) {
	// commentID := chi.URLParam(r, "commentId")
	// var params usecases.UpdateCommentParams
	// if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
	// 	https.BadRequest(w, err)
	// 	return
	// }

	// if err := h.validator.Struct(params); err != nil {
	// 	https.ValidationFailed(w, err)
	// 	return
	// }

	// updatedComment, err := h.app.UpdateBlogComment(r.Context(), commentID, &params)
	// if err != nil {
	// 	https.RespondWithFailure(w, err)
	// 	return
	// }

	// https.ResponseSuccess(w, http.StatusOK, "Comment updated successfully", updatedComment)
}

func (h *BlogHandler) deleteBlogComment(w http.ResponseWriter, r *http.Request) {
	// commentID := chi.URLParam(r, "commentId")
	// deletedResult, err := h.app.DeleteBlogComment(r.Context(), commentID)
	// if err != nil {
	// 	https.RespondWithFailure(w, err)
	// 	return
	// }

	// https.ResponseSuccess(w, http.StatusOK, "Comment deleted successfully", deletedResult)
}

func (h *BlogHandler) deleteMultipleBlogComments(w http.ResponseWriter, r *http.Request) {
	// var query usecases.DeleteCommentsQuery
	// if err := https.BindQuery(r, &query); err != nil {
	// 	https.BadRequest(w, err)
	// 	return
	// }

	// // Validate that at least one ID is provided
	// if len(query.IDs) == 0 {
	// 	https.BadRequest(w, errors.New("at least one id must be provided in ids parameter"))
	// 	return
	// }

	// deletedResult, err := h.app.DeleteMultipleBlogComments(r.Context(), &query)
	// if err != nil {
	// 	https.RespondWithFailure(w, err)
	// 	return
	// }

	// https.ResponseSuccess(w, http.StatusOK, "Comments deleted successfully", deletedResult)
}
