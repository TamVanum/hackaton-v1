package app

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

type TechnologyService struct {
	technologyRepository domain.TechnologyRepository
}

func NewTechnologyService(technologyRepository domain.TechnologyRepository) domain.TechnologyService {
	return &TechnologyService{
		technologyRepository: technologyRepository,
	}
}

func (s *TechnologyService) Get(ctx context.Context) ([]*domain.Technology, error) {
	return s.technologyRepository.FindAll(ctx)
}

func (s *TechnologyService) GetByIDs(ctx context.Context, ids []int) ([]*domain.Technology, error) {
	return s.technologyRepository.FindByIDs(ctx, ids)
}

func (s *TechnologyService) Make(ctx context.Context, name, description string) (*domain.Technology, error) {
	technology, err := domain.NewTechnology(name, description)
	if err != nil {
		return nil, err
	}

	return technology, nil
}

func (s *TechnologyService) Create(ctx context.Context, name, description string) (*domain.Technology, error) {
	technology, err := s.Make(ctx, name, description)
	if err != nil {
		return nil, err
	}

	persistedTechnology, err := s.Persist(ctx, technology)
	if err != nil {
		return nil, err
	}

	return persistedTechnology, nil
}

func (s *TechnologyService) Persist(ctx context.Context, technology *domain.Technology) (*domain.Technology, error) {
	return s.technologyRepository.Save(ctx, technology)
}
