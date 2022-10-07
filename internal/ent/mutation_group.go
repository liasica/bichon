// Code generated, DO NOT EDIT.

package ent

import (
    "context"
    "fmt"
    "sync"
    "errors"
    "time"
    "github.com/chatpuppy/puppychat/internal/ent/group"
    "github.com/chatpuppy/puppychat/internal/ent/predicate"

    "entgo.io/ent"
)


// GroupMutation represents an operation that mutates the Group nodes in the graph.
type GroupMutation struct {
	config
	op                   Op
	typ                  string
	id                   *string
	created_at           *time.Time
	name                 *string
	category             *string
	members_max          *int
	addmembers_max       *int
	members_count        *int
	addmembers_count     *int
	public               *bool
	address              *string
	intro                *string
	keys                 *string
	last_node            *int64
	addlast_node         *int64
	clearedFields        map[string]struct{}
	owner                *string
	clearedowner         bool
	messages             map[string]struct{}
	removedmessages      map[string]struct{}
	clearedmessages      bool
	members              map[string]struct{}
	removedmembers       map[string]struct{}
	clearedmembers       bool
	group_members        map[string]struct{}
	removedgroup_members map[string]struct{}
	clearedgroup_members bool
	done                 bool
	oldValue             func(context.Context) (*Group, error)
	predicates           []predicate.Group
}

var _ ent.Mutation = (*GroupMutation)(nil)

// groupOption allows management of the mutation configuration using functional options.
type groupOption func(*GroupMutation)

