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
	"github.com/chenningg/hermitboard-api/ent/predicate"
	"github.com/chenningg/hermitboard-api/ent/transaction"
	"github.com/chenningg/hermitboard-api/ent/transactiontype"
	"github.com/chenningg/hermitboard-api/pulid"
)

// TransactionTypeUpdate is the builder for updating TransactionType entities.
type TransactionTypeUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionTypeMutation
}

// Where appends a list predicates to the TransactionTypeUpdate builder.
func (ttu *TransactionTypeUpdate) Where(ps ...predicate.TransactionType) *TransactionTypeUpdate {
	ttu.mutation.Where(ps...)
	return ttu
}

// SetUpdatedAt sets the "updated_at" field.
func (ttu *TransactionTypeUpdate) SetUpdatedAt(t time.Time) *TransactionTypeUpdate {
	ttu.mutation.SetUpdatedAt(t)
	return ttu
}

// SetDeletedAt sets the "deleted_at" field.
func (ttu *TransactionTypeUpdate) SetDeletedAt(t time.Time) *TransactionTypeUpdate {
	ttu.mutation.SetDeletedAt(t)
	return ttu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ttu *TransactionTypeUpdate) ClearDeletedAt() *TransactionTypeUpdate {
	ttu.mutation.ClearDeletedAt()
	return ttu
}

// SetValue sets the "value" field.
func (ttu *TransactionTypeUpdate) SetValue(t transactiontype.Value) *TransactionTypeUpdate {
	ttu.mutation.SetValue(t)
	return ttu
}

// SetDescription sets the "description" field.
func (ttu *TransactionTypeUpdate) SetDescription(s string) *TransactionTypeUpdate {
	ttu.mutation.SetDescription(s)
	return ttu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ttu *TransactionTypeUpdate) SetNillableDescription(s *string) *TransactionTypeUpdate {
	if s != nil {
		ttu.SetDescription(*s)
	}
	return ttu
}

