// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/blockchain"
	"github.com/chenningg/hermitboard-api/ent/blockchaincryptocurrency"
	"github.com/chenningg/hermitboard-api/ent/cryptocurrency"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// BlockchainCryptocurrencyQuery is the builder for querying BlockchainCryptocurrency entities.
type BlockchainCryptocurrencyQuery struct {
	config
	limit              *int
	offset             *int
	unique             *bool
	order              []OrderFunc
	fields             []string
	predicates         []predicate.BlockchainCryptocurrency
	withBlockchain     *BlockchainQuery
	withCryptocurrency *CryptocurrencyQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*BlockchainCryptocurrency) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlockchainCryptocurrencyQuery builder.
func (bcq *BlockchainCryptocurrencyQuery) Where(ps ...predicate.BlockchainCryptocurrency) *BlockchainCryptocurrencyQuery {
	bcq.predicates = append(bcq.predicates, ps...)
	return bcq
}

// Limit adds a limit step to the query.
func (bcq *BlockchainCryptocurrencyQuery) Limit(limit int) *BlockchainCryptocurrencyQuery {
	bcq.limit = &limit
	return bcq
}

// Offset adds an offset step to the query.
func (bcq *BlockchainCryptocurrencyQuery) Offset(offset int) *BlockchainCryptocurrencyQuery {
	bcq.offset = &offset
	return bcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bcq *BlockchainCryptocurrencyQuery) Unique(unique bool) *BlockchainCryptocurrencyQuery {
	bcq.unique = &unique
	return bcq
}

// Order adds an order step to the query.
func (bcq *BlockchainCryptocurrencyQuery) Order(o ...OrderFunc) *BlockchainCryptocurrencyQuery {
	bcq.order = append(bcq.order, o...)
	return bcq
}

