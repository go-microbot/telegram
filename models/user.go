package models

// User represents a Telegram user or bot model.
// https://core.telegram.org/bots/api#user.
type User struct {
	// Unique identifier for this user or bot.
	ID int64 `json:"id"`
	// True, if this user is a bot.
	IsBot bool `json:"is_bot"`
	// User's or bot's first name.
	FirstName string `json:"first_name"`
	// Optional. User's or bot's last name.
	LastName string `json:"last_name,omitempty"`
	// Optional. User's or bot's username.
	Username string `json:"username,omitempty"`
	// Optional. IETF (https://en.wikipedia.org/wiki/IETF_language_tag)
	// language tag of the user's language.
	LanguageCode string `json:"language_code,omitempty"`
	// Optional. True, if the bot can be invited to groups.
	// Returned only in getMe (https://core.telegram.org/bots/api#getme).
	CanJoinGroups *bool `json:"can_join_groups,omitempty"`
	// Optional. True, if privacy mode (https://core.telegram.org/bots#privacy-mode)
	// is disabled for the bot.
	// Returned only in getMe (https://core.telegram.org/bots/api#getme).
	CanReadAllGroupMessages *bool `json:"can_read_all_group_messages,omitempty"`
	// Optional. True, if the bot supports inline queries.
	// Returned only in getMe (https://core.telegram.org/bots/api#getme).
	SupportsInlineQueries *bool `json:"supports_inline_queries,omitempty"`
}
