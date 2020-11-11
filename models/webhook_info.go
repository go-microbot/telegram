package models

// WebhookInfo represents information about the current status of a webhook.
// https://core.telegram.org/bots/api#webhookinfo.
type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up.
	URL string `json:"url"`
	// True, if a custom certificate was provided for webhook certificate checks.
	HasCustomCertificate bool `json:"has_custom_certificate"`
	// Number of updates awaiting delivery.
	PendingUpdateCount int32 `json:"pending_update_count"`
	// Optional. Currently used webhook IP address.
	IPAddress string `json:"ip_address,omitempty"`
	// Optional. Unix time for the most recent error that happened
	// when trying to deliver an update via webhook.
	LastErrorDate *int32 `json:"last_error_date,omitempty"`
	// Optional. Error message in human-readable format for the most
	// recent error that happened when trying to deliver an update via webhook.
	LastErrorMessage string `json:"last_error_message,omitempty"`
	// Optional. Maximum allowed number of simultaneous HTTPS connections
	// to the webhook for update delivery.
	MaxConnections *int32 `json:"max_connections,omitempty"`
	// Optional. A list of update types the bot is subscribed to. Defaults to all update types.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}
