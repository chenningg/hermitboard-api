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
	"github.com/chenningg/hermitboard-api/ent/blockchain"
	"github.com/chenningg/hermitboard-api/ent/cryptocurrency"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/ent/transaction"
	"github.com/chenningg/hermitboard-api/pulid"
)

// BlockchainQuery is the builder for querying Blockchain entities.
type BlockchainQuery struct {
	config
	limit                     *int
	offset                    *int
	unique                    *bool
	order                     []OrderFunc
	fields                    []string
	predicates                []predicate.Blockchain
	withCryptocurrencies      *CryptocurrencyQuery
	withTransactions          *TransactionQuery
	modifiers                 []func(*sql.Selector)
	loadTotal                 []func(context.Context, []*Blockchain) error
	withNamedCryptocurrencies map[string]*CryptocurrencyQuery
	withNamedTransactions     map[string]*TransactionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlockchainQuery builder.
func (bq *BlockchainQuery) Where(ps ...predicate.Blockchain) *BlockchainQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BlockchainQuery) Limit(limit int) *BlockchainQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BlockchainQuery) Offset(offset int) *BlockchainQuery {
	bq.offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BlockchainQuery) Unique(unique bool) *BlockchainQuery {
	bq.unique = &unique
	return bq
}

// Order adds an order step to the query.
func (bq *BlockchainQuery) Order(o ...OrderFunc) *BlockchainQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryCryptocurrencies chains the current query on the "cryptocurrencies" edge.
func (bq *BlockchainQuery) QueryCryptocurrencies() *CryptocurrencyQuery {
	query := &CryptocurrencyQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blockchain.Table, blockchain.FieldID, selector),
			sqlgraph.To(cryptocurrency.Table, cryptocurrency.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, blockchain.CryptocurrenciesTable, blockchain.CryptocurrenciesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransactions chains the current query on the "transactions" edge.
func (bq *BlockchainQuery) QueryTransactions() *TransactionQuery {
	query := &TransactionQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blockchain.Table, blockchain.FieldID, selector),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, blockchain.TransactionsTable, blockchain.TransactionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Blockchain entity from the query.
// Returns a *NotFoundError when no Blockchain was found.
func (bq *BlockchainQuery) First(ctx context.Context) (*Blockchain, error) {
	nodes, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{blockchain.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BlockchainQuery) FirstX(ctx context.Context) *Blockchain {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Blockchain ID from the query.
// Returns a *NotFoundError when no Blockchain ID was found.
func (bq *BlockchainQuery) FirstID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blockchain.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BlockchainQuery) FirstIDX(ctx context.Context) pulid.PULID {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Blockchain entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Blockchain entity is found.
// Returns a *NotFoundError when no Blockchain entities are found.
func (bq *BlockchainQuery) Only(ctx context.Context) (*Blockchain, error) {
	nodes, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{blockchain.Label}
	default:
		return nil, &NotSingularError{blockchain.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BlockchainQuery) OnlyX(ctx context.Context) *Blockchain {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Blockchain ID in the query.
// Returns a *NotSingularError when more than one Blockchain ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BlockchainQuery) OnlyID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blockchain.Label}
	default:
		err = &NotSingularError{blockchain.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BlockchainQuery) OnlyIDX(ctx context.Context) pulid.PULID {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Blockchains.
func (bq *BlockchainQuery) All(ctx context.Context) ([]*Blockchain, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bq *BlockchainQuery) AllX(ctx context.Context) []*Blockchain {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Blockchain IDs.
func (bq *BlockchainQuery) IDs(ctx context.Context) ([]pulid.PULID, error) {
	var ids []pulid.PULID
	if err := bq.Select(blockchain.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BlockchainQuery) IDsX(ctx context.Context) []pulid.PULID {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BlockchainQuery) Count(ctx context.Context) (int, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BlockchainQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BlockchainQuery) Exist(ctx context.Context) (bool, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BlockchainQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlockchainQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BlockchainQuery) Clone() *BlockchainQuery {
	if bq == nil {
		return nil
	}
	return &BlockchainQuery{
		config:               bq.config,
		limit:                bq.limit,
		offset:               bq.offset,
		order:                append([]OrderFunc{}, bq.order...),
		predicates:           append([]predicate.Blockchain{}, bq.predicates...),
		withCryptocurrencies: bq.withCryptocurrencies.Clone(),
		withTransactions:     bq.withTransactions.Clone(),
		// clone intermediate query.
		sql:    bq.sql.Clone(),
		path:   bq.path,
		unique: bq.unique,
	}
}

// WithCryptocurrencies tells the query-builder to eager-load the nodes that are connected to
// the "cryptocurrencies" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BlockchainQuery) WithCryptocurrencies(opts ...func(*CryptocurrencyQuery)) *BlockchainQuery {
	query := &CryptocurrencyQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withCryptocurrencies = query
	return bq
}

// WithTransactions tells the query-builder to eager-load the nodes that are connected to
// the "transactions" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BlockchainQuery) WithTransactions(opts ...func(*TransactionQuery)) *BlockchainQuery {
	query := &TransactionQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withTransactions = query
	return bq
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
//	client.Blockchain.Query().
//		GroupBy(blockchain.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BlockchainQuery) GroupBy(field string, fields ...string) *BlockchainGroupBy {
	grbuild := &BlockchainGroupBy{config: bq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(ctx), nil
	}
	grbuild.label = blockchain.Label
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
//	client.Blockchain.Query().
//		Select(blockchain.FieldCreatedAt).
//		Scan(ctx, &v)
func (bq *BlockchainQuery) Select(fields ...string) *BlockchainSelect {
	bq.fields = append(bq.fields, fields...)
	selbuild := &BlockchainSelect{BlockchainQuery: bq}
	selbuild.label = blockchain.Label
	selbuild.flds, selbuild.scan = &bq.fields, selbuild.Scan
	return selbuild
}

func (bq *BlockchainQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bq.fields {
		if !blockchain.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BlockchainQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Blockchain, error) {
	var (
		nodes       = []*Blockchain{}
		_spec       = bq.querySpec()
		loadedTypes = [2]bool{
			bq.withCryptocurrencies != nil,
			bq.withTransactions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Blockchain).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Blockchain{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withCryptocurrencies; query != nil {
		if err := bq.loadCryptocurrencies(ctx, query, nodes,
			func(n *Blockchain) { n.Edges.Cryptocurrencies = []*Cryptocurrency{} },
			func(n *Blockchain, e *Cryptocurrency) { n.Edges.Cryptocurrencies = append(n.Edges.Cryptocurrencies, e) }); err != nil {
			return nil, err
		}
	}
	if query := bq.withTransactions; query != nil {
		if err := bq.loadTransactions(ctx, query, nodes,
			func(n *Blockchain) { n.Edges.Transactions = []*Transaction{} },
			func(n *Blockchain, e *Transaction) { n.Edges.Transactions = append(n.Edges.Transactions, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bq.withNamedCryptocurrencies {
		if err := bq.loadCryptocurrencies(ctx, query, nodes,
			func(n *Blockchain) { n.appendNamedCryptocurrencies(name) },
			func(n *Blockchain, e *Cryptocurrency) { n.appendNamedCryptocurrencies(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bq.withNamedTransactions {
		if err := bq.loadTransactions(ctx, query, nodes,
			func(n *Blockchain) { n.appendNamedTransactions(name) },
			func(n *Blockchain, e *Transaction) { n.appendNamedTransactions(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range bq.loadTotal {
		if err := bq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BlockchainQuery) loadCryptocurrencies(ctx context.Context, query *CryptocurrencyQuery, nodes []*Blockchain, init func(*Blockchain), assign func(*Blockchain, *Cryptocurrency)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[pulid.PULID]*Blockchain)
	nids := make(map[pulid.PULID]map[*Blockchain]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(blockchain.CryptocurrenciesTable)
		s.Join(joinT).On(s.C(cryptocurrency.FieldID), joinT.C(blockchain.CryptocurrenciesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(blockchain.CryptocurrenciesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(blockchain.CryptocurrenciesPrimaryKey[0]))
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
				nids[inValue] = map[*Blockchain]struct{}{byID[outValue]: struct{}{}}
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
			return fmt.Errorf(`unexpected "cryptocurrencies" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (bq *BlockchainQuery) loadTransactions(ctx context.Context, query *TransactionQuery, nodes []*Blockchain, init func(*Blockchain), assign func(*Blockchain, *Transaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pulid.PULID]*Blockchain)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.InValues(blockchain.TransactionsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.transaction_blockchain
		if fk == nil {
			return fmt.Errorf(`foreign-key "transaction_blockchain" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "transaction_blockchain" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BlockchainQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	_spec.Node.Columns = bq.fields
	if len(bq.fields) > 0 {
		_spec.Unique = bq.unique != nil && *bq.unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BlockchainQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (bq *BlockchainQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blockchain.Table,
			Columns: blockchain.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: blockchain.FieldID,
			},
		},
		From:   bq.sql,
		Unique: true,
	}
	if unique := bq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blockchain.FieldID)
		for i := range fields {
			if fields[i] != blockchain.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BlockchainQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(blockchain.Table)
	columns := bq.fields
	if len(columns) == 0 {
		columns = blockchain.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.unique != nil && *bq.unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedCryptocurrencies tells the query-builder to eager-load the nodes that are connected to the "cryptocurrencies"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bq *BlockchainQuery) WithNamedCryptocurrencies(name string, opts ...func(*CryptocurrencyQuery)) *BlockchainQuery {
	query := &CryptocurrencyQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	if bq.withNamedCryptocurrencies == nil {
		bq.withNamedCryptocurrencies = make(map[string]*CryptocurrencyQuery)
	}
	bq.withNamedCryptocurrencies[name] = query
	return bq
}

// WithNamedTransactions tells the query-builder to eager-load the nodes that are connected to the "transactions"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bq *BlockchainQuery) WithNamedTransactions(name string, opts ...func(*TransactionQuery)) *BlockchainQuery {
	query := &TransactionQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	if bq.withNamedTransactions == nil {
		bq.withNamedTransactions = make(map[string]*TransactionQuery)
	}
	bq.withNamedTransactions[name] = query
	return bq
}

// BlockchainGroupBy is the group-by builder for Blockchain entities.
type BlockchainGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BlockchainGroupBy) Aggregate(fns ...AggregateFunc) *BlockchainGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bgb *BlockchainGroupBy) Scan(ctx context.Context, v any) error {
	query, err := bgb.path(ctx)
	if err != nil {
		return err
	}
	bgb.sql = query
	return bgb.sqlScan(ctx, v)
}

func (bgb *BlockchainGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range bgb.fields {
		if !blockchain.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BlockchainGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql.Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
		for _, f := range bgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bgb.fields...)...)
}

// BlockchainSelect is the builder for selecting fields of Blockchain entities.
type BlockchainSelect struct {
	*BlockchainQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BlockchainSelect) Scan(ctx context.Context, v any) error {
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	bs.sql = bs.BlockchainQuery.sqlQuery(ctx)
	return bs.sqlScan(ctx, v)
}

func (bs *BlockchainSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := bs.sql.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
