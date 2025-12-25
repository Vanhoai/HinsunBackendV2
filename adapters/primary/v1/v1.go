package v1

import "net/http"

func RegisterV1Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/auth/", http.StripPrefix("/api/v1/auth", AuthHandler()))
	mux.Handle("/api/v1/experiences/", http.StripPrefix("/api/v1/experiences", ExperienceHandler()))

	return mux
}
