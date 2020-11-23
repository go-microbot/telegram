package models

import "github.com/go-microbot/telegram/query"

// SendContactRequest represents `sendContact` request body.
type SendContactRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Contact's phone number.
	PhoneNumber string `json:"phone_number"`
	// Contact's first name.
	FirstName string `json:"first_name"`
	// Optional. Contact's last name.
	LastName string `json:"last_name,omitempty"`
	// Optional. Additional data about the contact in the form of
	// a vCard (https://en.wikipedia.org/wiki/VCard), 0-2048 bytes.
	VCard string `json:"vcard,omitempty"`
	// Optional. Sends the message
	// silently (https://telegram.org/blog/channels-2-0#silent-messages).
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message.
	ReplyToMessageID *int32 `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the
	// specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to remove keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
