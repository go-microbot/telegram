package models

// There are available chat member statuses.
const (
	ChatMemberStatusCreator       ChatMemberStatus = "creator"
	ChatMemberStatusAdministrator ChatMemberStatus = "administrator"
	ChatMemberStatusMember        ChatMemberStatus = "member"
	ChatMemberStatusRestricted    ChatMemberStatus = "restricted"
	ChatMemberStatusLeft          ChatMemberStatus = "left"
	ChatMemberStatusKicked        ChatMemberStatus = "kicked"
)

// ChatMemberStatus represents status for the chat member.
type ChatMemberStatus string

// ChatMember represents object contains information about one member of a chat.
// https://core.telegram.org/bots/api#chatmember.
type ChatMember struct {
	// Information about the user.
	User User `json:"user"`
	// The member's status in the chat.
	// Can be "creator", "administrator", "member", "restricted", "left" or "kicked".
	Status ChatMemberStatus `json:"status"`
	// Optional. Owner and administrators only. Custom title for this user.
	// TODO: add
}
