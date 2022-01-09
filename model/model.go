package model

import "errors"

var (
	ErrPersonCanNotBeNill = errors.New("Person can not be nil")
	ErrPersonNotFound     = errors.New("Person not found")
)
