package connection

import "fmt"

var ErrBadInput = fmt.Errorf("invalid or malformed input")
var ErrInternal = fmt.Errorf("internal server error")
var ErrUnauthorized = fmt.Errorf("not authorized")
var ErrNotFound = fmt.Errorf("not found")
