package persistance

import (
	"context"
	"database/sql"
	"errors"

	"hackathon-pvc-backend/internal/registration/application/ports"
	"hackathon-pvc-backend/internal/registration/domain"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) ports.RepositoryPort {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Save(ctx context.Context, reg *domain.Registration) (*domain.Registration, error) {
	query := `
		INSERT INTO registrations (name, nickname, project_idea, teammate, role, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	var teammateValue *string
	if reg.Teammate().HasValue() {
		value := reg.Teammate().Value()
		teammateValue = value
	}

	result, err := r.db.ExecContext(ctx, query,
		reg.Name().Value(),
		reg.Nickname().Value(),
		reg.ProjectIdea().Value(),
		teammateValue,
		reg.Role().Value(),
		reg.CreatedAt().Value(),
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	err = reg.SetID(int(id))
	if err != nil {
		return nil, err
	}

	return reg, nil
}

func (r *SQLiteRepository) FindByID(ctx context.Context, id int) (*domain.Registration, error) {
	query := `
		SELECT id, name, nickname, project_idea, teammate, role, created_at
		FROM registrations
		WHERE id = ?
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var dbID int
	var name, nickname, projectIdea, role string
	var teammate *string
	var createdAt string

	err := row.Scan(&dbID, &name, &nickname, &projectIdea, &teammate, &role, &createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("registration not found")
		}
		return nil, err
	}

	reg, err := domain.NewRegistration(name, nickname, projectIdea, teammate, role)
	if err != nil {
		return nil, err
	}

	err = reg.SetID(dbID)
	if err != nil {
		return nil, err
	}

	return reg, nil
}

func (r *SQLiteRepository) FindByNickname(ctx context.Context, nickname string) (*domain.Registration, error) {
	query := `
		SELECT id, name, nickname, project_idea, teammate, role, created_at
		FROM registrations
		WHERE nickname = ?
	`

	row := r.db.QueryRowContext(ctx, query, nickname)

	var dbID int
	var name, dbNickname, projectIdea, role string
	var teammate *string
	var createdAt string

	err := row.Scan(&dbID, &name, &dbNickname, &projectIdea, &teammate, &role, &createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("registration not found")
		}
		return nil, err
	}

	reg, err := domain.NewRegistration(name, dbNickname, projectIdea, teammate, role)
	if err != nil {
		return nil, err
	}

	err = reg.SetID(dbID)
	if err != nil {
		return nil, err
	}

	return reg, nil
}

func (r *SQLiteRepository) FindAll(ctx context.Context) ([]*domain.Registration, error) {
	query := `
		SELECT id, name, nickname, project_idea, teammate, role, created_at
		FROM registrations
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registrations []*domain.Registration

	for rows.Next() {
		var dbID int
		var name, nickname, projectIdea, role string
		var teammate *string
		var createdAt string

		err := rows.Scan(&dbID, &name, &nickname, &projectIdea, &teammate, &role, &createdAt)
		if err != nil {
			return nil, err
		}

		reg, err := domain.NewRegistration(name, nickname, projectIdea, teammate, role)
		if err != nil {
			return nil, err
		}

		err = reg.SetID(dbID)
		if err != nil {
			return nil, err
		}

		registrations = append(registrations, reg)
	}

	return registrations, nil
}

func (r *SQLiteRepository) Count(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM registrations`

	row := r.db.QueryRowContext(ctx, query)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
