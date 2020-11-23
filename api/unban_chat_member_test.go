package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type unbanChatMember struct{}

func (h unbanChatMember) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)
	adminUserID := ctx.Value(adminUserIDCtxKey)
	require.NotNil(t, adminUserID)

	err := localAPI.UnbanChatMember(ctx, apiModels.UnbanChatMemberRequest{
		ChatID:       query.NewParamAny(groupChatID.(int64)),
		UserID:       adminUserID.(int32),
		OnlyIfBanned: true,
	})
	require.NoError(t, err)

	return ctx
}
