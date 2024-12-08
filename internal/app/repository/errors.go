package repository

import "errors"

var (
	ErrNotFound       = errors.New("user not found")
	ErrDuplicatedUser = errors.New("email already exists")
	ErrOnDelete       = errors.New("delete operation was not successful")
)
