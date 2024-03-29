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
	"github.com/chenningg/hermitboard-api/ent/asset"
	"github.com/chenningg/hermitboard-api/ent/blockchain"
	"github.com/chenningg/hermitboard-api/ent/cryptocurrency"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// CryptocurrencyUpdate is the builder for updating Cryptocurrency entities.
type CryptocurrencyUpdate struct {
	config
	hooks    []Hook
	mutation *CryptocurrencyMutation
}

// Where appends a list predicates to the CryptocurrencyUpdate builder.
func (cu *CryptocurrencyUpdate) Where(ps ...predicate.Cryptocurrency) *CryptocurrencyUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CryptocurrencyUpdate) SetUpdatedAt(t time.Time) *CryptocurrencyUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CryptocurrencyUpdate) SetDeletedAt(t time.Time) *CryptocurrencyUpdate {
	cu.mutation.SetDeletedAt(t)
	return cu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cu *CryptocurrencyUpdate) ClearDeletedAt() *CryptocurrencyUpdate {
	cu.mutation.ClearDeletedAt()
	return cu
}

// SetSymbol sets the "symbol" field.
func (cu *CryptocurrencyUpdate) SetSymbol(s string) *CryptocurrencyUpdate {
	cu.mutation.SetSymbol(s)
	return cu
}

// SetIcon sets the "icon" field.
func (cu *CryptocurrencyUpdate) SetIcon(s string) *CryptocurrencyUpdate {
	cu.mutation.SetIcon(s)
	return cu
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (cu *CryptocurrencyUpdate) SetNillableIcon(s *string) *CryptocurrencyUpdate {
	if s != nil {
		cu.SetIcon(*s)
	}
	return cu
}

// ClearIcon clears the value of the "icon" field.
func (cu *CryptocurrencyUpdate) ClearIcon() *CryptocurrencyUpdate {
	cu.mutation.ClearIcon()
	return cu
}

// SetName sets the "name" field.
func (cu *CryptocurrencyUpdate) SetName(s string) *CryptocurrencyUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetAssetID sets the "asset_id" field.
func (cu *CryptocurrencyUpdate) SetAssetID(pu pulid.PULID) *CryptocurrencyUpdate {
	cu.mutation.SetAssetID(pu)
	return cu
}

// SetAsset sets the "asset" edge to the Asset entity.
func (cu *CryptocurrencyUpdate) SetAsset(a *Asset) *CryptocurrencyUpdate {
	return cu.SetAssetID(a.ID)
}

// AddBlockchainIDs adds the "blockchains" edge to the Blockchain entity by IDs.
func (cu *CryptocurrencyUpdate) AddBlockchainIDs(ids ...pulid.PULID) *CryptocurrencyUpdate {
	cu.mutation.AddBlockchainIDs(ids...)
	return cu
}

// AddBlockchains adds the "blockchains" edges to the Blockchain entity.
func (cu *CryptocurrencyUpdate) AddBlockchains(b ...*Blockchain) *CryptocurrencyUpdate {
	ids := make([]pulid.PULID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.AddBlockchainIDs(ids...)
}

// Mutation returns the CryptocurrencyMutation object of the builder.
func (cu *CryptocurrencyUpdate) Mutation() *CryptocurrencyMutation {
	return cu.mutation
}

// ClearAsset clears the "asset" edge to the Asset entity.
func (cu *CryptocurrencyUpdate) ClearAsset() *CryptocurrencyUpdate {
	cu.mutation.ClearAsset()
	return cu
}

// ClearBlockchains clears all "blockchains" edges to the Blockchain entity.
func (cu *CryptocurrencyUpdate) ClearBlockchains() *CryptocurrencyUpdate {
	cu.mutation.ClearBlockchains()
	return cu
}

// RemoveBlockchainIDs removes the "blockchains" edge to Blockchain entities by IDs.
func (cu *CryptocurrencyUpdate) RemoveBlockchainIDs(ids ...pulid.PULID) *CryptocurrencyUpdate {
	cu.mutation.RemoveBlockchainIDs(ids...)
	return cu
}

// RemoveBlockchains removes "blockchains" edges to Blockchain entities.
func (cu *CryptocurrencyUpdate) RemoveBlockchains(b ...*Blockchain) *CryptocurrencyUpdate {
	ids := make([]pulid.PULID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.RemoveBlockchainIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CryptocurrencyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CryptocurrencyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CryptocurrencyUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CryptocurrencyUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CryptocurrencyUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CryptocurrencyUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := cryptocurrency.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	if _, ok := cu.mutation.DeletedAt(); !ok && !cu.mutation.DeletedAtCleared() {
		v := cryptocurrency.UpdateDefaultDeletedAt()
		cu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CryptocurrencyUpdate) check() error {
	if v, ok := cu.mutation.Symbol(); ok {
		if err := cryptocurrency.SymbolValidator(v); err != nil {
			return &ValidationError{Name: "symbol", err: fmt.Errorf(`ent: validator failed for field "Cryptocurrency.symbol": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Icon(); ok {
		if err := cryptocurrency.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "Cryptocurrency.icon": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Name(); ok {
		if err := cryptocurrency.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Cryptocurrency.name": %w`, err)}
		}
	}
	if _, ok := cu.mutation.AssetID(); cu.mutation.AssetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Cryptocurrency.asset"`)
	}
	return nil
}

func (cu *CryptocurrencyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cryptocurrency.Table,
			Columns: cryptocurrency.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: cryptocurrency.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cryptocurrency.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cryptocurrency.FieldDeletedAt,
		})
	}
	if cu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: cryptocurrency.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.Symbol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cryptocurrency.FieldSymbol,
		})
	}
	if value, ok := cu.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cryptocurrency.FieldIcon,
		})
	}
	if cu.mutation.IconCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: cryptocurrency.FieldIcon,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cryptocurrency.FieldName,
		})
	}
	if cu.mutation.AssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   cryptocurrency.AssetTable,
			Columns: []string{cryptocurrency.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   cryptocurrency.AssetTable,
			Columns: []string{cryptocurrency.AssetColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.BlockchainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cryptocurrency.BlockchainsTable,
			Columns: cryptocurrency.BlockchainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: blockchain.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedBlockchainsIDs(); len(nodes) > 0 && !cu.mutation.BlockchainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cryptocurrency.BlockchainsTable,
			Columns: cryptocurrency.BlockchainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: blockchain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.BlockchainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cryptocurrency.BlockchainsTable,
			Columns: cryptocurrency.BlockchainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: blockchain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cryptocurrency.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CryptocurrencyUpdateOne is the builder for updating a single Cryptocurrency entity.
type CryptocurrencyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CryptocurrencyMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CryptocurrencyUpdateOne) SetUpdatedAt(t time.Time) *CryptocurrencyUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CryptocurrencyUpdateOne) SetDeletedAt(t time.Time) *CryptocurrencyUpdateOne {
	cuo.mutation.SetDeletedAt(t)
	return cuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cuo *CryptocurrencyUpdateOne) ClearDeletedAt() *CryptocurrencyUpdateOne {
	cuo.mutation.ClearDeletedAt()
	return cuo
}

