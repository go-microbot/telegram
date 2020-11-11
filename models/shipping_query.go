package models

// ShippingQuery contains information about an incoming shipping query.
// https://core.telegram.org/bots/api#shippingquery.
type ShippingQuery struct {
	// Unique query identifier.
	ID string `json:"id"`
	// User who sent the query.
	From User `json:"from"`
	// Bot specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`
	// User specified shipping address.
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}
