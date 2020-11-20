package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendVideo struct{}

func (h sendVideo) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendVideo(ctx, apiModels.SendVideoRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: form.NewPartAny(true),
		Video:               form.NewPartFile("./test_data/test_video.mp4"),
		Caption:             form.NewPartText("test video"),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
