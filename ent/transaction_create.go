// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/asset"
	"github.com/chenningg/hermitboard-api/ent/exchange"
	"github.com/chenningg/hermitboard-api/ent/portfolio"
	"github.com/chenningg/hermitboard-api/ent/transaction"
	"github.com/chenningg/hermitboard-api/ent/transactiontype"
	"github.com/chenningg/hermitboard-api/pulid"
)

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	mutation *TransactionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TransactionCreate) SetCreatedAt(t time.Time) *TransactionCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableCreatedAt(t *time.Time) *TransactionCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TransactionCreate) SetUpdatedAt(t time.Time) *TransactionCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableUpdatedAt(t *time.Time) *TransactionCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetDeletedAt sets the "deleted_at" field.
func (tc *TransactionCreate) SetDeletedAt(t time.Time) *TransactionCreate {
	tc.mutation.SetDeletedAt(t)
	return tc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableDeletedAt(t *time.Time) *TransactionCreate {
	if t != nil {
		tc.SetDeletedAt(*t)
	}
	return tc
}

// SetTime sets the "time" field.
func (tc *TransactionCreate) SetTime(t time.Time) *TransactionCreate {
	tc.mutation.SetTime(t)
	return tc
}

// SetUnits sets the "units" field.
func (tc *TransactionCreate) SetUnits(i int) *TransactionCreate {
	tc.mutation.SetUnits(i)
	return tc
}

// SetPricePerUnit sets the "price_per_unit" field.
func (tc *TransactionCreate) SetPricePerUnit(f float64) *TransactionCreate {
	tc.mutation.SetPricePerUnit(f)
	return tc
}

// SetTransactionTypeID sets the "transaction_type_id" field.
func (tc *TransactionCreate) SetTransactionTypeID(pu pulid.PULID) *TransactionCreate {
	tc.mutation.SetTransactionTypeID(pu)
	return tc
}

// SetBaseAssetID sets the "base_asset_id" field.
func (tc *TransactionCreate) SetBaseAssetID(pu pulid.PULID) *TransactionCreate {
	tc.mutation.SetBaseAssetID(pu)
	return tc
}

// SetQuoteAssetID sets the "quote_asset_id" field.
func (tc *TransactionCreate) SetQuoteAssetID(pu pulid.PULID) *TransactionCreate {
	tc.mutation.SetQuoteAssetID(pu)
	return tc
}

// SetNillableQuoteAssetID sets the "quote_asset_id" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableQuoteAssetID(pu *pulid.PULID) *TransactionCreate {
	if pu != nil {
		tc.SetQuoteAssetID(*pu)
	}
	return tc
}

// SetPortfolioID sets the "portfolio_id" field.
func (tc *TransactionCreate) SetPortfolioID(pu pulid.PULID) *TransactionCreate {
	tc.mutation.SetPortfolioID(pu)
	return tc
}

// SetExchangeID sets the "exchange_id" field.
func (tc *TransactionCreate) SetExchangeID(pu pulid.PULID) *TransactionCreate {
	tc.mutation.SetExchangeID(pu)
	return tc
}

