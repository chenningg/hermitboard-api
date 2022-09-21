// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/chenningg/hermitboard-api/ent"
)

// The AccountFunc type is an adapter to allow the use of ordinary
// function as Account mutator.
type AccountFunc func(context.Context, *ent.AccountMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AccountFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AccountMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AccountMutation", m)
	}
	return f(ctx, mv)
}

// The AccountAuthRoleFunc type is an adapter to allow the use of ordinary
// function as AccountAuthRole mutator.
type AccountAuthRoleFunc func(context.Context, *ent.AccountAuthRoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AccountAuthRoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AccountAuthRoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AccountAuthRoleMutation", m)
	}
	return f(ctx, mv)
}

// The AssetFunc type is an adapter to allow the use of ordinary
// function as Asset mutator.
type AssetFunc func(context.Context, *ent.AssetMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AssetFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AssetMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AssetMutation", m)
	}
	return f(ctx, mv)
}

// The AssetClassFunc type is an adapter to allow the use of ordinary
// function as AssetClass mutator.
type AssetClassFunc func(context.Context, *ent.AssetClassMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AssetClassFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AssetClassMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AssetClassMutation", m)
	}
	return f(ctx, mv)
}

// The AuthRoleFunc type is an adapter to allow the use of ordinary
// function as AuthRole mutator.
type AuthRoleFunc func(context.Context, *ent.AuthRoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AuthRoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AuthRoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AuthRoleMutation", m)
	}
	return f(ctx, mv)
}

// The BlockchainFunc type is an adapter to allow the use of ordinary
// function as Blockchain mutator.
type BlockchainFunc func(context.Context, *ent.BlockchainMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BlockchainFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BlockchainMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BlockchainMutation", m)
	}
	return f(ctx, mv)
}

// The BlockchainCryptocurrencyFunc type is an adapter to allow the use of ordinary
// function as BlockchainCryptocurrency mutator.
type BlockchainCryptocurrencyFunc func(context.Context, *ent.BlockchainCryptocurrencyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BlockchainCryptocurrencyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BlockchainCryptocurrencyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BlockchainCryptocurrencyMutation", m)
	}
	return f(ctx, mv)
}

// The CryptocurrencyFunc type is an adapter to allow the use of ordinary
// function as Cryptocurrency mutator.
type CryptocurrencyFunc func(context.Context, *ent.CryptocurrencyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CryptocurrencyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CryptocurrencyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CryptocurrencyMutation", m)
	}
	return f(ctx, mv)
}

// The DailyAssetPriceFunc type is an adapter to allow the use of ordinary
// function as DailyAssetPrice mutator.
type DailyAssetPriceFunc func(context.Context, *ent.DailyAssetPriceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DailyAssetPriceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DailyAssetPriceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DailyAssetPriceMutation", m)
	}
	return f(ctx, mv)
}

// The ExchangeFunc type is an adapter to allow the use of ordinary
// function as Exchange mutator.
type ExchangeFunc func(context.Context, *ent.ExchangeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ExchangeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ExchangeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExchangeMutation", m)
	}
	return f(ctx, mv)
}

// The PortfolioFunc type is an adapter to allow the use of ordinary
// function as Portfolio mutator.
type PortfolioFunc func(context.Context, *ent.PortfolioMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PortfolioFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PortfolioMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PortfolioMutation", m)
	}
	return f(ctx, mv)
}

// The StaffAccountFunc type is an adapter to allow the use of ordinary
// function as StaffAccount mutator.
type StaffAccountFunc func(context.Context, *ent.StaffAccountMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StaffAccountFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StaffAccountMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StaffAccountMutation", m)
	}
	return f(ctx, mv)
}

// The StaffAccountAuthRoleFunc type is an adapter to allow the use of ordinary
// function as StaffAccountAuthRole mutator.
type StaffAccountAuthRoleFunc func(context.Context, *ent.StaffAccountAuthRoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StaffAccountAuthRoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StaffAccountAuthRoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StaffAccountAuthRoleMutation", m)
	}
	return f(ctx, mv)
}

// The TransactionFunc type is an adapter to allow the use of ordinary
// function as Transaction mutator.
type TransactionFunc func(context.Context, *ent.TransactionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TransactionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TransactionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TransactionMutation", m)
	}
	return f(ctx, mv)
}

// The TransactionTypeFunc type is an adapter to allow the use of ordinary
// function as TransactionType mutator.
type TransactionTypeFunc func(context.Context, *ent.TransactionTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TransactionTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TransactionTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TransactionTypeMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}