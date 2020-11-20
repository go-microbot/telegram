package api

import (
	"context"
	"fmt"
	"testing"
	"time"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/stretchr/testify/require"
)

type createNewStickerSet struct{}

func (h createNewStickerSet) Test(ctx context.Context, t *testing.T) context.Context {
	adminUserID := ctx.Value(adminUserIDCtxKey)
	require.NotNil(t, adminUserID)
	botUserName := ctx.Value(botUsernameCtxKey)
	require.NotNil(t, botUserName)

	name := fmt.Sprintf("test_set_%d_by_%s", time.Now().UTC().Unix(), botUserName.(string))
	err := localAPI.CreateNewStickerSet(ctx, apiModels.CreateNewStickerSetRequest{
		UserID:     form.NewPartAny(adminUserID.(int32)),
		Emojis:     form.NewPartText("😀"),
		Name:       form.NewPartText(name),
		PngSticker: form.NewPartFile("./test_data/stickers/sticker1.png"),
		Title:      form.NewPartText("test title"),
	})
	require.NoError(t, err)

	return context.WithValue(ctx, TestDataKey(stickerSetNameCtxKey), name)
}
