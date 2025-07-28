package sql

import (
	"context"
	"database/sql"
	"hackathon-pvc-backend/internal/registration/domain"
)

type SqlRoleRepository struct {
	db *sql.DB
}

func NewSqlRoleRepository(db *sql.DB) domain.RoleRepositoryPort {
	return &SqlRoleRepository{db: db}
}

func (r *SqlRoleRepository) Save(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	query := `
		INSERT INTO roles (name)
		VALUES (?)
	`

	result, err := r.db.ExecContext(ctx, query, role.Name())
	if err != nil {
		return nil, err
	}

	roleID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	role.SetID(int(roleID))

	return role, nil
}
