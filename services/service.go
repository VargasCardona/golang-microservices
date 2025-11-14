package services

import (
	"context"
	"github.com/guitar-service/crud"
)

type GuitarService struct {
	repo crud.GuitarRepository
}

func NewGuitarService(repo crud.GuitarRepository) *GuitarService {
	return &GuitarService{repo: repo}
}

// CREATE
func (s *GuitarService) CreateGuitar(ctx context.Context, g *crud.GuitarModel) error {
	return s.repo.Create(ctx, g)
}

// READ ONE
func (s *GuitarService) GetGuitar(ctx context.Context, id uint) (*crud.GuitarModel, error) {
	return s.repo.GetByID(ctx, id)
}

// LIST
func (s *GuitarService) ListGuitars(ctx context.Context) ([]crud.GuitarModel, error) {
	return s.repo.List(ctx)
}

// UPDATE
func (s *GuitarService) UpdateGuitar(ctx context.Context, g *crud.GuitarModel) error {
	return s.repo.Update(ctx, g)
}

// DELETE
func (s *GuitarService) DeleteGuitar(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
