// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/blockchain"
	"github.com/chenningg/hermitboard-api/ent/blockchaincryptocurrency"
	"github.com/chenningg/hermitboard-api/ent/cryptocurrency"
	"github.com/chenningg/hermitboard-api/pulid"
)

// BlockchainCryptocurrencyCreate is the builder for creating a BlockchainCryptocurrency entity.
type BlockchainCryptocurrencyCreate struct {
	config
	mutation *BlockchainCryptocurrencyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bcc *BlockchainCryptocurrencyCreate) SetCreatedAt(t time.Time) *BlockchainCryptocurrencyCreate {
	bcc.mutation.SetCreatedAt(t)
	return bcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bcc *BlockchainCryptocurrencyCreate) SetNillableCreatedAt(t *time.Time) *BlockchainCryptocurrencyCreate {
	if t != nil {
		bcc.SetCreatedAt(*t)
	}
	return bcc
}

// SetUpdatedAt sets the "updated_at" field.
func (bcc *BlockchainCryptocurrencyCreate) SetUpdatedAt(t time.Time) *BlockchainCryptocurrencyCreate {
	bcc.mutation.SetUpdatedAt(t)
	return bcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bcc *BlockchainCryptocurrencyCreate) SetNillableUpdatedAt(t *time.Time) *BlockchainCryptocurrencyCreate {
	if t != nil {
		bcc.SetUpdatedAt(*t)
	}
	return bcc
}

// SetDeletedAt sets the "deleted_at" field.
func (bcc *BlockchainCryptocurrencyCreate) SetDeletedAt(t time.Time) *BlockchainCryptocurrencyCreate {
	bcc.mutation.SetDeletedAt(t)
	return bcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bcc *BlockchainCryptocurrencyCreate) SetNillableDeletedAt(t *time.Time) *BlockchainCryptocurrencyCreate {
	if t != nil {
		bcc.SetDeletedAt(*t)
	}
	return bcc
}

// SetBlockchainID sets the "blockchain_id" field.
func (bcc *BlockchainCryptocurrencyCreate) SetBlockchainID(pu pulid.PULID) *BlockchainCryptocurrencyCreate {
	bcc.mutation.SetBlockchainID(pu)
	return bcc
}

// SetCryptocurrencyID sets the "cryptocurrency_id" field.
func (bcc *BlockchainCryptocurrencyCreate) SetCryptocurrencyID(pu pulid.PULID) *BlockchainCryptocurrencyCreate {
	bcc.mutation.SetCryptocurrencyID(pu)
	return bcc
}

// SetID sets the "id" field.
func (bcc *BlockchainCryptocurrencyCreate) SetID(pu pulid.PULID) *BlockchainCryptocurrencyCreate {
	bcc.mutation.SetID(pu)
	return bcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bcc *BlockchainCryptocurrencyCreate) SetNillableID(pu *pulid.PULID) *BlockchainCryptocurrencyCreate {
	if pu != nil {
		bcc.SetID(*pu)
	}
	return bcc
}

// SetBlockchain sets the "blockchain" edge to the Blockchain entity.
func (bcc *BlockchainCryptocurrencyCreate) SetBlockchain(b *Blockchain) *BlockchainCryptocurrencyCreate {
	return bcc.SetBlockchainID(b.ID)
}

// SetCryptocurrency sets the "cryptocurrency" edge to the Cryptocurrency entity.
func (bcc *BlockchainCryptocurrencyCreate) SetCryptocurrency(c *Cryptocurrency) *BlockchainCryptocurrencyCreate {
	return bcc.SetCryptocurrencyID(c.ID)
}

// Mutation returns the BlockchainCryptocurrencyMutation object of the builder.
func (bcc *BlockchainCryptocurrencyCreate) Mutation() *BlockchainCryptocurrencyMutation {
	return bcc.mutation
}

