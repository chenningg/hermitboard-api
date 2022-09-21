// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/asset"
	"github.com/chenningg/hermitboard-api/ent/dailyassetprice"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// DailyAssetPriceQuery is the builder for querying DailyAssetPrice entities.
type DailyAssetPriceQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.DailyAssetPrice
	withBaseAsset *AssetQuery
	modifiers     []func(*sql.Selector)
	loadTotal     []func(context.Context, []*DailyAssetPrice) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DailyAssetPriceQuery builder.
func (dapq *DailyAssetPriceQuery) Where(ps ...predicate.DailyAssetPrice) *DailyAssetPriceQuery {
	dapq.predicates = append(dapq.predicates, ps...)
	return dapq
}

// Limit adds a limit step to the query.
func (dapq *DailyAssetPriceQuery) Limit(limit int) *DailyAssetPriceQuery {
	dapq.limit = &limit
	return dapq
}

// Offset adds an offset step to the query.
func (dapq *DailyAssetPriceQuery) Offset(offset int) *DailyAssetPriceQuery {
	dapq.offset = &offset
	return dapq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dapq *DailyAssetPriceQuery) Unique(unique bool) *DailyAssetPriceQuery {
	dapq.unique = &unique
	return dapq
}

// Order adds an order step to the query.
func (dapq *DailyAssetPriceQuery) Order(o ...OrderFunc) *DailyAssetPriceQuery {
	dapq.order = append(dapq.order, o...)
	return dapq
}

