package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type setChatAdministratorCustomTitle struct{}

func (h setChatAdministratorCustomTitle) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)
	botUserID := ctx.Value(botUserIDCtxKey)
	require.NotNil(t, botUserID)

	err := localAPI.SetChatAdministratorCustomTitle(ctx, apiModels.SetChatAdminCustomTitleRequest{
		ChatID:      query.NewParamAny(groupChatID.(int64)),
		CustomTitle: "root",
		UserID:      botUserID.(int32),
	})
	if err != nil {
		require.Contains(t, err.Error(), "Bad Request")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