// QueryBlockchain chains the current query on the "blockchain" edge.
func (bcq *BlockchainCryptocurrencyQuery) QueryBlockchain() *BlockchainQuery {
	query := &BlockchainQuery{config: bcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blockchaincryptocurrency.Table, blockchaincryptocurrency.FieldID, selector),
			sqlgraph.To(blockchain.Table, blockchain.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, blockchaincryptocurrency.BlockchainTable, blockchaincryptocurrency.BlockchainColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCryptocurrency chains the current query on the "cryptocurrency" edge.
func (bcq *BlockchainCryptocurrencyQuery) QueryCryptocurrency() *CryptocurrencyQuery {
	query := &CryptocurrencyQuery{config: bcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blockchaincryptocurrency.Table, blockchaincryptocurrency.FieldID, selector),
			sqlgraph.To(cryptocurrency.Table, cryptocurrency.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, blockchaincryptocurrency.CryptocurrencyTable, blockchaincryptocurrency.CryptocurrencyColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BlockchainCryptocurrency entity from the query.
// Returns a *NotFoundError when no BlockchainCryptocurrency was found.
func (bcq *BlockchainCryptocurrencyQuery) First(ctx context.Context) (*BlockchainCryptocurrency, error) {
	nodes, err := bcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{blockchaincryptocurrency.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bcq *BlockchainCryptocurrencyQuery) FirstX(ctx context.Context) *BlockchainCryptocurrency {
	node, err := bcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BlockchainCryptocurrency ID from the query.
// Returns a *NotFoundError when no BlockchainCryptocurrency ID was found.
func (bcq *BlockchainCryptocurrencyQuery) FirstID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = bcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blockchaincryptocurrency.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bcq *BlockchainCryptocurrencyQuery) FirstIDX(ctx context.Context) pulid.PULID {
	id, err := bcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BlockchainCryptocurrency entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BlockchainCryptocurrency entity is found.
// Returns a *NotFoundError when no BlockchainCryptocurrency entities are found.
func (bcq *BlockchainCryptocurrencyQuery) Only(ctx context.Context) (*BlockchainCryptocurrency, error) {
	nodes, err := bcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{blockchaincryptocurrency.Label}
	default:
		return nil, &NotSingularError{blockchaincryptocurrency.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bcq *BlockchainCryptocurrencyQuery) OnlyX(ctx context.Context) *BlockchainCryptocurrency {
	node, err := bcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BlockchainCryptocurrency ID in the query.
// Returns a *NotSingularError when more than one BlockchainCryptocurrency ID is found.
// Returns a *NotFoundError when no entities are found.
func (bcq *BlockchainCryptocurrencyQuery) OnlyID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = bcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blockchaincryptocurrency.Label}
	default:
		err = &NotSingularError{blockchaincryptocurrency.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bcq *BlockchainCryptocurrencyQuery) OnlyIDX(ctx context.Context) pulid.PULID {
	id, err := bcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BlockchainCryptocurrencies.
func (bcq *BlockchainCryptocurrencyQuery) All(ctx context.Context) ([]*BlockchainCryptocurrency, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bcq *BlockchainCryptocurrencyQuery) AllX(ctx context.Context) []*BlockchainCryptocurrency {
	nodes, err := bcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BlockchainCryptocurrency IDs.
func (bcq *BlockchainCryptocurrencyQuery) IDs(ctx context.Context) ([]pulid.PULID, error) {
	var ids []pulid.PULID
	if err := bcq.Select(blockchaincryptocurrency.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bcq *BlockchainCryptocurrencyQuery) IDsX(ctx context.Context) []pulid.PULID {
	ids, err := bcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bcq *BlockchainCryptocurrencyQuery) Count(ctx context.Context) (int, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bcq *BlockchainCryptocurrencyQuery) CountX(ctx context.Context) int {
	count, err := bcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bcq *BlockchainCryptocurrencyQuery) Exist(ctx context.Context) (bool, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bcq *BlockchainCryptocurrencyQuery) ExistX(ctx context.Context) bool {
	exist, err := bcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlockchainCryptocurrencyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bcq *BlockchainCryptocurrencyQuery) Clone() *BlockchainCryptocurrencyQuery {
	if bcq == nil {
		return nil
	}
	return &BlockchainCryptocurrencyQuery{
		config:             bcq.config,
		limit:              bcq.limit,
		offset:             bcq.offset,
		order:              append([]OrderFunc{}, bcq.order...),
		predicates:         append([]predicate.BlockchainCryptocurrency{}, bcq.predicates...),
		withBlockchain:     bcq.withBlockchain.Clone(),
		withCryptocurrency: bcq.withCryptocurrency.Clone(),
		// clone intermediate query.
		sql:    bcq.sql.Clone(),
		path:   bcq.path,
		unique: bcq.unique,
	}
}

// WithBlockchain tells the query-builder to eager-load the nodes that are connected to
// the "blockchain" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BlockchainCryptocurrencyQuery) WithBlockchain(opts ...func(*BlockchainQuery)) *BlockchainCryptocurrencyQuery {
	query := &BlockchainQuery{config: bcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcq.withBlockchain = query
	return bcq
}

// WithCryptocurrency tells the query-builder to eager-load the nodes that are connected to
// the "cryptocurrency" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BlockchainCryptocurrencyQuery) WithCryptocurrency(opts ...func(*CryptocurrencyQuery)) *BlockchainCryptocurrencyQuery {
	query := &CryptocurrencyQuery{config: bcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcq.withCryptocurrency = query
	return bcq
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
//	client.BlockchainCryptocurrency.Query().
//		GroupBy(blockchaincryptocurrency.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bcq *BlockchainCryptocurrencyQuery) GroupBy(field string, fields ...string) *BlockchainCryptocurrencyGroupBy {
	grbuild := &BlockchainCryptocurrencyGroupBy{config: bcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bcq.sqlQuery(ctx), nil
	}
	grbuild.label = blockchaincryptocurrency.Label
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
//	client.BlockchainCryptocurrency.Query().
//		Select(blockchaincryptocurrency.FieldCreatedAt).
//		Scan(ctx, &v)
func (bcq *BlockchainCryptocurrencyQuery) Select(fields ...string) *BlockchainCryptocurrencySelect {
	bcq.fields = append(bcq.fields, fields...)
	selbuild := &BlockchainCryptocurrencySelect{BlockchainCryptocurrencyQuery: bcq}
	selbuild.label = blockchaincryptocurrency.Label
	selbuild.flds, selbuild.scan = &bcq.fields, selbuild.Scan
	return selbuild
}

func (bcq *BlockchainCryptocurrencyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bcq.fields {
		if !blockchaincryptocurrency.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bcq.path != nil {
		prev, err := bcq.path(ctx)
		if err != nil {
			return err
		}
		bcq.sql = prev
	}
	return nil
}

func (bcq *BlockchainCryptocurrencyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BlockchainCryptocurrency, error) {
	var (
		nodes       = []*BlockchainCryptocurrency{}
		_spec       = bcq.querySpec()
		loadedTypes = [2]bool{
			bcq.withBlockchain != nil,
			bcq.withCryptocurrency != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BlockchainCryptocurrency).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BlockchainCryptocurrency{config: bcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bcq.modifiers) > 0 {
		_spec.Modifiers = bcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bcq.withBlockchain; query != nil {
		if err := bcq.loadBlockchain(ctx, query, nodes, nil,
			func(n *BlockchainCryptocurrency, e *Blockchain) { n.Edges.Blockchain = e }); err != nil {
			return nil, err
		}
	}
	if query := bcq.withCryptocurrency; query != nil {
		if err := bcq.loadCryptocurrency(ctx, query, nodes, nil,
			func(n *BlockchainCryptocurrency, e *Cryptocurrency) { n.Edges.Cryptocurrency = e }); err != nil {
			return nil, err
		}
	}
	for i := range bcq.loadTotal {
		if err := bcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bcq *BlockchainCryptocurrencyQuery) loadBlockchain(ctx context.Context, query *BlockchainQuery, nodes []*BlockchainCryptocurrency, init func(*BlockchainCryptocurrency), assign func(*BlockchainCryptocurrency, *Blockchain)) error {
	ids := make([]pulid.PULID, 0, len(nodes))
	nodeids := make(map[pulid.PULID][]*BlockchainCryptocurrency)
	for i := range nodes {
		fk := nodes[i].BlockchainID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(blockchain.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "blockchain_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bcq *BlockchainCryptocurrencyQuery) loadCryptocurrency(ctx context.Context, query *CryptocurrencyQuery, nodes []*BlockchainCryptocurrency, init func(*BlockchainCryptocurrency), assign func(*BlockchainCryptocurrency, *Cryptocurrency)) error {
	ids := make([]pulid.PULID, 0, len(nodes))
	nodeids := make(map[pulid.PULID][]*BlockchainCryptocurrency)
	for i := range nodes {
		fk := nodes[i].CryptocurrencyID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(cryptocurrency.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cryptocurrency_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (bcq *BlockchainCryptocurrencyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bcq.querySpec()
	if len(bcq.modifiers) > 0 {
		_spec.Modifiers = bcq.modifiers
	}
	_spec.Node.Columns = bcq.fields
	if len(bcq.fields) > 0 {
		_spec.Unique = bcq.unique != nil && *bcq.unique
	}
	return sqlgraph.CountNodes(ctx, bcq.driver, _spec)
}

func (bcq *BlockchainCryptocurrencyQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := bcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (bcq *BlockchainCryptocurrencyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blockchaincryptocurrency.Table,
			Columns: blockchaincryptocurrency.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: blockchaincryptocurrency.FieldID,
			},
		},
		From:   bcq.sql,
		Unique: true,
	}
	if unique := bcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blockchaincryptocurrency.FieldID)
		for i := range fields {
			if fields[i] != blockchaincryptocurrency.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bcq *BlockchainCryptocurrencyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bcq.driver.Dialect())
	t1 := builder.Table(blockchaincryptocurrency.Table)
	columns := bcq.fields
	if len(columns) == 0 {
		columns = blockchaincryptocurrency.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bcq.sql != nil {
		selector = bcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bcq.unique != nil && *bcq.unique {
		selector.Distinct()
	}
	for _, p := range bcq.predicates {
		p(selector)
	}
	for _, p := range bcq.order {
		p(selector)
	}
	if offset := bcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BlockchainCryptocurrencyGroupBy is the group-by builder for BlockchainCryptocurrency entities.
type BlockchainCryptocurrencyGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bcgb *BlockchainCryptocurrencyGroupBy) Aggregate(fns ...AggregateFunc) *BlockchainCryptocurrencyGroupBy {
	bcgb.fns = append(bcgb.fns, fns...)
	return bcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bcgb *BlockchainCryptocurrencyGroupBy) Scan(ctx context.Context, v any) error {
	query, err := bcgb.path(ctx)
	if err != nil {
		return err
	}
	bcgb.sql = query
	return bcgb.sqlScan(ctx, v)
}

func (bcgb *BlockchainCryptocurrencyGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range bcgb.fields {
		if !blockchaincryptocurrency.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bcgb *BlockchainCryptocurrencyGroupBy) sqlQuery() *sql.Selector {
	selector := bcgb.sql.Select()
	aggregation := make([]string, 0, len(bcgb.fns))
	for _, fn := range bcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bcgb.fields)+len(bcgb.fns))
		for _, f := range bcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bcgb.fields...)...)
}

// BlockchainCryptocurrencySelect is the builder for selecting fields of BlockchainCryptocurrency entities.
type BlockchainCryptocurrencySelect struct {
	*BlockchainCryptocurrencyQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bcs *BlockchainCryptocurrencySelect) Scan(ctx context.Context, v any) error {
	if err := bcs.prepareQuery(ctx); err != nil {
		return err
	}
	bcs.sql = bcs.BlockchainCryptocurrencyQuery.sqlQuery(ctx)
	return bcs.sqlScan(ctx, v)
}

func (bcs *BlockchainCryptocurrencySelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := bcs.sql.Query()
	if err := bcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}