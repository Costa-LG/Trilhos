package page

import (
	"context"
	"time"
)

type Repository interface {
	Create(ctx context.Context, page Page) error
	List(ctx context.Context) ([]Page, error)
}

type Clock interface {
	Now() time.Time
}

type IDGenerator interface {
	NewID() ID
}