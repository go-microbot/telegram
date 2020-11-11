package api

import (
	"context"

	"github.com/go-microbot/telegram/models"
)

// Bot represents general Bot API interface.
// https://core.telegram.org/bots/api#available-methods.
type Bot interface {
	// https://core.telegram.org/bots/api#getme.
	GetMe(ctx context.Context) (*models.User, error)
	// https://core.telegram.org/bots/api#getupdates.
	GetUpdates(ctx context.Context, req GetUpdatesRequest) ([]models.Update, error)
}
