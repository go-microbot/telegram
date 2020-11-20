package models

import "github.com/go-microbot/telegram/query"

// GetUserProfilePhotosRequest represents `getUserProfilePhotos` request body.
type GetUserProfilePhotosRequest struct {
	// Unique identifier of the target user.
	UserID query.ParamInt `query:"user_id"`
	// Optional. Sequential number of the first photo to be returned.
	// By default, all photos are returned.
	Offset query.ParamInt `query:"offset,omitempty"`
	// Optional. Limits the number of photos to be retrieved.
	// Values between 1-100 are accepted. Defaults to 100.
	Limit query.ParamInt `query:"limit,omitempty"`
}
