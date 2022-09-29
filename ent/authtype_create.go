// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"github.com/chenningg/hermitboard-api/pulid"
)

// AuthTypeCreate is the builder for creating a AuthType entity.
type AuthTypeCreate struct {
	config
	mutation *AuthTypeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (atc *AuthTypeCreate) SetCreatedAt(t time.Time) *AuthTypeCreate {
	atc.mutation.SetCreatedAt(t)
	return atc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atc *AuthTypeCreate) SetNillableCreatedAt(t *time.Time) *AuthTypeCreate {
	if t != nil {
		atc.SetCreatedAt(*t)
	}
	return atc
}

// SetUpdatedAt sets the "updated_at" field.
func (atc *AuthTypeCreate) SetUpdatedAt(t time.Time) *AuthTypeCreate {
	atc.mutation.SetUpdatedAt(t)
	return atc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atc *AuthTypeCreate) SetNillableUpdatedAt(t *time.Time) *AuthTypeCreate {
	if t != nil {
		atc.SetUpdatedAt(*t)
	}
	return atc
}

// SetDeletedAt sets the "deleted_at" field.
func (atc *AuthTypeCreate) SetDeletedAt(t time.Time) *AuthTypeCreate {
	atc.mutation.SetDeletedAt(t)
	return atc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (atc *AuthTypeCreate) SetNillableDeletedAt(t *time.Time) *AuthTypeCreate {
	if t != nil {
		atc.SetDeletedAt(*t)
	}
	return atc
}

// SetValue sets the "value" field.
func (atc *AuthTypeCreate) SetValue(a authtype.Value) *AuthTypeCreate {
	atc.mutation.SetValue(a)
	return atc
}

// SetDescription sets the "description" field.
func (atc *AuthTypeCreate) SetDescription(s string) *AuthTypeCreate {
	atc.mutation.SetDescription(s)
	return atc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (atc *AuthTypeCreate) SetNillableDescription(s *string) *AuthTypeCreate {
	if s != nil {
		atc.SetDescription(*s)
	}
	return atc
}

// SetID sets the "id" field.
func (atc *AuthTypeCreate) SetID(pu pulid.PULID) *AuthTypeCreate {
	atc.mutation.SetID(pu)
	return atc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atc *AuthTypeCreate) SetNillableID(pu *pulid.PULID) *AuthTypeCreate {
	if pu != nil {
		atc.SetID(*pu)
	}
	return atc
}

// Mutation returns the AuthTypeMutation object of the builder.
func (atc *AuthTypeCreate) Mutation() *AuthTypeMutation {
	return atc.mutation
}

// Save creates the AuthType in the database.
func (atc *AuthTypeCreate) Save(ctx context.Context) (*AuthType, error) {
	var (
		err  error
		node *AuthType
	)
	atc.defaults()
	if len(atc.hooks) == 0 {
		if err = atc.check(); err != nil {
			return nil, err
		}
		node, err = atc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atc.check(); err != nil {
				return nil, err
			}
			atc.mutation = mutation
			if node, err = atc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(atc.hooks) - 1; i >= 0; i-- {
			if atc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, atc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AuthType)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AuthTypeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (atc *AuthTypeCreate) SaveX(ctx context.Context) *AuthType {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *AuthTypeCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *AuthTypeCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *AuthTypeCreate) defaults() {
	if _, ok := atc.mutation.CreatedAt(); !ok {
		v := authtype.DefaultCreatedAt()
		atc.mutation.SetCreatedAt(v)
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		v := authtype.DefaultUpdatedAt()
		atc.mutation.SetUpdatedAt(v)
	}
	if _, ok := atc.mutation.ID(); !ok {
		v := authtype.DefaultID()
		atc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *AuthTypeCreate) check() error {
	if _, ok := atc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AuthType.created_at"`)}
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AuthType.updated_at"`)}
	}
	if _, ok := atc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "AuthType.value"`)}
	}
	if v, ok := atc.mutation.Value(); ok {
		if err := authtype.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "AuthType.value": %w`, err)}
		}
	}
	if v, ok := atc.mutation.Description(); ok {
		if err := authtype.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "AuthType.description": %w`, err)}
		}
	}
	return nil
}

func (atc *AuthTypeCreate) sqlSave(ctx context.Context) (*AuthType, error) {
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
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

func (atc *AuthTypeCreate) createSpec() (*AuthType, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthType{config: atc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: authtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authtype.FieldID,
			},
		}
	)
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := atc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authtype.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := atc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authtype.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := atc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authtype.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := atc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: authtype.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := atc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authtype.FieldDescription,
		})
		_node.Description = &value
	}
	return _node, _spec
}

// AuthTypeCreateBulk is the builder for creating many AuthType entities in bulk.
type AuthTypeCreateBulk struct {
	config
	builders []*AuthTypeCreate
}

// Save creates the AuthType entities in the database.
func (atcb *AuthTypeCreateBulk) Save(ctx context.Context) ([]*AuthType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*AuthType, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *AuthTypeCreateBulk) SaveX(ctx context.Context) []*AuthType {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *AuthTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *AuthTypeCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}