// newGroupMutation creates new mutation for the Group entity.
func newGroupMutation(c config, op Op, opts ...groupOption) *GroupMutation {
	m := &GroupMutation{
		config:        c,
		op:            op,
		typ:           TypeGroup,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGroupID sets the ID field of the mutation.
func withGroupID(id string) groupOption {
	return func(m *GroupMutation) {
		var (
			err   error
			once  sync.Once
			value *Group
		)
		m.oldValue = func(ctx context.Context) (*Group, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Group.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGroup sets the old Group of the mutation.
func withGroup(node *Group) groupOption {
	return func(m *GroupMutation) {
		m.oldValue = func(context.Context) (*Group, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GroupMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GroupMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Group entities.
func (m *GroupMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *GroupMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *GroupMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Group.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *GroupMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *GroupMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *GroupMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetName sets the "name" field.
func (m *GroupMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *GroupMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *GroupMutation) ResetName() {
	m.name = nil
}

// SetCategory sets the "category" field.
func (m *GroupMutation) SetCategory(s string) {
	m.category = &s
}

// Category returns the value of the "category" field in the mutation.
func (m *GroupMutation) Category() (r string, exists bool) {
	v := m.category
	if v == nil {
		return
	}
	return *v, true
}

// OldCategory returns the old "category" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldCategory(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCategory is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCategory requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCategory: %w", err)
	}
	return oldValue.Category, nil
}

// ResetCategory resets all changes to the "category" field.
func (m *GroupMutation) ResetCategory() {
	m.category = nil
}

// SetOwnerID sets the "owner_id" field.
func (m *GroupMutation) SetOwnerID(s string) {
	m.owner = &s
}

// OwnerID returns the value of the "owner_id" field in the mutation.
func (m *GroupMutation) OwnerID() (r string, exists bool) {
	v := m.owner
	if v == nil {
		return
	}
	return *v, true
}

// OldOwnerID returns the old "owner_id" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldOwnerID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOwnerID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOwnerID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwnerID: %w", err)
	}
	return oldValue.OwnerID, nil
}

// ResetOwnerID resets all changes to the "owner_id" field.
func (m *GroupMutation) ResetOwnerID() {
	m.owner = nil
}

// SetMembersMax sets the "members_max" field.
func (m *GroupMutation) SetMembersMax(i int) {
	m.members_max = &i
	m.addmembers_max = nil
}

// MembersMax returns the value of the "members_max" field in the mutation.
func (m *GroupMutation) MembersMax() (r int, exists bool) {
	v := m.members_max
	if v == nil {
		return
	}
	return *v, true
}

// OldMembersMax returns the old "members_max" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldMembersMax(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMembersMax is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMembersMax requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMembersMax: %w", err)
	}
	return oldValue.MembersMax, nil
}

// AddMembersMax adds i to the "members_max" field.
func (m *GroupMutation) AddMembersMax(i int) {
	if m.addmembers_max != nil {
		*m.addmembers_max += i
	} else {
		m.addmembers_max = &i
	}
}

// AddedMembersMax returns the value that was added to the "members_max" field in this mutation.
func (m *GroupMutation) AddedMembersMax() (r int, exists bool) {
	v := m.addmembers_max
	if v == nil {
		return
	}
	return *v, true
}

// ResetMembersMax resets all changes to the "members_max" field.
func (m *GroupMutation) ResetMembersMax() {
	m.members_max = nil
	m.addmembers_max = nil
}

// SetMembersCount sets the "members_count" field.
func (m *GroupMutation) SetMembersCount(i int) {
	m.members_count = &i
	m.addmembers_count = nil
}

// MembersCount returns the value of the "members_count" field in the mutation.
func (m *GroupMutation) MembersCount() (r int, exists bool) {
	v := m.members_count
	if v == nil {
		return
	}
	return *v, true
}

// OldMembersCount returns the old "members_count" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldMembersCount(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMembersCount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMembersCount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMembersCount: %w", err)
	}
	return oldValue.MembersCount, nil
}

// AddMembersCount adds i to the "members_count" field.
func (m *GroupMutation) AddMembersCount(i int) {
	if m.addmembers_count != nil {
		*m.addmembers_count += i
	} else {
		m.addmembers_count = &i
	}
}

// AddedMembersCount returns the value that was added to the "members_count" field in this mutation.
func (m *GroupMutation) AddedMembersCount() (r int, exists bool) {
	v := m.addmembers_count
	if v == nil {
		return
	}
	return *v, true
}

// ResetMembersCount resets all changes to the "members_count" field.
func (m *GroupMutation) ResetMembersCount() {
	m.members_count = nil
	m.addmembers_count = nil
}

// SetPublic sets the "public" field.
func (m *GroupMutation) SetPublic(b bool) {
	m.public = &b
}

// Public returns the value of the "public" field in the mutation.
func (m *GroupMutation) Public() (r bool, exists bool) {
	v := m.public
	if v == nil {
		return
	}
	return *v, true
}

// OldPublic returns the old "public" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldPublic(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPublic is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPublic requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPublic: %w", err)
	}
	return oldValue.Public, nil
}

// ResetPublic resets all changes to the "public" field.
func (m *GroupMutation) ResetPublic() {
	m.public = nil
}

// SetAddress sets the "address" field.
func (m *GroupMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the value of the "address" field in the mutation.
func (m *GroupMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// OldAddress returns the old "address" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldAddress(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAddress is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAddress requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddress: %w", err)
	}
	return oldValue.Address, nil
}

// ResetAddress resets all changes to the "address" field.
func (m *GroupMutation) ResetAddress() {
	m.address = nil
}

// SetIntro sets the "intro" field.
func (m *GroupMutation) SetIntro(s string) {
	m.intro = &s
}

// Intro returns the value of the "intro" field in the mutation.
func (m *GroupMutation) Intro() (r string, exists bool) {
	v := m.intro
	if v == nil {
		return
	}
	return *v, true
}

// OldIntro returns the old "intro" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldIntro(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIntro is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIntro requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIntro: %w", err)
	}
	return oldValue.Intro, nil
}

// ClearIntro clears the value of the "intro" field.
func (m *GroupMutation) ClearIntro() {
	m.intro = nil
	m.clearedFields[group.FieldIntro] = struct{}{}
}

