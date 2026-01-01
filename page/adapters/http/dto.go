package http

import "time"

// requests
type CreatePageRequest struct {
	Title string `json:"title"`
}

type UpdateTitleRequest struct {
	Title string `json:"title"`
}

// responses
type PageResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
