// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	models "github.com/go-microbot/telegram/api/models"

	telegrammodels "github.com/go-microbot/telegram/models"
)

// Bot is an autogenerated mock type for the Bot type
type Bot struct {
	mock.Mock
}

// Close provides a mock function with given fields: ctx
func (_m *Bot) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CopyMessage provides a mock function with given fields: ctx, req
func (_m *Bot) CopyMessage(ctx context.Context, req models.CopyMessageRequest) (*telegrammodels.MessageID, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.MessageID
	if rf, ok := ret.Get(0).(func(context.Context, models.CopyMessageRequest) *telegrammodels.MessageID); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.MessageID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.CopyMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNewStickerSet provides a mock function with given fields: ctx, req
func (_m *Bot) CreateNewStickerSet(ctx context.Context, req models.CreateNewStickerSetRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.CreateNewStickerSetRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteChatPhoto provides a mock function with given fields: ctx, req
func (_m *Bot) DeleteChatPhoto(ctx context.Context, req models.ChatID) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ChatID) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMessage provides a mock function with given fields: ctx, req
func (_m *Bot) DeleteMessage(ctx context.Context, req models.DeleteMessageRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.DeleteMessageRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWebhook provides a mock function with given fields: ctx
func (_m *Bot) DeleteWebhook(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExportChatInviteLink provides a mock function with given fields: ctx, req
func (_m *Bot) ExportChatInviteLink(ctx context.Context, req models.ChatID) (string, error) {
	ret := _m.Called(ctx, req)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, models.ChatID) string); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.ChatID) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForwardMessage provides a mock function with given fields: ctx, req
func (_m *Bot) ForwardMessage(ctx context.Context, req models.ForwardMessageRequest) (*telegrammodels.Message, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.Message
	if rf, ok := ret.Get(0).(func(context.Context, models.ForwardMessageRequest) *telegrammodels.Message); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.ForwardMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChat provides a mock function with given fields: ctx, req
func (_m *Bot) GetChat(ctx context.Context, req models.ChatID) (*telegrammodels.Chat, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.Chat
	if rf, ok := ret.Get(0).(func(context.Context, models.ChatID) *telegrammodels.Chat); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.Chat)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.ChatID) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChatAdministrators provides a mock function with given fields: ctx, req
func (_m *Bot) GetChatAdministrators(ctx context.Context, req models.ChatID) ([]telegrammodels.ChatMember, error) {
	ret := _m.Called(ctx, req)

	var r0 []telegrammodels.ChatMember
	if rf, ok := ret.Get(0).(func(context.Context, models.ChatID) []telegrammodels.ChatMember); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]telegrammodels.ChatMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.ChatID) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChatMember provides a mock function with given fields: ctx, req
func (_m *Bot) GetChatMember(ctx context.Context, req models.GetChatMemberRequest) (*telegrammodels.ChatMember, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.ChatMember
	if rf, ok := ret.Get(0).(func(context.Context, models.GetChatMemberRequest) *telegrammodels.ChatMember); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.ChatMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.GetChatMemberRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChatMembersCount provides a mock function with given fields: ctx, req
func (_m *Bot) GetChatMembersCount(ctx context.Context, req models.ChatID) (int32, error) {
	ret := _m.Called(ctx, req)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, models.ChatID) int32); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.ChatID) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFile provides a mock function with given fields: ctx, req
func (_m *Bot) GetFile(ctx context.Context, req models.FileID) (*telegrammodels.File, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.File
	if rf, ok := ret.Get(0).(func(context.Context, models.FileID) *telegrammodels.File); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.FileID) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMe provides a mock function with given fields: ctx
