// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/chenningg/hermitboard-api/graph/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNTime2timeᚐTime(ctx context.Context, v interface{}) (time.Time, error) {
	res, err := graphql.UnmarshalTime(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNTime2timeᚐTime(ctx context.Context, sel ast.SelectionSet, v time.Time) graphql.Marshaler {
	res := graphql.MarshalTime(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalOTime2ᚕtimeᚐTimeᚄ(ctx context.Context, v interface{}) ([]time.Time, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]time.Time, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNTime2timeᚐTime(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalOTime2ᚕtimeᚐTimeᚄ(ctx context.Context, sel ast.SelectionSet, v []time.Time) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNTime2timeᚐTime(ctx, sel, v[i])
	}

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) unmarshalOTime2ᚖtimeᚐTime(ctx context.Context, v interface{}) (*time.Time, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalTime(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOTime2ᚖtimeᚐTime(ctx context.Context, sel ast.SelectionSet, v *time.Time) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	res := graphql.MarshalTime(*v)
	return res
}

func (ec *executionContext) unmarshalOVoid2ᚖgithubᚗcomᚋchenninggᚋhermitboardᚑapiᚋgraphᚋmodelᚐVoid(ctx context.Context, v interface{}) (*model.Void, error) {
	if v == nil {
		return nil, nil
	}
	var res = new(model.Void)
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOVoid2ᚖgithubᚗcomᚋchenninggᚋhermitboardᚑapiᚋgraphᚋmodelᚐVoid(ctx context.Context, sel ast.SelectionSet, v *model.Void) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

// endregion ***************************** type.gotpl *****************************
