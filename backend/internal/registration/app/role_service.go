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

func (s *RoleService) Save(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	return s.roleRepository.Save(ctx, role)
}
