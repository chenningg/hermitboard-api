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
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/ent/staffaccount"
	"github.com/chenningg/hermitboard-api/ent/staffaccountauthrole"
	"github.com/chenningg/hermitboard-api/pulid"
)

// StaffAccountUpdate is the builder for updating StaffAccount entities.
type StaffAccountUpdate struct {
	config
	hooks    []Hook
	mutation *StaffAccountMutation
}

// Where appends a list predicates to the StaffAccountUpdate builder.
func (sau *StaffAccountUpdate) Where(ps ...predicate.StaffAccount) *StaffAccountUpdate {
	sau.mutation.Where(ps...)
	return sau
}

// SetUpdatedAt sets the "updated_at" field.
func (sau *StaffAccountUpdate) SetUpdatedAt(t time.Time) *StaffAccountUpdate {
	sau.mutation.SetUpdatedAt(t)
	return sau
}

// SetDeletedAt sets the "deleted_at" field.
func (sau *StaffAccountUpdate) SetDeletedAt(t time.Time) *StaffAccountUpdate {
	sau.mutation.SetDeletedAt(t)
	return sau
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sau *StaffAccountUpdate) ClearDeletedAt() *StaffAccountUpdate {
	sau.mutation.ClearDeletedAt()
	return sau
}

// SetAuthType sets the "auth_type" field.
func (sau *StaffAccountUpdate) SetAuthType(st staffaccount.AuthType) *StaffAccountUpdate {
	sau.mutation.SetAuthType(st)
	return sau
}

// SetNillableAuthType sets the "auth_type" field if the given value is not nil.
func (sau *StaffAccountUpdate) SetNillableAuthType(st *staffaccount.AuthType) *StaffAccountUpdate {
	if st != nil {
		sau.SetAuthType(*st)
	}
	return sau
}

// SetNickname sets the "nickname" field.
func (sau *StaffAccountUpdate) SetNickname(s string) *StaffAccountUpdate {
	sau.mutation.SetNickname(s)
	return sau
}

// SetEmail sets the "email" field.
func (sau *StaffAccountUpdate) SetEmail(s string) *StaffAccountUpdate {
	sau.mutation.SetEmail(s)
	return sau
}

