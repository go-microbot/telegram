package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// EditMessageLiveLocationRequest represents `editMessageLiveLocation` request body.
type EditMessageLiveLocationRequest struct {
	// Optional. Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id,omitempty" json:"-"`
	// Optional. Required if inline_message_id is not specified.
	// Identifier of the message to edit.
	MessageID query.ParamInt `query:"message_id,omitempty" json:"-"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message.
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// Latitude of new location.
	Latitude float32 `json:"latitude"`
	// Longitude of new location.
	Longitude float32 `json:"longitude"`
	// Optional. The radius of uncertainty for the location, measured in meters; 0-1500.
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`
	// Optional. Direction in which the user is moving, in degrees.
	// Must be between 1 and 360 if specified.
	Heading int32 `json:"heading,omitempty"`
	// Optional. Maximum distance for proximity alerts about approaching another chat member,
	// in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int32 `json:"proximity_alert_radius,omitempty"`
	// Optional. A JSON-serialized object for a new
	// inline keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	ReplyMarkup *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
