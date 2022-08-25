// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/message"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (gc *GroupCreate) SetCreatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableCreatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetName sets the "name" field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetMemberID sets the "member_id" field.
func (gc *GroupCreate) SetMemberID(u uint64) *GroupCreate {
	gc.mutation.SetMemberID(u)
	return gc
}

// SetMembersMax sets the "members_max" field.
func (gc *GroupCreate) SetMembersMax(i int) *GroupCreate {
	gc.mutation.SetMembersMax(i)
	return gc
}

// SetMembersCount sets the "members_count" field.
func (gc *GroupCreate) SetMembersCount(i int) *GroupCreate {
	gc.mutation.SetMembersCount(i)
	return gc
}

// SetPublic sets the "public" field.
func (gc *GroupCreate) SetPublic(b bool) *GroupCreate {
	gc.mutation.SetPublic(b)
	return gc
}

// SetSn sets the "sn" field.
func (gc *GroupCreate) SetSn(s string) *GroupCreate {
	gc.mutation.SetSn(s)
	return gc
}

// SetIntro sets the "intro" field.
func (gc *GroupCreate) SetIntro(s string) *GroupCreate {
	gc.mutation.SetIntro(s)
	return gc
}

// SetOwnerID sets the "owner" edge to the Member entity by ID.
func (gc *GroupCreate) SetOwnerID(id uint64) *GroupCreate {
	gc.mutation.SetOwnerID(id)
	return gc
}

// SetOwner sets the "owner" edge to the Member entity.
func (gc *GroupCreate) SetOwner(m *Member) *GroupCreate {
	return gc.SetOwnerID(m.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (gc *GroupCreate) AddMessageIDs(ids ...uint64) *GroupCreate {
	gc.mutation.AddMessageIDs(ids...)
	return gc
}

// AddMessages adds the "messages" edges to the Message entity.
func (gc *GroupCreate) AddMessages(m ...*Message) *GroupCreate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gc.AddMessageIDs(ids...)
}

// AddMemberIDs adds the "members" edge to the Member entity by IDs.
func (gc *GroupCreate) AddMemberIDs(ids ...uint64) *GroupCreate {
	gc.mutation.AddMemberIDs(ids...)
	return gc
}

// AddMembers adds the "members" edges to the Member entity.
func (gc *GroupCreate) AddMembers(m ...*Member) *GroupCreate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gc.AddMemberIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	var (
		err  error
		node *Group
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			if node, err = gc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			if gc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Group)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GroupMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GroupCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GroupCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GroupCreate) defaults() {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := group.DefaultCreatedAt
		gc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Group.created_at"`)}
	}
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Group.name"`)}
	}
	if _, ok := gc.mutation.MemberID(); !ok {
		return &ValidationError{Name: "member_id", err: errors.New(`ent: missing required field "Group.member_id"`)}
	}
	if _, ok := gc.mutation.MembersMax(); !ok {
		return &ValidationError{Name: "members_max", err: errors.New(`ent: missing required field "Group.members_max"`)}
	}
	if _, ok := gc.mutation.MembersCount(); !ok {
		return &ValidationError{Name: "members_count", err: errors.New(`ent: missing required field "Group.members_count"`)}
	}
	if _, ok := gc.mutation.Public(); !ok {
		return &ValidationError{Name: "public", err: errors.New(`ent: missing required field "Group.public"`)}
	}
	if _, ok := gc.mutation.Sn(); !ok {
		return &ValidationError{Name: "sn", err: errors.New(`ent: missing required field "Group.sn"`)}
	}
	if _, ok := gc.mutation.Intro(); !ok {
		return &ValidationError{Name: "intro", err: errors.New(`ent: missing required field "Group.intro"`)}
	}
	if _, ok := gc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Group.owner"`)}
	}
	return nil
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = uint64(id)
	return _node, nil
}

