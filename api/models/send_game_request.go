package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendGameRequest represents `sendGame` request body.
type SendGameRequest struct {
	// Unique identifier for the target chat.
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Short name of the game, serves as the unique identifier for the game.
	GameShortName string `json:"game_short_name"`
	// Optional. Sends the message
	// silently (https://telegram.org/blog/channels-2-0#silent-messages).
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message.
	ReplyToMessageID *int32 `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even
	// if the specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. A JSON-serialized object for an
	// inline keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	// If empty, one 'Play game_title' button will be shown.
	// If not empty, the first button must launch the game.
	ReplyMarkup *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
