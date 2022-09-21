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
	"github.com/chenningg/hermitboard-api/pulid"
)

// AssetClassCreate is the builder for creating a AssetClass entity.
type AssetClassCreate struct {
	config
	mutation *AssetClassMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (acc *AssetClassCreate) SetCreatedAt(t time.Time) *AssetClassCreate {
	acc.mutation.SetCreatedAt(t)
	return acc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acc *AssetClassCreate) SetNillableCreatedAt(t *time.Time) *AssetClassCreate {
	if t != nil {
		acc.SetCreatedAt(*t)
	}
	return acc
}

// SetUpdatedAt sets the "updated_at" field.
func (acc *AssetClassCreate) SetUpdatedAt(t time.Time) *AssetClassCreate {
	acc.mutation.SetUpdatedAt(t)
	return acc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (acc *AssetClassCreate) SetNillableUpdatedAt(t *time.Time) *AssetClassCreate {
	if t != nil {
		acc.SetUpdatedAt(*t)
	}
	return acc
}

// SetDeletedAt sets the "deleted_at" field.
func (acc *AssetClassCreate) SetDeletedAt(t time.Time) *AssetClassCreate {
	acc.mutation.SetDeletedAt(t)
	return acc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (acc *AssetClassCreate) SetNillableDeletedAt(t *time.Time) *AssetClassCreate {
	if t != nil {
		acc.SetDeletedAt(*t)
	}
	return acc
}

// SetAssetClass sets the "asset_class" field.
func (acc *AssetClassCreate) SetAssetClass(ac assetclass.AssetClass) *AssetClassCreate {
	acc.mutation.SetAssetClass(ac)
	return acc
}

// SetDescription sets the "description" field.
func (acc *AssetClassCreate) SetDescription(s string) *AssetClassCreate {
	acc.mutation.SetDescription(s)
	return acc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (acc *AssetClassCreate) SetNillableDescription(s *string) *AssetClassCreate {
	if s != nil {
		acc.SetDescription(*s)
	}
	return acc
}

// SetID sets the "id" field.
func (acc *AssetClassCreate) SetID(pu pulid.PULID) *AssetClassCreate {
	acc.mutation.SetID(pu)
	return acc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (acc *AssetClassCreate) SetNillableID(pu *pulid.PULID) *AssetClassCreate {
	if pu != nil {
		acc.SetID(*pu)
	}
	return acc
}

// AddAssetIDs adds the "assets" edge to the Asset entity by IDs.
func (acc *AssetClassCreate) AddAssetIDs(ids ...pulid.PULID) *AssetClassCreate {
	acc.mutation.AddAssetIDs(ids...)
	return acc
}

// AddAssets adds the "assets" edges to the Asset entity.
func (acc *AssetClassCreate) AddAssets(a ...*Asset) *AssetClassCreate {
	ids := make([]pulid.PULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return acc.AddAssetIDs(ids...)
}

// Mutation returns the AssetClassMutation object of the builder.
func (acc *AssetClassCreate) Mutation() *AssetClassMutation {
	return acc.mutation
}

// Save creates the AssetClass in the database.
func (acc *AssetClassCreate) Save(ctx context.Context) (*AssetClass, error) {
	var (
		err  error
		node *AssetClass
	)
	acc.defaults()
	if len(acc.hooks) == 0 {
		if err = acc.check(); err != nil {
			return nil, err
		}
		node, err = acc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetClassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acc.check(); err != nil {
				return nil, err
			}
			acc.mutation = mutation
			if node, err = acc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(acc.hooks) - 1; i >= 0; i-- {
			if acc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, acc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AssetClass)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AssetClassMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (acc *AssetClassCreate) SaveX(ctx context.Context) *AssetClass {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *AssetClassCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *AssetClassCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acc *AssetClassCreate) defaults() {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		v := assetclass.DefaultCreatedAt()
		acc.mutation.SetCreatedAt(v)
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		v := assetclass.DefaultUpdatedAt()
		acc.mutation.SetUpdatedAt(v)
	}
	if _, ok := acc.mutation.ID(); !ok {
		v := assetclass.DefaultID()
		acc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acc *AssetClassCreate) check() error {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AssetClass.created_at"`)}
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AssetClass.updated_at"`)}
	}
	if _, ok := acc.mutation.AssetClass(); !ok {
		return &ValidationError{Name: "asset_class", err: errors.New(`ent: missing required field "AssetClass.asset_class"`)}
	}
	if v, ok := acc.mutation.AssetClass(); ok {
		if err := assetclass.AssetClassValidator(v); err != nil {
			return &ValidationError{Name: "asset_class", err: fmt.Errorf(`ent: validator failed for field "AssetClass.asset_class": %w`, err)}
		}
	}
	if v, ok := acc.mutation.Description(); ok {
		if err := assetclass.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "AssetClass.description": %w`, err)}
		}
	}
	return nil
}

func (acc *AssetClassCreate) sqlSave(ctx context.Context) (*AssetClass, error) {
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
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

func (acc *AssetClassCreate) createSpec() (*AssetClass, *sqlgraph.CreateSpec) {
	var (
		_node = &AssetClass{config: acc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: assetclass.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: assetclass.FieldID,
			},
		}
	)
	if id, ok := acc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := acc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assetclass.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := acc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assetclass.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := acc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assetclass.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := acc.mutation.AssetClass(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: assetclass.FieldAssetClass,
		})
		_node.AssetClass = value
	}
	if value, ok := acc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assetclass.FieldDescription,
		})
		_node.Description = &value
	}
	if nodes := acc.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   assetclass.AssetsTable,
			Columns: []string{assetclass.AssetsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AssetClassCreateBulk is the builder for creating many AssetClass entities in bulk.
type AssetClassCreateBulk struct {
	config
	builders []*AssetClassCreate
}

// Save creates the AssetClass entities in the database.
func (accb *AssetClassCreateBulk) Save(ctx context.Context) ([]*AssetClass, error) {
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*AssetClass, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AssetClassMutation)
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
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *AssetClassCreateBulk) SaveX(ctx context.Context) []*AssetClass {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *AssetClassCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *AssetClassCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}
