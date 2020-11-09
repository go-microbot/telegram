package models

// There are list of available chat types.
const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSuperGroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

// ChatType represents chat type.
type ChatType string

// Chat represents a chat model.
type Chat struct {
	// Unique identifier for this chat.
	ID int64 `json:"id"`
	// Type of chat, can be either "private", "group", "supergroup" or "channel".
	Type ChatType `json:"type"`
	// Optional. Title, for supergroups, channels and group chats.
	Title string `json:"title,omitempty"`
	// Optional. Username, for private chats, supergroups and channels if available.
	Username string `json:"username,omitempty"`
	// Optional. First name of the other party in a private chat.
	FirstName string `json:"first_name,omitempty"`
	// Optional. Last name of the other party in a private chat.
	LastName string `json:"last_name,omitempty"`
	// Optional. Chat photo.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	Photo *ChatPhoto `json:"photo,omitempty"`
	// Optional. Bio of the other party in a private chat.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	BIO string `json:"bio,omitempty"`
	// Optional. Description, for groups, supergroups and channel chats.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	Description string `json:"description,omitempty"`
	// Optional. Chat invite link, for groups, supergroups and channel chats.
	// Each administrator in a chat generates their own invite links,
	// so the bot must first generate the link using
	// exportChatInviteLink (https://core.telegram.org/bots/api#exportchatinvitelink).
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	InviteLink string `json:"invite_link,omitempty"`
	// Optional. The most recent pinned message (by sending date).
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// Optional. Default chat member permissions, for groups and supergroups.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	Permissions *ChatPermissions `json:"permissions,omitempty"`
	// Optional. For supergroups, the minimum allowed delay
	// between consecutive messages sent by each unpriviledged user.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	SlowModeDelay int32 `json:"slow_mode_delay,omitempty"`
	// Optional. For supergroups, name of group sticker set.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	StickerSetName string `json:"sticker_set_name,omitempty"`
	// Optional. True, if the bot can change the group sticker set.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	CanSetStickerSet *bool `json:"can_set_sticker_set,omitempty"`
	// Optional. Unique identifier for the linked chat,
	// i.e. the discussion group identifier for a channel and vice versa;
	// for supergroups and channel chats.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	LinkedChatId *int64 `json:"linked_chat_id,omitempty"`
	// Optional. For supergroups, the location to which the supergroup is connected.
	// Returned only in getChat (https://core.telegram.org/bots/api#getchat).
	Location *ChatLocation `json:"location,omitempty"`
}
