package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendVoice struct{}

func (h sendVoice) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendVoice(ctx, apiModels.SendVoiceRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: form.NewPartAny(true),
		Voice:               form.NewPartFile("./test_data/test_voice.opus"),
		Caption:             form.NewPartText("test video"),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
