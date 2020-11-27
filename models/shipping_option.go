package models

// ShippingOption represents one shipping option.
// https://core.telegram.org/bots/api#shippingoption.
type ShippingOption struct {
	// Shipping option identifier.
	ID string `json:"id"`
	// Option title.
	Title string `json:"title"`
	// List of price portions.
	Prices []LabeledPrice `json:"prices"`
}
