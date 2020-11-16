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
