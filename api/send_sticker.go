package api

import (
	"context"
	"net/http"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendSticker represents method to send static .WEBP or animated .TGS stickers.
// On success, the sent Message is returned.
func (api *TelegramAPI) SendSticker(ctx context.Context, req apiModels.SendStickerRequest) (*models.Message, error) {
	resp, err := api.NewRequest("sendSticker").
		Query(query.AsMap(req)).
		Body(NewFormDataBody(req)).
		Method(http.MethodPost).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var msg models.Message
	if err := resp.Decode(&msg); err != nil {
		return nil, err
	}

	return &msg, nil
}
