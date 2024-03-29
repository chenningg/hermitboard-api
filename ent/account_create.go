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
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/authtype"
	"github.com/chenningg/hermitboard-api/ent/connection"
	"github.com/chenningg/hermitboard-api/ent/portfolio"
	"github.com/chenningg/hermitboard-api/pulid"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AccountCreate) SetCreatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AccountCreate) SetUpdatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AccountCreate) SetDeletedAt(t time.Time) *AccountCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableDeletedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetNickname sets the "nickname" field.
func (ac *AccountCreate) SetNickname(s string) *AccountCreate {
	ac.mutation.SetNickname(s)
	return ac
}

// SetEmail sets the "email" field.
func (ac *AccountCreate) SetEmail(s string) *AccountCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// SetEmailConfirmed sets the "email_confirmed" field.
func (ac *AccountCreate) SetEmailConfirmed(b bool) *AccountCreate {
	ac.mutation.SetEmailConfirmed(b)
	return ac
}

// SetNillableEmailConfirmed sets the "email_confirmed" field if the given value is not nil.
func (ac *AccountCreate) SetNillableEmailConfirmed(b *bool) *AccountCreate {
	if b != nil {
		ac.SetEmailConfirmed(*b)
	}
	return ac
}

// SetPassword sets the "password" field.
func (ac *AccountCreate) SetPassword(s string) *AccountCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (ac *AccountCreate) SetNillablePassword(s *string) *AccountCreate {
	if s != nil {
		ac.SetPassword(*s)
	}
	return ac
}

// SetPasswordUpdatedAt sets the "password_updated_at" field.
func (ac *AccountCreate) SetPasswordUpdatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetPasswordUpdatedAt(t)
	return ac
}

// SetNillablePasswordUpdatedAt sets the "password_updated_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillablePasswordUpdatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetPasswordUpdatedAt(*t)
	}
	return ac
}

// SetProfilePictureURL sets the "profile_picture_url" field.
func (ac *AccountCreate) SetProfilePictureURL(s string) *AccountCreate {
	ac.mutation.SetProfilePictureURL(s)
	return ac
}

// SetNillableProfilePictureURL sets the "profile_picture_url" field if the given value is not nil.
func (ac *AccountCreate) SetNillableProfilePictureURL(s *string) *AccountCreate {
	if s != nil {
		ac.SetProfilePictureURL(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AccountCreate) SetID(pu pulid.PULID) *AccountCreate {
	ac.mutation.SetID(pu)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AccountCreate) SetNillableID(pu *pulid.PULID) *AccountCreate {
	if pu != nil {
		ac.SetID(*pu)
	}
	return ac
}

// AddFriendIDs adds the "friends" edge to the Account entity by IDs.
func (ac *AccountCreate) AddFriendIDs(ids ...pulid.PULID) *AccountCreate {
	ac.mutation.AddFriendIDs(ids...)
	return ac
}

// AddFriends adds the "friends" edges to the Account entity.
func (ac *AccountCreate) AddFriends(a ...*Account) *AccountCreate {
	ids := make([]pulid.PULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddFriendIDs(ids...)
}

// AddAuthRoleIDs adds the "auth_roles" edge to the AuthRole entity by IDs.
func (ac *AccountCreate) AddAuthRoleIDs(ids ...pulid.PULID) *AccountCreate {
	ac.mutation.AddAuthRoleIDs(ids...)
	return ac
}

// AddAuthRoles adds the "auth_roles" edges to the AuthRole entity.
func (ac *AccountCreate) AddAuthRoles(a ...*AuthRole) *AccountCreate {
	ids := make([]pulid.PULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAuthRoleIDs(ids...)
}

// AddPortfolioIDs adds the "portfolios" edge to the Portfolio entity by IDs.
func (ac *AccountCreate) AddPortfolioIDs(ids ...pulid.PULID) *AccountCreate {
	ac.mutation.AddPortfolioIDs(ids...)
	return ac
}

// AddPortfolios adds the "portfolios" edges to the Portfolio entity.
func (ac *AccountCreate) AddPortfolios(p ...*Portfolio) *AccountCreate {
	ids := make([]pulid.PULID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddPortfolioIDs(ids...)
}

// SetAuthTypeID sets the "auth_type" edge to the AuthType entity by ID.
func (ac *AccountCreate) SetAuthTypeID(id pulid.PULID) *AccountCreate {
	ac.mutation.SetAuthTypeID(id)
	return ac
}

// SetAuthType sets the "auth_type" edge to the AuthType entity.
func (ac *AccountCreate) SetAuthType(a *AuthType) *AccountCreate {
	return ac.SetAuthTypeID(a.ID)
}

// AddConnectionIDs adds the "connections" edge to the Connection entity by IDs.
func (ac *AccountCreate) AddConnectionIDs(ids ...pulid.PULID) *AccountCreate {
	ac.mutation.AddConnectionIDs(ids...)
	return ac
}

// AddConnections adds the "connections" edges to the Connection entity.
func (ac *AccountCreate) AddConnections(c ...*Connection) *AccountCreate {
	ids := make([]pulid.PULID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ac.AddConnectionIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Account)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := account.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := account.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.EmailConfirmed(); !ok {
		v := account.DefaultEmailConfirmed
		ac.mutation.SetEmailConfirmed(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := account.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Account.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Account.updated_at"`)}
	}
	if _, ok := ac.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`ent: missing required field "Account.nickname"`)}
	}
	if v, ok := ac.mutation.Nickname(); ok {
		if err := account.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Account.nickname": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Account.email"`)}
	}
	if v, ok := ac.mutation.Email(); ok {
		if err := account.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Account.email": %w`, err)}
		}
	}
	if _, ok := ac.mutation.EmailConfirmed(); !ok {
		return &ValidationError{Name: "email_confirmed", err: errors.New(`ent: missing required field "Account.email_confirmed"`)}
	}
	if v, ok := ac.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	if len(ac.mutation.AuthRolesIDs()) == 0 {
		return &ValidationError{Name: "auth_roles", err: errors.New(`ent: missing required edge "Account.auth_roles"`)}
	}
	if _, ok := ac.mutation.AuthTypeID(); !ok {
		return &ValidationError{Name: "auth_type", err: errors.New(`ent: missing required edge "Account.auth_type"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: account.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: account.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := ac.mutation.Nickname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldNickname,
		})
		_node.Nickname = value
	}
	if value, ok := ac.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := ac.mutation.EmailConfirmed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldEmailConfirmed,
		})
		_node.EmailConfirmed = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPassword,
		})
		_node.Password = &value
	}
	if value, ok := ac.mutation.PasswordUpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldPasswordUpdatedAt,
		})
		_node.PasswordUpdatedAt = &value
	}
	if value, ok := ac.mutation.ProfilePictureURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldProfilePictureURL,
		})
		_node.ProfilePictureURL = &value
	}
	if nodes := ac.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.FriendsTable,
			Columns: account.FriendsPrimaryKey,
			Bidi:    true,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AuthRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.AuthRolesTable,
			Columns: account.AuthRolesPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.PortfoliosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PortfoliosTable,
			Columns: []string{account.PortfoliosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: portfolio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AuthTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   account.AuthTypeTable,
			Columns: []string{account.AuthTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: authtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_auth_type = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ConnectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.ConnectionsTable,
			Columns: []string{account.ConnectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: connection.FieldID,
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

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	builders []*AccountCreate
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
