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

type BlogHandler struct {
	globalAppService applications.GlobalAppService
	validator        *validator.Validate
}

func NewBlogHandler(globalAppService applications.GlobalAppService, validator *validator.Validate) *BlogHandler {
	return &BlogHandler{
		globalAppService: globalAppService,
		validator:        validator,
	}
}

func (h *BlogHandler) Handler() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.findAllBlogs)
	r.Post("/", h.createBlog)
	r.Delete("/", h.deleteMultipleBlogs)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.findBlogByID)
		r.Delete("/", h.deleteBlog)
		r.Put("/", h.updateBlog)
		r.Post("/views", h.incrementBlogViews)
		r.Post("/favorites", h.incrementBlogFavorites)
		r.Delete("/favorites", h.decrementBlogFavorites)
	})

	r.Get("/author/{authorId}", h.findBlogsByAuthor)

	return r
}

func (h *BlogHandler) findAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := h.globalAppService.FindBlogs(r.Context())
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blogs retrieved successfully", blogs)
}

func (h *BlogHandler) findBlogByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	blog, err := h.globalAppService.FindBlog(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blog retrieved successfully", blog)
}

func (h *BlogHandler) findBlogsByAuthor(w http.ResponseWriter, r *http.Request) {
	authorId := chi.URLParam(r, "authorId")
	blogs, err := h.globalAppService.FindBlogsByAuthor(r.Context(), authorId)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blogs retrieved successfully", blogs)
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

	blog, err := h.globalAppService.CreateBlog(r.Context(), &params)
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

	updatedBlog, err := h.globalAppService.UpdateBlog(r.Context(), id, &params)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blog updated successfully", updatedBlog)
}

func (h *BlogHandler) deleteBlog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	deletedResult, err := h.globalAppService.DeleteBlog(r.Context(), id)
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

	deletedResult, err := h.globalAppService.DeleteMultipleBlogs(r.Context(), &query)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blogs deleted successfully", deletedResult)
}

func (h *BlogHandler) incrementBlogViews(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.globalAppService.IncrementBlogViews(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blog views incremented successfully", nil)
}

func (h *BlogHandler) incrementBlogFavorites(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.globalAppService.IncrementBlogFavorites(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blog favorites incremented successfully", nil)
}

func (h *BlogHandler) decrementBlogFavorites(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.globalAppService.DecrementBlogFavorites(r.Context(), id)
	if err != nil {
		https.RespondWithFailure(w, err)
		return
	}

	https.ResponseSuccess(w, http.StatusOK, "Blog favorites decremented successfully", nil)
}
