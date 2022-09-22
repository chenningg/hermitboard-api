// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AccountQuery) CollectFields(ctx context.Context, satisfies ...string) (*AccountQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AccountQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "authRoles", "auth_roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AuthRoleQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.WithNamedAuthRoles(alias, func(wq *AuthRoleQuery) {
				*wq = *query
			})
		case "portfolios":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PortfolioQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.WithNamedPortfolios(alias, func(wq *PortfolioQuery) {
				*wq = *query
			})
		case "authType", "auth_type":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AuthTypeQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.withAuthType = query
		case "accountAuthRoles", "account_auth_roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AccountAuthRoleQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.WithNamedAccountAuthRoles(alias, func(wq *AccountAuthRoleQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type accountPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AccountPaginateOption
}

func newAccountPaginateArgs(rv map[string]interface{}) *accountPaginateArgs {
	args := &accountPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (aar *AccountAuthRoleQuery) CollectFields(ctx context.Context, satisfies ...string) (*AccountAuthRoleQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return aar, nil
	}
	if err := aar.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return aar, nil
}

func (aar *AccountAuthRoleQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "account":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AccountQuery{config: aar.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			aar.withAccount = query
		case "authRole", "auth_role":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AuthRoleQuery{config: aar.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			aar.withAuthRole = query
		}
	}
	return nil
}

type accountauthrolePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AccountAuthRolePaginateOption
}

