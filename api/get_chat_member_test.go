package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type getChatMember struct{}

func (h getChatMember) Test(ctx context.Context, t *testing.T) context.Context {
	chatID := ctx.Value(chatIDCtxKey)
	require.NotNil(t, chatID)
	botUserID := ctx.Value(botUserIDCtxKey)
	require.NotNil(t, botUserID)

	member, err := localAPI.GetChatMember(ctx, apiModels.GetChatMemberRequest{
		ChatID: query.NewParamAny(chatID.(int64)),
		UserID: query.NewParamInt(int(botUserID.(int32))),
	})
	require.NoError(t, err)
	require.NotNil(t, member)
	require.Equal(t, botUserID, member.User.ID)
	require.True(t, member.User.IsBot)

	return ctx
}