// IntroCleared returns if the "intro" field was cleared in this mutation.
func (m *GroupMutation) IntroCleared() bool {
	_, ok := m.clearedFields[group.FieldIntro]
	return ok
}

// ResetIntro resets all changes to the "intro" field.
func (m *GroupMutation) ResetIntro() {
	m.intro = nil
	delete(m.clearedFields, group.FieldIntro)
}

// SetKeys sets the "keys" field.
func (m *GroupMutation) SetKeys(s string) {
	m.keys = &s
}

// Keys returns the value of the "keys" field in the mutation.
func (m *GroupMutation) Keys() (r string, exists bool) {
	v := m.keys
	if v == nil {
		return
	}
	return *v, true
}

// OldKeys returns the old "keys" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldKeys(ctx context.Context) (v string, err error) {
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
func (m *GroupMutation) ResetKeys() {
	m.keys = nil
}

// SetLastNode sets the "last_node" field.
func (m *GroupMutation) SetLastNode(i int64) {
	m.last_node = &i
	m.addlast_node = nil
}

// LastNode returns the value of the "last_node" field in the mutation.
func (m *GroupMutation) LastNode() (r int64, exists bool) {
	v := m.last_node
	if v == nil {
		return
	}
	return *v, true
}

// OldLastNode returns the old "last_node" field's value of the Group entity.
// If the Group object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GroupMutation) OldLastNode(ctx context.Context) (v int64, err error) {
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
func (m *GroupMutation) AddLastNode(i int64) {
	if m.addlast_node != nil {
		*m.addlast_node += i
	} else {
		m.addlast_node = &i
	}
}

// AddedLastNode returns the value that was added to the "last_node" field in this mutation.
func (m *GroupMutation) AddedLastNode() (r int64, exists bool) {
	v := m.addlast_node
	if v == nil {
		return
	}
	return *v, true
}

// ResetLastNode resets all changes to the "last_node" field.
func (m *GroupMutation) ResetLastNode() {
	m.last_node = nil
	m.addlast_node = nil
}

// ClearOwner clears the "owner" edge to the Member entity.
func (m *GroupMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the Member entity was cleared.
func (m *GroupMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *GroupMutation) OwnerIDs() (ids []string) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *GroupMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// AddMessageIDs adds the "messages" edge to the Message entity by ids.
func (m *GroupMutation) AddMessageIDs(ids ...string) {
	if m.messages == nil {
		m.messages = make(map[string]struct{})
	}
	for i := range ids {
		m.messages[ids[i]] = struct{}{}
	}
}

// ClearMessages clears the "messages" edge to the Message entity.
func (m *GroupMutation) ClearMessages() {
	m.clearedmessages = true
}

// MessagesCleared reports if the "messages" edge to the Message entity was cleared.
func (m *GroupMutation) MessagesCleared() bool {
	return m.clearedmessages
}

// RemoveMessageIDs removes the "messages" edge to the Message entity by IDs.
func (m *GroupMutation) RemoveMessageIDs(ids ...string) {
	if m.removedmessages == nil {
		m.removedmessages = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.messages, ids[i])
		m.removedmessages[ids[i]] = struct{}{}
	}
}

// RemovedMessages returns the removed IDs of the "messages" edge to the Message entity.
func (m *GroupMutation) RemovedMessagesIDs() (ids []string) {
	for id := range m.removedmessages {
		ids = append(ids, id)
	}
	return
}

// MessagesIDs returns the "messages" edge IDs in the mutation.
func (m *GroupMutation) MessagesIDs() (ids []string) {
	for id := range m.messages {
		ids = append(ids, id)
	}
	return
}

// ResetMessages resets all changes to the "messages" edge.
func (m *GroupMutation) ResetMessages() {
	m.messages = nil
	m.clearedmessages = false
	m.removedmessages = nil
}

