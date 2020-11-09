package models

// PollOption represents a information about one answer option in a poll.
// https://core.telegram.org/bots/api#polloption.
type PollOption struct {
	// Option text, 1-100 characters.
	Text string `json:"text"`
	// Number of users that voted for this option.
	VoterCount int32 `json:"voter_count"`
}
