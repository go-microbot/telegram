package models

// MessageID represents a unique message identifier.
// https://core.telegram.org/bots/api#messageid.
type MessageID struct {
	// Unique message identifier.
	MessageID int32 `json:"message_id"`
}
