package model

import "errors"

var (
	ErrNotFound        = errors.New("Event not found ")
	ErrInvalidInterval = errors.New("Invalid interval ")
)
