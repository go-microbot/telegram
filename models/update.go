package models

// Update represents an incoming update.
// https://core.telegram.org/bots/api#update.
type Update struct {
	// The update's unique identifier.
	// Update identifiers start from a certain positive number and increase sequentially.
	// This ID becomes especially handy if you're using
	// Webhooks (https://core.telegram.org/bots/api#setwebhook),
	// since it allows you to ignore repeated updates or to restore the correct update sequence,
	// should they get out of order. If there are no new updates for at least a week,
	// then identifier of the next update will be chosen randomly instead of sequentially.
	UpdateID int32 `json:"update_id"`
	// Optional. New incoming message of any kind — text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`
	// Optional. New version of a message that is known to the bot and was edited.
	EditedMessage *Message `json:"edited_message,omitempty"`
	// Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`
	// Optional. New version of a channel post that is known to the bot and was edited.
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	// Optional. New incoming inline (https://core.telegram.org/bots/api#inline-mode) query.
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	// Optional. The result of an inline query that was chosen by a user and sent
	// to their chat partner. Please see our documentation on the feedback collecting
	// (https://core.telegram.org/bots/inline#collecting-feedback) for details on how
	// to enable these updates for your bot.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	// Optional. New incoming callback query.
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	// Optional. New incoming shipping query. Only for invoices with flexible price.
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`
	// Optional. New incoming pre-checkout query. Contains full information about checkout.
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
	// Optional. New poll state.
	// Bots receive only updates about stopped polls and polls, which are sent by the bot.
	Poll *Poll `json:"poll,omitempty"`
	// Optional. A user changed their answer in a non-anonymous poll.
	// Bots receive new votes only in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
}
