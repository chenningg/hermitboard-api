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
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/ent/transaction"
	"github.com/chenningg/hermitboard-api/ent/transactiontype"
	"github.com/chenningg/hermitboard-api/pulid"
)

// TransactionTypeQuery is the builder for querying TransactionType entities.
type TransactionTypeQuery struct {
	config
	limit                 *int
	offset                *int
	unique                *bool
	order                 []OrderFunc
	fields                []string
	predicates            []predicate.TransactionType
	withTransactions      *TransactionQuery
	modifiers             []func(*sql.Selector)
	loadTotal             []func(context.Context, []*TransactionType) error
	withNamedTransactions map[string]*TransactionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TransactionTypeQuery builder.
func (ttq *TransactionTypeQuery) Where(ps ...predicate.TransactionType) *TransactionTypeQuery {
	ttq.predicates = append(ttq.predicates, ps...)
	return ttq
}

// Limit adds a limit step to the query.
func (ttq *TransactionTypeQuery) Limit(limit int) *TransactionTypeQuery {
	ttq.limit = &limit
	return ttq
}

// Offset adds an offset step to the query.
func (ttq *TransactionTypeQuery) Offset(offset int) *TransactionTypeQuery {
	ttq.offset = &offset
	return ttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ttq *TransactionTypeQuery) Unique(unique bool) *TransactionTypeQuery {
	ttq.unique = &unique
	return ttq
}

// Order adds an order step to the query.
func (ttq *TransactionTypeQuery) Order(o ...OrderFunc) *TransactionTypeQuery {
	ttq.order = append(ttq.order, o...)
	return ttq
}

// QueryTransactions chains the current query on the "transactions" edge.
func (ttq *TransactionTypeQuery) QueryTransactions() *TransactionQuery {
	query := &TransactionQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transactiontype.Table, transactiontype.FieldID, selector),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, transactiontype.TransactionsTable, transactiontype.TransactionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TransactionType entity from the query.
// Returns a *NotFoundError when no TransactionType was found.
func (ttq *TransactionTypeQuery) First(ctx context.Context) (*TransactionType, error) {
	nodes, err := ttq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{transactiontype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ttq *TransactionTypeQuery) FirstX(ctx context.Context) *TransactionType {
	node, err := ttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TransactionType ID from the query.
// Returns a *NotFoundError when no TransactionType ID was found.
func (ttq *TransactionTypeQuery) FirstID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = ttq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{transactiontype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ttq *TransactionTypeQuery) FirstIDX(ctx context.Context) pulid.PULID {
	id, err := ttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TransactionType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TransactionType entity is found.
// Returns a *NotFoundError when no TransactionType entities are found.
func (ttq *TransactionTypeQuery) Only(ctx context.Context) (*TransactionType, error) {
	nodes, err := ttq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{transactiontype.Label}
	default:
		return nil, &NotSingularError{transactiontype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ttq *TransactionTypeQuery) OnlyX(ctx context.Context) *TransactionType {
	node, err := ttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TransactionType ID in the query.
// Returns a *NotSingularError when more than one TransactionType ID is found.
// Returns a *NotFoundError when no entities are found.
func (ttq *TransactionTypeQuery) OnlyID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = ttq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{transactiontype.Label}
	default:
		err = &NotSingularError{transactiontype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ttq *TransactionTypeQuery) OnlyIDX(ctx context.Context) pulid.PULID {
	id, err := ttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TransactionTypes.
func (ttq *TransactionTypeQuery) All(ctx context.Context) ([]*TransactionType, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ttq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ttq *TransactionTypeQuery) AllX(ctx context.Context) []*TransactionType {
	nodes, err := ttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TransactionType IDs.
func (ttq *TransactionTypeQuery) IDs(ctx context.Context) ([]pulid.PULID, error) {
	var ids []pulid.PULID
	if err := ttq.Select(transactiontype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ttq *TransactionTypeQuery) IDsX(ctx context.Context) []pulid.PULID {
	ids, err := ttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ttq *TransactionTypeQuery) Count(ctx context.Context) (int, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ttq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ttq *TransactionTypeQuery) CountX(ctx context.Context) int {
	count, err := ttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ttq *TransactionTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ttq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ttq *TransactionTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TransactionTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ttq *TransactionTypeQuery) Clone() *TransactionTypeQuery {
	if ttq == nil {
		return nil
	}
	return &TransactionTypeQuery{
		config:           ttq.config,
		limit:            ttq.limit,
		offset:           ttq.offset,
		order:            append([]OrderFunc{}, ttq.order...),
		predicates:       append([]predicate.TransactionType{}, ttq.predicates...),
		withTransactions: ttq.withTransactions.Clone(),
		// clone intermediate query.
		sql:    ttq.sql.Clone(),
		path:   ttq.path,
		unique: ttq.unique,
	}
}

// WithTransactions tells the query-builder to eager-load the nodes that are connected to
// the "transactions" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TransactionTypeQuery) WithTransactions(opts ...func(*TransactionQuery)) *TransactionTypeQuery {
	query := &TransactionQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withTransactions = query
	return ttq
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
//	client.TransactionType.Query().
//		GroupBy(transactiontype.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ttq *TransactionTypeQuery) GroupBy(field string, fields ...string) *TransactionTypeGroupBy {
	grbuild := &TransactionTypeGroupBy{config: ttq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ttq.sqlQuery(ctx), nil
	}
	grbuild.label = transactiontype.Label
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
//	client.TransactionType.Query().
//		Select(transactiontype.FieldCreatedAt).
//		Scan(ctx, &v)
func (ttq *TransactionTypeQuery) Select(fields ...string) *TransactionTypeSelect {
	ttq.fields = append(ttq.fields, fields...)
	selbuild := &TransactionTypeSelect{TransactionTypeQuery: ttq}
	selbuild.label = transactiontype.Label
	selbuild.flds, selbuild.scan = &ttq.fields, selbuild.Scan
	return selbuild
}

func (ttq *TransactionTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ttq.fields {
		if !transactiontype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ttq.path != nil {
		prev, err := ttq.path(ctx)
		if err != nil {
			return err
		}
		ttq.sql = prev
	}
	return nil
}

func (ttq *TransactionTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TransactionType, error) {
	var (
		nodes       = []*TransactionType{}
		_spec       = ttq.querySpec()
		loadedTypes = [1]bool{
			ttq.withTransactions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TransactionType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TransactionType{config: ttq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ttq.modifiers) > 0 {
		_spec.Modifiers = ttq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ttq.withTransactions; query != nil {
		if err := ttq.loadTransactions(ctx, query, nodes,
			func(n *TransactionType) { n.Edges.Transactions = []*Transaction{} },
			func(n *TransactionType, e *Transaction) { n.Edges.Transactions = append(n.Edges.Transactions, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ttq.withNamedTransactions {
		if err := ttq.loadTransactions(ctx, query, nodes,
			func(n *TransactionType) { n.appendNamedTransactions(name) },
			func(n *TransactionType, e *Transaction) { n.appendNamedTransactions(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range ttq.loadTotal {
		if err := ttq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ttq *TransactionTypeQuery) loadTransactions(ctx context.Context, query *TransactionQuery, nodes []*TransactionType, init func(*TransactionType), assign func(*TransactionType, *Transaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pulid.PULID]*TransactionType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.InValues(transactiontype.TransactionsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TransactionTypeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "transaction_type_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ttq *TransactionTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ttq.querySpec()
	if len(ttq.modifiers) > 0 {
		_spec.Modifiers = ttq.modifiers
	}
	_spec.Node.Columns = ttq.fields
	if len(ttq.fields) > 0 {
		_spec.Unique = ttq.unique != nil && *ttq.unique
	}
	return sqlgraph.CountNodes(ctx, ttq.driver, _spec)
}

func (ttq *TransactionTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := ttq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (ttq *TransactionTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactiontype.Table,
			Columns: transactiontype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transactiontype.FieldID,
			},
		},
		From:   ttq.sql,
		Unique: true,
	}
	if unique := ttq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ttq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transactiontype.FieldID)
		for i := range fields {
			if fields[i] != transactiontype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ttq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ttq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ttq *TransactionTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ttq.driver.Dialect())
	t1 := builder.Table(transactiontype.Table)
	columns := ttq.fields
	if len(columns) == 0 {
		columns = transactiontype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ttq.sql != nil {
		selector = ttq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ttq.unique != nil && *ttq.unique {
		selector.Distinct()
	}
	for _, p := range ttq.predicates {
		p(selector)
	}
	for _, p := range ttq.order {
		p(selector)
	}
	if offset := ttq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ttq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedTransactions tells the query-builder to eager-load the nodes that are connected to the "transactions"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ttq *TransactionTypeQuery) WithNamedTransactions(name string, opts ...func(*TransactionQuery)) *TransactionTypeQuery {
	query := &TransactionQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	if ttq.withNamedTransactions == nil {
		ttq.withNamedTransactions = make(map[string]*TransactionQuery)
	}
	ttq.withNamedTransactions[name] = query
	return ttq
}

// TransactionTypeGroupBy is the group-by builder for TransactionType entities.
type TransactionTypeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ttgb *TransactionTypeGroupBy) Aggregate(fns ...AggregateFunc) *TransactionTypeGroupBy {
	ttgb.fns = append(ttgb.fns, fns...)
	return ttgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ttgb *TransactionTypeGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ttgb.path(ctx)
	if err != nil {
		return err
	}
	ttgb.sql = query
	return ttgb.sqlScan(ctx, v)
}

func (ttgb *TransactionTypeGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ttgb.fields {
		if !transactiontype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ttgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ttgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ttgb *TransactionTypeGroupBy) sqlQuery() *sql.Selector {
	selector := ttgb.sql.Select()
	aggregation := make([]string, 0, len(ttgb.fns))
	for _, fn := range ttgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ttgb.fields)+len(ttgb.fns))
		for _, f := range ttgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ttgb.fields...)...)
}

// TransactionTypeSelect is the builder for selecting fields of TransactionType entities.
type TransactionTypeSelect struct {
	*TransactionTypeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tts *TransactionTypeSelect) Scan(ctx context.Context, v any) error {
	if err := tts.prepareQuery(ctx); err != nil {
		return err
	}
	tts.sql = tts.TransactionTypeQuery.sqlQuery(ctx)
	return tts.sqlScan(ctx, v)
}

func (tts *TransactionTypeSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := tts.sql.Query()
	if err := tts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
