package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type getChatMembersCount struct{}

func (h getChatMembersCount) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)

	count, err := localAPI.GetChatMembersCount(ctx, apiModels.ChatID{
		ChatID: query.NewParamAny(groupChatID.(int64)),
	})
	require.NoError(t, err)
	require.True(t, count >= 1)

	return ctx
}
