package gographqltwitter

import "errors"

var (
	ErrBadCredentials = errors.New("email/password wrong combination")
	ErrValidation     = errors.New("validation error")
	ErrNotFound       = errors.New("not found error")
)