func (gc *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: group.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: group.FieldID,
			},
		}
	)
	_spec.OnConflict = gc.conflict
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
		_node.Name = value
	}
	if value, ok := gc.mutation.MembersMax(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersMax,
		})
		_node.MembersMax = value
	}
	if value, ok := gc.mutation.MembersCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersCount,
		})
		_node.MembersCount = value
	}
	if value, ok := gc.mutation.Public(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: group.FieldPublic,
		})
		_node.Public = value
	}
	if value, ok := gc.mutation.Sn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldSn,
		})
		_node.Sn = value
	}
	if value, ok := gc.mutation.Intro(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldIntro,
		})
		_node.Intro = value
	}
	if nodes := gc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MemberID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: group.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Group.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (gc *GroupCreate) OnConflict(opts ...sql.ConflictOption) *GroupUpsertOne {
	gc.conflict = opts
	return &GroupUpsertOne{
		create: gc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gc *GroupCreate) OnConflictColumns(columns ...string) *GroupUpsertOne {
	gc.conflict = append(gc.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertOne{
		create: gc,
	}
}

type (
	// GroupUpsertOne is the builder for "upsert"-ing
	//  one Group node.
	GroupUpsertOne struct {
		create *GroupCreate
	}

	// GroupUpsert is the "OnConflict" setter.
	GroupUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *GroupUpsert) SetCreatedAt(v time.Time) *GroupUpsert {
	u.Set(group.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GroupUpsert) UpdateCreatedAt() *GroupUpsert {
	u.SetExcluded(group.FieldCreatedAt)
	return u
}

// SetName sets the "name" field.
func (u *GroupUpsert) SetName(v string) *GroupUpsert {
	u.Set(group.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsert) UpdateName() *GroupUpsert {
	u.SetExcluded(group.FieldName)
	return u
}

// SetMemberID sets the "member_id" field.
func (u *GroupUpsert) SetMemberID(v uint64) *GroupUpsert {
	u.Set(group.FieldMemberID, v)
	return u
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *GroupUpsert) UpdateMemberID() *GroupUpsert {
	u.SetExcluded(group.FieldMemberID)
	return u
}

// SetMembersMax sets the "members_max" field.
func (u *GroupUpsert) SetMembersMax(v int) *GroupUpsert {
	u.Set(group.FieldMembersMax, v)
	return u
}

// UpdateMembersMax sets the "members_max" field to the value that was provided on create.
func (u *GroupUpsert) UpdateMembersMax() *GroupUpsert {
	u.SetExcluded(group.FieldMembersMax)
	return u
}

// AddMembersMax adds v to the "members_max" field.
func (u *GroupUpsert) AddMembersMax(v int) *GroupUpsert {
	u.Add(group.FieldMembersMax, v)
	return u
}

// SetMembersCount sets the "members_count" field.
func (u *GroupUpsert) SetMembersCount(v int) *GroupUpsert {
	u.Set(group.FieldMembersCount, v)
	return u
}

// UpdateMembersCount sets the "members_count" field to the value that was provided on create.
func (u *GroupUpsert) UpdateMembersCount() *GroupUpsert {
	u.SetExcluded(group.FieldMembersCount)
	return u
}

// AddMembersCount adds v to the "members_count" field.
func (u *GroupUpsert) AddMembersCount(v int) *GroupUpsert {
	u.Add(group.FieldMembersCount, v)
	return u
}

// SetPublic sets the "public" field.
func (u *GroupUpsert) SetPublic(v bool) *GroupUpsert {
	u.Set(group.FieldPublic, v)
	return u
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GroupUpsert) UpdatePublic() *GroupUpsert {
	u.SetExcluded(group.FieldPublic)
	return u
}

// SetSn sets the "sn" field.
func (u *GroupUpsert) SetSn(v string) *GroupUpsert {
	u.Set(group.FieldSn, v)
	return u
}

// UpdateSn sets the "sn" field to the value that was provided on create.
func (u *GroupUpsert) UpdateSn() *GroupUpsert {
	u.SetExcluded(group.FieldSn)
	return u
}

// SetIntro sets the "intro" field.
func (u *GroupUpsert) SetIntro(v string) *GroupUpsert {
	u.Set(group.FieldIntro, v)
	return u
}

// UpdateIntro sets the "intro" field to the value that was provided on create.
func (u *GroupUpsert) UpdateIntro() *GroupUpsert {
	u.SetExcluded(group.FieldIntro)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GroupUpsertOne) UpdateNewValues() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(group.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GroupUpsertOne) Ignore() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupUpsertOne) DoNothing() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupCreate.OnConflict
// documentation for more info.
func (u *GroupUpsertOne) Update(set func(*GroupUpsert)) *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GroupUpsertOne) SetCreatedAt(v time.Time) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateCreatedAt() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetName sets the "name" field.
func (u *GroupUpsertOne) SetName(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateName() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateName()
	})
}

// SetMemberID sets the "member_id" field.
func (u *GroupUpsertOne) SetMemberID(v uint64) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetMemberID(v)
	})
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateMemberID() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMemberID()
	})
}

