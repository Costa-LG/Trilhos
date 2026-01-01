package page

import "errors"

var (
	ErrNotFound		= errors.New("page not found")
	ErrInvalidTime	= errors.New("invalid time")
	ErrInvalidTitle	= errors.New("invalid title")
	ErrInvalidValue	= errors.New("invalid value")
	ErrInvalidID	= errors.New("invalid id")
)
