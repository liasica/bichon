// Code generated, DO NOT EDIT.

package ent

import (
    "context"
    "fmt"
    "sync"
    "errors"
    "time"
    "github.com/chatpuppy/puppychat/internal/ent/member"
    "github.com/chatpuppy/puppychat/internal/ent/predicate"

    "entgo.io/ent"
)


// MemberMutation represents an operation that mutates the Member nodes in the graph.
type MemberMutation struct {
	config
	op                   Op
	typ                  string
	id                   *string
	created_at           *time.Time
	updated_at           *time.Time
	address              *string
	nickname             *string
	avatar               *string
	intro                *string
	public_key           *string
	nonce                *string
	show_nickname        *bool
	clearedFields        map[string]struct{}
	own_groups           map[string]struct{}
	removedown_groups    map[string]struct{}
	clearedown_groups    bool
	messages             map[string]struct{}
	removedmessages      map[string]struct{}
	clearedmessages      bool
	groups               map[string]struct{}
	removedgroups        map[string]struct{}
	clearedgroups        bool
	group_members        map[string]struct{}
	removedgroup_members map[string]struct{}
	clearedgroup_members bool
	done                 bool
	oldValue             func(context.Context) (*Member, error)
	predicates           []predicate.Member
}

var _ ent.Mutation = (*MemberMutation)(nil)

// memberOption allows management of the mutation configuration using functional options.
type memberOption func(*MemberMutation)

