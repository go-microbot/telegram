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
	KeyFile  string
	CertFile string
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
func (w UpdatesStrategyWebhook) Updates() chan models.Update {
	return w.updatesChan
}

// Errors returns errors channel.
func (w UpdatesStrategyWebhook) Errors() chan error {
	return w.errorsChan
}

// Listen starts webhook listening.
func (w UpdatesStrategyWebhook) Listen() {
	go w.listen()

	<-w.doneChan
}

// Stop stops webhook listening.
func (w UpdatesStrategyWebhook) Stop() {
	w.doneChan <- struct{}{}
	close(w.errorsChan)
	close(w.updatesChan)
}

func (w UpdatesStrategyWebhook) listen() {
	handler := webHookHandler{
		updates: w.updatesChan,
		errs:    w.errorsChan,
	}

	var err error
	switch {
	case w.cfg.CertFile != "" && w.cfg.KeyFile != "":
		err = http.ListenAndServeTLS(w.cfg.ServeURL, w.cfg.CertFile, w.cfg.KeyFile, &handler)
	default:
		err = http.ListenAndServe(w.cfg.ServeURL, &handler)
	}

	if err != nil {
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
