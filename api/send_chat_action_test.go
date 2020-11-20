package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type sendChatAction struct{}

func (h sendChatAction) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	msgID := ctx.Value(sentMessageIDCtxKey)
	require.NotNil(t, msgID)

	err := localAPI.SendChatAction(ctx, apiModels.SendChatActionRequest{
		ChatID: query.NewParamAny(chatID.(int64)),
		Action: "upload_photo",
	})
	require.NoError(t, err)

	return ctx
}
