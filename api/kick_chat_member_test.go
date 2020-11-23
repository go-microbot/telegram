package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type kickChatMember struct{}

func (h kickChatMember) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)
	adminUserID := ctx.Value(adminUserIDCtxKey)
	require.NotNil(t, adminUserID)

	err := localAPI.KickChatMember(ctx, apiModels.KickChatMemberRequest{
		ChatID: query.NewParamAny(groupChatID.(int64)),
		UserID: adminUserID.(int32),
	})
	if err != nil {
		require.Contains(t, err.Error(), "can't remove chat owner")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
