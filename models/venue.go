package models

// Venue represents a venue model.
// https://core.telegram.org/bots/api#venue.
type Venue struct {
	// Venue location. Can't be a live location.
	Location Location `json:"location"`
	// Name of the venue.
	Title string `json:"title"`
	// Address of the venue.
	Address string `json:"address"`
	// Optional. Foursquare identifier of the venue.
	FoursquareID string `json:"foursquare_id,omitempty"`
	// Optional. Foursquare type of the venue.
	// For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.
	FoursquareType string `json:"foursquare_type,omitempty"`
	// Optional. Google Places identifier of the venue.
	GooglePlaceID string `json:"google_place_id,omitempty"`
	// Optional. Google Places type of the venue.
	// (See supported types: https://developers.google.com/places/web-service/supported_types).
	GooglePlaceType string `json:"google_place_type,omitempty"`
}