// newMemberMutation creates new mutation for the Member entity.
func newMemberMutation(c config, op Op, opts ...memberOption) *MemberMutation {
	m := &MemberMutation{
		config:        c,
		op:            op,
		typ:           TypeMember,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMemberID sets the ID field of the mutation.
func withMemberID(id string) memberOption {
	return func(m *MemberMutation) {
		var (
			err   error
			once  sync.Once
			value *Member
		)
		m.oldValue = func(ctx context.Context) (*Member, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Member.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMember sets the old Member of the mutation.
func withMember(node *Member) memberOption {
	return func(m *MemberMutation) {
		m.oldValue = func(context.Context) (*Member, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MemberMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MemberMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Member entities.
func (m *MemberMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MemberMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MemberMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Member.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *MemberMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *MemberMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *MemberMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *MemberMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *MemberMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *MemberMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetAddress sets the "address" field.
func (m *MemberMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the value of the "address" field in the mutation.
func (m *MemberMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// OldAddress returns the old "address" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldAddress(ctx context.Context) (v string, err error) {
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
func (m *MemberMutation) ResetAddress() {
	m.address = nil
}

// SetNickname sets the "nickname" field.
func (m *MemberMutation) SetNickname(s string) {
	m.nickname = &s
}

// Nickname returns the value of the "nickname" field in the mutation.
func (m *MemberMutation) Nickname() (r string, exists bool) {
	v := m.nickname
	if v == nil {
		return
	}
	return *v, true
}

// OldNickname returns the old "nickname" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldNickname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNickname is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNickname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickname: %w", err)
	}
	return oldValue.Nickname, nil
}

// ClearNickname clears the value of the "nickname" field.
func (m *MemberMutation) ClearNickname() {
	m.nickname = nil
	m.clearedFields[member.FieldNickname] = struct{}{}
}

// NicknameCleared returns if the "nickname" field was cleared in this mutation.
func (m *MemberMutation) NicknameCleared() bool {
	_, ok := m.clearedFields[member.FieldNickname]
	return ok
}

// ResetNickname resets all changes to the "nickname" field.
func (m *MemberMutation) ResetNickname() {
	m.nickname = nil
	delete(m.clearedFields, member.FieldNickname)
}

// SetAvatar sets the "avatar" field.
func (m *MemberMutation) SetAvatar(s string) {
	m.avatar = &s
}

// Avatar returns the value of the "avatar" field in the mutation.
func (m *MemberMutation) Avatar() (r string, exists bool) {
	v := m.avatar
	if v == nil {
		return
	}
	return *v, true
}

// OldAvatar returns the old "avatar" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldAvatar(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAvatar is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAvatar requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAvatar: %w", err)
	}
	return oldValue.Avatar, nil
}

// ClearAvatar clears the value of the "avatar" field.
func (m *MemberMutation) ClearAvatar() {
	m.avatar = nil
	m.clearedFields[member.FieldAvatar] = struct{}{}
}

// AvatarCleared returns if the "avatar" field was cleared in this mutation.
func (m *MemberMutation) AvatarCleared() bool {
	_, ok := m.clearedFields[member.FieldAvatar]
	return ok
}

// ResetAvatar resets all changes to the "avatar" field.
func (m *MemberMutation) ResetAvatar() {
	m.avatar = nil
	delete(m.clearedFields, member.FieldAvatar)
}

// SetIntro sets the "intro" field.
func (m *MemberMutation) SetIntro(s string) {
	m.intro = &s
}

// Intro returns the value of the "intro" field in the mutation.
func (m *MemberMutation) Intro() (r string, exists bool) {
	v := m.intro
	if v == nil {
		return
	}
	return *v, true
}

// OldIntro returns the old "intro" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldIntro(ctx context.Context) (v string, err error) {
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
func (m *MemberMutation) ClearIntro() {
	m.intro = nil
	m.clearedFields[member.FieldIntro] = struct{}{}
}

// IntroCleared returns if the "intro" field was cleared in this mutation.
func (m *MemberMutation) IntroCleared() bool {
	_, ok := m.clearedFields[member.FieldIntro]
	return ok
}

// ResetIntro resets all changes to the "intro" field.
func (m *MemberMutation) ResetIntro() {
	m.intro = nil
	delete(m.clearedFields, member.FieldIntro)
}

// SetPublicKey sets the "public_key" field.
func (m *MemberMutation) SetPublicKey(s string) {
	m.public_key = &s
}

// PublicKey returns the value of the "public_key" field in the mutation.
func (m *MemberMutation) PublicKey() (r string, exists bool) {
	v := m.public_key
	if v == nil {
		return
	}
	return *v, true
}

// OldPublicKey returns the old "public_key" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldPublicKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPublicKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPublicKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPublicKey: %w", err)
	}
	return oldValue.PublicKey, nil
}

// ClearPublicKey clears the value of the "public_key" field.
func (m *MemberMutation) ClearPublicKey() {
	m.public_key = nil
	m.clearedFields[member.FieldPublicKey] = struct{}{}
}

// PublicKeyCleared returns if the "public_key" field was cleared in this mutation.
func (m *MemberMutation) PublicKeyCleared() bool {
	_, ok := m.clearedFields[member.FieldPublicKey]
	return ok
}

// ResetPublicKey resets all changes to the "public_key" field.
func (m *MemberMutation) ResetPublicKey() {
	m.public_key = nil
	delete(m.clearedFields, member.FieldPublicKey)
}

// SetNonce sets the "nonce" field.
func (m *MemberMutation) SetNonce(s string) {
	m.nonce = &s
}

// Nonce returns the value of the "nonce" field in the mutation.
func (m *MemberMutation) Nonce() (r string, exists bool) {
	v := m.nonce
	if v == nil {
		return
	}
	return *v, true
}

// OldNonce returns the old "nonce" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldNonce(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNonce is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNonce requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNonce: %w", err)
	}
	return oldValue.Nonce, nil
}

// ResetNonce resets all changes to the "nonce" field.
func (m *MemberMutation) ResetNonce() {
	m.nonce = nil
}

// SetShowNickname sets the "show_nickname" field.
func (m *MemberMutation) SetShowNickname(b bool) {
	m.show_nickname = &b
}

// ShowNickname returns the value of the "show_nickname" field in the mutation.
func (m *MemberMutation) ShowNickname() (r bool, exists bool) {
	v := m.show_nickname
	if v == nil {
		return
	}
	return *v, true
}

// OldShowNickname returns the old "show_nickname" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldShowNickname(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldShowNickname is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldShowNickname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldShowNickname: %w", err)
	}
	return oldValue.ShowNickname, nil
}

// ResetShowNickname resets all changes to the "show_nickname" field.
func (m *MemberMutation) ResetShowNickname() {
	m.show_nickname = nil
}

// AddOwnGroupIDs adds the "own_groups" edge to the Group entity by ids.
func (m *MemberMutation) AddOwnGroupIDs(ids ...string) {
	if m.own_groups == nil {
		m.own_groups = make(map[string]struct{})
	}
	for i := range ids {
		m.own_groups[ids[i]] = struct{}{}
	}
}

// ClearOwnGroups clears the "own_groups" edge to the Group entity.
func (m *MemberMutation) ClearOwnGroups() {
	m.clearedown_groups = true
}

// OwnGroupsCleared reports if the "own_groups" edge to the Group entity was cleared.
func (m *MemberMutation) OwnGroupsCleared() bool {
	return m.clearedown_groups
}

// RemoveOwnGroupIDs removes the "own_groups" edge to the Group entity by IDs.
func (m *MemberMutation) RemoveOwnGroupIDs(ids ...string) {
	if m.removedown_groups == nil {
		m.removedown_groups = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.own_groups, ids[i])
		m.removedown_groups[ids[i]] = struct{}{}
	}
}

