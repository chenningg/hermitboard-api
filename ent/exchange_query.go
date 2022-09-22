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
	"github.com/chenningg/hermitboard-api/ent/exchange"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/ent/transaction"
	"github.com/chenningg/hermitboard-api/pulid"
)

// ExchangeQuery is the builder for querying Exchange entities.
type ExchangeQuery struct {
	config
	limit                 *int
	offset                *int
	unique                *bool
	order                 []OrderFunc
	fields                []string
	predicates            []predicate.Exchange
	withTransactions      *TransactionQuery
	modifiers             []func(*sql.Selector)
	loadTotal             []func(context.Context, []*Exchange) error
	withNamedTransactions map[string]*TransactionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExchangeQuery builder.
func (eq *ExchangeQuery) Where(ps ...predicate.Exchange) *ExchangeQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit adds a limit step to the query.
func (eq *ExchangeQuery) Limit(limit int) *ExchangeQuery {
	eq.limit = &limit
	return eq
}

// Offset adds an offset step to the query.
func (eq *ExchangeQuery) Offset(offset int) *ExchangeQuery {
	eq.offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *ExchangeQuery) Unique(unique bool) *ExchangeQuery {
	eq.unique = &unique
	return eq
}

// Order adds an order step to the query.
func (eq *ExchangeQuery) Order(o ...OrderFunc) *ExchangeQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryTransactions chains the current query on the "transactions" edge.
func (eq *ExchangeQuery) QueryTransactions() *TransactionQuery {
	query := &TransactionQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exchange.Table, exchange.FieldID, selector),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, exchange.TransactionsTable, exchange.TransactionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Exchange entity from the query.
// Returns a *NotFoundError when no Exchange was found.
func (eq *ExchangeQuery) First(ctx context.Context) (*Exchange, error) {
	nodes, err := eq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{exchange.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *ExchangeQuery) FirstX(ctx context.Context) *Exchange {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Exchange ID from the query.
// Returns a *NotFoundError when no Exchange ID was found.
func (eq *ExchangeQuery) FirstID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = eq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{exchange.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *ExchangeQuery) FirstIDX(ctx context.Context) pulid.PULID {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Exchange entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Exchange entity is found.
// Returns a *NotFoundError when no Exchange entities are found.
func (eq *ExchangeQuery) Only(ctx context.Context) (*Exchange, error) {
	nodes, err := eq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{exchange.Label}
	default:
		return nil, &NotSingularError{exchange.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *ExchangeQuery) OnlyX(ctx context.Context) *Exchange {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Exchange ID in the query.
// Returns a *NotSingularError when more than one Exchange ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *ExchangeQuery) OnlyID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = eq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{exchange.Label}
	default:
		err = &NotSingularError{exchange.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *ExchangeQuery) OnlyIDX(ctx context.Context) pulid.PULID {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Exchanges.
func (eq *ExchangeQuery) All(ctx context.Context) ([]*Exchange, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eq *ExchangeQuery) AllX(ctx context.Context) []*Exchange {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Exchange IDs.
func (eq *ExchangeQuery) IDs(ctx context.Context) ([]pulid.PULID, error) {
	var ids []pulid.PULID
	if err := eq.Select(exchange.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *ExchangeQuery) IDsX(ctx context.Context) []pulid.PULID {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *ExchangeQuery) Count(ctx context.Context) (int, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eq *ExchangeQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *ExchangeQuery) Exist(ctx context.Context) (bool, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *ExchangeQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExchangeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *ExchangeQuery) Clone() *ExchangeQuery {
	if eq == nil {
		return nil
	}
	return &ExchangeQuery{
		config:           eq.config,
		limit:            eq.limit,
		offset:           eq.offset,
		order:            append([]OrderFunc{}, eq.order...),
		predicates:       append([]predicate.Exchange{}, eq.predicates...),
		withTransactions: eq.withTransactions.Clone(),
		// clone intermediate query.
		sql:    eq.sql.Clone(),
		path:   eq.path,
		unique: eq.unique,
	}
}

// WithTransactions tells the query-builder to eager-load the nodes that are connected to
// the "transactions" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *ExchangeQuery) WithTransactions(opts ...func(*TransactionQuery)) *ExchangeQuery {
	query := &TransactionQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withTransactions = query
	return eq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Exchange.Query().
//		GroupBy(exchange.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *ExchangeQuery) GroupBy(field string, fields ...string) *ExchangeGroupBy {
	grbuild := &ExchangeGroupBy{config: eq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eq.sqlQuery(ctx), nil
	}
	grbuild.label = exchange.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Exchange.Query().
//		Select(exchange.FieldCreatedAt).
//		Scan(ctx, &v)
func (eq *ExchangeQuery) Select(fields ...string) *ExchangeSelect {
	eq.fields = append(eq.fields, fields...)
	selbuild := &ExchangeSelect{ExchangeQuery: eq}
	selbuild.label = exchange.Label
	selbuild.flds, selbuild.scan = &eq.fields, selbuild.Scan
	return selbuild
}

func (eq *ExchangeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range eq.fields {
		if !exchange.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *ExchangeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Exchange, error) {
	var (
		nodes       = []*Exchange{}
		_spec       = eq.querySpec()
		loadedTypes = [1]bool{
			eq.withTransactions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Exchange).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Exchange{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(eq.modifiers) > 0 {
		_spec.Modifiers = eq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withTransactions; query != nil {
		if err := eq.loadTransactions(ctx, query, nodes,
			func(n *Exchange) { n.Edges.Transactions = []*Transaction{} },
			func(n *Exchange, e *Transaction) { n.Edges.Transactions = append(n.Edges.Transactions, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range eq.withNamedTransactions {
		if err := eq.loadTransactions(ctx, query, nodes,
			func(n *Exchange) { n.appendNamedTransactions(name) },
			func(n *Exchange, e *Transaction) { n.appendNamedTransactions(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range eq.loadTotal {
		if err := eq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *ExchangeQuery) loadTransactions(ctx context.Context, query *TransactionQuery, nodes []*Exchange, init func(*Exchange), assign func(*Exchange, *Transaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pulid.PULID]*Exchange)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.InValues(exchange.TransactionsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exchange_transactions
		if fk == nil {
			return fmt.Errorf(`foreign-key "exchange_transactions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "exchange_transactions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (eq *ExchangeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	if len(eq.modifiers) > 0 {
		_spec.Modifiers = eq.modifiers
	}
	_spec.Node.Columns = eq.fields
	if len(eq.fields) > 0 {
		_spec.Unique = eq.unique != nil && *eq.unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *ExchangeQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (eq *ExchangeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   exchange.Table,
			Columns: exchange.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: exchange.FieldID,
			},
		},
		From:   eq.sql,
		Unique: true,
	}
	if unique := eq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := eq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exchange.FieldID)
		for i := range fields {
			if fields[i] != exchange.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *ExchangeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(exchange.Table)
	columns := eq.fields
	if len(columns) == 0 {
		columns = exchange.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.unique != nil && *eq.unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedTransactions tells the query-builder to eager-load the nodes that are connected to the "transactions"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (eq *ExchangeQuery) WithNamedTransactions(name string, opts ...func(*TransactionQuery)) *ExchangeQuery {
	query := &TransactionQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	if eq.withNamedTransactions == nil {
		eq.withNamedTransactions = make(map[string]*TransactionQuery)
	}
	eq.withNamedTransactions[name] = query
	return eq
}

// ExchangeGroupBy is the group-by builder for Exchange entities.
type ExchangeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *ExchangeGroupBy) Aggregate(fns ...AggregateFunc) *ExchangeGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the group-by query and scans the result into the given value.
func (egb *ExchangeGroupBy) Scan(ctx context.Context, v any) error {
	query, err := egb.path(ctx)
	if err != nil {
		return err
	}
	egb.sql = query
	return egb.sqlScan(ctx, v)
}

func (egb *ExchangeGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range egb.fields {
		if !exchange.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := egb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (egb *ExchangeGroupBy) sqlQuery() *sql.Selector {
	selector := egb.sql.Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(egb.fields)+len(egb.fns))
		for _, f := range egb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(egb.fields...)...)
}

// ExchangeSelect is the builder for selecting fields of Exchange entities.
type ExchangeSelect struct {
	*ExchangeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (es *ExchangeSelect) Scan(ctx context.Context, v any) error {
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	es.sql = es.ExchangeQuery.sqlQuery(ctx)
	return es.sqlScan(ctx, v)
}

func (es *ExchangeSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := es.sql.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
