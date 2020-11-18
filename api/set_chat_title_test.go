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

type setChatTitle struct{}

func (h setChatTitle) Test(ctx context.Context, t *testing.T) context.Context {
	groupChatID := ctx.Value(groupChatIDCtxKey)
	require.NotNil(t, groupChatID)

	newTitle := fmt.Sprintf("New title changed at %d", time.Now().UTC().Unix())

	err := localAPI.SetChatTitle(ctx, apiModels.SetChatTitleRequest{
		ChatID: query.NewParamAny(groupChatID.(int64)),
		Title:  newTitle,
	})
	require.NoError(t, err)

	return context.WithValue(ctx, TestDataKey(newGroupChatTitleCtxKey), newTitle)
}
