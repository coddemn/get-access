package apperrors

import "errors"

var (
	ErrNameTaken     = errors.New("name is already taken")
	ErrIncorrectAuth = errors.New("incorrect login or password")
)
