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
	"github.com/chenningg/hermitboard-api/ent/assetclass"
	"github.com/chenningg/hermitboard-api/ent/cryptocurrency"
	"github.com/chenningg/hermitboard-api/ent/dailyassetprice"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/ent/transaction"
	"github.com/oklog/ulid/v2"
	ulid "github.com/oklog/ulid/v2"
)

// AssetUpdate is the builder for updating Asset entities.
type AssetUpdate struct {
	config
	hooks    []Hook
	mutation *AssetMutation
}

// Where appends a list predicates to the AssetUpdate builder.
func (au *AssetUpdate) Where(ps ...predicate.Asset) *AssetUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AssetUpdate) SetUpdatedAt(t time.Time) *AssetUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetAssetClassID sets the "asset_class" edge to the AssetClass entity by ID.
func (au *AssetUpdate) SetAssetClassID(id ulid.ULID) *AssetUpdate {
	au.mutation.SetAssetClassID(id)
	return au
}

// SetAssetClass sets the "asset_class" edge to the AssetClass entity.
func (au *AssetUpdate) SetAssetClass(a *AssetClass) *AssetUpdate {
	return au.SetAssetClassID(a.ID)
}

// SetCryptocurrencyID sets the "cryptocurrency" edge to the Cryptocurrency entity by ID.
func (au *AssetUpdate) SetCryptocurrencyID(id ulid.ULID) *AssetUpdate {
	au.mutation.SetCryptocurrencyID(id)
	return au
}

// SetNillableCryptocurrencyID sets the "cryptocurrency" edge to the Cryptocurrency entity by ID if the given value is not nil.
func (au *AssetUpdate) SetNillableCryptocurrencyID(id *ulid.ULID) *AssetUpdate {
	if id != nil {
		au = au.SetCryptocurrencyID(*id)
	}
	return au
}

// SetCryptocurrency sets the "cryptocurrency" edge to the Cryptocurrency entity.
func (au *AssetUpdate) SetCryptocurrency(c *Cryptocurrency) *AssetUpdate {
	return au.SetCryptocurrencyID(c.ID)
}

// AddTransactionBaseIDs adds the "transaction_base" edge to the Transaction entity by IDs.
func (au *AssetUpdate) AddTransactionBaseIDs(ids ...ulid.ULID) *AssetUpdate {
	au.mutation.AddTransactionBaseIDs(ids...)
	return au
}

