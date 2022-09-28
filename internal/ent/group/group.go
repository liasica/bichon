// Code generated by ent, DO NOT EDIT.

package group

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldMembersMax holds the string denoting the members_max field in the database.
	FieldMembersMax = "members_max"
	// FieldMembersCount holds the string denoting the members_count field in the database.
	FieldMembersCount = "members_count"
	// FieldPublic holds the string denoting the public field in the database.
	FieldPublic = "public"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldIntro holds the string denoting the intro field in the database.
	FieldIntro = "intro"
	// FieldKeys holds the string denoting the keys field in the database.
	FieldKeys = "keys"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeGroupMembers holds the string denoting the group_members edge name in mutations.
	EdgeGroupMembers = "group_members"
	// Table holds the table name of the group in the database.
	Table = "group"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "group"
	// OwnerInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	OwnerInverseTable = "member"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "message"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "message"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "group_id"
	// MembersTable is the table that holds the members relation/edge. The primary key declared below.
	MembersTable = "group_member"
	// MembersInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MembersInverseTable = "member"
	// GroupMembersTable is the table that holds the group_members relation/edge.
	GroupMembersTable = "group_member"
	// GroupMembersInverseTable is the table name for the GroupMember entity.
	// It exists in this package in order to avoid circular dependency with the "groupmember" package.
	GroupMembersInverseTable = "group_member"
	// GroupMembersColumn is the table column denoting the group_members relation/edge.
	GroupMembersColumn = "group_id"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldCategory,
	FieldOwnerID,
	FieldMembersMax,
	FieldMembersCount,
	FieldPublic,
	FieldAddress,
	FieldIntro,
	FieldKeys,
}

var (
	// MembersPrimaryKey and MembersColumn2 are the table columns denoting the
	// primary key for the members relation (M2M).
	MembersPrimaryKey = []string{"group_id", "member_id"}
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
	// DefaultMembersCount holds the default value on creation for the "members_count" field.
	DefaultMembersCount int
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
