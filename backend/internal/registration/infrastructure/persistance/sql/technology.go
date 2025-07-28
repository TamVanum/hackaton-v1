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
