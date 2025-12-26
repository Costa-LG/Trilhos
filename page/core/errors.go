package page

import "errors"

var (
	ErrNotFound     = errors.New("page not found")
	ErrInvalidTitle = errors.New("invalid title")
)
