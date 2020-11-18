package api

import (
	"context"
	"testing"

	"github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type setChatPhoto struct{}

func (h setChatPhoto) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)

	err := localAPI.SetChatPhoto(ctx, models.SetChatPhotoRequest{
		ChatID: query.NewParamAny(groupChatID.(int64)),
		Photo:  form.NewPartFile("./test_data/chat_photo.png"),
	})
	require.NoError(t, err)

	return ctx
}
