package models

import (
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
)

// SetChatPhotoRequest represents `setChatPhoto` body.
type SetChatPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id"`
	// New chat photo, uploaded using multipart/form-data.
	Photo form.PartFile `form:"photo"`
}
