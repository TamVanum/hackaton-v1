package sql

import (
	"context"
	"database/sql"
	"hackathon-pvc-backend/internal/registration/domain"
)

type SqlTechnologyRepository struct {
	db *sql.DB
}

func NewSqlTechnologyRepository(db *sql.DB) domain.TechnologyRepository {
	return &SqlTechnologyRepository{db: db}
}

func (r *SqlTechnologyRepository) Save(ctx context.Context, technology *domain.Technology) (*domain.Technology, error) {
	query := `
		INSERT INTO technologies (name, description)
		VALUES (?, ?)
	`

	result, err := r.db.ExecContext(ctx, query, technology.Name(), technology.Description())
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
		SELECT id, name, description
		FROM technologies
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	technologies := []*domain.Technology{}

	for rows.Next() {
		var id int
		var name, description string

		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		// Create technology and set values
		domainTechnology, err := domain.NewTechnology(name, description)
		if err != nil {
			return nil, err
		}
		domainTechnology.SetID(id)

		technologies = append(technologies, domainTechnology)
	}

	return technologies, nil
}

func (r *SqlTechnologyRepository) FindByIDs(ctx context.Context, ids []int) ([]*domain.Technology, error) {
	if len(ids) == 0 {
		return []*domain.Technology{}, nil
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
		FROM technologies
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

	technologies := []*domain.Technology{}

	for rows.Next() {
		var id int
		var name, description string

		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		// Create technology and set values
		domainTechnology, err := domain.NewTechnology(name, description)
		if err != nil {
			return nil, err
		}
		domainTechnology.SetID(id)

		technologies = append(technologies, domainTechnology)
	}

	return technologies, nil
}