// SetSymbol sets the "symbol" field.
func (cuo *CryptocurrencyUpdateOne) SetSymbol(s string) *CryptocurrencyUpdateOne {
	cuo.mutation.SetSymbol(s)
	return cuo
}

// SetIcon sets the "icon" field.
func (cuo *CryptocurrencyUpdateOne) SetIcon(s string) *CryptocurrencyUpdateOne {
	cuo.mutation.SetIcon(s)
	return cuo
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (cuo *CryptocurrencyUpdateOne) SetNillableIcon(s *string) *CryptocurrencyUpdateOne {
	if s != nil {
		cuo.SetIcon(*s)
	}
	return cuo
}

// ClearIcon clears the value of the "icon" field.
func (cuo *CryptocurrencyUpdateOne) ClearIcon() *CryptocurrencyUpdateOne {
	cuo.mutation.ClearIcon()
	return cuo
}

// SetName sets the "name" field.
func (cuo *CryptocurrencyUpdateOne) SetName(s string) *CryptocurrencyUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetAssetID sets the "asset_id" field.
func (cuo *CryptocurrencyUpdateOne) SetAssetID(pu pulid.PULID) *CryptocurrencyUpdateOne {
	cuo.mutation.SetAssetID(pu)
	return cuo
}

// SetAsset sets the "asset" edge to the Asset entity.
func (cuo *CryptocurrencyUpdateOne) SetAsset(a *Asset) *CryptocurrencyUpdateOne {
	return cuo.SetAssetID(a.ID)
}

// AddBlockchainIDs adds the "blockchains" edge to the Blockchain entity by IDs.
func (cuo *CryptocurrencyUpdateOne) AddBlockchainIDs(ids ...pulid.PULID) *CryptocurrencyUpdateOne {
	cuo.mutation.AddBlockchainIDs(ids...)
	return cuo
}

// AddBlockchains adds the "blockchains" edges to the Blockchain entity.
func (cuo *CryptocurrencyUpdateOne) AddBlockchains(b ...*Blockchain) *CryptocurrencyUpdateOne {
	ids := make([]pulid.PULID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.AddBlockchainIDs(ids...)
}

// Mutation returns the CryptocurrencyMutation object of the builder.
func (cuo *CryptocurrencyUpdateOne) Mutation() *CryptocurrencyMutation {
	return cuo.mutation
}

// ClearAsset clears the "asset" edge to the Asset entity.
func (cuo *CryptocurrencyUpdateOne) ClearAsset() *CryptocurrencyUpdateOne {
	cuo.mutation.ClearAsset()
	return cuo
}

// ClearBlockchains clears all "blockchains" edges to the Blockchain entity.
func (cuo *CryptocurrencyUpdateOne) ClearBlockchains() *CryptocurrencyUpdateOne {
	cuo.mutation.ClearBlockchains()
	return cuo
}

// RemoveBlockchainIDs removes the "blockchains" edge to Blockchain entities by IDs.
func (cuo *CryptocurrencyUpdateOne) RemoveBlockchainIDs(ids ...pulid.PULID) *CryptocurrencyUpdateOne {
	cuo.mutation.RemoveBlockchainIDs(ids...)
	return cuo
}

// RemoveBlockchains removes "blockchains" edges to Blockchain entities.
func (cuo *CryptocurrencyUpdateOne) RemoveBlockchains(b ...*Blockchain) *CryptocurrencyUpdateOne {
	ids := make([]pulid.PULID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.RemoveBlockchainIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CryptocurrencyUpdateOne) Select(field string, fields ...string) *CryptocurrencyUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cryptocurrency entity.
func (cuo *CryptocurrencyUpdateOne) Save(ctx context.Context) (*Cryptocurrency, error) {
	var (
		err  error
		node *Cryptocurrency
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CryptocurrencyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Cryptocurrency)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CryptocurrencyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CryptocurrencyUpdateOne) SaveX(ctx context.Context) *Cryptocurrency {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CryptocurrencyUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CryptocurrencyUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CryptocurrencyUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := cryptocurrency.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := cuo.mutation.DeletedAt(); !ok && !cuo.mutation.DeletedAtCleared() {
		v := cryptocurrency.UpdateDefaultDeletedAt()
		cuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CryptocurrencyUpdateOne) check() error {
	if v, ok := cuo.mutation.Symbol(); ok {
		if err := cryptocurrency.SymbolValidator(v); err != nil {
			return &ValidationError{Name: "symbol", err: fmt.Errorf(`ent: validator failed for field "Cryptocurrency.symbol": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Icon(); ok {
		if err := cryptocurrency.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "Cryptocurrency.icon": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Name(); ok {
		if err := cryptocurrency.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Cryptocurrency.name": %w`, err)}
		}
	}
	if _, ok := cuo.mutation.AssetID(); cuo.mutation.AssetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Cryptocurrency.asset"`)
	}
	return nil
}

func (cuo *CryptocurrencyUpdateOne) sqlSave(ctx context.Context) (_node *Cryptocurrency, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cryptocurrency.Table,
			Columns: cryptocurrency.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: cryptocurrency.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cryptocurrency.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cryptocurrency.FieldID)
		for _, f := range fields {
			if !cryptocurrency.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cryptocurrency.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cryptocurrency.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cryptocurrency.FieldDeletedAt,
		})
	}
	if cuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: cryptocurrency.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.Symbol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cryptocurrency.FieldSymbol,
		})
	}
	if value, ok := cuo.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cryptocurrency.FieldIcon,
		})
	}
	if cuo.mutation.IconCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: cryptocurrency.FieldIcon,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cryptocurrency.FieldName,
		})
	}
	if cuo.mutation.AssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   cryptocurrency.AssetTable,
			Columns: []string{cryptocurrency.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   cryptocurrency.AssetTable,
			Columns: []string{cryptocurrency.AssetColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.BlockchainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cryptocurrency.BlockchainsTable,
			Columns: cryptocurrency.BlockchainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: blockchain.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedBlockchainsIDs(); len(nodes) > 0 && !cuo.mutation.BlockchainsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cryptocurrency.BlockchainsTable,
			Columns: cryptocurrency.BlockchainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: blockchain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.BlockchainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cryptocurrency.BlockchainsTable,
			Columns: cryptocurrency.BlockchainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: blockchain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Cryptocurrency{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cryptocurrency.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
