package mixin

import (
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/chenningg/hermitboard-api/ent/schema/pulid"
)

// MixinWithPrefix creates a Mixin that encodes the provided prefix.
func PULIDMixinWithPrefix(prefix string) *PULIDMixin {
	return &PULIDMixin{prefix: strings.ToUpper(prefix)}
}

type PULIDMixin struct {
	mixin.Schema
	prefix string
}

func (m PULIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(pulid.PULID("")).
			DefaultFunc(func() pulid.PULID { return pulid.NewPULID(m.prefix) }),
	}
}

// Annotation captures the id prefix for a type.
type Annotation struct {
	Prefix string
}

// Name implements the ent Annotation interface.
func (a Annotation) Name() string {
	return "PULID"
}

// Annotations returns the annotations for a Mixin instance.
func (m PULIDMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{Prefix: m.prefix},
	}
}
