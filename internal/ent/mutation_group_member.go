// Code generated, DO NOT EDIT.

package ent

import (
    "context"
    "fmt"
    "sync"
    "errors"
    "time"
    "github.com/chatpuppy/puppychat/internal/ent/groupmember"
    "github.com/chatpuppy/puppychat/internal/ent/predicate"

    "entgo.io/ent"
)


// GroupMemberMutation represents an operation that mutates the GroupMember nodes in the graph.
type GroupMemberMutation struct {
	config
	op            Op
	typ           string
	id            *string
	created_at    *time.Time
	permission    *uint8
	addpermission *int8
	sn            *string
	clearedFields map[string]struct{}
	member        *string
	clearedmember bool
	group         *string
	clearedgroup  bool
	done          bool
	oldValue      func(context.Context) (*GroupMember, error)
	predicates    []predicate.GroupMember
}

var _ ent.Mutation = (*GroupMemberMutation)(nil)

// groupmemberOption allows management of the mutation configuration using functional options.
type groupmemberOption func(*GroupMemberMutation)

// newGroupMemberMutation creates new mutation for the GroupMember entity.
func newGroupMemberMutation(c config, op Op, opts ...groupmemberOption) *GroupMemberMutation {
	m := &GroupMemberMutation{
		config:        c,
		op:            op,
		typ:           TypeGroupMember,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGroupMemberID sets the ID field of the mutation.
func withGroupMemberID(id string) groupmemberOption {
	return func(m *GroupMemberMutation) {
		var (
			err   error
			once  sync.Once
			value *GroupMember
		)
		m.oldValue = func(ctx context.Context) (*GroupMember, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().GroupMember.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGroupMember sets the old GroupMember of the mutation.
func withGroupMember(node *GroupMember) groupmemberOption {
	return func(m *GroupMemberMutation) {
		m.oldValue = func(context.Context) (*GroupMember, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GroupMemberMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GroupMemberMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of GroupMember entities.
func (m *GroupMemberMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *GroupMemberMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *GroupMemberMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().GroupMember.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *GroupMemberMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *GroupMemberMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the GroupMember entity.
// If the GroupMember object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMemberMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *GroupMemberMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetMemberID sets the "member_id" field.
func (m *GroupMemberMutation) SetMemberID(s string) {
	m.member = &s
}

// MemberID returns the value of the "member_id" field in the mutation.
func (m *GroupMemberMutation) MemberID() (r string, exists bool) {
	v := m.member
	if v == nil {
		return
	}
	return *v, true
}

// OldMemberID returns the old "member_id" field's value of the GroupMember entity.
// If the GroupMember object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMemberMutation) OldMemberID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMemberID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMemberID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMemberID: %w", err)
	}
	return oldValue.MemberID, nil
}

// ResetMemberID resets all changes to the "member_id" field.
func (m *GroupMemberMutation) ResetMemberID() {
	m.member = nil
}

// SetGroupID sets the "group_id" field.
func (m *GroupMemberMutation) SetGroupID(s string) {
	m.group = &s
}

// GroupID returns the value of the "group_id" field in the mutation.
func (m *GroupMemberMutation) GroupID() (r string, exists bool) {
	v := m.group
	if v == nil {
		return
	}
	return *v, true
}

// OldGroupID returns the old "group_id" field's value of the GroupMember entity.
// If the GroupMember object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMemberMutation) OldGroupID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGroupID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGroupID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGroupID: %w", err)
	}
	return oldValue.GroupID, nil
}

// ResetGroupID resets all changes to the "group_id" field.
func (m *GroupMemberMutation) ResetGroupID() {
	m.group = nil
}

// SetPermission sets the "permission" field.
func (m *GroupMemberMutation) SetPermission(u uint8) {
	m.permission = &u
	m.addpermission = nil
}

// Permission returns the value of the "permission" field in the mutation.
func (m *GroupMemberMutation) Permission() (r uint8, exists bool) {
	v := m.permission
	if v == nil {
		return
	}
	return *v, true
}

// OldPermission returns the old "permission" field's value of the GroupMember entity.
// If the GroupMember object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMemberMutation) OldPermission(ctx context.Context) (v uint8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPermission is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPermission requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPermission: %w", err)
	}
	return oldValue.Permission, nil
}

// AddPermission adds u to the "permission" field.
func (m *GroupMemberMutation) AddPermission(u int8) {
	if m.addpermission != nil {
		*m.addpermission += u
	} else {
		m.addpermission = &u
	}
}

// AddedPermission returns the value that was added to the "permission" field in this mutation.
func (m *GroupMemberMutation) AddedPermission() (r int8, exists bool) {
	v := m.addpermission
	if v == nil {
		return
	}
	return *v, true
}

// ResetPermission resets all changes to the "permission" field.
func (m *GroupMemberMutation) ResetPermission() {
	m.permission = nil
	m.addpermission = nil
}

// SetSn sets the "sn" field.
func (m *GroupMemberMutation) SetSn(s string) {
	m.sn = &s
}

// Sn returns the value of the "sn" field in the mutation.
func (m *GroupMemberMutation) Sn() (r string, exists bool) {
	v := m.sn
	if v == nil {
		return
	}
	return *v, true
}

// OldSn returns the old "sn" field's value of the GroupMember entity.
// If the GroupMember object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMemberMutation) OldSn(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSn is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSn requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSn: %w", err)
	}
	return oldValue.Sn, nil
}

// ResetSn resets all changes to the "sn" field.
func (m *GroupMemberMutation) ResetSn() {
	m.sn = nil
}

// ClearMember clears the "member" edge to the Member entity.
func (m *GroupMemberMutation) ClearMember() {
	m.clearedmember = true
}

// MemberCleared reports if the "member" edge to the Member entity was cleared.
func (m *GroupMemberMutation) MemberCleared() bool {
	return m.clearedmember
}

// MemberIDs returns the "member" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// MemberID instead. It exists only for internal usage by the builders.
func (m *GroupMemberMutation) MemberIDs() (ids []string) {
	if id := m.member; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetMember resets all changes to the "member" edge.
func (m *GroupMemberMutation) ResetMember() {
	m.member = nil
	m.clearedmember = false
}

// ClearGroup clears the "group" edge to the Group entity.
func (m *GroupMemberMutation) ClearGroup() {
	m.clearedgroup = true
}

// GroupCleared reports if the "group" edge to the Group entity was cleared.
func (m *GroupMemberMutation) GroupCleared() bool {
	return m.clearedgroup
}

// GroupIDs returns the "group" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// GroupID instead. It exists only for internal usage by the builders.
func (m *GroupMemberMutation) GroupIDs() (ids []string) {
	if id := m.group; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetGroup resets all changes to the "group" edge.
func (m *GroupMemberMutation) ResetGroup() {
	m.group = nil
	m.clearedgroup = false
}

// Where appends a list predicates to the GroupMemberMutation builder.
func (m *GroupMemberMutation) Where(ps ...predicate.GroupMember) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *GroupMemberMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (GroupMember).
func (m *GroupMemberMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *GroupMemberMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, groupmember.FieldCreatedAt)
	}
	if m.member != nil {
		fields = append(fields, groupmember.FieldMemberID)
	}
	if m.group != nil {
		fields = append(fields, groupmember.FieldGroupID)
	}
	if m.permission != nil {
		fields = append(fields, groupmember.FieldPermission)
	}
	if m.sn != nil {
		fields = append(fields, groupmember.FieldSn)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *GroupMemberMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case groupmember.FieldCreatedAt:
		return m.CreatedAt()
	case groupmember.FieldMemberID:
		return m.MemberID()
	case groupmember.FieldGroupID:
		return m.GroupID()
	case groupmember.FieldPermission:
		return m.Permission()
	case groupmember.FieldSn:
		return m.Sn()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *GroupMemberMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case groupmember.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case groupmember.FieldMemberID:
		return m.OldMemberID(ctx)
	case groupmember.FieldGroupID:
		return m.OldGroupID(ctx)
	case groupmember.FieldPermission:
		return m.OldPermission(ctx)
	case groupmember.FieldSn:
		return m.OldSn(ctx)
	}
	return nil, fmt.Errorf("unknown GroupMember field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GroupMemberMutation) SetField(name string, value ent.Value) error {
	switch name {
	case groupmember.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case groupmember.FieldMemberID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMemberID(v)
		return nil
	case groupmember.FieldGroupID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGroupID(v)
		return nil
	case groupmember.FieldPermission:
		v, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPermission(v)
		return nil
	case groupmember.FieldSn:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSn(v)
		return nil
	}
	return fmt.Errorf("unknown GroupMember field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *GroupMemberMutation) AddedFields() []string {
	var fields []string
	if m.addpermission != nil {
		fields = append(fields, groupmember.FieldPermission)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *GroupMemberMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case groupmember.FieldPermission:
		return m.AddedPermission()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GroupMemberMutation) AddField(name string, value ent.Value) error {
	switch name {
	case groupmember.FieldPermission:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPermission(v)
		return nil
	}
	return fmt.Errorf("unknown GroupMember numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *GroupMemberMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *GroupMemberMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *GroupMemberMutation) ClearField(name string) error {
	return fmt.Errorf("unknown GroupMember nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *GroupMemberMutation) ResetField(name string) error {
	switch name {
	case groupmember.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case groupmember.FieldMemberID:
		m.ResetMemberID()
		return nil
	case groupmember.FieldGroupID:
		m.ResetGroupID()
		return nil
	case groupmember.FieldPermission:
		m.ResetPermission()
		return nil
	case groupmember.FieldSn:
		m.ResetSn()
		return nil
	}
	return fmt.Errorf("unknown GroupMember field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *GroupMemberMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.member != nil {
		edges = append(edges, groupmember.EdgeMember)
	}
	if m.group != nil {
		edges = append(edges, groupmember.EdgeGroup)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *GroupMemberMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case groupmember.EdgeMember:
		if id := m.member; id != nil {
			return []ent.Value{*id}
		}
	case groupmember.EdgeGroup:
		if id := m.group; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *GroupMemberMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *GroupMemberMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *GroupMemberMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedmember {
		edges = append(edges, groupmember.EdgeMember)
	}
	if m.clearedgroup {
		edges = append(edges, groupmember.EdgeGroup)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *GroupMemberMutation) EdgeCleared(name string) bool {
	switch name {
	case groupmember.EdgeMember:
		return m.clearedmember
	case groupmember.EdgeGroup:
		return m.clearedgroup
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *GroupMemberMutation) ClearEdge(name string) error {
	switch name {
	case groupmember.EdgeMember:
		m.ClearMember()
		return nil
	case groupmember.EdgeGroup:
		m.ClearGroup()
		return nil
	}
	return fmt.Errorf("unknown GroupMember unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *GroupMemberMutation) ResetEdge(name string) error {
	switch name {
	case groupmember.EdgeMember:
		m.ResetMember()
		return nil
	case groupmember.EdgeGroup:
		m.ResetGroup()
		return nil
	}
	return fmt.Errorf("unknown GroupMember edge %s", name)
}

