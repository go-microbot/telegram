package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewTelegramAPI(t *testing.T) {
	tAPI := NewTelegramAPI("123")
	require.NotNil(t, tAPI)
	require.Equal(t, "123", tAPI.token)
	require.NotNil(t, tAPI.client)
	require.Equal(t, baseURL, tAPI.url)
}
