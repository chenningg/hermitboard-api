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
	"github.com/chenningg/hermitboard-api/ent/account"
	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/ent/portfolio"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/oklog/ulid/v2"
	ulid "github.com/oklog/ulid/v2"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetAuthID sets the "auth_id" field.
func (au *AccountUpdate) SetAuthID(s string) *AccountUpdate {
	au.mutation.SetAuthID(s)
	return au
}

// SetNillableAuthID sets the "auth_id" field if the given value is not nil.
func (au *AccountUpdate) SetNillableAuthID(s *string) *AccountUpdate {
	if s != nil {
		au.SetAuthID(*s)
	}
	return au
}

// ClearAuthID clears the value of the "auth_id" field.
func (au *AccountUpdate) ClearAuthID() *AccountUpdate {
	au.mutation.ClearAuthID()
	return au
}

// SetNickname sets the "nickname" field.
func (au *AccountUpdate) SetNickname(s string) *AccountUpdate {
	au.mutation.SetNickname(s)
	return au
}

// SetEmail sets the "email" field.
func (au *AccountUpdate) SetEmail(s string) *AccountUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetPassword sets the "password" field.
func (au *AccountUpdate) SetPassword(s string) *AccountUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePassword(s *string) *AccountUpdate {
	if s != nil {
		au.SetPassword(*s)
	}
	return au
}

// ClearPassword clears the value of the "password" field.
func (au *AccountUpdate) ClearPassword() *AccountUpdate {
	au.mutation.ClearPassword()
	return au
}

// SetPasswordUpdatedAt sets the "password_updated_at" field.
func (au *AccountUpdate) SetPasswordUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetPasswordUpdatedAt(t)
	return au
}

// AddAuthRoleIDs adds the "auth_roles" edge to the AuthRole entity by IDs.
func (au *AccountUpdate) AddAuthRoleIDs(ids ...ulid.ULID) *AccountUpdate {
	au.mutation.AddAuthRoleIDs(ids...)
	return au
}

