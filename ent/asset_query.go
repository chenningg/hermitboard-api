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
	"github.com/chenningg/hermitboard-api/ent/asset"
	"github.com/chenningg/hermitboard-api/ent/assetclass"
	"github.com/chenningg/hermitboard-api/ent/cryptocurrency"
	"github.com/chenningg/hermitboard-api/ent/dailyassetprice"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// AssetQuery is the builder for querying Asset entities.
type AssetQuery struct {
	config
	limit                     *int
	offset                    *int
	unique                    *bool
	order                     []OrderFunc
	fields                    []string
	predicates                []predicate.Asset
	withAssetClass            *AssetClassQuery
	withCryptocurrency        *CryptocurrencyQuery
	withDailyAssetPrices      *DailyAssetPriceQuery
	withFKs                   bool
	modifiers                 []func(*sql.Selector)
	loadTotal                 []func(context.Context, []*Asset) error
	withNamedDailyAssetPrices map[string]*DailyAssetPriceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AssetQuery builder.
func (aq *AssetQuery) Where(ps ...predicate.Asset) *AssetQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AssetQuery) Limit(limit int) *AssetQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AssetQuery) Offset(offset int) *AssetQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AssetQuery) Unique(unique bool) *AssetQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AssetQuery) Order(o ...OrderFunc) *AssetQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryAssetClass chains the current query on the "asset_class" edge.
func (aq *AssetQuery) QueryAssetClass() *AssetClassQuery {
	query := &AssetClassQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(assetclass.Table, assetclass.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, asset.AssetClassTable, asset.AssetClassColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCryptocurrency chains the current query on the "cryptocurrency" edge.
func (aq *AssetQuery) QueryCryptocurrency() *CryptocurrencyQuery {
	query := &CryptocurrencyQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(cryptocurrency.Table, cryptocurrency.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, asset.CryptocurrencyTable, asset.CryptocurrencyColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDailyAssetPrices chains the current query on the "daily_asset_prices" edge.
func (aq *AssetQuery) QueryDailyAssetPrices() *DailyAssetPriceQuery {
	query := &DailyAssetPriceQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(dailyassetprice.Table, dailyassetprice.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, asset.DailyAssetPricesTable, asset.DailyAssetPricesColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Asset entity from the query.
// Returns a *NotFoundError when no Asset was found.
func (aq *AssetQuery) First(ctx context.Context) (*Asset, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asset.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AssetQuery) FirstX(ctx context.Context) *Asset {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Asset ID from the query.
// Returns a *NotFoundError when no Asset ID was found.
func (aq *AssetQuery) FirstID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asset.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AssetQuery) FirstIDX(ctx context.Context) pulid.PULID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Asset entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Asset entity is found.
// Returns a *NotFoundError when no Asset entities are found.
func (aq *AssetQuery) Only(ctx context.Context) (*Asset, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asset.Label}
	default:
		return nil, &NotSingularError{asset.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AssetQuery) OnlyX(ctx context.Context) *Asset {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Asset ID in the query.
// Returns a *NotSingularError when more than one Asset ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AssetQuery) OnlyID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = &NotSingularError{asset.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AssetQuery) OnlyIDX(ctx context.Context) pulid.PULID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Assets.
func (aq *AssetQuery) All(ctx context.Context) ([]*Asset, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AssetQuery) AllX(ctx context.Context) []*Asset {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Asset IDs.
func (aq *AssetQuery) IDs(ctx context.Context) ([]pulid.PULID, error) {
	var ids []pulid.PULID
	if err := aq.Select(asset.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AssetQuery) IDsX(ctx context.Context) []pulid.PULID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AssetQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AssetQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AssetQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AssetQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AssetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AssetQuery) Clone() *AssetQuery {
	if aq == nil {
		return nil
	}
	return &AssetQuery{
		config:               aq.config,
		limit:                aq.limit,
		offset:               aq.offset,
		order:                append([]OrderFunc{}, aq.order...),
		predicates:           append([]predicate.Asset{}, aq.predicates...),
		withAssetClass:       aq.withAssetClass.Clone(),
		withCryptocurrency:   aq.withCryptocurrency.Clone(),
		withDailyAssetPrices: aq.withDailyAssetPrices.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithAssetClass tells the query-builder to eager-load the nodes that are connected to
// the "asset_class" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithAssetClass(opts ...func(*AssetClassQuery)) *AssetQuery {
	query := &AssetClassQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withAssetClass = query
	return aq
}

// WithCryptocurrency tells the query-builder to eager-load the nodes that are connected to
// the "cryptocurrency" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithCryptocurrency(opts ...func(*CryptocurrencyQuery)) *AssetQuery {
	query := &CryptocurrencyQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withCryptocurrency = query
	return aq
}

// WithDailyAssetPrices tells the query-builder to eager-load the nodes that are connected to
// the "daily_asset_prices" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithDailyAssetPrices(opts ...func(*DailyAssetPriceQuery)) *AssetQuery {
	query := &DailyAssetPriceQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withDailyAssetPrices = query
	return aq
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
//	client.Asset.Query().
//		GroupBy(asset.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AssetQuery) GroupBy(field string, fields ...string) *AssetGroupBy {
	grbuild := &AssetGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = asset.Label
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
//	client.Asset.Query().
//		Select(asset.FieldCreatedAt).
//		Scan(ctx, &v)
func (aq *AssetQuery) Select(fields ...string) *AssetSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AssetSelect{AssetQuery: aq}
	selbuild.label = asset.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

func (aq *AssetQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !asset.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AssetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Asset, error) {
	var (
		nodes       = []*Asset{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [3]bool{
			aq.withAssetClass != nil,
			aq.withCryptocurrency != nil,
			aq.withDailyAssetPrices != nil,
		}
	)
	if aq.withAssetClass != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, asset.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Asset).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Asset{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withAssetClass; query != nil {
		if err := aq.loadAssetClass(ctx, query, nodes, nil,
			func(n *Asset, e *AssetClass) { n.Edges.AssetClass = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withCryptocurrency; query != nil {
		if err := aq.loadCryptocurrency(ctx, query, nodes, nil,
			func(n *Asset, e *Cryptocurrency) { n.Edges.Cryptocurrency = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withDailyAssetPrices; query != nil {
		if err := aq.loadDailyAssetPrices(ctx, query, nodes,
			func(n *Asset) { n.Edges.DailyAssetPrices = []*DailyAssetPrice{} },
			func(n *Asset, e *DailyAssetPrice) { n.Edges.DailyAssetPrices = append(n.Edges.DailyAssetPrices, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedDailyAssetPrices {
		if err := aq.loadDailyAssetPrices(ctx, query, nodes,
			func(n *Asset) { n.appendNamedDailyAssetPrices(name) },
			func(n *Asset, e *DailyAssetPrice) { n.appendNamedDailyAssetPrices(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range aq.loadTotal {
		if err := aq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AssetQuery) loadAssetClass(ctx context.Context, query *AssetClassQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *AssetClass)) error {
	ids := make([]pulid.PULID, 0, len(nodes))
	nodeids := make(map[pulid.PULID][]*Asset)
	for i := range nodes {
		if nodes[i].asset_asset_class == nil {
			continue
		}
		fk := *nodes[i].asset_asset_class
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(assetclass.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "asset_asset_class" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AssetQuery) loadCryptocurrency(ctx context.Context, query *CryptocurrencyQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *Cryptocurrency)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pulid.PULID]*Asset)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.Where(predicate.Cryptocurrency(func(s *sql.Selector) {
		s.Where(sql.InValues(asset.CryptocurrencyColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AssetID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "asset_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AssetQuery) loadDailyAssetPrices(ctx context.Context, query *DailyAssetPriceQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *DailyAssetPrice)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pulid.PULID]*Asset)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.DailyAssetPrice(func(s *sql.Selector) {
		s.Where(sql.InValues(asset.DailyAssetPricesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AssetID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "asset_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AssetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AssetQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (aq *AssetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asset.Table,
			Columns: asset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: asset.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asset.FieldID)
		for i := range fields {
			if fields[i] != asset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AssetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(asset.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = asset.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedDailyAssetPrices tells the query-builder to eager-load the nodes that are connected to the "daily_asset_prices"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithNamedDailyAssetPrices(name string, opts ...func(*DailyAssetPriceQuery)) *AssetQuery {
	query := &DailyAssetPriceQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedDailyAssetPrices == nil {
		aq.withNamedDailyAssetPrices = make(map[string]*DailyAssetPriceQuery)
	}
	aq.withNamedDailyAssetPrices[name] = query
	return aq
}

// AssetGroupBy is the group-by builder for Asset entities.
type AssetGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AssetGroupBy) Aggregate(fns ...AggregateFunc) *AssetGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AssetGroupBy) Scan(ctx context.Context, v any) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AssetGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range agb.fields {
		if !asset.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AssetGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AssetSelect is the builder for selecting fields of Asset entities.
type AssetSelect struct {
	*AssetQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *AssetSelect) Scan(ctx context.Context, v any) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AssetQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AssetSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
