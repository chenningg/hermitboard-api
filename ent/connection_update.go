// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/connection"
	"github.com/chenningg/hermitboard-api/ent/portfolio"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// ConnectionUpdate is the builder for updating Connection entities.
type ConnectionUpdate struct {
	config
	hooks    []Hook
	mutation *ConnectionMutation
}

// Where appends a list predicates to the ConnectionUpdate builder.
func (cu *ConnectionUpdate) Where(ps ...predicate.Connection) *ConnectionUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ConnectionUpdate) SetUpdatedAt(t time.Time) *ConnectionUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *ConnectionUpdate) SetDeletedAt(t time.Time) *ConnectionUpdate {
	cu.mutation.SetDeletedAt(t)
	return cu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cu *ConnectionUpdate) ClearDeletedAt() *ConnectionUpdate {
	cu.mutation.ClearDeletedAt()
	return cu
}

// SetName sets the "name" field.
func (cu *ConnectionUpdate) SetName(s string) *ConnectionUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetAccessToken sets the "access_token" field.
func (cu *ConnectionUpdate) SetAccessToken(s string) *ConnectionUpdate {
	cu.mutation.SetAccessToken(s)
	return cu
}

// SetAccountID sets the "account_id" field.
func (cu *ConnectionUpdate) SetAccountID(pu pulid.PULID) *ConnectionUpdate {
	cu.mutation.SetAccountID(pu)
	return cu
}

// SetAccount sets the "account" edge to the Account entity.
func (cu *ConnectionUpdate) SetAccount(a *Account) *ConnectionUpdate {
	return cu.SetAccountID(a.ID)
}

// AddPortfolioIDs adds the "portfolios" edge to the Portfolio entity by IDs.
func (cu *ConnectionUpdate) AddPortfolioIDs(ids ...pulid.PULID) *ConnectionUpdate {
	cu.mutation.AddPortfolioIDs(ids...)
	return cu
}

// AddPortfolios adds the "portfolios" edges to the Portfolio entity.
func (cu *ConnectionUpdate) AddPortfolios(p ...*Portfolio) *ConnectionUpdate {
	ids := make([]pulid.PULID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddPortfolioIDs(ids...)
}

// Mutation returns the ConnectionMutation object of the builder.
func (cu *ConnectionUpdate) Mutation() *ConnectionMutation {
	return cu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (cu *ConnectionUpdate) ClearAccount() *ConnectionUpdate {
	cu.mutation.ClearAccount()
	return cu
}

// ClearPortfolios clears all "portfolios" edges to the Portfolio entity.
func (cu *ConnectionUpdate) ClearPortfolios() *ConnectionUpdate {
	cu.mutation.ClearPortfolios()
	return cu
}

// RemovePortfolioIDs removes the "portfolios" edge to Portfolio entities by IDs.
func (cu *ConnectionUpdate) RemovePortfolioIDs(ids ...pulid.PULID) *ConnectionUpdate {
	cu.mutation.RemovePortfolioIDs(ids...)
	return cu
}

// RemovePortfolios removes "portfolios" edges to Portfolio entities.
func (cu *ConnectionUpdate) RemovePortfolios(p ...*Portfolio) *ConnectionUpdate {
	ids := make([]pulid.PULID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemovePortfolioIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConnectionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConnectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConnectionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConnectionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConnectionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ConnectionUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := connection.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	if _, ok := cu.mutation.DeletedAt(); !ok && !cu.mutation.DeletedAtCleared() {
		v := connection.UpdateDefaultDeletedAt()
		cu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ConnectionUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := connection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Connection.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.AccessToken(); ok {
		if err := connection.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "Connection.access_token": %w`, err)}
		}
	}
	if _, ok := cu.mutation.AccountID(); cu.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Connection.account"`)
	}
	return nil
}

func (cu *ConnectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   connection.Table,
			Columns: connection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: connection.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: connection.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: connection.FieldDeletedAt,
		})
	}
	if cu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: connection.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: connection.FieldName,
		})
	}
	if value, ok := cu.mutation.AccessToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: connection.FieldAccessToken,
		})
	}
	if cu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connection.AccountTable,
			Columns: []string{connection.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connection.AccountTable,
			Columns: []string{connection.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.PortfoliosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connection.PortfoliosTable,
			Columns: connection.PortfoliosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: portfolio.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedPortfoliosIDs(); len(nodes) > 0 && !cu.mutation.PortfoliosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connection.PortfoliosTable,
			Columns: connection.PortfoliosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: portfolio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PortfoliosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connection.PortfoliosTable,
			Columns: connection.PortfoliosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: portfolio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ConnectionUpdateOne is the builder for updating a single Connection entity.
type ConnectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConnectionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ConnectionUpdateOne) SetUpdatedAt(t time.Time) *ConnectionUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *ConnectionUpdateOne) SetDeletedAt(t time.Time) *ConnectionUpdateOne {
	cuo.mutation.SetDeletedAt(t)
	return cuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cuo *ConnectionUpdateOne) ClearDeletedAt() *ConnectionUpdateOne {
	cuo.mutation.ClearDeletedAt()
	return cuo
}

