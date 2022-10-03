package model

import (
	"io"
)

type Void bool

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (void *Void) UnmarshalGQL(v interface{}) error {
	*void = false
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (void Void) MarshalGQL(w io.Writer) {
	w.Write(nil)
}
