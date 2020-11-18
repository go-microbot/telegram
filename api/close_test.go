package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type close struct{}

func (h close) Test(ctx context.Context, t *testing.T) context.Context {
	err := localAPI.Close(ctx)
	require.NoError(t, err)

	return ctx
}
