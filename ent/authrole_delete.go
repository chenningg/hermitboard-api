// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/predicate"
)

// AuthRoleDelete is the builder for deleting a AuthRole entity.
type AuthRoleDelete struct {
	config
	hooks    []Hook
	mutation *AuthRoleMutation
}

// Where appends a list predicates to the AuthRoleDelete builder.
func (ard *AuthRoleDelete) Where(ps ...predicate.AuthRole) *AuthRoleDelete {
	ard.mutation.Where(ps...)
	return ard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ard *AuthRoleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ard.hooks) == 0 {
		affected, err = ard.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ard.mutation = mutation
			affected, err = ard.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ard.hooks) - 1; i >= 0; i-- {
			if ard.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ard.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ard.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ard *AuthRoleDelete) ExecX(ctx context.Context) int {
	n, err := ard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ard *AuthRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: authrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authrole.FieldID,
			},
		},
	}
	if ps := ard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ard.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AuthRoleDeleteOne is the builder for deleting a single AuthRole entity.
type AuthRoleDeleteOne struct {
	ard *AuthRoleDelete
}

// Exec executes the deletion query.
func (ardo *AuthRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := ardo.ard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ardo *AuthRoleDeleteOne) ExecX(ctx context.Context) {
	ardo.ard.ExecX(ctx)
}