// SetPassword sets the "password" field.
func (sau *StaffAccountUpdate) SetPassword(s string) *StaffAccountUpdate {
	sau.mutation.SetPassword(s)
	return sau
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (sau *StaffAccountUpdate) SetNillablePassword(s *string) *StaffAccountUpdate {
	if s != nil {
		sau.SetPassword(*s)
	}
	return sau
}

// ClearPassword clears the value of the "password" field.
func (sau *StaffAccountUpdate) ClearPassword() *StaffAccountUpdate {
	sau.mutation.ClearPassword()
	return sau
}

// SetPasswordUpdatedAt sets the "password_updated_at" field.
func (sau *StaffAccountUpdate) SetPasswordUpdatedAt(t time.Time) *StaffAccountUpdate {
	sau.mutation.SetPasswordUpdatedAt(t)
	return sau
}

// AddAuthRoleIDs adds the "auth_roles" edge to the AuthRole entity by IDs.
func (sau *StaffAccountUpdate) AddAuthRoleIDs(ids ...pulid.PULID) *StaffAccountUpdate {
	sau.mutation.AddAuthRoleIDs(ids...)
	return sau
}

// AddAuthRoles adds the "auth_roles" edges to the AuthRole entity.
func (sau *StaffAccountUpdate) AddAuthRoles(a ...*AuthRole) *StaffAccountUpdate {
	ids := make([]pulid.PULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sau.AddAuthRoleIDs(ids...)
}

// AddStaffAccountAuthRoleIDs adds the "staff_account_auth_roles" edge to the StaffAccountAuthRole entity by IDs.
func (sau *StaffAccountUpdate) AddStaffAccountAuthRoleIDs(ids ...pulid.PULID) *StaffAccountUpdate {
	sau.mutation.AddStaffAccountAuthRoleIDs(ids...)
	return sau
}

// AddStaffAccountAuthRoles adds the "staff_account_auth_roles" edges to the StaffAccountAuthRole entity.
func (sau *StaffAccountUpdate) AddStaffAccountAuthRoles(s ...*StaffAccountAuthRole) *StaffAccountUpdate {
	ids := make([]pulid.PULID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sau.AddStaffAccountAuthRoleIDs(ids...)
}

// Mutation returns the StaffAccountMutation object of the builder.
func (sau *StaffAccountUpdate) Mutation() *StaffAccountMutation {
	return sau.mutation
}

// ClearAuthRoles clears all "auth_roles" edges to the AuthRole entity.
func (sau *StaffAccountUpdate) ClearAuthRoles() *StaffAccountUpdate {
	sau.mutation.ClearAuthRoles()
	return sau
}

// RemoveAuthRoleIDs removes the "auth_roles" edge to AuthRole entities by IDs.
func (sau *StaffAccountUpdate) RemoveAuthRoleIDs(ids ...pulid.PULID) *StaffAccountUpdate {
	sau.mutation.RemoveAuthRoleIDs(ids...)
	return sau
}

// RemoveAuthRoles removes "auth_roles" edges to AuthRole entities.
func (sau *StaffAccountUpdate) RemoveAuthRoles(a ...*AuthRole) *StaffAccountUpdate {
	ids := make([]pulid.PULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sau.RemoveAuthRoleIDs(ids...)
}

// ClearStaffAccountAuthRoles clears all "staff_account_auth_roles" edges to the StaffAccountAuthRole entity.
func (sau *StaffAccountUpdate) ClearStaffAccountAuthRoles() *StaffAccountUpdate {
	sau.mutation.ClearStaffAccountAuthRoles()
	return sau
}

// RemoveStaffAccountAuthRoleIDs removes the "staff_account_auth_roles" edge to StaffAccountAuthRole entities by IDs.
func (sau *StaffAccountUpdate) RemoveStaffAccountAuthRoleIDs(ids ...pulid.PULID) *StaffAccountUpdate {
	sau.mutation.RemoveStaffAccountAuthRoleIDs(ids...)
	return sau
}

// RemoveStaffAccountAuthRoles removes "staff_account_auth_roles" edges to StaffAccountAuthRole entities.
func (sau *StaffAccountUpdate) RemoveStaffAccountAuthRoles(s ...*StaffAccountAuthRole) *StaffAccountUpdate {
	ids := make([]pulid.PULID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sau.RemoveStaffAccountAuthRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *StaffAccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sau.defaults()
	if len(sau.hooks) == 0 {
		if err = sau.check(); err != nil {
			return 0, err
		}
		affected, err = sau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StaffAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sau.check(); err != nil {
				return 0, err
			}
			sau.mutation = mutation
			affected, err = sau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sau.hooks) - 1; i >= 0; i-- {
			if sau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sau *StaffAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *StaffAccountUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *StaffAccountUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sau *StaffAccountUpdate) defaults() {
	if _, ok := sau.mutation.UpdatedAt(); !ok {
		v := staffaccount.UpdateDefaultUpdatedAt()
		sau.mutation.SetUpdatedAt(v)
	}
	if _, ok := sau.mutation.DeletedAt(); !ok && !sau.mutation.DeletedAtCleared() {
		v := staffaccount.UpdateDefaultDeletedAt()
		sau.mutation.SetDeletedAt(v)
	}
	if _, ok := sau.mutation.PasswordUpdatedAt(); !ok {
		v := staffaccount.UpdateDefaultPasswordUpdatedAt()
		sau.mutation.SetPasswordUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sau *StaffAccountUpdate) check() error {
	if v, ok := sau.mutation.AuthType(); ok {
		if err := staffaccount.AuthTypeValidator(v); err != nil {
			return &ValidationError{Name: "auth_type", err: fmt.Errorf(`ent: validator failed for field "StaffAccount.auth_type": %w`, err)}
		}
	}
	if v, ok := sau.mutation.Nickname(); ok {
		if err := staffaccount.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "StaffAccount.nickname": %w`, err)}
		}
	}
	if v, ok := sau.mutation.Email(); ok {
		if err := staffaccount.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "StaffAccount.email": %w`, err)}
		}
	}
	if v, ok := sau.mutation.Password(); ok {
		if err := staffaccount.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "StaffAccount.password": %w`, err)}
		}
	}
	return nil
}

func (sau *StaffAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   staffaccount.Table,
			Columns: staffaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: staffaccount.FieldID,
			},
		},
	}
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: staffaccount.FieldUpdatedAt,
		})
	}
	if value, ok := sau.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: staffaccount.FieldDeletedAt,
		})
	}
	if sau.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: staffaccount.FieldDeletedAt,
		})
	}
	if value, ok := sau.mutation.AuthType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: staffaccount.FieldAuthType,
		})
	}
	if value, ok := sau.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: staffaccount.FieldNickname,
		})
	}
	if value, ok := sau.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: staffaccount.FieldEmail,
		})
	}
	if value, ok := sau.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: staffaccount.FieldPassword,
		})
	}
	if sau.mutation.PasswordCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: staffaccount.FieldPassword,
		})
	}
	if value, ok := sau.mutation.PasswordUpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: staffaccount.FieldPasswordUpdatedAt,
		})
	}
	if sau.mutation.AuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   staffaccount.AuthRolesTable,
			Columns: staffaccount.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: authrole.FieldID,
				},
			},
		}
		createE := &StaffAccountAuthRoleCreate{config: sau.config, mutation: newStaffAccountAuthRoleMutation(sau.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.RemovedAuthRolesIDs(); len(nodes) > 0 && !sau.mutation.AuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   staffaccount.AuthRolesTable,
			Columns: staffaccount.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: authrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &StaffAccountAuthRoleCreate{config: sau.config, mutation: newStaffAccountAuthRoleMutation(sau.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.AuthRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   staffaccount.AuthRolesTable,
			Columns: staffaccount.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: authrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &StaffAccountAuthRoleCreate{config: sau.config, mutation: newStaffAccountAuthRoleMutation(sau.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sau.mutation.StaffAccountAuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   staffaccount.StaffAccountAuthRolesTable,
			Columns: []string{staffaccount.StaffAccountAuthRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: staffaccountauthrole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.RemovedStaffAccountAuthRolesIDs(); len(nodes) > 0 && !sau.mutation.StaffAccountAuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   staffaccount.StaffAccountAuthRolesTable,
			Columns: []string{staffaccount.StaffAccountAuthRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: staffaccountauthrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.StaffAccountAuthRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   staffaccount.StaffAccountAuthRolesTable,
			Columns: []string{staffaccount.StaffAccountAuthRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: staffaccountauthrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{staffaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// StaffAccountUpdateOne is the builder for updating a single StaffAccount entity.
type StaffAccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StaffAccountMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (sauo *StaffAccountUpdateOne) SetUpdatedAt(t time.Time) *StaffAccountUpdateOne {
	sauo.mutation.SetUpdatedAt(t)
	return sauo
}

// SetDeletedAt sets the "deleted_at" field.
func (sauo *StaffAccountUpdateOne) SetDeletedAt(t time.Time) *StaffAccountUpdateOne {
	sauo.mutation.SetDeletedAt(t)
	return sauo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (sauo *StaffAccountUpdateOne) ClearDeletedAt() *StaffAccountUpdateOne {
	sauo.mutation.ClearDeletedAt()
	return sauo
}

// SetAuthType sets the "auth_type" field.
func (sauo *StaffAccountUpdateOne) SetAuthType(st staffaccount.AuthType) *StaffAccountUpdateOne {
	sauo.mutation.SetAuthType(st)
	return sauo
}

// SetNillableAuthType sets the "auth_type" field if the given value is not nil.
func (sauo *StaffAccountUpdateOne) SetNillableAuthType(st *staffaccount.AuthType) *StaffAccountUpdateOne {
	if st != nil {
		sauo.SetAuthType(*st)
	}
	return sauo
}

// SetNickname sets the "nickname" field.
func (sauo *StaffAccountUpdateOne) SetNickname(s string) *StaffAccountUpdateOne {
	sauo.mutation.SetNickname(s)
	return sauo
}

// SetEmail sets the "email" field.
func (sauo *StaffAccountUpdateOne) SetEmail(s string) *StaffAccountUpdateOne {
	sauo.mutation.SetEmail(s)
	return sauo
}

// SetPassword sets the "password" field.
func (sauo *StaffAccountUpdateOne) SetPassword(s string) *StaffAccountUpdateOne {
	sauo.mutation.SetPassword(s)
	return sauo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (sauo *StaffAccountUpdateOne) SetNillablePassword(s *string) *StaffAccountUpdateOne {
	if s != nil {
		sauo.SetPassword(*s)
	}
	return sauo
}

// ClearPassword clears the value of the "password" field.
func (sauo *StaffAccountUpdateOne) ClearPassword() *StaffAccountUpdateOne {
	sauo.mutation.ClearPassword()
	return sauo
}

// SetPasswordUpdatedAt sets the "password_updated_at" field.
func (sauo *StaffAccountUpdateOne) SetPasswordUpdatedAt(t time.Time) *StaffAccountUpdateOne {
	sauo.mutation.SetPasswordUpdatedAt(t)
	return sauo
}

// AddAuthRoleIDs adds the "auth_roles" edge to the AuthRole entity by IDs.
func (sauo *StaffAccountUpdateOne) AddAuthRoleIDs(ids ...pulid.PULID) *StaffAccountUpdateOne {
	sauo.mutation.AddAuthRoleIDs(ids...)
	return sauo
}

// AddAuthRoles adds the "auth_roles" edges to the AuthRole entity.
func (sauo *StaffAccountUpdateOne) AddAuthRoles(a ...*AuthRole) *StaffAccountUpdateOne {
	ids := make([]pulid.PULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sauo.AddAuthRoleIDs(ids...)
}

// AddStaffAccountAuthRoleIDs adds the "staff_account_auth_roles" edge to the StaffAccountAuthRole entity by IDs.
func (sauo *StaffAccountUpdateOne) AddStaffAccountAuthRoleIDs(ids ...pulid.PULID) *StaffAccountUpdateOne {
	sauo.mutation.AddStaffAccountAuthRoleIDs(ids...)
	return sauo
}

// AddStaffAccountAuthRoles adds the "staff_account_auth_roles" edges to the StaffAccountAuthRole entity.
func (sauo *StaffAccountUpdateOne) AddStaffAccountAuthRoles(s ...*StaffAccountAuthRole) *StaffAccountUpdateOne {
	ids := make([]pulid.PULID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sauo.AddStaffAccountAuthRoleIDs(ids...)
}

// Mutation returns the StaffAccountMutation object of the builder.
func (sauo *StaffAccountUpdateOne) Mutation() *StaffAccountMutation {
	return sauo.mutation
}

// ClearAuthRoles clears all "auth_roles" edges to the AuthRole entity.
func (sauo *StaffAccountUpdateOne) ClearAuthRoles() *StaffAccountUpdateOne {
	sauo.mutation.ClearAuthRoles()
	return sauo
}

// RemoveAuthRoleIDs removes the "auth_roles" edge to AuthRole entities by IDs.
func (sauo *StaffAccountUpdateOne) RemoveAuthRoleIDs(ids ...pulid.PULID) *StaffAccountUpdateOne {
	sauo.mutation.RemoveAuthRoleIDs(ids...)
	return sauo
}

// RemoveAuthRoles removes "auth_roles" edges to AuthRole entities.
func (sauo *StaffAccountUpdateOne) RemoveAuthRoles(a ...*AuthRole) *StaffAccountUpdateOne {
	ids := make([]pulid.PULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sauo.RemoveAuthRoleIDs(ids...)
}

// ClearStaffAccountAuthRoles clears all "staff_account_auth_roles" edges to the StaffAccountAuthRole entity.
func (sauo *StaffAccountUpdateOne) ClearStaffAccountAuthRoles() *StaffAccountUpdateOne {
	sauo.mutation.ClearStaffAccountAuthRoles()
	return sauo
}

// RemoveStaffAccountAuthRoleIDs removes the "staff_account_auth_roles" edge to StaffAccountAuthRole entities by IDs.
func (sauo *StaffAccountUpdateOne) RemoveStaffAccountAuthRoleIDs(ids ...pulid.PULID) *StaffAccountUpdateOne {
	sauo.mutation.RemoveStaffAccountAuthRoleIDs(ids...)
	return sauo
}

// RemoveStaffAccountAuthRoles removes "staff_account_auth_roles" edges to StaffAccountAuthRole entities.
func (sauo *StaffAccountUpdateOne) RemoveStaffAccountAuthRoles(s ...*StaffAccountAuthRole) *StaffAccountUpdateOne {
	ids := make([]pulid.PULID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sauo.RemoveStaffAccountAuthRoleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *StaffAccountUpdateOne) Select(field string, fields ...string) *StaffAccountUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated StaffAccount entity.
func (sauo *StaffAccountUpdateOne) Save(ctx context.Context) (*StaffAccount, error) {
	var (
		err  error
		node *StaffAccount
	)
	sauo.defaults()
	if len(sauo.hooks) == 0 {
		if err = sauo.check(); err != nil {
			return nil, err
		}
		node, err = sauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StaffAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sauo.check(); err != nil {
				return nil, err
			}
			sauo.mutation = mutation
			node, err = sauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sauo.hooks) - 1; i >= 0; i-- {
			if sauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*StaffAccount)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from StaffAccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *StaffAccountUpdateOne) SaveX(ctx context.Context) *StaffAccount {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *StaffAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *StaffAccountUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sauo *StaffAccountUpdateOne) defaults() {
	if _, ok := sauo.mutation.UpdatedAt(); !ok {
		v := staffaccount.UpdateDefaultUpdatedAt()
		sauo.mutation.SetUpdatedAt(v)
	}
	if _, ok := sauo.mutation.DeletedAt(); !ok && !sauo.mutation.DeletedAtCleared() {
		v := staffaccount.UpdateDefaultDeletedAt()
		sauo.mutation.SetDeletedAt(v)
	}
	if _, ok := sauo.mutation.PasswordUpdatedAt(); !ok {
		v := staffaccount.UpdateDefaultPasswordUpdatedAt()
		sauo.mutation.SetPasswordUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sauo *StaffAccountUpdateOne) check() error {
	if v, ok := sauo.mutation.AuthType(); ok {
		if err := staffaccount.AuthTypeValidator(v); err != nil {
			return &ValidationError{Name: "auth_type", err: fmt.Errorf(`ent: validator failed for field "StaffAccount.auth_type": %w`, err)}
		}
	}
	if v, ok := sauo.mutation.Nickname(); ok {
		if err := staffaccount.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "StaffAccount.nickname": %w`, err)}
		}
	}
	if v, ok := sauo.mutation.Email(); ok {
		if err := staffaccount.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "StaffAccount.email": %w`, err)}
		}
	}
	if v, ok := sauo.mutation.Password(); ok {
		if err := staffaccount.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "StaffAccount.password": %w`, err)}
		}
	}
	return nil
}

