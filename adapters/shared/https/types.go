package https

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Payload any    `json:"payload,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

// FailureResponse extends APIResponse with error details
type FailureResponse struct {
	Code     string         `json:"code"`
	Message  string         `json:"message"`
	CausedBy string         `json:"causedBy,omitempty"`
	Details  map[string]any `json:"details,omitempty"`
}
