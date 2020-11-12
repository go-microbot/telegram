package bot

import (
	"encoding/json"
	"net/http"

	"github.com/go-microbot/telegram/models"
)

// UpdatesStrategyWebhook represents object to receive icoming messages
// using webhook (https://core.telegram.org/bots/api#setwebhook).
type UpdatesStrategyWebhook struct {
	updatesChan chan models.Update
	errorsChan  chan error
	doneChan    chan struct{}
	cfg         WebhookConfig
}

// WebhookConfig represents webhook configuration.
type WebhookConfig struct {
	ServeURL string
}

type webHookHandler struct {
	updates chan models.Update
	errs    chan error
}

// NewUpdatesStrategyWebhook returns new instance of the UpdatesStrategyWebhook.
func NewUpdatesStrategyWebhook(cfg WebhookConfig) UpdatesStrategyWebhook {
	return UpdatesStrategyWebhook{
		cfg:         cfg,
		updatesChan: make(chan models.Update, defaultUpdatesLimit),
		doneChan:    make(chan struct{}),
		errorsChan:  make(chan error),
	}
}

// Updates returns updates channel.
func (poll UpdatesStrategyWebhook) Updates() chan models.Update {
	return poll.updatesChan
}

// Errors returns errors channel.
func (poll UpdatesStrategyWebhook) Errors() chan error {
	return poll.errorsChan
}

// Listen starts webhook listening.
func (poll UpdatesStrategyWebhook) Listen() {
	go poll.listen()

	<-poll.doneChan
}

// Stop stops webhook listening.
func (poll UpdatesStrategyWebhook) Stop() {
	poll.doneChan <- struct{}{}
	close(poll.errorsChan)
	close(poll.updatesChan)
}

func (poll UpdatesStrategyWebhook) listen() {
	if err := http.ListenAndServe(poll.cfg.ServeURL, &webHookHandler{
		updates: poll.updatesChan,
		errs:    poll.errorsChan,
	}); err != nil {
		panic(err)
	}
}

func (h *webHookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var update models.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		h.errs <- err
		return
	}

	h.updates <- update
}
