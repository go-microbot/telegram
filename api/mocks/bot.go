// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	apimodels "github.com/go-microbot/telegram/api/models"

	http "net/http"

	mock "github.com/stretchr/testify/mock"

	models "github.com/go-microbot/telegram/models"
)

// Bot is an autogenerated mock type for the Bot type
type Bot struct {
	mock.Mock
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

// GetMe provides a mock function with given fields: ctx
func (_m *Bot) GetMe(ctx context.Context) (*models.User, error) {
	ret := _m.Called(ctx)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context) *models.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
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
func (_m *Bot) GetPollUpdates(ctx context.Context, req apimodels.GetUpdatesRequest, client *http.Client) ([]models.Update, error) {
	ret := _m.Called(ctx, req, client)

	var r0 []models.Update
	if rf, ok := ret.Get(0).(func(context.Context, apimodels.GetUpdatesRequest, *http.Client) []models.Update); ok {
		r0 = rf(ctx, req, client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Update)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, apimodels.GetUpdatesRequest, *http.Client) error); ok {
		r1 = rf(ctx, req, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpdates provides a mock function with given fields: ctx, req
func (_m *Bot) GetUpdates(ctx context.Context, req apimodels.GetUpdatesRequest) ([]models.Update, error) {
	ret := _m.Called(ctx, req)

	var r0 []models.Update
	if rf, ok := ret.Get(0).(func(context.Context, apimodels.GetUpdatesRequest) []models.Update); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Update)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, apimodels.GetUpdatesRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWebhookInfo provides a mock function with given fields: ctx
func (_m *Bot) GetWebhookInfo(ctx context.Context) (*models.WebhookInfo, error) {
	ret := _m.Called(ctx)

	var r0 *models.WebhookInfo
	if rf, ok := ret.Get(0).(func(context.Context) *models.WebhookInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.WebhookInfo)
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

// SendMessage provides a mock function with given fields: ctx, req
func (_m *Bot) SendMessage(ctx context.Context, req apimodels.SendMessageRequest) (*models.Message, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.Message
	if rf, ok := ret.Get(0).(func(context.Context, apimodels.SendMessageRequest) *models.Message); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, apimodels.SendMessageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetWebhook provides a mock function with given fields: ctx, req
func (_m *Bot) SetWebhook(ctx context.Context, req apimodels.SetWebhookRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, apimodels.SetWebhookRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}