package storage

import "errors"

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("User not found")
	ErrAppNotFound  = errors.New("App not found")
)
