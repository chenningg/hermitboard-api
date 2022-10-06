package auth

import "fmt"

var ErrInvalidPassword = fmt.Errorf("invalid password")
var ErrIncorrectUsernameOrPassword = fmt.Errorf("incorrect username or password")
var ErrWeakPassword = fmt.Errorf("password is too weak")
var ErrNotFound = fmt.Errorf("requested resource could not be found")
var ErrBadInput = fmt.Errorf("invalid or malformed input")
var ErrInternal = fmt.Errorf("internal server error")
var ErrUnauthorized = fmt.Errorf("not authorized")
