package https

import "net/http"

// MethodHandler helps route different HTTP methods to different handlers
type MethodHandler map[string]http.HandlerFunc

// ServeHTTP implements http.Handler interface
func (m MethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := m[r.Method]; ok {
		handler(w, r)
	} else {
		// Return 405 Method Not Allowed
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method not allowed"}`))
	}
}

// Helper to set JSON response
func JsonResponse(w http.ResponseWriter, status int, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(data))
}
