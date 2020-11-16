package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type deleteWebhook struct{}

func (h deleteWebhook) Test(ctx context.Context, t *testing.T) context.Context {
	err := localAPI.DeleteWebhook(ctx)
	require.NoError(t, err)
	return ctx
}
