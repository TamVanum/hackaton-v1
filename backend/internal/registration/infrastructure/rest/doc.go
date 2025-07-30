package rest

import "hackathon-pvc-backend/internal/registration/domain"

type CreateRegistrationRequest struct {
	Name            string  `json:"name" binding:"required"`
	Email           string  `json:"email" binding:"required"`
	Nickname        string  `json:"nickname" binding:"required"`
	Region          string  `json:"region" binding:"required"`
	ProjectIdea     string  `json:"project_idea" binding:"required"`
	TeamPreference  bool    `json:"team_preference" binding:"required"`
	DesiredTeammate *string `json:"desired_teammate,omitempty"`
	RoleIDs         []int   `json:"role_ids" binding:"required"`
	TechnologyIDs   []int   `json:"technology_ids" binding:"required"`
}

type CreateRegistrationResponse struct {
	ID              int                  `json:"id"`
	Name            string               `json:"name"`
	Nickname        string               `json:"nickname"`
	Email           string               `json:"email"`
	Region          string               `json:"region"`
	ProjectIdea     string               `json:"project_idea"`
	TeamPreference  bool                 `json:"team_preference"`
	DesiredTeammate *string              `json:"desired_teammate,omitempty"`
	RoleIDs         []*domain.Role       `json:"role_ids"`
	TechnologyIDs   []*domain.Technology `json:"technology_ids"`
	CreatedAt       string               `json:"created_at"`
}

type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateRoleResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateTechnologyRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateTechnologyResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
