package sql

import (
	"context"
	"database/sql"
	"hackathon-pvc-backend/internal/registration/domain"
)

type SqlTechnologyRepository struct {
	db *sql.DB
}

func NewSqlTechnologyRepository(db *sql.DB) domain.TechnologyRepositoryPort {
	return &SqlTechnologyRepository{db: db}
}

func (r *SqlTechnologyRepository) Save(ctx context.Context, technology *domain.Technology) (*domain.Technology, error) {
	query := `
		INSERT INTO technologies (name)
		VALUES (?)
	`

	result, err := r.db.ExecContext(ctx, query, technology.Name())
	if err != nil {
		return nil, err
	}

	technologyID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	technology.SetID(int(technologyID))

	return technology, nil
}

func (r *SqlTechnologyRepository) FindAll(ctx context.Context) ([]*domain.Technology, error) {
	query := `
		SELECT id, name
		FROM technologies
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	technologies := []*domain.Technology{}

	for rows.Next() {
		var technology domain.Technology
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		technologies = append(technologies, &technology)
	}

	return technologies, nil
}

func (r *SqlTechnologyRepository) BulkFindByIDs(ctx context.Context, ids []int) ([]*domain.Technology, error) {
	query := `
		SELECT id, name
		FROM technologies
		WHERE id IN (?)
	`

	rows, err := r.db.QueryContext(ctx, query, ids)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	technologies := []*domain.Technology{}

	for rows.Next() {
		var technology domain.Technology
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		technologies = append(technologies, &technology)
	}

	return technologies, nil
}
