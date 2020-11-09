package bot

import (
	"context"
	"fmt"

	"github.com/go-microbot/telegram/api"
)

// Bot represents Telegram Bot model.
// https://core.telegram.org/bots.
type Bot struct {
	api api.ServiceAPI
}

// New returns new Bot instance.
func New(botAPI api.ServiceAPI) Bot {
	return Bot{
		api: botAPI,
	}
}

// GetAPI returns bot's API instance.
func (b *Bot) GetAPI() api.ServiceAPI {
	return b.api
}

func (b *Bot) Test() {
	user, err := b.api.GetMe(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.Username)
}
