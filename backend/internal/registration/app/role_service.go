package app

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

type RoleService struct {
	roleRepository domain.RoleRepository
}

func NewRoleService(roleRepository domain.RoleRepository) domain.RoleService {
	return &RoleService{
		roleRepository: roleRepository,
	}
}

func (s *RoleService) Get(ctx context.Context) ([]*domain.Role, error) {
	return s.roleRepository.FindAll(ctx)
}

func (s *RoleService) GetByIDs(ctx context.Context, ids []int) ([]*domain.Role, error) {
	return s.roleRepository.FindByIDs(ctx, ids)
}

func (s *RoleService) Make(ctx context.Context, name, description string) (*domain.Role, error) {
	role, err := domain.NewRole(name, description)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *RoleService) Persist(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	return s.roleRepository.Save(ctx, role)
}

func (s *RoleService) Create(ctx context.Context, name, description string) (*domain.Role, error) {
	role, err := s.Make(ctx, name, description)
	if err != nil {
		return nil, err
	}

	persistedRole, err := s.Persist(ctx, role)
	if err != nil {
		return nil, err
	}

	return persistedRole, nil
}