func (sauo *StaffAccountUpdateOne) sqlSave(ctx context.Context) (_node *StaffAccount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   staffaccount.Table,
			Columns: staffaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: staffaccount.FieldID,
			},
		},
	}
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StaffAccount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, staffaccount.FieldID)
		for _, f := range fields {
			if !staffaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != staffaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: staffaccount.FieldUpdatedAt,
		})
	}
	if value, ok := sauo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: staffaccount.FieldDeletedAt,
		})
	}
	if sauo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: staffaccount.FieldDeletedAt,
		})
	}
	if value, ok := sauo.mutation.AuthType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: staffaccount.FieldAuthType,
		})
	}
	if value, ok := sauo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: staffaccount.FieldNickname,
		})
	}
	if value, ok := sauo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: staffaccount.FieldEmail,
		})
	}
	if value, ok := sauo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: staffaccount.FieldPassword,
		})
	}
	if sauo.mutation.PasswordCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: staffaccount.FieldPassword,
		})
	}
	if value, ok := sauo.mutation.PasswordUpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: staffaccount.FieldPasswordUpdatedAt,
		})
	}
	if sauo.mutation.AuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   staffaccount.AuthRolesTable,
			Columns: staffaccount.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: authrole.FieldID,
				},
			},
		}
		createE := &StaffAccountAuthRoleCreate{config: sauo.config, mutation: newStaffAccountAuthRoleMutation(sauo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.RemovedAuthRolesIDs(); len(nodes) > 0 && !sauo.mutation.AuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   staffaccount.AuthRolesTable,
			Columns: staffaccount.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: authrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &StaffAccountAuthRoleCreate{config: sauo.config, mutation: newStaffAccountAuthRoleMutation(sauo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.AuthRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   staffaccount.AuthRolesTable,
			Columns: staffaccount.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: authrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &StaffAccountAuthRoleCreate{config: sauo.config, mutation: newStaffAccountAuthRoleMutation(sauo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sauo.mutation.StaffAccountAuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   staffaccount.StaffAccountAuthRolesTable,
			Columns: []string{staffaccount.StaffAccountAuthRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: staffaccountauthrole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.RemovedStaffAccountAuthRolesIDs(); len(nodes) > 0 && !sauo.mutation.StaffAccountAuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   staffaccount.StaffAccountAuthRolesTable,
			Columns: []string{staffaccount.StaffAccountAuthRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: staffaccountauthrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.StaffAccountAuthRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   staffaccount.StaffAccountAuthRolesTable,
			Columns: []string{staffaccount.StaffAccountAuthRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: staffaccountauthrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &StaffAccount{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{staffaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}