// AddAuthRoles adds the "auth_roles" edges to the AuthRole entity.
func (au *AccountUpdate) AddAuthRoles(a ...*AuthRole) *AccountUpdate {
	ids := make([]ulid.ULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAuthRoleIDs(ids...)
}

// AddPortfolioIDs adds the "portfolios" edge to the Portfolio entity by IDs.
func (au *AccountUpdate) AddPortfolioIDs(ids ...ulid.ULID) *AccountUpdate {
	au.mutation.AddPortfolioIDs(ids...)
	return au
}

// AddPortfolios adds the "portfolios" edges to the Portfolio entity.
func (au *AccountUpdate) AddPortfolios(p ...*Portfolio) *AccountUpdate {
	ids := make([]ulid.ULID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.AddPortfolioIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearAuthRoles clears all "auth_roles" edges to the AuthRole entity.
func (au *AccountUpdate) ClearAuthRoles() *AccountUpdate {
	au.mutation.ClearAuthRoles()
	return au
}

// RemoveAuthRoleIDs removes the "auth_roles" edge to AuthRole entities by IDs.
func (au *AccountUpdate) RemoveAuthRoleIDs(ids ...ulid.ULID) *AccountUpdate {
	au.mutation.RemoveAuthRoleIDs(ids...)
	return au
}

// RemoveAuthRoles removes "auth_roles" edges to AuthRole entities.
func (au *AccountUpdate) RemoveAuthRoles(a ...*AuthRole) *AccountUpdate {
	ids := make([]ulid.ULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAuthRoleIDs(ids...)
}

// ClearPortfolios clears all "portfolios" edges to the Portfolio entity.
func (au *AccountUpdate) ClearPortfolios() *AccountUpdate {
	au.mutation.ClearPortfolios()
	return au
}

// RemovePortfolioIDs removes the "portfolios" edge to Portfolio entities by IDs.
func (au *AccountUpdate) RemovePortfolioIDs(ids ...ulid.ULID) *AccountUpdate {
	au.mutation.RemovePortfolioIDs(ids...)
	return au
}

// RemovePortfolios removes "portfolios" edges to Portfolio entities.
func (au *AccountUpdate) RemovePortfolios(p ...*Portfolio) *AccountUpdate {
	ids := make([]ulid.ULID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.RemovePortfolioIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
	if _, ok := au.mutation.PasswordUpdatedAt(); !ok {
		v := account.UpdateDefaultPasswordUpdatedAt()
		au.mutation.SetPasswordUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.AuthID(); ok {
		if err := account.AuthIDValidator(v); err != nil {
			return &ValidationError{Name: "auth_id", err: fmt.Errorf(`ent: validator failed for field "Account.auth_id": %w`, err)}
		}
	}
	if v, ok := au.mutation.Nickname(); ok {
		if err := account.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Account.nickname": %w`, err)}
		}
	}
	if v, ok := au.mutation.Email(); ok {
		if err := account.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Account.email": %w`, err)}
		}
	}
	if v, ok := au.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.AuthID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAuthID,
		})
	}
	if au.mutation.AuthIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldAuthID,
		})
	}
	if value, ok := au.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldNickname,
		})
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmail,
		})
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPassword,
		})
	}
	if au.mutation.PasswordCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldPassword,
		})
	}
	if value, ok := au.mutation.PasswordUpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldPasswordUpdatedAt,
		})
	}
	if au.mutation.AuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.AuthRolesTable,
			Columns: account.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: authrole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAuthRolesIDs(); len(nodes) > 0 && !au.mutation.AuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.AuthRolesTable,
			Columns: account.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: authrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AuthRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.AuthRolesTable,
			Columns: account.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: authrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.PortfoliosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PortfoliosTable,
			Columns: []string{account.PortfoliosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: portfolio.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedPortfoliosIDs(); len(nodes) > 0 && !au.mutation.PortfoliosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PortfoliosTable,
			Columns: []string{account.PortfoliosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: portfolio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.PortfoliosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PortfoliosTable,
			Columns: []string{account.PortfoliosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: portfolio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetAuthID sets the "auth_id" field.
func (auo *AccountUpdateOne) SetAuthID(s string) *AccountUpdateOne {
	auo.mutation.SetAuthID(s)
	return auo
}

// SetNillableAuthID sets the "auth_id" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableAuthID(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetAuthID(*s)
	}
	return auo
}

// ClearAuthID clears the value of the "auth_id" field.
func (auo *AccountUpdateOne) ClearAuthID() *AccountUpdateOne {
	auo.mutation.ClearAuthID()
	return auo
}

// SetNickname sets the "nickname" field.
func (auo *AccountUpdateOne) SetNickname(s string) *AccountUpdateOne {
	auo.mutation.SetNickname(s)
	return auo
}

// SetEmail sets the "email" field.
func (auo *AccountUpdateOne) SetEmail(s string) *AccountUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetPassword sets the "password" field.
func (auo *AccountUpdateOne) SetPassword(s string) *AccountUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePassword(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetPassword(*s)
	}
	return auo
}

// ClearPassword clears the value of the "password" field.
func (auo *AccountUpdateOne) ClearPassword() *AccountUpdateOne {
	auo.mutation.ClearPassword()
	return auo
}

// SetPasswordUpdatedAt sets the "password_updated_at" field.
func (auo *AccountUpdateOne) SetPasswordUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetPasswordUpdatedAt(t)
	return auo
}

// AddAuthRoleIDs adds the "auth_roles" edge to the AuthRole entity by IDs.
func (auo *AccountUpdateOne) AddAuthRoleIDs(ids ...ulid.ULID) *AccountUpdateOne {
	auo.mutation.AddAuthRoleIDs(ids...)
	return auo
}