// AddMemberIDs adds the "members" edge to the Member entity by ids.
func (m *GroupMutation) AddMemberIDs(ids ...string) {
	if m.members == nil {
		m.members = make(map[string]struct{})
	}
	for i := range ids {
		m.members[ids[i]] = struct{}{}
	}
}

// ClearMembers clears the "members" edge to the Member entity.
func (m *GroupMutation) ClearMembers() {
	m.clearedmembers = true
}

// MembersCleared reports if the "members" edge to the Member entity was cleared.
func (m *GroupMutation) MembersCleared() bool {
	return m.clearedmembers
}

// RemoveMemberIDs removes the "members" edge to the Member entity by IDs.
func (m *GroupMutation) RemoveMemberIDs(ids ...string) {
	if m.removedmembers == nil {
		m.removedmembers = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.members, ids[i])
		m.removedmembers[ids[i]] = struct{}{}
	}
}

// RemovedMembers returns the removed IDs of the "members" edge to the Member entity.
func (m *GroupMutation) RemovedMembersIDs() (ids []string) {
	for id := range m.removedmembers {
		ids = append(ids, id)
	}
	return
}

// MembersIDs returns the "members" edge IDs in the mutation.
func (m *GroupMutation) MembersIDs() (ids []string) {
	for id := range m.members {
		ids = append(ids, id)
	}
	return
}

// ResetMembers resets all changes to the "members" edge.
func (m *GroupMutation) ResetMembers() {
	m.members = nil
	m.clearedmembers = false
	m.removedmembers = nil
}

// AddGroupMemberIDs adds the "group_members" edge to the GroupMember entity by ids.
func (m *GroupMutation) AddGroupMemberIDs(ids ...string) {
	if m.group_members == nil {
		m.group_members = make(map[string]struct{})
	}
	for i := range ids {
		m.group_members[ids[i]] = struct{}{}
	}
}

// ClearGroupMembers clears the "group_members" edge to the GroupMember entity.
func (m *GroupMutation) ClearGroupMembers() {
	m.clearedgroup_members = true
}

// GroupMembersCleared reports if the "group_members" edge to the GroupMember entity was cleared.
func (m *GroupMutation) GroupMembersCleared() bool {
	return m.clearedgroup_members
}

// RemoveGroupMemberIDs removes the "group_members" edge to the GroupMember entity by IDs.
func (m *GroupMutation) RemoveGroupMemberIDs(ids ...string) {
	if m.removedgroup_members == nil {
		m.removedgroup_members = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.group_members, ids[i])
		m.removedgroup_members[ids[i]] = struct{}{}
	}
}

// RemovedGroupMembers returns the removed IDs of the "group_members" edge to the GroupMember entity.
func (m *GroupMutation) RemovedGroupMembersIDs() (ids []string) {
	for id := range m.removedgroup_members {
		ids = append(ids, id)
	}
	return
}

// GroupMembersIDs returns the "group_members" edge IDs in the mutation.
func (m *GroupMutation) GroupMembersIDs() (ids []string) {
	for id := range m.group_members {
		ids = append(ids, id)
	}
	return
}

// ResetGroupMembers resets all changes to the "group_members" edge.
func (m *GroupMutation) ResetGroupMembers() {
	m.group_members = nil
	m.clearedgroup_members = false
	m.removedgroup_members = nil
}

