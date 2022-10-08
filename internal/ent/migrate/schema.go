// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroupColumns holds the columns for the "group" table.
	GroupColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 25},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "category", Type: field.TypeString},
		{Name: "members_max", Type: field.TypeInt},
		{Name: "members_count", Type: field.TypeInt, Default: 1},
		{Name: "public", Type: field.TypeBool},
		{Name: "address", Type: field.TypeString, Unique: true},
		{Name: "intro", Type: field.TypeString, Nullable: true},
		{Name: "keys", Type: field.TypeString, Size: 2147483647},
		{Name: "owner_id", Type: field.TypeString, Size: 25},
	}
	// GroupTable holds the schema information for the "group" table.
	GroupTable = &schema.Table{
		Name:       "group",
		Columns:    GroupColumns,
		PrimaryKey: []*schema.Column{GroupColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_member_own_groups",
				Columns:    []*schema.Column{GroupColumns[10]},
				RefColumns: []*schema.Column{MemberColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "group_created_at",
				Unique:  false,
				Columns: []*schema.Column{GroupColumns[1]},
			},
			{
				Name:    "group_name",
				Unique:  false,
				Columns: []*schema.Column{GroupColumns[2]},
				Annotation: &entsql.IndexAnnotation{
					Types: map[string]string{
						"postgres": "GIN",
					},
				},
			},
			{
				Name:    "group_owner_id",
				Unique:  false,
				Columns: []*schema.Column{GroupColumns[10]},
			},
		},
	}
	// GroupMemberColumns holds the columns for the "group_member" table.
	GroupMemberColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 25},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "permission", Type: field.TypeOther, SchemaType: map[string]string{"postgres": "smallint"}},
		{Name: "invite_code", Type: field.TypeString, Unique: true},
		{Name: "invite_expire", Type: field.TypeTime},
		{Name: "read_id", Type: field.TypeString, Nullable: true},
		{Name: "read_time", Type: field.TypeTime, Nullable: true},
		{Name: "member_id", Type: field.TypeString, Size: 25},
		{Name: "group_id", Type: field.TypeString, Size: 25},
		{Name: "inviter_id", Type: field.TypeString, Nullable: true, Size: 25},
	}
	// GroupMemberTable holds the schema information for the "group_member" table.
	GroupMemberTable = &schema.Table{
		Name:       "group_member",
		Columns:    GroupMemberColumns,
		PrimaryKey: []*schema.Column{GroupMemberColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_member_member_member",
				Columns:    []*schema.Column{GroupMemberColumns[7]},
				RefColumns: []*schema.Column{MemberColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "group_member_group_group",
				Columns:    []*schema.Column{GroupMemberColumns[8]},
				RefColumns: []*schema.Column{GroupColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "group_member_member_inviter",
				Columns:    []*schema.Column{GroupMemberColumns[9]},
				RefColumns: []*schema.Column{MemberColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "groupmember_group_id_member_id",
				Unique:  true,
				Columns: []*schema.Column{GroupMemberColumns[8], GroupMemberColumns[7]},
			},
			{
				Name:    "groupmember_created_at",
				Unique:  false,
				Columns: []*schema.Column{GroupMemberColumns[1]},
			},
			{
				Name:    "groupmember_member_id",
				Unique:  false,
				Columns: []*schema.Column{GroupMemberColumns[7]},
			},
			{
				Name:    "groupmember_group_id",
				Unique:  false,
				Columns: []*schema.Column{GroupMemberColumns[8]},
			},
			{
				Name:    "groupmember_read_id",
				Unique:  false,
				Columns: []*schema.Column{GroupMemberColumns[5]},
			},
			{
				Name:    "groupmember_read_time",
				Unique:  false,
				Columns: []*schema.Column{GroupMemberColumns[6]},
			},
		},
	}
	// KeyColumns holds the columns for the "key" table.
	KeyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 25},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "keys", Type: field.TypeString, Size: 2147483647},
		{Name: "member_id", Type: field.TypeString, Size: 25},
		{Name: "group_id", Type: field.TypeString, Size: 25},
	}
	// KeyTable holds the schema information for the "key" table.
	KeyTable = &schema.Table{
		Name:       "key",
		Columns:    KeyColumns,
		PrimaryKey: []*schema.Column{KeyColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "key_member_member",
				Columns:    []*schema.Column{KeyColumns[3]},
				RefColumns: []*schema.Column{MemberColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "key_group_group",
				Columns:    []*schema.Column{KeyColumns[4]},
				RefColumns: []*schema.Column{GroupColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "key_created_at",
				Unique:  false,
				Columns: []*schema.Column{KeyColumns[1]},
			},
			{
				Name:    "key_member_id",
				Unique:  false,
				Columns: []*schema.Column{KeyColumns[3]},
			},
			{
				Name:    "key_group_id",
				Unique:  false,
				Columns: []*schema.Column{KeyColumns[4]},
			},
		},
	}
	// MemberColumns holds the columns for the "member" table.
	MemberColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 25},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "address", Type: field.TypeString, Unique: true},
		{Name: "nickname", Type: field.TypeString, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "intro", Type: field.TypeString, Nullable: true},
		{Name: "public_key", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "nonce", Type: field.TypeString},
		{Name: "show_nickname", Type: field.TypeBool, Default: true},
	}
	// MemberTable holds the schema information for the "member" table.
	MemberTable = &schema.Table{
		Name:       "member",
		Columns:    MemberColumns,
		PrimaryKey: []*schema.Column{MemberColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "member_created_at",
				Unique:  false,
				Columns: []*schema.Column{MemberColumns[1]},
			},
			{
				Name:    "member_nonce",
				Unique:  false,
				Columns: []*schema.Column{MemberColumns[7]},
			},
		},
	}
	// MessageColumns holds the columns for the "message" table.
	MessageColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 25},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeBytes},
		{Name: "owner", Type: field.TypeJSON},
		{Name: "group_id", Type: field.TypeString, Size: 25},
		{Name: "member_id", Type: field.TypeString, Size: 25},
		{Name: "parent_id", Type: field.TypeString, Nullable: true, Size: 25},
	}
	// MessageTable holds the schema information for the "message" table.
	MessageTable = &schema.Table{
		Name:       "message",
		Columns:    MessageColumns,
		PrimaryKey: []*schema.Column{MessageColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "message_group_messages",
				Columns:    []*schema.Column{MessageColumns[4]},
				RefColumns: []*schema.Column{GroupColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "message_member_messages",
				Columns:    []*schema.Column{MessageColumns[5]},
				RefColumns: []*schema.Column{MemberColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "message_message_children",
				Columns:    []*schema.Column{MessageColumns[6]},
				RefColumns: []*schema.Column{MessageColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "message_created_at",
				Unique:  false,
				Columns: []*schema.Column{MessageColumns[1]},
			},
			{
				Name:    "message_group_id",
				Unique:  false,
				Columns: []*schema.Column{MessageColumns[4]},
			},
			{
				Name:    "message_member_id",
				Unique:  false,
				Columns: []*schema.Column{MessageColumns[5]},
			},
			{
				Name:    "message_parent_id",
				Unique:  false,
				Columns: []*schema.Column{MessageColumns[6]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupTable,
		GroupMemberTable,
		KeyTable,
		MemberTable,
		MessageTable,
	}
)

func init() {
	GroupTable.ForeignKeys[0].RefTable = MemberTable
	GroupTable.Annotation = &entsql.Annotation{
		Table: "group",
	}
	GroupMemberTable.ForeignKeys[0].RefTable = MemberTable
	GroupMemberTable.ForeignKeys[1].RefTable = GroupTable
	GroupMemberTable.ForeignKeys[2].RefTable = MemberTable
	GroupMemberTable.Annotation = &entsql.Annotation{
		Table: "group_member",
	}
	KeyTable.ForeignKeys[0].RefTable = MemberTable
	KeyTable.ForeignKeys[1].RefTable = GroupTable
	KeyTable.Annotation = &entsql.Annotation{
		Table: "key",
	}
	MemberTable.Annotation = &entsql.Annotation{
		Table: "member",
	}
	MessageTable.ForeignKeys[0].RefTable = GroupTable
	MessageTable.ForeignKeys[1].RefTable = MemberTable
	MessageTable.ForeignKeys[2].RefTable = MessageTable
	MessageTable.Annotation = &entsql.Annotation{
		Table: "message",
	}
}
