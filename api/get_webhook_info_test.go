package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type getWebhookInfo struct{}

func (h getWebhookInfo) Test(ctx context.Context, t *testing.T) context.Context {
	info, err := localAPI.GetWebhookInfo(ctx)
	require.NoError(t, err)
	require.NotNil(t, info)
	webhookURL := ctx.Value(webhookURLCtxKey)
	require.NotEmpty(t, webhookURL)
	require.Equal(t, webhookURL, info.URL)
	return ctx
}
