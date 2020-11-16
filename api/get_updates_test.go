package api

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-microbot/telegram/api/models"
	"github.com/stretchr/testify/require"
)

type getUpdates struct{}

func (h getUpdates) Test(ctx context.Context, t *testing.T) {
	// with default HTTP client.
	updates, err := localAPI.GetUpdates(context.Background(), models.GetUpdatesRequest{})
	require.NoError(t, err)
	require.NotNil(t, updates)

	// with custom HTTP client.
	updates, err = localAPI.GetPollUpdates(context.Background(), models.GetUpdatesRequest{}, http.DefaultClient)
	require.NoError(t, err)
	require.NotNil(t, updates)
}
