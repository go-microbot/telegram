package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type getChat struct{}

func (h getChat) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)
	groupChatTitle := ctx.Value(newGroupChatTitleCtxKey)
	require.NotNil(t, groupChatTitle)
	groupChatDescription := ctx.Value(newGroupChatDescriptionCtxKey)
	require.NotNil(t, groupChatDescription)

	chat, err := localAPI.GetChat(ctx, apiModels.ChatID{
		ChatID: query.NewParamAny(groupChatID.(int64)),
	})
	require.NoError(t, err)
	require.NotNil(t, chat)

	require.Equal(t, groupChatID, chat.ID)
	require.Equal(t, groupChatTitle, chat.Title)
	require.Equal(t, groupChatDescription, chat.Description)

	return ctx
}