func (_m *Bot) GetMe(ctx context.Context) (*telegrammodels.User, error) {
	ret := _m.Called(ctx)

	var r0 *telegrammodels.User
	if rf, ok := ret.Get(0).(func(context.Context) *telegrammodels.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMyCommands provides a mock function with given fields: ctx
func (_m *Bot) GetMyCommands(ctx context.Context) ([]telegrammodels.BotCommand, error) {
	ret := _m.Called(ctx)

	var r0 []telegrammodels.BotCommand
	if rf, ok := ret.Get(0).(func(context.Context) []telegrammodels.BotCommand); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]telegrammodels.BotCommand)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPollUpdates provides a mock function with given fields: ctx, req, client
func (_m *Bot) GetPollUpdates(ctx context.Context, req models.GetUpdatesRequest, client *http.Client) ([]telegrammodels.Update, error) {
	ret := _m.Called(ctx, req, client)

	var r0 []telegrammodels.Update
	if rf, ok := ret.Get(0).(func(context.Context, models.GetUpdatesRequest, *http.Client) []telegrammodels.Update); ok {
		r0 = rf(ctx, req, client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]telegrammodels.Update)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.GetUpdatesRequest, *http.Client) error); ok {
		r1 = rf(ctx, req, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpdates provides a mock function with given fields: ctx, req
func (_m *Bot) GetUpdates(ctx context.Context, req models.GetUpdatesRequest) ([]telegrammodels.Update, error) {
	ret := _m.Called(ctx, req)

	var r0 []telegrammodels.Update
	if rf, ok := ret.Get(0).(func(context.Context, models.GetUpdatesRequest) []telegrammodels.Update); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]telegrammodels.Update)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.GetUpdatesRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserProfilePhotos provides a mock function with given fields: ctx, req
func (_m *Bot) GetUserProfilePhotos(ctx context.Context, req models.GetUserProfilePhotosRequest) (*telegrammodels.UserProfilePhotos, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.UserProfilePhotos
	if rf, ok := ret.Get(0).(func(context.Context, models.GetUserProfilePhotosRequest) *telegrammodels.UserProfilePhotos); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.UserProfilePhotos)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.GetUserProfilePhotosRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWebhookInfo provides a mock function with given fields: ctx
func (_m *Bot) GetWebhookInfo(ctx context.Context) (*telegrammodels.WebhookInfo, error) {
	ret := _m.Called(ctx)

	var r0 *telegrammodels.WebhookInfo
	if rf, ok := ret.Get(0).(func(context.Context) *telegrammodels.WebhookInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.WebhookInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LeaveChat provides a mock function with given fields: ctx, req
func (_m *Bot) LeaveChat(ctx context.Context, req models.ChatID) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ChatID) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Logout provides a mock function with given fields: ctx
func (_m *Bot) Logout(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PinChatMessage provides a mock function with given fields: ctx, req
func (_m *Bot) PinChatMessage(ctx context.Context, req models.PinChatMessageRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PinChatMessageRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendAnimation provides a mock function with given fields: ctx, req
func (_m *Bot) SendAnimation(ctx context.Context, req models.SendAnimationRequest) (*telegrammodels.Message, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.Message
	if rf, ok := ret.Get(0).(func(context.Context, models.SendAnimationRequest) *telegrammodels.Message); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendAnimationRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendAudio provides a mock function with given fields: ctx, req
func (_m *Bot) SendAudio(ctx context.Context, req models.SendAudioRequest) (*telegrammodels.Message, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.Message
	if rf, ok := ret.Get(0).(func(context.Context, models.SendAudioRequest) *telegrammodels.Message); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendAudioRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendChatAction provides a mock function with given fields: ctx, req
func (_m *Bot) SendChatAction(ctx context.Context, req models.SendChatActionRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SendChatActionRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendDocument provides a mock function with given fields: ctx, req
func (_m *Bot) SendDocument(ctx context.Context, req models.SendDocumentRequest) (*telegrammodels.Message, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.Message
	if rf, ok := ret.Get(0).(func(context.Context, models.SendDocumentRequest) *telegrammodels.Message); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendDocumentRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendLocation provides a mock function with given fields: ctx, req
func (_m *Bot) SendLocation(ctx context.Context, req models.SendLocationRequest) (*telegrammodels.Message, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.Message
	if rf, ok := ret.Get(0).(func(context.Context, models.SendLocationRequest) *telegrammodels.Message); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendLocationRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: ctx, req
func (_m *Bot) SendMessage(ctx context.Context, req models.SendMessageRequest) (*telegrammodels.Message, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.Message
	if rf, ok := ret.Get(0).(func(context.Context, models.SendMessageRequest) *telegrammodels.Message); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendPhoto provides a mock function with given fields: ctx, req
func (_m *Bot) SendPhoto(ctx context.Context, req models.SendPhotoRequest) (*telegrammodels.Message, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.Message
	if rf, ok := ret.Get(0).(func(context.Context, models.SendPhotoRequest) *telegrammodels.Message); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendPhotoRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendVideo provides a mock function with given fields: ctx, req
func (_m *Bot) SendVideo(ctx context.Context, req models.SendVideoRequest) (*telegrammodels.Message, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.Message
	if rf, ok := ret.Get(0).(func(context.Context, models.SendVideoRequest) *telegrammodels.Message); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendVideoRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendVoice provides a mock function with given fields: ctx, req
func (_m *Bot) SendVoice(ctx context.Context, req models.SendVoiceRequest) (*telegrammodels.Message, error) {
	ret := _m.Called(ctx, req)

	var r0 *telegrammodels.Message
	if rf, ok := ret.Get(0).(func(context.Context, models.SendVoiceRequest) *telegrammodels.Message); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telegrammodels.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SendVoiceRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetChatAdministratorCustomTitle provides a mock function with given fields: ctx, req
func (_m *Bot) SetChatAdministratorCustomTitle(ctx context.Context, req models.SetChatAdminCustomTitleRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SetChatAdminCustomTitleRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetChatDescription provides a mock function with given fields: ctx, req
func (_m *Bot) SetChatDescription(ctx context.Context, req models.SetChatDescriptionRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SetChatDescriptionRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetChatPermissions provides a mock function with given fields: ctx, req
func (_m *Bot) SetChatPermissions(ctx context.Context, req models.SetChatPermissionsRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SetChatPermissionsRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetChatPhoto provides a mock function with given fields: ctx, req
func (_m *Bot) SetChatPhoto(ctx context.Context, req models.SetChatPhotoRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SetChatPhotoRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetChatStickerSet provides a mock function with given fields: ctx, req
func (_m *Bot) SetChatStickerSet(ctx context.Context, req models.SetChatStickerSetRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SetChatStickerSetRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetChatTitle provides a mock function with given fields: ctx, req
func (_m *Bot) SetChatTitle(ctx context.Context, req models.SetChatTitleRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SetChatTitleRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetMyCommands provides a mock function with given fields: ctx, req
func (_m *Bot) SetMyCommands(ctx context.Context, req models.SetMyCommandsRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SetMyCommandsRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWebhook provides a mock function with given fields: ctx, req
func (_m *Bot) SetWebhook(ctx context.Context, req models.SetWebhookRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SetWebhookRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnpinAllChatMessages provides a mock function with given fields: ctx, req
func (_m *Bot) UnpinAllChatMessages(ctx context.Context, req models.ChatID) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ChatID) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnpinChatMessage provides a mock function with given fields: ctx, req
func (_m *Bot) UnpinChatMessage(ctx context.Context, req models.UnpinChatMessageRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UnpinChatMessageRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
