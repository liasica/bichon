// Code generated by ent, DO NOT EDIT.

package member

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the member type in the database.
	Label = "member"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldIntro holds the string denoting the intro field in the database.
	FieldIntro = "intro"
	// FieldPublicKey holds the string denoting the public_key field in the database.
	FieldPublicKey = "public_key"
	// FieldNonce holds the string denoting the nonce field in the database.
	FieldNonce = "nonce"
	// FieldShowNickname holds the string denoting the show_nickname field in the database.
	FieldShowNickname = "show_nickname"
	// EdgeOwnGroups holds the string denoting the own_groups edge name in mutations.
	EdgeOwnGroups = "own_groups"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeGroupMembers holds the string denoting the group_members edge name in mutations.
	EdgeGroupMembers = "group_members"
	// Table holds the table name of the member in the database.
	Table = "member"
	// OwnGroupsTable is the table that holds the own_groups relation/edge.
	OwnGroupsTable = "group"
	// OwnGroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	OwnGroupsInverseTable = "group"
	// OwnGroupsColumn is the table column denoting the own_groups relation/edge.
	OwnGroupsColumn = "owner_id"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "message"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "message"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "member_id"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "group_member"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "group"
	// GroupMembersTable is the table that holds the group_members relation/edge.
	GroupMembersTable = "group_member"
	// GroupMembersInverseTable is the table name for the GroupMember entity.
	// It exists in this package in order to avoid circular dependency with the "groupmember" package.
	GroupMembersInverseTable = "group_member"
	// GroupMembersColumn is the table column denoting the group_members relation/edge.
	GroupMembersColumn = "member_id"
)

// Columns holds all SQL columns for member fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAddress,
	FieldNickname,
	FieldAvatar,
	FieldIntro,
	FieldPublicKey,
	FieldNonce,
	FieldShowNickname,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"group_id", "member_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/chatpuppy/puppychat/internal/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultShowNickname holds the default value on creation for the "show_nickname" field.
	DefaultShowNickname bool
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
