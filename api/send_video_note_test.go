package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendVideoNote struct{}

func (h sendVideoNote) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendVideoNote(ctx, apiModels.SendVideoNoteRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: form.NewPartAny(true),
		VideoNote:           form.NewPartFile("./test_data/test_video.mp4"),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
