package api

import (
	"context"

	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// GetUserProfilePhotos represents method to get a list of profile pictures for a user.
// Returns a UserProfilePhotos object.
func (api *TelegramAPI) GetUserProfilePhotos(ctx context.Context, req apiModels.GetUserProfilePhotosRequest) (*models.UserProfilePhotos, error) {
	resp, err := api.NewRequest("getUserProfilePhotos").
		Query(query.AsMap(req)).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var photos models.UserProfilePhotos
	if err := resp.Decode(&photos); err != nil {
		return nil, err
	}

	return &photos, nil
}
