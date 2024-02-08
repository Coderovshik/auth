package usecase

import (
	"errors"
)

var (
	ErrEmailEmpty         = errors.New("email is empty")
	ErrPasswordEmpty      = errors.New("password is empty")
	ErrIdEmpty            = errors.New("id is empty")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppInfo     = errors.New("invalid app info")
	ErrInvalidUserId      = errors.New("invalid user id")
)
