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
	// https://core.telegram.org/bots/api#setchatpermissions
	SetChatPermissions(ctx context.Context, req apiModels.SetChatPermissionsRequest) error
	// https://core.telegram.org/bots/api#setchatphoto.
	SetChatPhoto(ctx context.Context, req apiModels.SetChatPhotoRequest) error
	// https://core.telegram.org/bots/api#setchattitle.
	SetChatTitle(ctx context.Context, req apiModels.SetChatTitleRequest) error
	// https://core.telegram.org/bots/api#getchat.
	GetChat(ctx context.Context, req apiModels.ChatID) (*models.Chat, error)
	// https://core.telegram.org/bots/api#leavechat.
	LeaveChat(ctx context.Context, req apiModels.ChatID) error
	// https://core.telegram.org/bots/api#forwardmessage.
	ForwardMessage(ctx context.Context, req apiModels.ForwardMessageRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#setmycommands.
	SetMyCommands(ctx context.Context, req apiModels.SetMyCommandsRequest) error
	// https://core.telegram.org/bots/api#getmycommands.
	GetMyCommands(ctx context.Context) ([]models.BotCommand, error)
	// https://core.telegram.org/bots/api#setchatdescription.
	SetChatDescription(ctx context.Context, req apiModels.SetChatDescriptionRequest) error
	// https://core.telegram.org/bots/api#deletechatphoto.
	DeleteChatPhoto(ctx context.Context, req apiModels.ChatID) error
	// https://core.telegram.org/bots/api#sendlocation.
	SendLocation(ctx context.Context, req apiModels.SendLocationRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#getfile.
	GetFile(ctx context.Context, req apiModels.FileID) (*models.File, error)
	// https://core.telegram.org/bots/api#pinchatmessage.
	PinChatMessage(ctx context.Context, req apiModels.PinChatMessageRequest) error
	// https://core.telegram.org/bots/api#unpinchatmessage.
	UnpinChatMessage(ctx context.Context, req apiModels.UnpinChatMessageRequest) error
	// https://core.telegram.org/bots/api#unpinallchatmessages.
	UnpinAllChatMessages(ctx context.Context, req apiModels.ChatID) error
	// https://core.telegram.org/bots/api#getchatmember.
	GetChatMember(ctx context.Context, req apiModels.GetChatMemberRequest) (*models.ChatMember, error)
	// https://core.telegram.org/bots/api#exportchatinvitelink.
	ExportChatInviteLink(ctx context.Context, req apiModels.ChatID) (string, error)
	// https://core.telegram.org/bots/api#setchatadministratorcustomtitle.
	SetChatAdministratorCustomTitle(ctx context.Context, req apiModels.SetChatAdminCustomTitleRequest) error
	// https://core.telegram.org/bots/api#getchatadministrators.
	GetChatAdministrators(ctx context.Context, req apiModels.ChatID) ([]models.ChatMember, error)
	// https://core.telegram.org/bots/api#getchatmemberscount.
	GetChatMembersCount(ctx context.Context, req apiModels.ChatID) (int32, error)
	// https://core.telegram.org/bots/api#sendaudio.
	SendAudio(ctx context.Context, req apiModels.SendAudioRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#senddocument.
	SendDocument(ctx context.Context, req apiModels.SendDocumentRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#sendvideo.
	SendVideo(ctx context.Context, req apiModels.SendVideoRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#sendanimation.
	SendAnimation(ctx context.Context, req apiModels.SendAnimationRequest) (*models.Message, error)
	// https://core.telegram.org/bots/api#sendvoice.
	SendVoice(ctx context.Context, req apiModels.SendVoiceRequest) (*models.Message, error)
}
