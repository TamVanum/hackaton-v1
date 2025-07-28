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

func (r *SqlRoleRepository) FindAll(ctx context.Context) ([]*domain.Role, error) {
	query := `
		SELECT id, name
		FROM roles
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	roles := []*domain.Role{}

	for rows.Next() {
		var role domain.Role
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}

	return roles, nil
}

func (r *SqlRoleRepository) BulkFindByIDs(ctx context.Context, ids []int) ([]*domain.Role, error) {
	query := `
		SELECT id, name
		FROM roles
		WHERE id IN (?)
	`

	rows, err := r.db.QueryContext(ctx, query, ids)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	roles := []*domain.Role{}

	for rows.Next() {
		var role domain.Role
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}

	return roles, nil
}
