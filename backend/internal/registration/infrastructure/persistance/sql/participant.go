package sql

import (
	"context"
	"database/sql"
	"time"

	"hackathon-pvc-backend/internal/registration/domain"
)

type SqlParticipantRepository struct {
	db *sql.DB
}

func NewSqlParticipantRepository(db *sql.DB) domain.ParticipantRepositoryPort {
	return &SqlParticipantRepository{db: db}
}

func (r *SqlParticipantRepository) Save(ctx context.Context, participant *domain.Participant) (*domain.Participant, error) {
	query := `
		INSERT INTO participants (name, nickname, email, region, project_idea, team_preference, desired_teammate, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query,
		participant.Name(),
		participant.Nickname(),
		participant.Email(),
		participant.Region(),
		participant.ProjectIdea(),
		participant.TeamPreference(),
		participant.DesiredTeammate(),
		participant.CreatedAt().Format(time.RFC3339),
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	err = participant.SetID(int(id))
	if err != nil {
		return nil, err
	}

	return participant, nil
}

func (r *SqlParticipantRepository) SetRoles(ctx context.Context, participant *domain.Participant, roles []*domain.Role) error {
	query := `
		INSERT INTO participant_roles (participant_id, role_id)
		VALUES (?, ?)
	`

	for _, role := range roles {
		_, err := r.db.ExecContext(ctx, query, participant.ID(), role.ID())
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SqlParticipantRepository) SetTechnologies(ctx context.Context, participant *domain.Participant, technologies []*domain.Technology) error {
	query := `
		INSERT INTO participant_technologies (participant_id, technology_id)
		VALUES (?, ?)
	`

	for _, technology := range technologies {
		_, err := r.db.ExecContext(ctx, query, participant.ID(), technology.ID())
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SqlParticipantRepository) FindByID(ctx context.Context, id int) (*domain.Participant, error) {
	// TODO: Implement when needed
	return nil, nil
}

func (r *SqlParticipantRepository) FindByNickname(ctx context.Context, nickname string) (*domain.Participant, error) {
	// TODO: Implement when needed
	return nil, nil
}

func (r *SqlParticipantRepository) FindAll(ctx context.Context) ([]*domain.Participant, error) {
	// TODO: Implement when needed
	return nil, nil
}

func (r *SqlParticipantRepository) Count(ctx context.Context) (int, error) {
	// TODO: Implement when needed
	return 0, nil
}
