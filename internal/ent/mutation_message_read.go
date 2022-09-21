// Code generated, DO NOT EDIT.

package ent

import (
    "context"
    "fmt"
    "sync"
    "errors"
    "time"
    "github.com/chatpuppy/puppychat/internal/ent/messageread"
    "github.com/chatpuppy/puppychat/internal/ent/predicate"

    "entgo.io/ent"
)


// MessageReadMutation represents an operation that mutates the MessageRead nodes in the graph.
type MessageReadMutation struct {
	config
	op            Op
	typ           string
	id            *string
	last_id       *string
	last_time     *time.Time
	clearedFields map[string]struct{}
	member        *string
	clearedmember bool
	group         *string
	clearedgroup  bool
	done          bool
	oldValue      func(context.Context) (*MessageRead, error)
	predicates    []predicate.MessageRead
}

var _ ent.Mutation = (*MessageReadMutation)(nil)

// messagereadOption allows management of the mutation configuration using functional options.
type messagereadOption func(*MessageReadMutation)

// newMessageReadMutation creates new mutation for the MessageRead entity.
func newMessageReadMutation(c config, op Op, opts ...messagereadOption) *MessageReadMutation {
	m := &MessageReadMutation{
		config:        c,
		op:            op,
		typ:           TypeMessageRead,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMessageReadID sets the ID field of the mutation.
func withMessageReadID(id string) messagereadOption {
	return func(m *MessageReadMutation) {
		var (
			err   error
			once  sync.Once
			value *MessageRead
		)
		m.oldValue = func(ctx context.Context) (*MessageRead, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().MessageRead.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMessageRead sets the old MessageRead of the mutation.
func withMessageRead(node *MessageRead) messagereadOption {
	return func(m *MessageReadMutation) {
		m.oldValue = func(context.Context) (*MessageRead, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MessageReadMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MessageReadMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of MessageRead entities.
func (m *MessageReadMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MessageReadMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MessageReadMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().MessageRead.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetMemberID sets the "member_id" field.
func (m *MessageReadMutation) SetMemberID(s string) {
	m.member = &s
}

// MemberID returns the value of the "member_id" field in the mutation.
func (m *MessageReadMutation) MemberID() (r string, exists bool) {
	v := m.member
	if v == nil {
		return
	}
	return *v, true
}

// OldMemberID returns the old "member_id" field's value of the MessageRead entity.
// If the MessageRead object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageReadMutation) OldMemberID(ctx context.Context) (v string, err error) {
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
func (m *MessageReadMutation) ResetMemberID() {
	m.member = nil
}

// SetGroupID sets the "group_id" field.
func (m *MessageReadMutation) SetGroupID(s string) {
	m.group = &s
}

// GroupID returns the value of the "group_id" field in the mutation.
func (m *MessageReadMutation) GroupID() (r string, exists bool) {
	v := m.group
	if v == nil {
		return
	}
	return *v, true
}

// OldGroupID returns the old "group_id" field's value of the MessageRead entity.
// If the MessageRead object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageReadMutation) OldGroupID(ctx context.Context) (v string, err error) {
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
func (m *MessageReadMutation) ResetGroupID() {
	m.group = nil
}

// SetLastID sets the "last_id" field.
func (m *MessageReadMutation) SetLastID(s string) {
	m.last_id = &s
}

// LastID returns the value of the "last_id" field in the mutation.
func (m *MessageReadMutation) LastID() (r string, exists bool) {
	v := m.last_id
	if v == nil {
		return
	}
	return *v, true
}

// OldLastID returns the old "last_id" field's value of the MessageRead entity.
// If the MessageRead object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageReadMutation) OldLastID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastID: %w", err)
	}
	return oldValue.LastID, nil
}

// ResetLastID resets all changes to the "last_id" field.
func (m *MessageReadMutation) ResetLastID() {
	m.last_id = nil
}

// SetLastTime sets the "last_time" field.
func (m *MessageReadMutation) SetLastTime(t time.Time) {
	m.last_time = &t
}

// LastTime returns the value of the "last_time" field in the mutation.
func (m *MessageReadMutation) LastTime() (r time.Time, exists bool) {
	v := m.last_time
	if v == nil {
		return
	}
	return *v, true
}

// OldLastTime returns the old "last_time" field's value of the MessageRead entity.
// If the MessageRead object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageReadMutation) OldLastTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastTime: %w", err)
	}
	return oldValue.LastTime, nil
}

// ResetLastTime resets all changes to the "last_time" field.
func (m *MessageReadMutation) ResetLastTime() {
	m.last_time = nil
}

// ClearMember clears the "member" edge to the Member entity.
func (m *MessageReadMutation) ClearMember() {
	m.clearedmember = true
}

// MemberCleared reports if the "member" edge to the Member entity was cleared.
func (m *MessageReadMutation) MemberCleared() bool {
	return m.clearedmember
}

// MemberIDs returns the "member" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// MemberID instead. It exists only for internal usage by the builders.
func (m *MessageReadMutation) MemberIDs() (ids []string) {
	if id := m.member; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetMember resets all changes to the "member" edge.
func (m *MessageReadMutation) ResetMember() {
	m.member = nil
	m.clearedmember = false
}

// ClearGroup clears the "group" edge to the Group entity.
func (m *MessageReadMutation) ClearGroup() {
	m.clearedgroup = true
}

// GroupCleared reports if the "group" edge to the Group entity was cleared.
func (m *MessageReadMutation) GroupCleared() bool {
	return m.clearedgroup
}

// GroupIDs returns the "group" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// GroupID instead. It exists only for internal usage by the builders.
func (m *MessageReadMutation) GroupIDs() (ids []string) {
	if id := m.group; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetGroup resets all changes to the "group" edge.
func (m *MessageReadMutation) ResetGroup() {
	m.group = nil
	m.clearedgroup = false
}

// Where appends a list predicates to the MessageReadMutation builder.
func (m *MessageReadMutation) Where(ps ...predicate.MessageRead) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *MessageReadMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (MessageRead).
func (m *MessageReadMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MessageReadMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.member != nil {
		fields = append(fields, messageread.FieldMemberID)
	}
	if m.group != nil {
		fields = append(fields, messageread.FieldGroupID)
	}
	if m.last_id != nil {
		fields = append(fields, messageread.FieldLastID)
	}
	if m.last_time != nil {
		fields = append(fields, messageread.FieldLastTime)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MessageReadMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case messageread.FieldMemberID:
		return m.MemberID()
	case messageread.FieldGroupID:
		return m.GroupID()
	case messageread.FieldLastID:
		return m.LastID()
	case messageread.FieldLastTime:
		return m.LastTime()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MessageReadMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case messageread.FieldMemberID:
		return m.OldMemberID(ctx)
	case messageread.FieldGroupID:
		return m.OldGroupID(ctx)
	case messageread.FieldLastID:
		return m.OldLastID(ctx)
	case messageread.FieldLastTime:
		return m.OldLastTime(ctx)
	}
	return nil, fmt.Errorf("unknown MessageRead field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageReadMutation) SetField(name string, value ent.Value) error {
	switch name {
	case messageread.FieldMemberID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMemberID(v)
		return nil
	case messageread.FieldGroupID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGroupID(v)
		return nil
	case messageread.FieldLastID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastID(v)
		return nil
	case messageread.FieldLastTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastTime(v)
		return nil
	}
	return fmt.Errorf("unknown MessageRead field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MessageReadMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MessageReadMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageReadMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown MessageRead numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MessageReadMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MessageReadMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MessageReadMutation) ClearField(name string) error {
	return fmt.Errorf("unknown MessageRead nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MessageReadMutation) ResetField(name string) error {
	switch name {
	case messageread.FieldMemberID:
		m.ResetMemberID()
		return nil
	case messageread.FieldGroupID:
		m.ResetGroupID()
		return nil
	case messageread.FieldLastID:
		m.ResetLastID()
		return nil
	case messageread.FieldLastTime:
		m.ResetLastTime()
		return nil
	}
	return fmt.Errorf("unknown MessageRead field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MessageReadMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.member != nil {
		edges = append(edges, messageread.EdgeMember)
	}
	if m.group != nil {
		edges = append(edges, messageread.EdgeGroup)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MessageReadMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case messageread.EdgeMember:
		if id := m.member; id != nil {
			return []ent.Value{*id}
		}
	case messageread.EdgeGroup:
		if id := m.group; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MessageReadMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MessageReadMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MessageReadMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedmember {
		edges = append(edges, messageread.EdgeMember)
	}
	if m.clearedgroup {
		edges = append(edges, messageread.EdgeGroup)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MessageReadMutation) EdgeCleared(name string) bool {
	switch name {
	case messageread.EdgeMember:
		return m.clearedmember
	case messageread.EdgeGroup:
		return m.clearedgroup
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MessageReadMutation) ClearEdge(name string) error {
	switch name {
	case messageread.EdgeMember:
		m.ClearMember()
		return nil
	case messageread.EdgeGroup:
		m.ClearGroup()
		return nil
	}
	return fmt.Errorf("unknown MessageRead unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MessageReadMutation) ResetEdge(name string) error {
	switch name {
	case messageread.EdgeMember:
		m.ResetMember()
		return nil
	case messageread.EdgeGroup:
		m.ResetGroup()
		return nil
	}
	return fmt.Errorf("unknown MessageRead edge %s", name)
}

