package models

import (
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

// SendInvoiceRequest represents `sendInvoice` request body.
type SendInvoiceRequest struct {
	// Unique identifier for the target private chat.
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Product name, 1-32 characters.
	Title string `json:"title"`
	// Product description, 1-255 characters.
	Description string `json:"description"`
	// Bot-defined invoice payload, 1-128 bytes.
	// This will not be displayed to the user, use for your internal processes.
	Payload string `json:"payload"`
	// Payments provider token, obtained via Botfather (https://t.me/botfather).
	ProviderToken string `json:"provider_token"`
	// Unique deep-linking parameter that can be used to generate this invoice
	// when used as a start parameter.
	StartParameter string `json:"start_parameter"`
	// Three-letter ISO 4217 currency code, see more on currencies
	// (https://core.telegram.org/bots/payments#supported-currencies).
	Currency string `json:"currency"`
	// Price breakdown, a JSON-serialized list of components
	// (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.).
	Prices []models.LabeledPrice `json:"prices"`
	// Optional. A JSON-serialized data about the invoice,
	// which will be shared with the payment provider.
	// A detailed description of required fields should be provided by the payment provider.
	ProviderData interface{} `json:"provider_data,omitempty"`
	// Optional. URL of the product photo for the invoice.
	// Can be a photo of the goods or a marketing image for a service.
	// People like it better when they see what they are paying for.
	PhotoURL string `json:"photo_url,omitempty"`
	// Optional. Photo size.
	PhotoSize int32 `json:"photo_size,omitempty"`
	// Optional. Photo width.
	PhotoWidth int32 `json:"photo_width,omitempty"`
	// Optional. Photo height.
	PhotoHeight int32 `json:"photo_height,omitempty"`
	// Optional. Pass True, if you require the user's full name to complete the order.
	NeedName bool `json:"need_name,omitempty"`
	// Optional. Pass True, if you require the user's phone number to complete the order.
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
	// Optional. Pass True, if you require the user's email address to complete the order.
	NeedEmail bool `json:"need_email,omitempty"`
	// Optional. Pass True, if you require the user's shipping address to complete the order.
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
	// Optional. Pass True, if user's phone number should be sent to provider.
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
	// Optional. Pass True, if user's email address should be sent to provider.
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
	// Optional. Pass True, if the final price depends on the shipping method.
	IsFlexible bool `json:"is_flexible,omitempty"`
	// Optional. Sends the message
	// silently (https://telegram.org/blog/channels-2-0#silent-messages).
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message.
	ReplyToMessageID *int32 `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even
	// if the specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. A JSON-serialized object for an
	// inline keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
	// If empty, one 'Pay total price' button will be shown.
	// If not empty, the first button must be a Pay button.
	ReplyMarkup *models.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
