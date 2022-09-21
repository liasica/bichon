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
	"github.com/chatpuppy/puppychat/internal/ent/messageread"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// MessageReadUpdate is the builder for updating MessageRead entities.
type MessageReadUpdate struct {
	config
	hooks     []Hook
	mutation  *MessageReadMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MessageReadUpdate builder.
func (mru *MessageReadUpdate) Where(ps ...predicate.MessageRead) *MessageReadUpdate {
	mru.mutation.Where(ps...)
	return mru
}

// SetMemberID sets the "member_id" field.
func (mru *MessageReadUpdate) SetMemberID(s string) *MessageReadUpdate {
	mru.mutation.SetMemberID(s)
	return mru
}

// SetGroupID sets the "group_id" field.
func (mru *MessageReadUpdate) SetGroupID(s string) *MessageReadUpdate {
	mru.mutation.SetGroupID(s)
	return mru
}

// SetLastID sets the "last_id" field.
func (mru *MessageReadUpdate) SetLastID(s string) *MessageReadUpdate {
	mru.mutation.SetLastID(s)
	return mru
}

// SetLastTime sets the "last_time" field.
func (mru *MessageReadUpdate) SetLastTime(t time.Time) *MessageReadUpdate {
	mru.mutation.SetLastTime(t)
	return mru
}

// SetMember sets the "member" edge to the Member entity.
func (mru *MessageReadUpdate) SetMember(m *Member) *MessageReadUpdate {
	return mru.SetMemberID(m.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (mru *MessageReadUpdate) SetGroup(g *Group) *MessageReadUpdate {
	return mru.SetGroupID(g.ID)
}

// Mutation returns the MessageReadMutation object of the builder.
func (mru *MessageReadUpdate) Mutation() *MessageReadMutation {
	return mru.mutation
}

// ClearMember clears the "member" edge to the Member entity.
func (mru *MessageReadUpdate) ClearMember() *MessageReadUpdate {
	mru.mutation.ClearMember()
	return mru
}

// ClearGroup clears the "group" edge to the Group entity.
func (mru *MessageReadUpdate) ClearGroup() *MessageReadUpdate {
	mru.mutation.ClearGroup()
	return mru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mru *MessageReadUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mru.hooks) == 0 {
		if err = mru.check(); err != nil {
			return 0, err
		}
		affected, err = mru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageReadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mru.check(); err != nil {
				return 0, err
			}
			mru.mutation = mutation
			affected, err = mru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mru.hooks) - 1; i >= 0; i-- {
			if mru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mru *MessageReadUpdate) SaveX(ctx context.Context) int {
	affected, err := mru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mru *MessageReadUpdate) Exec(ctx context.Context) error {
	_, err := mru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mru *MessageReadUpdate) ExecX(ctx context.Context) {
	if err := mru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mru *MessageReadUpdate) check() error {
	if _, ok := mru.mutation.MemberID(); mru.mutation.MemberCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MessageRead.member"`)
	}
	if _, ok := mru.mutation.GroupID(); mru.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MessageRead.group"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mru *MessageReadUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MessageReadUpdate {
	mru.modifiers = append(mru.modifiers, modifiers...)
	return mru
}

func (mru *MessageReadUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageread.Table,
			Columns: messageread.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: messageread.FieldID,
			},
		},
	}
	if ps := mru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mru.mutation.LastID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messageread.FieldLastID,
		})
	}
	if value, ok := mru.mutation.LastTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: messageread.FieldLastTime,
		})
	}
	if mru.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   messageread.MemberTable,
			Columns: []string{messageread.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: member.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mru.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   messageread.MemberTable,
			Columns: []string{messageread.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mru.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   messageread.GroupTable,
			Columns: []string{messageread.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mru.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   messageread.GroupTable,
			Columns: []string{messageread.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = mru.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, mru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messageread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MessageReadUpdateOne is the builder for updating a single MessageRead entity.
type MessageReadUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MessageReadMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetMemberID sets the "member_id" field.
func (mruo *MessageReadUpdateOne) SetMemberID(s string) *MessageReadUpdateOne {
	mruo.mutation.SetMemberID(s)
	return mruo
}

// SetGroupID sets the "group_id" field.
func (mruo *MessageReadUpdateOne) SetGroupID(s string) *MessageReadUpdateOne {
	mruo.mutation.SetGroupID(s)
	return mruo
}

// SetLastID sets the "last_id" field.
func (mruo *MessageReadUpdateOne) SetLastID(s string) *MessageReadUpdateOne {
	mruo.mutation.SetLastID(s)
	return mruo
}

// SetLastTime sets the "last_time" field.
func (mruo *MessageReadUpdateOne) SetLastTime(t time.Time) *MessageReadUpdateOne {
	mruo.mutation.SetLastTime(t)
	return mruo
}

// SetMember sets the "member" edge to the Member entity.
func (mruo *MessageReadUpdateOne) SetMember(m *Member) *MessageReadUpdateOne {
	return mruo.SetMemberID(m.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (mruo *MessageReadUpdateOne) SetGroup(g *Group) *MessageReadUpdateOne {
	return mruo.SetGroupID(g.ID)
}

// Mutation returns the MessageReadMutation object of the builder.
func (mruo *MessageReadUpdateOne) Mutation() *MessageReadMutation {
	return mruo.mutation
}

// ClearMember clears the "member" edge to the Member entity.
func (mruo *MessageReadUpdateOne) ClearMember() *MessageReadUpdateOne {
	mruo.mutation.ClearMember()
	return mruo
}

// ClearGroup clears the "group" edge to the Group entity.
func (mruo *MessageReadUpdateOne) ClearGroup() *MessageReadUpdateOne {
	mruo.mutation.ClearGroup()
	return mruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mruo *MessageReadUpdateOne) Select(field string, fields ...string) *MessageReadUpdateOne {
	mruo.fields = append([]string{field}, fields...)
	return mruo
}

// Save executes the query and returns the updated MessageRead entity.
func (mruo *MessageReadUpdateOne) Save(ctx context.Context) (*MessageRead, error) {
	var (
		err  error
		node *MessageRead
	)
	if len(mruo.hooks) == 0 {
		if err = mruo.check(); err != nil {
			return nil, err
		}
		node, err = mruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageReadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mruo.check(); err != nil {
				return nil, err
			}
			mruo.mutation = mutation
			node, err = mruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mruo.hooks) - 1; i >= 0; i-- {
			if mruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MessageRead)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageReadMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mruo *MessageReadUpdateOne) SaveX(ctx context.Context) *MessageRead {
	node, err := mruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mruo *MessageReadUpdateOne) Exec(ctx context.Context) error {
	_, err := mruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mruo *MessageReadUpdateOne) ExecX(ctx context.Context) {
	if err := mruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mruo *MessageReadUpdateOne) check() error {
	if _, ok := mruo.mutation.MemberID(); mruo.mutation.MemberCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MessageRead.member"`)
	}
	if _, ok := mruo.mutation.GroupID(); mruo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MessageRead.group"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mruo *MessageReadUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MessageReadUpdateOne {
	mruo.modifiers = append(mruo.modifiers, modifiers...)
	return mruo
}

func (mruo *MessageReadUpdateOne) sqlSave(ctx context.Context) (_node *MessageRead, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageread.Table,
			Columns: messageread.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: messageread.FieldID,
			},
		},
	}
	id, ok := mruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MessageRead.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messageread.FieldID)
		for _, f := range fields {
			if !messageread.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != messageread.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mruo.mutation.LastID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messageread.FieldLastID,
		})
	}
	if value, ok := mruo.mutation.LastTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: messageread.FieldLastTime,
		})
	}
	if mruo.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   messageread.MemberTable,
			Columns: []string{messageread.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: member.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mruo.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   messageread.MemberTable,
			Columns: []string{messageread.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mruo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   messageread.GroupTable,
			Columns: []string{messageread.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mruo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   messageread.GroupTable,
			Columns: []string{messageread.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = mruo.modifiers
	_node = &MessageRead{config: mruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messageread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
