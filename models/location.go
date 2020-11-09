package models

// Location represents a point on the map.
// https://core.telegram.org/bots/api#location.
type Location struct {
	// Longitude as defined by sender.
	Longitude float64 `json:"longitude"`
	// Latitude as defined by sender.
	Latitude float64 `json:"latitude"`
	// Optional. The radius of uncertainty for the location, measured in meters; 0-1500.
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	// Optional. Time relative to the message sending date,
	// during which the location can be updated, in seconds. For active live locations only.
	LivePeriod int32 `json:"live_period,omitempty"`
	// Optional. The direction in which user is moving, in degrees; 1-360.
	// For active live locations only.
	Heading int32 `json:"heading,omitempty"`
	// Optional. Maximum distance for proximity alerts
	// about approaching another chat member, in meters. For sent live locations only.
	ProximityAlertRadius int32 `json:"proximity_alert_radius,omitempty"`
}
