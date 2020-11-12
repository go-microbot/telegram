package bot

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/go-microbot/telegram/api"
	"github.com/go-microbot/telegram/models"
	"github.com/stretchr/testify/require"
)

var errMock = errors.New("error")

type testStrategy struct {
	fn      func(strategy *testStrategy)
	updates chan models.Update
	errs    chan error
}

func (t testStrategy) Listen() {
	if t.fn != nil {
		t.fn(&t)
	}
}

func (t testStrategy) Updates() chan models.Update {
	return t.updates
}

func (t testStrategy) Errors() chan error {
	return t.errs
}

func (t testStrategy) Stop() {
	close(t.updates)
	close(t.errs)
}

func newTestStrategy(fn func(strategy *testStrategy)) testStrategy {
	return testStrategy{
		fn:      fn,
		updates: make(chan models.Update),
		errs:    make(chan error),
	}
}

func Test_NewTelegramBot(t *testing.T) {
	botAPI := api.NewTelegramAPI("")
	require.NotNil(t, botAPI)
	bot := NewTelegramBot(&botAPI)
	require.NotNil(t, bot)
	require.NotNil(t, bot.api)
	require.NotNil(t, bot.updatesChan)
	require.NotNil(t, bot.updatesErrChan)
}

func TestTelegramBot_API(t *testing.T) {
	botAPI := api.NewTelegramAPI("")
	require.NotNil(t, botAPI)
	bot := NewTelegramBot(&botAPI)
	require.NotNil(t, bot.api)
	require.Equal(t, &botAPI, bot.API())
}

func TestTelegramBot_WaitForUpdates(t *testing.T) {
	botAPI := api.NewTelegramAPI("")
	require.NotNil(t, botAPI)
	bot := NewTelegramBot(&botAPI)
	require.NotNil(t, bot.api)
	var wg sync.WaitGroup
	wg.Add(1)
	strategy := newTestStrategy(func(strategy *testStrategy) {
		time.Sleep(time.Second)
		strategy.updates <- models.Update{
			Message: &models.Message{
				Text: "first message",
			},
		}
		time.Sleep(time.Second)

		strategy.errs <- errors.New("some error")
		time.Sleep(time.Second)

		strategy.updates <- models.Update{
			Message: &models.Message{
				Text: "second message",
			},
		}
		time.Sleep(time.Second)
		wg.Done()
	})

	messages := make([]string, 0)
	errs := make([]error, 0)
	go func() {
		for {
			select {
			case msg, ok := <-bot.updatesChan:
				require.True(t, ok)
				require.NotNil(t, msg)
				require.NotNil(t, msg.Message)
				messages = append(messages, msg.Message.Text)
			case err, ok := <-bot.updatesErrChan:
				require.True(t, ok)
				require.Error(t, err)
				errs = append(errs, err)
			}
		}
	}()

	go bot.WaitForUpdates(strategy)
	wg.Wait()

	expErrors := []error{
		errors.New("some error"),
	}
	expMessages := []string{"first message", "second message"}
	require.Equal(t, expMessages, messages)
	require.Equal(t, expErrors, errs)
}

func TestTelegramBot_Stop(t *testing.T) {
	botAPI := api.NewTelegramAPI("")
	require.NotNil(t, botAPI)
	bot := NewTelegramBot(&botAPI)
	require.NotNil(t, bot.api)
	var wg sync.WaitGroup
	wg.Add(1)
	strategy := newTestStrategy(func(strategy *testStrategy) {
		_, ok := <-strategy.updates
		require.False(t, ok)
		wg.Done()
	})
	go bot.WaitForUpdates(strategy)
	time.Sleep(time.Second * 3)
	bot.Stop()
	wg.Wait()
}

func TestTelegramBot_Updates(t *testing.T) {
	botAPI := api.NewTelegramAPI("")
	require.NotNil(t, botAPI)
	bot := NewTelegramBot(&botAPI)
	require.NotNil(t, bot.api)
	updates, errs := bot.Updates()
	require.NotNil(t, updates)
	require.NotNil(t, errs)
	require.Equal(t, bot.updatesChan, updates)
	require.Equal(t, bot.updatesErrChan, errs)
}
