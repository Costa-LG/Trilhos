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
	id, err := validateID(id)
	if err != nil{
		return Page{}, err
	}
	cleaned, err := validateTitle(title)
	if err != nil{
		return Page{}, err
	}
	
	return Page{
		ID:        id,
		Title:     cleaned,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (p *Page) Rename(newTitle string, now time.Time) error {
	cleaned, err := validateTitle(newTitle)
	if err != nil{
		return err
	}
	if now.Before(p.CreatedAt) {
		return ErrInvalidTime
	}
	p.Title = cleaned
	p.UpdatedAt = now
	return nil
}

func validateID(id ID) (ID, error){
	if id == "" {
		return id, ErrInvalidID
	}
	return id, nil
}

func validateTitle(title string) (string, error) {
	cleaned := strings.TrimSpace(title)
	if cleaned == "" {
		return cleaned, ErrInvalidTitle
	}
	length := len(cleaned)
	if length < 1 || length > 100 {
		return cleaned, ErrInvalidValue
	}
	return cleaned, nil
}