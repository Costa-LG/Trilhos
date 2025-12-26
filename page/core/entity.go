package page

import (
	"strings"
	"time"
)

type ID string

type Page struct {
	ID        ID
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(id ID, title string, now time.Time) (Page, error) {
	cleaned := strings.TrimSpace(title)
	if cleaned == "" {
		return Page{}, ErrInvalidTitle
	}

	return Page{
		ID:        id,
		Title:     cleaned,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}