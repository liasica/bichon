// Code generated, DO NOT EDIT.

package ent

import (
    "context"
    "fmt"
    "sync"
    "errors"
    "time"
    "github.com/chatpuppy/puppychat/internal/ent/key"
    "github.com/chatpuppy/puppychat/internal/ent/predicate"

    "entgo.io/ent"
)


// KeyMutation represents an operation that mutates the Key nodes in the graph.
type KeyMutation struct {
	config
	op            Op
	typ           string
	id            *string
	created_at    *time.Time
	keys          *string
	last_node     *int64
	addlast_node  *int64
	clearedFields map[string]struct{}
	member        *string
	clearedmember bool
	group         *string
	clearedgroup  bool
	done          bool
	oldValue      func(context.Context) (*Key, error)
	predicates    []predicate.Key
}

var _ ent.Mutation = (*KeyMutation)(nil)

// keyOption allows management of the mutation configuration using functional options.
type keyOption func(*KeyMutation)

// newKeyMutation creates new mutation for the Key entity.
func newKeyMutation(c config, op Op, opts ...keyOption) *KeyMutation {
	m := &KeyMutation{
		config:        c,
		op:            op,
		typ:           TypeKey,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withKeyID sets the ID field of the mutation.
func withKeyID(id string) keyOption {
	return func(m *KeyMutation) {
		var (
			err   error
			once  sync.Once
			value *Key
		)
		m.oldValue = func(ctx context.Context) (*Key, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Key.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withKey sets the old Key of the mutation.
func withKey(node *Key) keyOption {
	return func(m *KeyMutation) {
		m.oldValue = func(context.Context) (*Key, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m KeyMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m KeyMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Key entities.
func (m *KeyMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *KeyMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *KeyMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Key.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *KeyMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *KeyMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Key entity.
// If the Key object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KeyMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *KeyMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetMemberID sets the "member_id" field.
func (m *KeyMutation) SetMemberID(s string) {
	m.member = &s
}

// MemberID returns the value of the "member_id" field in the mutation.
func (m *KeyMutation) MemberID() (r string, exists bool) {
	v := m.member
	if v == nil {
		return
	}
	return *v, true
}

// OldMemberID returns the old "member_id" field's value of the Key entity.
// If the Key object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KeyMutation) OldMemberID(ctx context.Context) (v string, err error) {
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
func (m *KeyMutation) ResetMemberID() {
	m.member = nil
}

// SetGroupID sets the "group_id" field.
func (m *KeyMutation) SetGroupID(s string) {
	m.group = &s
}

// GroupID returns the value of the "group_id" field in the mutation.
func (m *KeyMutation) GroupID() (r string, exists bool) {
	v := m.group
	if v == nil {
		return
	}
	return *v, true
}

// OldGroupID returns the old "group_id" field's value of the Key entity.
// If the Key object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KeyMutation) OldGroupID(ctx context.Context) (v string, err error) {
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
func (m *KeyMutation) ResetGroupID() {
	m.group = nil
}

// SetKeys sets the "keys" field.
func (m *KeyMutation) SetKeys(s string) {
	m.keys = &s
}

// Keys returns the value of the "keys" field in the mutation.
func (m *KeyMutation) Keys() (r string, exists bool) {
	v := m.keys
	if v == nil {
		return
	}
	return *v, true
}

// OldKeys returns the old "keys" field's value of the Key entity.
// If the Key object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KeyMutation) OldKeys(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldKeys is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldKeys requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldKeys: %w", err)
	}
	return oldValue.Keys, nil
}

// ResetKeys resets all changes to the "keys" field.
func (m *KeyMutation) ResetKeys() {
	m.keys = nil
}

// SetLastNode sets the "last_node" field.
func (m *KeyMutation) SetLastNode(i int64) {
	m.last_node = &i
	m.addlast_node = nil
}

// LastNode returns the value of the "last_node" field in the mutation.
func (m *KeyMutation) LastNode() (r int64, exists bool) {
	v := m.last_node
	if v == nil {
		return
	}
	return *v, true
}

// OldLastNode returns the old "last_node" field's value of the Key entity.
// If the Key object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KeyMutation) OldLastNode(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastNode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastNode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastNode: %w", err)
	}
	return oldValue.LastNode, nil
}

// AddLastNode adds i to the "last_node" field.
func (m *KeyMutation) AddLastNode(i int64) {
	if m.addlast_node != nil {
		*m.addlast_node += i
	} else {
		m.addlast_node = &i
	}
}

// AddedLastNode returns the value that was added to the "last_node" field in this mutation.
func (m *KeyMutation) AddedLastNode() (r int64, exists bool) {
	v := m.addlast_node
	if v == nil {
		return
	}
	return *v, true
}

// ResetLastNode resets all changes to the "last_node" field.
func (m *KeyMutation) ResetLastNode() {
	m.last_node = nil
	m.addlast_node = nil
}

// ClearMember clears the "member" edge to the Member entity.
func (m *KeyMutation) ClearMember() {
	m.clearedmember = true
}

// MemberCleared reports if the "member" edge to the Member entity was cleared.
func (m *KeyMutation) MemberCleared() bool {
	return m.clearedmember
}

// MemberIDs returns the "member" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// MemberID instead. It exists only for internal usage by the builders.
func (m *KeyMutation) MemberIDs() (ids []string) {
	if id := m.member; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetMember resets all changes to the "member" edge.
func (m *KeyMutation) ResetMember() {
	m.member = nil
	m.clearedmember = false
}

// ClearGroup clears the "group" edge to the Group entity.
func (m *KeyMutation) ClearGroup() {
	m.clearedgroup = true
}

// GroupCleared reports if the "group" edge to the Group entity was cleared.
func (m *KeyMutation) GroupCleared() bool {
	return m.clearedgroup
}

// GroupIDs returns the "group" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// GroupID instead. It exists only for internal usage by the builders.
func (m *KeyMutation) GroupIDs() (ids []string) {
	if id := m.group; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetGroup resets all changes to the "group" edge.
func (m *KeyMutation) ResetGroup() {
	m.group = nil
	m.clearedgroup = false
}

// Where appends a list predicates to the KeyMutation builder.
func (m *KeyMutation) Where(ps ...predicate.Key) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *KeyMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Key).
func (m *KeyMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *KeyMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, key.FieldCreatedAt)
	}
	if m.member != nil {
		fields = append(fields, key.FieldMemberID)
	}
	if m.group != nil {
		fields = append(fields, key.FieldGroupID)
	}
	if m.keys != nil {
		fields = append(fields, key.FieldKeys)
	}
	if m.last_node != nil {
		fields = append(fields, key.FieldLastNode)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *KeyMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case key.FieldCreatedAt:
		return m.CreatedAt()
	case key.FieldMemberID:
		return m.MemberID()
	case key.FieldGroupID:
		return m.GroupID()
	case key.FieldKeys:
		return m.Keys()
	case key.FieldLastNode:
		return m.LastNode()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *KeyMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case key.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case key.FieldMemberID:
		return m.OldMemberID(ctx)
	case key.FieldGroupID:
		return m.OldGroupID(ctx)
	case key.FieldKeys:
		return m.OldKeys(ctx)
	case key.FieldLastNode:
		return m.OldLastNode(ctx)
	}
	return nil, fmt.Errorf("unknown Key field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *KeyMutation) SetField(name string, value ent.Value) error {
	switch name {
	case key.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case key.FieldMemberID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMemberID(v)
		return nil
	case key.FieldGroupID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGroupID(v)
		return nil
	case key.FieldKeys:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetKeys(v)
		return nil
	case key.FieldLastNode:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastNode(v)
		return nil
	}
	return fmt.Errorf("unknown Key field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *KeyMutation) AddedFields() []string {
	var fields []string
	if m.addlast_node != nil {
		fields = append(fields, key.FieldLastNode)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *KeyMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case key.FieldLastNode:
		return m.AddedLastNode()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *KeyMutation) AddField(name string, value ent.Value) error {
	switch name {
	case key.FieldLastNode:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLastNode(v)
		return nil
	}
	return fmt.Errorf("unknown Key numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *KeyMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *KeyMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *KeyMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Key nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *KeyMutation) ResetField(name string) error {
	switch name {
	case key.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case key.FieldMemberID:
		m.ResetMemberID()
		return nil
	case key.FieldGroupID:
		m.ResetGroupID()
		return nil
	case key.FieldKeys:
		m.ResetKeys()
		return nil
	case key.FieldLastNode:
		m.ResetLastNode()
		return nil
	}
	return fmt.Errorf("unknown Key field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *KeyMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.member != nil {
		edges = append(edges, key.EdgeMember)
	}
	if m.group != nil {
		edges = append(edges, key.EdgeGroup)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *KeyMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case key.EdgeMember:
		if id := m.member; id != nil {
			return []ent.Value{*id}
		}
	case key.EdgeGroup:
		if id := m.group; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *KeyMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *KeyMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *KeyMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedmember {
		edges = append(edges, key.EdgeMember)
	}
	if m.clearedgroup {
		edges = append(edges, key.EdgeGroup)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *KeyMutation) EdgeCleared(name string) bool {
	switch name {
	case key.EdgeMember:
		return m.clearedmember
	case key.EdgeGroup:
		return m.clearedgroup
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *KeyMutation) ClearEdge(name string) error {
	switch name {
	case key.EdgeMember:
		m.ClearMember()
		return nil
	case key.EdgeGroup:
		m.ClearGroup()
		return nil
	}
	return fmt.Errorf("unknown Key unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *KeyMutation) ResetEdge(name string) error {
	switch name {
	case key.EdgeMember:
		m.ResetMember()
		return nil
	case key.EdgeGroup:
		m.ResetGroup()
		return nil
	}
	return fmt.Errorf("unknown Key edge %s", name)
}

