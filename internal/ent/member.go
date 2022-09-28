// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chatpuppy/puppychat/internal/ent/member"
)

// Member is the model entity for the Member schema.
type Member struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Intro holds the value of the "intro" field.
	Intro string `json:"intro,omitempty"`
	// PublicKey holds the value of the "public_key" field.
	PublicKey string `json:"public_key,omitempty"`
	// Nonce holds the value of the "nonce" field.
	Nonce string `json:"nonce,omitempty"`
	// ShowNickname holds the value of the "show_nickname" field.
	ShowNickname bool `json:"show_nickname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberQuery when eager-loading is set.
	Edges MemberEdges `json:"edges"`
}

// MemberEdges holds the relations/edges for other nodes in the graph.
type MemberEdges struct {
	// OwnGroups holds the value of the own_groups edge.
	OwnGroups []*Group `json:"own_groups,omitempty"`
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// GroupMembers holds the value of the group_members edge.
	GroupMembers []*GroupMember `json:"group_members,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// OwnGroupsOrErr returns the OwnGroups value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) OwnGroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[0] {
		return e.OwnGroups, nil
	}
	return nil, &NotLoadedError{edge: "own_groups"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[1] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[2] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// GroupMembersOrErr returns the GroupMembers value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) GroupMembersOrErr() ([]*GroupMember, error) {
	if e.loadedTypes[3] {
		return e.GroupMembers, nil
	}
	return nil, &NotLoadedError{edge: "group_members"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Member) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case member.FieldShowNickname:
			values[i] = new(sql.NullBool)
		case member.FieldID, member.FieldAddress, member.FieldNickname, member.FieldAvatar, member.FieldIntro, member.FieldPublicKey, member.FieldNonce:
			values[i] = new(sql.NullString)
		case member.FieldCreatedAt, member.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Member", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Member fields.
func (m *Member) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case member.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				m.ID = value.String
			}
		case member.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case member.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case member.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				m.Address = value.String
			}
		case member.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				m.Nickname = value.String
			}
		case member.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				m.Avatar = value.String
			}
		case member.FieldIntro:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field intro", values[i])
			} else if value.Valid {
				m.Intro = value.String
			}
		case member.FieldPublicKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_key", values[i])
			} else if value.Valid {
				m.PublicKey = value.String
			}
		case member.FieldNonce:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nonce", values[i])
			} else if value.Valid {
				m.Nonce = value.String
			}
		case member.FieldShowNickname:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field show_nickname", values[i])
			} else if value.Valid {
				m.ShowNickname = value.Bool
			}
		}
	}
	return nil
}

// QueryOwnGroups queries the "own_groups" edge of the Member entity.
func (m *Member) QueryOwnGroups() *GroupQuery {
	return (&MemberClient{config: m.config}).QueryOwnGroups(m)
}

// QueryMessages queries the "messages" edge of the Member entity.
func (m *Member) QueryMessages() *MessageQuery {
	return (&MemberClient{config: m.config}).QueryMessages(m)
}

// QueryGroups queries the "groups" edge of the Member entity.
func (m *Member) QueryGroups() *GroupQuery {
	return (&MemberClient{config: m.config}).QueryGroups(m)
}

// QueryGroupMembers queries the "group_members" edge of the Member entity.
func (m *Member) QueryGroupMembers() *GroupMemberQuery {
	return (&MemberClient{config: m.config}).QueryGroupMembers(m)
}

// Update returns a builder for updating this Member.
// Note that you need to call Member.Unwrap() before calling this method if this Member
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Member) Update() *MemberUpdateOne {
	return (&MemberClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Member entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Member) Unwrap() *Member {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Member is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Member) String() string {
	var builder strings.Builder
	builder.WriteString("Member(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(m.Address)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(m.Nickname)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(m.Avatar)
	builder.WriteString(", ")
	builder.WriteString("intro=")
	builder.WriteString(m.Intro)
	builder.WriteString(", ")
	builder.WriteString("public_key=")
	builder.WriteString(m.PublicKey)
	builder.WriteString(", ")
	builder.WriteString("nonce=")
	builder.WriteString(m.Nonce)
	builder.WriteString(", ")
	builder.WriteString("show_nickname=")
	builder.WriteString(fmt.Sprintf("%v", m.ShowNickname))
	builder.WriteByte(')')
	return builder.String()
}

// Members is a parsable slice of Member.
type Members []*Member

func (m Members) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
