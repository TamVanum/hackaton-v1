package app

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

type RoleService struct {
	roleRepository domain.RoleRepositoryPort
}

func NewRoleService(roleRepository domain.RoleRepositoryPort) *RoleService {
	return &RoleService{
		roleRepository: roleRepository,
	}
}

func (s *RoleService) Get(ctx context.Context) ([]*domain.Role, error) {
	return s.roleRepository.FindAll(ctx)
}

func (s *RoleService) GetByIDs(ctx context.Context, ids []int) ([]*domain.Role, error) {
	return s.roleRepository.BulkFindByIDs(ctx, ids)
}

func (s *RoleService) Save(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	return s.roleRepository.Save(ctx, role)
}
