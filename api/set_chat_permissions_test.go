package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type setChatPermissions struct{}

func (h setChatPermissions) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)

	err := localAPI.SetChatPermissions(ctx, apiModels.SetChatPermissionsRequest{
		ChatID: query.NewParamAny(groupChatID.(int64)),
		Permissions: models.ChatPermissions{
			CanChangeInfo:         true,
			CanAddWebPagePreviews: true,
			CanInviteUsers:        true,
			CanPinMessages:        true,
			CanSendMediaMessages:  true,
			CanSendMessages:       true,
			CanSendOtherMessages:  true,
			CanSendPolls:          true,
		},
	})
	require.NoError(t, err)

	return ctx
}
