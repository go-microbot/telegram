package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// GetStickerSet represents method to get a sticker set.
// On success, a StickerSet object is returned.
func (api *TelegramAPI) GetStickerSet(ctx context.Context, req apiModels.GetStickerSetRequest) (*models.StickerSet, error) {
	resp, err := api.NewRequest("getStickerSet").
		Query(query.AsMap(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var set models.StickerSet
	if err := resp.Decode(&set); err != nil {
		return nil, err
	}

	return &set, nil
}
