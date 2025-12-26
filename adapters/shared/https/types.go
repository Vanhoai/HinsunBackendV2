package https

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ErrorDetail provides additional error information
// type ErrorDetail struct {
// 	Field   string `json:"field,omitempty"`
// 	Message string `json:"message"`
// }

// FailureResponse extends APIResponse with error details
type FailureResponse struct {
	Code     string         `json:"code"`
	Message  string         `json:"message"`
	CausedBy string         `json:"causedBy,omitempty"`
	Details  map[string]any `json:"details,omitempty"`
}
