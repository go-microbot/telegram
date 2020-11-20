package models

// UserProfilePhotos represent a user's profile pictures.
// https://core.telegram.org/bots/api#userprofilephotos.
type UserProfilePhotos struct {
	// Total number of profile pictures the target user has.
	TotalCount int32 `json:"total_count"`
	// Requested profile pictures (in up to 4 sizes each).
	Photos []PhotoSize `json:"photos"`
}