// ClearDescription clears the value of the "description" field.
func (ttu *TransactionTypeUpdate) ClearDescription() *TransactionTypeUpdate {
	ttu.mutation.ClearDescription()
	return ttu
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (ttu *TransactionTypeUpdate) AddTransactionIDs(ids ...pulid.PULID) *TransactionTypeUpdate {
	ttu.mutation.AddTransactionIDs(ids...)
	return ttu
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (ttu *TransactionTypeUpdate) AddTransactions(t ...*Transaction) *TransactionTypeUpdate {
	ids := make([]pulid.PULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ttu.AddTransactionIDs(ids...)
}

// Mutation returns the TransactionTypeMutation object of the builder.
func (ttu *TransactionTypeUpdate) Mutation() *TransactionTypeMutation {
	return ttu.mutation
}

// ClearTransactions clears all "transactions" edges to the Transaction entity.
func (ttu *TransactionTypeUpdate) ClearTransactions() *TransactionTypeUpdate {
	ttu.mutation.ClearTransactions()
	return ttu
}

// RemoveTransactionIDs removes the "transactions" edge to Transaction entities by IDs.
func (ttu *TransactionTypeUpdate) RemoveTransactionIDs(ids ...pulid.PULID) *TransactionTypeUpdate {
	ttu.mutation.RemoveTransactionIDs(ids...)
	return ttu
}

// RemoveTransactions removes "transactions" edges to Transaction entities.
func (ttu *TransactionTypeUpdate) RemoveTransactions(t ...*Transaction) *TransactionTypeUpdate {
	ids := make([]pulid.PULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ttu.RemoveTransactionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ttu *TransactionTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ttu.defaults()
	if len(ttu.hooks) == 0 {
		if err = ttu.check(); err != nil {
			return 0, err
		}
		affected, err = ttu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttu.check(); err != nil {
				return 0, err
			}
			ttu.mutation = mutation
			affected, err = ttu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttu.hooks) - 1; i >= 0; i-- {
			if ttu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttu *TransactionTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ttu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ttu *TransactionTypeUpdate) Exec(ctx context.Context) error {
	_, err := ttu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttu *TransactionTypeUpdate) ExecX(ctx context.Context) {
	if err := ttu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttu *TransactionTypeUpdate) defaults() {
	if _, ok := ttu.mutation.UpdatedAt(); !ok {
		v := transactiontype.UpdateDefaultUpdatedAt()
		ttu.mutation.SetUpdatedAt(v)
	}
	if _, ok := ttu.mutation.DeletedAt(); !ok && !ttu.mutation.DeletedAtCleared() {
		v := transactiontype.UpdateDefaultDeletedAt()
		ttu.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttu *TransactionTypeUpdate) check() error {
	if v, ok := ttu.mutation.Value(); ok {
		if err := transactiontype.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "TransactionType.value": %w`, err)}
		}
	}
	if v, ok := ttu.mutation.Description(); ok {
		if err := transactiontype.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "TransactionType.description": %w`, err)}
		}
	}
	return nil
}

func (ttu *TransactionTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactiontype.Table,
			Columns: transactiontype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transactiontype.FieldID,
			},
		},
	}
	if ps := ttu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transactiontype.FieldUpdatedAt,
		})
	}
	if value, ok := ttu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transactiontype.FieldDeletedAt,
		})
	}
	if ttu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: transactiontype.FieldDeletedAt,
		})
	}
	if value, ok := ttu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transactiontype.FieldValue,
		})
	}
	if value, ok := ttu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactiontype.FieldDescription,
		})
	}
	if ttu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactiontype.FieldDescription,
		})
	}
	if ttu.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactiontype.TransactionsTable,
			Columns: []string{transactiontype.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !ttu.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactiontype.TransactionsTable,
			Columns: []string{transactiontype.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactiontype.TransactionsTable,
			Columns: []string{transactiontype.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ttu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactiontype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TransactionTypeUpdateOne is the builder for updating a single TransactionType entity.
type TransactionTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionTypeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ttuo *TransactionTypeUpdateOne) SetUpdatedAt(t time.Time) *TransactionTypeUpdateOne {
	ttuo.mutation.SetUpdatedAt(t)
	return ttuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ttuo *TransactionTypeUpdateOne) SetDeletedAt(t time.Time) *TransactionTypeUpdateOne {
	ttuo.mutation.SetDeletedAt(t)
	return ttuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ttuo *TransactionTypeUpdateOne) ClearDeletedAt() *TransactionTypeUpdateOne {
	ttuo.mutation.ClearDeletedAt()
	return ttuo
}

// SetValue sets the "value" field.
func (ttuo *TransactionTypeUpdateOne) SetValue(t transactiontype.Value) *TransactionTypeUpdateOne {
	ttuo.mutation.SetValue(t)
	return ttuo
}

// SetDescription sets the "description" field.
func (ttuo *TransactionTypeUpdateOne) SetDescription(s string) *TransactionTypeUpdateOne {
	ttuo.mutation.SetDescription(s)
	return ttuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ttuo *TransactionTypeUpdateOne) SetNillableDescription(s *string) *TransactionTypeUpdateOne {
	if s != nil {
		ttuo.SetDescription(*s)
	}
	return ttuo
}

// ClearDescription clears the value of the "description" field.
func (ttuo *TransactionTypeUpdateOne) ClearDescription() *TransactionTypeUpdateOne {
	ttuo.mutation.ClearDescription()
	return ttuo
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (ttuo *TransactionTypeUpdateOne) AddTransactionIDs(ids ...pulid.PULID) *TransactionTypeUpdateOne {
	ttuo.mutation.AddTransactionIDs(ids...)
	return ttuo
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (ttuo *TransactionTypeUpdateOne) AddTransactions(t ...*Transaction) *TransactionTypeUpdateOne {
	ids := make([]pulid.PULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ttuo.AddTransactionIDs(ids...)
}

// Mutation returns the TransactionTypeMutation object of the builder.
func (ttuo *TransactionTypeUpdateOne) Mutation() *TransactionTypeMutation {
	return ttuo.mutation
}

// ClearTransactions clears all "transactions" edges to the Transaction entity.
func (ttuo *TransactionTypeUpdateOne) ClearTransactions() *TransactionTypeUpdateOne {
	ttuo.mutation.ClearTransactions()
	return ttuo
}

// RemoveTransactionIDs removes the "transactions" edge to Transaction entities by IDs.
func (ttuo *TransactionTypeUpdateOne) RemoveTransactionIDs(ids ...pulid.PULID) *TransactionTypeUpdateOne {
	ttuo.mutation.RemoveTransactionIDs(ids...)
	return ttuo
}

// RemoveTransactions removes "transactions" edges to Transaction entities.
func (ttuo *TransactionTypeUpdateOne) RemoveTransactions(t ...*Transaction) *TransactionTypeUpdateOne {
	ids := make([]pulid.PULID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ttuo.RemoveTransactionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ttuo *TransactionTypeUpdateOne) Select(field string, fields ...string) *TransactionTypeUpdateOne {
	ttuo.fields = append([]string{field}, fields...)
	return ttuo
}

// Save executes the query and returns the updated TransactionType entity.
func (ttuo *TransactionTypeUpdateOne) Save(ctx context.Context) (*TransactionType, error) {
	var (
		err  error
		node *TransactionType
	)
	ttuo.defaults()
	if len(ttuo.hooks) == 0 {
		if err = ttuo.check(); err != nil {
			return nil, err
		}
		node, err = ttuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttuo.check(); err != nil {
				return nil, err
			}
			ttuo.mutation = mutation
			node, err = ttuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ttuo.hooks) - 1; i >= 0; i-- {
			if ttuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ttuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TransactionType)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TransactionTypeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttuo *TransactionTypeUpdateOne) SaveX(ctx context.Context) *TransactionType {
	node, err := ttuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ttuo *TransactionTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ttuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttuo *TransactionTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ttuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttuo *TransactionTypeUpdateOne) defaults() {
	if _, ok := ttuo.mutation.UpdatedAt(); !ok {
		v := transactiontype.UpdateDefaultUpdatedAt()
		ttuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := ttuo.mutation.DeletedAt(); !ok && !ttuo.mutation.DeletedAtCleared() {
		v := transactiontype.UpdateDefaultDeletedAt()
		ttuo.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttuo *TransactionTypeUpdateOne) check() error {
	if v, ok := ttuo.mutation.Value(); ok {
		if err := transactiontype.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "TransactionType.value": %w`, err)}
		}
	}
	if v, ok := ttuo.mutation.Description(); ok {
		if err := transactiontype.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "TransactionType.description": %w`, err)}
		}
	}
	return nil
}

func (ttuo *TransactionTypeUpdateOne) sqlSave(ctx context.Context) (_node *TransactionType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactiontype.Table,
			Columns: transactiontype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transactiontype.FieldID,
			},
		},
	}
	id, ok := ttuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TransactionType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ttuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transactiontype.FieldID)
		for _, f := range fields {
			if !transactiontype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transactiontype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ttuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transactiontype.FieldUpdatedAt,
		})
	}
	if value, ok := ttuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transactiontype.FieldDeletedAt,
		})
	}
	if ttuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: transactiontype.FieldDeletedAt,
		})
	}
	if value, ok := ttuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transactiontype.FieldValue,
		})
	}
	if value, ok := ttuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactiontype.FieldDescription,
		})
	}
	if ttuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactiontype.FieldDescription,
		})
	}
	if ttuo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactiontype.TransactionsTable,
			Columns: []string{transactiontype.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !ttuo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactiontype.TransactionsTable,
			Columns: []string{transactiontype.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactiontype.TransactionsTable,
			Columns: []string{transactiontype.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TransactionType{config: ttuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ttuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactiontype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
