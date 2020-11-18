package api

import (
	"context"
	"testing"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/stretchr/testify/require"
)

type setMyCommands struct{}

func (h setMyCommands) Test(ctx context.Context, t *testing.T) context.Context {
	commands := []models.BotCommand{
		{
			Command:     "test1",
			Description: "Test command #1",
		},
		{
			Command:     "test2",
			Description: "Test command #2",
		},
	}

	err := localAPI.SetMyCommands(ctx, apiModels.SetMyCommandsRequest{
		Commands: commands,
	})
	require.NoError(t, err)

	return context.WithValue(ctx, TestDataKey(botCommandsCtxKey), commands)
}
