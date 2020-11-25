package models

// AnswerCallbackQueryRequest represents `answerCallbackQuery` request body.
type AnswerCallbackQueryRequest struct {
	// Unique identifier for the query to be answered.
	CallbackQueryID string `json:"callback_query_id"`
	// Optional. Text of the notification. If not specified,
	// nothing will be shown to the user, 0-200 characters.
	Text string `json:"text,omitempty"`
	// Optional. If true, an alert will be shown by the client
	// instead of a notification at the top of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`
	// Optional. URL that will be opened by the user's client.
	// If you have created a Game and accepted the conditions via @Botfather,
	// specify the URL that opens your game — note that this will only work
	// if the query comes from a callback_game button.
	//
	// Otherwise, you may use links like t.me/your_bot?start=XXXX
	// that open your bot with a parameter.
	URL string `json:"url,omitempty"`
	// Optional. The maximum amount of time in seconds that the result of the callback query
	// may be cached client-side. Telegram apps will support caching starting
	// in version 3.14. Defaults to 0.
	CacheTime int32 `json:"cache_time,omitempty"`
}
