package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type getMyCommands struct{}

func (h getMyCommands) Test(ctx context.Context, t *testing.T) context.Context {
	botCommands := ctx.Value(botCommandsCtxKey)
	require.NotNil(t, botCommands)

	commands, err := localAPI.GetMyCommands(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, commands)
	require.Equal(t, botCommands, commands)

	return ctx
}