// RemovedOwnGroups returns the removed IDs of the "own_groups" edge to the Group entity.
func (m *MemberMutation) RemovedOwnGroupsIDs() (ids []string) {
	for id := range m.removedown_groups {
		ids = append(ids, id)
	}
	return
}

// OwnGroupsIDs returns the "own_groups" edge IDs in the mutation.
func (m *MemberMutation) OwnGroupsIDs() (ids []string) {
	for id := range m.own_groups {
		ids = append(ids, id)
	}
	return
}

// ResetOwnGroups resets all changes to the "own_groups" edge.
func (m *MemberMutation) ResetOwnGroups() {
	m.own_groups = nil
	m.clearedown_groups = false
	m.removedown_groups = nil
}

// AddMessageIDs adds the "messages" edge to the Message entity by ids.
func (m *MemberMutation) AddMessageIDs(ids ...string) {
	if m.messages == nil {
		m.messages = make(map[string]struct{})
	}
	for i := range ids {
		m.messages[ids[i]] = struct{}{}
	}
}

// ClearMessages clears the "messages" edge to the Message entity.
func (m *MemberMutation) ClearMessages() {
	m.clearedmessages = true
}

// MessagesCleared reports if the "messages" edge to the Message entity was cleared.
func (m *MemberMutation) MessagesCleared() bool {
	return m.clearedmessages
}

// RemoveMessageIDs removes the "messages" edge to the Message entity by IDs.
func (m *MemberMutation) RemoveMessageIDs(ids ...string) {
	if m.removedmessages == nil {
		m.removedmessages = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.messages, ids[i])
		m.removedmessages[ids[i]] = struct{}{}
	}
}

// RemovedMessages returns the removed IDs of the "messages" edge to the Message entity.
func (m *MemberMutation) RemovedMessagesIDs() (ids []string) {
	for id := range m.removedmessages {
		ids = append(ids, id)
	}
	return
}

// MessagesIDs returns the "messages" edge IDs in the mutation.
func (m *MemberMutation) MessagesIDs() (ids []string) {
	for id := range m.messages {
		ids = append(ids, id)
	}
	return
}

// ResetMessages resets all changes to the "messages" edge.
func (m *MemberMutation) ResetMessages() {
	m.messages = nil
	m.clearedmessages = false
	m.removedmessages = nil
}

// AddGroupIDs adds the "groups" edge to the Group entity by ids.
func (m *MemberMutation) AddGroupIDs(ids ...string) {
	if m.groups == nil {
		m.groups = make(map[string]struct{})
	}
	for i := range ids {
		m.groups[ids[i]] = struct{}{}
	}
}

// ClearGroups clears the "groups" edge to the Group entity.
func (m *MemberMutation) ClearGroups() {
	m.clearedgroups = true
}

// GroupsCleared reports if the "groups" edge to the Group entity was cleared.
func (m *MemberMutation) GroupsCleared() bool {
	return m.clearedgroups
}

// RemoveGroupIDs removes the "groups" edge to the Group entity by IDs.
func (m *MemberMutation) RemoveGroupIDs(ids ...string) {
	if m.removedgroups == nil {
		m.removedgroups = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.groups, ids[i])
		m.removedgroups[ids[i]] = struct{}{}
	}
}

// RemovedGroups returns the removed IDs of the "groups" edge to the Group entity.
func (m *MemberMutation) RemovedGroupsIDs() (ids []string) {
	for id := range m.removedgroups {
		ids = append(ids, id)
	}
	return
}

// GroupsIDs returns the "groups" edge IDs in the mutation.
func (m *MemberMutation) GroupsIDs() (ids []string) {
	for id := range m.groups {
		ids = append(ids, id)
	}
	return
}

// ResetGroups resets all changes to the "groups" edge.
func (m *MemberMutation) ResetGroups() {
	m.groups = nil
	m.clearedgroups = false
	m.removedgroups = nil
}

// AddGroupMemberIDs adds the "group_members" edge to the GroupMember entity by ids.
func (m *MemberMutation) AddGroupMemberIDs(ids ...string) {
	if m.group_members == nil {
		m.group_members = make(map[string]struct{})
	}
	for i := range ids {
		m.group_members[ids[i]] = struct{}{}
	}
}

