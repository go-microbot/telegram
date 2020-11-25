package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/stretchr/testify/require"
)

type setStickerSetThumb struct{}

func (h setStickerSetThumb) Test(ctx context.Context, t *testing.T) context.Context {
	stickerSetName := ctx.Value(stickerSetNameCtxKey)
	require.NotNil(t, stickerSetName)
	adminUserID := ctx.Value(adminUserIDCtxKey)
	require.NotNil(t, adminUserID)

	err := localAPI.SetStickerSetThumb(ctx, apiModels.SetStickerSetThumbRequest{
		Name:   form.NewPartText(stickerSetName.(string)),
		UserID: form.NewPartAny(adminUserID.(int32)),
		Thumb:  form.NewPartFile("./test_data/stickers/thumb.png"),
	})
	if err != nil {
		require.Contains(t, err.Error(), "STICKER_INVALID")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
