// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/messageread"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// MessageReadQuery is the builder for querying MessageRead entities.
type MessageReadQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MessageRead
	withMember *MemberQuery
	withGroup  *GroupQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageReadQuery builder.
func (mrq *MessageReadQuery) Where(ps ...predicate.MessageRead) *MessageReadQuery {
	mrq.predicates = append(mrq.predicates, ps...)
	return mrq
}

// Limit adds a limit step to the query.
func (mrq *MessageReadQuery) Limit(limit int) *MessageReadQuery {
	mrq.limit = &limit
	return mrq
}

// Offset adds an offset step to the query.
func (mrq *MessageReadQuery) Offset(offset int) *MessageReadQuery {
	mrq.offset = &offset
	return mrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mrq *MessageReadQuery) Unique(unique bool) *MessageReadQuery {
	mrq.unique = &unique
	return mrq
}

// Order adds an order step to the query.
func (mrq *MessageReadQuery) Order(o ...OrderFunc) *MessageReadQuery {
	mrq.order = append(mrq.order, o...)
	return mrq
}

// QueryMember chains the current query on the "member" edge.
func (mrq *MessageReadQuery) QueryMember() *MemberQuery {
	query := &MemberQuery{config: mrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(messageread.Table, messageread.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, messageread.MemberTable, messageread.MemberColumn),
		)
		fromU = sqlgraph.SetNeighbors(mrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroup chains the current query on the "group" edge.
func (mrq *MessageReadQuery) QueryGroup() *GroupQuery {
	query := &GroupQuery{config: mrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(messageread.Table, messageread.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, messageread.GroupTable, messageread.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(mrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MessageRead entity from the query.
// Returns a *NotFoundError when no MessageRead was found.
func (mrq *MessageReadQuery) First(ctx context.Context) (*MessageRead, error) {
	nodes, err := mrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{messageread.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mrq *MessageReadQuery) FirstX(ctx context.Context) *MessageRead {
	node, err := mrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MessageRead ID from the query.
// Returns a *NotFoundError when no MessageRead ID was found.
func (mrq *MessageReadQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{messageread.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mrq *MessageReadQuery) FirstIDX(ctx context.Context) string {
	id, err := mrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MessageRead entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MessageRead entity is found.
// Returns a *NotFoundError when no MessageRead entities are found.
func (mrq *MessageReadQuery) Only(ctx context.Context) (*MessageRead, error) {
	nodes, err := mrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{messageread.Label}
	default:
		return nil, &NotSingularError{messageread.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mrq *MessageReadQuery) OnlyX(ctx context.Context) *MessageRead {
	node, err := mrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MessageRead ID in the query.
// Returns a *NotSingularError when more than one MessageRead ID is found.
// Returns a *NotFoundError when no entities are found.
func (mrq *MessageReadQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{messageread.Label}
	default:
		err = &NotSingularError{messageread.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mrq *MessageReadQuery) OnlyIDX(ctx context.Context) string {
	id, err := mrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MessageReads.
func (mrq *MessageReadQuery) All(ctx context.Context) ([]*MessageRead, error) {
	if err := mrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mrq *MessageReadQuery) AllX(ctx context.Context) []*MessageRead {
	nodes, err := mrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MessageRead IDs.
func (mrq *MessageReadQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := mrq.Select(messageread.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mrq *MessageReadQuery) IDsX(ctx context.Context) []string {
	ids, err := mrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mrq *MessageReadQuery) Count(ctx context.Context) (int, error) {
	if err := mrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mrq *MessageReadQuery) CountX(ctx context.Context) int {
	count, err := mrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mrq *MessageReadQuery) Exist(ctx context.Context) (bool, error) {
	if err := mrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mrq *MessageReadQuery) ExistX(ctx context.Context) bool {
	exist, err := mrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageReadQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mrq *MessageReadQuery) Clone() *MessageReadQuery {
	if mrq == nil {
		return nil
	}
	return &MessageReadQuery{
		config:     mrq.config,
		limit:      mrq.limit,
		offset:     mrq.offset,
		order:      append([]OrderFunc{}, mrq.order...),
		predicates: append([]predicate.MessageRead{}, mrq.predicates...),
		withMember: mrq.withMember.Clone(),
		withGroup:  mrq.withGroup.Clone(),
		// clone intermediate query.
		sql:    mrq.sql.Clone(),
		path:   mrq.path,
		unique: mrq.unique,
	}
}

// WithMember tells the query-builder to eager-load the nodes that are connected to
// the "member" edge. The optional arguments are used to configure the query builder of the edge.
func (mrq *MessageReadQuery) WithMember(opts ...func(*MemberQuery)) *MessageReadQuery {
	query := &MemberQuery{config: mrq.config}
	for _, opt := range opts {
		opt(query)
	}
	mrq.withMember = query
	return mrq
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (mrq *MessageReadQuery) WithGroup(opts ...func(*GroupQuery)) *MessageReadQuery {
	query := &GroupQuery{config: mrq.config}
	for _, opt := range opts {
		opt(query)
	}
	mrq.withGroup = query
	return mrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MemberID string `json:"member_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MessageRead.Query().
//		GroupBy(messageread.FieldMemberID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mrq *MessageReadQuery) GroupBy(field string, fields ...string) *MessageReadGroupBy {
	grbuild := &MessageReadGroupBy{config: mrq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mrq.sqlQuery(ctx), nil
	}
	grbuild.label = messageread.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MemberID string `json:"member_id,omitempty"`
//	}
//
//	client.MessageRead.Query().
//		Select(messageread.FieldMemberID).
//		Scan(ctx, &v)
func (mrq *MessageReadQuery) Select(fields ...string) *MessageReadSelect {
	mrq.fields = append(mrq.fields, fields...)
	selbuild := &MessageReadSelect{MessageReadQuery: mrq}
	selbuild.label = messageread.Label
	selbuild.flds, selbuild.scan = &mrq.fields, selbuild.Scan
	return selbuild
}

func (mrq *MessageReadQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mrq.fields {
		if !messageread.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mrq.path != nil {
		prev, err := mrq.path(ctx)
		if err != nil {
			return err
		}
		mrq.sql = prev
	}
	return nil
}

func (mrq *MessageReadQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MessageRead, error) {
	var (
		nodes       = []*MessageRead{}
		_spec       = mrq.querySpec()
		loadedTypes = [2]bool{
			mrq.withMember != nil,
			mrq.withGroup != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*MessageRead).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &MessageRead{config: mrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mrq.modifiers) > 0 {
		_spec.Modifiers = mrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mrq.withMember; query != nil {
		if err := mrq.loadMember(ctx, query, nodes, nil,
			func(n *MessageRead, e *Member) { n.Edges.Member = e }); err != nil {
			return nil, err
		}
	}
	if query := mrq.withGroup; query != nil {
		if err := mrq.loadGroup(ctx, query, nodes, nil,
			func(n *MessageRead, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mrq *MessageReadQuery) loadMember(ctx context.Context, query *MemberQuery, nodes []*MessageRead, init func(*MessageRead), assign func(*MessageRead, *Member)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*MessageRead)
	for i := range nodes {
		fk := nodes[i].MemberID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(member.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "member_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mrq *MessageReadQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*MessageRead, init func(*MessageRead), assign func(*MessageRead, *Group)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*MessageRead)
	for i := range nodes {
		fk := nodes[i].GroupID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mrq *MessageReadQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mrq.querySpec()
	if len(mrq.modifiers) > 0 {
		_spec.Modifiers = mrq.modifiers
	}
	_spec.Node.Columns = mrq.fields
	if len(mrq.fields) > 0 {
		_spec.Unique = mrq.unique != nil && *mrq.unique
	}
	return sqlgraph.CountNodes(ctx, mrq.driver, _spec)
}

func (mrq *MessageReadQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mrq *MessageReadQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageread.Table,
			Columns: messageread.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: messageread.FieldID,
			},
		},
		From:   mrq.sql,
		Unique: true,
	}
	if unique := mrq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messageread.FieldID)
		for i := range fields {
			if fields[i] != messageread.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mrq *MessageReadQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mrq.driver.Dialect())
	t1 := builder.Table(messageread.Table)
	columns := mrq.fields
	if len(columns) == 0 {
		columns = messageread.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mrq.sql != nil {
		selector = mrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mrq.unique != nil && *mrq.unique {
		selector.Distinct()
	}
	for _, m := range mrq.modifiers {
		m(selector)
	}
	for _, p := range mrq.predicates {
		p(selector)
	}
	for _, p := range mrq.order {
		p(selector)
	}
	if offset := mrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mrq *MessageReadQuery) Modify(modifiers ...func(s *sql.Selector)) *MessageReadSelect {
	mrq.modifiers = append(mrq.modifiers, modifiers...)
	return mrq.Select()
}

// MessageReadGroupBy is the group-by builder for MessageRead entities.
type MessageReadGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mrgb *MessageReadGroupBy) Aggregate(fns ...AggregateFunc) *MessageReadGroupBy {
	mrgb.fns = append(mrgb.fns, fns...)
	return mrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mrgb *MessageReadGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mrgb.path(ctx)
	if err != nil {
		return err
	}
	mrgb.sql = query
	return mrgb.sqlScan(ctx, v)
}

func (mrgb *MessageReadGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mrgb.fields {
		if !messageread.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mrgb *MessageReadGroupBy) sqlQuery() *sql.Selector {
	selector := mrgb.sql.Select()
	aggregation := make([]string, 0, len(mrgb.fns))
	for _, fn := range mrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mrgb.fields)+len(mrgb.fns))
		for _, f := range mrgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mrgb.fields...)...)
}

// MessageReadSelect is the builder for selecting fields of MessageRead entities.
type MessageReadSelect struct {
	*MessageReadQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mrs *MessageReadSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mrs.prepareQuery(ctx); err != nil {
		return err
	}
	mrs.sql = mrs.MessageReadQuery.sqlQuery(ctx)
	return mrs.sqlScan(ctx, v)
}

func (mrs *MessageReadSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mrs.sql.Query()
	if err := mrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mrs *MessageReadSelect) Modify(modifiers ...func(s *sql.Selector)) *MessageReadSelect {
	mrs.modifiers = append(mrs.modifiers, modifiers...)
	return mrs
}
