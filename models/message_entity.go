package models

// There are list of available message entity types.
const (
	MessageEntityTypeMention       MessageEntityType = "mention"
	MessageEntityTypeHashtag       MessageEntityType = "hashtag"
	MessageEntityTypeCashtag       MessageEntityType = "cashtag"
	MessageEntityTypeBotCommand    MessageEntityType = "bot_command"
	MessageEntityTypeURL           MessageEntityType = "url"
	MessageEntityTypeEmail         MessageEntityType = "email"
	MessageEntityTypePhoneNumber   MessageEntityType = "phone_number"
	MessageEntityTypeBold          MessageEntityType = "bold"
	MessageEntityTypeItalic        MessageEntityType = "italic"
	MessageEntityTypeUnderline     MessageEntityType = "underline"
	MessageEntityTypeStrikethrough MessageEntityType = "strikethrough"
	MessageEntityTypeCode          MessageEntityType = "code"
	MessageEntityTypePre           MessageEntityType = "pre"
	MessageEntityTypeTextLink      MessageEntityType = "text_link"
	MessageEntityTypeTextMention   MessageEntityType = "text_mention"
)

// MessageEntityType represents message entity type.
type MessageEntityType string

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
// https://core.telegram.org/bots/api#messageentity.
type MessageEntity struct {
	// Type of the entity.
	// Can be "mention" (@username), "hashtag" (#hashtag), "cashtag" ($USD),
	// "bot_command" (/start@jobs_bot), "url" (https://telegram.org),
	// "email" (do-not-reply@telegram.org), "phone_number" (+1-212-555-0123),
	// "bold" (bold text), "italic" (italic text), "underline" (underlined text),
	// "strikethrough" (strikethrough text), "code" (monowidth string), "pre" (monowidth block),
	// "text_link" (for clickable text URLs), "text_mention" (for users without usernames).
	Type MessageEntityType `json:"type"`
	// Offset in UTF-16 code units to the start of the entity.
	Offset int32 `json:"offset"`
	// Length of the entity in UTF-16 code units.
	Length int32 `json:"length"`
	// Optional. For “text_link” only, url that will be opened after user taps on the text.
	URL string `json:"url,omitempty"`
	// Optional. For “text_mention” only, the mentioned user.
	User *User `json:"user,omitempty"`
	// Optional. For “pre” only, the programming language of the entity text.
	Language string `json:"language,omitempty"`
}
