package models

import "github.com/go-microbot/telegram/query"

// PromoteChatMemberRequest represents `promoteChatMember` request body.
type PromoteChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername).
	ChatID query.ParamAny `query:"chat_id" json:"-"`
	// Unique identifier of the target user.
	UserID query.ParamInt `query:"user_id" json:"-"`
	// Optional. Pass True, if the administrator's presence in the chat is hidden.
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// Optional. Pass True, if the administrator can change chat title,
	// photo and other settings.
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// Optional. Pass True, if the administrator can create channel posts, channels only.
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	// Optional. Pass True, if the administrator can edit messages of other users
	// and can pin messages, channels only.
	CanEditMessages bool `json:"can_edit_messages,omitempty"`
	// Optional. Pass True, if the administrator can delete messages of other users.
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
	// Optional. Pass True, if the administrator can invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// Optional. Pass True, if the administrator can restrict, ban or unban chat members.
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	// Optional. Pass True, if the administrator can pin messages, supergroups only.
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
	// Optional. Pass True, if the administrator can add new administrators
	// with a subset of their own privileges or demote administrators that he has promoted,
	// directly or indirectly (promoted by administrators that were appointed by him).
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`
}
