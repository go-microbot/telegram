package api

import (
	"net/http"
)

const baseURL = "https://api.telegram.org"

// TelegramAPI represents Telegram Bot API.
type TelegramAPI struct {
	client *http.Client
	token  string
	url    string
}

// NewTelegramAPI returns new TelegramAPI instance.
func NewTelegramAPI(token string) TelegramAPI {
	return TelegramAPI{
		token:  token,
		client: http.DefaultClient,
		url:    baseURL,
	}
}
