package models

// SetPassportDataErrorsRequest represents `setPassportDataErrors` request body.
type SetPassportDataErrorsRequest struct {
	// User identifier.
	UserID int32 `json:"user_id"`
	// A JSON-serialized array describing the errors.
	Errors []interface{} `json:"errors"`
}
