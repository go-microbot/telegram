package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SetWebhookRequest represents `setWebhook` query params.
type SetWebhookRequest struct {
	// HTTPS url to send updates to. Use an empty string to remove webhook integration.
	URL query.ParamString `query:"url"`
	// Optional. Upload your public key certificate so
	// that the root certificate in use can be checked.
	// See our self-signed (https://core.telegram.org/bots/self-signed) guide for details.
	Certificate *models.InputFile
	// Optional. The fixed IP address which will be used to send webhook requests
	// instead of the IP address resolved through DNS.
	IPAddress query.ParamString `query:"ip_address,omitempty"`
	// Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook
	// for update delivery, 1-100. Defaults to 40.
	// Use lower values to limit the load on your bot's server,
	// and higher values to increase your bot's throughput.
	MaxConnections query.ParamInt `query:"max_connections,omitempty"`
	// Optional. A JSON-serialized list of the update types you want your bot to receive.
	// For example, specify ["message", "edited_channel_post", "callback_query"]
	// to only receive updates of these types.
	// See Update (https://core.telegram.org/bots/api#update) for a complete
	// list of available update types.
	// Specify an empty list to receive all updates regardless of type (default).
	// If not specified, the previous setting will be used.
	// Please note that this parameter doesn't affect updates created before the call
	// to the setWebhook, so unwanted updates may be received for a short period of time.
	AllowedUpdates query.ParamStringSlice `query:"allowed_updates,omitempty"`
	// Optional. Pass True to drop all pending updates.
	DropPendingUpdates query.ParamBool `query:"drop_pending_updates,omitempty"`
}