// ClearGroupMembers clears the "group_members" edge to the GroupMember entity.
func (m *MemberMutation) ClearGroupMembers() {
	m.clearedgroup_members = true
}

// GroupMembersCleared reports if the "group_members" edge to the GroupMember entity was cleared.
func (m *MemberMutation) GroupMembersCleared() bool {
	return m.clearedgroup_members
}

// RemoveGroupMemberIDs removes the "group_members" edge to the GroupMember entity by IDs.
func (m *MemberMutation) RemoveGroupMemberIDs(ids ...string) {
	if m.removedgroup_members == nil {
		m.removedgroup_members = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.group_members, ids[i])
		m.removedgroup_members[ids[i]] = struct{}{}
	}
}

// RemovedGroupMembers returns the removed IDs of the "group_members" edge to the GroupMember entity.
func (m *MemberMutation) RemovedGroupMembersIDs() (ids []string) {
	for id := range m.removedgroup_members {
		ids = append(ids, id)
	}
	return
}

// GroupMembersIDs returns the "group_members" edge IDs in the mutation.
func (m *MemberMutation) GroupMembersIDs() (ids []string) {
	for id := range m.group_members {
		ids = append(ids, id)
	}
	return
}

// ResetGroupMembers resets all changes to the "group_members" edge.
func (m *MemberMutation) ResetGroupMembers() {
	m.group_members = nil
	m.clearedgroup_members = false
	m.removedgroup_members = nil
}

