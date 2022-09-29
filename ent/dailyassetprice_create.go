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
	"github.com/chenningg/hermitboard-api/ent/dailyassetprice"
	"github.com/chenningg/hermitboard-api/pulid"
)

// DailyAssetPriceCreate is the builder for creating a DailyAssetPrice entity.
type DailyAssetPriceCreate struct {
	config
	mutation *DailyAssetPriceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dapc *DailyAssetPriceCreate) SetCreatedAt(t time.Time) *DailyAssetPriceCreate {
	dapc.mutation.SetCreatedAt(t)
	return dapc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dapc *DailyAssetPriceCreate) SetNillableCreatedAt(t *time.Time) *DailyAssetPriceCreate {
	if t != nil {
		dapc.SetCreatedAt(*t)
	}
	return dapc
}

// SetUpdatedAt sets the "updated_at" field.
func (dapc *DailyAssetPriceCreate) SetUpdatedAt(t time.Time) *DailyAssetPriceCreate {
	dapc.mutation.SetUpdatedAt(t)
	return dapc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dapc *DailyAssetPriceCreate) SetNillableUpdatedAt(t *time.Time) *DailyAssetPriceCreate {
	if t != nil {
		dapc.SetUpdatedAt(*t)
	}
	return dapc
}

// SetDeletedAt sets the "deleted_at" field.
func (dapc *DailyAssetPriceCreate) SetDeletedAt(t time.Time) *DailyAssetPriceCreate {
	dapc.mutation.SetDeletedAt(t)
	return dapc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dapc *DailyAssetPriceCreate) SetNillableDeletedAt(t *time.Time) *DailyAssetPriceCreate {
	if t != nil {
		dapc.SetDeletedAt(*t)
	}
	return dapc
}

// SetTime sets the "time" field.
func (dapc *DailyAssetPriceCreate) SetTime(t time.Time) *DailyAssetPriceCreate {
	dapc.mutation.SetTime(t)
	return dapc
}

// SetOpen sets the "open" field.
func (dapc *DailyAssetPriceCreate) SetOpen(f float64) *DailyAssetPriceCreate {
	dapc.mutation.SetOpen(f)
	return dapc
}

// SetNillableOpen sets the "open" field if the given value is not nil.
func (dapc *DailyAssetPriceCreate) SetNillableOpen(f *float64) *DailyAssetPriceCreate {
	if f != nil {
		dapc.SetOpen(*f)
	}
	return dapc
}

// SetHigh sets the "high" field.
func (dapc *DailyAssetPriceCreate) SetHigh(f float64) *DailyAssetPriceCreate {
	dapc.mutation.SetHigh(f)
	return dapc
}

// SetNillableHigh sets the "high" field if the given value is not nil.
func (dapc *DailyAssetPriceCreate) SetNillableHigh(f *float64) *DailyAssetPriceCreate {
	if f != nil {
		dapc.SetHigh(*f)
	}
	return dapc
}

// SetLow sets the "low" field.
func (dapc *DailyAssetPriceCreate) SetLow(f float64) *DailyAssetPriceCreate {
	dapc.mutation.SetLow(f)
	return dapc
}

// SetNillableLow sets the "low" field if the given value is not nil.
func (dapc *DailyAssetPriceCreate) SetNillableLow(f *float64) *DailyAssetPriceCreate {
	if f != nil {
		dapc.SetLow(*f)
	}
	return dapc
}

// SetClose sets the "close" field.
func (dapc *DailyAssetPriceCreate) SetClose(f float64) *DailyAssetPriceCreate {
	dapc.mutation.SetClose(f)
	return dapc
}

// SetNillableClose sets the "close" field if the given value is not nil.
func (dapc *DailyAssetPriceCreate) SetNillableClose(f *float64) *DailyAssetPriceCreate {
	if f != nil {
		dapc.SetClose(*f)
	}
	return dapc
}

// SetAdjustedClose sets the "adjusted_close" field.
func (dapc *DailyAssetPriceCreate) SetAdjustedClose(f float64) *DailyAssetPriceCreate {
	dapc.mutation.SetAdjustedClose(f)
	return dapc
}

// SetAssetID sets the "asset_id" field.
func (dapc *DailyAssetPriceCreate) SetAssetID(pu pulid.PULID) *DailyAssetPriceCreate {
	dapc.mutation.SetAssetID(pu)
	return dapc
}

// SetID sets the "id" field.
func (dapc *DailyAssetPriceCreate) SetID(pu pulid.PULID) *DailyAssetPriceCreate {
	dapc.mutation.SetID(pu)
	return dapc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dapc *DailyAssetPriceCreate) SetNillableID(pu *pulid.PULID) *DailyAssetPriceCreate {
	if pu != nil {
		dapc.SetID(*pu)
	}
	return dapc
}

// SetAsset sets the "asset" edge to the Asset entity.
func (dapc *DailyAssetPriceCreate) SetAsset(a *Asset) *DailyAssetPriceCreate {
	return dapc.SetAssetID(a.ID)
}

// Mutation returns the DailyAssetPriceMutation object of the builder.
func (dapc *DailyAssetPriceCreate) Mutation() *DailyAssetPriceMutation {
	return dapc.mutation
}