// QueryBaseAsset chains the current query on the "base_asset" edge.
func (dapq *DailyAssetPriceQuery) QueryBaseAsset() *AssetQuery {
	query := &AssetQuery{config: dapq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dapq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dapq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dailyassetprice.Table, dailyassetprice.FieldID, selector),
			sqlgraph.To(asset.Table, asset.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dailyassetprice.BaseAssetTable, dailyassetprice.BaseAssetColumn),
		)
		fromU = sqlgraph.SetNeighbors(dapq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DailyAssetPrice entity from the query.
// Returns a *NotFoundError when no DailyAssetPrice was found.
func (dapq *DailyAssetPriceQuery) First(ctx context.Context) (*DailyAssetPrice, error) {
	nodes, err := dapq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dailyassetprice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dapq *DailyAssetPriceQuery) FirstX(ctx context.Context) *DailyAssetPrice {
	node, err := dapq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DailyAssetPrice ID from the query.
// Returns a *NotFoundError when no DailyAssetPrice ID was found.
func (dapq *DailyAssetPriceQuery) FirstID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = dapq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dailyassetprice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dapq *DailyAssetPriceQuery) FirstIDX(ctx context.Context) pulid.PULID {
	id, err := dapq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DailyAssetPrice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DailyAssetPrice entity is found.
// Returns a *NotFoundError when no DailyAssetPrice entities are found.
func (dapq *DailyAssetPriceQuery) Only(ctx context.Context) (*DailyAssetPrice, error) {
	nodes, err := dapq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dailyassetprice.Label}
	default:
		return nil, &NotSingularError{dailyassetprice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dapq *DailyAssetPriceQuery) OnlyX(ctx context.Context) *DailyAssetPrice {
	node, err := dapq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DailyAssetPrice ID in the query.
// Returns a *NotSingularError when more than one DailyAssetPrice ID is found.
// Returns a *NotFoundError when no entities are found.
func (dapq *DailyAssetPriceQuery) OnlyID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = dapq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dailyassetprice.Label}
	default:
		err = &NotSingularError{dailyassetprice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dapq *DailyAssetPriceQuery) OnlyIDX(ctx context.Context) pulid.PULID {
	id, err := dapq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DailyAssetPrices.
func (dapq *DailyAssetPriceQuery) All(ctx context.Context) ([]*DailyAssetPrice, error) {
	if err := dapq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dapq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dapq *DailyAssetPriceQuery) AllX(ctx context.Context) []*DailyAssetPrice {
	nodes, err := dapq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DailyAssetPrice IDs.
func (dapq *DailyAssetPriceQuery) IDs(ctx context.Context) ([]pulid.PULID, error) {
	var ids []pulid.PULID
	if err := dapq.Select(dailyassetprice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dapq *DailyAssetPriceQuery) IDsX(ctx context.Context) []pulid.PULID {
	ids, err := dapq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dapq *DailyAssetPriceQuery) Count(ctx context.Context) (int, error) {
	if err := dapq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dapq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dapq *DailyAssetPriceQuery) CountX(ctx context.Context) int {
	count, err := dapq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dapq *DailyAssetPriceQuery) Exist(ctx context.Context) (bool, error) {
	if err := dapq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dapq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dapq *DailyAssetPriceQuery) ExistX(ctx context.Context) bool {
	exist, err := dapq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DailyAssetPriceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dapq *DailyAssetPriceQuery) Clone() *DailyAssetPriceQuery {
	if dapq == nil {
		return nil
	}
	return &DailyAssetPriceQuery{
		config:        dapq.config,
		limit:         dapq.limit,
		offset:        dapq.offset,
		order:         append([]OrderFunc{}, dapq.order...),
		predicates:    append([]predicate.DailyAssetPrice{}, dapq.predicates...),
		withBaseAsset: dapq.withBaseAsset.Clone(),
		// clone intermediate query.
		sql:    dapq.sql.Clone(),
		path:   dapq.path,
		unique: dapq.unique,
	}
}

// WithBaseAsset tells the query-builder to eager-load the nodes that are connected to
// the "base_asset" edge. The optional arguments are used to configure the query builder of the edge.
func (dapq *DailyAssetPriceQuery) WithBaseAsset(opts ...func(*AssetQuery)) *DailyAssetPriceQuery {
	query := &AssetQuery{config: dapq.config}
	for _, opt := range opts {
		opt(query)
	}
	dapq.withBaseAsset = query
	return dapq
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
//	client.DailyAssetPrice.Query().
//		GroupBy(dailyassetprice.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dapq *DailyAssetPriceQuery) GroupBy(field string, fields ...string) *DailyAssetPriceGroupBy {
	grbuild := &DailyAssetPriceGroupBy{config: dapq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dapq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dapq.sqlQuery(ctx), nil
	}
	grbuild.label = dailyassetprice.Label
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
//	client.DailyAssetPrice.Query().
//		Select(dailyassetprice.FieldCreatedAt).
//		Scan(ctx, &v)
func (dapq *DailyAssetPriceQuery) Select(fields ...string) *DailyAssetPriceSelect {
	dapq.fields = append(dapq.fields, fields...)
	selbuild := &DailyAssetPriceSelect{DailyAssetPriceQuery: dapq}
	selbuild.label = dailyassetprice.Label
	selbuild.flds, selbuild.scan = &dapq.fields, selbuild.Scan
	return selbuild
}

func (dapq *DailyAssetPriceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dapq.fields {
		if !dailyassetprice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dapq.path != nil {
		prev, err := dapq.path(ctx)
		if err != nil {
			return err
		}
		dapq.sql = prev
	}
	return nil
}

func (dapq *DailyAssetPriceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DailyAssetPrice, error) {
	var (
		nodes       = []*DailyAssetPrice{}
		_spec       = dapq.querySpec()
		loadedTypes = [1]bool{
			dapq.withBaseAsset != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DailyAssetPrice).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DailyAssetPrice{config: dapq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dapq.modifiers) > 0 {
		_spec.Modifiers = dapq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dapq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dapq.withBaseAsset; query != nil {
		if err := dapq.loadBaseAsset(ctx, query, nodes, nil,
			func(n *DailyAssetPrice, e *Asset) { n.Edges.BaseAsset = e }); err != nil {
			return nil, err
		}
	}
	for i := range dapq.loadTotal {
		if err := dapq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dapq *DailyAssetPriceQuery) loadBaseAsset(ctx context.Context, query *AssetQuery, nodes []*DailyAssetPrice, init func(*DailyAssetPrice), assign func(*DailyAssetPrice, *Asset)) error {
	ids := make([]pulid.PULID, 0, len(nodes))
	nodeids := make(map[pulid.PULID][]*DailyAssetPrice)
	for i := range nodes {
		fk := nodes[i].BaseAssetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(asset.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "base_asset_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dapq *DailyAssetPriceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dapq.querySpec()
	if len(dapq.modifiers) > 0 {
		_spec.Modifiers = dapq.modifiers
	}
	_spec.Node.Columns = dapq.fields
	if len(dapq.fields) > 0 {
		_spec.Unique = dapq.unique != nil && *dapq.unique
	}
	return sqlgraph.CountNodes(ctx, dapq.driver, _spec)
}

func (dapq *DailyAssetPriceQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := dapq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (dapq *DailyAssetPriceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dailyassetprice.Table,
			Columns: dailyassetprice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: dailyassetprice.FieldID,
			},
		},
		From:   dapq.sql,
		Unique: true,
	}
	if unique := dapq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dapq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dailyassetprice.FieldID)
		for i := range fields {
			if fields[i] != dailyassetprice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dapq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dapq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dapq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dapq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dapq *DailyAssetPriceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dapq.driver.Dialect())
	t1 := builder.Table(dailyassetprice.Table)
	columns := dapq.fields
	if len(columns) == 0 {
		columns = dailyassetprice.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dapq.sql != nil {
		selector = dapq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dapq.unique != nil && *dapq.unique {
		selector.Distinct()
	}
	for _, p := range dapq.predicates {
		p(selector)
	}
	for _, p := range dapq.order {
		p(selector)
	}
	if offset := dapq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dapq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DailyAssetPriceGroupBy is the group-by builder for DailyAssetPrice entities.
type DailyAssetPriceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dapgb *DailyAssetPriceGroupBy) Aggregate(fns ...AggregateFunc) *DailyAssetPriceGroupBy {
	dapgb.fns = append(dapgb.fns, fns...)
	return dapgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dapgb *DailyAssetPriceGroupBy) Scan(ctx context.Context, v any) error {
	query, err := dapgb.path(ctx)
	if err != nil {
		return err
	}
	dapgb.sql = query
	return dapgb.sqlScan(ctx, v)
}

func (dapgb *DailyAssetPriceGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range dapgb.fields {
		if !dailyassetprice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dapgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dapgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dapgb *DailyAssetPriceGroupBy) sqlQuery() *sql.Selector {
	selector := dapgb.sql.Select()
	aggregation := make([]string, 0, len(dapgb.fns))
	for _, fn := range dapgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dapgb.fields)+len(dapgb.fns))
		for _, f := range dapgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dapgb.fields...)...)
}

// DailyAssetPriceSelect is the builder for selecting fields of DailyAssetPrice entities.
type DailyAssetPriceSelect struct {
	*DailyAssetPriceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (daps *DailyAssetPriceSelect) Scan(ctx context.Context, v any) error {
	if err := daps.prepareQuery(ctx); err != nil {
		return err
	}
	daps.sql = daps.DailyAssetPriceQuery.sqlQuery(ctx)
	return daps.sqlScan(ctx, v)
}

func (daps *DailyAssetPriceSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := daps.sql.Query()
	if err := daps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
