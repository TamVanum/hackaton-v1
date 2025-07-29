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
		INSERT INTO roles (name, description)
		VALUES (?, ?)
	`

	result, err := r.db.ExecContext(ctx, query, role.Name(), role.Description())
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
		SELECT id, name, description
		FROM roles
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []*domain.Role{}

	for rows.Next() {
		var id int
		var name, description string

		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		// Create role and set values
		domainRole, err := domain.NewRole(name, description)
		if err != nil {
			return nil, err
		}
		domainRole.SetID(id)

		roles = append(roles, domainRole)
	}

	return roles, nil
}

func (r *SqlRoleRepository) FindByIDs(ctx context.Context, ids []int) ([]*domain.Role, error) {
	if len(ids) == 0 {
		return []*domain.Role{}, nil
	}

	// Create placeholders for IN clause
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		SELECT id, name, description
		FROM roles
		WHERE id IN (` + placeholders[0]
	for i := 1; i < len(placeholders); i++ {
		query += ", " + placeholders[i]
	}
	query += ")"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []*domain.Role{}

	for rows.Next() {
		var id int
		var name, description string

		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		// Create role and set values
		domainRole, err := domain.NewRole(name, description)
		if err != nil {
			return nil, err
		}
		domainRole.SetID(id)

		roles = append(roles, domainRole)
	}

	return roles, nil
}
