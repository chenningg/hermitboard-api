// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/exchange"
	"github.com/chenningg/hermitboard-api/ent/transaction"
	"github.com/chenningg/hermitboard-api/pulid"
)

// ExchangeCreate is the builder for creating a Exchange entity.
type ExchangeCreate struct {
	config
	mutation *ExchangeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ec *ExchangeCreate) SetCreatedAt(t time.Time) *ExchangeCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *ExchangeCreate) SetNillableCreatedAt(t *time.Time) *ExchangeCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *ExchangeCreate) SetUpdatedAt(t time.Time) *ExchangeCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *ExchangeCreate) SetNillableUpdatedAt(t *time.Time) *ExchangeCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetDeletedAt sets the "deleted_at" field.
func (ec *ExchangeCreate) SetDeletedAt(t time.Time) *ExchangeCreate {
	ec.mutation.SetDeletedAt(t)
	return ec
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ec *ExchangeCreate) SetNillableDeletedAt(t *time.Time) *ExchangeCreate {
	if t != nil {
		ec.SetDeletedAt(*t)
	}
	return ec
}

// SetName sets the "name" field.
func (ec *ExchangeCreate) SetName(s string) *ExchangeCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetIcon sets the "icon" field.
func (ec *ExchangeCreate) SetIcon(s string) *ExchangeCreate {
	ec.mutation.SetIcon(s)
	return ec
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (ec *ExchangeCreate) SetNillableIcon(s *string) *ExchangeCreate {
	if s != nil {
		ec.SetIcon(*s)
	}
	return ec
}

// SetURL sets the "url" field.
func (ec *ExchangeCreate) SetURL(s string) *ExchangeCreate {
	ec.mutation.SetURL(s)
	return ec
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (ec *ExchangeCreate) SetNillableURL(s *string) *ExchangeCreate {
	if s != nil {
		ec.SetURL(*s)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *ExchangeCreate) SetID(pu pulid.PULID) *ExchangeCreate {
	ec.mutation.SetID(pu)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *ExchangeCreate) SetNillableID(pu *pulid.PULID) *ExchangeCreate {
	if pu != nil {
		ec.SetID(*pu)
	}
	return ec
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (ec *ExchangeCreate) AddTransactionIDs(ids ...pulid.PULID) *ExchangeCreate {
	ec.mutation.AddTransactionIDs(ids...)
	return ec
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (ec *ExchangeCreate) AddTransactions(t ...*Transaction) *ExchangeCreate {
	ids := make([]pulid.PULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ec.AddTransactionIDs(ids...)
}

// Mutation returns the ExchangeMutation object of the builder.
func (ec *ExchangeCreate) Mutation() *ExchangeMutation {
	return ec.mutation
}

// Save creates the Exchange in the database.
func (ec *ExchangeCreate) Save(ctx context.Context) (*Exchange, error) {
	var (
		err  error
		node *Exchange
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExchangeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Exchange)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ExchangeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *ExchangeCreate) SaveX(ctx context.Context) *Exchange {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *ExchangeCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *ExchangeCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *ExchangeCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := exchange.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		v := exchange.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		v := exchange.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *ExchangeCreate) check() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Exchange.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Exchange.updated_at"`)}
	}
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Exchange.name"`)}
	}
	if v, ok := ec.mutation.Name(); ok {
		if err := exchange.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Exchange.name": %w`, err)}
		}
	}
	if v, ok := ec.mutation.Icon(); ok {
		if err := exchange.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "Exchange.icon": %w`, err)}
		}
	}
	if v, ok := ec.mutation.URL(); ok {
		if err := exchange.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Exchange.url": %w`, err)}
		}
	}
	return nil
}

func (ec *ExchangeCreate) sqlSave(ctx context.Context) (*Exchange, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
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

func (ec *ExchangeCreate) createSpec() (*Exchange, *sqlgraph.CreateSpec) {
	var (
		_node = &Exchange{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: exchange.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: exchange.FieldID,
			},
		}
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: exchange.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: exchange.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: exchange.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := ec.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: exchange.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ec.mutation.Icon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: exchange.FieldIcon,
		})
		_node.Icon = &value
	}
	if value, ok := ec.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: exchange.FieldURL,
		})
		_node.URL = &value
	}
	if nodes := ec.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exchange.TransactionsTable,
			Columns: []string{exchange.TransactionsColumn},
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

// ExchangeCreateBulk is the builder for creating many Exchange entities in bulk.
type ExchangeCreateBulk struct {
	config
	builders []*ExchangeCreate
}

// Save creates the Exchange entities in the database.
func (ecb *ExchangeCreateBulk) Save(ctx context.Context) ([]*Exchange, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Exchange, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExchangeMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *ExchangeCreateBulk) SaveX(ctx context.Context) []*Exchange {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *ExchangeCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *ExchangeCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}