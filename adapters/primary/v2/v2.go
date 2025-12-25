package v2

import "net/http"

func RegisterV2Routes() http.Handler {
	mux := http.NewServeMux()
	return mux
}
