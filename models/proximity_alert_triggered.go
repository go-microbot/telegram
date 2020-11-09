package models

// ProximityAlertTriggered represents the content of a service message,
// sent whenever a user in the chat triggers a proximity alert set by another user.
// https://core.telegram.org/bots/api#proximityalerttriggered.
type ProximityAlertTriggered struct {
	// User that triggered the alert.
	Traveler User `json:"traveler"`
	// User that set the alert.
	Watcher User `json:"watcher"`
	// The distance between the users.
	Distance int32 `json:"distance"`
}
