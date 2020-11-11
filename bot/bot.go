package bot

import (
	"github.com/go-microbot/telegram/api"
	"github.com/go-microbot/telegram/models"
)

// TelegramBot represents Telegram Bot handler.
// https://core.telegram.org/bots.
type TelegramBot struct {
	api            api.Bot
	updatesChan    chan models.Update
	updatesErrChan chan error
	strategy       UpdatesStrategy
}

// NewTelegramBot returns new Telegram Bot instance.
func NewTelegramBot(botAPI api.Bot) TelegramBot {
	return TelegramBot{
		api:            botAPI,
		updatesChan:    make(chan models.Update, defaultUpdatesLimit),
		updatesErrChan: make(chan error),
	}
}

// API returns bot's API instance.
func (b *TelegramBot) API() api.Bot {
	return b.api
}

// WaitForUpdates starts listening of new messages using provided strategy.
func (b *TelegramBot) WaitForUpdates(strategy UpdatesStrategy) {
	b.strategy = strategy

	go strategy.Listen()

	for {
		select {
		case err, ok := <-strategy.Errors():
			if ok {
				b.updatesErrChan <- err
			}
		case update, ok := <-strategy.Updates():
			if ok {
				b.updatesChan <- update
			}
		}
	}
}

// Stop stops listening of new messages.
func (b *TelegramBot) Stop() {
	b.strategy.Stop()
}

// Updates returns updates and errors channels.
func (b *TelegramBot) Updates() (upds chan models.Update, errs chan error) {
	return b.updatesChan, b.updatesErrChan
}
