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
	"github.com/chenningg/hermitboard-api/ent/assetclass"
	"github.com/chenningg/hermitboard-api/ent/cryptocurrency"
	"github.com/chenningg/hermitboard-api/ent/dailyassetprice"
	"github.com/chenningg/hermitboard-api/ent/transaction"
	"github.com/oklog/ulid/v2"
	ulid "github.com/oklog/ulid/v2"
)

// AssetCreate is the builder for creating a Asset entity.
type AssetCreate struct {
	config
	mutation *AssetMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AssetCreate) SetCreatedAt(t time.Time) *AssetCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AssetCreate) SetNillableCreatedAt(t *time.Time) *AssetCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AssetCreate) SetUpdatedAt(t time.Time) *AssetCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AssetCreate) SetNillableUpdatedAt(t *time.Time) *AssetCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AssetCreate) SetID(u ulid.ULID) *AssetCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AssetCreate) SetNillableID(u *ulid.ULID) *AssetCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetAssetClassID sets the "asset_class" edge to the AssetClass entity by ID.
func (ac *AssetCreate) SetAssetClassID(id ulid.ULID) *AssetCreate {
	ac.mutation.SetAssetClassID(id)
	return ac
}

// SetAssetClass sets the "asset_class" edge to the AssetClass entity.
func (ac *AssetCreate) SetAssetClass(a *AssetClass) *AssetCreate {
	return ac.SetAssetClassID(a.ID)
}

// SetCryptocurrencyID sets the "cryptocurrency" edge to the Cryptocurrency entity by ID.
func (ac *AssetCreate) SetCryptocurrencyID(id ulid.ULID) *AssetCreate {
	ac.mutation.SetCryptocurrencyID(id)
	return ac
}

// SetNillableCryptocurrencyID sets the "cryptocurrency" edge to the Cryptocurrency entity by ID if the given value is not nil.
func (ac *AssetCreate) SetNillableCryptocurrencyID(id *ulid.ULID) *AssetCreate {
	if id != nil {
		ac = ac.SetCryptocurrencyID(*id)
	}
	return ac
}

// SetCryptocurrency sets the "cryptocurrency" edge to the Cryptocurrency entity.
func (ac *AssetCreate) SetCryptocurrency(c *Cryptocurrency) *AssetCreate {
	return ac.SetCryptocurrencyID(c.ID)
}

// AddTransactionBaseIDs adds the "transaction_base" edge to the Transaction entity by IDs.
func (ac *AssetCreate) AddTransactionBaseIDs(ids ...ulid.ULID) *AssetCreate {
	ac.mutation.AddTransactionBaseIDs(ids...)
	return ac
}

// AddTransactionBase adds the "transaction_base" edges to the Transaction entity.
func (ac *AssetCreate) AddTransactionBase(t ...*Transaction) *AssetCreate {
	ids := make([]ulid.ULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTransactionBaseIDs(ids...)
}

// AddTransactionQuoteIDs adds the "transaction_quote" edge to the Transaction entity by IDs.
func (ac *AssetCreate) AddTransactionQuoteIDs(ids ...ulid.ULID) *AssetCreate {
	ac.mutation.AddTransactionQuoteIDs(ids...)
	return ac
}

// AddTransactionQuote adds the "transaction_quote" edges to the Transaction entity.
func (ac *AssetCreate) AddTransactionQuote(t ...*Transaction) *AssetCreate {
	ids := make([]ulid.ULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTransactionQuoteIDs(ids...)
}

// AddDailyAssetPriceIDs adds the "daily_asset_price" edge to the DailyAssetPrice entity by IDs.
func (ac *AssetCreate) AddDailyAssetPriceIDs(ids ...ulid.ULID) *AssetCreate {
	ac.mutation.AddDailyAssetPriceIDs(ids...)
	return ac
}

// AddDailyAssetPrice adds the "daily_asset_price" edges to the DailyAssetPrice entity.
func (ac *AssetCreate) AddDailyAssetPrice(d ...*DailyAssetPrice) *AssetCreate {
	ids := make([]ulid.ULID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ac.AddDailyAssetPriceIDs(ids...)
}

// Mutation returns the AssetMutation object of the builder.
func (ac *AssetCreate) Mutation() *AssetMutation {
	return ac.mutation
}

// Save creates the Asset in the database.
func (ac *AssetCreate) Save(ctx context.Context) (*Asset, error) {
	var (
		err  error
		node *Asset
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Asset)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AssetMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AssetCreate) SaveX(ctx context.Context) *Asset {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AssetCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AssetCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AssetCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := asset.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := asset.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := asset.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AssetCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Asset.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Asset.updated_at"`)}
	}
	if _, ok := ac.mutation.AssetClassID(); !ok {
		return &ValidationError{Name: "asset_class", err: errors.New(`ent: missing required edge "Asset.asset_class"`)}
	}
	return nil
}

func (ac *AssetCreate) sqlSave(ctx context.Context) (*Asset, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*ulid.ULID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ac *AssetCreate) createSpec() (*Asset, *sqlgraph.CreateSpec) {
	var (
		_node = &Asset{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: asset.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: asset.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asset.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asset.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ac.mutation.AssetClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   asset.AssetClassTable,
			Columns: []string{asset.AssetClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: assetclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.asset_asset_class = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CryptocurrencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   asset.CryptocurrencyTable,
			Columns: []string{asset.CryptocurrencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: cryptocurrency.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.TransactionBaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionBaseTable,
			Columns: asset.TransactionBasePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.TransactionQuoteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionQuoteTable,
			Columns: asset.TransactionQuotePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.DailyAssetPriceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.DailyAssetPriceTable,
			Columns: []string{asset.DailyAssetPriceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dailyassetprice.FieldID,
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

// AssetCreateBulk is the builder for creating many Asset entities in bulk.
type AssetCreateBulk struct {
	config
	builders []*AssetCreate
}

// Save creates the Asset entities in the database.
func (acb *AssetCreateBulk) Save(ctx context.Context) ([]*Asset, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Asset, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AssetMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AssetCreateBulk) SaveX(ctx context.Context) []*Asset {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AssetCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AssetCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
