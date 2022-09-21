// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/ent/staffaccountauthrole"
)

// StaffAccountAuthRoleDelete is the builder for deleting a StaffAccountAuthRole entity.
type StaffAccountAuthRoleDelete struct {
	config
	hooks    []Hook
	mutation *StaffAccountAuthRoleMutation
}

// Where appends a list predicates to the StaffAccountAuthRoleDelete builder.
func (saard *StaffAccountAuthRoleDelete) Where(ps ...predicate.StaffAccountAuthRole) *StaffAccountAuthRoleDelete {
	saard.mutation.Where(ps...)
	return saard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (saard *StaffAccountAuthRoleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(saard.hooks) == 0 {
		affected, err = saard.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StaffAccountAuthRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			saard.mutation = mutation
			affected, err = saard.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(saard.hooks) - 1; i >= 0; i-- {
			if saard.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = saard.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, saard.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (saard *StaffAccountAuthRoleDelete) ExecX(ctx context.Context) int {
	n, err := saard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (saard *StaffAccountAuthRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: staffaccountauthrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: staffaccountauthrole.FieldID,
			},
		},
	}
	if ps := saard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, saard.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// StaffAccountAuthRoleDeleteOne is the builder for deleting a single StaffAccountAuthRole entity.
type StaffAccountAuthRoleDeleteOne struct {
	saard *StaffAccountAuthRoleDelete
}

// Exec executes the deletion query.
func (saardo *StaffAccountAuthRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := saardo.saard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{staffaccountauthrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (saardo *StaffAccountAuthRoleDeleteOne) ExecX(ctx context.Context) {
	saardo.saard.ExecX(ctx)
}