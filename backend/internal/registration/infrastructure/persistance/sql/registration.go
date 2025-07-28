package sql

import (
	"context"
	"database/sql"
	"time"

	"hackathon-pvc-backend/internal/registration/domain"
)

type SqlRegistrationRepository struct {
	db *sql.DB
}

func NewSqlRegistrationRepository(db *sql.DB) domain.RegistrationRepositoryPort {
	return &SqlRegistrationRepository{db: db}
}

func (r *SqlRegistrationRepository) Save(ctx context.Context, reg *domain.Registration) (*domain.Registration, error) {
	query := `
		INSERT INTO registrations (name, nickname, email, region, project_idea, team_preference, desired_teammate, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		reg.Name(),
		reg.Nickname(),
		reg.Email(),
		reg.Region(),
		reg.ProjectIdea(),
		reg.TeamPreference(),
		reg.DesiredTeammate(),
		reg.CreatedAt().Format(time.RFC3339),
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