// AddTransactionBase adds the "transaction_base" edges to the Transaction entity.
func (au *AssetUpdate) AddTransactionBase(t ...*Transaction) *AssetUpdate {
	ids := make([]ulid.ULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTransactionBaseIDs(ids...)
}

// AddTransactionQuoteIDs adds the "transaction_quote" edge to the Transaction entity by IDs.
func (au *AssetUpdate) AddTransactionQuoteIDs(ids ...ulid.ULID) *AssetUpdate {
	au.mutation.AddTransactionQuoteIDs(ids...)
	return au
}

// AddTransactionQuote adds the "transaction_quote" edges to the Transaction entity.
func (au *AssetUpdate) AddTransactionQuote(t ...*Transaction) *AssetUpdate {
	ids := make([]ulid.ULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTransactionQuoteIDs(ids...)
}

// AddDailyAssetPriceIDs adds the "daily_asset_price" edge to the DailyAssetPrice entity by IDs.
func (au *AssetUpdate) AddDailyAssetPriceIDs(ids ...ulid.ULID) *AssetUpdate {
	au.mutation.AddDailyAssetPriceIDs(ids...)
	return au
}

// AddDailyAssetPrice adds the "daily_asset_price" edges to the DailyAssetPrice entity.
func (au *AssetUpdate) AddDailyAssetPrice(d ...*DailyAssetPrice) *AssetUpdate {
	ids := make([]ulid.ULID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return au.AddDailyAssetPriceIDs(ids...)
}

// Mutation returns the AssetMutation object of the builder.
func (au *AssetUpdate) Mutation() *AssetMutation {
	return au.mutation
}

// ClearAssetClass clears the "asset_class" edge to the AssetClass entity.
func (au *AssetUpdate) ClearAssetClass() *AssetUpdate {
	au.mutation.ClearAssetClass()
	return au
}

// ClearCryptocurrency clears the "cryptocurrency" edge to the Cryptocurrency entity.
func (au *AssetUpdate) ClearCryptocurrency() *AssetUpdate {
	au.mutation.ClearCryptocurrency()
	return au
}

// ClearTransactionBase clears all "transaction_base" edges to the Transaction entity.
func (au *AssetUpdate) ClearTransactionBase() *AssetUpdate {
	au.mutation.ClearTransactionBase()
	return au
}

// RemoveTransactionBaseIDs removes the "transaction_base" edge to Transaction entities by IDs.
func (au *AssetUpdate) RemoveTransactionBaseIDs(ids ...ulid.ULID) *AssetUpdate {
	au.mutation.RemoveTransactionBaseIDs(ids...)
	return au
}

// RemoveTransactionBase removes "transaction_base" edges to Transaction entities.
func (au *AssetUpdate) RemoveTransactionBase(t ...*Transaction) *AssetUpdate {
	ids := make([]ulid.ULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTransactionBaseIDs(ids...)
}

// ClearTransactionQuote clears all "transaction_quote" edges to the Transaction entity.
func (au *AssetUpdate) ClearTransactionQuote() *AssetUpdate {
	au.mutation.ClearTransactionQuote()
	return au
}

// RemoveTransactionQuoteIDs removes the "transaction_quote" edge to Transaction entities by IDs.
func (au *AssetUpdate) RemoveTransactionQuoteIDs(ids ...ulid.ULID) *AssetUpdate {
	au.mutation.RemoveTransactionQuoteIDs(ids...)
	return au
}

// RemoveTransactionQuote removes "transaction_quote" edges to Transaction entities.
func (au *AssetUpdate) RemoveTransactionQuote(t ...*Transaction) *AssetUpdate {
	ids := make([]ulid.ULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTransactionQuoteIDs(ids...)
}

// ClearDailyAssetPrice clears all "daily_asset_price" edges to the DailyAssetPrice entity.
func (au *AssetUpdate) ClearDailyAssetPrice() *AssetUpdate {
	au.mutation.ClearDailyAssetPrice()
	return au
}

// RemoveDailyAssetPriceIDs removes the "daily_asset_price" edge to DailyAssetPrice entities by IDs.
func (au *AssetUpdate) RemoveDailyAssetPriceIDs(ids ...ulid.ULID) *AssetUpdate {
	au.mutation.RemoveDailyAssetPriceIDs(ids...)
	return au
}

// RemoveDailyAssetPrice removes "daily_asset_price" edges to DailyAssetPrice entities.
func (au *AssetUpdate) RemoveDailyAssetPrice(d ...*DailyAssetPrice) *AssetUpdate {
	ids := make([]ulid.ULID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return au.RemoveDailyAssetPriceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AssetUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*AssetMutation)
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
func (au *AssetUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AssetUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AssetUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AssetUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := asset.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AssetUpdate) check() error {
	if _, ok := au.mutation.AssetClassID(); au.mutation.AssetClassCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Asset.asset_class"`)
	}
	return nil
}

func (au *AssetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asset.Table,
			Columns: asset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: asset.FieldID,
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
			Column: asset.FieldUpdatedAt,
		})
	}
	if au.mutation.AssetClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   asset.AssetClassTable,
			Columns: []string{asset.AssetClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: assetclass.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AssetClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   asset.AssetClassTable,
			Columns: []string{asset.AssetClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: assetclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.CryptocurrencyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   asset.CryptocurrencyTable,
			Columns: []string{asset.CryptocurrencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: cryptocurrency.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.CryptocurrencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   asset.CryptocurrencyTable,
			Columns: []string{asset.CryptocurrencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: cryptocurrency.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.TransactionBaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionBaseTable,
			Columns: asset.TransactionBasePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTransactionBaseIDs(); len(nodes) > 0 && !au.mutation.TransactionBaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionBaseTable,
			Columns: asset.TransactionBasePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TransactionBaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionBaseTable,
			Columns: asset.TransactionBasePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.TransactionQuoteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionQuoteTable,
			Columns: asset.TransactionQuotePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTransactionQuoteIDs(); len(nodes) > 0 && !au.mutation.TransactionQuoteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionQuoteTable,
			Columns: asset.TransactionQuotePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TransactionQuoteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionQuoteTable,
			Columns: asset.TransactionQuotePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.DailyAssetPriceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.DailyAssetPriceTable,
			Columns: []string{asset.DailyAssetPriceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dailyassetprice.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedDailyAssetPriceIDs(); len(nodes) > 0 && !au.mutation.DailyAssetPriceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.DailyAssetPriceTable,
			Columns: []string{asset.DailyAssetPriceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dailyassetprice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.DailyAssetPriceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.DailyAssetPriceTable,
			Columns: []string{asset.DailyAssetPriceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dailyassetprice.FieldID,
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
			err = &NotFoundError{asset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AssetUpdateOne is the builder for updating a single Asset entity.
type AssetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AssetMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AssetUpdateOne) SetUpdatedAt(t time.Time) *AssetUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetAssetClassID sets the "asset_class" edge to the AssetClass entity by ID.
func (auo *AssetUpdateOne) SetAssetClassID(id ulid.ULID) *AssetUpdateOne {
	auo.mutation.SetAssetClassID(id)
	return auo
}

// SetAssetClass sets the "asset_class" edge to the AssetClass entity.
func (auo *AssetUpdateOne) SetAssetClass(a *AssetClass) *AssetUpdateOne {
	return auo.SetAssetClassID(a.ID)
}

// SetCryptocurrencyID sets the "cryptocurrency" edge to the Cryptocurrency entity by ID.
func (auo *AssetUpdateOne) SetCryptocurrencyID(id ulid.ULID) *AssetUpdateOne {
	auo.mutation.SetCryptocurrencyID(id)
	return auo
}

// SetNillableCryptocurrencyID sets the "cryptocurrency" edge to the Cryptocurrency entity by ID if the given value is not nil.
func (auo *AssetUpdateOne) SetNillableCryptocurrencyID(id *ulid.ULID) *AssetUpdateOne {
	if id != nil {
		auo = auo.SetCryptocurrencyID(*id)
	}
	return auo
}

// SetCryptocurrency sets the "cryptocurrency" edge to the Cryptocurrency entity.
func (auo *AssetUpdateOne) SetCryptocurrency(c *Cryptocurrency) *AssetUpdateOne {
	return auo.SetCryptocurrencyID(c.ID)
}

// AddTransactionBaseIDs adds the "transaction_base" edge to the Transaction entity by IDs.
func (auo *AssetUpdateOne) AddTransactionBaseIDs(ids ...ulid.ULID) *AssetUpdateOne {
	auo.mutation.AddTransactionBaseIDs(ids...)
	return auo
}

// AddTransactionBase adds the "transaction_base" edges to the Transaction entity.
func (auo *AssetUpdateOne) AddTransactionBase(t ...*Transaction) *AssetUpdateOne {
	ids := make([]ulid.ULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTransactionBaseIDs(ids...)
}

// AddTransactionQuoteIDs adds the "transaction_quote" edge to the Transaction entity by IDs.
func (auo *AssetUpdateOne) AddTransactionQuoteIDs(ids ...ulid.ULID) *AssetUpdateOne {
	auo.mutation.AddTransactionQuoteIDs(ids...)
	return auo
}

// AddTransactionQuote adds the "transaction_quote" edges to the Transaction entity.
func (auo *AssetUpdateOne) AddTransactionQuote(t ...*Transaction) *AssetUpdateOne {
	ids := make([]ulid.ULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTransactionQuoteIDs(ids...)
}

// AddDailyAssetPriceIDs adds the "daily_asset_price" edge to the DailyAssetPrice entity by IDs.
func (auo *AssetUpdateOne) AddDailyAssetPriceIDs(ids ...ulid.ULID) *AssetUpdateOne {
	auo.mutation.AddDailyAssetPriceIDs(ids...)
	return auo
}

// AddDailyAssetPrice adds the "daily_asset_price" edges to the DailyAssetPrice entity.
func (auo *AssetUpdateOne) AddDailyAssetPrice(d ...*DailyAssetPrice) *AssetUpdateOne {
	ids := make([]ulid.ULID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return auo.AddDailyAssetPriceIDs(ids...)
}

// Mutation returns the AssetMutation object of the builder.
func (auo *AssetUpdateOne) Mutation() *AssetMutation {
	return auo.mutation
}

// ClearAssetClass clears the "asset_class" edge to the AssetClass entity.
func (auo *AssetUpdateOne) ClearAssetClass() *AssetUpdateOne {
	auo.mutation.ClearAssetClass()
	return auo
}

// ClearCryptocurrency clears the "cryptocurrency" edge to the Cryptocurrency entity.
func (auo *AssetUpdateOne) ClearCryptocurrency() *AssetUpdateOne {
	auo.mutation.ClearCryptocurrency()
	return auo
}

// ClearTransactionBase clears all "transaction_base" edges to the Transaction entity.
func (auo *AssetUpdateOne) ClearTransactionBase() *AssetUpdateOne {
	auo.mutation.ClearTransactionBase()
	return auo
}

// RemoveTransactionBaseIDs removes the "transaction_base" edge to Transaction entities by IDs.
func (auo *AssetUpdateOne) RemoveTransactionBaseIDs(ids ...ulid.ULID) *AssetUpdateOne {
	auo.mutation.RemoveTransactionBaseIDs(ids...)
	return auo
}

// RemoveTransactionBase removes "transaction_base" edges to Transaction entities.
func (auo *AssetUpdateOne) RemoveTransactionBase(t ...*Transaction) *AssetUpdateOne {
	ids := make([]ulid.ULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTransactionBaseIDs(ids...)
}

// ClearTransactionQuote clears all "transaction_quote" edges to the Transaction entity.
func (auo *AssetUpdateOne) ClearTransactionQuote() *AssetUpdateOne {
	auo.mutation.ClearTransactionQuote()
	return auo
}

// RemoveTransactionQuoteIDs removes the "transaction_quote" edge to Transaction entities by IDs.
func (auo *AssetUpdateOne) RemoveTransactionQuoteIDs(ids ...ulid.ULID) *AssetUpdateOne {
	auo.mutation.RemoveTransactionQuoteIDs(ids...)
	return auo
}

// RemoveTransactionQuote removes "transaction_quote" edges to Transaction entities.
func (auo *AssetUpdateOne) RemoveTransactionQuote(t ...*Transaction) *AssetUpdateOne {
	ids := make([]ulid.ULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTransactionQuoteIDs(ids...)
}

// ClearDailyAssetPrice clears all "daily_asset_price" edges to the DailyAssetPrice entity.
func (auo *AssetUpdateOne) ClearDailyAssetPrice() *AssetUpdateOne {
	auo.mutation.ClearDailyAssetPrice()
	return auo
}

// RemoveDailyAssetPriceIDs removes the "daily_asset_price" edge to DailyAssetPrice entities by IDs.
func (auo *AssetUpdateOne) RemoveDailyAssetPriceIDs(ids ...ulid.ULID) *AssetUpdateOne {
	auo.mutation.RemoveDailyAssetPriceIDs(ids...)
	return auo
}

// RemoveDailyAssetPrice removes "daily_asset_price" edges to DailyAssetPrice entities.
func (auo *AssetUpdateOne) RemoveDailyAssetPrice(d ...*DailyAssetPrice) *AssetUpdateOne {
	ids := make([]ulid.ULID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return auo.RemoveDailyAssetPriceIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AssetUpdateOne) Select(field string, fields ...string) *AssetUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Asset entity.
func (auo *AssetUpdateOne) Save(ctx context.Context) (*Asset, error) {
	var (
		err  error
		node *Asset
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetMutation)
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
		nv, ok := v.(*Asset)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AssetMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AssetUpdateOne) SaveX(ctx context.Context) *Asset {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AssetUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AssetUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AssetUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := asset.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AssetUpdateOne) check() error {
	if _, ok := auo.mutation.AssetClassID(); auo.mutation.AssetClassCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Asset.asset_class"`)
	}
	return nil
}

func (auo *AssetUpdateOne) sqlSave(ctx context.Context) (_node *Asset, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asset.Table,
			Columns: asset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: asset.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Asset.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asset.FieldID)
		for _, f := range fields {
			if !asset.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != asset.FieldID {
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
			Column: asset.FieldUpdatedAt,
		})
	}
	if auo.mutation.AssetClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   asset.AssetClassTable,
			Columns: []string{asset.AssetClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: assetclass.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AssetClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   asset.AssetClassTable,
			Columns: []string{asset.AssetClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: assetclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.CryptocurrencyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   asset.CryptocurrencyTable,
			Columns: []string{asset.CryptocurrencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: cryptocurrency.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.CryptocurrencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   asset.CryptocurrencyTable,
			Columns: []string{asset.CryptocurrencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: cryptocurrency.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.TransactionBaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionBaseTable,
			Columns: asset.TransactionBasePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTransactionBaseIDs(); len(nodes) > 0 && !auo.mutation.TransactionBaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionBaseTable,
			Columns: asset.TransactionBasePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TransactionBaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionBaseTable,
			Columns: asset.TransactionBasePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.TransactionQuoteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionQuoteTable,
			Columns: asset.TransactionQuotePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTransactionQuoteIDs(); len(nodes) > 0 && !auo.mutation.TransactionQuoteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionQuoteTable,
			Columns: asset.TransactionQuotePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TransactionQuoteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.TransactionQuoteTable,
			Columns: asset.TransactionQuotePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.DailyAssetPriceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.DailyAssetPriceTable,
			Columns: []string{asset.DailyAssetPriceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dailyassetprice.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedDailyAssetPriceIDs(); len(nodes) > 0 && !auo.mutation.DailyAssetPriceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.DailyAssetPriceTable,
			Columns: []string{asset.DailyAssetPriceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dailyassetprice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.DailyAssetPriceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asset.DailyAssetPriceTable,
			Columns: []string{asset.DailyAssetPriceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dailyassetprice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Asset{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
