// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"github.com/chenningg/hermitboard-api/ent/predicate"
)

// AuthTypeUpdate is the builder for updating AuthType entities.
type AuthTypeUpdate struct {
	config
	hooks    []Hook
	mutation *AuthTypeMutation
}

// Where appends a list predicates to the AuthTypeUpdate builder.
func (atu *AuthTypeUpdate) Where(ps ...predicate.AuthType) *AuthTypeUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetUpdatedAt sets the "updated_at" field.
func (atu *AuthTypeUpdate) SetUpdatedAt(t time.Time) *AuthTypeUpdate {
	atu.mutation.SetUpdatedAt(t)
	return atu
}

// SetDeletedAt sets the "deleted_at" field.
func (atu *AuthTypeUpdate) SetDeletedAt(t time.Time) *AuthTypeUpdate {
	atu.mutation.SetDeletedAt(t)
	return atu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (atu *AuthTypeUpdate) ClearDeletedAt() *AuthTypeUpdate {
	atu.mutation.ClearDeletedAt()
	return atu
}

// SetValue sets the "value" field.
func (atu *AuthTypeUpdate) SetValue(a authtype.Value) *AuthTypeUpdate {
	atu.mutation.SetValue(a)
	return atu
}

// SetDescription sets the "description" field.
func (atu *AuthTypeUpdate) SetDescription(s string) *AuthTypeUpdate {
	atu.mutation.SetDescription(s)
	return atu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (atu *AuthTypeUpdate) SetNillableDescription(s *string) *AuthTypeUpdate {
	if s != nil {
		atu.SetDescription(*s)
	}
	return atu
}

// ClearDescription clears the value of the "description" field.
func (atu *AuthTypeUpdate) ClearDescription() *AuthTypeUpdate {
	atu.mutation.ClearDescription()
	return atu
}

// Mutation returns the AuthTypeMutation object of the builder.
func (atu *AuthTypeUpdate) Mutation() *AuthTypeMutation {
	return atu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AuthTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	atu.defaults()
	if len(atu.hooks) == 0 {
		if err = atu.check(); err != nil {
			return 0, err
		}
		affected, err = atu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atu.check(); err != nil {
				return 0, err
			}
			atu.mutation = mutation
			affected, err = atu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atu.hooks) - 1; i >= 0; i-- {
			if atu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AuthTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AuthTypeUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AuthTypeUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atu *AuthTypeUpdate) defaults() {
	if _, ok := atu.mutation.UpdatedAt(); !ok {
		v := authtype.UpdateDefaultUpdatedAt()
		atu.mutation.SetUpdatedAt(v)
	}
	if _, ok := atu.mutation.DeletedAt(); !ok && !atu.mutation.DeletedAtCleared() {
		v := authtype.UpdateDefaultDeletedAt()
		atu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *AuthTypeUpdate) check() error {
	if v, ok := atu.mutation.Value(); ok {
		if err := authtype.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "AuthType.value": %w`, err)}
		}
	}
	if v, ok := atu.mutation.Description(); ok {
		if err := authtype.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "AuthType.description": %w`, err)}
		}
	}
	return nil
}

func (atu *AuthTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authtype.Table,
			Columns: authtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authtype.FieldID,
			},
		},
	}
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authtype.FieldUpdatedAt,
		})
	}
	if value, ok := atu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authtype.FieldDeletedAt,
		})
	}
	if atu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: authtype.FieldDeletedAt,
		})
	}
	if value, ok := atu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: authtype.FieldValue,
		})
	}
	if value, ok := atu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authtype.FieldDescription,
		})
	}
	if atu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: authtype.FieldDescription,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AuthTypeUpdateOne is the builder for updating a single AuthType entity.
type AuthTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthTypeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (atuo *AuthTypeUpdateOne) SetUpdatedAt(t time.Time) *AuthTypeUpdateOne {
	atuo.mutation.SetUpdatedAt(t)
	return atuo
}

// SetDeletedAt sets the "deleted_at" field.
func (atuo *AuthTypeUpdateOne) SetDeletedAt(t time.Time) *AuthTypeUpdateOne {
	atuo.mutation.SetDeletedAt(t)
	return atuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (atuo *AuthTypeUpdateOne) ClearDeletedAt() *AuthTypeUpdateOne {
	atuo.mutation.ClearDeletedAt()
	return atuo
}

// SetValue sets the "value" field.
func (atuo *AuthTypeUpdateOne) SetValue(a authtype.Value) *AuthTypeUpdateOne {
	atuo.mutation.SetValue(a)
	return atuo
}

// SetDescription sets the "description" field.
func (atuo *AuthTypeUpdateOne) SetDescription(s string) *AuthTypeUpdateOne {
	atuo.mutation.SetDescription(s)
	return atuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (atuo *AuthTypeUpdateOne) SetNillableDescription(s *string) *AuthTypeUpdateOne {
	if s != nil {
		atuo.SetDescription(*s)
	}
	return atuo
}

// ClearDescription clears the value of the "description" field.
func (atuo *AuthTypeUpdateOne) ClearDescription() *AuthTypeUpdateOne {
	atuo.mutation.ClearDescription()
	return atuo
}

// Mutation returns the AuthTypeMutation object of the builder.
func (atuo *AuthTypeUpdateOne) Mutation() *AuthTypeMutation {
	return atuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AuthTypeUpdateOne) Select(field string, fields ...string) *AuthTypeUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AuthType entity.
func (atuo *AuthTypeUpdateOne) Save(ctx context.Context) (*AuthType, error) {
	var (
		err  error
		node *AuthType
	)
	atuo.defaults()
	if len(atuo.hooks) == 0 {
		if err = atuo.check(); err != nil {
			return nil, err
		}
		node, err = atuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atuo.check(); err != nil {
				return nil, err
			}
			atuo.mutation = mutation
			node, err = atuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atuo.hooks) - 1; i >= 0; i-- {
			if atuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, atuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (atuo *AuthTypeUpdateOne) SaveX(ctx context.Context) *AuthType {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AuthTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AuthTypeUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atuo *AuthTypeUpdateOne) defaults() {
	if _, ok := atuo.mutation.UpdatedAt(); !ok {
		v := authtype.UpdateDefaultUpdatedAt()
		atuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := atuo.mutation.DeletedAt(); !ok && !atuo.mutation.DeletedAtCleared() {
		v := authtype.UpdateDefaultDeletedAt()
		atuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *AuthTypeUpdateOne) check() error {
	if v, ok := atuo.mutation.Value(); ok {
		if err := authtype.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "AuthType.value": %w`, err)}
		}
	}
	if v, ok := atuo.mutation.Description(); ok {
		if err := authtype.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "AuthType.description": %w`, err)}
		}
	}
	return nil
}

func (atuo *AuthTypeUpdateOne) sqlSave(ctx context.Context) (_node *AuthType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authtype.Table,
			Columns: authtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authtype.FieldID,
			},
		},
	}
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AuthType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authtype.FieldID)
		for _, f := range fields {
			if !authtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != authtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authtype.FieldUpdatedAt,
		})
	}
	if value, ok := atuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authtype.FieldDeletedAt,
		})
	}
	if atuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: authtype.FieldDeletedAt,
		})
	}
	if value, ok := atuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: authtype.FieldValue,
		})
	}
	if value, ok := atuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authtype.FieldDescription,
		})
	}
	if atuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: authtype.FieldDescription,
		})
	}
	_node = &AuthType{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
