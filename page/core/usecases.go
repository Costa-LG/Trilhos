package page

import (
	"context"
)

type Service struct {
	repo  Repository
	clock Clock
	idGen IDGenerator
}

// Construtor do servico de usecases.
// Recebe as dependencias que o Service precisa, retornando uma instancia
// (camada suja da hexagonal)
func NewService(repo Repository, clock Clock, idGen IDGenerator) *Service {
	return &Service{
		repo:  repo,
		clock: clock,
		idGen: idGen,
	}
}

func (s *Service) Create(ctx context.Context, title string) (Page, error) {
	page, err := New(s.idGen.NewID(), title, s.clock.Now())
	if err != nil {
		return Page{}, err
	}

	if err := s.repo.Create(ctx, page); err != nil {
		return Page{}, err
	}

	return page, nil
}

func (s *Service) List(ctx context.Context) ([]Page, error) {
	return s.repo.List(ctx)
}

