package models

// Message represents a telegram message model.
type Message struct {
	// Unique message identifier inside this chat.
	ID int32 `json:"message_id"`
	// Optional. Sender, empty for messages sent to channels.
	From *User `json:"from,omitempty"`
	// Optional. Sender of the message, sent on behalf of a chat.
	// The channel itself for channel messages.
	// The supergroup itself for messages from anonymous group administrators.
	// The linked channel for messages automatically forwarded to the discussion group.
	SenderChat *Chat `json:"sender_chat,omitempty"`
	// Optional. For forwarded messages, sender of the original message.
	ForwardFrom *User `json:"forward_from,omitempty"`
	// Optional. For messages forwarded from channels or from anonymous administrators,
	// information about the original sender chat.
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`
	// Optional. For messages forwarded from channels,
	// identifier of the original message in the channel.
	ForwardFromMessageID *int32 `json:"forward_from_message_id,omitempty"`
	// Optional. For messages forwarded from channels, signature of the post author if present.
	ForwardSignature string `json:"forward_signature,omitempty"`
	// Optional. Sender's name for messages forwarded from users
	// who disallow adding a link to their account in forwarded messages.
	ForwardSenderName string `json:"forward_sender_name,omitempty"`
	// Optional. For forwarded messages, date the original message was sent in Unix time.
	ForwardDate int32 `json:"forward_date,omitempty"`
	// Optional. For replies, the original message.
	// Note that the Message object in this field will not contain further reply_to_message fields
	// even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	// Optional. Bot through which the message was sent.
	ViaBot *User `json:"via_bot,omitempty"`
	// Optional. Date the message was last edited in Unix time.
	EditDate int32 `json:"edit_date,omitempty"`
	// Optional. The unique identifier of a media message group this message belongs to.
	MediaGroupID string `json:"media_group_id,omitempty"`
	// Optional. Signature of the post author for messages in channels,
	// or the custom title of an anonymous group administrator.
	AuthorSignature string `json:"author_signature,omitempty"`
	// Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Text string `json:"text,omitempty"`
	// Optional. For text messages, special entities like usernames,
	// URLs, bot commands, etc. that appear in the text.
	Entities []MessageEntity `json:"entities,omitempty"`
	// Optional. Message is an animation, information about the animation.
	// For backward compatibility, when this field is set, the document field will also be set.
	Animation *Animation `json:"animation,omitempty"`
	// Optional. Message is an audio file, information about the file.
	Audio *Audio `json:"audio,omitempty"`
}