// Where appends a list predicates to the GroupMutation builder.
func (m *GroupMutation) Where(ps ...predicate.Group) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *GroupMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Group).
func (m *GroupMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *GroupMutation) Fields() []string {
	fields := make([]string, 0, 11)
	if m.created_at != nil {
		fields = append(fields, group.FieldCreatedAt)
	}
	if m.name != nil {
		fields = append(fields, group.FieldName)
	}
	if m.category != nil {
		fields = append(fields, group.FieldCategory)
	}
	if m.owner != nil {
		fields = append(fields, group.FieldOwnerID)
	}
	if m.members_max != nil {
		fields = append(fields, group.FieldMembersMax)
	}
	if m.members_count != nil {
		fields = append(fields, group.FieldMembersCount)
	}
	if m.public != nil {
		fields = append(fields, group.FieldPublic)
	}
	if m.address != nil {
		fields = append(fields, group.FieldAddress)
	}
	if m.intro != nil {
		fields = append(fields, group.FieldIntro)
	}
	if m.keys != nil {
		fields = append(fields, group.FieldKeys)
	}
	if m.last_node != nil {
		fields = append(fields, group.FieldLastNode)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *GroupMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case group.FieldCreatedAt:
		return m.CreatedAt()
	case group.FieldName:
		return m.Name()
	case group.FieldCategory:
		return m.Category()
	case group.FieldOwnerID:
		return m.OwnerID()
	case group.FieldMembersMax:
		return m.MembersMax()
	case group.FieldMembersCount:
		return m.MembersCount()
	case group.FieldPublic:
		return m.Public()
	case group.FieldAddress:
		return m.Address()
	case group.FieldIntro:
		return m.Intro()
	case group.FieldKeys:
		return m.Keys()
	case group.FieldLastNode:
		return m.LastNode()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *GroupMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case group.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case group.FieldName:
		return m.OldName(ctx)
	case group.FieldCategory:
		return m.OldCategory(ctx)
	case group.FieldOwnerID:
		return m.OldOwnerID(ctx)
	case group.FieldMembersMax:
		return m.OldMembersMax(ctx)
	case group.FieldMembersCount:
		return m.OldMembersCount(ctx)
	case group.FieldPublic:
		return m.OldPublic(ctx)
	case group.FieldAddress:
		return m.OldAddress(ctx)
	case group.FieldIntro:
		return m.OldIntro(ctx)
	case group.FieldKeys:
		return m.OldKeys(ctx)
	case group.FieldLastNode:
		return m.OldLastNode(ctx)
	}
	return nil, fmt.Errorf("unknown Group field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GroupMutation) SetField(name string, value ent.Value) error {
	switch name {
	case group.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case group.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case group.FieldCategory:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCategory(v)
		return nil
	case group.FieldOwnerID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwnerID(v)
		return nil
	case group.FieldMembersMax:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMembersMax(v)
		return nil
	case group.FieldMembersCount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMembersCount(v)
		return nil
	case group.FieldPublic:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPublic(v)
		return nil
	case group.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	case group.FieldIntro:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIntro(v)
		return nil
	case group.FieldKeys:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetKeys(v)
		return nil
	case group.FieldLastNode:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastNode(v)
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *GroupMutation) AddedFields() []string {
	var fields []string
	if m.addmembers_max != nil {
		fields = append(fields, group.FieldMembersMax)
	}
	if m.addmembers_count != nil {
		fields = append(fields, group.FieldMembersCount)
	}
	if m.addlast_node != nil {
		fields = append(fields, group.FieldLastNode)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *GroupMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case group.FieldMembersMax:
		return m.AddedMembersMax()
	case group.FieldMembersCount:
		return m.AddedMembersCount()
	case group.FieldLastNode:
		return m.AddedLastNode()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GroupMutation) AddField(name string, value ent.Value) error {
	switch name {
	case group.FieldMembersMax:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddMembersMax(v)
		return nil
	case group.FieldMembersCount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddMembersCount(v)
		return nil
	case group.FieldLastNode:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLastNode(v)
		return nil
	}
	return fmt.Errorf("unknown Group numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *GroupMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(group.FieldIntro) {
		fields = append(fields, group.FieldIntro)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *GroupMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *GroupMutation) ClearField(name string) error {
	switch name {
	case group.FieldIntro:
		m.ClearIntro()
		return nil
	}
	return fmt.Errorf("unknown Group nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *GroupMutation) ResetField(name string) error {
	switch name {
	case group.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case group.FieldName:
		m.ResetName()
		return nil
	case group.FieldCategory:
		m.ResetCategory()
		return nil
	case group.FieldOwnerID:
		m.ResetOwnerID()
		return nil
	case group.FieldMembersMax:
		m.ResetMembersMax()
		return nil
	case group.FieldMembersCount:
		m.ResetMembersCount()
		return nil
	case group.FieldPublic:
		m.ResetPublic()
		return nil
	case group.FieldAddress:
		m.ResetAddress()
		return nil
	case group.FieldIntro:
		m.ResetIntro()
		return nil
	case group.FieldKeys:
		m.ResetKeys()
		return nil
	case group.FieldLastNode:
		m.ResetLastNode()
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *GroupMutation) AddedEdges() []string {
	edges := make([]string, 0, 4)
	if m.owner != nil {
		edges = append(edges, group.EdgeOwner)
	}
	if m.messages != nil {
		edges = append(edges, group.EdgeMessages)
	}
	if m.members != nil {
		edges = append(edges, group.EdgeMembers)
	}
	if m.group_members != nil {
		edges = append(edges, group.EdgeGroupMembers)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *GroupMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case group.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	case group.EdgeMessages:
		ids := make([]ent.Value, 0, len(m.messages))
		for id := range m.messages {
			ids = append(ids, id)
		}
		return ids
	case group.EdgeMembers:
		ids := make([]ent.Value, 0, len(m.members))
		for id := range m.members {
			ids = append(ids, id)
		}
		return ids
	case group.EdgeGroupMembers:
		ids := make([]ent.Value, 0, len(m.group_members))
		for id := range m.group_members {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *GroupMutation) RemovedEdges() []string {
	edges := make([]string, 0, 4)
	if m.removedmessages != nil {
		edges = append(edges, group.EdgeMessages)
	}
	if m.removedmembers != nil {
		edges = append(edges, group.EdgeMembers)
	}
	if m.removedgroup_members != nil {
		edges = append(edges, group.EdgeGroupMembers)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *GroupMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case group.EdgeMessages:
		ids := make([]ent.Value, 0, len(m.removedmessages))
		for id := range m.removedmessages {
			ids = append(ids, id)
		}
		return ids
	case group.EdgeMembers:
		ids := make([]ent.Value, 0, len(m.removedmembers))
		for id := range m.removedmembers {
			ids = append(ids, id)
		}
		return ids
	case group.EdgeGroupMembers:
		ids := make([]ent.Value, 0, len(m.removedgroup_members))
		for id := range m.removedgroup_members {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *GroupMutation) ClearedEdges() []string {
	edges := make([]string, 0, 4)
	if m.clearedowner {
		edges = append(edges, group.EdgeOwner)
	}
	if m.clearedmessages {
		edges = append(edges, group.EdgeMessages)
	}
	if m.clearedmembers {
		edges = append(edges, group.EdgeMembers)
	}
	if m.clearedgroup_members {
		edges = append(edges, group.EdgeGroupMembers)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *GroupMutation) EdgeCleared(name string) bool {
	switch name {
	case group.EdgeOwner:
		return m.clearedowner
	case group.EdgeMessages:
		return m.clearedmessages
	case group.EdgeMembers:
		return m.clearedmembers
	case group.EdgeGroupMembers:
		return m.clearedgroup_members
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *GroupMutation) ClearEdge(name string) error {
	switch name {
	case group.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Group unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *GroupMutation) ResetEdge(name string) error {
	switch name {
	case group.EdgeOwner:
		m.ResetOwner()
		return nil
	case group.EdgeMessages:
		m.ResetMessages()
		return nil
	case group.EdgeMembers:
		m.ResetMembers()
		return nil
	case group.EdgeGroupMembers:
		m.ResetGroupMembers()
		return nil
	}
	return fmt.Errorf("unknown Group edge %s", name)
}

