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
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/connection"
	"github.com/chenningg/hermitboard-api/ent/portfolio"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/ent/source"
	"github.com/chenningg/hermitboard-api/pulid"
)

// ConnectionQuery is the builder for querying Connection entities.
type ConnectionQuery struct {
	config
	limit               *int
	offset              *int
	unique              *bool
	order               []OrderFunc
	fields              []string
	predicates          []predicate.Connection
	withSource          *SourceQuery
	withAccount         *AccountQuery
	withPortfolios      *PortfolioQuery
	modifiers           []func(*sql.Selector)
	loadTotal           []func(context.Context, []*Connection) error
	withNamedPortfolios map[string]*PortfolioQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConnectionQuery builder.
func (cq *ConnectionQuery) Where(ps ...predicate.Connection) *ConnectionQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *ConnectionQuery) Limit(limit int) *ConnectionQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *ConnectionQuery) Offset(offset int) *ConnectionQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ConnectionQuery) Unique(unique bool) *ConnectionQuery {
	cq.unique = &unique
	return cq
}

// Order adds an order step to the query.
func (cq *ConnectionQuery) Order(o ...OrderFunc) *ConnectionQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QuerySource chains the current query on the "source" edge.
func (cq *ConnectionQuery) QuerySource() *SourceQuery {
	query := &SourceQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connection.Table, connection.FieldID, selector),
			sqlgraph.To(source.Table, source.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, connection.SourceTable, connection.SourceColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAccount chains the current query on the "account" edge.
func (cq *ConnectionQuery) QueryAccount() *AccountQuery {
	query := &AccountQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connection.Table, connection.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, connection.AccountTable, connection.AccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPortfolios chains the current query on the "portfolios" edge.
func (cq *ConnectionQuery) QueryPortfolios() *PortfolioQuery {
	query := &PortfolioQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connection.Table, connection.FieldID, selector),
			sqlgraph.To(portfolio.Table, portfolio.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, connection.PortfoliosTable, connection.PortfoliosPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Connection entity from the query.
// Returns a *NotFoundError when no Connection was found.
func (cq *ConnectionQuery) First(ctx context.Context) (*Connection, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{connection.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ConnectionQuery) FirstX(ctx context.Context) *Connection {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Connection ID from the query.
// Returns a *NotFoundError when no Connection ID was found.
func (cq *ConnectionQuery) FirstID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{connection.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ConnectionQuery) FirstIDX(ctx context.Context) pulid.PULID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Connection entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Connection entity is found.
// Returns a *NotFoundError when no Connection entities are found.
func (cq *ConnectionQuery) Only(ctx context.Context) (*Connection, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{connection.Label}
	default:
		return nil, &NotSingularError{connection.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ConnectionQuery) OnlyX(ctx context.Context) *Connection {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Connection ID in the query.
// Returns a *NotSingularError when more than one Connection ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ConnectionQuery) OnlyID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{connection.Label}
	default:
		err = &NotSingularError{connection.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ConnectionQuery) OnlyIDX(ctx context.Context) pulid.PULID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Connections.
func (cq *ConnectionQuery) All(ctx context.Context) ([]*Connection, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *ConnectionQuery) AllX(ctx context.Context) []*Connection {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Connection IDs.
func (cq *ConnectionQuery) IDs(ctx context.Context) ([]pulid.PULID, error) {
	var ids []pulid.PULID
	if err := cq.Select(connection.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ConnectionQuery) IDsX(ctx context.Context) []pulid.PULID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ConnectionQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ConnectionQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ConnectionQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ConnectionQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConnectionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ConnectionQuery) Clone() *ConnectionQuery {
	if cq == nil {
		return nil
	}
	return &ConnectionQuery{
		config:         cq.config,
		limit:          cq.limit,
		offset:         cq.offset,
		order:          append([]OrderFunc{}, cq.order...),
		predicates:     append([]predicate.Connection{}, cq.predicates...),
		withSource:     cq.withSource.Clone(),
		withAccount:    cq.withAccount.Clone(),
		withPortfolios: cq.withPortfolios.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithSource tells the query-builder to eager-load the nodes that are connected to
// the "source" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectionQuery) WithSource(opts ...func(*SourceQuery)) *ConnectionQuery {
	query := &SourceQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withSource = query
	return cq
}

// WithAccount tells the query-builder to eager-load the nodes that are connected to
// the "account" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectionQuery) WithAccount(opts ...func(*AccountQuery)) *ConnectionQuery {
	query := &AccountQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withAccount = query
	return cq
}

// WithPortfolios tells the query-builder to eager-load the nodes that are connected to
// the "portfolios" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectionQuery) WithPortfolios(opts ...func(*PortfolioQuery)) *ConnectionQuery {
	query := &PortfolioQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withPortfolios = query
	return cq
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
//	client.Connection.Query().
//		GroupBy(connection.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ConnectionQuery) GroupBy(field string, fields ...string) *ConnectionGroupBy {
	grbuild := &ConnectionGroupBy{config: cq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(ctx), nil
	}
	grbuild.label = connection.Label
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
//	client.Connection.Query().
//		Select(connection.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *ConnectionQuery) Select(fields ...string) *ConnectionSelect {
	cq.fields = append(cq.fields, fields...)
	selbuild := &ConnectionSelect{ConnectionQuery: cq}
	selbuild.label = connection.Label
	selbuild.flds, selbuild.scan = &cq.fields, selbuild.Scan
	return selbuild
}

func (cq *ConnectionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !connection.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ConnectionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Connection, error) {
	var (
		nodes       = []*Connection{}
		_spec       = cq.querySpec()
		loadedTypes = [3]bool{
			cq.withSource != nil,
			cq.withAccount != nil,
			cq.withPortfolios != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Connection).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Connection{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withSource; query != nil {
		if err := cq.loadSource(ctx, query, nodes, nil,
			func(n *Connection, e *Source) { n.Edges.Source = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withAccount; query != nil {
		if err := cq.loadAccount(ctx, query, nodes, nil,
			func(n *Connection, e *Account) { n.Edges.Account = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withPortfolios; query != nil {
		if err := cq.loadPortfolios(ctx, query, nodes,
			func(n *Connection) { n.Edges.Portfolios = []*Portfolio{} },
			func(n *Connection, e *Portfolio) { n.Edges.Portfolios = append(n.Edges.Portfolios, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedPortfolios {
		if err := cq.loadPortfolios(ctx, query, nodes,
			func(n *Connection) { n.appendNamedPortfolios(name) },
			func(n *Connection, e *Portfolio) { n.appendNamedPortfolios(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ConnectionQuery) loadSource(ctx context.Context, query *SourceQuery, nodes []*Connection, init func(*Connection), assign func(*Connection, *Source)) error {
	ids := make([]pulid.PULID, 0, len(nodes))
	nodeids := make(map[pulid.PULID][]*Connection)
	for i := range nodes {
		fk := nodes[i].SourceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(source.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "source_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ConnectionQuery) loadAccount(ctx context.Context, query *AccountQuery, nodes []*Connection, init func(*Connection), assign func(*Connection, *Account)) error {
	ids := make([]pulid.PULID, 0, len(nodes))
	nodeids := make(map[pulid.PULID][]*Connection)
	for i := range nodes {
		fk := nodes[i].AccountID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(account.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ConnectionQuery) loadPortfolios(ctx context.Context, query *PortfolioQuery, nodes []*Connection, init func(*Connection), assign func(*Connection, *Portfolio)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[pulid.PULID]*Connection)
	nids := make(map[pulid.PULID]map[*Connection]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(connection.PortfoliosTable)
		s.Join(joinT).On(s.C(portfolio.FieldID), joinT.C(connection.PortfoliosPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(connection.PortfoliosPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(connection.PortfoliosPrimaryKey[1]))
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
				nids[inValue] = map[*Connection]struct{}{byID[outValue]: struct{}{}}
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
			return fmt.Errorf(`unexpected "portfolios" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *ConnectionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ConnectionQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (cq *ConnectionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   connection.Table,
			Columns: connection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: connection.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, connection.FieldID)
		for i := range fields {
			if fields[i] != connection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ConnectionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(connection.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = connection.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedPortfolios tells the query-builder to eager-load the nodes that are connected to the "portfolios"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectionQuery) WithNamedPortfolios(name string, opts ...func(*PortfolioQuery)) *ConnectionQuery {
	query := &PortfolioQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedPortfolios == nil {
		cq.withNamedPortfolios = make(map[string]*PortfolioQuery)
	}
	cq.withNamedPortfolios[name] = query
	return cq
}

// ConnectionGroupBy is the group-by builder for Connection entities.
type ConnectionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ConnectionGroupBy) Aggregate(fns ...AggregateFunc) *ConnectionGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *ConnectionGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

func (cgb *ConnectionGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cgb.fields {
		if !connection.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *ConnectionGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql.Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
		for _, f := range cgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cgb.fields...)...)
}

// ConnectionSelect is the builder for selecting fields of Connection entities.
type ConnectionSelect struct {
	*ConnectionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ConnectionSelect) Scan(ctx context.Context, v any) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.ConnectionQuery.sqlQuery(ctx)
	return cs.sqlScan(ctx, v)
}

func (cs *ConnectionSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := cs.sql.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