// Save creates the BlockchainCryptocurrency in the database.
func (bcc *BlockchainCryptocurrencyCreate) Save(ctx context.Context) (*BlockchainCryptocurrency, error) {
	var (
		err  error
		node *BlockchainCryptocurrency
	)
	bcc.defaults()
	if len(bcc.hooks) == 0 {
		if err = bcc.check(); err != nil {
			return nil, err
		}
		node, err = bcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlockchainCryptocurrencyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bcc.check(); err != nil {
				return nil, err
			}
			bcc.mutation = mutation
			if node, err = bcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bcc.hooks) - 1; i >= 0; i-- {
			if bcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bcc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, bcc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*BlockchainCryptocurrency)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BlockchainCryptocurrencyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bcc *BlockchainCryptocurrencyCreate) SaveX(ctx context.Context) *BlockchainCryptocurrency {
	v, err := bcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcc *BlockchainCryptocurrencyCreate) Exec(ctx context.Context) error {
	_, err := bcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcc *BlockchainCryptocurrencyCreate) ExecX(ctx context.Context) {
	if err := bcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bcc *BlockchainCryptocurrencyCreate) defaults() {
	if _, ok := bcc.mutation.CreatedAt(); !ok {
		v := blockchaincryptocurrency.DefaultCreatedAt()
		bcc.mutation.SetCreatedAt(v)
	}
	if _, ok := bcc.mutation.UpdatedAt(); !ok {
		v := blockchaincryptocurrency.DefaultUpdatedAt()
		bcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := bcc.mutation.ID(); !ok {
		v := blockchaincryptocurrency.DefaultID()
		bcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcc *BlockchainCryptocurrencyCreate) check() error {
	if _, ok := bcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BlockchainCryptocurrency.created_at"`)}
	}
	if _, ok := bcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "BlockchainCryptocurrency.updated_at"`)}
	}
	if _, ok := bcc.mutation.BlockchainID(); !ok {
		return &ValidationError{Name: "blockchain_id", err: errors.New(`ent: missing required field "BlockchainCryptocurrency.blockchain_id"`)}
	}
	if _, ok := bcc.mutation.CryptocurrencyID(); !ok {
		return &ValidationError{Name: "cryptocurrency_id", err: errors.New(`ent: missing required field "BlockchainCryptocurrency.cryptocurrency_id"`)}
	}
	if _, ok := bcc.mutation.BlockchainID(); !ok {
		return &ValidationError{Name: "blockchain", err: errors.New(`ent: missing required edge "BlockchainCryptocurrency.blockchain"`)}
	}
	if _, ok := bcc.mutation.CryptocurrencyID(); !ok {
		return &ValidationError{Name: "cryptocurrency", err: errors.New(`ent: missing required edge "BlockchainCryptocurrency.cryptocurrency"`)}
	}
	return nil
}

func (bcc *BlockchainCryptocurrencyCreate) sqlSave(ctx context.Context) (*BlockchainCryptocurrency, error) {
	_node, _spec := bcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bcc.driver, _spec); err != nil {
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

func (bcc *BlockchainCryptocurrencyCreate) createSpec() (*BlockchainCryptocurrency, *sqlgraph.CreateSpec) {
	var (
		_node = &BlockchainCryptocurrency{config: bcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: blockchaincryptocurrency.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: blockchaincryptocurrency.FieldID,
			},
		}
	)
	if id, ok := bcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: blockchaincryptocurrency.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := bcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: blockchaincryptocurrency.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := bcc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: blockchaincryptocurrency.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := bcc.mutation.BlockchainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blockchaincryptocurrency.BlockchainTable,
			Columns: []string{blockchaincryptocurrency.BlockchainColumn},
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
		_node.BlockchainID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bcc.mutation.CryptocurrencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blockchaincryptocurrency.CryptocurrencyTable,
			Columns: []string{blockchaincryptocurrency.CryptocurrencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: cryptocurrency.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CryptocurrencyID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BlockchainCryptocurrencyCreateBulk is the builder for creating many BlockchainCryptocurrency entities in bulk.
type BlockchainCryptocurrencyCreateBulk struct {
	config
	builders []*BlockchainCryptocurrencyCreate
}

// Save creates the BlockchainCryptocurrency entities in the database.
func (bccb *BlockchainCryptocurrencyCreateBulk) Save(ctx context.Context) ([]*BlockchainCryptocurrency, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bccb.builders))
	nodes := make([]*BlockchainCryptocurrency, len(bccb.builders))
	mutators := make([]Mutator, len(bccb.builders))
	for i := range bccb.builders {
		func(i int, root context.Context) {
			builder := bccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlockchainCryptocurrencyMutation)
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
					_, err = mutators[i+1].Mutate(root, bccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bccb *BlockchainCryptocurrencyCreateBulk) SaveX(ctx context.Context) []*BlockchainCryptocurrency {
	v, err := bccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bccb *BlockchainCryptocurrencyCreateBulk) Exec(ctx context.Context) error {
	_, err := bccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bccb *BlockchainCryptocurrencyCreateBulk) ExecX(ctx context.Context) {
	if err := bccb.Exec(ctx); err != nil {
		panic(err)
	}
}
