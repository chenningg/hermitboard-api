package auth

import "fmt"

var ErrInvalidPassword = fmt.Errorf("auth: password must be a minimum of 8 alphanumeric characters and symbols")
var ErrIncorrectUsernameOrPassword = fmt.Errorf("auth: incorrect username or password")
var ErrBadInput = fmt.Errorf("auth: invalid or malformed input")
var ErrAlreadyLoggedIn = fmt.Errorf("auth: already logged in, cannot create account")
var ErrInternal = fmt.Errorf("auth: internal server error")
var ErrUnauthorized = fmt.Errorf("auth: not authorized")
