package api

/*func TestTelegramAPI_GetWebhookInfo(t *testing.T) {
	// enable hook.
	cert := models.InputFile(webhookCert)
	err := localAPI.SetWebhook(context.Background(), apiModels.SetWebhookRequest{
		Certificate: &cert,
		URL:         query.NewParamString("localhost:1012"),
	})
	require.NoError(t, err)

	info, err := localAPI.GetWebhookInfo(context.Background())
	require.NoError(t, err)
	require.NotNil(t, info)
	require.Equal(t, "localhost:1012", info.URL)

	// disable hook.
	err = localAPI.SetWebhook(context.Background(), apiModels.SetWebhookRequest{})
	require.NoError(t, err)
}*/
