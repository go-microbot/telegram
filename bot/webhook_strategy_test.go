package bot

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/go-microbot/telegram/models"
	"github.com/stretchr/testify/require"
)

func Test_NewUpdatesStrategyWebhook(t *testing.T) {
	hook := NewUpdatesStrategyWebhook(WebhookConfig{})
	require.NotNil(t, hook)
	require.NotNil(t, hook.cfg)
	require.NotNil(t, hook.errorsChan)
	require.NotNil(t, hook.updatesChan)
	require.NotNil(t, hook.doneChan)
}

func TestUpdatesStrategyWebhook_Updates(t *testing.T) {
	hook := NewUpdatesStrategyWebhook(WebhookConfig{})
	require.NotNil(t, hook)
	require.NotNil(t, hook.updatesChan)
	require.Equal(t, hook.updatesChan, hook.Updates())
}

func TestUpdatesStrategyWebhook_Errors(t *testing.T) {
	hook := NewUpdatesStrategyWebhook(WebhookConfig{})
	require.NotNil(t, hook)
	require.NotNil(t, hook.errorsChan)
	require.Equal(t, hook.errorsChan, hook.Errors())
}

func TestUpdatesStrategyWebhook_Listen(t *testing.T) {
	hook := NewUpdatesStrategyWebhook(WebhookConfig{
		ServeURL: "localhost:8443",
	})
	require.NotNil(t, hook)

	errs := make([]error, 0)
	messages := make([]string, 0)
	go func() {
		for {
			select {
			case msg, ok := <-hook.updatesChan:
				require.True(t, ok)
				require.NotNil(t, msg)
				require.NotNil(t, msg.Message)
				messages = append(messages, msg.Message.Text)
			case err, ok := <-hook.errorsChan:
				require.True(t, ok)
				require.Error(t, err)
				errs = append(errs, err)
			}
		}
	}()

	go hook.Listen()

	url := "http://" + hook.cfg.ServeURL
	// message 1.
	data, err := json.Marshal(models.Update{
		Message: &models.Message{
			Text: "message 1",
		},
	})
	require.NoError(t, err)
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	require.NoError(t, err)
	closeBody(t, resp.Body)

	// error 1.
	require.NoError(t, err)
	resp, err = http.Post(url, "application/json", bytes.NewReader([]byte("invalid")))
	require.NoError(t, err)
	closeBody(t, resp.Body)

	// message 2.
	data, err = json.Marshal(models.Update{
		Message: &models.Message{
			Text: "message 2",
		},
	})
	require.NoError(t, err)
	resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	require.NoError(t, err)
	closeBody(t, resp.Body)

	// error 2.
	require.NoError(t, err)
	resp, err = http.Post(url, "application/json", bytes.NewReader([]byte("data invalid")))
	require.NoError(t, err)
	closeBody(t, resp.Body)

	expMessages := []string{"message 1", "message 2"}
	require.Equal(t, expMessages, messages)
	require.Equal(t, 2, len(errs))
	for i := range errs {
		_, ok := errs[i].(*json.SyntaxError)
		require.True(t, ok)
	}
}

func TestUpdatesStrategyWebhook_Stop(t *testing.T) {
	hook := NewUpdatesStrategyWebhook(WebhookConfig{
		ServeURL: "localhost:8440",
	})
	require.NotNil(t, hook)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, ok := <-hook.updatesChan
		require.False(t, ok)
		wg.Done()
	}()

	go hook.Listen()
	time.Sleep(time.Second * 2)
	hook.Stop()
	wg.Wait()
}

func closeBody(t *testing.T, body io.ReadCloser) {
	err := body.Close()
	require.NoError(t, err)
}
