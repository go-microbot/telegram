package models

import "github.com/go-microbot/telegram/query"

// SendVenueRequest represents `sendVenue` request body.
type SendVenueRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Latitude of the venue.
	Latitude float32 `json:"latitude"`
	// Longitude of the venue.
	Longitude float32 `json:"longitude"`
	// Name of the venue.
	Title string `json:"title"`
	// Address of the venue.
	Address string `json:"address"`
	// Optional. Foursquare identifier of the venue.
	FoursquareID string `json:"foursquare_id,omitempty"`
	// Optional. Foursquare type of the venue, if known.
	// For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".
	FoursquareType string `json:"foursquare_type,omitempty"`
	// Optional. Google Places identifier of the venue.
	GooglePlaceID string `json:"google_place_id,omitempty"`
	// Optional. Google Places type of the venue.
	// See supported types: https://developers.google.com/places/web-service/supported_types.
	GooglePlaceType string `json:"google_place_type,omitempty"`
	// Optional. Sends the message
	// silently (https://telegram.org/blog/channels-2-0#silent-messages).
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message.
	ReplyToMessageID *int32 `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to remove reply keyboard
	// or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