// Save creates the DailyAssetPrice in the database.
func (dapc *DailyAssetPriceCreate) Save(ctx context.Context) (*DailyAssetPrice, error) {
	var (
		err  error
		node *DailyAssetPrice
	)
	dapc.defaults()
	if len(dapc.hooks) == 0 {
		if err = dapc.check(); err != nil {
			return nil, err
		}
		node, err = dapc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DailyAssetPriceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dapc.check(); err != nil {
				return nil, err
			}
			dapc.mutation = mutation
			if node, err = dapc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dapc.hooks) - 1; i >= 0; i-- {
			if dapc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dapc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dapc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DailyAssetPrice)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DailyAssetPriceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dapc *DailyAssetPriceCreate) SaveX(ctx context.Context) *DailyAssetPrice {
	v, err := dapc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dapc *DailyAssetPriceCreate) Exec(ctx context.Context) error {
	_, err := dapc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dapc *DailyAssetPriceCreate) ExecX(ctx context.Context) {
	if err := dapc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dapc *DailyAssetPriceCreate) defaults() {
	if _, ok := dapc.mutation.CreatedAt(); !ok {
		v := dailyassetprice.DefaultCreatedAt()
		dapc.mutation.SetCreatedAt(v)
	}
	if _, ok := dapc.mutation.UpdatedAt(); !ok {
		v := dailyassetprice.DefaultUpdatedAt()
		dapc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dapc.mutation.ID(); !ok {
		v := dailyassetprice.DefaultID()
		dapc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dapc *DailyAssetPriceCreate) check() error {
	if _, ok := dapc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DailyAssetPrice.created_at"`)}
	}
	if _, ok := dapc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DailyAssetPrice.updated_at"`)}
	}
	if _, ok := dapc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "DailyAssetPrice.time"`)}
	}
	if _, ok := dapc.mutation.AdjustedClose(); !ok {
		return &ValidationError{Name: "adjusted_close", err: errors.New(`ent: missing required field "DailyAssetPrice.adjusted_close"`)}
	}
	if _, ok := dapc.mutation.AssetID(); !ok {
		return &ValidationError{Name: "asset_id", err: errors.New(`ent: missing required field "DailyAssetPrice.asset_id"`)}
	}
	if _, ok := dapc.mutation.AssetID(); !ok {
		return &ValidationError{Name: "asset", err: errors.New(`ent: missing required edge "DailyAssetPrice.asset"`)}
	}
	return nil
}

func (dapc *DailyAssetPriceCreate) sqlSave(ctx context.Context) (*DailyAssetPrice, error) {
	_node, _spec := dapc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dapc.driver, _spec); err != nil {
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

func (dapc *DailyAssetPriceCreate) createSpec() (*DailyAssetPrice, *sqlgraph.CreateSpec) {
	var (
		_node = &DailyAssetPrice{config: dapc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dailyassetprice.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: dailyassetprice.FieldID,
			},
		}
	)
	if id, ok := dapc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dapc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailyassetprice.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dapc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailyassetprice.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := dapc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailyassetprice.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := dapc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailyassetprice.FieldTime,
		})
		_node.Time = value
	}
	if value, ok := dapc.mutation.Open(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldOpen,
		})
		_node.Open = &value
	}
	if value, ok := dapc.mutation.High(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldHigh,
		})
		_node.High = &value
	}
	if value, ok := dapc.mutation.Low(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldLow,
		})
		_node.Low = &value
	}
	if value, ok := dapc.mutation.Close(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldClose,
		})
		_node.Close = &value
	}
	if value, ok := dapc.mutation.AdjustedClose(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldAdjustedClose,
		})
		_node.AdjustedClose = value
	}
	if nodes := dapc.mutation.AssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyassetprice.AssetTable,
			Columns: []string{dailyassetprice.AssetColumn},
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
	return _node, _spec
}

// DailyAssetPriceCreateBulk is the builder for creating many DailyAssetPrice entities in bulk.
type DailyAssetPriceCreateBulk struct {
	config
	builders []*DailyAssetPriceCreate
}

// Save creates the DailyAssetPrice entities in the database.
func (dapcb *DailyAssetPriceCreateBulk) Save(ctx context.Context) ([]*DailyAssetPrice, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dapcb.builders))
	nodes := make([]*DailyAssetPrice, len(dapcb.builders))
	mutators := make([]Mutator, len(dapcb.builders))
	for i := range dapcb.builders {
		func(i int, root context.Context) {
			builder := dapcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DailyAssetPriceMutation)
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
					_, err = mutators[i+1].Mutate(root, dapcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dapcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dapcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dapcb *DailyAssetPriceCreateBulk) SaveX(ctx context.Context) []*DailyAssetPrice {
	v, err := dapcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dapcb *DailyAssetPriceCreateBulk) Exec(ctx context.Context) error {
	_, err := dapcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dapcb *DailyAssetPriceCreateBulk) ExecX(ctx context.Context) {
	if err := dapcb.Exec(ctx); err != nil {
		panic(err)
	}
}
