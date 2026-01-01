# Trilhos
Trilhos é uma aplicação inspirada em notion, desenvolvida em Go, para organização pessoal por meio de blocos e páginas reutilizáveis.
O sistema será construído em arquitetura Hexagonal.

## Schema

pages (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

blocks (
    id TEXT PRIMARY KEY,
    page_id TEXT NOT NULL REFERENCES pages(id),
    parent_block_id TEXT REFERENCES blocks(id),
    block_type TEXT NOT NULL,
    props JSONB,
    position INT NOT NULL
);
TODO
adicionar on delete cascade

## API (Pages)

PageResponse:
```json
{
  "id": "string",
  "title": "string",
  "created_at": "RFC3339 timestamp",
  "updated_at": "RFC3339 timestamp"
}
```

ErrorResponse:
```json
{
  "error": "invalid_title",
  "message": "title cannot be empty"
}
```

POST /pages
- Request: `{ "title": "..." }`
- Response 201: PageResponse

GET /pages
- Response 200: [PageResponse]
- Order: created_at desc (most recent first)

GET /pages/{id}
- Response 200: PageResponse

PATCH /pages/{id}
- Request: `{ "title": "..." }`
- Response 200: PageResponse

DELETE /pages/{id}
- Response 204 (no body)

```
Trilhos/
  api/
    main.go
  
  internal/
    shared/
      domain/
        id.go
      errors/
        errors.go
      tx/
        tx.go

  page/
    core/
      entity.go
      ports.go
      usecases.go
      errors.go
    adapters/
      handler.go
      dto.go
      repo.go
  block/
    core/
      entity.go
      ports.go
      usecases.go
      errors.go
    adapters/
      handler.go
      dto.go
      repo.go
  app/
    http/
      router.go
      middleware.go
    postgres/
      db.go
      migrations/
        001_init.sql
```