// SetMembersMax sets the "members_max" field.
func (u *GroupUpsertOne) SetMembersMax(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetMembersMax(v)
	})
}

// AddMembersMax adds v to the "members_max" field.
func (u *GroupUpsertOne) AddMembersMax(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.AddMembersMax(v)
	})
}

// UpdateMembersMax sets the "members_max" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateMembersMax() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMembersMax()
	})
}

// SetMembersCount sets the "members_count" field.
func (u *GroupUpsertOne) SetMembersCount(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetMembersCount(v)
	})
}

// AddMembersCount adds v to the "members_count" field.
func (u *GroupUpsertOne) AddMembersCount(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.AddMembersCount(v)
	})
}

// UpdateMembersCount sets the "members_count" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateMembersCount() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMembersCount()
	})
}

// SetPublic sets the "public" field.
func (u *GroupUpsertOne) SetPublic(v bool) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdatePublic() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdatePublic()
	})
}

// SetSn sets the "sn" field.
func (u *GroupUpsertOne) SetSn(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetSn(v)
	})
}

// UpdateSn sets the "sn" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateSn() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateSn()
	})
}

// SetIntro sets the "intro" field.
func (u *GroupUpsertOne) SetIntro(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetIntro(v)
	})
}

// UpdateIntro sets the "intro" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateIntro() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateIntro()
	})
}

// Exec executes the query.
func (u *GroupUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GroupUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GroupUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GroupCreateBulk is the builder for creating many Group entities in bulk.
type GroupCreateBulk struct {
	config
	builders []*GroupCreate
	conflict []sql.ConflictOption
}

// Save creates the Group entities in the database.
func (gcb *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Group, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GroupCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GroupCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Group.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (gcb *GroupCreateBulk) OnConflict(opts ...sql.ConflictOption) *GroupUpsertBulk {
	gcb.conflict = opts
	return &GroupUpsertBulk{
		create: gcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gcb *GroupCreateBulk) OnConflictColumns(columns ...string) *GroupUpsertBulk {
	gcb.conflict = append(gcb.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertBulk{
		create: gcb,
	}
}

// GroupUpsertBulk is the builder for "upsert"-ing
// a bulk of Group nodes.
type GroupUpsertBulk struct {
	create *GroupCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GroupUpsertBulk) UpdateNewValues() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(group.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GroupUpsertBulk) Ignore() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupUpsertBulk) DoNothing() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupCreateBulk.OnConflict
// documentation for more info.
func (u *GroupUpsertBulk) Update(set func(*GroupUpsert)) *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GroupUpsertBulk) SetCreatedAt(v time.Time) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateCreatedAt() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetName sets the "name" field.
func (u *GroupUpsertBulk) SetName(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateName() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateName()
	})
}

// SetMemberID sets the "member_id" field.
func (u *GroupUpsertBulk) SetMemberID(v uint64) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetMemberID(v)
	})
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateMemberID() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMemberID()
	})
}

// SetMembersMax sets the "members_max" field.
func (u *GroupUpsertBulk) SetMembersMax(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetMembersMax(v)
	})
}

// AddMembersMax adds v to the "members_max" field.
func (u *GroupUpsertBulk) AddMembersMax(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.AddMembersMax(v)
	})
}

// UpdateMembersMax sets the "members_max" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateMembersMax() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMembersMax()
	})
}

// SetMembersCount sets the "members_count" field.
func (u *GroupUpsertBulk) SetMembersCount(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetMembersCount(v)
	})
}

// AddMembersCount adds v to the "members_count" field.
func (u *GroupUpsertBulk) AddMembersCount(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.AddMembersCount(v)
	})
}

// UpdateMembersCount sets the "members_count" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateMembersCount() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMembersCount()
	})
}

// SetPublic sets the "public" field.
func (u *GroupUpsertBulk) SetPublic(v bool) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdatePublic() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdatePublic()
	})
}

// SetSn sets the "sn" field.
func (u *GroupUpsertBulk) SetSn(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetSn(v)
	})
}

// UpdateSn sets the "sn" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateSn() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateSn()
	})
}

// SetIntro sets the "intro" field.
func (u *GroupUpsertBulk) SetIntro(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetIntro(v)
	})
}

// UpdateIntro sets the "intro" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateIntro() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateIntro()
	})
}

// Exec executes the query.
func (u *GroupUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GroupCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}