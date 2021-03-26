package api

import (
	"context"
	"os"
	"path"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type setWebhook struct{}

func (h setWebhook) Test(ctx context.Context, t *testing.T) context.Context {
	// load certificate.
	wd, err := os.Getwd()
	if err != nil {
		require.NoError(t, err)
	}

	// set webhook.
	url := "localhost:1012"
	err = localAPI.SetWebhook(ctx, apiModels.SetWebhookRequest{
		Certificate: form.NewPartFile(path.Join(wd, "./test_data/telegram_test.key")),
		URL:         query.NewParamString(url),
	})
	require.NoError(t, err)

	return context.WithValue(ctx, TestDataKey(webhookURLCtxKey), url)
}
