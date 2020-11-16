package bot

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-microbot/telegram/api"
	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// UpdatesStrategyLongPolling represents object to receive icoming messages
// using long polling (https://en.wikipedia.org/wiki/Push_technology#Long_polling).
type UpdatesStrategyLongPolling struct {
	client      *http.Client
	updatesChan chan models.Update
	errorsChan  chan error
	doneChan    chan struct{}
	cfg         LongPollingConfig
}

// LongPollingConfig represents long polling configuration.
type LongPollingConfig struct {
	Offset int
	Limit  int
	// In seconds.
	Timeout        int
	AllowedUpdates []string
	BotAPI         api.Bot
}

// Updates returns updates channel.
func (poll UpdatesStrategyLongPolling) Updates() chan models.Update {
	return poll.updatesChan
}

// Errors returns errors channel.
func (poll UpdatesStrategyLongPolling) Errors() chan error {
	return poll.errorsChan
}

// Listen starts polling listening.
func (poll UpdatesStrategyLongPolling) Listen() {
	req := prepareUpdatesRequest(poll.cfg)

	go func(req apiModels.GetUpdatesRequest) {
		for {
			poll.polling(&req)
			// wait 1s before next polling.
			time.Sleep(time.Second)
		}
	}(req)

	<-poll.doneChan
}

// Stop stops long polling listening.
func (poll UpdatesStrategyLongPolling) Stop() {
	poll.doneChan <- struct{}{}
	close(poll.errorsChan)
	close(poll.updatesChan)
}

func (poll UpdatesStrategyLongPolling) polling(req *apiModels.GetUpdatesRequest) {
	ctx := context.Background()
	updates, err := poll.cfg.BotAPI.GetPollUpdates(ctx, *req, poll.client)
	if err != nil {
		if errors.Is(err, api.ErrReqTimeout) {
			return
		}

		poll.errorsChan <- err
		return
	}

	if nUpdates := len(updates); nUpdates != 0 {
		req.Offset = query.NewParamInt(int(updates[nUpdates-1].UpdateID + 1))
	}

	for i := range updates {
		poll.updatesChan <- updates[i]
	}
}

// NewUpdatesStrategyLongPolling returns new instance of the UpdatesStrategyLongPolling.
func NewUpdatesStrategyLongPolling(cfg LongPollingConfig) UpdatesStrategyLongPolling {
	if cfg.Limit <= 0 {
		cfg.Limit = defaultUpdatesLimit
	}

	return UpdatesStrategyLongPolling{
		cfg: cfg,
		client: &http.Client{
			Timeout: time.Second * time.Duration(cfg.Timeout),
		},
		updatesChan: make(chan models.Update, cfg.Limit),
		doneChan:    make(chan struct{}),
		errorsChan:  make(chan error),
	}
}

func prepareUpdatesRequest(cfg LongPollingConfig) apiModels.GetUpdatesRequest {
	req := apiModels.GetUpdatesRequest{
		Timeout: query.NewParamInt(cfg.Timeout),
	}

	if cfg.Offset != 0 {
		req.Offset = query.NewParamInt(cfg.Offset)
	}

	if cfg.Limit > 0 {
		req.Limit = query.NewParamInt(cfg.Limit)
	}

	if len(cfg.AllowedUpdates) != 0 {
		req.AllowedUpdates = query.NewParamStringSlice(cfg.AllowedUpdates)
	}

	return req
}