func newAccountAuthRolePaginateArgs(rv map[string]interface{}) *accountauthrolePaginateArgs {
	args := &accountauthrolePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AssetQuery) CollectFields(ctx context.Context, satisfies ...string) (*AssetQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AssetQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "assetClass", "asset_class":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AssetClassQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.withAssetClass = query
		case "cryptocurrency":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CryptocurrencyQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.withCryptocurrency = query
		case "transactionBase", "transaction_base":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TransactionQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.WithNamedTransactionBase(alias, func(wq *TransactionQuery) {
				*wq = *query
			})
		case "transactionQuote", "transaction_quote":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TransactionQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.WithNamedTransactionQuote(alias, func(wq *TransactionQuery) {
				*wq = *query
			})
		case "dailyAssetPrice", "daily_asset_price":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &DailyAssetPriceQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.WithNamedDailyAssetPrice(alias, func(wq *DailyAssetPriceQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type assetPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AssetPaginateOption
}

func newAssetPaginateArgs(rv map[string]interface{}) *assetPaginateArgs {
	args := &assetPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ac *AssetClassQuery) CollectFields(ctx context.Context, satisfies ...string) (*AssetClassQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ac, nil
	}
	if err := ac.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ac, nil
}

func (ac *AssetClassQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "assets":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AssetQuery{config: ac.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ac.WithNamedAssets(alias, func(wq *AssetQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type assetclassPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AssetClassPaginateOption
}

func newAssetClassPaginateArgs(rv map[string]interface{}) *assetclassPaginateArgs {
	args := &assetclassPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &AssetClassOrder{Field: &AssetClassOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithAssetClassOrder(order))
			}
		case *AssetClassOrder:
			if v != nil {
				args.opts = append(args.opts, WithAssetClassOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ar *AuthRoleQuery) CollectFields(ctx context.Context, satisfies ...string) (*AuthRoleQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ar, nil
	}
	if err := ar.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ar, nil
}

func (ar *AuthRoleQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "accounts":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AccountQuery{config: ar.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ar.WithNamedAccounts(alias, func(wq *AccountQuery) {
				*wq = *query
			})
		case "staffAccounts", "staff_accounts":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &StaffAccountQuery{config: ar.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ar.WithNamedStaffAccounts(alias, func(wq *StaffAccountQuery) {
				*wq = *query
			})
		case "accountAuthRoles", "account_auth_roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AccountAuthRoleQuery{config: ar.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ar.WithNamedAccountAuthRoles(alias, func(wq *AccountAuthRoleQuery) {
				*wq = *query
			})
		case "staffAccountAuthRoles", "staff_account_auth_roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &StaffAccountAuthRoleQuery{config: ar.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ar.WithNamedStaffAccountAuthRoles(alias, func(wq *StaffAccountAuthRoleQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type authrolePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AuthRolePaginateOption
}

func newAuthRolePaginateArgs(rv map[string]interface{}) *authrolePaginateArgs {
	args := &authrolePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &AuthRoleOrder{Field: &AuthRoleOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithAuthRoleOrder(order))
			}
		case *AuthRoleOrder:
			if v != nil {
				args.opts = append(args.opts, WithAuthRoleOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (at *AuthTypeQuery) CollectFields(ctx context.Context, satisfies ...string) (*AuthTypeQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return at, nil
	}
	if err := at.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return at, nil
}

func (at *AuthTypeQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "account":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AccountQuery{config: at.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			at.withAccount = query
		case "staffAccount", "staff_account":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &StaffAccountQuery{config: at.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			at.withStaffAccount = query
		}
	}
	return nil
}

type authtypePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AuthTypePaginateOption
}

func newAuthTypePaginateArgs(rv map[string]interface{}) *authtypePaginateArgs {
	args := &authtypePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &AuthTypeOrder{Field: &AuthTypeOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithAuthTypeOrder(order))
			}
		case *AuthTypeOrder:
			if v != nil {
				args.opts = append(args.opts, WithAuthTypeOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (b *BlockchainQuery) CollectFields(ctx context.Context, satisfies ...string) (*BlockchainQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return b, nil
	}
	if err := b.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return b, nil
}

func (b *BlockchainQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "cryptocurrencies":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CryptocurrencyQuery{config: b.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			b.WithNamedCryptocurrencies(alias, func(wq *CryptocurrencyQuery) {
				*wq = *query
			})
		case "blockchainCryptocurrencies", "blockchain_cryptocurrencies":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &BlockchainCryptocurrencyQuery{config: b.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			b.WithNamedBlockchainCryptocurrencies(alias, func(wq *BlockchainCryptocurrencyQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type blockchainPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []BlockchainPaginateOption
}

func newBlockchainPaginateArgs(rv map[string]interface{}) *blockchainPaginateArgs {
	args := &blockchainPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (bc *BlockchainCryptocurrencyQuery) CollectFields(ctx context.Context, satisfies ...string) (*BlockchainCryptocurrencyQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return bc, nil
	}
	if err := bc.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return bc, nil
}

func (bc *BlockchainCryptocurrencyQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "blockchain":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &BlockchainQuery{config: bc.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			bc.withBlockchain = query
		case "cryptocurrency":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CryptocurrencyQuery{config: bc.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			bc.withCryptocurrency = query
		}
	}
	return nil
}

type blockchaincryptocurrencyPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []BlockchainCryptocurrencyPaginateOption
}

func newBlockchainCryptocurrencyPaginateArgs(rv map[string]interface{}) *blockchaincryptocurrencyPaginateArgs {
	args := &blockchaincryptocurrencyPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CryptocurrencyQuery) CollectFields(ctx context.Context, satisfies ...string) (*CryptocurrencyQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CryptocurrencyQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "asset":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AssetQuery{config: c.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			c.withAsset = query
		case "blockchains":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &BlockchainQuery{config: c.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			c.WithNamedBlockchains(alias, func(wq *BlockchainQuery) {
				*wq = *query
			})
		case "blockchainCryptocurrencies", "blockchain_cryptocurrencies":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &BlockchainCryptocurrencyQuery{config: c.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			c.WithNamedBlockchainCryptocurrencies(alias, func(wq *BlockchainCryptocurrencyQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type cryptocurrencyPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CryptocurrencyPaginateOption
}

func newCryptocurrencyPaginateArgs(rv map[string]interface{}) *cryptocurrencyPaginateArgs {
	args := &cryptocurrencyPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (dap *DailyAssetPriceQuery) CollectFields(ctx context.Context, satisfies ...string) (*DailyAssetPriceQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return dap, nil
	}
	if err := dap.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return dap, nil
}

func (dap *DailyAssetPriceQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "baseAsset", "base_asset":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AssetQuery{config: dap.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			dap.withBaseAsset = query
		}
	}
	return nil
}

type dailyassetpricePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []DailyAssetPricePaginateOption
}

func newDailyAssetPricePaginateArgs(rv map[string]interface{}) *dailyassetpricePaginateArgs {
	args := &dailyassetpricePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (e *ExchangeQuery) CollectFields(ctx context.Context, satisfies ...string) (*ExchangeQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return e, nil
	}
	if err := e.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return e, nil
}

func (e *ExchangeQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "transactions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TransactionQuery{config: e.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			e.WithNamedTransactions(alias, func(wq *TransactionQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type exchangePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ExchangePaginateOption
}

func newExchangePaginateArgs(rv map[string]interface{}) *exchangePaginateArgs {
	args := &exchangePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (po *PortfolioQuery) CollectFields(ctx context.Context, satisfies ...string) (*PortfolioQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return po, nil
	}
	if err := po.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return po, nil
}

func (po *PortfolioQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "account":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AccountQuery{config: po.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			po.withAccount = query
		case "transactions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TransactionQuery{config: po.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			po.WithNamedTransactions(alias, func(wq *TransactionQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type portfolioPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PortfolioPaginateOption
}

func newPortfolioPaginateArgs(rv map[string]interface{}) *portfolioPaginateArgs {
	args := &portfolioPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sa *StaffAccountQuery) CollectFields(ctx context.Context, satisfies ...string) (*StaffAccountQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sa, nil
	}
	if err := sa.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sa, nil
}

func (sa *StaffAccountQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "authRoles", "auth_roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AuthRoleQuery{config: sa.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sa.WithNamedAuthRoles(alias, func(wq *AuthRoleQuery) {
				*wq = *query
			})
		case "authType", "auth_type":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AuthTypeQuery{config: sa.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sa.withAuthType = query
		case "staffAccountAuthRoles", "staff_account_auth_roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &StaffAccountAuthRoleQuery{config: sa.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sa.WithNamedStaffAccountAuthRoles(alias, func(wq *StaffAccountAuthRoleQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type staffaccountPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []StaffAccountPaginateOption
}

func newStaffAccountPaginateArgs(rv map[string]interface{}) *staffaccountPaginateArgs {
	args := &staffaccountPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (saar *StaffAccountAuthRoleQuery) CollectFields(ctx context.Context, satisfies ...string) (*StaffAccountAuthRoleQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return saar, nil
	}
	if err := saar.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return saar, nil
}

func (saar *StaffAccountAuthRoleQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "staffAccount", "staff_account":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &StaffAccountQuery{config: saar.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			saar.withStaffAccount = query
		case "authRole", "auth_role":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AuthRoleQuery{config: saar.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			saar.withAuthRole = query
		}
	}
	return nil
}

type staffaccountauthrolePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []StaffAccountAuthRolePaginateOption
}

func newStaffAccountAuthRolePaginateArgs(rv map[string]interface{}) *staffaccountauthrolePaginateArgs {
	args := &staffaccountauthrolePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TransactionQuery) CollectFields(ctx context.Context, satisfies ...string) (*TransactionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TransactionQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "transactionType", "transaction_type":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TransactionTypeQuery{config: t.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.withTransactionType = query
		case "baseAsset", "base_asset":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AssetQuery{config: t.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.withBaseAsset = query
		case "quoteAsset", "quote_asset":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AssetQuery{config: t.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.withQuoteAsset = query
		case "portfolio":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PortfolioQuery{config: t.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.withPortfolio = query
		case "exchange":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &ExchangeQuery{config: t.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.withExchange = query
		}
	}
	return nil
}

type transactionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TransactionPaginateOption
}

func newTransactionPaginateArgs(rv map[string]interface{}) *transactionPaginateArgs {
	args := &transactionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (tt *TransactionTypeQuery) CollectFields(ctx context.Context, satisfies ...string) (*TransactionTypeQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return tt, nil
	}
	if err := tt.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return tt, nil
}

func (tt *TransactionTypeQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "transactions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TransactionQuery{config: tt.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			tt.WithNamedTransactions(alias, func(wq *TransactionQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type transactiontypePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TransactionTypePaginateOption
}

func newTransactionTypePaginateArgs(rv map[string]interface{}) *transactiontypePaginateArgs {
	args := &transactiontypePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &TransactionTypeOrder{Field: &TransactionTypeOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTransactionTypeOrder(order))
			}
		case *TransactionTypeOrder:
			if v != nil {
				args.opts = append(args.opts, WithTransactionTypeOrder(v))
			}
		}
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}
