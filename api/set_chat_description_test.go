package api

import (
	"context"
	"fmt"
	"testing"
	"time"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
	"github.com/stretchr/testify/require"
)

type setChatDescription struct{}

func (h setChatDescription) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)

	newDescription := fmt.Sprintf("New description changed at %d", time.Now().UTC().Unix())

	err := localAPI.SetChatDescription(ctx, apiModels.SetChatDescriptionRequest{
		ChatID:      query.NewParamAny(groupChatID.(int64)),
		Description: newDescription,
	})
	require.NoError(t, err)

	return context.WithValue(ctx, TestDataKey(newGroupChatDescriptionCtxKey), newDescription)
}
