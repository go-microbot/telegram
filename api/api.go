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
	// https://core.telegram.org/bots/api#logout.
	Logout(ctx context.Context) error
	// https://core.telegram.org/bots/api#close.
	Close(ctx context.Context) error
	// https://core.telegram.org/bots/api#sendphoto.
	SendPhoto(ctx context.Context, req apiModels.SendPhotoRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#setchatphoto.
	SetChatPhoto(ctx context.Context, req apiModels.SetChatPhotoRequest) error
	// https://core.telegram.org/bots/api#setchattitle.
	SetChatTitle(ctx context.Context, req apiModels.SetChatTitleRequest) error
	// https://core.telegram.org/bots/api#getchat.
	GetChat(ctx context.Context, req apiModels.ChatID) (*models.Chat, error)
	// https://core.telegram.org/bots/api#leavechat.
	LeaveChat(ctx context.Context, req apiModels.ChatID) error
}
