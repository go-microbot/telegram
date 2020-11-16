package api

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-microbot/telegram/api/models"
	"github.com/stretchr/testify/require"
)

func TestTelegramAPI_GetUpdates(t *testing.T) {
	updates, err := localAPI.GetUpdates(context.Background(), models.GetUpdatesRequest{})
	require.NoError(t, err)
	require.NotNil(t, updates)
}

func TestTelegramAPI_GetPollUpdates(t *testing.T) {
	updates, err := localAPI.GetPollUpdates(context.Background(), models.GetUpdatesRequest{}, http.DefaultClient)
	require.NoError(t, err)
	require.NotNil(t, updates)
}
