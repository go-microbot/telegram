package models

import "github.com/go-microbot/telegram/models"

// AnswerShippingQueryRequest represnts `answerShippingQuery` request body.
type AnswerShippingQueryRequest struct {
	// Unique identifier for the query to be answered.
	ShippingQueryID string `json:"shipping_query_id"`
	// Specify True if delivery to the specified address is possible and False
	// if there are any problems
	// (for example, if delivery to the specified address is not possible).
	OK bool `json:"ok"`
	// Optional. Required if ok is True. A JSON-serialized array of available shipping options.
	ShippingOptions []models.ShippingOption `json:"shipping_options,omitempty"`
	// Optional. Required if ok is False.
	// Error message in human readable form that explains why it is impossible
	// to complete the order (e.g. "Sorry, delivery to your desired address is unavailable").
	// Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}