// SetName sets the "name" field.
func (cuo *ConnectionUpdateOne) SetName(s string) *ConnectionUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetAccessToken sets the "access_token" field.
func (cuo *ConnectionUpdateOne) SetAccessToken(s string) *ConnectionUpdateOne {
	cuo.mutation.SetAccessToken(s)
	return cuo
}

// SetAccountID sets the "account_id" field.
func (cuo *ConnectionUpdateOne) SetAccountID(pu pulid.PULID) *ConnectionUpdateOne {
	cuo.mutation.SetAccountID(pu)
	return cuo
}

// SetAccount sets the "account" edge to the Account entity.
func (cuo *ConnectionUpdateOne) SetAccount(a *Account) *ConnectionUpdateOne {
	return cuo.SetAccountID(a.ID)
}

// AddPortfolioIDs adds the "portfolios" edge to the Portfolio entity by IDs.
func (cuo *ConnectionUpdateOne) AddPortfolioIDs(ids ...pulid.PULID) *ConnectionUpdateOne {
	cuo.mutation.AddPortfolioIDs(ids...)
	return cuo
}

// AddPortfolios adds the "portfolios" edges to the Portfolio entity.
func (cuo *ConnectionUpdateOne) AddPortfolios(p ...*Portfolio) *ConnectionUpdateOne {
	ids := make([]pulid.PULID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddPortfolioIDs(ids...)
}

// Mutation returns the ConnectionMutation object of the builder.
func (cuo *ConnectionUpdateOne) Mutation() *ConnectionMutation {
	return cuo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (cuo *ConnectionUpdateOne) ClearAccount() *ConnectionUpdateOne {
	cuo.mutation.ClearAccount()
	return cuo
}

// ClearPortfolios clears all "portfolios" edges to the Portfolio entity.
func (cuo *ConnectionUpdateOne) ClearPortfolios() *ConnectionUpdateOne {
	cuo.mutation.ClearPortfolios()
	return cuo
}

// RemovePortfolioIDs removes the "portfolios" edge to Portfolio entities by IDs.
func (cuo *ConnectionUpdateOne) RemovePortfolioIDs(ids ...pulid.PULID) *ConnectionUpdateOne {
	cuo.mutation.RemovePortfolioIDs(ids...)
	return cuo
}

// RemovePortfolios removes "portfolios" edges to Portfolio entities.
func (cuo *ConnectionUpdateOne) RemovePortfolios(p ...*Portfolio) *ConnectionUpdateOne {
	ids := make([]pulid.PULID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemovePortfolioIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConnectionUpdateOne) Select(field string, fields ...string) *ConnectionUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Connection entity.
func (cuo *ConnectionUpdateOne) Save(ctx context.Context) (*Connection, error) {
	var (
		err  error
		node *Connection
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConnectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Connection)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ConnectionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConnectionUpdateOne) SaveX(ctx context.Context) *Connection {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConnectionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConnectionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ConnectionUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := connection.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := cuo.mutation.DeletedAt(); !ok && !cuo.mutation.DeletedAtCleared() {
		v := connection.UpdateDefaultDeletedAt()
		cuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ConnectionUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := connection.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Connection.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.AccessToken(); ok {
		if err := connection.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "Connection.access_token": %w`, err)}
		}
	}
	if _, ok := cuo.mutation.AccountID(); cuo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Connection.account"`)
	}
	return nil
}

func (cuo *ConnectionUpdateOne) sqlSave(ctx context.Context) (_node *Connection, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   connection.Table,
			Columns: connection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: connection.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Connection.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, connection.FieldID)
		for _, f := range fields {
			if !connection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != connection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: connection.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: connection.FieldDeletedAt,
		})
	}
	if cuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: connection.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: connection.FieldName,
		})
	}
	if value, ok := cuo.mutation.AccessToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: connection.FieldAccessToken,
		})
	}
	if cuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connection.AccountTable,
			Columns: []string{connection.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connection.AccountTable,
			Columns: []string{connection.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.PortfoliosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connection.PortfoliosTable,
			Columns: connection.PortfoliosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: portfolio.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedPortfoliosIDs(); len(nodes) > 0 && !cuo.mutation.PortfoliosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connection.PortfoliosTable,
			Columns: connection.PortfoliosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: portfolio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PortfoliosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   connection.PortfoliosTable,
			Columns: connection.PortfoliosPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: portfolio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Connection{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}