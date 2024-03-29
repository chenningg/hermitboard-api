// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	hermitboard_api "github.com/chenningg/hermitboard-api"
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

func (ec *executionContext) unmarshalInputDeleteConnectionInput(ctx context.Context, obj interface{}) (hermitboard_api.DeleteConnectionInput, error) {
	var it hermitboard_api.DeleteConnectionInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			it.ID, err = ec.unmarshalNID2githubᚗcomᚋchenninggᚋhermitboardᚑapiᚋpulidᚐPULID(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

// endregion ***************************** type.gotpl *****************************
