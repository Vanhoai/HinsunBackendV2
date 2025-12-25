package v1

import (
	"hinsun-backend/adapters/shared/https"
	"net/http"
)

func AuthHandler() http.Handler {
	authMux := http.NewServeMux()

	authMux.HandleFunc("/email-auth", authEmail)
	authMux.HandleFunc("/oauth2", oauth2)

	return authMux
}

// authEmail godoc
// @Summary      Sign in with email and password
// @Description  Authenticate an account using email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      AuthEmailParams  true  "Login credentials"
// @Success      200  {object}  AuthResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      401  {object}  ErrorResponse
// @Router       /auth [post]
func authEmail(w http.ResponseWriter, r *http.Request) {
	https.JsonResponse(w, http.StatusOK, `{"message": "You All Signed In Mister Wick 🧘🏽🧘🏽🧘🏽"}`)
}

func oauth2(w http.ResponseWriter, r *http.Request) {
	https.JsonResponse(w, http.StatusOK, `{"message": "OAuth2 endpoint"}`)
}
