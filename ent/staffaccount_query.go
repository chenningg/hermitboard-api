// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/ent/staffaccount"
	"github.com/chenningg/hermitboard-api/pulid"
)

// StaffAccountQuery is the builder for querying StaffAccount entities.
type StaffAccountQuery struct {
	config
	limit              *int
	offset             *int
	unique             *bool
	order              []OrderFunc
	fields             []string
	predicates         []predicate.StaffAccount
	withAuthRoles      *AuthRoleQuery
	withAuthType       *AuthTypeQuery
	withFKs            bool
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*StaffAccount) error
	withNamedAuthRoles map[string]*AuthRoleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StaffAccountQuery builder.
func (saq *StaffAccountQuery) Where(ps ...predicate.StaffAccount) *StaffAccountQuery {
	saq.predicates = append(saq.predicates, ps...)
	return saq
}

// Limit adds a limit step to the query.
func (saq *StaffAccountQuery) Limit(limit int) *StaffAccountQuery {
	saq.limit = &limit
	return saq
}

// Offset adds an offset step to the query.
func (saq *StaffAccountQuery) Offset(offset int) *StaffAccountQuery {
	saq.offset = &offset
	return saq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (saq *StaffAccountQuery) Unique(unique bool) *StaffAccountQuery {
	saq.unique = &unique
	return saq
}

// Order adds an order step to the query.
func (saq *StaffAccountQuery) Order(o ...OrderFunc) *StaffAccountQuery {
	saq.order = append(saq.order, o...)
	return saq
}

// QueryAuthRoles chains the current query on the "auth_roles" edge.
func (saq *StaffAccountQuery) QueryAuthRoles() *AuthRoleQuery {
	query := &AuthRoleQuery{config: saq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := saq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(staffaccount.Table, staffaccount.FieldID, selector),
			sqlgraph.To(authrole.Table, authrole.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, staffaccount.AuthRolesTable, staffaccount.AuthRolesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(saq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAuthType chains the current query on the "auth_type" edge.
func (saq *StaffAccountQuery) QueryAuthType() *AuthTypeQuery {
	query := &AuthTypeQuery{config: saq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := saq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(staffaccount.Table, staffaccount.FieldID, selector),
			sqlgraph.To(authtype.Table, authtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, staffaccount.AuthTypeTable, staffaccount.AuthTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(saq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StaffAccount entity from the query.
// Returns a *NotFoundError when no StaffAccount was found.
func (saq *StaffAccountQuery) First(ctx context.Context) (*StaffAccount, error) {
	nodes, err := saq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{staffaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (saq *StaffAccountQuery) FirstX(ctx context.Context) *StaffAccount {
	node, err := saq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StaffAccount ID from the query.
// Returns a *NotFoundError when no StaffAccount ID was found.
func (saq *StaffAccountQuery) FirstID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = saq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{staffaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (saq *StaffAccountQuery) FirstIDX(ctx context.Context) pulid.PULID {
	id, err := saq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StaffAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StaffAccount entity is found.
// Returns a *NotFoundError when no StaffAccount entities are found.
func (saq *StaffAccountQuery) Only(ctx context.Context) (*StaffAccount, error) {
	nodes, err := saq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{staffaccount.Label}
	default:
		return nil, &NotSingularError{staffaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (saq *StaffAccountQuery) OnlyX(ctx context.Context) *StaffAccount {
	node, err := saq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StaffAccount ID in the query.
// Returns a *NotSingularError when more than one StaffAccount ID is found.
// Returns a *NotFoundError when no entities are found.
func (saq *StaffAccountQuery) OnlyID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = saq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{staffaccount.Label}
	default:
		err = &NotSingularError{staffaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (saq *StaffAccountQuery) OnlyIDX(ctx context.Context) pulid.PULID {
	id, err := saq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StaffAccounts.
func (saq *StaffAccountQuery) All(ctx context.Context) ([]*StaffAccount, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return saq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (saq *StaffAccountQuery) AllX(ctx context.Context) []*StaffAccount {
	nodes, err := saq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StaffAccount IDs.
func (saq *StaffAccountQuery) IDs(ctx context.Context) ([]pulid.PULID, error) {
	var ids []pulid.PULID
	if err := saq.Select(staffaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (saq *StaffAccountQuery) IDsX(ctx context.Context) []pulid.PULID {
	ids, err := saq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (saq *StaffAccountQuery) Count(ctx context.Context) (int, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return saq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (saq *StaffAccountQuery) CountX(ctx context.Context) int {
	count, err := saq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (saq *StaffAccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return saq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (saq *StaffAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := saq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StaffAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (saq *StaffAccountQuery) Clone() *StaffAccountQuery {
	if saq == nil {
		return nil
	}
	return &StaffAccountQuery{
		config:        saq.config,
		limit:         saq.limit,
		offset:        saq.offset,
		order:         append([]OrderFunc{}, saq.order...),
		predicates:    append([]predicate.StaffAccount{}, saq.predicates...),
		withAuthRoles: saq.withAuthRoles.Clone(),
		withAuthType:  saq.withAuthType.Clone(),
		// clone intermediate query.
		sql:    saq.sql.Clone(),
		path:   saq.path,
		unique: saq.unique,
	}
}

// WithAuthRoles tells the query-builder to eager-load the nodes that are connected to
// the "auth_roles" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *StaffAccountQuery) WithAuthRoles(opts ...func(*AuthRoleQuery)) *StaffAccountQuery {
	query := &AuthRoleQuery{config: saq.config}
	for _, opt := range opts {
		opt(query)
	}
	saq.withAuthRoles = query
	return saq
}

// WithAuthType tells the query-builder to eager-load the nodes that are connected to
// the "auth_type" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *StaffAccountQuery) WithAuthType(opts ...func(*AuthTypeQuery)) *StaffAccountQuery {
	query := &AuthTypeQuery{config: saq.config}
	for _, opt := range opts {
		opt(query)
	}
	saq.withAuthType = query
	return saq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StaffAccount.Query().
//		GroupBy(staffaccount.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (saq *StaffAccountQuery) GroupBy(field string, fields ...string) *StaffAccountGroupBy {
	grbuild := &StaffAccountGroupBy{config: saq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return saq.sqlQuery(ctx), nil
	}
	grbuild.label = staffaccount.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//	}
//
//	client.StaffAccount.Query().
//		Select(staffaccount.FieldCreatedAt).
//		Scan(ctx, &v)
func (saq *StaffAccountQuery) Select(fields ...string) *StaffAccountSelect {
	saq.fields = append(saq.fields, fields...)
	selbuild := &StaffAccountSelect{StaffAccountQuery: saq}
	selbuild.label = staffaccount.Label
	selbuild.flds, selbuild.scan = &saq.fields, selbuild.Scan
	return selbuild
}

func (saq *StaffAccountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range saq.fields {
		if !staffaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if saq.path != nil {
		prev, err := saq.path(ctx)
		if err != nil {
			return err
		}
		saq.sql = prev
	}
	return nil
}

func (saq *StaffAccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*StaffAccount, error) {
	var (
		nodes       = []*StaffAccount{}
		withFKs     = saq.withFKs
		_spec       = saq.querySpec()
		loadedTypes = [2]bool{
			saq.withAuthRoles != nil,
			saq.withAuthType != nil,
		}
	)
	if saq.withAuthType != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, staffaccount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*StaffAccount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &StaffAccount{config: saq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(saq.modifiers) > 0 {
		_spec.Modifiers = saq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, saq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := saq.withAuthRoles; query != nil {
		if err := saq.loadAuthRoles(ctx, query, nodes,
			func(n *StaffAccount) { n.Edges.AuthRoles = []*AuthRole{} },
			func(n *StaffAccount, e *AuthRole) { n.Edges.AuthRoles = append(n.Edges.AuthRoles, e) }); err != nil {
			return nil, err
		}
	}
	if query := saq.withAuthType; query != nil {
		if err := saq.loadAuthType(ctx, query, nodes, nil,
			func(n *StaffAccount, e *AuthType) { n.Edges.AuthType = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range saq.withNamedAuthRoles {
		if err := saq.loadAuthRoles(ctx, query, nodes,
			func(n *StaffAccount) { n.appendNamedAuthRoles(name) },
			func(n *StaffAccount, e *AuthRole) { n.appendNamedAuthRoles(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range saq.loadTotal {
		if err := saq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (saq *StaffAccountQuery) loadAuthRoles(ctx context.Context, query *AuthRoleQuery, nodes []*StaffAccount, init func(*StaffAccount), assign func(*StaffAccount, *AuthRole)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[pulid.PULID]*StaffAccount)
	nids := make(map[pulid.PULID]map[*StaffAccount]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(staffaccount.AuthRolesTable)
		s.Join(joinT).On(s.C(authrole.FieldID), joinT.C(staffaccount.AuthRolesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(staffaccount.AuthRolesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(staffaccount.AuthRolesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(pulid.PULID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*pulid.PULID)
			inValue := *values[1].(*pulid.PULID)
			if nids[inValue] == nil {
				nids[inValue] = map[*StaffAccount]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "auth_roles" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (saq *StaffAccountQuery) loadAuthType(ctx context.Context, query *AuthTypeQuery, nodes []*StaffAccount, init func(*StaffAccount), assign func(*StaffAccount, *AuthType)) error {
	ids := make([]pulid.PULID, 0, len(nodes))
	nodeids := make(map[pulid.PULID][]*StaffAccount)
	for i := range nodes {
		if nodes[i].staff_account_auth_type == nil {
			continue
		}
		fk := *nodes[i].staff_account_auth_type
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(authtype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "staff_account_auth_type" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (saq *StaffAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := saq.querySpec()
	if len(saq.modifiers) > 0 {
		_spec.Modifiers = saq.modifiers
	}
	_spec.Node.Columns = saq.fields
	if len(saq.fields) > 0 {
		_spec.Unique = saq.unique != nil && *saq.unique
	}
	return sqlgraph.CountNodes(ctx, saq.driver, _spec)
}

func (saq *StaffAccountQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := saq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (saq *StaffAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   staffaccount.Table,
			Columns: staffaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: staffaccount.FieldID,
			},
		},
		From:   saq.sql,
		Unique: true,
	}
	if unique := saq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := saq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, staffaccount.FieldID)
		for i := range fields {
			if fields[i] != staffaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := saq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := saq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := saq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := saq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (saq *StaffAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(saq.driver.Dialect())
	t1 := builder.Table(staffaccount.Table)
	columns := saq.fields
	if len(columns) == 0 {
		columns = staffaccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if saq.sql != nil {
		selector = saq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if saq.unique != nil && *saq.unique {
		selector.Distinct()
	}
	for _, p := range saq.predicates {
		p(selector)
	}
	for _, p := range saq.order {
		p(selector)
	}
	if offset := saq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := saq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAuthRoles tells the query-builder to eager-load the nodes that are connected to the "auth_roles"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (saq *StaffAccountQuery) WithNamedAuthRoles(name string, opts ...func(*AuthRoleQuery)) *StaffAccountQuery {
	query := &AuthRoleQuery{config: saq.config}
	for _, opt := range opts {
		opt(query)
	}
	if saq.withNamedAuthRoles == nil {
		saq.withNamedAuthRoles = make(map[string]*AuthRoleQuery)
	}
	saq.withNamedAuthRoles[name] = query
	return saq
}

// StaffAccountGroupBy is the group-by builder for StaffAccount entities.
type StaffAccountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sagb *StaffAccountGroupBy) Aggregate(fns ...AggregateFunc) *StaffAccountGroupBy {
	sagb.fns = append(sagb.fns, fns...)
	return sagb
}

// Scan applies the group-by query and scans the result into the given value.
func (sagb *StaffAccountGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sagb.path(ctx)
	if err != nil {
		return err
	}
	sagb.sql = query
	return sagb.sqlScan(ctx, v)
}

func (sagb *StaffAccountGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sagb.fields {
		if !staffaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sagb *StaffAccountGroupBy) sqlQuery() *sql.Selector {
	selector := sagb.sql.Select()
	aggregation := make([]string, 0, len(sagb.fns))
	for _, fn := range sagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sagb.fields)+len(sagb.fns))
		for _, f := range sagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sagb.fields...)...)
}

// StaffAccountSelect is the builder for selecting fields of StaffAccount entities.
type StaffAccountSelect struct {
	*StaffAccountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sas *StaffAccountSelect) Scan(ctx context.Context, v any) error {
	if err := sas.prepareQuery(ctx); err != nil {
		return err
	}
	sas.sql = sas.StaffAccountQuery.sqlQuery(ctx)
	return sas.sqlScan(ctx, v)
}

func (sas *StaffAccountSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := sas.sql.Query()
	if err := sas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
