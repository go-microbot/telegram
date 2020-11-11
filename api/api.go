package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
)

// Bot represents general Bot API interface.
// https://core.telegram.org/bots/api#available-methods.
type Bot interface {
	// https://core.telegram.org/bots/api#getme.
	GetMe(ctx context.Context) (*models.User, error)
	// https://core.telegram.org/bots/api#getupdates.
	GetUpdates(ctx context.Context, req apiModels.GetUpdatesRequest) ([]models.Update, error)
	GetPollUpdates(ctx context.Context, req apiModels.GetUpdatesRequest, client *http.Client) ([]models.Update, error)
	// https://core.telegram.org/bots/api#setwebhook.
	SetWebhook(ctx context.Context, req apiModels.SetWebhookRequest) error
	// https://core.telegram.org/bots/api#getwebhookinfo.
	GetWebhookInfo(ctx context.Context) (*models.WebhookInfo, error)
	// https://core.telegram.org/bots/api#deletewebhook.
	DeleteWebhook(ctx context.Context) error
	// https://core.telegram.org/bots/api#sendmessage.
	SendMessage(ctx context.Context, req apiModels.SendMessageRequest) (*models.Message, error)
}
