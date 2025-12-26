package v2

import (
	"github.com/go-chi/chi/v5"
)

type V2Routes struct{}

func NewV2Routes() *V2Routes {
	return &V2Routes{}
}

func (vr *V2Routes) RegisterRoutes() chi.Router {
	r := chi.NewRouter()
	return r
}