// AddAuthRoles adds the "auth_roles" edges to the AuthRole entity.
func (auo *AccountUpdateOne) AddAuthRoles(a ...*AuthRole) *AccountUpdateOne {
	ids := make([]ulid.ULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAuthRoleIDs(ids...)
}

// AddPortfolioIDs adds the "portfolios" edge to the Portfolio entity by IDs.
func (auo *AccountUpdateOne) AddPortfolioIDs(ids ...ulid.ULID) *AccountUpdateOne {
	auo.mutation.AddPortfolioIDs(ids...)
	return auo
}

// AddPortfolios adds the "portfolios" edges to the Portfolio entity.
func (auo *AccountUpdateOne) AddPortfolios(p ...*Portfolio) *AccountUpdateOne {
	ids := make([]ulid.ULID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.AddPortfolioIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearAuthRoles clears all "auth_roles" edges to the AuthRole entity.
func (auo *AccountUpdateOne) ClearAuthRoles() *AccountUpdateOne {
	auo.mutation.ClearAuthRoles()
	return auo
}

// RemoveAuthRoleIDs removes the "auth_roles" edge to AuthRole entities by IDs.
func (auo *AccountUpdateOne) RemoveAuthRoleIDs(ids ...ulid.ULID) *AccountUpdateOne {
	auo.mutation.RemoveAuthRoleIDs(ids...)
	return auo
}

// RemoveAuthRoles removes "auth_roles" edges to AuthRole entities.
func (auo *AccountUpdateOne) RemoveAuthRoles(a ...*AuthRole) *AccountUpdateOne {
	ids := make([]ulid.ULID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAuthRoleIDs(ids...)
}

// ClearPortfolios clears all "portfolios" edges to the Portfolio entity.
func (auo *AccountUpdateOne) ClearPortfolios() *AccountUpdateOne {
	auo.mutation.ClearPortfolios()
	return auo
}

// RemovePortfolioIDs removes the "portfolios" edge to Portfolio entities by IDs.
func (auo *AccountUpdateOne) RemovePortfolioIDs(ids ...ulid.ULID) *AccountUpdateOne {
	auo.mutation.RemovePortfolioIDs(ids...)
	return auo
}

// RemovePortfolios removes "portfolios" edges to Portfolio entities.
func (auo *AccountUpdateOne) RemovePortfolios(p ...*Portfolio) *AccountUpdateOne {
	ids := make([]ulid.ULID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.RemovePortfolioIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
	if _, ok := auo.mutation.PasswordUpdatedAt(); !ok {
		v := account.UpdateDefaultPasswordUpdatedAt()
		auo.mutation.SetPasswordUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.AuthID(); ok {
		if err := account.AuthIDValidator(v); err != nil {
			return &ValidationError{Name: "auth_id", err: fmt.Errorf(`ent: validator failed for field "Account.auth_id": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Nickname(); ok {
		if err := account.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Account.nickname": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Email(); ok {
		if err := account.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Account.email": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.AuthID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAuthID,
		})
	}
	if auo.mutation.AuthIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldAuthID,
		})
	}
	if value, ok := auo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldNickname,
		})
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmail,
		})
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPassword,
		})
	}
	if auo.mutation.PasswordCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldPassword,
		})
	}
	if value, ok := auo.mutation.PasswordUpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldPasswordUpdatedAt,
		})
	}
	if auo.mutation.AuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.AuthRolesTable,
			Columns: account.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: authrole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAuthRolesIDs(); len(nodes) > 0 && !auo.mutation.AuthRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.AuthRolesTable,
			Columns: account.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: authrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AuthRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   account.AuthRolesTable,
			Columns: account.AuthRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: authrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.PortfoliosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PortfoliosTable,
			Columns: []string{account.PortfoliosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: portfolio.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedPortfoliosIDs(); len(nodes) > 0 && !auo.mutation.PortfoliosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PortfoliosTable,
			Columns: []string{account.PortfoliosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: portfolio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.PortfoliosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PortfoliosTable,
			Columns: []string{account.PortfoliosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: portfolio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
