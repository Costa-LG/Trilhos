package page

import (
	"context"
	"time"
)

// Sobre o context.Contex, ele é um objeto do Go para carregar prazo (deadline), cancelamento e valores do request.
// Package context defines the Context type, which carries deadlines, cancellation signals,
// and other request-scoped values across API boundaries and between processes.

type Repository interface {
	Create(ctx context.Context, page Page) error
	Get(ctx context.Context, id ID) (Page, error)
	List(ctx context.Context) ([]Page, error)
	Update(ctx context.Context, page Page) (Page, error)
	Delete(ctx context.Context, id ID) error
}

type Clock interface {
	Now() time.Time
}

type IDGenerator interface {
	NewID() ID
}
