package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTelegramAPI_GetMe(t *testing.T) {
	me, err := localAPI.GetMe(context.Background())
	require.NoError(t, err)
	require.NotNil(t, me)
	require.True(t, me.IsBot)
	require.Equal(t, "go_micro_bot", me.Username)
}
