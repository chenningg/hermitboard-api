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
	"github.com/chenningg/hermitboard-api/ent/accountauthrole"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"github.com/chenningg/hermitboard-api/ent/portfolio"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// AccountQuery is the builder for querying Account entities.
type AccountQuery struct {
	config
	limit                     *int
	offset                    *int
	unique                    *bool
	order                     []OrderFunc
	fields                    []string
	predicates                []predicate.Account
	withAuthRoles             *AuthRoleQuery
	withPortfolios            *PortfolioQuery
	withAuthType              *AuthTypeQuery
	withAccountAuthRoles      *AccountAuthRoleQuery
	modifiers                 []func(*sql.Selector)
	loadTotal                 []func(context.Context, []*Account) error
	withNamedAuthRoles        map[string]*AuthRoleQuery
	withNamedPortfolios       map[string]*PortfolioQuery
	withNamedAccountAuthRoles map[string]*AccountAuthRoleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccountQuery builder.
func (aq *AccountQuery) Where(ps ...predicate.Account) *AccountQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AccountQuery) Limit(limit int) *AccountQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AccountQuery) Offset(offset int) *AccountQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AccountQuery) Unique(unique bool) *AccountQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AccountQuery) Order(o ...OrderFunc) *AccountQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryAuthRoles chains the current query on the "auth_roles" edge.
func (aq *AccountQuery) QueryAuthRoles() *AuthRoleQuery {
	query := &AuthRoleQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(authrole.Table, authrole.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, account.AuthRolesTable, account.AuthRolesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPortfolios chains the current query on the "portfolios" edge.
func (aq *AccountQuery) QueryPortfolios() *PortfolioQuery {
	query := &PortfolioQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(portfolio.Table, portfolio.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.PortfoliosTable, account.PortfoliosColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAuthType chains the current query on the "auth_type" edge.
func (aq *AccountQuery) QueryAuthType() *AuthTypeQuery {
	query := &AuthTypeQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(authtype.Table, authtype.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, account.AuthTypeTable, account.AuthTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAccountAuthRoles chains the current query on the "account_auth_roles" edge.
func (aq *AccountQuery) QueryAccountAuthRoles() *AccountAuthRoleQuery {
	query := &AccountAuthRoleQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, selector),
			sqlgraph.To(accountauthrole.Table, accountauthrole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, account.AccountAuthRolesTable, account.AccountAuthRolesColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Account entity from the query.
// Returns a *NotFoundError when no Account was found.
func (aq *AccountQuery) First(ctx context.Context) (*Account, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{account.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AccountQuery) FirstX(ctx context.Context) *Account {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Account ID from the query.
// Returns a *NotFoundError when no Account ID was found.
func (aq *AccountQuery) FirstID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{account.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AccountQuery) FirstIDX(ctx context.Context) pulid.PULID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Account entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Account entity is found.
// Returns a *NotFoundError when no Account entities are found.
func (aq *AccountQuery) Only(ctx context.Context) (*Account, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{account.Label}
	default:
		return nil, &NotSingularError{account.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AccountQuery) OnlyX(ctx context.Context) *Account {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Account ID in the query.
// Returns a *NotSingularError when more than one Account ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AccountQuery) OnlyID(ctx context.Context) (id pulid.PULID, err error) {
	var ids []pulid.PULID
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{account.Label}
	default:
		err = &NotSingularError{account.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AccountQuery) OnlyIDX(ctx context.Context) pulid.PULID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Accounts.
func (aq *AccountQuery) All(ctx context.Context) ([]*Account, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AccountQuery) AllX(ctx context.Context) []*Account {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Account IDs.
func (aq *AccountQuery) IDs(ctx context.Context) ([]pulid.PULID, error) {
	var ids []pulid.PULID
	if err := aq.Select(account.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AccountQuery) IDsX(ctx context.Context) []pulid.PULID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AccountQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AccountQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AccountQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AccountQuery) Clone() *AccountQuery {
	if aq == nil {
		return nil
	}
	return &AccountQuery{
		config:               aq.config,
		limit:                aq.limit,
		offset:               aq.offset,
		order:                append([]OrderFunc{}, aq.order...),
		predicates:           append([]predicate.Account{}, aq.predicates...),
		withAuthRoles:        aq.withAuthRoles.Clone(),
		withPortfolios:       aq.withPortfolios.Clone(),
		withAuthType:         aq.withAuthType.Clone(),
		withAccountAuthRoles: aq.withAccountAuthRoles.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithAuthRoles tells the query-builder to eager-load the nodes that are connected to
// the "auth_roles" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithAuthRoles(opts ...func(*AuthRoleQuery)) *AccountQuery {
	query := &AuthRoleQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withAuthRoles = query
	return aq
}

// WithPortfolios tells the query-builder to eager-load the nodes that are connected to
// the "portfolios" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithPortfolios(opts ...func(*PortfolioQuery)) *AccountQuery {
	query := &PortfolioQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withPortfolios = query
	return aq
}

// WithAuthType tells the query-builder to eager-load the nodes that are connected to
// the "auth_type" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithAuthType(opts ...func(*AuthTypeQuery)) *AccountQuery {
	query := &AuthTypeQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withAuthType = query
	return aq
}

// WithAccountAuthRoles tells the query-builder to eager-load the nodes that are connected to
// the "account_auth_roles" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithAccountAuthRoles(opts ...func(*AccountAuthRoleQuery)) *AccountQuery {
	query := &AccountAuthRoleQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withAccountAuthRoles = query
	return aq
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
//	client.Account.Query().
//		GroupBy(account.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AccountQuery) GroupBy(field string, fields ...string) *AccountGroupBy {
	grbuild := &AccountGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = account.Label
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
//	client.Account.Query().
//		Select(account.FieldCreatedAt).
//		Scan(ctx, &v)
func (aq *AccountQuery) Select(fields ...string) *AccountSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AccountSelect{AccountQuery: aq}
	selbuild.label = account.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

func (aq *AccountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !account.ValidColumn(f) {
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

func (aq *AccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Account, error) {
	var (
		nodes       = []*Account{}
		_spec       = aq.querySpec()
		loadedTypes = [4]bool{
			aq.withAuthRoles != nil,
			aq.withPortfolios != nil,
			aq.withAuthType != nil,
			aq.withAccountAuthRoles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Account).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Account{config: aq.config}
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
	if query := aq.withAuthRoles; query != nil {
		if err := aq.loadAuthRoles(ctx, query, nodes,
			func(n *Account) { n.Edges.AuthRoles = []*AuthRole{} },
			func(n *Account, e *AuthRole) { n.Edges.AuthRoles = append(n.Edges.AuthRoles, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withPortfolios; query != nil {
		if err := aq.loadPortfolios(ctx, query, nodes,
			func(n *Account) { n.Edges.Portfolios = []*Portfolio{} },
			func(n *Account, e *Portfolio) { n.Edges.Portfolios = append(n.Edges.Portfolios, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withAuthType; query != nil {
		if err := aq.loadAuthType(ctx, query, nodes, nil,
			func(n *Account, e *AuthType) { n.Edges.AuthType = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withAccountAuthRoles; query != nil {
		if err := aq.loadAccountAuthRoles(ctx, query, nodes,
			func(n *Account) { n.Edges.AccountAuthRoles = []*AccountAuthRole{} },
			func(n *Account, e *AccountAuthRole) { n.Edges.AccountAuthRoles = append(n.Edges.AccountAuthRoles, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedAuthRoles {
		if err := aq.loadAuthRoles(ctx, query, nodes,
			func(n *Account) { n.appendNamedAuthRoles(name) },
			func(n *Account, e *AuthRole) { n.appendNamedAuthRoles(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedPortfolios {
		if err := aq.loadPortfolios(ctx, query, nodes,
			func(n *Account) { n.appendNamedPortfolios(name) },
			func(n *Account, e *Portfolio) { n.appendNamedPortfolios(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedAccountAuthRoles {
		if err := aq.loadAccountAuthRoles(ctx, query, nodes,
			func(n *Account) { n.appendNamedAccountAuthRoles(name) },
			func(n *Account, e *AccountAuthRole) { n.appendNamedAccountAuthRoles(name, e) }); err != nil {
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

func (aq *AccountQuery) loadAuthRoles(ctx context.Context, query *AuthRoleQuery, nodes []*Account, init func(*Account), assign func(*Account, *AuthRole)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[pulid.PULID]*Account)
	nids := make(map[pulid.PULID]map[*Account]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(account.AuthRolesTable)
		s.Join(joinT).On(s.C(authrole.FieldID), joinT.C(account.AuthRolesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(account.AuthRolesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(account.AuthRolesPrimaryKey[0]))
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
				nids[inValue] = map[*Account]struct{}{byID[outValue]: struct{}{}}
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
func (aq *AccountQuery) loadPortfolios(ctx context.Context, query *PortfolioQuery, nodes []*Account, init func(*Account), assign func(*Account, *Portfolio)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pulid.PULID]*Account)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Portfolio(func(s *sql.Selector) {
		s.Where(sql.InValues(account.PortfoliosColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AccountID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AccountQuery) loadAuthType(ctx context.Context, query *AuthTypeQuery, nodes []*Account, init func(*Account), assign func(*Account, *AuthType)) error {
	ids := make([]pulid.PULID, 0, len(nodes))
	nodeids := make(map[pulid.PULID][]*Account)
	for i := range nodes {
		fk := nodes[i].AuthTypeID
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
			return fmt.Errorf(`unexpected foreign-key "auth_type_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AccountQuery) loadAccountAuthRoles(ctx context.Context, query *AccountAuthRoleQuery, nodes []*Account, init func(*Account), assign func(*Account, *AccountAuthRole)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[pulid.PULID]*Account)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.AccountAuthRole(func(s *sql.Selector) {
		s.Where(sql.InValues(account.AccountAuthRolesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AccountID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AccountQuery) sqlCount(ctx context.Context) (int, error) {
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

func (aq *AccountQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (aq *AccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: account.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for i := range fields {
			if fields[i] != account.FieldID {
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

func (aq *AccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(account.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = account.Columns
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

// WithNamedAuthRoles tells the query-builder to eager-load the nodes that are connected to the "auth_roles"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithNamedAuthRoles(name string, opts ...func(*AuthRoleQuery)) *AccountQuery {
	query := &AuthRoleQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedAuthRoles == nil {
		aq.withNamedAuthRoles = make(map[string]*AuthRoleQuery)
	}
	aq.withNamedAuthRoles[name] = query
	return aq
}

// WithNamedPortfolios tells the query-builder to eager-load the nodes that are connected to the "portfolios"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithNamedPortfolios(name string, opts ...func(*PortfolioQuery)) *AccountQuery {
	query := &PortfolioQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedPortfolios == nil {
		aq.withNamedPortfolios = make(map[string]*PortfolioQuery)
	}
	aq.withNamedPortfolios[name] = query
	return aq
}

// WithNamedAccountAuthRoles tells the query-builder to eager-load the nodes that are connected to the "account_auth_roles"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AccountQuery) WithNamedAccountAuthRoles(name string, opts ...func(*AccountAuthRoleQuery)) *AccountQuery {
	query := &AccountAuthRoleQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedAccountAuthRoles == nil {
		aq.withNamedAccountAuthRoles = make(map[string]*AccountAuthRoleQuery)
	}
	aq.withNamedAccountAuthRoles[name] = query
	return aq
}

// AccountGroupBy is the group-by builder for Account entities.
type AccountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AccountGroupBy) Aggregate(fns ...AggregateFunc) *AccountGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AccountGroupBy) Scan(ctx context.Context, v any) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AccountGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range agb.fields {
		if !account.ValidColumn(f) {
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

func (agb *AccountGroupBy) sqlQuery() *sql.Selector {
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

// AccountSelect is the builder for selecting fields of Account entities.
type AccountSelect struct {
	*AccountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *AccountSelect) Scan(ctx context.Context, v any) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AccountQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AccountSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
