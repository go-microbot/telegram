package api

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const localAPIURL = "http://localhost:8081"

var (
	localAPI TelegramAPI
)

type Testable interface {
	Test(ctx context.Context, t *testing.T)
}

func TestMain(m *testing.M) {
	localAPI = NewTelegramAPI( /*os.Getenv("TEST_BOT_TOKEN")*/ "1256583982:AAHoS3RanoCsXtKhNJCQSOftJXWSRJnLg2o")
	localAPI.url = localAPIURL

	os.Exit(m.Run())
}

func Test_NewTelegramAPI(t *testing.T) {
	tAPI := NewTelegramAPI("123")
	require.NotNil(t, tAPI)
	require.Equal(t, "123", tAPI.token)
	require.NotNil(t, tAPI.client)
	require.Equal(t, baseURL, tAPI.url)
}

func TestTelegramAPI_Integration(t *testing.T) {
	// ctx := context.Background()
	testCases := []struct {
		name        string
		testHandler Testable
		ctx         context.Context
		sleep       int
	}{
		{
			name:        "getMe",
			testHandler: getMe{},
		},
		{
			name:        "getUpdates",
			testHandler: getUpdates{},
		},
		{
			name:        "setWebhook",
			testHandler: setWebhook{},
		},
	}
	for i := range testCases {
		tc := &testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			tc.testHandler.Test(tc.ctx, t)
			time.Sleep(time.Second * time.Duration(tc.sleep))
		})
	}
}
