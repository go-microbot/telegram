package models

// Contact represents a phone contact.
// https://core.telegram.org/bots/api#contact.
type Contact struct {
	// Contact's phone number.
	PhoneNumber string `json:"phone_number"`
	// Contact's first name.
	FirstName string `json:"first_name"`
	// Optional. Contact's last name.
	LastName string `json:"last_name,omitempty"`
	// Optional. Contact's user identifier in Telegram.
	UserID int32 `json:"user_id,omitempty"`
	// Optional. Additional data about the contact in the form of a vCard
	// (https://en.wikipedia.org/wiki/VCard).
	VCard string `json:"vcard,omitempty"`
}
