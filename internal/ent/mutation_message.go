// Code generated, DO NOT EDIT.

package ent

import (
    "context"
    "fmt"
    "sync"
    "errors"
    "time"
    "github.com/chatpuppy/puppychat/internal/ent/message"
    "github.com/chatpuppy/puppychat/internal/ent/predicate"

    "entgo.io/ent"
)


// MessageMutation represents an operation that mutates the Message nodes in the graph.
type MessageMutation struct {
	config
	op            Op
	typ           string
	id            *uint64
	created_at    *time.Time
	key_id        *uint64
	addkey_id     *int64
	content       *string
	clearedFields map[string]struct{}
	owner         *uint64
	clearedowner  bool
	group         *uint64
	clearedgroup  bool
	done          bool
	oldValue      func(context.Context) (*Message, error)
	predicates    []predicate.Message
}

var _ ent.Mutation = (*MessageMutation)(nil)

// messageOption allows management of the mutation configuration using functional options.
type messageOption func(*MessageMutation)

// newMessageMutation creates new mutation for the Message entity.
func newMessageMutation(c config, op Op, opts ...messageOption) *MessageMutation {
	m := &MessageMutation{
		config:        c,
		op:            op,
		typ:           TypeMessage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMessageID sets the ID field of the mutation.
func withMessageID(id uint64) messageOption {
	return func(m *MessageMutation) {
		var (
			err   error
			once  sync.Once
			value *Message
		)
		m.oldValue = func(ctx context.Context) (*Message, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Message.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMessage sets the old Message of the mutation.
func withMessage(node *Message) messageOption {
	return func(m *MessageMutation) {
		m.oldValue = func(context.Context) (*Message, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MessageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MessageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MessageMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MessageMutation) IDs(ctx context.Context) ([]uint64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Message.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *MessageMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *MessageMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *MessageMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetKeyID sets the "key_id" field.
func (m *MessageMutation) SetKeyID(u uint64) {
	m.key_id = &u
	m.addkey_id = nil
}

// KeyID returns the value of the "key_id" field in the mutation.
func (m *MessageMutation) KeyID() (r uint64, exists bool) {
	v := m.key_id
	if v == nil {
		return
	}
	return *v, true
}

// OldKeyID returns the old "key_id" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldKeyID(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldKeyID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldKeyID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldKeyID: %w", err)
	}
	return oldValue.KeyID, nil
}

// AddKeyID adds u to the "key_id" field.
func (m *MessageMutation) AddKeyID(u int64) {
	if m.addkey_id != nil {
		*m.addkey_id += u
	} else {
		m.addkey_id = &u
	}
}

// AddedKeyID returns the value that was added to the "key_id" field in this mutation.
func (m *MessageMutation) AddedKeyID() (r int64, exists bool) {
	v := m.addkey_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetKeyID resets all changes to the "key_id" field.
func (m *MessageMutation) ResetKeyID() {
	m.key_id = nil
	m.addkey_id = nil
}

// SetGroupID sets the "group_id" field.
func (m *MessageMutation) SetGroupID(u uint64) {
	m.group = &u
}

// GroupID returns the value of the "group_id" field in the mutation.
func (m *MessageMutation) GroupID() (r uint64, exists bool) {
	v := m.group
	if v == nil {
		return
	}
	return *v, true
}

// OldGroupID returns the old "group_id" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldGroupID(ctx context.Context) (v uint64, err error) {
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
func (m *MessageMutation) ResetGroupID() {
	m.group = nil
}

// SetMemberID sets the "member_id" field.
func (m *MessageMutation) SetMemberID(u uint64) {
	m.owner = &u
}

// MemberID returns the value of the "member_id" field in the mutation.
func (m *MessageMutation) MemberID() (r uint64, exists bool) {
	v := m.owner
	if v == nil {
		return
	}
	return *v, true
}

// OldMemberID returns the old "member_id" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldMemberID(ctx context.Context) (v uint64, err error) {
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
func (m *MessageMutation) ResetMemberID() {
	m.owner = nil
}

// SetContent sets the "content" field.
func (m *MessageMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *MessageMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *MessageMutation) ResetContent() {
	m.content = nil
}

// SetOwnerID sets the "owner" edge to the Member entity by id.
func (m *MessageMutation) SetOwnerID(id uint64) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the Member entity.
func (m *MessageMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the Member entity was cleared.
func (m *MessageMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the "owner" edge ID in the mutation.
func (m *MessageMutation) OwnerID() (id uint64, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *MessageMutation) OwnerIDs() (ids []uint64) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *MessageMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// ClearGroup clears the "group" edge to the Group entity.
func (m *MessageMutation) ClearGroup() {
	m.clearedgroup = true
}

// GroupCleared reports if the "group" edge to the Group entity was cleared.
func (m *MessageMutation) GroupCleared() bool {
	return m.clearedgroup
}

// GroupIDs returns the "group" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// GroupID instead. It exists only for internal usage by the builders.
func (m *MessageMutation) GroupIDs() (ids []uint64) {
	if id := m.group; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetGroup resets all changes to the "group" edge.
func (m *MessageMutation) ResetGroup() {
	m.group = nil
	m.clearedgroup = false
}

// Where appends a list predicates to the MessageMutation builder.
func (m *MessageMutation) Where(ps ...predicate.Message) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *MessageMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Message).
func (m *MessageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MessageMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, message.FieldCreatedAt)
	}
	if m.key_id != nil {
		fields = append(fields, message.FieldKeyID)
	}
	if m.group != nil {
		fields = append(fields, message.FieldGroupID)
	}
	if m.owner != nil {
		fields = append(fields, message.FieldMemberID)
	}
	if m.content != nil {
		fields = append(fields, message.FieldContent)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MessageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case message.FieldCreatedAt:
		return m.CreatedAt()
	case message.FieldKeyID:
		return m.KeyID()
	case message.FieldGroupID:
		return m.GroupID()
	case message.FieldMemberID:
		return m.MemberID()
	case message.FieldContent:
		return m.Content()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MessageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case message.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case message.FieldKeyID:
		return m.OldKeyID(ctx)
	case message.FieldGroupID:
		return m.OldGroupID(ctx)
	case message.FieldMemberID:
		return m.OldMemberID(ctx)
	case message.FieldContent:
		return m.OldContent(ctx)
	}
	return nil, fmt.Errorf("unknown Message field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case message.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case message.FieldKeyID:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetKeyID(v)
		return nil
	case message.FieldGroupID:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGroupID(v)
		return nil
	case message.FieldMemberID:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMemberID(v)
		return nil
	case message.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	}
	return fmt.Errorf("unknown Message field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MessageMutation) AddedFields() []string {
	var fields []string
	if m.addkey_id != nil {
		fields = append(fields, message.FieldKeyID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MessageMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case message.FieldKeyID:
		return m.AddedKeyID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageMutation) AddField(name string, value ent.Value) error {
	switch name {
	case message.FieldKeyID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddKeyID(v)
		return nil
	}
	return fmt.Errorf("unknown Message numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MessageMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MessageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MessageMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Message nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MessageMutation) ResetField(name string) error {
	switch name {
	case message.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case message.FieldKeyID:
		m.ResetKeyID()
		return nil
	case message.FieldGroupID:
		m.ResetGroupID()
		return nil
	case message.FieldMemberID:
		m.ResetMemberID()
		return nil
	case message.FieldContent:
		m.ResetContent()
		return nil
	}
	return fmt.Errorf("unknown Message field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MessageMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.owner != nil {
		edges = append(edges, message.EdgeOwner)
	}
	if m.group != nil {
		edges = append(edges, message.EdgeGroup)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MessageMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case message.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	case message.EdgeGroup:
		if id := m.group; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MessageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MessageMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MessageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedowner {
		edges = append(edges, message.EdgeOwner)
	}
	if m.clearedgroup {
		edges = append(edges, message.EdgeGroup)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MessageMutation) EdgeCleared(name string) bool {
	switch name {
	case message.EdgeOwner:
		return m.clearedowner
	case message.EdgeGroup:
		return m.clearedgroup
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MessageMutation) ClearEdge(name string) error {
	switch name {
	case message.EdgeOwner:
		m.ClearOwner()
		return nil
	case message.EdgeGroup:
		m.ClearGroup()
		return nil
	}
	return fmt.Errorf("unknown Message unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MessageMutation) ResetEdge(name string) error {
	switch name {
	case message.EdgeOwner:
		m.ResetOwner()
		return nil
	case message.EdgeGroup:
		m.ResetGroup()
		return nil
	}
	return fmt.Errorf("unknown Message edge %s", name)
}

