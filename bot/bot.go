package bot

import (
	"context"
	"fmt"

	"github.com/go-microbot/telegram/api"
)

// TelegramBot represents Telegram Bot handler.
// https://core.telegram.org/bots.
type TelegramBot struct {
	api api.Bot
}

// NewTelegramBot returns new Telegram Bot instance.
func NewTelegramBot(botAPI api.Bot) TelegramBot {
	return TelegramBot{
		api: botAPI,
	}
}

// API returns bot's API instance.
func (b *TelegramBot) API() api.Bot {
	return b.api
}

func (b *TelegramBot) Test() {
	user, err := b.api.GetMe(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.Username)
}
