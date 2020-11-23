package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type restrictChatMember struct{}

func (h restrictChatMember) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)
	adminUserID := ctx.Value(adminUserIDCtxKey)
	require.NotNil(t, adminUserID)

	err := localAPI.RestrictChatMember(ctx, apiModels.RestrictChatMemberRequest{
		ChatID: query.NewParamAny(groupChatID.(int64)),
		UserID: query.NewParamInt(int(adminUserID.(int32))),
		Permissions: models.ChatPermissions{
			CanAddWebPagePreviews: false,
		},
	})
	if err != nil {
		require.Contains(t, err.Error(), "can't remove chat owner")
		return ctx
	}
	require.NoError(t, err)

	return ctx
}
