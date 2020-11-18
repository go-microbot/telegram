package models

import "github.com/go-microbot/telegram/query"

// SendLocationRequest represents `sendLocation` request body.
type SendLocationRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Latitude of the location.
	Latitude float32 `json:"latitude"`
	// Longitude of the location.
	Longitude float32 `json:"longitude"`
	// Optional. The radius of uncertainty for the location, measured in meters; 0-1500.
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`
	// Optional. Period in seconds for which the location will be updated
	// (see Live Locations (https://telegram.org/blog/live-locations),
	// should be between 60 and 86400.
	LivePeriod int32 `json:"live_period,omitempty"`
	// Optional. For live locations, a direction in which the user is moving, in degrees.
	// Must be between 1 and 360 if specified.
	Heading int32 `json:"heading,omitempty"`
	// Optional. For live locations, a maximum distance for proximity alerts about
	// approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int32 `json:"proximity_alert_radius,omitempty"`
	// Optional. Sends the message
	// silently (https://telegram.org/blog/channels-2-0#silent-messages).
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message.
	ReplyToMessageID *int32 `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent
	// even if the specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to remove reply keyboard
	// or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
