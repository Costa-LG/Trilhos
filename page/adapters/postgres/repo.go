package postgres

import (
	"context"
	"database/sql"
	"errors"

	page "Trilhos/page/core"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, p page.Page) error {
	const query = `
		INSERT INTO pages (id, title, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.ExecContext(ctx, query, string(p.ID), p.Title, p.CreatedAt, p.UpdatedAt)
	return err
}

func (r *Repository) Get(ctx context.Context, id page.ID) (page.Page, error) {
	const query = `
		SELECT id, title, created_at, updated_at
		FROM pages
		WHERE id = $1
	`
	var (
		p     page.Page
		rawID string
	)
	err := r.db.QueryRowContext(ctx, query, string(id)).
		Scan(&rawID, &p.Title, &p.CreatedAt, &p.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return page.Page{}, page.ErrNotFound
	}
	if err != nil {
		return page.Page{}, err
	}
	p.ID = page.ID(rawID)
	return p, nil
}

func (r *Repository) List(ctx context.Context) ([]page.Page, error) {
	const query = `
		SELECT id, title, created_at, updated_at
		FROM pages
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pages := make([]page.Page, 0)
	// Hydrate domain pages from each row returned by the query.
	for rows.Next() {
		var (
			id string
			p  page.Page
		)
		if err := rows.Scan(&id, &p.Title, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		p.ID = page.ID(id)
		pages = append(pages, p)
	}
	// rows.Err captures any scan/iteration errors that happened in the loop.
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return pages, nil
}

func (r *Repository) Update(ctx context.Context, p page.Page) (page.Page, error) {
	const query = `
		UPDATE pages
		SET
			title = $1,
			updated_at = $2
		WHERE id = $3
	`

	result, err := r.db.ExecContext(ctx, query, p.Title, p.UpdatedAt, string(p.ID))
	if err != nil {
		return page.Page{}, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return page.Page{}, err
	}
	if affected == 0 {
		return page.ErrNotFound
	}
	return p, nil
}

func (r *Repository) Delete(ctx context.Context, id page.ID) error {
	const query = `
		DELETE
		FROM pages
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, string(id))
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return page.ErrNotFound
	}
	return nil
}
