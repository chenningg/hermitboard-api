// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/portfolio"
	"github.com/chenningg/hermitboard-api/ent/transaction"
	"github.com/chenningg/hermitboard-api/pulid"
)

// PortfolioCreate is the builder for creating a Portfolio entity.
type PortfolioCreate struct {
	config
	mutation *PortfolioMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PortfolioCreate) SetCreatedAt(t time.Time) *PortfolioCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PortfolioCreate) SetNillableCreatedAt(t *time.Time) *PortfolioCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PortfolioCreate) SetUpdatedAt(t time.Time) *PortfolioCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PortfolioCreate) SetNillableUpdatedAt(t *time.Time) *PortfolioCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *PortfolioCreate) SetDeletedAt(t time.Time) *PortfolioCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *PortfolioCreate) SetNillableDeletedAt(t *time.Time) *PortfolioCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PortfolioCreate) SetName(s string) *PortfolioCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetIsPublic sets the "is_public" field.
func (pc *PortfolioCreate) SetIsPublic(b bool) *PortfolioCreate {
	pc.mutation.SetIsPublic(b)
	return pc
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (pc *PortfolioCreate) SetNillableIsPublic(b *bool) *PortfolioCreate {
	if b != nil {
		pc.SetIsPublic(*b)
	}
	return pc
}

// SetIsVisible sets the "is_visible" field.
func (pc *PortfolioCreate) SetIsVisible(b bool) *PortfolioCreate {
	pc.mutation.SetIsVisible(b)
	return pc
}

// SetNillableIsVisible sets the "is_visible" field if the given value is not nil.
func (pc *PortfolioCreate) SetNillableIsVisible(b *bool) *PortfolioCreate {
	if b != nil {
		pc.SetIsVisible(*b)
	}
	return pc
}

// SetAccountID sets the "account_id" field.
func (pc *PortfolioCreate) SetAccountID(pu pulid.PULID) *PortfolioCreate {
	pc.mutation.SetAccountID(pu)
	return pc
}

// SetID sets the "id" field.
func (pc *PortfolioCreate) SetID(pu pulid.PULID) *PortfolioCreate {
	pc.mutation.SetID(pu)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PortfolioCreate) SetNillableID(pu *pulid.PULID) *PortfolioCreate {
	if pu != nil {
		pc.SetID(*pu)
	}
	return pc
}

// SetAccount sets the "account" edge to the Account entity.
func (pc *PortfolioCreate) SetAccount(a *Account) *PortfolioCreate {
	return pc.SetAccountID(a.ID)
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (pc *PortfolioCreate) AddTransactionIDs(ids ...pulid.PULID) *PortfolioCreate {
	pc.mutation.AddTransactionIDs(ids...)
	return pc
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (pc *PortfolioCreate) AddTransactions(t ...*Transaction) *PortfolioCreate {
	ids := make([]pulid.PULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTransactionIDs(ids...)
}

// Mutation returns the PortfolioMutation object of the builder.
func (pc *PortfolioCreate) Mutation() *PortfolioMutation {
	return pc.mutation
}

// Save creates the Portfolio in the database.
func (pc *PortfolioCreate) Save(ctx context.Context) (*Portfolio, error) {
	var (
		err  error
		node *Portfolio
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PortfolioMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Portfolio)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PortfolioMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PortfolioCreate) SaveX(ctx context.Context) *Portfolio {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PortfolioCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PortfolioCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PortfolioCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := portfolio.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := portfolio.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.IsPublic(); !ok {
		v := portfolio.DefaultIsPublic
		pc.mutation.SetIsPublic(v)
	}
	if _, ok := pc.mutation.IsVisible(); !ok {
		v := portfolio.DefaultIsVisible
		pc.mutation.SetIsVisible(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := portfolio.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PortfolioCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Portfolio.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Portfolio.updated_at"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Portfolio.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := portfolio.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Portfolio.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.IsPublic(); !ok {
		return &ValidationError{Name: "is_public", err: errors.New(`ent: missing required field "Portfolio.is_public"`)}
	}
	if _, ok := pc.mutation.IsVisible(); !ok {
		return &ValidationError{Name: "is_visible", err: errors.New(`ent: missing required field "Portfolio.is_visible"`)}
	}
	if _, ok := pc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`ent: missing required field "Portfolio.account_id"`)}
	}
	if _, ok := pc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required edge "Portfolio.account"`)}
	}
	return nil
}

func (pc *PortfolioCreate) sqlSave(ctx context.Context) (*Portfolio, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*pulid.PULID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (pc *PortfolioCreate) createSpec() (*Portfolio, *sqlgraph.CreateSpec) {
	var (
		_node = &Portfolio{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: portfolio.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: portfolio.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: portfolio.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: portfolio.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: portfolio.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portfolio.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.IsPublic(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: portfolio.FieldIsPublic,
		})
		_node.IsPublic = value
	}
	if value, ok := pc.mutation.IsVisible(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: portfolio.FieldIsVisible,
		})
		_node.IsVisible = value
	}
	if nodes := pc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   portfolio.AccountTable,
			Columns: []string{portfolio.AccountColumn},
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
		_node.AccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   portfolio.TransactionsTable,
			Columns: []string{portfolio.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PortfolioCreateBulk is the builder for creating many Portfolio entities in bulk.
type PortfolioCreateBulk struct {
	config
	builders []*PortfolioCreate
}

// Save creates the Portfolio entities in the database.
func (pcb *PortfolioCreateBulk) Save(ctx context.Context) ([]*Portfolio, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Portfolio, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PortfolioMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PortfolioCreateBulk) SaveX(ctx context.Context) []*Portfolio {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PortfolioCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PortfolioCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
