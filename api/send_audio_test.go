package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendAudio struct{}

func (h sendAudio) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendAudio(ctx, apiModels.SendAudioRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: form.NewPartAny(true),
		Audio:               form.NewPartFile("./test_data/test_audio.mp3"),
		Caption:             form.NewPartText("test audio"),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
