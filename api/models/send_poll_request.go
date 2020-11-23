package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendPollRequest represents `sendPoll` request body.
type SendPollRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Poll question, 1-300 characters.
	Question string `json:"question"`
	// A JSON-serialized list of answer options, 2-10 strings 1-100 characters each.
	Options []string `json:"options"`
	// Optional. True, if the poll needs to be anonymous, defaults to True.
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// Optional. Poll type, "quiz" or "regular", defaults to "regular".
	Type string `json:"type,omitempty"`
	// Optional. True, if the poll allows multiple answers,
	// ignored for polls in quiz mode, defaults to False.
	AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`
	// Optional. 0-based identifier of the correct answer option, required for polls in quiz mode.
	CorrectOptionID int32 `json:"correct_option_id,omitempty"`
	// Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp
	// icon in a quiz-style poll, 0-200 characters with at most 2 line feeds
	// after entities parsing.
	Explanation string `json:"explanation,omitempty"`
	// Optional. Mode for parsing entities in the explanation.
	// See formatting options (https://core.telegram.org/bots/api#formatting-options)
	// for more details.
	ExplanationParseMode string `json:"explanation_parse_mode,omitempty"`
	// Optional. List of special entities that appear in the poll explanation,
	// which can be specified instead of parse_mode.
	ExplanationEntities []models.MessageEntity `json:"explanation_entities,omitempty"`
	// Optional. Amount of time in seconds the poll will be active after creation, 5-600.
	// Can't be used together with close_date.
	OpenPeriod int32 `json:"open_period,omitempty"`
	// Optional. Point in time (Unix timestamp) when the poll will be automatically closed.
	// Must be at least 5 and no more than 600 seconds in the future.
	// Can't be used together with open_period.
	CloseDate int32 `json:"close_date,omitempty"`
	// Optional. Pass True, if the poll needs to be immediately closed.
	// This can be useful for poll preview.
	IsClosed bool `json:"is_closed,omitempty"`
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
