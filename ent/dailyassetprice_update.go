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
	"github.com/chenningg/hermitboard-api/ent/dailyassetprice"
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/pulid"
)

// DailyAssetPriceUpdate is the builder for updating DailyAssetPrice entities.
type DailyAssetPriceUpdate struct {
	config
	hooks    []Hook
	mutation *DailyAssetPriceMutation
}

// Where appends a list predicates to the DailyAssetPriceUpdate builder.
func (dapu *DailyAssetPriceUpdate) Where(ps ...predicate.DailyAssetPrice) *DailyAssetPriceUpdate {
	dapu.mutation.Where(ps...)
	return dapu
}

// SetUpdatedAt sets the "updated_at" field.
func (dapu *DailyAssetPriceUpdate) SetUpdatedAt(t time.Time) *DailyAssetPriceUpdate {
	dapu.mutation.SetUpdatedAt(t)
	return dapu
}

// SetDeletedAt sets the "deleted_at" field.
func (dapu *DailyAssetPriceUpdate) SetDeletedAt(t time.Time) *DailyAssetPriceUpdate {
	dapu.mutation.SetDeletedAt(t)
	return dapu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (dapu *DailyAssetPriceUpdate) ClearDeletedAt() *DailyAssetPriceUpdate {
	dapu.mutation.ClearDeletedAt()
	return dapu
}

// SetTime sets the "time" field.
func (dapu *DailyAssetPriceUpdate) SetTime(t time.Time) *DailyAssetPriceUpdate {
	dapu.mutation.SetTime(t)
	return dapu
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (dapu *DailyAssetPriceUpdate) SetNillableTime(t *time.Time) *DailyAssetPriceUpdate {
	if t != nil {
		dapu.SetTime(*t)
	}
	return dapu
}

// SetOpen sets the "open" field.
func (dapu *DailyAssetPriceUpdate) SetOpen(f float64) *DailyAssetPriceUpdate {
	dapu.mutation.ResetOpen()
	dapu.mutation.SetOpen(f)
	return dapu
}

// SetNillableOpen sets the "open" field if the given value is not nil.
func (dapu *DailyAssetPriceUpdate) SetNillableOpen(f *float64) *DailyAssetPriceUpdate {
	if f != nil {
		dapu.SetOpen(*f)
	}
	return dapu
}

// AddOpen adds f to the "open" field.
func (dapu *DailyAssetPriceUpdate) AddOpen(f float64) *DailyAssetPriceUpdate {
	dapu.mutation.AddOpen(f)
	return dapu
}

// ClearOpen clears the value of the "open" field.
func (dapu *DailyAssetPriceUpdate) ClearOpen() *DailyAssetPriceUpdate {
	dapu.mutation.ClearOpen()
	return dapu
}

// SetHigh sets the "high" field.
func (dapu *DailyAssetPriceUpdate) SetHigh(f float64) *DailyAssetPriceUpdate {
	dapu.mutation.ResetHigh()
	dapu.mutation.SetHigh(f)
	return dapu
}

// SetNillableHigh sets the "high" field if the given value is not nil.
func (dapu *DailyAssetPriceUpdate) SetNillableHigh(f *float64) *DailyAssetPriceUpdate {
	if f != nil {
		dapu.SetHigh(*f)
	}
	return dapu
}

// AddHigh adds f to the "high" field.
func (dapu *DailyAssetPriceUpdate) AddHigh(f float64) *DailyAssetPriceUpdate {
	dapu.mutation.AddHigh(f)
	return dapu
}

// ClearHigh clears the value of the "high" field.
func (dapu *DailyAssetPriceUpdate) ClearHigh() *DailyAssetPriceUpdate {
	dapu.mutation.ClearHigh()
	return dapu
}

// SetLow sets the "low" field.
func (dapu *DailyAssetPriceUpdate) SetLow(f float64) *DailyAssetPriceUpdate {
	dapu.mutation.ResetLow()
	dapu.mutation.SetLow(f)
	return dapu
}

// SetNillableLow sets the "low" field if the given value is not nil.
func (dapu *DailyAssetPriceUpdate) SetNillableLow(f *float64) *DailyAssetPriceUpdate {
	if f != nil {
		dapu.SetLow(*f)
	}
	return dapu
}

// AddLow adds f to the "low" field.
func (dapu *DailyAssetPriceUpdate) AddLow(f float64) *DailyAssetPriceUpdate {
	dapu.mutation.AddLow(f)
	return dapu
}

// ClearLow clears the value of the "low" field.
func (dapu *DailyAssetPriceUpdate) ClearLow() *DailyAssetPriceUpdate {
	dapu.mutation.ClearLow()
	return dapu
}

// SetClose sets the "close" field.
func (dapu *DailyAssetPriceUpdate) SetClose(f float64) *DailyAssetPriceUpdate {
	dapu.mutation.ResetClose()
	dapu.mutation.SetClose(f)
	return dapu
}

// SetNillableClose sets the "close" field if the given value is not nil.
func (dapu *DailyAssetPriceUpdate) SetNillableClose(f *float64) *DailyAssetPriceUpdate {
	if f != nil {
		dapu.SetClose(*f)
	}
	return dapu
}

// AddClose adds f to the "close" field.
func (dapu *DailyAssetPriceUpdate) AddClose(f float64) *DailyAssetPriceUpdate {
	dapu.mutation.AddClose(f)
	return dapu
}

// ClearClose clears the value of the "close" field.
func (dapu *DailyAssetPriceUpdate) ClearClose() *DailyAssetPriceUpdate {
	dapu.mutation.ClearClose()
	return dapu
}

// SetAdjustedClose sets the "adjusted_close" field.
func (dapu *DailyAssetPriceUpdate) SetAdjustedClose(f float64) *DailyAssetPriceUpdate {
	dapu.mutation.ResetAdjustedClose()
	dapu.mutation.SetAdjustedClose(f)
	return dapu
}

// AddAdjustedClose adds f to the "adjusted_close" field.
func (dapu *DailyAssetPriceUpdate) AddAdjustedClose(f float64) *DailyAssetPriceUpdate {
	dapu.mutation.AddAdjustedClose(f)
	return dapu
}

// SetBaseAssetID sets the "base_asset_id" field.
func (dapu *DailyAssetPriceUpdate) SetBaseAssetID(pu pulid.PULID) *DailyAssetPriceUpdate {
	dapu.mutation.SetBaseAssetID(pu)
	return dapu
}

// SetBaseAsset sets the "base_asset" edge to the Asset entity.
func (dapu *DailyAssetPriceUpdate) SetBaseAsset(a *Asset) *DailyAssetPriceUpdate {
	return dapu.SetBaseAssetID(a.ID)
}

// Mutation returns the DailyAssetPriceMutation object of the builder.
func (dapu *DailyAssetPriceUpdate) Mutation() *DailyAssetPriceMutation {
	return dapu.mutation
}

// ClearBaseAsset clears the "base_asset" edge to the Asset entity.
func (dapu *DailyAssetPriceUpdate) ClearBaseAsset() *DailyAssetPriceUpdate {
	dapu.mutation.ClearBaseAsset()
	return dapu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dapu *DailyAssetPriceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	dapu.defaults()
	if len(dapu.hooks) == 0 {
		if err = dapu.check(); err != nil {
			return 0, err
		}
		affected, err = dapu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DailyAssetPriceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dapu.check(); err != nil {
				return 0, err
			}
			dapu.mutation = mutation
			affected, err = dapu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dapu.hooks) - 1; i >= 0; i-- {
			if dapu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dapu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dapu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dapu *DailyAssetPriceUpdate) SaveX(ctx context.Context) int {
	affected, err := dapu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dapu *DailyAssetPriceUpdate) Exec(ctx context.Context) error {
	_, err := dapu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dapu *DailyAssetPriceUpdate) ExecX(ctx context.Context) {
	if err := dapu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dapu *DailyAssetPriceUpdate) defaults() {
	if _, ok := dapu.mutation.UpdatedAt(); !ok {
		v := dailyassetprice.UpdateDefaultUpdatedAt()
		dapu.mutation.SetUpdatedAt(v)
	}
	if _, ok := dapu.mutation.DeletedAt(); !ok && !dapu.mutation.DeletedAtCleared() {
		v := dailyassetprice.UpdateDefaultDeletedAt()
		dapu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dapu *DailyAssetPriceUpdate) check() error {
	if _, ok := dapu.mutation.BaseAssetID(); dapu.mutation.BaseAssetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DailyAssetPrice.base_asset"`)
	}
	return nil
}

func (dapu *DailyAssetPriceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dailyassetprice.Table,
			Columns: dailyassetprice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: dailyassetprice.FieldID,
			},
		},
	}
	if ps := dapu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dapu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailyassetprice.FieldUpdatedAt,
		})
	}
	if value, ok := dapu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailyassetprice.FieldDeletedAt,
		})
	}
	if dapu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: dailyassetprice.FieldDeletedAt,
		})
	}
	if value, ok := dapu.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailyassetprice.FieldTime,
		})
	}
	if value, ok := dapu.mutation.Open(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldOpen,
		})
	}
	if value, ok := dapu.mutation.AddedOpen(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldOpen,
		})
	}
	if dapu.mutation.OpenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: dailyassetprice.FieldOpen,
		})
	}
	if value, ok := dapu.mutation.High(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldHigh,
		})
	}
	if value, ok := dapu.mutation.AddedHigh(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldHigh,
		})
	}
	if dapu.mutation.HighCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: dailyassetprice.FieldHigh,
		})
	}
	if value, ok := dapu.mutation.Low(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldLow,
		})
	}
	if value, ok := dapu.mutation.AddedLow(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldLow,
		})
	}
	if dapu.mutation.LowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: dailyassetprice.FieldLow,
		})
	}
	if value, ok := dapu.mutation.Close(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldClose,
		})
	}
	if value, ok := dapu.mutation.AddedClose(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldClose,
		})
	}
	if dapu.mutation.CloseCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: dailyassetprice.FieldClose,
		})
	}
	if value, ok := dapu.mutation.AdjustedClose(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldAdjustedClose,
		})
	}
	if value, ok := dapu.mutation.AddedAdjustedClose(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldAdjustedClose,
		})
	}
	if dapu.mutation.BaseAssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyassetprice.BaseAssetTable,
			Columns: []string{dailyassetprice.BaseAssetColumn},
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
	if nodes := dapu.mutation.BaseAssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyassetprice.BaseAssetTable,
			Columns: []string{dailyassetprice.BaseAssetColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, dapu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dailyassetprice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DailyAssetPriceUpdateOne is the builder for updating a single DailyAssetPrice entity.
type DailyAssetPriceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DailyAssetPriceMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (dapuo *DailyAssetPriceUpdateOne) SetUpdatedAt(t time.Time) *DailyAssetPriceUpdateOne {
	dapuo.mutation.SetUpdatedAt(t)
	return dapuo
}

// SetDeletedAt sets the "deleted_at" field.
func (dapuo *DailyAssetPriceUpdateOne) SetDeletedAt(t time.Time) *DailyAssetPriceUpdateOne {
	dapuo.mutation.SetDeletedAt(t)
	return dapuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (dapuo *DailyAssetPriceUpdateOne) ClearDeletedAt() *DailyAssetPriceUpdateOne {
	dapuo.mutation.ClearDeletedAt()
	return dapuo
}

// SetTime sets the "time" field.
func (dapuo *DailyAssetPriceUpdateOne) SetTime(t time.Time) *DailyAssetPriceUpdateOne {
	dapuo.mutation.SetTime(t)
	return dapuo
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (dapuo *DailyAssetPriceUpdateOne) SetNillableTime(t *time.Time) *DailyAssetPriceUpdateOne {
	if t != nil {
		dapuo.SetTime(*t)
	}
	return dapuo
}

// SetOpen sets the "open" field.
func (dapuo *DailyAssetPriceUpdateOne) SetOpen(f float64) *DailyAssetPriceUpdateOne {
	dapuo.mutation.ResetOpen()
	dapuo.mutation.SetOpen(f)
	return dapuo
}

// SetNillableOpen sets the "open" field if the given value is not nil.
func (dapuo *DailyAssetPriceUpdateOne) SetNillableOpen(f *float64) *DailyAssetPriceUpdateOne {
	if f != nil {
		dapuo.SetOpen(*f)
	}
	return dapuo
}

// AddOpen adds f to the "open" field.
func (dapuo *DailyAssetPriceUpdateOne) AddOpen(f float64) *DailyAssetPriceUpdateOne {
	dapuo.mutation.AddOpen(f)
	return dapuo
}

// ClearOpen clears the value of the "open" field.
func (dapuo *DailyAssetPriceUpdateOne) ClearOpen() *DailyAssetPriceUpdateOne {
	dapuo.mutation.ClearOpen()
	return dapuo
}

// SetHigh sets the "high" field.
func (dapuo *DailyAssetPriceUpdateOne) SetHigh(f float64) *DailyAssetPriceUpdateOne {
	dapuo.mutation.ResetHigh()
	dapuo.mutation.SetHigh(f)
	return dapuo
}

// SetNillableHigh sets the "high" field if the given value is not nil.
func (dapuo *DailyAssetPriceUpdateOne) SetNillableHigh(f *float64) *DailyAssetPriceUpdateOne {
	if f != nil {
		dapuo.SetHigh(*f)
	}
	return dapuo
}

// AddHigh adds f to the "high" field.
func (dapuo *DailyAssetPriceUpdateOne) AddHigh(f float64) *DailyAssetPriceUpdateOne {
	dapuo.mutation.AddHigh(f)
	return dapuo
}

// ClearHigh clears the value of the "high" field.
func (dapuo *DailyAssetPriceUpdateOne) ClearHigh() *DailyAssetPriceUpdateOne {
	dapuo.mutation.ClearHigh()
	return dapuo
}

// SetLow sets the "low" field.
func (dapuo *DailyAssetPriceUpdateOne) SetLow(f float64) *DailyAssetPriceUpdateOne {
	dapuo.mutation.ResetLow()
	dapuo.mutation.SetLow(f)
	return dapuo
}

// SetNillableLow sets the "low" field if the given value is not nil.
func (dapuo *DailyAssetPriceUpdateOne) SetNillableLow(f *float64) *DailyAssetPriceUpdateOne {
	if f != nil {
		dapuo.SetLow(*f)
	}
	return dapuo
}

// AddLow adds f to the "low" field.
func (dapuo *DailyAssetPriceUpdateOne) AddLow(f float64) *DailyAssetPriceUpdateOne {
	dapuo.mutation.AddLow(f)
	return dapuo
}

// ClearLow clears the value of the "low" field.
func (dapuo *DailyAssetPriceUpdateOne) ClearLow() *DailyAssetPriceUpdateOne {
	dapuo.mutation.ClearLow()
	return dapuo
}

// SetClose sets the "close" field.
func (dapuo *DailyAssetPriceUpdateOne) SetClose(f float64) *DailyAssetPriceUpdateOne {
	dapuo.mutation.ResetClose()
	dapuo.mutation.SetClose(f)
	return dapuo
}

// SetNillableClose sets the "close" field if the given value is not nil.
func (dapuo *DailyAssetPriceUpdateOne) SetNillableClose(f *float64) *DailyAssetPriceUpdateOne {
	if f != nil {
		dapuo.SetClose(*f)
	}
	return dapuo
}

// AddClose adds f to the "close" field.
func (dapuo *DailyAssetPriceUpdateOne) AddClose(f float64) *DailyAssetPriceUpdateOne {
	dapuo.mutation.AddClose(f)
	return dapuo
}

// ClearClose clears the value of the "close" field.
func (dapuo *DailyAssetPriceUpdateOne) ClearClose() *DailyAssetPriceUpdateOne {
	dapuo.mutation.ClearClose()
	return dapuo
}

// SetAdjustedClose sets the "adjusted_close" field.
func (dapuo *DailyAssetPriceUpdateOne) SetAdjustedClose(f float64) *DailyAssetPriceUpdateOne {
	dapuo.mutation.ResetAdjustedClose()
	dapuo.mutation.SetAdjustedClose(f)
	return dapuo
}

// AddAdjustedClose adds f to the "adjusted_close" field.
func (dapuo *DailyAssetPriceUpdateOne) AddAdjustedClose(f float64) *DailyAssetPriceUpdateOne {
	dapuo.mutation.AddAdjustedClose(f)
	return dapuo
}

// SetBaseAssetID sets the "base_asset_id" field.
func (dapuo *DailyAssetPriceUpdateOne) SetBaseAssetID(pu pulid.PULID) *DailyAssetPriceUpdateOne {
	dapuo.mutation.SetBaseAssetID(pu)
	return dapuo
}

// SetBaseAsset sets the "base_asset" edge to the Asset entity.
func (dapuo *DailyAssetPriceUpdateOne) SetBaseAsset(a *Asset) *DailyAssetPriceUpdateOne {
	return dapuo.SetBaseAssetID(a.ID)
}

// Mutation returns the DailyAssetPriceMutation object of the builder.
func (dapuo *DailyAssetPriceUpdateOne) Mutation() *DailyAssetPriceMutation {
	return dapuo.mutation
}

// ClearBaseAsset clears the "base_asset" edge to the Asset entity.
func (dapuo *DailyAssetPriceUpdateOne) ClearBaseAsset() *DailyAssetPriceUpdateOne {
	dapuo.mutation.ClearBaseAsset()
	return dapuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dapuo *DailyAssetPriceUpdateOne) Select(field string, fields ...string) *DailyAssetPriceUpdateOne {
	dapuo.fields = append([]string{field}, fields...)
	return dapuo
}

// Save executes the query and returns the updated DailyAssetPrice entity.
func (dapuo *DailyAssetPriceUpdateOne) Save(ctx context.Context) (*DailyAssetPrice, error) {
	var (
		err  error
		node *DailyAssetPrice
	)
	dapuo.defaults()
	if len(dapuo.hooks) == 0 {
		if err = dapuo.check(); err != nil {
			return nil, err
		}
		node, err = dapuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DailyAssetPriceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dapuo.check(); err != nil {
				return nil, err
			}
			dapuo.mutation = mutation
			node, err = dapuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dapuo.hooks) - 1; i >= 0; i-- {
			if dapuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dapuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dapuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DailyAssetPrice)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DailyAssetPriceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dapuo *DailyAssetPriceUpdateOne) SaveX(ctx context.Context) *DailyAssetPrice {
	node, err := dapuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dapuo *DailyAssetPriceUpdateOne) Exec(ctx context.Context) error {
	_, err := dapuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dapuo *DailyAssetPriceUpdateOne) ExecX(ctx context.Context) {
	if err := dapuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dapuo *DailyAssetPriceUpdateOne) defaults() {
	if _, ok := dapuo.mutation.UpdatedAt(); !ok {
		v := dailyassetprice.UpdateDefaultUpdatedAt()
		dapuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := dapuo.mutation.DeletedAt(); !ok && !dapuo.mutation.DeletedAtCleared() {
		v := dailyassetprice.UpdateDefaultDeletedAt()
		dapuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dapuo *DailyAssetPriceUpdateOne) check() error {
	if _, ok := dapuo.mutation.BaseAssetID(); dapuo.mutation.BaseAssetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DailyAssetPrice.base_asset"`)
	}
	return nil
}

func (dapuo *DailyAssetPriceUpdateOne) sqlSave(ctx context.Context) (_node *DailyAssetPrice, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dailyassetprice.Table,
			Columns: dailyassetprice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: dailyassetprice.FieldID,
			},
		},
	}
	id, ok := dapuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DailyAssetPrice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dapuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dailyassetprice.FieldID)
		for _, f := range fields {
			if !dailyassetprice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dailyassetprice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dapuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dapuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailyassetprice.FieldUpdatedAt,
		})
	}
	if value, ok := dapuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailyassetprice.FieldDeletedAt,
		})
	}
	if dapuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: dailyassetprice.FieldDeletedAt,
		})
	}
	if value, ok := dapuo.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dailyassetprice.FieldTime,
		})
	}
	if value, ok := dapuo.mutation.Open(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldOpen,
		})
	}
	if value, ok := dapuo.mutation.AddedOpen(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldOpen,
		})
	}
	if dapuo.mutation.OpenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: dailyassetprice.FieldOpen,
		})
	}
	if value, ok := dapuo.mutation.High(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldHigh,
		})
	}
	if value, ok := dapuo.mutation.AddedHigh(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldHigh,
		})
	}
	if dapuo.mutation.HighCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: dailyassetprice.FieldHigh,
		})
	}
	if value, ok := dapuo.mutation.Low(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldLow,
		})
	}
	if value, ok := dapuo.mutation.AddedLow(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldLow,
		})
	}
	if dapuo.mutation.LowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: dailyassetprice.FieldLow,
		})
	}
	if value, ok := dapuo.mutation.Close(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldClose,
		})
	}
	if value, ok := dapuo.mutation.AddedClose(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldClose,
		})
	}
	if dapuo.mutation.CloseCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: dailyassetprice.FieldClose,
		})
	}
	if value, ok := dapuo.mutation.AdjustedClose(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldAdjustedClose,
		})
	}
	if value, ok := dapuo.mutation.AddedAdjustedClose(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: dailyassetprice.FieldAdjustedClose,
		})
	}
	if dapuo.mutation.BaseAssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyassetprice.BaseAssetTable,
			Columns: []string{dailyassetprice.BaseAssetColumn},
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
	if nodes := dapuo.mutation.BaseAssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyassetprice.BaseAssetTable,
			Columns: []string{dailyassetprice.BaseAssetColumn},
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
	_node = &DailyAssetPrice{config: dapuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dapuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dailyassetprice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}