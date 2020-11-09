package api

import (
	"context"

	"github.com/go-microbot/telegram/models"
)

// ServiceAPI represents general API interface.
// https://core.telegram.org/bots/api#available-methods.
type ServiceAPI interface {
	GetMe(ctx context.Context) (*models.User, error)
}
