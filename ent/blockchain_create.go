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

// BlockchainCreate is the builder for creating a Blockchain entity.
type BlockchainCreate struct {
	config
	mutation *BlockchainMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bc *BlockchainCreate) SetCreatedAt(t time.Time) *BlockchainCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BlockchainCreate) SetNillableCreatedAt(t *time.Time) *BlockchainCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BlockchainCreate) SetUpdatedAt(t time.Time) *BlockchainCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BlockchainCreate) SetNillableUpdatedAt(t *time.Time) *BlockchainCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetDeletedAt sets the "deleted_at" field.
func (bc *BlockchainCreate) SetDeletedAt(t time.Time) *BlockchainCreate {
	bc.mutation.SetDeletedAt(t)
	return bc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bc *BlockchainCreate) SetNillableDeletedAt(t *time.Time) *BlockchainCreate {
	if t != nil {
		bc.SetDeletedAt(*t)
	}
	return bc
}

// SetName sets the "name" field.
func (bc *BlockchainCreate) SetName(s string) *BlockchainCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetSymbol sets the "symbol" field.
func (bc *BlockchainCreate) SetSymbol(s string) *BlockchainCreate {
	bc.mutation.SetSymbol(s)
	return bc
}

// SetIcon sets the "icon" field.
func (bc *BlockchainCreate) SetIcon(s string) *BlockchainCreate {
	bc.mutation.SetIcon(s)
	return bc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (bc *BlockchainCreate) SetNillableIcon(s *string) *BlockchainCreate {
	if s != nil {
		bc.SetIcon(*s)
	}
	return bc
}

// SetChainID sets the "chain_id" field.
func (bc *BlockchainCreate) SetChainID(i int64) *BlockchainCreate {
	bc.mutation.SetChainID(i)
	return bc
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (bc *BlockchainCreate) SetNillableChainID(i *int64) *BlockchainCreate {
	if i != nil {
		bc.SetChainID(*i)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BlockchainCreate) SetID(pu pulid.PULID) *BlockchainCreate {
	bc.mutation.SetID(pu)
	return bc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bc *BlockchainCreate) SetNillableID(pu *pulid.PULID) *BlockchainCreate {
	if pu != nil {
		bc.SetID(*pu)
	}
	return bc
}

// AddCryptocurrencyIDs adds the "cryptocurrencies" edge to the Cryptocurrency entity by IDs.
func (bc *BlockchainCreate) AddCryptocurrencyIDs(ids ...pulid.PULID) *BlockchainCreate {
	bc.mutation.AddCryptocurrencyIDs(ids...)
	return bc
}

// AddCryptocurrencies adds the "cryptocurrencies" edges to the Cryptocurrency entity.
func (bc *BlockchainCreate) AddCryptocurrencies(c ...*Cryptocurrency) *BlockchainCreate {
	ids := make([]pulid.PULID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bc.AddCryptocurrencyIDs(ids...)
}

// AddBlockchainCryptocurrencyIDs adds the "blockchain_cryptocurrencies" edge to the BlockchainCryptocurrency entity by IDs.
func (bc *BlockchainCreate) AddBlockchainCryptocurrencyIDs(ids ...pulid.PULID) *BlockchainCreate {
	bc.mutation.AddBlockchainCryptocurrencyIDs(ids...)
	return bc
}

// AddBlockchainCryptocurrencies adds the "blockchain_cryptocurrencies" edges to the BlockchainCryptocurrency entity.
func (bc *BlockchainCreate) AddBlockchainCryptocurrencies(b ...*BlockchainCryptocurrency) *BlockchainCreate {
	ids := make([]pulid.PULID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddBlockchainCryptocurrencyIDs(ids...)
}

// Mutation returns the BlockchainMutation object of the builder.
func (bc *BlockchainCreate) Mutation() *BlockchainMutation {
	return bc.mutation
}

// Save creates the Blockchain in the database.
func (bc *BlockchainCreate) Save(ctx context.Context) (*Blockchain, error) {
	var (
		err  error
		node *Blockchain
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlockchainMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, bc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Blockchain)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BlockchainMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlockchainCreate) SaveX(ctx context.Context) *Blockchain {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlockchainCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlockchainCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BlockchainCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := blockchain.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := blockchain.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
	if _, ok := bc.mutation.ID(); !ok {
		v := blockchain.DefaultID()
		bc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlockchainCreate) check() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Blockchain.created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Blockchain.updated_at"`)}
	}
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Blockchain.name"`)}
	}
	if v, ok := bc.mutation.Name(); ok {
		if err := blockchain.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Blockchain.name": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Symbol(); !ok {
		return &ValidationError{Name: "symbol", err: errors.New(`ent: missing required field "Blockchain.symbol"`)}
	}
	if v, ok := bc.mutation.Symbol(); ok {
		if err := blockchain.SymbolValidator(v); err != nil {
			return &ValidationError{Name: "symbol", err: fmt.Errorf(`ent: validator failed for field "Blockchain.symbol": %w`, err)}
		}
	}
	if v, ok := bc.mutation.Icon(); ok {
		if err := blockchain.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "Blockchain.icon": %w`, err)}
		}
	}
	return nil
}

func (bc *BlockchainCreate) sqlSave(ctx context.Context) (*Blockchain, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
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

func (bc *BlockchainCreate) createSpec() (*Blockchain, *sqlgraph.CreateSpec) {
	var (
		_node = &Blockchain{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: blockchain.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: blockchain.FieldID,
			},
		}
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: blockchain.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: blockchain.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := bc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: blockchain.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: blockchain.FieldName,
		})
		_node.Name = value
	}
	if value, ok := bc.mutation.Symbol(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: blockchain.FieldSymbol,
		})
		_node.Symbol = value
	}
	if value, ok := bc.mutation.Icon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: blockchain.FieldIcon,
		})
		_node.Icon = &value
	}
	if value, ok := bc.mutation.ChainID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: blockchain.FieldChainID,
		})
		_node.ChainID = &value
	}
	if nodes := bc.mutation.CryptocurrenciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blockchain.CryptocurrenciesTable,
			Columns: blockchain.CryptocurrenciesPrimaryKey,
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
		createE := &BlockchainCryptocurrencyCreate{config: bc.config, mutation: newBlockchainCryptocurrencyMutation(bc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.BlockchainCryptocurrenciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blockchain.BlockchainCryptocurrenciesTable,
			Columns: []string{blockchain.BlockchainCryptocurrenciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: blockchaincryptocurrency.FieldID,
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

// BlockchainCreateBulk is the builder for creating many Blockchain entities in bulk.
type BlockchainCreateBulk struct {
	config
	builders []*BlockchainCreate
}

// Save creates the Blockchain entities in the database.
func (bcb *BlockchainCreateBulk) Save(ctx context.Context) ([]*Blockchain, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blockchain, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlockchainMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlockchainCreateBulk) SaveX(ctx context.Context) []*Blockchain {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlockchainCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlockchainCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
