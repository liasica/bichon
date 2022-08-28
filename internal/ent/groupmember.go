// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/groupmember"
	"github.com/chatpuppy/puppychat/internal/ent/member"
)

// GroupMember is the model entity for the GroupMember schema.
type GroupMember struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// MemberID holds the value of the "member_id" field.
	MemberID string `json:"member_id,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID string `json:"group_id,omitempty"`
	// member's permission in group
	Permission uint8 `json:"permission,omitempty"`
	// user's share sn
	Sn string `json:"sn,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupMemberQuery when eager-loading is set.
	Edges GroupMemberEdges `json:"edges"`
}

// GroupMemberEdges holds the relations/edges for other nodes in the graph.
type GroupMemberEdges struct {
	// Member holds the value of the member edge.
	Member *Member `json:"member,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MemberOrErr returns the Member value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMemberEdges) MemberOrErr() (*Member, error) {
	if e.loadedTypes[0] {
		if e.Member == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: member.Label}
		}
		return e.Member, nil
	}
	return nil, &NotLoadedError{edge: "member"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMemberEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[1] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupMember) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupmember.FieldPermission:
			values[i] = new(sql.NullInt64)
		case groupmember.FieldID, groupmember.FieldMemberID, groupmember.FieldGroupID, groupmember.FieldSn:
			values[i] = new(sql.NullString)
		case groupmember.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GroupMember", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupMember fields.
func (gm *GroupMember) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupmember.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				gm.ID = value.String
			}
		case groupmember.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gm.CreatedAt = value.Time
			}
		case groupmember.FieldMemberID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				gm.MemberID = value.String
			}
		case groupmember.FieldGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				gm.GroupID = value.String
			}
		case groupmember.FieldPermission:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field permission", values[i])
			} else if value.Valid {
				gm.Permission = uint8(value.Int64)
			}
		case groupmember.FieldSn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sn", values[i])
			} else if value.Valid {
				gm.Sn = value.String
			}
		}
	}
	return nil
}

// QueryMember queries the "member" edge of the GroupMember entity.
func (gm *GroupMember) QueryMember() *MemberQuery {
	return (&GroupMemberClient{config: gm.config}).QueryMember(gm)
}

// QueryGroup queries the "group" edge of the GroupMember entity.
func (gm *GroupMember) QueryGroup() *GroupQuery {
	return (&GroupMemberClient{config: gm.config}).QueryGroup(gm)
}

// Update returns a builder for updating this GroupMember.
// Note that you need to call GroupMember.Unwrap() before calling this method if this GroupMember
// was returned from a transaction, and the transaction was committed or rolled back.
func (gm *GroupMember) Update() *GroupMemberUpdateOne {
	return (&GroupMemberClient{config: gm.config}).UpdateOne(gm)
}

// Unwrap unwraps the GroupMember entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gm *GroupMember) Unwrap() *GroupMember {
	_tx, ok := gm.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupMember is not a transactional entity")
	}
	gm.config.driver = _tx.drv
	return gm
}

// String implements the fmt.Stringer.
func (gm *GroupMember) String() string {
	var builder strings.Builder
	builder.WriteString("GroupMember(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(gm.MemberID)
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(gm.GroupID)
	builder.WriteString(", ")
	builder.WriteString("permission=")
	builder.WriteString(fmt.Sprintf("%v", gm.Permission))
	builder.WriteString(", ")
	builder.WriteString("sn=")
	builder.WriteString(gm.Sn)
	builder.WriteByte(')')
	return builder.String()
}

// GroupMembers is a parsable slice of GroupMember.
type GroupMembers []*GroupMember

func (gm GroupMembers) config(cfg config) {
	for _i := range gm {
		gm[_i].config = cfg
	}
}