// Where appends a list predicates to the MemberMutation builder.
func (m *MemberMutation) Where(ps ...predicate.Member) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *MemberMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Member).
func (m *MemberMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MemberMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.created_at != nil {
		fields = append(fields, member.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, member.FieldUpdatedAt)
	}
	if m.address != nil {
		fields = append(fields, member.FieldAddress)
	}
	if m.nickname != nil {
		fields = append(fields, member.FieldNickname)
	}
	if m.avatar != nil {
		fields = append(fields, member.FieldAvatar)
	}
	if m.intro != nil {
		fields = append(fields, member.FieldIntro)
	}
	if m.public_key != nil {
		fields = append(fields, member.FieldPublicKey)
	}
	if m.nonce != nil {
		fields = append(fields, member.FieldNonce)
	}
	if m.show_nickname != nil {
		fields = append(fields, member.FieldShowNickname)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MemberMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case member.FieldCreatedAt:
		return m.CreatedAt()
	case member.FieldUpdatedAt:
		return m.UpdatedAt()
	case member.FieldAddress:
		return m.Address()
	case member.FieldNickname:
		return m.Nickname()
	case member.FieldAvatar:
		return m.Avatar()
	case member.FieldIntro:
		return m.Intro()
	case member.FieldPublicKey:
		return m.PublicKey()
	case member.FieldNonce:
		return m.Nonce()
	case member.FieldShowNickname:
		return m.ShowNickname()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MemberMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case member.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case member.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case member.FieldAddress:
		return m.OldAddress(ctx)
	case member.FieldNickname:
		return m.OldNickname(ctx)
	case member.FieldAvatar:
		return m.OldAvatar(ctx)
	case member.FieldIntro:
		return m.OldIntro(ctx)
	case member.FieldPublicKey:
		return m.OldPublicKey(ctx)
	case member.FieldNonce:
		return m.OldNonce(ctx)
	case member.FieldShowNickname:
		return m.OldShowNickname(ctx)
	}
	return nil, fmt.Errorf("unknown Member field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MemberMutation) SetField(name string, value ent.Value) error {
	switch name {
	case member.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case member.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case member.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	case member.FieldNickname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickname(v)
		return nil
	case member.FieldAvatar:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAvatar(v)
		return nil
	case member.FieldIntro:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIntro(v)
		return nil
	case member.FieldPublicKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPublicKey(v)
		return nil
	case member.FieldNonce:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNonce(v)
		return nil
	case member.FieldShowNickname:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetShowNickname(v)
		return nil
	}
	return fmt.Errorf("unknown Member field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MemberMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MemberMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MemberMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Member numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MemberMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(member.FieldNickname) {
		fields = append(fields, member.FieldNickname)
	}
	if m.FieldCleared(member.FieldAvatar) {
		fields = append(fields, member.FieldAvatar)
	}
	if m.FieldCleared(member.FieldIntro) {
		fields = append(fields, member.FieldIntro)
	}
	if m.FieldCleared(member.FieldPublicKey) {
		fields = append(fields, member.FieldPublicKey)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MemberMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MemberMutation) ClearField(name string) error {
	switch name {
	case member.FieldNickname:
		m.ClearNickname()
		return nil
	case member.FieldAvatar:
		m.ClearAvatar()
		return nil
	case member.FieldIntro:
		m.ClearIntro()
		return nil
	case member.FieldPublicKey:
		m.ClearPublicKey()
		return nil
	}
	return fmt.Errorf("unknown Member nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MemberMutation) ResetField(name string) error {
	switch name {
	case member.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case member.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case member.FieldAddress:
		m.ResetAddress()
		return nil
	case member.FieldNickname:
		m.ResetNickname()
		return nil
	case member.FieldAvatar:
		m.ResetAvatar()
		return nil
	case member.FieldIntro:
		m.ResetIntro()
		return nil
	case member.FieldPublicKey:
		m.ResetPublicKey()
		return nil
	case member.FieldNonce:
		m.ResetNonce()
		return nil
	case member.FieldShowNickname:
		m.ResetShowNickname()
		return nil
	}
	return fmt.Errorf("unknown Member field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MemberMutation) AddedEdges() []string {
	edges := make([]string, 0, 4)
	if m.own_groups != nil {
		edges = append(edges, member.EdgeOwnGroups)
	}
	if m.messages != nil {
		edges = append(edges, member.EdgeMessages)
	}
	if m.groups != nil {
		edges = append(edges, member.EdgeGroups)
	}
	if m.group_members != nil {
		edges = append(edges, member.EdgeGroupMembers)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MemberMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case member.EdgeOwnGroups:
		ids := make([]ent.Value, 0, len(m.own_groups))
		for id := range m.own_groups {
			ids = append(ids, id)
		}
		return ids
	case member.EdgeMessages:
		ids := make([]ent.Value, 0, len(m.messages))
		for id := range m.messages {
			ids = append(ids, id)
		}
		return ids
	case member.EdgeGroups:
		ids := make([]ent.Value, 0, len(m.groups))
		for id := range m.groups {
			ids = append(ids, id)
		}
		return ids
	case member.EdgeGroupMembers:
		ids := make([]ent.Value, 0, len(m.group_members))
		for id := range m.group_members {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MemberMutation) RemovedEdges() []string {
	edges := make([]string, 0, 4)
	if m.removedown_groups != nil {
		edges = append(edges, member.EdgeOwnGroups)
	}
	if m.removedmessages != nil {
		edges = append(edges, member.EdgeMessages)
	}
	if m.removedgroups != nil {
		edges = append(edges, member.EdgeGroups)
	}
	if m.removedgroup_members != nil {
		edges = append(edges, member.EdgeGroupMembers)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MemberMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case member.EdgeOwnGroups:
		ids := make([]ent.Value, 0, len(m.removedown_groups))
		for id := range m.removedown_groups {
			ids = append(ids, id)
		}
		return ids
	case member.EdgeMessages:
		ids := make([]ent.Value, 0, len(m.removedmessages))
		for id := range m.removedmessages {
			ids = append(ids, id)
		}
		return ids
	case member.EdgeGroups:
		ids := make([]ent.Value, 0, len(m.removedgroups))
		for id := range m.removedgroups {
			ids = append(ids, id)
		}
		return ids
	case member.EdgeGroupMembers:
		ids := make([]ent.Value, 0, len(m.removedgroup_members))
		for id := range m.removedgroup_members {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MemberMutation) ClearedEdges() []string {
	edges := make([]string, 0, 4)
	if m.clearedown_groups {
		edges = append(edges, member.EdgeOwnGroups)
	}
	if m.clearedmessages {
		edges = append(edges, member.EdgeMessages)
	}
	if m.clearedgroups {
		edges = append(edges, member.EdgeGroups)
	}
	if m.clearedgroup_members {
		edges = append(edges, member.EdgeGroupMembers)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MemberMutation) EdgeCleared(name string) bool {
	switch name {
	case member.EdgeOwnGroups:
		return m.clearedown_groups
	case member.EdgeMessages:
		return m.clearedmessages
	case member.EdgeGroups:
		return m.clearedgroups
	case member.EdgeGroupMembers:
		return m.clearedgroup_members
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MemberMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Member unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MemberMutation) ResetEdge(name string) error {
	switch name {
	case member.EdgeOwnGroups:
		m.ResetOwnGroups()
		return nil
	case member.EdgeMessages:
		m.ResetMessages()
		return nil
	case member.EdgeGroups:
		m.ResetGroups()
		return nil
	case member.EdgeGroupMembers:
		m.ResetGroupMembers()
		return nil
	}
	return fmt.Errorf("unknown Member edge %s", name)
}

