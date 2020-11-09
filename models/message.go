package models

// Message represents a telegram message model.
// https://core.telegram.org/bots/api#message.
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
	// Optional. Message is a general file, information about the file.
	Document *Document `json:"document,omitempty"`
	// Optional. Message is a photo, available sizes of the photo.
	Photo []PhotoSize `json:"photo,omitempty"`
	// Optional. Message is a sticker, information about the sticker.
	Sticker *Sticker `json:"sticker,omitempty"`
	// Optional. Message is a video, information about the video.
	Video *Video `json:"video,omitempty"`
	// Optional. Message is a video note (https://telegram.org/blog/video-messages-and-telescope),
	// information about the video message.
	VideoNote *VideoNote `json:"video_note,omitempty"`
	// Optional. Message is a voice message, information about the file.
	Voice *Voice `json:"voice,omitempty"`
	// Optional. Caption for the animation, audio, document,
	// photo, video or voice, 0-1024 characters.
	Caption string `json:"caption,omitempty"`
	// Optional. For messages with a caption, special entities like
	// usernames, URLs, bot commands, etc. that appear in the caption.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Optional. Message is a shared contact, information about the contact.
	Contact *Contact `json:"contact,omitempty"`
	// Optional. Message is a dice with random value from 1 to 6.
	Dice *Dice `json:"dice,omitempty"`
	// Optional. Message is a game, information about the game.
	Game *Game `json:"game,omitempty"`
	// Optional. Message is a native poll, information about the poll.
	Poll *Poll `json:"poll,omitempty"`
	// Optional. Message is a venue, information about the venue.
	// For backward compatibility, when this field is set, the location field will also be set.
	Venue *Venue `json:"venue,omitempty"`
	// Optional. Message is a shared location, information about the location.
	Location *Location `json:"location,omitempty"`
	// Optional. New members that were added to the group or supergroup
	// and information about them (the bot itself may be one of these members).
	NewChatMembers []User `json:"new_chat_members,omitempty"`
	// Optional. A member was removed from the group,
	// information about them (this member may be the bot itself).
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	// Optional. A chat title was changed to this value.
	NewChatTitle string `json:"new_chat_title,omitempty"`
	// Optional. A chat photo was change to this value.
	NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`
	// Optional. Service message: the chat photo was deleted.
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
	// Optional. Service message: the group has been created.
	GroupChatCreated bool `json:"group_chat_created,omitempty"`
	// Optional. Service message: the supergroup has been created.
	// This field can't be received in a message coming through updates,
	// because bot can't be a member of a supergroup when it is created.
	// It can only be found in reply_to_message if someone replies
	// to a very first message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
	// Optional. Service message: the channel has been created.
	// This field can't be received in a message coming through updates,
	// because bot can't be a member of a channel when it is created.
	// It can only be found in reply_to_message if someone replies
	// to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`
	// Optional. The group has been migrated to a supergroup with the specified identifier.
	MigrateToChatID *int64 `json:"migrate_to_chat_id,omitempty"`
	// Optional. The supergroup has been migrated from a group with the specified identifier.
	MigrateFromChatID *int64 `json:"migrate_from_chat_id,omitempty"`
	// Optional. Specified message was pinned.
	// Note that the Message object in this field will not contain
	// further reply_to_message fields even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// Optional. Message is an invoice for a payment (https://core.telegram.org/bots/api#payments),
	// information about the invoice.
	Invoice *Invoice `json:"invoice,omitempty"`
	// Optional. Message is a service message about a successful payment,
	// information about the payment.
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
	// Optional. The domain name of the website on which the user has logged in.
	ConnectedWebsite string `json:"connected_website,omitempty"`
	// Optional. Telegram Passport data.
	PassportData *PassportData `json:"passport_data,omitempty"`
	// Optional. Service message.
	// A user in the chat triggered another user's proximity alert while sharing Live Location.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
	// Optional. Inline keyboard attached to the message.
	// login_url buttons are represented as ordinary url buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
