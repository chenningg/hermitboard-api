// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/accountauthrole"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/staffaccount"
	"github.com/chenningg/hermitboard-api/ent/staffaccountauthrole"
	"github.com/chenningg/hermitboard-api/pulid"
)

// AuthRoleCreate is the builder for creating a AuthRole entity.
type AuthRoleCreate struct {
	config
	mutation *AuthRoleMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (arc *AuthRoleCreate) SetCreatedAt(t time.Time) *AuthRoleCreate {
	arc.mutation.SetCreatedAt(t)
	return arc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (arc *AuthRoleCreate) SetNillableCreatedAt(t *time.Time) *AuthRoleCreate {
	if t != nil {
		arc.SetCreatedAt(*t)
	}
	return arc
}

// SetUpdatedAt sets the "updated_at" field.
func (arc *AuthRoleCreate) SetUpdatedAt(t time.Time) *AuthRoleCreate {
	arc.mutation.SetUpdatedAt(t)
	return arc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (arc *AuthRoleCreate) SetNillableUpdatedAt(t *time.Time) *AuthRoleCreate {
	if t != nil {
		arc.SetUpdatedAt(*t)
	}
	return arc
}

// SetDeletedAt sets the "deleted_at" field.
func (arc *AuthRoleCreate) SetDeletedAt(t time.Time) *AuthRoleCreate {
	arc.mutation.SetDeletedAt(t)
	return arc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (arc *AuthRoleCreate) SetNillableDeletedAt(t *time.Time) *AuthRoleCreate {
	if t != nil {
		arc.SetDeletedAt(*t)
	}
	return arc
}

// SetAuthRole sets the "auth_role" field.
func (arc *AuthRoleCreate) SetAuthRole(ar authrole.AuthRole) *AuthRoleCreate {
	arc.mutation.SetAuthRole(ar)
	return arc
}

// SetDescription sets the "description" field.
func (arc *AuthRoleCreate) SetDescription(s string) *AuthRoleCreate {
	arc.mutation.SetDescription(s)
	return arc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (arc *AuthRoleCreate) SetNillableDescription(s *string) *AuthRoleCreate {
	if s != nil {
		arc.SetDescription(*s)
	}
	return arc
}

// SetID sets the "id" field.
func (arc *AuthRoleCreate) SetID(pu pulid.PULID) *AuthRoleCreate {
	arc.mutation.SetID(pu)
	return arc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (arc *AuthRoleCreate) SetNillableID(pu *pulid.PULID) *AuthRoleCreate {
	if pu != nil {
		arc.SetID(*pu)
	}
	return arc
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (arc *AuthRoleCreate) AddAccountIDs(ids ...pulid.PULID) *AuthRoleCreate {
	arc.mutation.AddAccountIDs(ids...)
	return arc
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (arc *AuthRoleCreate) AddAccounts(a ...*Account) *AuthRoleCreate {
	ids := make([]pulid.PULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return arc.AddAccountIDs(ids...)
}

// AddStaffAccountIDs adds the "staff_accounts" edge to the StaffAccount entity by IDs.
func (arc *AuthRoleCreate) AddStaffAccountIDs(ids ...pulid.PULID) *AuthRoleCreate {
	arc.mutation.AddStaffAccountIDs(ids...)
	return arc
}

// AddStaffAccounts adds the "staff_accounts" edges to the StaffAccount entity.
func (arc *AuthRoleCreate) AddStaffAccounts(s ...*StaffAccount) *AuthRoleCreate {
	ids := make([]pulid.PULID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return arc.AddStaffAccountIDs(ids...)
}

// AddAccountAuthRoleIDs adds the "account_auth_roles" edge to the AccountAuthRole entity by IDs.
func (arc *AuthRoleCreate) AddAccountAuthRoleIDs(ids ...pulid.PULID) *AuthRoleCreate {
	arc.mutation.AddAccountAuthRoleIDs(ids...)
	return arc
}

// AddAccountAuthRoles adds the "account_auth_roles" edges to the AccountAuthRole entity.
func (arc *AuthRoleCreate) AddAccountAuthRoles(a ...*AccountAuthRole) *AuthRoleCreate {
	ids := make([]pulid.PULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return arc.AddAccountAuthRoleIDs(ids...)
}

// AddStaffAccountAuthRoleIDs adds the "staff_account_auth_roles" edge to the StaffAccountAuthRole entity by IDs.
func (arc *AuthRoleCreate) AddStaffAccountAuthRoleIDs(ids ...pulid.PULID) *AuthRoleCreate {
	arc.mutation.AddStaffAccountAuthRoleIDs(ids...)
	return arc
}

// AddStaffAccountAuthRoles adds the "staff_account_auth_roles" edges to the StaffAccountAuthRole entity.
func (arc *AuthRoleCreate) AddStaffAccountAuthRoles(s ...*StaffAccountAuthRole) *AuthRoleCreate {
	ids := make([]pulid.PULID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return arc.AddStaffAccountAuthRoleIDs(ids...)
}

// Mutation returns the AuthRoleMutation object of the builder.
func (arc *AuthRoleCreate) Mutation() *AuthRoleMutation {
	return arc.mutation
}

// Save creates the AuthRole in the database.
func (arc *AuthRoleCreate) Save(ctx context.Context) (*AuthRole, error) {
	var (
		err  error
		node *AuthRole
	)
	arc.defaults()
	if len(arc.hooks) == 0 {
		if err = arc.check(); err != nil {
			return nil, err
		}
		node, err = arc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = arc.check(); err != nil {
				return nil, err
			}
			arc.mutation = mutation
			if node, err = arc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(arc.hooks) - 1; i >= 0; i-- {
			if arc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = arc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, arc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AuthRole)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AuthRoleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AuthRoleCreate) SaveX(ctx context.Context) *AuthRole {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *AuthRoleCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *AuthRoleCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arc *AuthRoleCreate) defaults() {
	if _, ok := arc.mutation.CreatedAt(); !ok {
		v := authrole.DefaultCreatedAt()
		arc.mutation.SetCreatedAt(v)
	}
	if _, ok := arc.mutation.UpdatedAt(); !ok {
		v := authrole.DefaultUpdatedAt()
		arc.mutation.SetUpdatedAt(v)
	}
	if _, ok := arc.mutation.ID(); !ok {
		v := authrole.DefaultID()
		arc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *AuthRoleCreate) check() error {
	if _, ok := arc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AuthRole.created_at"`)}
	}
	if _, ok := arc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AuthRole.updated_at"`)}
	}
	if _, ok := arc.mutation.AuthRole(); !ok {
		return &ValidationError{Name: "auth_role", err: errors.New(`ent: missing required field "AuthRole.auth_role"`)}
	}
	if v, ok := arc.mutation.AuthRole(); ok {
		if err := authrole.AuthRoleValidator(v); err != nil {
			return &ValidationError{Name: "auth_role", err: fmt.Errorf(`ent: validator failed for field "AuthRole.auth_role": %w`, err)}
		}
	}
	if v, ok := arc.mutation.Description(); ok {
		if err := authrole.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "AuthRole.description": %w`, err)}
		}
	}
	return nil
}

func (arc *AuthRoleCreate) sqlSave(ctx context.Context) (*AuthRole, error) {
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
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

func (arc *AuthRoleCreate) createSpec() (*AuthRole, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthRole{config: arc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: authrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authrole.FieldID,
			},
		}
	)
	if id, ok := arc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := arc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authrole.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := arc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authrole.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := arc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authrole.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := arc.mutation.AuthRole(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: authrole.FieldAuthRole,
		})
		_node.AuthRole = value
	}
	if value, ok := arc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrole.FieldDescription,
		})
		_node.Description = &value
	}
	if nodes := arc.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   authrole.AccountsTable,
			Columns: authrole.AccountsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AccountAuthRoleCreate{config: arc.config, mutation: newAccountAuthRoleMutation(arc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.StaffAccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   authrole.StaffAccountsTable,
			Columns: authrole.StaffAccountsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: staffaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &StaffAccountAuthRoleCreate{config: arc.config, mutation: newStaffAccountAuthRoleMutation(arc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.AccountAuthRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   authrole.AccountAuthRolesTable,
			Columns: []string{authrole.AccountAuthRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: accountauthrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.StaffAccountAuthRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   authrole.StaffAccountAuthRolesTable,
			Columns: []string{authrole.StaffAccountAuthRolesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AuthRoleCreateBulk is the builder for creating many AuthRole entities in bulk.
type AuthRoleCreateBulk struct {
	config
	builders []*AuthRoleCreate
}

// Save creates the AuthRole entities in the database.
func (arcb *AuthRoleCreateBulk) Save(ctx context.Context) ([]*AuthRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AuthRole, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AuthRoleCreateBulk) SaveX(ctx context.Context) []*AuthRole {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *AuthRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *AuthRoleCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}