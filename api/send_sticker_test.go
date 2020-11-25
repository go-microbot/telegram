package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendSticker struct{}

func (h sendSticker) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)

	msg, err := localAPI.SendSticker(ctx, apiModels.SendStickerRequest{
		ChatID:              query.NewParamAny(chatID.(int64)),
		DisableNotification: form.NewPartAny(true),
		Sticker:             form.NewPartFile("./test_data/stickers/sticker1.png"),
	})
	require.NoError(t, err)
	require.NotNil(t, msg)

	return ctx
}
