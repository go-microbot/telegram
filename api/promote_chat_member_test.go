package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type promoteChatMember struct{}

func (h promoteChatMember) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)
	adminUserID := ctx.Value(adminUserIDCtxKey)
	require.NotNil(t, adminUserID)

	err := localAPI.PromoteChatMember(ctx, apiModels.PromoteChatMemberRequest{
		ChatID:          query.NewParamAny(groupChatID.(int64)),
		UserID:          query.NewParamInt(int(adminUserID.(int32))),
		CanPostMessages: true,
	})
	if err != nil {
		require.Contains(t, err.Error(), "can't remove chat owner")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
