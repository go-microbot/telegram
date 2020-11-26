package models

// PassportElementErrorDataField represents an issue in one of the data fields
// that was provided by the user. The error is considered resolved when the field's value changes.
// https://core.telegram.org/bots/api#passportelementerrordatafield.
type PassportElementErrorDataField struct {
	// Error source, must be data.
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the error,
	// one of "personal_details", "passport", "driver_license", "identity_card",
	// "internal_passport", "address".
	Type PassportElementType `json:"type"`
	// Name of the data field which has the error.
	FieldName string `json:"field_name"`
	// Base64-encoded data hash.
	DataHash string `json:"data_hash"`
	// Error message.
	Message string `json:"message"`
}
