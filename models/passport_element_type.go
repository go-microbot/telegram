package models

// There are available passport element types.
const (
	PassportElementTypePersonalDetails  PassportElementType = "personal_details"
	PassportElementTypePassport         PassportElementType = "passport"
	PassportElementTypeDriverLicense    PassportElementType = "driver_license"
	PassportElementTypeIdentityCard     PassportElementType = "identity_card"
	PassportElementTypeInternalPassport PassportElementType = "internal_passport"
	PassportElementTypeAddress          PassportElementType = "address"
)

// PassportElementType represents passport element type.
type PassportElementType string
