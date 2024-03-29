// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"github.com/chenningg/hermitboard-api/ent/predicate"
)

// AuthTypeDelete is the builder for deleting a AuthType entity.
type AuthTypeDelete struct {
	config
	hooks    []Hook
	mutation *AuthTypeMutation
}

// Where appends a list predicates to the AuthTypeDelete builder.
func (atd *AuthTypeDelete) Where(ps ...predicate.AuthType) *AuthTypeDelete {
	atd.mutation.Where(ps...)
	return atd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atd *AuthTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atd.hooks) == 0 {
		affected, err = atd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atd.mutation = mutation
			affected, err = atd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atd.hooks) - 1; i >= 0; i-- {
			if atd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (atd *AuthTypeDelete) ExecX(ctx context.Context) int {
	n, err := atd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atd *AuthTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: authtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authtype.FieldID,
			},
		},
	}
	if ps := atd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, atd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AuthTypeDeleteOne is the builder for deleting a single AuthType entity.
type AuthTypeDeleteOne struct {
	atd *AuthTypeDelete
}

// Exec executes the deletion query.
func (atdo *AuthTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := atdo.atd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authtype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atdo *AuthTypeDeleteOne) ExecX(ctx context.Context) {
	atdo.atd.ExecX(ctx)
}