// SetID sets the "id" field.
func (tc *TransactionCreate) SetID(pu pulid.PULID) *TransactionCreate {
	tc.mutation.SetID(pu)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableID(pu *pulid.PULID) *TransactionCreate {
	if pu != nil {
		tc.SetID(*pu)
	}
	return tc
}

// SetTransactionType sets the "transaction_type" edge to the TransactionType entity.
func (tc *TransactionCreate) SetTransactionType(t *TransactionType) *TransactionCreate {
	return tc.SetTransactionTypeID(t.ID)
}

// SetBaseAsset sets the "base_asset" edge to the Asset entity.
func (tc *TransactionCreate) SetBaseAsset(a *Asset) *TransactionCreate {
	return tc.SetBaseAssetID(a.ID)
}

// SetQuoteAsset sets the "quote_asset" edge to the Asset entity.
func (tc *TransactionCreate) SetQuoteAsset(a *Asset) *TransactionCreate {
	return tc.SetQuoteAssetID(a.ID)
}

// SetPortfolio sets the "portfolio" edge to the Portfolio entity.
func (tc *TransactionCreate) SetPortfolio(p *Portfolio) *TransactionCreate {
	return tc.SetPortfolioID(p.ID)
}

// SetExchange sets the "exchange" edge to the Exchange entity.
func (tc *TransactionCreate) SetExchange(e *Exchange) *TransactionCreate {
	return tc.SetExchangeID(e.ID)
}

// Mutation returns the TransactionMutation object of the builder.
func (tc *TransactionCreate) Mutation() *TransactionMutation {
	return tc.mutation
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	var (
		err  error
		node *Transaction
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Transaction)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TransactionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TransactionCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := transaction.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := transaction.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := transaction.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Transaction.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Transaction.updated_at"`)}
	}
	if _, ok := tc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "Transaction.time"`)}
	}
	if _, ok := tc.mutation.Units(); !ok {
		return &ValidationError{Name: "units", err: errors.New(`ent: missing required field "Transaction.units"`)}
	}
	if _, ok := tc.mutation.PricePerUnit(); !ok {
		return &ValidationError{Name: "price_per_unit", err: errors.New(`ent: missing required field "Transaction.price_per_unit"`)}
	}
	if _, ok := tc.mutation.TransactionTypeID(); !ok {
		return &ValidationError{Name: "transaction_type_id", err: errors.New(`ent: missing required field "Transaction.transaction_type_id"`)}
	}
	if _, ok := tc.mutation.BaseAssetID(); !ok {
		return &ValidationError{Name: "base_asset_id", err: errors.New(`ent: missing required field "Transaction.base_asset_id"`)}
	}
	if _, ok := tc.mutation.PortfolioID(); !ok {
		return &ValidationError{Name: "portfolio_id", err: errors.New(`ent: missing required field "Transaction.portfolio_id"`)}
	}
	if _, ok := tc.mutation.ExchangeID(); !ok {
		return &ValidationError{Name: "exchange_id", err: errors.New(`ent: missing required field "Transaction.exchange_id"`)}
	}
	if _, ok := tc.mutation.TransactionTypeID(); !ok {
		return &ValidationError{Name: "transaction_type", err: errors.New(`ent: missing required edge "Transaction.transaction_type"`)}
	}
	if _, ok := tc.mutation.BaseAssetID(); !ok {
		return &ValidationError{Name: "base_asset", err: errors.New(`ent: missing required edge "Transaction.base_asset"`)}
	}
	if _, ok := tc.mutation.PortfolioID(); !ok {
		return &ValidationError{Name: "portfolio", err: errors.New(`ent: missing required edge "Transaction.portfolio"`)}
	}
	if _, ok := tc.mutation.ExchangeID(); !ok {
		return &ValidationError{Name: "exchange", err: errors.New(`ent: missing required edge "Transaction.exchange"`)}
	}
	return nil
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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

func (tc *TransactionCreate) createSpec() (*Transaction, *sqlgraph.CreateSpec) {
	var (
		_node = &Transaction{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transaction.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transaction.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transaction.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transaction.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := tc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transaction.FieldTime,
		})
		_node.Time = value
	}
	if value, ok := tc.mutation.Units(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transaction.FieldUnits,
		})
		_node.Units = value
	}
	if value, ok := tc.mutation.PricePerUnit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: transaction.FieldPricePerUnit,
		})
		_node.PricePerUnit = value
	}
	if nodes := tc.mutation.TransactionTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transaction.TransactionTypeTable,
			Columns: []string{transaction.TransactionTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transactiontype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TransactionTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.BaseAssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transaction.BaseAssetTable,
			Columns: []string{transaction.BaseAssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BaseAssetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.QuoteAssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transaction.QuoteAssetTable,
			Columns: []string{transaction.QuoteAssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.QuoteAssetID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.PortfolioIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.PortfolioTable,
			Columns: []string{transaction.PortfolioColumn},
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
		_node.PortfolioID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ExchangeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.ExchangeTable,
			Columns: []string{transaction.ExchangeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: exchange.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ExchangeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TransactionCreateBulk is the builder for creating many Transaction entities in bulk.
type TransactionCreateBulk struct {
	config
	builders []*TransactionCreate
}

// Save creates the Transaction entities in the database.
func (tcb *TransactionCreateBulk) Save(ctx context.Context) ([]*Transaction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transaction, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionCreateBulk) SaveX(ctx context.Context) []*Transaction {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}