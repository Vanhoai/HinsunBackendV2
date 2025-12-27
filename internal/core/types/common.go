package types

type DeletedResult struct {
	RowsAffected int `json:"rowsAffected"`      // number of rows affected by the delete operation
	Payload      any `json:"payload,omitempty"` // optional payload with additional information
}
