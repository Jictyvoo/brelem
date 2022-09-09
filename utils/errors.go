package utils

import "errors"

var (
	ErrEmptyString            = errors.New("empty string")
	ErrInvalidElement         = errors.New("given element is invalid using the requested validator")
	ErrElementIncorrectLength = errors.New("given element has not a correct length")
	ErrExpiredDocument        = errors.New("document is already expred")
)
