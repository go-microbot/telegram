package models

import "github.com/go-microbot/telegram/models"

// SetMyCommandsRequest represents `setMyCommands` request body.
type SetMyCommandsRequest struct {
	// A JSON-serialized list of bot commands to be set as the list of the bot's commands.
	// At most 100 commands can be specified.
	Commands []models.BotCommand `json:"commands"`
}
