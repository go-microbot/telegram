package bot

import (
	"sync"
	"testing"
	"time"

	"github.com/go-microbot/telegram/api"
	apiMocks "github.com/go-microbot/telegram/api/mocks"
	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_NewUpdatesStrategyLongPolling(t *testing.T) {
	poll := NewUpdatesStrategyLongPolling(LongPollingConfig{
		Timeout: 1,
	})
	require.NotNil(t, poll)
	require.NotNil(t, poll.cfg)
	require.NotNil(t, poll.client)
	require.NotNil(t, poll.updatesChan)
	require.NotNil(t, poll.errorsChan)
	require.NotNil(t, poll.doneChan)
	require.Equal(t, defaultUpdatesLimit, poll.cfg.Limit)
}

func TestUpdatesStrategyLongPolling_Updates(t *testing.T) {
	poll := NewUpdatesStrategyLongPolling(LongPollingConfig{})
	require.NotNil(t, poll)
	require.NotNil(t, poll.updatesChan)
	require.Equal(t, poll.updatesChan, poll.Updates())
}

func TestUpdatesStrategyLongPolling_Errors(t *testing.T) {
	poll := NewUpdatesStrategyLongPolling(LongPollingConfig{})
	require.NotNil(t, poll)
	require.NotNil(t, poll.errorsChan)
	require.Equal(t, poll.errorsChan, poll.Errors())
}

func TestUpdatesStrategyLongPolling_Listen(t *testing.T) {
	t.Run("with errors", func(t *testing.T) {
		tAPI := new(apiMocks.Bot)
		defer tAPI.AssertExpectations(t)
		poll := NewUpdatesStrategyLongPolling(LongPollingConfig{
			Timeout:        1,
			Offset:         1,
			Limit:          2,
			AllowedUpdates: []string{"all"},
			BotAPI:         tAPI,
		})
		require.NotNil(t, poll)

		// timeout error.
		tAPI.On("GetPollUpdates", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, api.ErrReqTimeout).Once()

		// with error.
		tAPI.On("GetPollUpdates", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errMock)

		var wg sync.WaitGroup
		wg.Add(1)
		errs := make([]error, 0)
		go func() {
			for err := range poll.errorsChan {
				require.Error(t, err)
				errs = append(errs, err)
				wg.Done()
				break
			}
		}()

		go poll.Listen()
		wg.Wait()

		expErrors := []error{errMock}
		require.Equal(t, expErrors, errs)
	})
	t.Run("all ok", func(t *testing.T) {
		tAPI := new(apiMocks.Bot)
		defer tAPI.AssertExpectations(t)
		poll := NewUpdatesStrategyLongPolling(LongPollingConfig{
			Timeout:        1,
			Offset:         1,
			Limit:          2,
			AllowedUpdates: []string{"all"},
			BotAPI:         tAPI,
		})
		require.NotNil(t, poll)

		// with updates.
		tAPI.On("GetPollUpdates", mock.Anything, apiModels.GetUpdatesRequest{
			Limit:          query.NewParamInt(2),
			AllowedUpdates: query.NewParamStringSlice([]string{"all"}),
			Offset:         query.NewParamInt(1),
			Timeout:        query.NewParamInt(1),
		}, mock.Anything).Return([]models.Update{
			{
				UpdateID: 1,
				Message: &models.Message{
					Text: "message 1",
				},
			},
		}, nil)
		tAPI.On("GetPollUpdates", mock.Anything, apiModels.GetUpdatesRequest{
			Limit:          query.NewParamInt(2),
			AllowedUpdates: query.NewParamStringSlice([]string{"all"}),
			Offset:         query.NewParamInt(2),
			Timeout:        query.NewParamInt(1),
		}, mock.Anything).Return([]models.Update{
			{
				UpdateID: 2,
				Message: &models.Message{
					Text: "message 2",
				},
			},
		}, nil)
		tAPI.On("GetPollUpdates", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, nil).Maybe()

		messages := make([]string, 0)
		var wg sync.WaitGroup
		wg.Add(1)
		expMessages := []string{"message 1", "message 2"}
		go func() {
			for msg := range poll.updatesChan {
				require.NotNil(t, msg)
				require.NotNil(t, msg.Message)
				messages = append(messages, msg.Message.Text)
				if len(messages) == len(expMessages) {
					wg.Done()
					break
				}
			}
		}()

		go poll.Listen()
		wg.Wait()

		require.Equal(t, expMessages, messages)
	})
}

func TestUpdatesStrategyLongPolling_Stop(t *testing.T) {
	tAPI := new(apiMocks.Bot)
	defer tAPI.AssertExpectations(t)
	poll := NewUpdatesStrategyLongPolling(LongPollingConfig{
		BotAPI: tAPI,
	})
	require.NotNil(t, poll)

	tAPI.On("GetPollUpdates", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, nil)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		_, ok := <-poll.updatesChan
		require.False(t, ok)
		wg.Done()
	}()

	go poll.Listen()
	time.Sleep(time.Second * 2)
	poll.Stop()
	wg.Wait()
}
