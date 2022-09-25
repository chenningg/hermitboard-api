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
	"github.com/chenningg/hermitboard-api/ent/blockchain"
	"github.com/chenningg/hermitboard-api/ent/cryptocurrency"
	"github.com/chenningg/hermitboard-api/pulid"
)

// CryptocurrencyCreate is the builder for creating a Cryptocurrency entity.
type CryptocurrencyCreate struct {
	config
	mutation *CryptocurrencyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CryptocurrencyCreate) SetCreatedAt(t time.Time) *CryptocurrencyCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CryptocurrencyCreate) SetNillableCreatedAt(t *time.Time) *CryptocurrencyCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CryptocurrencyCreate) SetUpdatedAt(t time.Time) *CryptocurrencyCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CryptocurrencyCreate) SetNillableUpdatedAt(t *time.Time) *CryptocurrencyCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CryptocurrencyCreate) SetDeletedAt(t time.Time) *CryptocurrencyCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CryptocurrencyCreate) SetNillableDeletedAt(t *time.Time) *CryptocurrencyCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetSymbol sets the "symbol" field.
func (cc *CryptocurrencyCreate) SetSymbol(s string) *CryptocurrencyCreate {
	cc.mutation.SetSymbol(s)
	return cc
}

// SetIcon sets the "icon" field.
func (cc *CryptocurrencyCreate) SetIcon(s string) *CryptocurrencyCreate {
	cc.mutation.SetIcon(s)
	return cc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (cc *CryptocurrencyCreate) SetNillableIcon(s *string) *CryptocurrencyCreate {
	if s != nil {
		cc.SetIcon(*s)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CryptocurrencyCreate) SetName(s string) *CryptocurrencyCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetAssetID sets the "asset_id" field.
func (cc *CryptocurrencyCreate) SetAssetID(pu pulid.PULID) *CryptocurrencyCreate {
	cc.mutation.SetAssetID(pu)
	return cc
}

// SetID sets the "id" field.
func (cc *CryptocurrencyCreate) SetID(pu pulid.PULID) *CryptocurrencyCreate {
	cc.mutation.SetID(pu)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CryptocurrencyCreate) SetNillableID(pu *pulid.PULID) *CryptocurrencyCreate {
	if pu != nil {
		cc.SetID(*pu)
	}
	return cc
}

// SetAsset sets the "asset" edge to the Asset entity.
func (cc *CryptocurrencyCreate) SetAsset(a *Asset) *CryptocurrencyCreate {
	return cc.SetAssetID(a.ID)
}

// AddBlockchainIDs adds the "blockchains" edge to the Blockchain entity by IDs.
func (cc *CryptocurrencyCreate) AddBlockchainIDs(ids ...pulid.PULID) *CryptocurrencyCreate {
	cc.mutation.AddBlockchainIDs(ids...)
	return cc
}

// AddBlockchains adds the "blockchains" edges to the Blockchain entity.
func (cc *CryptocurrencyCreate) AddBlockchains(b ...*Blockchain) *CryptocurrencyCreate {
	ids := make([]pulid.PULID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cc.AddBlockchainIDs(ids...)
}

// Mutation returns the CryptocurrencyMutation object of the builder.
func (cc *CryptocurrencyCreate) Mutation() *CryptocurrencyMutation {
	return cc.mutation
}

// Save creates the Cryptocurrency in the database.
func (cc *CryptocurrencyCreate) Save(ctx context.Context) (*Cryptocurrency, error) {
	var (
		err  error
		node *Cryptocurrency
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CryptocurrencyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Cryptocurrency)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CryptocurrencyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CryptocurrencyCreate) SaveX(ctx context.Context) *Cryptocurrency {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CryptocurrencyCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CryptocurrencyCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CryptocurrencyCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := cryptocurrency.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := cryptocurrency.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := cryptocurrency.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CryptocurrencyCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Cryptocurrency.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Cryptocurrency.updated_at"`)}
	}
	if _, ok := cc.mutation.Symbol(); !ok {
		return &ValidationError{Name: "symbol", err: errors.New(`ent: missing required field "Cryptocurrency.symbol"`)}
	}
	if v, ok := cc.mutation.Symbol(); ok {
		if err := cryptocurrency.SymbolValidator(v); err != nil {
			return &ValidationError{Name: "symbol", err: fmt.Errorf(`ent: validator failed for field "Cryptocurrency.symbol": %w`, err)}
		}
	}
	if v, ok := cc.mutation.Icon(); ok {
		if err := cryptocurrency.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "Cryptocurrency.icon": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Cryptocurrency.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := cryptocurrency.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Cryptocurrency.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.AssetID(); !ok {
		return &ValidationError{Name: "asset_id", err: errors.New(`ent: missing required field "Cryptocurrency.asset_id"`)}
	}
	if _, ok := cc.mutation.AssetID(); !ok {
		return &ValidationError{Name: "asset", err: errors.New(`ent: missing required edge "Cryptocurrency.asset"`)}
	}
	if len(cc.mutation.BlockchainsIDs()) == 0 {
		return &ValidationError{Name: "blockchains", err: errors.New(`ent: missing required edge "Cryptocurrency.blockchains"`)}
	}
	return nil
}

func (cc *CryptocurrencyCreate) sqlSave(ctx context.Context) (*Cryptocurrency, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *CryptocurrencyCreate) createSpec() (*Cryptocurrency, *sqlgraph.CreateSpec) {
	var (
		_node = &Cryptocurrency{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cryptocurrency.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: cryptocurrency.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cryptocurrency.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cryptocurrency.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cryptocurrency.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := cc.mutation.Symbol(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cryptocurrency.FieldSymbol,
		})
		_node.Symbol = value
	}
	if value, ok := cc.mutation.Icon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cryptocurrency.FieldIcon,
		})
		_node.Icon = &value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cryptocurrency.FieldName,
		})
		_node.Name = value
	}
	if nodes := cc.mutation.AssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   cryptocurrency.AssetTable,
			Columns: []string{cryptocurrency.AssetColumn},
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
		_node.AssetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.BlockchainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cryptocurrency.BlockchainsTable,
			Columns: cryptocurrency.BlockchainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: blockchain.FieldID,
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

// CryptocurrencyCreateBulk is the builder for creating many Cryptocurrency entities in bulk.
type CryptocurrencyCreateBulk struct {
	config
	builders []*CryptocurrencyCreate
}

// Save creates the Cryptocurrency entities in the database.
func (ccb *CryptocurrencyCreateBulk) Save(ctx context.Context) ([]*Cryptocurrency, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Cryptocurrency, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CryptocurrencyMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CryptocurrencyCreateBulk) SaveX(ctx context.Context) []*Cryptocurrency {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CryptocurrencyCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CryptocurrencyCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
