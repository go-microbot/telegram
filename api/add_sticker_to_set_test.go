package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/stretchr/testify/require"
)

type addStickerToSet struct{}

func (h addStickerToSet) Test(ctx context.Context, t *testing.T) context.Context {
	stickerFileID := ctx.Value(uploadedStickerFileIDCtxKey)
	require.NotNil(t, stickerFileID)
	adminUserID := ctx.Value(adminUserIDCtxKey)
	require.NotNil(t, adminUserID)
	stickerSetName := ctx.Value(stickerSetNameCtxKey)
	require.NotNil(t, stickerSetName)

	err := localAPI.AddStickerToSet(ctx, apiModels.AddStickerToSetRequest{
		UserID:     form.NewPartAny(adminUserID.(int32)),
		Name:       form.NewPartText(stickerSetName.(string)),
		Emojis:     form.NewPartText("😃"),
		PngSticker: form.NewPartText(stickerFileID.(string)),
	})
	require.NoError(t, err)

	return ctx
}
