package models

// AnswerPreCheckoutQueryRequest represents `answerPreCheckoutQuery` request body.
type AnswerPreCheckoutQueryRequest struct {
	// Unique identifier for the query to be answered.
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`
	// Specify True if everything is alright (goods are available, etc.)
	// and the bot is ready to proceed with the order. Use False if there are any problems.
	OK bool `json:"ok"`
	// Optional. Required if ok is False.
	// Error message in human readable form that explains the reason for failure
	// to proceed with the checkout (e.g. "Sorry, somebody just bought the last of
	// our amazing black T-shirts while you were busy filling out your payment details.
	// Please choose a different color or garment!").
	// Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}
