# Trilhos
Trilhos é uma aplicação inspirada em notion, desenvolvida em Go, para organização pessoal por meio de blocos e páginas reutilizáveis.
O sistema será construído em arquitetura Hexagonal.

## Schema

Page
{
  id,
  title,
  created_at,
  updated_at,
}

Block
{
  id,
  page_id,
  parent_block_id (optional),
  type (ex: paragraph, heading1, todo, image, code),
  props (JSON com dados do tipo: {text}, {url, caption}, {code, language}),
  position,
}

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
      http/
        handler.go
        dto.go
      postgres/
        repo.go
  block/
    core/
      entity.go
      ports.go
      usecases.go
      errors.go
    adapters/
      http/
        handler.go
        dto.go
      postgres/
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