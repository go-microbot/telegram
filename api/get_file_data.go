package api

import (
	"context"
	"fmt"
	"io"

	"github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/query"
)

// GetFileData downloads file by ID from Telegram server.
func (api *TelegramAPI) GetFileData(ctx context.Context, fileID string, dst io.Writer) error {
	f, err := api.GetFile(context.Background(), models.FileID{
		FileID: query.NewParamString(fileID),
	})
	if err != nil {
		return fmt.Errorf("could not get file by ID: %v", err)
	}

	resp, err := api.client.Get(fmt.Sprintf("%s/file/bot%s/%s", api.url, api.token, f.FilePath))
	if err != nil {
		return fmt.Errorf("could not download file: %v", err)
	}

	_, err = io.Copy(dst, resp.Body)
	if err != nil {
		return fmt.Errorf("could not copy file data to dst stream: %v", err)
	}

	return nil
}
